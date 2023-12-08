package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("day1/input.txt")
	check(err)

	var lines []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	var total int = 0

	for _, line := range lines {
		var first rune
		var last rune

		for _, char := range line {
			if unicode.IsDigit(char) {
				if first == 0 {
					first = char
					last = char
				} else {
					last = char
				}
			}
		}

		value, err := strconv.Atoi(string(first) + string(last))
		check(err)

		total += value
	}

	fmt.Println("total: ", total)
}
