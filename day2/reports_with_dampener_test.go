package day2

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerateReportsWithDampener(t *testing.T) {
	// Example test data

	testCases := []struct {
		reports         []int
		expectedChanges []int
	}{
		{reports: []int{1, 2, 3, 4, 5}, expectedChanges: []int{1, 1, 1, 1}},
		{reports: []int{1, 2, 3, 2, 1}, expectedChanges: []int{1, 1, -1, -1}},
	}
	for _, test := range testCases {
		t.Run(fmt.Sprintf("Test Report %v", test.reports), func(t *testing.T) {
			report := GenerateReports(test.reports, true)
			if !reflect.DeepEqual(report.Changes(), test.expectedChanges) {
				t.Errorf("expected Changes %v, got %v", test.expectedChanges, report.Changes())
			}
		})
	}
}

func TestIsSafeWithDampener(t *testing.T) {
	testCases := []struct {
		name     string
		reports  []int
		expected bool
	}{
		{name: "Safe without removing any level", reports: []int{7, 6, 4, 2, 1}, expected: true},
		{name: "Unsafe regardless of which level is removed", reports: []int{1, 2, 7, 8, 9}, expected: false},
		{name: "Unsafe regardless of which level is removed", reports: []int{9, 7, 6, 2, 1}, expected: false},
		{name: "Safe by removing the second level, 3", reports: []int{1, 3, 2, 4, 5}, expected: true},
		{name: "Safe by removing the third level, 4", reports: []int{8, 6, 4, 4, 1}, expected: true},
		{name: "Safe without removing any level", reports: []int{1, 3, 6, 7, 9}, expected: true},
	}
	for _, test := range testCases {
		t.Run(fmt.Sprintf("Test Report %v", test.name), func(t *testing.T) {
			report := GenerateReports(test.reports, true)
			if !reflect.DeepEqual(report.IsSafe(), test.expected) {
				t.Errorf("expected isSafe %v, got %v", test.expected, report.IsSafe())
			}
		})
	}
}
