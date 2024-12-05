package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	precedingPage int
	followingPage int
}

func NewRule(input string) Rule {
	fmt.Println("CONVERTING", input)
	rules := strings.Split(input, "|")
	precedingPage, _ := strconv.Atoi(rules[0])
	followingPage, _ := strconv.Atoi(rules[1])
	return Rule{
		precedingPage: precedingPage,
		followingPage: followingPage,
	}
}

func (r Rule) IsValidUpdate(updates []int) bool {
	followingIndex := slices.Index(updates, r.followingPage)
	precedingIndex := slices.Index(updates, r.precedingPage)
	fmt.Printf("Comparing P %d to F %d for %d\n", r.precedingPage, r.followingPage, updates)
	if followingIndex == -1 || precedingIndex == -1 {
		return true
	}
	if precedingIndex > followingIndex {
		return false
	}
	return true
}
