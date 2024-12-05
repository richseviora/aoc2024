package main

import (
	"fmt"
	"os"
	"regexp"
)

const day5Name = "input.txt"

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
	input := ReadInput(day5Name)
	rulesReg := regexp.MustCompile("\\d+\\|\\d+")
	rulesMatches := rulesReg.FindAllString(input, -1)
	rules := make([]Rule, 0)
	for _, match := range rulesMatches {
		rules = append(rules, NewRule(match))
	}

	updatesReg := regexp.MustCompile("(?:\\d+,)+\\d+")
	updates := make([]PageUpdate, 0)
	updatesMatches := updatesReg.FindAllString(input, -1)
	for _, match := range updatesMatches {
		updates = append(updates, NewUpdate(match, rules))
	}
	total := 0
	for _, update := range updates {
		if update.IsValid() {
			fmt.Println("Valid Update", update)
			total += update.GetMiddleUpdate()
		}
	}
	fmt.Println(total)
}
