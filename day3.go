package main

import (
	"aoc2024/m/v2/day3"
	"fmt"
	"os"
)

const day3Name = "day3/input.txt"

func readInput(fname string) string {
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
	fmt.Println("Day 3 Begin")
	result := readInput(day3Name)
	day3.ParseString(result)
}
