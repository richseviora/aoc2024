package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCanBeComposedFrom(t *testing.T) {
	testCases := []struct {
		input     string
		sequences []string
		expected  [][]string
	}{
		{
			"a",
			[]string{"a", "b", "c"},
			[][]string{
				{"a"},
			},
		},
		{
			"ab",
			[]string{"a", "b"},
			[][]string{
				{"a", "b"},
			},
		},
		{
			"abc",
			[]string{"a", "b", "c"},
			[][]string{
				{"a", "b", "c"},
			},
		},
		{
			"aaa",
			[]string{"a", "b", "c"},
			[][]string{
				{"a", "a", "a"},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Test %v", tc.input), func(t *testing.T) {
			sequences := make([]Sequence, len(tc.sequences))
			outputCombinations := make([][]Sequence, 0)
			for _, sequence := range tc.expected {
				outputSlices := make([]Sequence, len(sequence))
				for i, s := range sequence {
					outputSlices[i] = NewSequence(s)
				}
				outputCombinations = append(outputCombinations, outputSlices)
			}
			for i, s := range tc.sequences {
				sequences[i] = NewSequence(s)
			}
			result := CanBeComposedFromMemoized(tc.input, sequences)
			if !reflect.DeepEqual(result, outputCombinations) {
				t.Errorf("expected %+v, got %+v", outputCombinations, result)
			}
		})
	}
}
