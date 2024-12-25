package main

import (
	"math"
)

type Hop struct {
	from    Coordinate
	to      Coordinate
	savings int
}

func (h *Hop) Distance() int {
	return int(math.Abs(float64(h.from.x-h.to.x)) + math.Abs(float64(h.from.y-h.to.y)))
}

func (g *Grid) CalculateLongSkips(grid1, grid2 map[Coordinate]int, threshold int) []Hop {
	// get the maximum score here so we know if we're beating it
	hopLength := 21
	maxValue := math.MinInt
	for _, value := range grid1 {
		if value > maxValue {
			maxValue = value
		}
	}
	getCoords := func(coordinate Coordinate) []Coordinate {
		coords := []Coordinate{}
		for x := -hopLength; x <= hopLength; x++ {
			for y := -hopLength; y <= hopLength; y++ {
				if math.Abs(float64(x))+math.Abs(float64(y)) >= float64(hopLength) {
					continue
				}
				coord := Coordinate{x: coordinate.x + x, y: coordinate.y + y}
				coords = append(coords, coord)
			}
		}
		return coords
	}
	hops := []Hop{}
	for coord1, value1 := range grid1 {
		possibleCoords := getCoords(coord1)
		for _, neighbor := range possibleCoords {
			if emptyContent, ok := g.cells[neighbor]; !ok || ok && emptyContent == "#" {
				continue
			}
			if value2, exists := grid2[neighbor]; exists {
				totalDistance := value1 + value2
				cost := math.Abs(float64(coord1.x-neighbor.x)) + math.Abs(float64(coord1.y-neighbor.y))
				savings := maxValue - totalDistance - int(cost)
				if savings >= threshold {
					hop := Hop{from: coord1, to: neighbor, savings: savings}
					//fmt.Printf("HOP FOUND: %+v, DIST: %d\n", hop, hop.Distance())
					hops = append(hops, hop)
				}
			}
		}
	}
	return hops
}

func (g *Grid) CalculateSkips(grid1, grid2 map[Coordinate]int, threshold int, pt2Mode bool) []Hop {
	if pt2Mode {
		return g.CalculateLongSkips(grid1, grid2, threshold)
	}
	// get the maximum score here so we know if we're beating it
	maxValue := math.MinInt
	for _, value := range grid1 {
		if value > maxValue {
			maxValue = value
		}
	}
	getCoords := func(coordinate, deltaCoordinate Coordinate) (Coordinate, Coordinate) {
		wallCheck := Coordinate{x: coordinate.x + deltaCoordinate.x, y: coordinate.y + deltaCoordinate.y}
		emptyCheck := Coordinate{x: coordinate.x + deltaCoordinate.x*2, y: coordinate.y + deltaCoordinate.y*2}
		return wallCheck, emptyCheck
	}
	hops := []Hop{}
	for coord1, value1 := range grid1 {
		ordinalDirections := []Coordinate{
			{x: +1, y: 0},
			{x: -1, y: 0},
			{x: 0, y: +1},
			{x: 0, y: -1},
		}
		for _, delta := range ordinalDirections {
			wallCoord, neighbor := getCoords(coord1, delta)
			if wallCell, ok := g.cells[wallCoord]; !ok || ok && wallCell != "#" {
				continue
			}
			if emptyContent, ok := g.cells[neighbor]; !ok || ok && emptyContent == "#" {
				continue
			}
			if value2, exists := grid2[neighbor]; exists {
				totalDistance := value1 + value2
				if maxValue-totalDistance >= threshold {
					hop := Hop{from: coord1, to: neighbor, savings: maxValue - totalDistance}
					hops = append(hops, hop)
				}
			}
		}
	}
	return hops
}
