package day2

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerateReports(t *testing.T) {
	// Example test data

	testCases := []struct {
		reports         []int
		expectedChanges []int
	}{
		{reports: []int{1, 2, 3, 4, 5}, expectedChanges: []int{1, 1, 1, 1}},
		{reports: []int{1, 2, 3, 2, 1}, expectedChanges: []int{1, 1, -1, -1}},
	}
	for _, test := range testCases {
		t.Run(fmt.Sprintf("Test Reports %v", test.reports), func(t *testing.T) {
			report := GenerateReports(test.reports)
			if !reflect.DeepEqual(report.changes, test.expectedChanges) {
				t.Errorf("expected Changes %v, got %v", test.expectedChanges, report.changes)
			}
		})
	}
}

func TestIsSafe(t *testing.T) {
	testCases := []struct {
		name     string
		reports  []int
		expected bool
	}{
		{name: "Safe - All Ascending by 1", reports: []int{1, 2, 3, 4, 5}, expected: true},
		{name: "Unsafe - example 1", reports: []int{1, 2, 7, 8, 9}, expected: false},
		{name: "Unsafe - example 2", reports: []int{9, 7, 6, 2, 1}, expected: false},
	}
	for _, test := range testCases {
		t.Run(fmt.Sprintf("Test Reports %v", test.name), func(t *testing.T) {
			report := GenerateReports(test.reports)
			if !reflect.DeepEqual(report.IsSafe(), test.expected) {
				t.Errorf("expected Changes %v, got %v", test.expected, report.IsSafe())
			}
		})
	}
}
