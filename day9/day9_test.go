package main

import (
	"reflect"
	"testing"
)

func Test_ParseFileMap(t *testing.T) {
	testCases := []struct {
		input    string
		expected []Block
	}{
		{"111", []Block{Block{0}, Block{-1}, Block{1}}},
		{"213", []Block{Block{0}, Block{0}, Block{-1}, Block{1}, Block{1}, Block{1}}},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := ParseFileMap(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
