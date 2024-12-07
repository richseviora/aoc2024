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

func TestEquation_Evaluate(t *testing.T) {
	tests := []struct {
		input     string
		result    int
		operators []Operator
	}{
		{"190: 10 19", 29, []Operator{Add}},
		{"190: 10 19", 190, []Operator{Multiply}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			equation := NewEquation(tt.input)
			actual := equation.Evaluate(tt.operators)
			if tt.result != actual {
				t.Errorf("NewEquation(%v) = %v, want %v", tt.input, actual, tt.result)
			}
		})
	}
}

func TestEquation_Solve(t *testing.T) {
	tests := []struct {
		input     string
		solvable  bool
		operators []Operator
	}{
		{"190: 10 19", true, []Operator{Multiply}},
		{"29: 10 19", true, []Operator{Add}},
		{"39: 10 19", false, []Operator{}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			equation := NewEquation(tt.input)
			actual := equation.Solve()
			if tt.solvable != actual {
				t.Errorf("NewEquation(%v) = %v, want %v", tt.input, actual, tt.solvable)
			}
			//bothEmpty := len(tt.operators) == 0 && len(equation.Operators) == 0
			if !reflect.DeepEqual(tt.operators, equation.Operators) {
				t.Errorf("NewEquation(%v) = %v, want %v", tt.input, equation.Operators, tt.operators)
			}
		})
	}
}
