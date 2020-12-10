package aoc

import (
	"testing"
)

func TestBagStuff(t *testing.T) {
	a, b := Day7("./input_test/bags")
	if a != 0 {
		t.Log("expected ", 0, ", got ", a)
		t.Fail()
	}
	if b != 126 {
		t.Log("expected ", 126, ", got ", b)
		t.Fail()
	}
}
