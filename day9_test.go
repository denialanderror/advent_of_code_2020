package aoc

import (
	"testing"
)

func TestPart2(t *testing.T) {
	a, b := Day9("./input_test/XMAS", 5)
	if a != 127 {
		t.Log("expected ", 127, ", got ", a)
		t.Fail()
	}
	if b != 62 {
		t.Log("expected ", 62, ", got ", b)
		t.Fail()
	}
}
