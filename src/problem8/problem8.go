package problem8

import (
	"fmt"
	"io/ioutil"
	"strings"
	"utils/arrays"
)

// Run ...
func Run() {
	text := readInput()
	stringArr := strings.Split(text, "")
	numArr := arrays.StringToIntArr(&stringArr)

	width := 25
	height := 6

	// println(validateData(&numArr, width, height))

	image := renderImage(&numArr, width, height)
	printImage((&image))
}

func printImage(image *[][]int) {
	for _, row := range *image {
		for _, col := range row {
			fmt.Printf("%c", rune(col))
		}
		print("\n")
	}
}

type layer struct {
	zeros int
	ones  int
	twos  int
}

func renderImage(data *[]int, width int, height int) [][]int {
	image := make([][]int, height)

	col, row := 0, 0

	for _, v := range *data {
		if col == width {
			col = 0
			row++
		}

		if row == height {
			row = 0
		}

		if image[row] == nil {
			image[row] = make([]int, width)
		}

		if image[row][col] == 0 {
			switch v {
			case 0:
				image[row][col] = ' '
			case 1:
				image[row][col] = '0'
			}
		}

		col++
	}

	return image
}

func validateData(data *[]int, width int, height int) int {
	targetLayer := layer{zeros: width * height}
	currentLayer := layer{}

	col, row := 0, 0

	for _, v := range *data {
		if col == width {
			col = 0
			row++
		}

		if row == height {
			row = 0
			if currentLayer.zeros < targetLayer.zeros {
				targetLayer = currentLayer
			}

			currentLayer = layer{}
		}

		switch v {
		case 0:
			currentLayer.zeros++
		case 1:
			currentLayer.ones++
		case 2:
			currentLayer.twos++
		}

		col++
	}

	if currentLayer.zeros < targetLayer.zeros {
		targetLayer = currentLayer
	}

	return targetLayer.ones * targetLayer.twos
}

func readInput() string {
	path := "./problem8/input.txt"
	dat, _ := ioutil.ReadFile(path)
	return string(dat)
}
