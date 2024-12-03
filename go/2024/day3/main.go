package main

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
)

// uncontrolledScanCorruptedMemory is initial solution to puzzle
func uncontrolledScanCorruptedMemory(input []string) {
	// sample is string while actual input is []string
	// sampleInput := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	// Regex pattern
	pattern := `mul\((\d+),(\d+)\)` // also be in the format: `mul\((\d{1,3}),(\d{1,3})\)`

	// Compile the regex
	regex := regexp.MustCompile(pattern)

	// Find all matches in sample input
	// matches := regex.FindAllStringSubmatch(sampleInput, -1)

	// Find all matches actual input
	var matches [][]string
	for _, inputLine := range input {
		match := regex.FindAllStringSubmatch(inputLine, -1)
		matches = append(matches, match...)
	}

	// Calculate the sum of products
	sum := 0
	for _, match := range matches {
		// match[1] is the first number, match[2] is the second number
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2
	}

	// Print the result
	fmt.Printf("initial puzzle answer: %d\n", sum)
}

// controlledScanCorruptedMemory is final solution to puzzle
func controlledScanCorruptedMemory(input []string, isControlled bool) {
	// controlledSampleInput := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	controlledPattern := `(do\(\)|don't\(\))|mul\((\d+),(\d+)\)` // can also be `(do\(\)|don't\(\))|mul\((\d{1,3}),(\d{1,3})\)`

	// Compile the regex
	regex := regexp.MustCompile(controlledPattern)

	// Find all matches in sample input
	// matches := regex.FindAllStringSubmatch(controlledSampleInput, -1)

	// // Find all matches actual input
	var matches [][]string
	for _, inputLine := range input {
		match := regex.FindAllStringSubmatch(inputLine, -1)
		matches = append(matches, match...)
	}

	// Calculate the controlledSum of products
	controlledSum := 0
	mulEnabled := true
	// match is in the form [... do/dont num1 num2]
	// if do/dont match [... do/dont emptyString emptyString]
	// if mul match match [... emptyString num1 num2]
	for _, match := range matches {
		if match[1] == "don't()" && isControlled {
			mulEnabled = false
			continue
		}
		if match[1] == "do()" && isControlled {
			mulEnabled = true
			continue
		}

		if mulEnabled {
			num1, _ := strconv.Atoi(match[2])
			num2, _ := strconv.Atoi(match[3])
			controlledSum += (num1 * num2)
		}
	}

	// Print the result
	fmt.Printf("puzzle answer: %d\n", controlledSum)
}

func main() {
	// input: corrupted memory
	input, _ := util.ReadMultiLinesFile()
	// scanCorruptedMemory(input)
	controlledScanCorruptedMemory(input, false)
	controlledScanCorruptedMemory(input, true)
}
