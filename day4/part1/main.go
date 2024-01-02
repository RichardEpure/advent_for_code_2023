package main

import (
	"bufio"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("day4/input.txt")
	check(err)

	var lines []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	sum := 0

	for _, card := range lines {
		points := 0
		split := strings.Split(card, ":")[1]
		numbers := strings.Split(split, "|")
		winningNumbers := strings.Fields(numbers[0])
		chosenNumbers := strings.Fields(numbers[1])
		winningMap := make(map[string]bool)

		for _, winningNumber := range winningNumbers {
			winningMap[winningNumber] = true
		}

		for _, chosenNumber := range chosenNumbers {
			_, ok := winningMap[chosenNumber]
			if ok {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}

		sum += points
	}

	println("sum of points:", sum)
}
