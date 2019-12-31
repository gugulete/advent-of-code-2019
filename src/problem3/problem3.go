package problem3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type input struct {
	direction string
	positions int
}

// Run ...
func Run() {
	a, b := readInput()

	aPath := make(map[string]int)

	x, y, step := 0, 0, 0

	for _, v := range a {
		for i := 0; i < v.positions; i++ {
			switch v.direction {
			case "L":
				x--
				step++
			case "R":
				x++
				step++
			case "U":
				y++
				step++
			case "D":
				y--
				step++
			}

			addPoint(x, y, step, aPath)
		}
	}

	const MaxUint = ^uint(0)
	dist := int(MaxUint >> 1)
	x, y, step = 0, 0, 0

	for _, v := range b {
		for i := 0; i < v.positions; i++ {
			switch v.direction {
			case "L":
				x--
				step++
			case "R":
				x++
				step++
			case "U":
				y++
				step++
			case "D":
				y--
				step++
			}

			dist = minDist(x, y, step, aPath, dist)
		}
	}

	println(dist)
}

func minDist(x int, y int, step int, path map[string]int, dist int) int {
	val, ok := getPointValue(x, y, path)
	if ok {
		d := val + step
		if d < dist {
			return d
		}
	}

	return dist
}

func addPoint(x int, y int, step int, path map[string]int) {
	path[fmt.Sprintf("%d:%d", x, y)] = step
}

func getPointValue(x int, y int, path map[string]int) (int, bool) {
	val, ok := path[fmt.Sprintf("%d:%d", x, y)]
	return val, ok
}

func readInput() ([]input, []input) {
	path := "./problem3/input.txt"

	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	x := toInput(scanner.Text())
	scanner.Scan()
	y := toInput(scanner.Text())

	return x, y
}

func toInput(str string) []input {
	arr := strings.Split(str, ",")

	res := make([]input, len(arr))
	for i, v := range arr {
		direction := v[0:1]
		positions, _ := strconv.Atoi(v[1:])

		res[i] = input{
			direction: direction,
			positions: positions,
		}
	}

	return res
}
