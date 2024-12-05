package main

import (
	"strconv"
	"strings"
)

type PageUpdate struct {
	updates []int
	rules   []Rule
}

func NewUpdate(input string, rules []Rule) PageUpdate {
	inputSlice := strings.Split(input, ",")
	inputInts := make([]int, 0)
	for _, i := range inputSlice {
		inputInt, _ := strconv.Atoi(i)
		inputInts = append(inputInts, inputInt)
	}
	return PageUpdate{updates: inputInts, rules: rules}
}

func (p PageUpdate) IsValid() bool {
	for _, rule := range p.rules {
		if !rule.IsValidUpdate(p.updates) {
			return false
		}
	}
	return true
}

func (p PageUpdate) GetMiddleUpdate() int {
	middle := len(p.updates) / 2
	return p.updates[middle]
}
