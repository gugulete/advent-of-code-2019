package problem4

import (
	"math"
)

//Run ...
func Run() {
	min := 136888
	max := 685979
	count := 0

	nextDigit(0, 6, 0, 7, 7, min, max, &count)

	println(count)
}

func nextDigit(prevDigit int, position int, prevNumber int, adjStart int, adjEnd int, min int, max int, count *int) {
	if position < 1 {
		if adjStart-adjEnd == 1 {
			*count++
		}

		return
	}

	currentMin := min / int(math.Pow10(position))
	currentMax := max / int(math.Pow10(position))

	if prevNumber < currentMin || prevNumber > currentMax {
		return
	}

	for i := prevDigit; i <= 9; i++ {
		areAdj := i == prevDigit

		if areAdj {
			if adjEnd == position+1 {
				nextDigit(i, position-1, prevNumber*10+i, adjStart, position, min, max, count)
			} else if adjStart-adjEnd == 1 {
				nextDigit(i, position-1, prevNumber*10+i, adjStart, adjEnd, min, max, count)
			} else {
				nextDigit(i, position-1, prevNumber*10+i, position+1, position, min, max, count)
			}
		} else {
			if adjStart-adjEnd == 1 {
				nextDigit(i, position-1, prevNumber*10+i, adjStart, adjEnd, min, max, count)
			} else {
				nextDigit(i, position-1, prevNumber*10+i, position, position, min, max, count)
			}
		}
	}
}
