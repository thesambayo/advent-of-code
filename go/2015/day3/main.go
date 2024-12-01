package main

import (
	"aoc/util"
	"fmt"
	"slices"
)

type Position struct {
	x int
	y int
}

type House struct {
	position      Position
	presentsCount int
}

func NewHouse(position Position) House {
	return House{
		position:      position,
		presentsCount: 1,
	}
}

func (house *House) UpdateNumberOfPresents() {
	house.presentsCount += 1
}

func (house House) isAtSameLocationAsPosition(position Position) bool {
	if house.position.x == position.x && house.position.y == position.y {
		return true
	}
	return false
}

type Santa struct {
	lastPosition Position
}

func createSanta(startingPosition Position) Santa {
	return Santa{
		lastPosition: startingPosition,
	}
}

func (santa *Santa) updatePosition(direction rune) {
	switch dir := string(direction); dir {
	// upward
	case "^":
		santa.lastPosition = Position{x: santa.lastPosition.x, y: santa.lastPosition.y + 1}
		// right
	case ">":
		santa.lastPosition = Position{x: santa.lastPosition.x + 1, y: santa.lastPosition.y}
	// down
	case "v":
		santa.lastPosition = Position{x: santa.lastPosition.x, y: santa.lastPosition.y - 1}
		// left
	default:
		santa.lastPosition = Position{x: santa.lastPosition.x - 1, y: santa.lastPosition.y}
	}
}

func main() {
	inputLine, _ := util.ReadOneLineFile()

	startingPosition := Position{x: 0, y: 0}
	santa := createSanta(startingPosition)
	roboSanta := createSanta(startingPosition)
	houses := []House{NewHouse(startingPosition)}
	houses[0].UpdateNumberOfPresents()

	for idx, movement := range inputLine {
		currentSanta := Santa{}
		if isEven := idx%2 != 0; isEven {
			santa.updatePosition(movement)
			currentSanta = santa
		} else {
			roboSanta.updatePosition(movement)
			currentSanta = roboSanta
		}

		houseIndexSantaMovedTo := slices.IndexFunc(houses, func(house House) bool {
			return house.isAtSameLocationAsPosition(currentSanta.lastPosition)
		})

		if houseIndexSantaMovedTo < 0 {
			houses = append(houses, NewHouse(currentSanta.lastPosition))
		} else {
			houses[houseIndexSantaMovedTo].UpdateNumberOfPresents()
		}
	}

	fmt.Println(len(houses))
}
