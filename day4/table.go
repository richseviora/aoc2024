package main

import "strings"

type Direction int64

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Table struct {
	rows [][]Cell
}

func NewTable(input string) Table {
	rows := strings.Split(input, "\n")
	table := Table{
		rows: make([][]Cell, 0),
	}
	for ri, row := range rows {
		newCellValues := make([]Cell, 0, len(row))
		for ci, char := range row {
			newCellValues = append(newCellValues, Cell{Value: string(char), Table: &table, Row: ri, Column: ci})
		}
		table.rows = append(table.rows, newCellValues)
	}
	return table

}

func (t Table) GetCellAt(row int, column int) *Cell {
	return &t.rows[row][column]
}

func (t Table) IsInRange(row int, column int) bool {
	if (row < 0) || (row >= len(t.rows)) {
		return false
	}
	if (column < 0) || (column >= len(t.rows[row])) {
		return false
	}
	return true
}

type Cell struct {
	Value  string
	Table  *Table
	Row    int
	Column int
}

func (c Cell) GetCellInDirection(d Direction) *Cell {
	switch d {
	case Up:
		return c.Table.GetCellAt(c.Row+1, c.Column)
	case Down:
		return c.Table.GetCellAt(c.Row-1, c.Column)
	case Left:
		return c.Table.GetCellAt(c.Row, c.Column-1)
	case Right:
		return c.Table.GetCellAt(c.Row, c.Column+1)
	}
	panic("Invalid Direction")
}
