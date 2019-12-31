package problem7

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"utils/arrays"
)

type valueMode int
type seq []int

const debug bool = false

const (
	pos valueMode = 0
	imm valueMode = 1
)

// Run ...
func Run() {
	text := readInput()
	stringArr := strings.Split(text, ",")

	numArr := arrays.StringToIntArr(&stringArr)

	// print(maxSignal(numArr, execSeq, []int{0, 1, 2, 3, 4}))
	print(maxSignal(&numArr, processSequenceWithFeedback, []int{5, 6, 7, 8, 9}))
}

func maxSignal(data *[]int, exec func(*[]int, seq, int) int, phases []int) int {
	max := 0
	for index, seq := range permutations(phases) {
		if debug {
			arr := seqToArray(seq)
			println("seq:", strings.Join(arrays.IntToStringArr(&arr), ","))
		}

		output := exec(data, seq, index)
		if debug {
			println("output:", output)
		}

		if output > max {
			max = output
		}
	}

	return max
}

func processSequenceWithFeedback(data *[]int, phases seq, seqID int) int {
	in := make(chan int, 1)
	in <- 0

	// pipe the amplifiers
	out := in
	for _, phase := range phases {
		out = processPhase(arrays.CopyIntArr(data), phase, out, seqID)
	}

	// pipe the output of the last amplifier back into the input of the first one
	var output int
	for output = range out {
		in <- output
	}

	// last output ever is the output signal
	return output
}

func permutations(input []int) []seq {
	allInserts := func(current []int, ins int) []seq {
		result := []seq{}

		for i := 0; i <= len(current); i++ {
			result = append(result, append(append(append([]int{}, current[:i]...), ins), current[i:]...))
		}

		return result
	}

	var generate func(inp []int, perms []seq) []seq
	generate = func(inp []int, perms []seq) []seq {
		if len(inp) == 0 {
			return perms
		}

		result := []seq{}
		for _, e := range perms {
			result = append(result, allInserts(e, inp[0])...)
		}

		return generate(inp[1:], result)
	}

	return generate(input[1:], []seq{arrayToSeq([]int{input[0]})})
}

func arrayToSeq(arr []int) seq {
	return arr
}

func seqToArray(s seq) []int {
	return s
}

func processPhase(data []int, phase int, input <-chan int, seqID int) chan int {
	in, out := make(chan int, 1), make(chan int)
	go func() {
		in <- phase
		for signal := range input {
			in <- signal
		}
		close(in)
	}()

	go executeCode(&data, in, out, phase, seqID)
	return out
}

func executeCode(data *[]int, in <-chan int, out chan<- int, phase int, seqID int) {
	cursor := 0

	printDebugMessage := func(cursor int, length int, command string) {
		if debug {
			slice := (*data)[cursor : cursor+length]
			println(strconv.Itoa(seqID)+"["+strconv.Itoa(phase)+"]", command, "\t>>", strings.Join(arrays.IntToStringArr(&slice), ","))
		}
	}

	for cursor < len(*data) {
		instriction := (*data)[cursor]
		code := getInstructionCode(instriction)

		switch code {
		case 1:
			//add
			printDebugMessage(cursor, 4, "add")

			value1 := getValue(cursor+1, getParamMode(instriction, 1), data)
			value2 := getValue(cursor+2, getParamMode(instriction, 2), data)

			writeValue(
				value1+value2,
				cursor+3,
				getParamMode(instriction, 3),
				data,
			)

			cursor += 4
		case 2:
			// multiply
			printDebugMessage(cursor, 4, "mtply")

			value1 := getValue(cursor+1, getParamMode(instriction, 1), data)
			value2 := getValue(cursor+2, getParamMode(instriction, 2), data)

			writeValue(
				value1*value2,
				cursor+3,
				getParamMode(instriction, 3),
				data,
			)

			cursor += 4
		case 3:
			// input
			printDebugMessage(cursor, 2, "input")

			writeValue(
				<-in,
				cursor+1,
				getParamMode(instriction, 1),
				data,
			)

			cursor += 2
		case 4:
			// output
			printDebugMessage(cursor, 2, "output")

			out <- getValue(
				cursor+1,
				getParamMode(instriction, 1),
				data,
			)
			cursor += 2
		case 5:
			// jump if true
			printDebugMessage(cursor, 3, "jumpT")

			value := getValue(cursor+1, getParamMode(instriction, 1), data)

			if value != 0 {
				cursor = getValue(cursor+2, getParamMode(instriction, 2), data)
			} else {
				cursor += 3
			}
		case 6:
			// jump if false
			printDebugMessage(cursor, 3, "jumpF")

			value := getValue(cursor+1, getParamMode(instriction, 1), data)

			if value == 0 {
				cursor = getValue(cursor+2, getParamMode(instriction, 2), data)
			} else {
				cursor += 3
			}
		case 7:
			// less than
			printDebugMessage(cursor, 4, "lt")

			value1 := getValue(cursor+1, getParamMode(instriction, 1), data)
			value2 := getValue(cursor+2, getParamMode(instriction, 2), data)

			value := 0
			if value1 < value2 {
				value = 1
			}

			writeValue(value, cursor+3, getParamMode(instriction, 3), data)
			cursor += 4
		case 8:
			// equal
			printDebugMessage(cursor, 4, "eq")

			value1 := getValue(cursor+1, getParamMode(instriction, 1), data)
			value2 := getValue(cursor+2, getParamMode(instriction, 2), data)

			value := 0
			if value1 == value2 {
				value = 1
			}

			writeValue(value, cursor+3, getParamMode(instriction, 3), data)
			cursor += 4
		case 99:
			printDebugMessage(cursor, 1, "end")

			close(out)
			return
		default:
			panic(fmt.Errorf("unexpected code: %d", code))
		}
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

func writeValue(value int, index int, mode valueMode, data *[]int) {
	switch mode {
	case pos:
		(*data)[(*data)[index]] = value
	case imm:
		(*data)[index] = value
	default:
		panic(fmt.Errorf("unexpected mode: %d", mode))
	}
}

func getValue(index int, mode valueMode, data *[]int) int {
	switch mode {
	case pos:
		return (*data)[(*data)[index]]
	case imm:
		return (*data)[index]
	default:
		panic(fmt.Errorf("unexpected mode: %d", mode))
	}
}

func readInput() string {
	path := "./problem7/input.txt"
	dat, _ := ioutil.ReadFile(path)
	return string(dat)
}
