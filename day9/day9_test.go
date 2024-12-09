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

func Test_CompactBlocks(t *testing.T) {
	testCases := []struct {
		name     string
		input    []Block
		expected []Block
	}{
		{"single non-empty block", []Block{Block{0}}, []Block{Block{0}}},
		{"1.1", []Block{Block{0}, Block{-1}, Block{1}}, []Block{Block{0}, Block{1}}},
		{"11111", ParseFileMap("11111"), []Block{Block{0}, Block{2}, Block{1}, Block{-1}}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := CompactBlocks(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
