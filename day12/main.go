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
	id     string
	length int
}

func (r *Region) Area() int {
	return len(r.points)
}

func (r *Region) Price() int {
	return r.Area() * r.PerimeterLength()
}

func (r *Region) PerimeterLength() int {
	if r.length != 0 {
		return r.length
	}
	sides := 0
	for _, point := range r.points {
		for _, dir := range directions {
			nr, nc := point.y+dir.dy, point.x+dir.dx
			value := grid[point.y][point.x]
			if grid[nr] == nil {
				fmt.Printf("For Point %+v, Side: %d %d value %s\n", point, nc, nr, "NO VALUE")
				sides++
			}
			if (grid[nr] != nil) && grid[nr][nc] != value {
				fmt.Printf("For Point %+v, Side: %d %d value %s\n", point, nc, nr, grid[nr][nc])
				sides++
			}
		}
	}
	r.length = sides
	return sides
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
				value := grid[ri][ci]
				region := Region{id: value}
				queue := [][]int{{ri, ci}}
				visited[ri][ci] = true

				for len(queue) > 0 {
					point := queue[0]
					queue = queue[1:]
					r, c := point[0], point[1]
					region.points = append(region.points, Coordinate{x: c, y: r})

					// Check all 4 possible directions
					for _, dir := range directions {
						nr, nc := r+dir.dy, c+dir.dx

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
	total := 0
	totalPrice := 0
	for _, r := range regions {
		area := r.Area()
		total += area
		price := r.Price()
		perimeter := r.PerimeterLength()
		totalPrice += price
		fmt.Printf("Region: %+v, Area: %d Perimeter: %d Price: %d, Detail: %+v\n", r.id, area, perimeter, price, r)
	}
	fmt.Println(total, totalPrice)
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
	//HandleFile(actualFileName)
}
