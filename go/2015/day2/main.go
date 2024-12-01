package main

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

type Cubebox struct {
	length int
	width  int
	height int
}

func (cube Cubebox) getSurfaceArea() int {
	return 2 * (cube.height*cube.length + cube.height*cube.width + cube.length*cube.width)
}

func (cube Cubebox) getVolume() int {
	return cube.length * cube.height * cube.width
}

func (cube Cubebox) getSmallestArea() int {
	areaLW := cube.length * cube.width
	areaLH := cube.length * cube.height
	areaWH := cube.width * cube.height

	// Find and return the smallest area
	if areaLW <= areaLH && areaLW <= areaWH {
		return areaLW
	} else if areaLH <= areaLW && areaLH <= areaWH {
		return areaLH
	}
	return areaWH
}

func (cube Cubebox) getSmallestPerimeter() int {
	perimeterLW := 2 * (cube.length + cube.width)
	perimeterLH := 2 * (cube.length + cube.height)
	perimeterWH := 2 * (cube.width + cube.height)

	// Find and return the smallest area
	if perimeterLW <= perimeterLH && perimeterLW <= perimeterWH {
		return perimeterLW
	} else if perimeterLH <= perimeterLW && perimeterLH <= perimeterWH {
		return perimeterLH
	}
	return perimeterWH
}

// dimension => 2X3X4 (lXwXh)
func NewCubeBox(dimension string) Cubebox {
	dimensions := strings.Split(dimension, "x")
	length, _ := strconv.Atoi(dimensions[0])
	width, _ := strconv.Atoi(dimensions[1])
	height, _ := strconv.Atoi(dimensions[2])
	return Cubebox{
		length: length,
		width:  width,
		height: height,
	}
}

func processWrapperPaperNeeded(cube Cubebox) int {
	return cube.getSurfaceArea() + cube.getSmallestArea()
}

func proceddRibbonNeeded(cube Cubebox) int {
	return cube.getSmallestPerimeter() + cube.getVolume()
}

func main() {
	input, _ := util.ReadMultiLinesFile()
	var totalOrder int
	var totalRibbon int

	for _, inputLine := range input {
		cube := NewCubeBox(inputLine)
		totalOrder += processWrapperPaperNeeded(cube)
		totalRibbon += proceddRibbonNeeded(cube)
	}

	fmt.Println(totalOrder)
	fmt.Printf("Total wrapper needed: %d\n", totalOrder)
	fmt.Printf("Total ribbon needed: %d\n", totalRibbon)
}
