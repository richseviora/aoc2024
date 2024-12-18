package main

import (
	"github.com/fatih/color"
	"os"
)

func ProcessChallenge(input string, limit int) int {
	grid := NewGrid(71, 71)
	grid.PopulateGridFromInput(input, limit)
	grid.PrintGrid([]*Cell{})
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

func HandleFile(fname string, limit, expected int) {
	fileContent := ReadInput(fname)
	for _, pt2 := range []bool{true} {
		color.Yellow("Filename: %s, Part 2: %t, File Length: %d", fname, pt2, len(fileContent))
		result := ProcessChallenge(fileContent, limit)
		if result == expected {
			color.Green("PASSED")
		} else {
			color.Red("FAILED")
		}
	}
}

func main() {
	HandleFile("input.txt", 1024, 124)
}
