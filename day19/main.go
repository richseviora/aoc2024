package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func ProcessChallenge(fname string) int {
	input := ReadInput(fname)
	puzzle := NewPuzzle(input)
	fmt.Printf("Result: %+v\n", puzzle)
	return 0
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

func HandleFile(fname string, expected int) {
	for _, pt2 := range []bool{true} {
		color.Yellow("Filename: %s, Part 2: %t", fname, pt2)
		result := ProcessChallenge(fname)
		if result == expected {
			color.Green("PASSED - Expected: %d, Actual: %d", expected, result)
		} else {
			color.Red("FAILED - Expected: %d, Actual: %d", expected, result)
		}
	}
}

func main() {
	HandleFile("test1.txt", 6)
}
