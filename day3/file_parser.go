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

func (p Pair) isStart() bool {
	return false
}

func (p Pair) isStop() bool {
	return false
}

func (p Pair) getMultiple() int {
	return p.left * p.right
}

type Instruction struct {
	stop  bool
	start bool
}

func (i Instruction) isStart() bool {
	return i.start
}

func (i Instruction) isStop() bool {
	return i.stop
}

func (i Instruction) getMultiple() int {
	return 0
}

type Token interface {
	isStart() bool
	isStop() bool
	getMultiple() int
}

func GeneratePair(inputs []string) Token {
	if inputs[0] == "don't()" {
		return Instruction{stop: true, start: false}
	}
	if inputs[0] == "do()" {
		return Instruction{stop: false, start: true}
	}
	left, _ := strconv.Atoi(inputs[1])
	right, _ := strconv.Atoi(inputs[2])
	return Pair{left: left, right: right}
}

func ParseString(input string) {
	regexp := regexp2.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := regexp.FindAllStringSubmatch(input, -1)
	pairs := make([]Token, 0)
	total := 0
	for _, match := range matches {
		newPair := GeneratePair(match)
		total += newPair.getMultiple()
		pairs = append(pairs, newPair)
	}
	fmt.Println(matches)
	fmt.Println(pairs)
	fmt.Println("Total of Multiples is", total)
}

func ParseStringV2(input string) {
	regexp := regexp2.MustCompile(`(?:mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)
	matches := regexp.FindAllStringSubmatch(input, -1)
	fmt.Println("MATCHES", matches)
	pairs := make([]Token, 0)
	total := 0
	enabled := true
	for _, match := range matches {
		newPair := GeneratePair(match)
		pairs = append(pairs, newPair)
		if newPair.isStart() {
			enabled = true
		}
		if newPair.isStop() {
			enabled = false
		}
		if !enabled {
			continue
		}
		total += newPair.getMultiple()
	}

	fmt.Println(pairs)
	fmt.Println("Total of Multiples is", total)
}
