package main

import (
	"aoc/util"
	"fmt"
)

func getFinalFloor(input string) int {
	floorCount := 0
	directions := map[string]int{
		"(": 1,
		")": -1,
	}
	for idx, stepRune := range input {
		floorCount += directions[string(stepRune)]
		if floorCount == -1 {
			fmt.Println("position of negative floor: ", idx+1)
		}
	}
	return floorCount
}

func main() {
	input, _ := util.ReadOneLineFile()
	fmt.Println(getFinalFloor(input))
}
