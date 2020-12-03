package aoc

import (
	"regexp"
	"strconv"
)

type passwordPolicy struct {
	min      int
	max      int
	letter   string
	password string
}

// Day2 - https://adventofcode.com/2020/day/2
func Day2(filename string) (int, int) {
	rawInput := ReadInput(filename)

	input := []passwordPolicy{}

	for _, s := range rawInput {
		re := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
		match := re.FindStringSubmatch(s)

		if len(match) > 0 {
			min, _ := strconv.Atoi(match[1])
			max, _ := strconv.Atoi(match[2])
			letter := match[3]
			password := match[4]

			policy := passwordPolicy{min, max, letter, password}
			input = append(input, policy)
		}
	}
	return countValidPasswords1(input), countValidPasswords2(input)
}

func countValidPasswords1(policies []passwordPolicy) int {
	count := 0
	for _, policy := range policies {
		if passwordIsValid1(policy) {
			count++
		}
	}
	return count
}

func passwordIsValid1(policy passwordPolicy) bool {
	re := regexp.MustCompile(policy.letter)
	match := re.FindAllString(policy.password, -1)
	count := len(match)
	return count >= policy.min && count <= policy.max
}

func countValidPasswords2(policies []passwordPolicy) int {
	count := 0
	for _, policy := range policies {
		if passwordIsValid2(policy) {
			count++
		}
	}
	return count
}

func passwordIsValid2(policy passwordPolicy) bool {
	firstMatches := string(policy.password[policy.min-1]) == policy.letter
	secondMatches := string(policy.password[policy.max-1]) == policy.letter

	return firstMatches != secondMatches
}
