package aoc

import (
	"strconv"
)

type instruction struct {
	op   string
	arg  int
	read bool
}

// Day8 - https://adventofcode.com/2020/day/8
func Day8(filename string) (int, int) {
	rawInput := ReadInput(filename)
	instructions := parseInstructions(rawInput)

	infiniteAcc, _ := run(instructions)
	var terminatedAcc int

	for i := range instructions {
		instructionSet := parseInstructions(rawInput)
		instruction := instructionSet[i]

		var terminated bool
		switch instruction.op {
		case "nop":
			instruction.op = "jmp"
			instructionSet[i] = instruction
			terminatedAcc, terminated = run(instructionSet)
		case "jmp":
			instruction.op = "nop"
			instructionSet[i] = instruction
			terminatedAcc, terminated = run(instructionSet)
		}

		if terminated {
			break
		}
	}

	return infiniteAcc, terminatedAcc
}

func run(instructions []instruction) (int, bool) {
	var acc int
	var pointer int
	var terminated bool

	for {
		if pointer >= len(instructions) {
			terminated = true
			break
		}

		instruction := instructions[pointer]
		if instruction.read {
			break
		}
		instruction.read = true
		instructions[pointer] = instruction
		switch instruction.op {
		case "nop":
			pointer++
		case "acc":
			acc += instruction.arg
			pointer++
		case "jmp":
			pointer += instruction.arg
		default:
			break
		}
	}
	return acc, terminated
}

func parseInstructions(input []string) []instruction {
	instructions := []instruction{}

	for _, row := range input {
		if len(row) > 0 {
			op := row[:3]
			arg, _ := strconv.Atoi(row[4:])
			instructions = append(instructions, instruction{op, arg, false})
		}
	}

	return instructions
}
