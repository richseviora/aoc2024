package main

import (
	"os"
	"strings"
)

type Coordinate struct {
	x, y int
}

type Grid struct {
	cells      map[Coordinate]string
	maxX, maxY int
}

var enableDetail = os.Getenv("ENABLE_DETAIL") == "true"

func NewGridFromInput(input string) *Grid {
	lines := strings.Split(input, "\n")

	cells := make(map[Coordinate]string)
	maxY := 0
	maxX := 0

	for i, line := range lines {
		if i > maxY {
			maxY = i
		}
		for j, char := range line {
			if j > maxX {
				maxX = j
			}
			coordinate := Coordinate{x: j, y: i}
			cells[coordinate] = string(char)
		}
	}
	grid := &Grid{
		cells: cells,
		maxX:  maxX,
		maxY:  maxY,
	}
	return grid
}
