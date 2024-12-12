package main

import (
	"fmt"
	"os"
	"strings"
)

var testFileName = "test.txt"
var actualFileName = "input.txt"

var grid = make(map[int]map[int]string)

func ProcessInput(input string) map[int]map[int]string {
	for ri, line := range strings.Split(input, "\n") {
		for ci, c := range line {
			if grid[ri] == nil {
				grid[ri] = make(map[int]string)
			}
			grid[ri][ci] = string(c)
		}
	}
	return grid
}

var directions = []struct{ dx, dy int }{
	{dx: -1, dy: 0}, // Up
	{dx: 1, dy: 0},  // Down
	{dx: 0, dy: -1}, // Left
	{dx: 0, dy: 1},  // Right
}

type Coordinate struct {
	x, y int
}

type Region struct {
	points []Coordinate
}

func FindContiguousRegions(grid map[int]map[int]string) []Region {
	visited := make(map[int]map[int]bool)
	var regions []Region

	for ri := range grid {
		for ci := range grid[ri] {
			if visited[ri] == nil {
				visited[ri] = make(map[int]bool)
			}

			if !visited[ri][ci] {
				// Perform BFS starting from this cell
				region := Region{}
				value := grid[ri][ci]
				queue := [][]int{{ri, ci}}
				visited[ri][ci] = true

				for len(queue) > 0 {
					point := queue[0]
					queue = queue[1:]
					r, c := point[0], point[1]
					region.points = append(region.points, Coordinate{x: r, y: c})

					// Check all 4 possible directions
					for _, dir := range directions {
						nr, nc := r+dir.dx, c+dir.dy

						if grid[nr] != nil && grid[nr][nc] != "" && !visited[nr][nc] && grid[nr][nc] == value {
							if visited[nr] == nil {
								visited[nr] = make(map[int]bool)
							}
							queue = append(queue, []int{nr, nc})
							visited[nr][nc] = true
						}
					}
				}
				regions = append(regions, region)
			}
		}
	}

	return regions
}

func ProcessChallenge(input string) {
	grid := ProcessInput(input)
	regions := FindContiguousRegions(grid)
	fmt.Println(len(regions))
}

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

func HandleFile(fname string) {
	fileContent := ReadInput(fname)
	for _, pt2 := range []bool{false} {
		fmt.Println(fname, pt2, len(fileContent))
		ProcessChallenge(fileContent)
	}
}

func main() {
	HandleFile(testFileName)
	HandleFile(actualFileName)
}
