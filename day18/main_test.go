package main

import "testing"

func TestProcessChallenge(t *testing.T) {
	testCases := []struct {
		fileName string
		size     int
		limit    int
		expected int
	}{
		{"test1.txt", 7, 12, 22},
		{"input.txt", 71, 2869, 0},
	}
	for _, tc := range testCases {
		t.Run(tc.fileName, func(t *testing.T) {
			result := ProcessChallenge(tc.fileName, tc.size, tc.limit)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestGrid_FindFirstFailingInput(t *testing.T) {
	t.Run("find location", func(t *testing.T) {
		result := ProcessChallengePart2("input.txt", 71, 12)
		if result != 1024 {
			t.Errorf("expected %v, got %v", 1024, result)
		}
	})
}
