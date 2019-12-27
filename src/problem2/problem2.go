package problem2

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"utils/arrays"
)

type input struct {
	noun int
	verb int
}

//Run ...
func Run() {
	text := readInput()
	stringArr := strings.Split(text, ",")
	numArr := arrays.StringToIntArr(&stringArr)

	inputs := []input{
		{noun: 54, verb: 85},
	}

	for _, inp := range inputs {
		fmt.Println("res: ", inp.noun, inp.verb, process(inp.noun, inp.verb, numArr))
	}
}

func process(noun int, verb int, arrInput []int) int {
	arr := make([]int, len(arrInput))
	for i, v := range arrInput {
		arr[i] = v
	}

	arr[1] = noun
	arr[2] = verb

	i := 0
	for i < len(arr) {
		if arr[i] == 99 {
			break
		}

		switch arr[i] {
		case 1:
			{
				arr[getPosition(i+3, arr)] = getValue(i+1, arr) + getValue(i+2, arr)
			}
		case 2:
			{
				arr[getPosition(i+3, arr)] = getValue(i+1, arr) * getValue(i+2, arr)
			}
		default:
			panic(fmt.Errorf("unexpected command: %d", arr[i]))
		}

		i += 4
	}

	return arr[0]
}

func getValue(i int, arr []int) int {
	return arr[getPosition(i, arr)]
}

func getPosition(i int, arr []int) int {
	return arr[i]
}

func readInput() string {
	const inputPath = "problem2/input.txt"

	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, inputPath)
	dat, _ := ioutil.ReadFile(path)

	return string(dat)
}
