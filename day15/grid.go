package main

import (
	"slices"
	"strings"
)

type Grid struct {
	Coordinates   map[Coordinate]Cell
	height, width int
}

func (g *Grid) AddCell(c Coordinate, content string) {
	g.Coordinates[c] = Cell{Content: content}
	if c.x > g.width {
		g.width = c.x
	}
	if c.y > g.height {
		g.height = c.y
	}
}

func (g *Grid) GetCell(c Coordinate) Cell {
	return g.Coordinates[c]
}

func (g *Grid) ToString() string {
	output := ""
	for y := 0; y <= g.height; y++ {
		for x := 0; x <= g.width; x++ {
			output += string(g.GetCell(Coordinate{x, y}).Content)
		}
		output += "\n"
	}
	return output
}

func (g *Grid) GetRobot() Coordinate {
	for coordinate, cell := range g.Coordinates {
		if cell.IsRobot() {
			return coordinate
		}
	}
	panic("No robot found")
}

func (g *Grid) HandleDirection(d Direction) bool {
	robot := g.GetRobot()
	nextCell := robot.GoDirection(d)
	cellsToMove := []Coordinate{robot}
	possibleMove := false
	for nextCell != nil {
		// check if valid move, empty
		if g.GetCell(*nextCell).IsEmpty() {
			possibleMove = true
			nextCell = nil
			continue
		}
		if g.GetCell(*nextCell).IsOccupied() {
			cellsToMove = append(cellsToMove, *nextCell)
			nextCell = nextCell.GoDirection(d)
			continue
		}
		if g.GetCell(*nextCell).IsWall() {
			possibleMove = false
			nextCell = nil
			continue
		}
	}
	// move cells {
	if !possibleMove {
		return false
	}
	slices.Reverse(cellsToMove)
	for _, cell := range cellsToMove {
		destCell := cell.GoDirection(d)
		if destCell == nil {
			panic("unexpected nil")
		}
		g.Coordinates[*destCell] = g.Coordinates[cell]
		g.Coordinates[cell] = Cell{Content: "."}
	}
	return true
}

type Cell struct {
	Content string
}

func (c Cell) IsWall() bool {
	return c.Content == "#"
}

func (c Cell) IsEmpty() bool {
	return c.Content == "."
}

func (c Cell) IsOccupied() bool {
	return c.Content == "O"
}

func (c Cell) IsRobot() bool {
	return c.Content == "@"
}

type Coordinate struct {
	x, y int
}

func (c Coordinate) GoDirection(direction Direction) *Coordinate {
	switch direction {
	case Up:
		return &Coordinate{x: c.x, y: c.y - 1}
	case Down:
		return &Coordinate{x: c.x, y: c.y + 1}
	case Left:
		return &Coordinate{x: c.x - 1, y: c.y}
	case Right:
		return &Coordinate{x: c.x + 1, y: c.y}
	}
	panic("Invalid Direction")
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func DirectionFromChar(c rune) Direction {
	switch c {
	case '^':
		return Up
	case 'v':
		return Down
	case '<':
		return Left
	case '>':
		return Right
	}
	panic("Invalid Direction")
}

func NewGridFromInput(input string) (*Grid, []Direction) {
	inputs := strings.Split(input, "\n")
	directions := make([]Direction, 0)
	grid := &Grid{
		Coordinates: map[Coordinate]Cell{},
	}
	readingGrid := true
	for y, line := range inputs {
		if line == "" {
			readingGrid = false
			continue
		}
		for x, char := range line {
			if readingGrid {
				grid.AddCell(Coordinate{x, y}, string(char))
			} else {
				directions = append(directions, DirectionFromChar(char))
			}

		}

	}
	return grid, directions
}
