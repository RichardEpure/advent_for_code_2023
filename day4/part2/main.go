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

func findMatches(cards []string, i int, memo map[int]int) (matches int) {
	matches, ok := memo[i]

	if ok {
		return matches
	}

	localMatches := 0
	split := strings.Split(cards[i], ":")[1]
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
			localMatches += 1
		}
	}
	memo[i] = localMatches

	for j := 1; j <= localMatches; j++ {
		memo[i] += findMatches(cards, i+j, memo)
	}

	return memo[i]
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

	sum := len(lines)
	memo := make(map[int]int)

	for i := range lines {
		sum += findMatches(lines, i, memo)
	}

	println("total scratchcards: ", sum)
}
