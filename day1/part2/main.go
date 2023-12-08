package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	var set map[string]string = make(map[string]string)
	set["one"] = "1"
	set["two"] = "2"
	set["three"] = "3"
	set["four"] = "4"
	set["five"] = "5"
	set["six"] = "6"
	set["seven"] = "7"
	set["eight"] = "8"
	set["nine"] = "9"

	var total int = 0

	for _, line := range lines {
		var first rune
		var last rune
		var firstReplace string
		var lastReplace string
		var firstIndex int = -1
		var lastIndex int = -1

		for key := range set {
			var fi int = strings.Index(line, key)
			var li int = strings.LastIndex(line, key)
			var indices []int = []int{fi, li}

			for _, index := range indices {
				if index == -1 {
					continue
				}

				if firstIndex == -1 {
					firstIndex = index
					firstReplace = key
				}

				if index < firstIndex {
					firstIndex = index
					firstReplace = key
				}

				if index > lastIndex {
					lastIndex = index
					lastReplace = key
				}
			}
		}

		if lastIndex != -1 && lastIndex != firstIndex {
			line = line[:lastIndex+1] + set[lastReplace] + line[lastIndex+1:]
		}

		if firstIndex != -1 {
			line = line[:firstIndex+1] + set[firstReplace] + line[firstIndex+1:]
		}

		for _, char := range line {
			if unicode.IsDigit(char) {
				if first == 0 {
					first, last = char, char
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
