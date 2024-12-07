package main

import (
	"reflect"
	"testing"
)

func TestNewEquation(t *testing.T) {
	tests := []struct {
		input    string
		result   int
		operands []int
	}{
		{"190: 10 19", 190, []int{10, 19}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			equation := NewEquation(tt.input)
			if !reflect.DeepEqual(tt.operands, equation.Operands) || !reflect.DeepEqual(tt.result, equation.Result) {
				t.Errorf("NewEquation(%v) = %v, want %v", tt.input, equation, tt)
			}
		})
	}
}
