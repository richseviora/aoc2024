package main

import (
	"github.com/fatih/color"
	"os"
)

func ProcessChallenge(fname string, threshold int) int {
	input := ReadInput(fname)
	g := NewGridFromInput(input)
	_, scoreFromStart := g.GetDistancesFromCellToCell(g.startCell, g.endCell)
	_, scoreFromEnd := g.GetDistancesFromCellToCell(g.endCell, g.startCell)
	cheats := g.CalculateSkips(scoreFromStart, scoreFromEnd, threshold)
	g.PrintCellWithScoreAndSkips(scoreFromStart, cheats)
	return len(cheats)
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

func HandleFile(fname string, expected int, threshold int) {

	for _, pt2 := range []bool{true} {
		color.Yellow("Filename: %s, Part 2: %t", fname, pt2)
		result := ProcessChallenge(fname, threshold)
		if result == expected {
			color.Green("PASSED - Expected: %d, Actual: %d", expected, result)
		} else {
			color.Red("FAILED - Expected: %d, Actual: %d", expected, result)
		}
	}
}

func main() {
	HandleFile("input.txt", 1024, 102)
}
