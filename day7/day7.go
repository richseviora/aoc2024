package main

import (
	"fmt"
	"github.com/mowshon/iterium"
	"os"
	"strconv"
	"strings"
)

type Operator int64

const (
	Add Operator = iota
	Multiply
	Append
)

const fileName = "input.txt"

type Equation struct {
	Result    int
	Operands  []int
	IsValid   bool
	Operators []Operator
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
			case Append:
				total, _ = strconv.Atoi(fmt.Sprintf("%d%d", total, e.Operands[i+1]))
			}
		}
	}
	return total
}

func (e *Equation) Solve() bool {
	operators := []Operator{Add, Multiply, Append}
	generator := iterium.Product(operators, len(e.Operands)-1)
	permutations, err := generator.Slice()
	if err != nil {
		panic(err)
	}
	for _, permutation := range permutations {
		if e.Evaluate(permutation) == e.Result {
			e.IsValid = true
			e.Operators = permutation
			return true
		}
	}
	e.IsValid = false
	return false
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
	fileContent := ReadInput(fileName)
	equations := make([]Equation, 0)
	total := 0
	for _, line := range strings.Split(fileContent, "\n") {
		if line == "" {
			continue
		}
		equations = append(equations, NewEquation(line))
	}
	for i, equation := range equations {
		if equation.Solve() {
			total += equation.Result

		}
		fmt.Printf("%d: %+v\n", i, equation)
	}
	fmt.Println(len(equations))
	fmt.Println("Part 1 Total", total)
}
