package aoc

import (
	"sort"
	"strconv"
)

// Day9 - https://adventofcode.com/2020/day/9
func Day9(filename string, preambleLength int) (int, int) {
	rawInput := ReadInput(filename)

	XMAS := []int{}
	for _, input := range rawInput {
		if len(input) > 0 {
			num, _ := strconv.Atoi(input)
			XMAS = append(XMAS, num)
		}
	}

	firstInvalid := -1

	for i := preambleLength; i < len(XMAS); i++ {
		valid := isValid(XMAS[i], XMAS[i-preambleLength:i])

		if !valid {
			firstInvalid = XMAS[i]
			break
		}
	}

	return firstInvalid, findWeakness(firstInvalid, XMAS)
}

func isValid(target int, preamble []int) bool {
	var valid bool

	start := 0
	end := len(preamble)

	for start < end {
		for i := start + 1; i < end; i++ {
			if preamble[start] == preamble[i] {
				continue
			}
			match := preamble[start] + preamble[i]
			if match == target {
				valid = true
				break
			}
		}
		if valid {
			break
		}
		start++
	}

	return valid
}

func findWeakness(target int, XMAS []int) int {
	contiguous := []int{}
	weakness := -1

	for i := 0; i < len(XMAS); i++ {
		for j := i; i < len(XMAS); j++ {
			contiguous = append(contiguous, XMAS[j])
			all := addAll(contiguous)

			if all == target && len(contiguous) > 1 {
				found := append([]int{}, contiguous...)
				sort.Ints(found)
				weakness = found[0] + found[len(found)-1]
				break
			}
			if all > target {
				contiguous = []int{}
				i++
				break
			}
		}
		if weakness > 0 {
			break
		}
	}

	return weakness
}

func addAll(nums []int) int {
	var all int

	for _, n := range nums {
		all += n
	}

	return all
}
