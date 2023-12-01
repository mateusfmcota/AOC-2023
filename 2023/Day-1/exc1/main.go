package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

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

		first := -1
		last := -1

		for _, char := range text {
			if unicode.IsDigit(char) {
				if first == -1 {
					first = (int)(char - '0')
					last = (int)(char - '0')
				} else {
					last = (int)(char - '0')
				}

			}
		}

		number := first*10 + last

		sum += number
	}

	fmt.Println(sum)
}
