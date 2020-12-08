package aoc

import (
	"testing"
)

func TestParseValidPassport(t *testing.T) {
	partA, _ := Day4("./input_test/single_line_valid")
	expected := 1
	if partA != expected {
		t.Log("expected ", expected, ", got ", partA)
		t.Fail()
	}
}

func TestParseInvalidPassport(t *testing.T) {
	partA, _ := Day4("./input_test/single_line_invalid")
	expected := 0
	if partA != expected {
		t.Log("expected ", expected, ", got ", partA)
		t.Fail()
	}
}

func TestValidPassports(t *testing.T) {
	_, partB := Day4("./input_test/valid_passports")
	expected := 4
	if partB != expected {
		t.Log("expected ", expected, ", got ", partB)
		t.Fail()
	}
}

func TestInvalidPassports(t *testing.T) {
	_, partB := Day4("./input_test/invalid_passports")
	expected := 0
	if partB != expected {
		t.Log("expected ", expected, ", got ", partB)
		t.Fail()
	}
}
