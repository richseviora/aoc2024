package main

import (
	"reflect"
	"testing"
)

func TestNewTable(t *testing.T) {
	table := NewTable("12345\n67890\nABCDE\nFGHIJ")
	testCases := []struct {
		name     string
		row      int
		column   int
		expected *Cell
	}{
		{"safe 0,0", 1, 1, &Cell{
			Value:  "7",
			Table:  &table,
			Row:    1,
			Column: 1,
		}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cell := table.GetCellAt(tc.row, tc.column)
			if !reflect.DeepEqual(cell, tc.expected) {
				t.Errorf("expected %v,\n got %v", tc.expected, cell)
			}
		})
	}
}

func TestTable_IsInRange(t *testing.T) {
	table := NewTable("12345\n67890\nABCDE\nFGHIJ")
	testCases := []struct {
		name     string
		row      int
		column   int
		expected bool
	}{
		{"safe 0,0", 1, 1, true},
		{"unsafe -1,0", -1, 1, false},
		{"unsafe 5,0", 5, 1, false},
		{"safe 4,0", 3, 1, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expected != table.IsInRange(tc.row, tc.column) {
				t.Errorf("expected %v, got %v", tc.expected, table.IsInRange(tc.row, tc.column))
			}
		})
	}
}
