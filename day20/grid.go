package main

import (
	"fmt"
	"github.com/fatih/color"
	heap2 "github.com/richseviora/aoc2024/day18/heap"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y int
}

type Grid struct {
	cells      map[Coordinate]string
	maxX, maxY int
	startCell  Coordinate
	endCell    Coordinate
}

var enableDetail = os.Getenv("ENABLE_DETAIL") == "true"

func NewGridFromInput(input string) *Grid {
	lines := strings.Split(input, "\n")

	cells := make(map[Coordinate]string)
	maxY := 0
	maxX := 0
	var startCell Coordinate
	var endCell Coordinate

	for i, line := range lines {
		if i > maxY {
			maxY = i
		}
		for j, char := range line {
			if j > maxX {
				maxX = j
			}
			coordinate := Coordinate{x: j, y: i}

			s := string(char)
			if s == "S" {
				startCell = coordinate
			} else if s == "E" {
				endCell = coordinate
			}
			cells[coordinate] = s
		}
	}
	grid := &Grid{
		cells:     cells,
		maxX:      maxX,
		maxY:      maxY,
		startCell: startCell,
		endCell:   endCell,
	}
	return grid
}

func (g *Grid) GetPossibleMoves(c Coordinate) []Coordinate {
	var validMoves []Coordinate

	// Define the possible directions: up, down, left, right
	directions := []Coordinate{
		{x: 0, y: -1}, // Up
		{x: 0, y: 1},  // Down
		{x: -1, y: 0}, // Left
		{x: 1, y: 0},  // Right
	}

	// Iterate over each direction to check adjacent cells
	for _, dir := range directions {
		adjacent := Coordinate{
			x: c.x + dir.x,
			y: c.y + dir.y,
		}

		// Check if the adjacent cell is valid and has a value of "."
		if cellValue, exists := g.cells[adjacent]; exists && (cellValue == "." || cellValue == "S" || cellValue == "E") {
			validMoves = append(validMoves, adjacent)
		}
	}

	return validMoves
}

func (g *Grid) GetShortestPath(scores map[Coordinate]int) []Coordinate {
	var path []Coordinate

	// Start from the high cell
	current := g.endCell

	for {
		path = append(path, current)

		// If we've reached the end cell with a score of 0, stop
		if scores[current] == 0 {
			break
		}

		// Look at all adjacent cells
		directions := []Coordinate{
			{x: 0, y: -1}, // Up
			{x: 0, y: 1},  // Down
			{x: -1, y: 0}, // Left
			{x: 1, y: 0},  // Right
		}

		// Find the adjacent cell where the score is exactly 1 less
		for _, dir := range directions {
			adjacent := Coordinate{
				x: current.x + dir.x,
				y: current.y + dir.y,
			}
			if score, exists := scores[adjacent]; exists && score == scores[current]-1 {
				current = adjacent
				break
			}
		}
	}
	//fmt.Printf("PATH: %+v\n", path)

	return path
}

func (g *Grid) GetDistancesFromCellToCell(startCell, endCell Coordinate) (int, map[Coordinate]int) {
	heap := heap2.NewHeapQueue[Coordinate]()
	heap.Push(startCell)
	cellScores := map[Coordinate]int{
		startCell: 0,
	}
	for heap.Len() > 0 {

		coord := heap.PopSafe()
		if reflect.ValueOf(coord).IsZero() {
			break
		}

		cellScore, _ := cellScores[coord]
		if enableDetail {
			fmt.Printf("EVAL: %+v SCORE: %d\n", coord, cellScore)
		}

		if coord == endCell {
			continue
		}
		possibleMoves := g.GetPossibleMoves(coord)
		for _, possibleMove := range possibleMoves {
			scoreDelta := ScoreSuggestion(coord, possibleMove)
			newScore := 1 + cellScore
			priority := newScore - scoreDelta
			if enableDetail {
				fmt.Printf("TEST: %+v SCORE: %d\n", possibleMove, newScore)
			}

			if existingScore, ok := cellScores[possibleMove]; (ok && existingScore > newScore) || !ok {
				cellScores[possibleMove] = newScore
				heap.Upsert(possibleMove, priority)
				if enableDetail {
					fmt.Printf("QUEUE: %+v SCORE: %d HEAP SIZE: %d\n", possibleMove, newScore, heap.Len())
				}
			} else {
				if enableDetail {
					fmt.Printf("SKIP: %+v SCORE: %d\n", possibleMove, newScore)
				}
			}
		}
	}
	//if enableDetail {
	//g.PrintCellWithScore(cellScores)
	//}

	result := cellScores[endCell]
	//fmt.Printf("RESULT: %d\n", result)
	return result, cellScores
}

func (g *Grid) PrintCellWithScoreAndSkips(scores map[Coordinate]int, hops []Hop) {
	maxScore := 0
	for _, score := range scores {
		if score > maxScore {
			maxScore = score
		}
	}
	cells := make(map[Coordinate]string)
	for coord, cell := range scores {
		cells[coord] = g.FormatCell(cell, coord, maxScore, hops)
	}
	g.PrintGrid(cells)
}

func (g *Grid) PrintCellWithScore(scores map[Coordinate]int) {
	maxScore := 0
	for _, score := range scores {
		if score > maxScore {
			maxScore = score
		}
	}
	cells := make(map[Coordinate]string)
	for coord, cell := range scores {
		cells[coord] = g.FormatCell(cell, coord, maxScore, nil)
	}
	g.PrintGrid(cells)
}

func (g *Grid) FormatCell(s int, c Coordinate, maxScore int, hops []Hop) string {
	isStart := slices.IndexFunc(hops, func(hop Hop) bool {
		return hop.from == c
	}) >= 0
	isEnd := slices.IndexFunc(hops, func(hop Hop) bool {
		return hop.to == c
	}) >= 0
	if isStart {
		return color.RGB(0, 128, 255).Sprint("U")
	}
	if isEnd {
		return color.RGB(0, 128, 255).Sprint("D")
	}
	if c == g.endCell {
		return color.RGB(255, 0, 0).Add(color.BlinkSlow).Sprint("E")
	}
	if c == g.startCell {
		return color.RGB(255, 0, 0).Add(color.BlinkSlow).Sprint("S")
	}
	return ColorizeScore(s, maxScore)
}

func ColorizeScore(score int, maxScore int) string {
	itoa := strconv.Itoa(score)
	txt := itoa[len(itoa)-1 : len(itoa)]

	// Calculate the intensity of red and green based on the score relative to maxScore
	// maxScore becomes fully red, score 0 is fully green
	red := (float32(score) * 255) / float32(maxScore)
	green := 255 - red

	// Create the color with dynamic red and green, no blue component
	return color.RGB(int(red), int(green), 0).Sprint(txt)
}

func (g *Grid) PrintGrid(path map[Coordinate]string) {
	fmt.Println("GRID >>>")
	for y := 0; y <= g.maxY; y++ {
		for x := 0; x <= g.maxX; x++ {
			c := Coordinate{x: x, y: y}
			content, ok := g.cells[c]
			if ok == false {
				panic("invalid coordinate")
			}
			override, ok := path[c]
			if ok {
				fmt.Print(override)
			} else {
				fmt.Print(content)
			}
		}
		fmt.Println()
	}
	fmt.Println(">>> GRID")
}

func ScoreSuggestion(origin, target Coordinate) int {
	return ((target.x - origin.x) + (target.y - origin.y)) * -1
}
