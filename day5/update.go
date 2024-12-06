package main

import (
	"slices"
	"strconv"
	"strings"
)

type PageUpdate struct {
	updates []int
	rules   map[int][]int
}

func NewUpdate(input string, rules map[int][]int) PageUpdate {
	inputSlice := strings.Split(input, ",")
	inputInts := make([]int, 0)
	for _, i := range inputSlice {
		inputInt, _ := strconv.Atoi(i)
		inputInts = append(inputInts, inputInt)
	}
	return PageUpdate{updates: inputInts, rules: rules}
}

func (p *PageUpdate) IsValid() bool {
	cmp := func(a, b int) int {
		for _, v := range p.rules[b] {
			if v == a {
				return 1
			}
		}
		return -1
	}

	return slices.IsSortedFunc(p.updates, cmp)
}

func (p *PageUpdate) GetMiddleUpdate() int {
	middle := len(p.updates) / 2
	return p.updates[middle]
}

func (p *PageUpdate) GetOrderedUpdate() *PageUpdate {
	if p.IsValid() {
		return p
	}
	newUpdates := make([]int, len(p.updates))
	copy(newUpdates, p.updates)
	cmp := func(a, b int) int {
		for _, v := range p.rules[b] {
			if v == a {
				return 1
			}
		}
		return -1
	}
	slices.SortFunc(newUpdates, cmp)
	return &PageUpdate{updates: newUpdates, rules: p.rules}
}
