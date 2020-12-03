package aoc

import (
	"io/ioutil"
	"strings"
)

// ReadInput reads a file and splits the lines into an array
func ReadInput(fileName string) []string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}
