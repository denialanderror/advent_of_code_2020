package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := getInput()

	twoMatch, threeMatch := findMissingExpenses(input, 2020)

	fmt.Println("The two matching values total: ", twoMatch)
	fmt.Println("The three matching values total: ", threeMatch)
}

func getInput() []int {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	dataString := strings.Split(string(data), "\n")

	input := []int{}

	for _, s := range dataString {
		v, _ := strconv.Atoi(s)
		if v > 0 && v < 2019 {
			input = append(input, v)
		}
	}

	sort.Ints(input)

	return input
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
