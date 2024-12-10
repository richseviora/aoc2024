package main

import (
	"errors"
	"github.com/kr/pretty"
	"reflect"
	"strconv"
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
			result, _ := CompactBlocks(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func CreateBlocksFromMap(input string) []Block {
	blocks := make([]Block, 0)
	for _, c := range input {
		if c == '.' {
			blocks = append(blocks, Block{-1})
		} else {
			blocks = append(blocks, Block{int(c - '0')})
		}
	}
	return blocks
}

func CreateMapFromBlocks(blocks []Block) string {
	string := ""
	for _, b := range blocks {
		if b.IsEmpty() {
			string += "."
		} else {
			string += strconv.Itoa(b.Id)
		}

	}
	return string
}

func Test_CompactFiles(t *testing.T) {
	testCases := []struct {
		name          string
		input         []Block
		expected      []Block
		untilComplete bool
	}{
		{"single non-empty block", []Block{Block{0}}, []Block{Block{0}}, false},
		{"1.1", []Block{Block{0}, Block{-1}, Block{1}}, CreateBlocksFromMap("01."), false},
		{"11111", ParseFileMap("11111"), CreateBlocksFromMap("021.."), false},
		{"133", ParseFileMap("133"), CreateBlocksFromMap("0111..."), false},
		{"135", ParseFileMap("135"), []Block{Block{0}, Block{-1}, Block{-1}, Block{-1}, Block{1}, Block{1}, Block{1}, Block{1}, Block{1}}, false},
		{"example", ParseFileMap("2333133121414131402"), CreateBlocksFromMap("00992111777.44.333....5555.6666.....8888.."), true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			callFunc := func() []Block {
				if tc.untilComplete {
					return CompactUntilComplete(tc.input, true)
				}
				result, _ := CompactBlocksWholeFiles(tc.input)
				return result
			}
			result := callFunc()
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v\ndiff: %v\nEXP: %v\nACT: %v", tc.expected, result, pretty.Diff(tc.expected, result), CreateMapFromBlocks(tc.expected), CreateMapFromBlocks(result))
			}
		})
	}
}

func Test_SwapBlocks(t *testing.T) {
	testCases := []struct {
		name           string
		input          []Block
		sourceStart    int
		destStart      int
		length         int
		expectedBlocks []Block
		expectedError  error
	}{
		{"invalid source", CreateBlocksFromMap("01"), 0, 0, 3, nil, LengthExceedsBlockError},
		{"simple swap", CreateBlocksFromMap("01"), 1, 0, 1, CreateBlocksFromMap("10"), nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := SwapBlocks(tc.input, tc.sourceStart, tc.destStart, tc.length)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("expected %v, got %v", tc.expectedError, err)
			}
			if !reflect.DeepEqual(result, tc.expectedBlocks) {
				t.Errorf("expected %v, got %v", tc.expectedBlocks, result)
			}
		})
	}
}
