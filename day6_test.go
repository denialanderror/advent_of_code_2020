package aoc

import (
	"testing"
)

func TestCountAnswers(t *testing.T) {
	union, intersection := Day6("./input_test/answer_set")
	if union != 11 {
		t.Log("expected ", 11, ", got ", union)
		t.Fail()
	}
	if intersection != 6 {
		t.Log("expected ", 6, ", got ", intersection)
		t.Fail()
	}
}
