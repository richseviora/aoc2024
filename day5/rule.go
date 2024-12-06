package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Rule struct {
	precedingPage  int
	followingPages []int
}

func NewRule(input string) Rule {
	fmt.Println("CONVERTING", input)
	rules := strings.Split(input, "|")
	precedingPage, _ := strconv.Atoi(rules[0])
	followingPage, _ := strconv.Atoi(rules[1])
	return Rule{
		precedingPage:  precedingPage,
		followingPages: []int{followingPage},
	}
}
