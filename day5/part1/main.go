package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type RangeMap struct {
	destination int
	source      int
	length      int
}

func (rm *RangeMap) Get(val int) int {
	if val >= rm.source && val < rm.source+rm.length {
		return (val - rm.source) + rm.destination
	}
	return val
}

type Map struct {
	rangeMaps []RangeMap
}

func (m *Map) Get(val int) int {
	newVal := val
	for _, rm := range m.rangeMaps {
		newVal = rm.Get(val)
		if newVal != val {
			break
		}
	}
	return newVal
}

func (m *Map) AddRange(rm RangeMap) {
	m.rangeMaps = append(m.rangeMaps, rm)
}

func main() {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		panic(err)
	}

	var lines []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	seeds := strings.Split(lines[0], " ")[1:]
	lines = lines[1:]

	maps := []Map{}

	for _, line := range lines {
		if len(line) == 0 {
			maps = append(maps, Map{rangeMaps: []RangeMap{}})
			continue
		}

		split := strings.Split(line, " ")
		if len(split) != 3 {
			continue
		}

		destination, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		source, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		length, err := strconv.Atoi(split[2])
		if err != nil {
			panic(err)
		}

		maps[len(maps)-1].AddRange(RangeMap{destination, source, length})
	}

	locations := []int{}
	smallestLocation := 0
	for _, seed := range seeds {
		value, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}

		for _, m := range maps {
			value = m.Get(value)
		}

		if smallestLocation == 0 || value < smallestLocation {
			smallestLocation = value
		}

		locations = append(locations, value)
	}
	println("smallest location:", smallestLocation)
}
