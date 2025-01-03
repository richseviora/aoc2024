package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func ProcessChallenge(fname string, size, limit int) int {
	input := ReadInput(fname)
	grid := NewGrid(size, size)
	grid.PopulateGridFromInput(input, limit)
	grid.PrintGrid([]*Cell{})
	return grid.GetPathDistanceToEnd()
}

func ProcessChallengePart2(fname string, size, limit int) int {
	input := ReadInput(fname)
	result := FindFirstFailingInput(input, size)
	inputLines := strings.Split(input, "\n")
	fmt.Printf("Failing Input: %s\n", inputLines[result])
	return result
}

func ReadInput(fname string) string {
	file, err := os.Open(fname)
	defer file.Close()
	if err != nil {
		color.Red("error received: %e", err)
		panic(err)
	}

	content, err := os.ReadFile(fname)
	if err != nil {
		color.Red("error received: %v", err)
		panic(err)
	}
	return string(content)

}

func HandleFile(fname string, size, limit, expected int) {

	for _, pt2 := range []bool{true} {
		color.Yellow("Filename: %s, Part 2: %t", fname, pt2)
		result := ProcessChallenge(fname, size, limit)
		if result == expected {
			color.Green("PASSED - Expected: %d, Actual: %d", expected, result)
		} else {
			color.Red("FAILED - Expected: %d, Actual: %d", expected, result)
		}
	}
}

func main() {
	HandleFile("input.txt", 1024, 124, 0)
}
