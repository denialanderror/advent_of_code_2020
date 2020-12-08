package aoc

import (
	"io/ioutil"
	"strings"
)

// Day6 - https://adventofcode.com/2020/day/6
func Day6(filename string) (int, int) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	answerSets := strings.Split(string(data), "\n\n")

	var union int
	var intersection int

	for _, set := range answerSets {
		answerMap := map[rune]int{}
		answerSet := strings.Split(strings.TrimSpace(set), "\n")

		for _, answers := range answerSet {
			for _, answer := range answers {
				answerMap[answer]++
			}
		}
		union += len(answerMap)

		for _, answer := range answerMap {
			if answer == len(answerSet) {
				intersection++
			}
		}
	}

	return union, intersection
}

func countDistinctAnswers(answerSet []string) int {
	distinctAnswers := map[rune]bool{}

	for _, answers := range answerSet {
		for _, answer := range answers {
			distinctAnswers[answer] = true
		}
	}

	return len(distinctAnswers)
}
