package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var testFileName = "test.txt"
var test2FileName = "test2.txt"
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
	points          []Coordinate
	id              string
	length          int
	perimeterLength []Segment
	segments        [][]Segment
}

type Segment struct {
	interior, exterior Coordinate
}

func (s *Segment) Horizontal() bool {
	return s.interior.x == s.exterior.x
}

func (s *Segment) Vertical() bool {
	return !s.Horizontal()
}

func (s *Segment) IsContinuing(other *Segment) bool {
	// if vertical, next segment msut be in same X/col, and up or down Y by 1
	if s.Vertical() != other.Vertical() {
		return false
	}
	if s.Vertical() && (s.interior.x == other.interior.x && s.exterior.x == other.exterior.x) {
		return (s.exterior.y == other.exterior.y-1 && s.interior.y == other.interior.y-1) ||
			(s.exterior.y == other.exterior.y+1 && s.interior.y == other.interior.y+1)
	} else if s.Horizontal() && s.interior.y == other.interior.y && s.exterior.y == other.exterior.y {
		return (s.interior.x == other.interior.x-1 && s.exterior.x == other.exterior.x-1) ||
			(s.interior.x == other.interior.x+1 && s.exterior.x == other.exterior.x+1)
	}
	return false
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
	perimeterSegments := make([]Segment, 0)
	for _, point := range r.points {
		for _, dir := range directions {
			nr, nc := point.y+dir.dy, point.x+dir.dx
			value := grid[point.y][point.x]
			if grid[nr] == nil {
				//fmt.Printf("For Point %+v, Side: %d %d value %s\n", point, nc, nr, "NO VALUE")
				perimeterSegments = append(perimeterSegments, Segment{interior: point, exterior: Coordinate{x: nc, y: nr}})
				continue
			}
			if (grid[nr] != nil) && grid[nr][nc] != value {
				//fmt.Printf("For Point %+v, Side: %d %d value %s\n", point, nc, nr, grid[nr][nc])
				perimeterSegments = append(perimeterSegments, Segment{interior: point, exterior: Coordinate{x: nc, y: nr}})
				continue
			}
		}
	}
	r.length = len(perimeterSegments)
	r.perimeterLength = perimeterSegments
	return len(perimeterSegments)
}

func (r *Region) PerimeterSides() int {
	if r.perimeterLength == nil {
		r.PerimeterLength()
	}
	segments := make([][]Segment, 0)
	visitedSegments := make(map[Segment]bool)
	for _, s := range r.perimeterLength {
		if visitedSegments[s] {
			continue
		}
		segments = append(segments, make([]Segment, 0))
		segments[len(segments)-1] = append(segments[len(segments)-1], s)
		for _, other := range r.perimeterLength {
			if slices.IndexFunc(segments[len(segments)-1], func(s Segment) bool { return s.IsContinuing(&other) }) != -1 {
				segments[len(segments)-1] = append(segments[len(segments)-1], other)
				visitedSegments[other] = true
			}
		}
	}
	r.segments = segments
	return len(segments)
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
	totalPricePt2 := 0
	for _, r := range regions {
		area := r.Area()
		total += area
		price := r.Price()
		perimeter := r.PerimeterLength()
		totalPrice += price
		sides := r.PerimeterSides()
		pricePt2 := sides * area
		totalPricePt2 += pricePt2
		fmt.Printf("Region: %+v, Area: %d Perimeter: %d Price: %d PT2 %d, Sides: %d Detail: %+v\n", r.id, area, perimeter, price, pricePt2, sides, r)
	}
	fmt.Println(total, totalPrice, totalPricePt2)
	fmt.Println("------")
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
	//HandleFile(testFileName)
	//HandleFile(test2FileName)
	HandleFile(actualFileName)
}
