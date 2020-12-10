package aoc

import (
	"regexp"
	"strconv"
)

type bagNode struct {
	name   string
	volume int
}

type bagMap = map[string][]bagNode

// Day7 - https://adventofcode.com/2020/day/7
func Day7(filename string) (int, int) {
	rawInput := ReadInput(filename)
	target := "shiny gold"
	bagMap := buildBagMap(rawInput)

	holdingGold := bagsThatHold(target, bagMap)

	uniqueGoldHolders := map[string]bool{}
	for _, bag := range holdingGold {
		uniqueGoldHolders[bag] = true
	}

	return len(holdingGold), requiredWithin(target, bagMap) - 1
}

func bagsThatHold(bagName string, bagMap bagMap) []string {
	containing := []string{}

	for key, bag := range bagMap {
		for _, child := range bag {
			if child.name == bagName {
				containing = append(containing, key)
				containing = append(containing, bagsThatHold(key, bagMap)...)
				break
			}
		}
	}

	return containing
}

func requiredWithin(bagName string, bagMap bagMap) int {
	requiredCount := 1

	bags := bagMap[bagName]

	for _, child := range bags {
		if child.volume > 0 {
			requiredCount += child.volume * requiredWithin(child.name, bagMap)
		}
	}

	return requiredCount
}

func buildBagMap(rows []string) bagMap {
	bagMap := map[string][]bagNode{}

	for _, row := range rows {
		pattern := regexp.MustCompile(`((\d?)\s?([a-z]+\s[a-z]+))`)
		match := pattern.FindAllStringSubmatch(row, -1)

		bags := []bagNode{}
		if len(match) > 0 {
			key := match[0][0]

			for _, bag := range match[2:] {
				if bag[0] == "no other" {
					break
				}

				name := bag[3]
				volume, _ := strconv.Atoi(bag[2])
				bags = append(bags, bagNode{name, volume})
			}

			bagMap[key] = bags
		}
	}

	return bagMap
}
