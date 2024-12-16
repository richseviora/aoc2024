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

func DropFirstSlice(s []*Coordinate) []*Coordinate {
	if len(s) == 0 {
		return s
	}
	if len(s) == 1 {
		return []*Coordinate{}
	}
	return s[1:]
}

func (g *Grid) HandleDirection(d Direction) bool {
	robot := g.GetRobot()
	nextCoordinates := []*Coordinate{robot.GoDirection(d)}
	cellsToMove := []Coordinate{robot}
	possibleMove := true
	visited := map[Coordinate]bool{}
	for len(nextCoordinates) > 0 && nextCoordinates[0] != nil {
		// check if valid move, empty
		nextCoordinate := nextCoordinates[0]
		if visited[*nextCoordinate] {
			nextCoordinates = DropFirstSlice(nextCoordinates)
			continue
		}
		nextCoordinates = DropFirstSlice(nextCoordinates)
		visited[*nextCoordinate] = true
		cell := g.GetCell(*nextCoordinate)
		if cell.IsEmpty() {
			continue
		}
		if cell.IsOccupied() {
			if cell.IsOccupiedSingle() {
				cellsToMove = append(cellsToMove, *nextCoordinate)
				nextCoordinates = append(nextCoordinates, nextCoordinate.GoDirection(d))
				continue
			} else if cell.IsOccupiedLeft() {
				occupiedRight := nextCoordinate.GetOccupiedRight(g)
				visited[occupiedRight] = true
				occupiedLeftNext := nextCoordinate.GoDirection(d)
				occupiedRightNext := occupiedRight.GoDirection(d)
				cellsToMove = append(cellsToMove, *nextCoordinate, occupiedRight)
				nextCoordinates = append(nextCoordinates, occupiedLeftNext, occupiedRightNext)
				continue
			} else if cell.IsOccupiedRight() {
				occupiedLeft := nextCoordinate.GetOccupiedLeft(g)
				visited[occupiedLeft] = true
				occupiedLeftNext := nextCoordinate.GoDirection(d)
				occupiedRightNext := occupiedLeft.GoDirection(d)
				cellsToMove = append(cellsToMove, *nextCoordinate, occupiedLeft)
				nextCoordinates = append(nextCoordinates, occupiedLeftNext, occupiedRightNext)
				continue
			}

		}
		if cell.IsWall() {
			possibleMove = false
			nextCoordinates = make([]*Coordinate, 0)
			continue
		}

	}
	// move cells {
	if !possibleMove {
		return false
	}
	slices.Reverse(cellsToMove)
	changes := make(map[Coordinate]Cell)
	for _, cell := range cellsToMove {
		destCell := cell.GoDirection(d)
		if destCell == nil {
			panic("unexpected nil")
		}
		changes[*destCell] = g.Coordinates[cell]
	}
	for _, cell := range cellsToMove {
		if _, ok := changes[cell]; !ok {
			changes[cell] = Cell{Content: "."}
		}

	}
	changes[robot] = Cell{Content: "."}
	for coordinate, cell := range changes {
		g.Coordinates[coordinate] = cell
	}
	return true
}

func (g *Grid) GetBoxGPSTotal() int {
	total := 0
	for coordinate, cell := range g.Coordinates {
		if cell.IsOccupiedSingle() || cell.IsOccupiedLeft() {
			total += coordinate.CalculateGPS()
		}
	}
	return total
}

type Cell struct {
	Content string
}

func (c *Cell) IsWall() bool {
	return c.Content == "#"
}

func (c *Cell) IsEmpty() bool {
	return c.Content == "."
}

func (c *Cell) IsOccupied() bool {
	return c.Content == "[" || c.Content == "]" || c.Content == "O"
}

func (c *Cell) IsOccupiedSingle() bool {
	return c.Content == "O"
}

func (c *Cell) IsOccupiedLeft() bool {
	return c.Content == "["
}

func (c *Cell) IsOccupiedRight() bool {
	return c.Content == "]"
}

func (c *Cell) IsRobot() bool {
	return c.Content == "@"
}

type Coordinate struct {
	x, y int
}

func (c *Coordinate) GetOccupiedLeft(g *Grid) Coordinate {
	cell := g.GetCell(*c)
	if cell.IsOccupiedLeft() {
		return *c
	}
	if cell.IsOccupiedRight() {
		return *c.GoDirection(Left)
	}
	panic("No occupied left for this cell")
}

func (c *Coordinate) GetOccupiedRight(g *Grid) Coordinate {
	cell := g.GetCell(*c)
	if cell.IsOccupiedRight() {
		return *c
	}
	if cell.IsOccupiedLeft() {
		return *c.GoDirection(Right)
	}
	panic("No occupied right for this cell")
}

func (c *Coordinate) CalculateGPS() int {
	return c.x + c.y*100
}

func (c *Coordinate) GoDirection(direction Direction) *Coordinate {
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

func NewGridFromInput(input string, doubleWide bool) (*Grid, []Direction) {
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
				if doubleWide {
					firstChar, secondChar := func(input string) (string, string) {
						if input == "#" {
							return "#", "#"
						} else if input == "O" {
							return "[", "]"
						} else if input == "." {
							return ".", "."
						} else if input == "@" {
							return "@", "."
						}
						panic("Invalid input")
					}(string(char))
					grid.AddCell(Coordinate{x * 2, y}, firstChar)
					grid.AddCell(Coordinate{(x * 2) + 1, y}, secondChar)
				} else {
					grid.AddCell(Coordinate{x, y}, string(char))
				}

			} else {
				directions = append(directions, DirectionFromChar(char))
			}

		}

	}
	return grid, directions
}
