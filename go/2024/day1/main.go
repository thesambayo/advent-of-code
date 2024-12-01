package main

import (
	"aoc/util"
	"fmt"
	"math"
	"sort"
)

func getTotalDistance(input [][]int) int {
	totalDistance := 0

	sort.Ints(input[0])
	sort.Ints(input[1])

	for idx := 0; idx < len(input[0]); idx++ {
		totalDistance += int(math.Abs(float64(input[0][idx] - input[1][idx])))
	}
	return totalDistance
}

func findNumOfOccurencesInSlice(searchList []int) map[int]int {
	searchedIds := map[int]int{}
	for _, id := range searchList {
		searchedIds[id] += 1
	}

	return searchedIds
}

func getSimilarityScore(input [][]int) int {
	var score int

	// searchedIds keeps track of ids that already have their num of occurrence checked.
	// This will be used to check if an id has already been checked so the num of occurence can be reused
	searchedIds := findNumOfOccurencesInSlice(input[1])

	for _, id := range input[0] {
		if occurrence, ok := searchedIds[id]; ok {
			score += id * occurrence
		}
	}

	return score
}

func main() {
	input, _ := util.ReadMultiColumnsIntFile()
	fmt.Println("answer for part 1: ", getTotalDistance(input))
	fmt.Println("answer for part 2: ", getSimilarityScore(input))
}
