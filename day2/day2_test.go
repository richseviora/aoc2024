package day2

import (
	"reflect"
	"testing"
)

func TestReports(t *testing.T) {
	// Example test data
	report := GenerateReports([]int{1, 2, 3, 4, 5})
	expectedChanges := []int{1, 1, 1, 1}

	if !reflect.DeepEqual(report.changes, expectedChanges) {
		t.Errorf("expected Changes %v, got %v", expectedChanges, report.changes)
	}
}
