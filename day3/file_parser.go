package day3

import (
	"fmt"
	regexp2 "regexp"
	"strconv"
)

type Pair struct {
	left  int
	right int
}

func GeneratePair(inputs []string) Pair {
	left, _ := strconv.Atoi(inputs[1])
	right, _ := strconv.Atoi(inputs[2])
	return Pair{left: left, right: right}
}

func (p Pair) Multiply() int {
	return p.left * p.right
}

func ParseString(input string) {
	regexp := regexp2.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := regexp.FindAllStringSubmatch(input, -1)
	pairs := make([]Pair, 0)
	total := 0
	for _, match := range matches {
		newPair := GeneratePair(match)
		total += newPair.Multiply()
		pairs = append(pairs, newPair)
	}
	fmt.Println(matches)
	fmt.Println(pairs)
	fmt.Println("Total of Multiples is", total)
}
