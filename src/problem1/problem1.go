package problem1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Run ...
func Run() {
	path := "./input.txt"

	inFile, _ := os.Open(path)
	defer inFile.Close()

	sum := 0

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())

		sum += getTotalFuel(num)
	}

	fmt.Println(sum)
}

func getTotalFuel(m int) int {
	result := getFuel(m)
	sum := result

	for result > 0 {
		result = getFuel(result)
		sum += result
	}

	return sum
}

func getFuel(m int) int {
	result := (m / 3) - 2

	if result < 0 {
		return 0
	}

	return result
}
