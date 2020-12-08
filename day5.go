package aoc

import (
	"math"
	"sort"
)

type BoardingPass struct {
	row    int
	column int
	seatID int
}

type bySeatID []BoardingPass

func (a bySeatID) Len() int           { return len(a) }
func (a bySeatID) Less(i, j int) bool { return a[i].seatID > a[j].seatID }
func (a bySeatID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Day5 - https://adventofcode.com/2020/day/5
func Day5(filename string) (int, int) {
	rawInput := ReadInput(filename)

	passes := []BoardingPass{}

	for _, row := range rawInput {
		passes = append(passes, ParseCodeToBoardingPass(row))
	}

	sort.Sort(bySeatID(passes))

	return passes[0].seatID, findMissingSeat(passes)
}

func ParseCodeToBoardingPass(code string) BoardingPass {
	pass := BoardingPass{-1, -1, -1}
	min := 0
	max := 127
	mid := func() float64 {
		return float64(max-min) / 2
	}

	for _, c := range code {
		codePoint := string(c)
		switch codePoint {
		case "F":
			max -= int(math.Ceil(mid()))
		case "B":
			min += int(mid())
		case "L":
			{
				if pass.row == -1 {
					pass.row = max
					min = 0
					max = 7
				}
				max -= int(math.Ceil(mid()))
			}
		case "R":
			{
				if pass.row == -1 {
					pass.row = max
					min = 0
					max = 7
				}
				min += int(mid())
			}
		default:
			break
		}
	}

	pass.column = max
	pass.seatID = pass.row*8 + pass.column

	return pass
}

func findMissingSeat(sortedPasses []BoardingPass) int {
	var missingSeatID int
	for i := 0; i < len(sortedPasses)-1; i++ {
		if sortedPasses[i].seatID-sortedPasses[i+1].seatID > 1 {
			missingSeatID = sortedPasses[i].seatID - 1
			break
		}
	}
	return missingSeatID
}
