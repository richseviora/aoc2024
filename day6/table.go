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
