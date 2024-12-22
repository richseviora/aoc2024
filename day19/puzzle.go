package main

import (
	"fmt"
	"github.com/kofalt/go-memoize"
	"strings"
)

type Puzzle struct {
	Sequences []Sequence
	Requests  []Request
	Memoizer  *memoize.Memoizer
}

func NewPuzzle(input string) Puzzle {
	lines := strings.Split(input, "\n")
	sequences := make([]Sequence, 0)
	requests := make([]Request, 0)
	for i, line := range lines {
		if i == 0 {
			sequenceStrings := strings.Split(line, ", ")
			for _, sequenceString := range sequenceStrings {
				sequences = append(sequences, NewSequence(sequenceString))
			}
		} else if i > 1 {
			requests = append(requests, Request{Pattern: line})
		}
	}
	return Puzzle{
		Sequences: sequences,
		Requests:  requests,
	}
}

func (p *Puzzle) GetPossiblePatterns() int {
	solutions := make(map[Request]int)
	totalOptions := 0
	for _, request := range p.Requests {
		fmt.Printf("REQUEST EVAL BEGIN: %s\n", request.Pattern)
		solution := CanBeComposedFromMemoized(request.Pattern, p.Sequences)
		if solution > 0 {
			solutions[request] = solution
			totalOptions += solution
		}
		fmt.Printf("REQUEST EVAL END: %s %d\n", request.Pattern, solution)
	}
	fmt.Printf("TOTAL OPTIONS: %d\n", totalOptions)
	return len(solutions)
}
