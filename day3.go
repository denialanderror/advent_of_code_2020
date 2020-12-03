package aoc

import (
	"strings"
)

// Day3 - https://adventofcode.com/2020/day/3
func Day3(filename string) (int, int) {
	rawInput := ReadInput(filename)

	slope := generateSlope(rawInput)

	route1_1 := howManyTrees(slope, 1, 1)
	route3_1 := howManyTrees(slope, 3, 1)
	route5_1 := howManyTrees(slope, 5, 1)
	route7_1 := howManyTrees(slope, 7, 1)
	route1_2 := howManyTrees(slope, 1, 2)

	return route3_1, route1_1 * route3_1 * route5_1 * route7_1 * route1_2
}

func generateSlope(input []string) [][]string {
	slope := [][]string{}

	for _, row := range input {
		if len(row) > 0 {
			slope = append(slope, strings.Split(row, ""))
		}
	}

	return slope
}

func howManyTrees(slope [][]string, right int, down int) int {
	slopeLength := len(slope)
	rowLength := len(slope[0])
	var x, y, trees int

	for y < slopeLength {
		location := slope[y][x]
		if location == "#" {
			trees++
		}
		x = (x + right) % rowLength
		y += down
	}
	return trees
}
