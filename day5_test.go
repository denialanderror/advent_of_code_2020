package aoc

import (
	"testing"
)

func TestFindSeats(t *testing.T) {
	expectations := map[string]BoardingPass{
		"BFFFBBFRRR": {70, 7, 567},
		"FFFBBBFRRR": {14, 7, 119},
		"BBFFBBFRLL": {102, 4, 820},
		"FBFBBFFRLR": {44, 5, 357},
	}

	for code, expected := range expectations {
		response := ParseCodeToBoardingPass(code)
		if response != expected {
			t.Log("expected ", expected, ", got ", response)
			t.Fail()
		}
	}
}
