package aoc

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

// Day1 - https://adventofcode.com/2020/day/1
func Day1(filename string) (int, int) {
	rawInput := ReadInput(filename)

	input := []int{}

	for _, s := range rawInput {
		v, _ := strconv.Atoi(s)
		if v > 0 && v < 2019 {
			input = append(input, v)
		}
	}

	sort.Ints(input)

	twoMatch, threeMatch := findMissingExpenses(input, 2020)

	log.Println("The two matching values total: ", twoMatch)
	log.Println("The three matching values total: ", threeMatch)

	return twoMatch, threeMatch
}

func findMissingExpenses(input []int, target int) (int, int) {
	twoMatch := -1
	threeMatch := -1
	for _, i := range input {
		if twoMatch > -1 && threeMatch > -1 {
			break
		}
		for _, j := range input {
			if twoMatch > -1 && threeMatch > -1 {
				break
			}
			if i+j == target {
				twoMatch = i * j
			}
			for _, k := range input {
				if threeMatch > -1 {
					break
				}
				if i+j+k == target {
					fmt.Println(i, j, k)
					threeMatch = i * j * k
				}
			}
		}
	}

	return twoMatch, threeMatch
}
