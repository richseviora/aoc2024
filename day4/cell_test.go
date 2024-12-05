package main

import (
	"reflect"
	"testing"
)

func TestCell_GetCellsInDirection(t *testing.T) {
	table := NewTable("12345\n67890\nABCDE\nFGHIJ")
	testCases := []struct {
		name      string
		row       int
		column    int
		direction Direction
		length    int
		expected  []*Cell
	}{
		{"safe 0,0", 1, 1, Down, 2, []*Cell{
			table.GetCellAt(1, 1),
			table.GetCellAt(2, 1),
		}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			firstCell := table.GetCellAt(tc.row, tc.column)
			cells := firstCell.GetCellsInDirection(tc.direction, tc.length)
			if !reflect.DeepEqual(tc.expected, cells) {
				t.Errorf("expected %v,\n got %v", tc.expected, cells)
			}
		})
	}
}
