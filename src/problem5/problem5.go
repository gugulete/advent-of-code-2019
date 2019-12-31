package problem5

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type valueMode int

const (
	pos valueMode = 0
	imm valueMode = 1
)

type param struct {
	index int
	mode  valueMode
}

// Run ...
func Run() {
	text := readInput()
	stringArr := strings.Split(text, ",")

	numArr := make([]int, len(stringArr))
	for i, v := range stringArr {
		num, _ := strconv.Atoi(v)
		numArr[i] = num
	}

	input := 5
	diagnose(input, numArr)
}

func diagnose(input int, arrInput []int) {
	arr := make([]int, len(arrInput))
	for i, v := range arrInput {
		arr[i] = v
	}

	i := 0
	for i >= 0 {
		i = handleInstruction(i, input, arr)
	}
}

func handleInstruction(index int, input int, arr []int) int {
	instriction := arr[index]

	code := getInstructionCode(instriction)
	switch code {
	case 1:
		add(
			valueIndex{mode: getParamMode(instriction, 1), index: index + 1},
			valueIndex{mode: getParamMode(instriction, 2), index: index + 2},
			valueIndex{mode: getParamMode(instriction, 3), index: index + 3},
			arr,
		)
		return index + 4
	case 2:
		multiply(
			valueIndex{mode: getParamMode(instriction, 1), index: index + 1},
			valueIndex{mode: getParamMode(instriction, 2), index: index + 2},
			valueIndex{mode: getParamMode(instriction, 3), index: index + 3},
			arr,
		)
		return index + 4
	case 3:
		saveInput(
			input,
			valueIndex{mode: getParamMode(instriction, 1), index: index + 1},
			arr,
		)
		return index + 2
	case 4:
		fmt.Printf("result: %d\n", getValue(
			valueIndex{mode: getParamMode(instriction, 1), index: index + 1},
			arr,
		))
		return index + 2
	case 5:
		return jumpIfTrue(
			valueIndex{mode: getParamMode(instriction, 1), index: index + 1},
			valueIndex{mode: getParamMode(instriction, 2), index: index + 2},
			index+3,
			arr,
		)
	case 6:
		return jumpIfFalse(
			valueIndex{mode: getParamMode(instriction, 1), index: index + 1},
			valueIndex{mode: getParamMode(instriction, 2), index: index + 2},
			index+3,
			arr,
		)
	case 7:
		lessThan(
			valueIndex{mode: getParamMode(instriction, 1), index: index + 1},
			valueIndex{mode: getParamMode(instriction, 2), index: index + 2},
			valueIndex{mode: getParamMode(instriction, 3), index: index + 3},
			arr,
		)

		return index + 4
	case 8:
		equal(
			valueIndex{mode: getParamMode(instriction, 1), index: index + 1},
			valueIndex{mode: getParamMode(instriction, 2), index: index + 2},
			valueIndex{mode: getParamMode(instriction, 3), index: index + 3},
			arr,
		)

		return index + 4
	case 99:
		return -1
	default:
		panic(fmt.Errorf("unexpected code: %d", code))
	}
}

func getInstructionCode(instr int) int {
	return instr % 100
}

func getParamMode(instr int, index int) valueMode {
	mode := (instr / int(math.Pow10(index+1))) % 10

	switch mode {
	case 0:
		return pos
	case 1:
		return imm
	default:
		panic(fmt.Errorf("unexpected mode: %d", mode))
	}
}

func equal(a valueIndex, b valueIndex, c valueIndex, arr []int) {
	if getValue(a, arr) == getValue(b, arr) {
		writeValue(1, c, arr)
	} else {
		writeValue(0, c, arr)
	}
}

func lessThan(a valueIndex, b valueIndex, c valueIndex, arr []int) {
	if getValue(a, arr) < getValue(b, arr) {
		writeValue(1, c, arr)
	} else {
		writeValue(0, c, arr)
	}
}

func jumpIfFalse(condition valueIndex, position valueIndex, defaultIndex int, arr []int) int {
	if getValue(condition, arr) == 0 {
		return getValue(position, arr)
	}

	return defaultIndex
}

func jumpIfTrue(condition valueIndex, position valueIndex, defaultIndex int, arr []int) int {
	if getValue(condition, arr) != 0 {
		return getValue(position, arr)
	}

	return defaultIndex
}

type valueIndex struct {
	mode  valueMode
	index int
}

func add(a valueIndex, b valueIndex, result valueIndex, arr []int) {
	writeValue(
		getValue(a, arr)+getValue(b, arr),
		result,
		arr,
	)
}

func multiply(a valueIndex, b valueIndex, result valueIndex, arr []int) {
	writeValue(
		getValue(a, arr)*getValue(b, arr),
		result,
		arr,
	)
}

func saveInput(input int, index valueIndex, arr []int) {
	writeValue(
		input,
		index,
		arr,
	)
}

func writeValue(value int, index valueIndex, arr []int) {
	switch index.mode {
	case pos:
		arr[arr[index.index]] = value
	case imm:
		arr[index.index] = value
	default:
		panic(fmt.Errorf("unexpected mode: %d", index.mode))
	}
}

func getValue(a valueIndex, arr []int) int {
	switch a.mode {
	case pos:
		return arr[arr[a.index]]
	case imm:
		return arr[a.index]
	default:
		panic(fmt.Errorf("unexpected mode: %d", a.mode))
	}
}

func readInput() string {
	path := "./problem5/input.txt"
	dat, _ := ioutil.ReadFile(path)
	return string(dat)
}
