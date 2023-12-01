package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

var strDigits = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var intDigits = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func getNumbers(text string) (int, int) {

	first := -1
	firstPos := math.MaxInt
	last := -1
	lastPos := 0

	for index, char := range text {
		if unicode.IsDigit(char) {
			if first == -1 {
				first = (int)(char - '0')
				firstPos = index
				last = (int)(char - '0')
				lastPos = index
			} else {
				last = (int)(char - '0')
				lastPos = index
			}
		}
	}

	for index, num := range strDigits {
		i := strings.Index(text, num)
		j := strings.LastIndex(text, num)
		if i < firstPos && i != -1 {
			first = intDigits[index]
			firstPos = i
		}
		if j > lastPos && j != -1 {
			last = intDigits[index]
			lastPos = j
		}
	}

	// fmt.Println()
	// fmt.Println("first: ", first, "last:", last)

	// fmt.Printf("%d %d %d %d\n", firstPos, first, lastPos, last)

	return first, last

}

func main() {

	input_file, err := os.Open("input")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input_file)
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		first, last := getNumbers(text)

		number := first*10 + last

		sum += number
	}

	fmt.Println(sum)
}
