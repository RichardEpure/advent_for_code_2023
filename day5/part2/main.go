package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	value  int
	length int
}

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

func (rm *RangeMap) GetRanges(r Range) (unmapped []Range, mapped Range) {
	if r.value+r.length <= rm.source || r.value >= rm.source+rm.length {
		return []Range{r}, Range{}
	}
	start := max(rm.source, r.value)
	end := min(rm.source+rm.length, r.value+r.length) - 1
	mapped = Range{rm.Get(start), rm.Get(end) - rm.Get(start) + 1}

	if r.value < start {
		unmapped = append(unmapped, Range{r.value, start - r.value})
	}
	if r.value+r.length-1 > end {
		unmapped = append(unmapped, Range{end + 1, (r.value + r.length) - end})
	}
	return unmapped, mapped
}

type Map struct {
	rangeMaps []RangeMap
}

func (m *Map) Get(val int) int {
	newVal := val
	for _, r := range m.rangeMaps {
		newVal = r.Get(val)
		if newVal != val {
			break
		}
	}
	return newVal
}

func (m *Map) GetRanges(r Range) []Range {
	unmappedRanges := []Range{r}
	mappedRanges := []Range{}
	for _, rm := range m.rangeMaps {
		newUnmappedRanges := []Range{}
		for _, r := range unmappedRanges {
			unmapped, mapped := rm.GetRanges(r)
			if (Range{}) != mapped {
				mappedRanges = append(mappedRanges, mapped)
			}
			newUnmappedRanges = append(newUnmappedRanges, unmapped...)
		}
		unmappedRanges = newUnmappedRanges
	}

	ranges := []Range{}
	ranges = append(ranges, mappedRanges...)
	ranges = append(ranges, unmappedRanges...)
	return ranges
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

	seedData := strings.Split(lines[0], " ")[1:]
	seeds := []Range{}
	for i := 0; i < len(seedData); i += 2 {
		value, err := strconv.Atoi(seedData[i])
		if err != nil {
			panic(err)
		}
		length, err := strconv.Atoi(seedData[i+1])
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, Range{value, length})
	}
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

	locations := []Range{}
	smallestLocation := 0
	for _, seed := range seeds {
		locationsForSeed := []Range{seed}
		for _, m := range maps {
			newLocations := []Range{}
			for _, l := range locationsForSeed {
				newLocations = append(newLocations, m.GetRanges(l)...)
			}
			locationsForSeed = newLocations
		}
		locations = append(locations, locationsForSeed...)
	}

	for _, l := range locations {
		if smallestLocation == 0 || l.value < smallestLocation {
			smallestLocation = l.value
		}
	}
	println("smallest location:", smallestLocation)
}
