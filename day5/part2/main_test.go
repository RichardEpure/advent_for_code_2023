package main

import (
	"testing"
)

func TestRangeMapGetRange(t *testing.T) {
	range1 := Range{
		value:  0,
		length: 20,
	}
	rangeMap1 := RangeMap{
		destination: 105,
		source:      5,
		length:      10,
	}
	unmapped, mapped := rangeMap1.GetRanges(range1)

	if (Range{}) == mapped || mapped.value != 105 || mapped.length != 10 {
		t.Fatalf("rangeMap failed to map range1. expected=%v, got=%v",
			Range{value: 105, length: 10}, mapped)
	}

	if len(unmapped) != 2 {
		t.Fatalf("unexpected number of unmapped Ranges. expected=%v, got=%v", 2, len(unmapped))
	}

	if unmapped[0].value != 0 || unmapped[0].length != 5 {
		t.Fatalf("unexpected unmapped range. expected=%v, got=%v",
			Range{value: 0, length: 5}, unmapped[0])
	}

	if unmapped[1].value != 15 || unmapped[1].length != 5 {
		t.Fatalf("unexpected unmapped range. expected=%v, got=%v",
			Range{value: 15, length: 5}, unmapped[1])
	}

	range2 := Range{
		value:  0,
		length: 5,
	}
	rangeMap2 := RangeMap{
		destination: 100,
		source:      0,
		length:      5,
	}
	unmapped, mapped = rangeMap2.GetRanges(range2)

	if (Range{}) == mapped || mapped.value != 100 || mapped.length != 5 {
		t.Fatalf("rangeMap failed to map range2. expected=%v, got=%v",
			Range{value: 100, length: 5}, mapped)
	}

	if len(unmapped) != 0 {
		t.Fatalf("unexpected number of unmapped Ranges. expected=%v, got=%v", 0, len(unmapped))
	}
}
