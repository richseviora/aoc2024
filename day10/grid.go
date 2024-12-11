package main

import (
	"strings"
)

// Grid represents a collection of cells organized in a nested map
type Grid struct {
	Cells map[int]map[int]*Cell
}

// NewGrid creates and returns a new Grid instance
func NewGrid() *Grid {
	return &Grid{
		Cells: make(map[int]map[int]*Cell),
	}
}

// NewGridFromString creates a new Grid from an input string
// Each line in the string represents a row, and each character in a line represents a cell
func NewGridFromString(input string) *Grid {
	grid := NewGrid()
	rows := strings.Split(input, "\n") // Split input into lines
	for rowIndex, line := range rows {
		for colIndex, char := range line {
			value := int(char - '0')                // Convert the char rune to an integer
			grid.SetCell(rowIndex, colIndex, value) // Set each character as a cell value
		}
	}
	return grid
}

// GetCell returns the cell at the specified row and column
// If the cell does not exist, it returns nil
func (g *Grid) GetCell(row, col int) *Cell {
	if g.Cells[row] != nil {
		return g.Cells[row][col]
	}
	return nil
}

// SetCell sets a cell at the specified row and column with the given value
func (g *Grid) SetCell(row, col int, value int) {
	if g.Cells[row] == nil {
		g.Cells[row] = make(map[int]*Cell)
	}
	g.Cells[row][col] = &Cell{Value: value, Coordinates: Coordinates{Row: row, Column: col}, Grid: g}
}

// IterateCells executes the provided callback function for each cell in the grid
func (g *Grid) IterateCells(callback func(row, col int, cell *Cell)) {
	for row, cols := range g.Cells {
		for col, cell := range cols {
			callback(row, col, cell)
		}
	}
}

func (g *Grid) FindTrailheads() []*Cell {
	trailheads := make([]*Cell, 0)
	g.IterateCells(func(row, col int, cell *Cell) {
		if cell == nil {
			panic("nil cell")
		}
		if cell.Value == 0 {
			trailheads = append(trailheads, cell)
		}
	})
	return trailheads
}
