package aoc

import (
	"regexp"
	"strconv"
)

// Day4 - https://adventofcode.com/2020/day/4
func Day4(filename string) (int, int) {
	rawInput := ReadInput(filename)
	passports := buildPassports(rawInput)

	var requiredFieldsPresent int
	var allFieldsValid int

	for _, passport := range passports {
		re := regexp.MustCompile(`(\w{3}):(\S+)`)
		match := re.FindAllStringSubmatch(passport, -1)

		keys := []string{}
		fields := map[string]string{}

		for _, m := range match {
			keys = append(keys, m[1])
			fields[m[1]] = m[2]
		}

		if containsRequiredFields(keys) {
			requiredFieldsPresent++
		}

		if meetsRequirements(fields) {
			allFieldsValid++
		}

	}

	return requiredFieldsPresent, allFieldsValid
}

func buildPassports(input []string) []string {
	passports := []string{}
	var passport string

	for _, p := range input {
		passport += (" " + p)

		if p == "" {
			passports = append(passports, passport)
			passport = ""
		}
	}
	return passports
}

func meetsRequirements(fields map[string]string) bool {
	requirements := map[string]func(string) bool{
		"byr": isWithinRange(1920, 2002),
		"iyr": isWithinRange(2010, 2020),
		"eyr": isWithinRange(2020, 2030),
		"hgt": isValidHeight,
		"hcl": matchesPattern(`^#[a-f0-9]{6}$`),
		"ecl": matchesPattern(`^amb|blu|brn|gry|grn|hzl|oth$`),
		"pid": matchesPattern(`^\d{9}$`),
	}

	isValid := true

	for requiredKey, validator := range requirements {
		contains := false

		for k, v := range fields {
			if requiredKey == k {
				contains = true
				isValid = validator(v)
				break
			}
		}

		if !contains || !isValid {
			isValid = false
			break
		}
	}

	return isValid
}

func isWithinRange(least int, most int) func(string) bool {
	return func(dateString string) bool {
		date, _ := strconv.Atoi(dateString)
		return date >= least && date <= most
	}
}

func isValidHeight(heightString string) bool {
	value := heightString[:len(heightString)-2]
	unit := heightString[len(heightString)-2:]

	var isValid bool

	switch unit {
	case "cm":
		isValid = isWithinRange(150, 193)(value)
	case "in":
		isValid = isWithinRange(59, 76)(value)
	default:
		isValid = false
	}

	return isValid
}

func matchesPattern(pattern string) func(string) bool {
	re := regexp.MustCompile(pattern)
	return func(value string) bool {
		return re.MatchString(value)
	}
}

func containsRequiredFields(fields []string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	isValid := true

	for _, r := range requiredFields {
		contains := false
		for _, f := range fields {
			if r == f {
				contains = true
				break
			}
		}

		if !contains {
			isValid = false
			break
		}
	}

	return isValid
}
