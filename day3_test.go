package aoc

import (
	"testing"
)

func TestSmallRun(t *testing.T) {
	partA, _ := Day3("./input_test/small_run")
	expected := 3
	if partA != expected {
		t.Fail()
	}
}

func TestOverflow(t *testing.T) {
	partA, _ := Day3("./input_test/overflow")
	expected := 6
	if partA != expected {
		t.Log("expected ", expected, ", got ", partA)
		t.Fail()
	}
}
