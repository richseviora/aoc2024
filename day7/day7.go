package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operator int64

const (
	Add Operator = iota
	Multiply
)

const fileName = "input.txt"

type Equation struct {
	Result   int
	Operands []int
	IsValid  *bool
	Operator []Operator
}

func NewEquation(input string) Equation {
	inputs := strings.Split(input, " ")
	resultStr := inputs[0][:len(inputs[0])-1]
	result, _ := strconv.Atoi(resultStr)
	operands := make([]int, len(inputs)-1)
	for i, input := range inputs[1:] {
		operands[i], _ = strconv.Atoi(input)
	}
	return Equation{
		Result:   result,
		Operands: operands,
	}
}

func (e *Equation) Evaluate(operators []Operator) int {
	if len(operators) != len(e.Operands)-1 {
		panic("operators must be one shorter than operands")
	}
	total := e.Operands[0]
	for i, operator := range operators {
		{
			switch operator {
			case Add:
				total += e.Operands[i+1]
			case Multiply:
				total *= e.Operands[i+1]
			}
		}
	}
	return total
}

func ReadInput(fname string) string {
	file, err := os.Open(fname)
	defer file.Close()
	if err != nil {
		fmt.Printf("error received: %e", err)
		panic(err)
	}

	content, err := os.ReadFile(fname)
	if err != nil {
		fmt.Printf("error received: %v", err)
		panic(err)
	}
	return string(content)
}

func main() {

}
