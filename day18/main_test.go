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
		{"input.txt", 71, 1024, 0},
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
