package main

import "strings"

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

type ReturnType struct {
	Cell      *Cell
	Direction Direction
}

func (t Table) IterateOverTable(teststring string) []ReturnType {
	returns := make([]ReturnType, 0)
	directions := []Direction{Up, Down, Left, Right, UpLeft, UpRight, DownLeft, DownRight}
	for ri, row := range t.rows {
		for ci, _ := range row {
			cell := t.GetCellAt(ri, ci)
			for _, direction := range directions {
				if cell.GetCellValuesInDirection(direction, len(teststring)) == teststring {
					returns = append(returns, ReturnType{Cell: cell, Direction: direction})
				}
			}
		}
	}
	return returns
}
