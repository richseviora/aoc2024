package main

import (
	"math"
)

type Hop struct {
	from Coordinate
	to   Coordinate
}

func (g *Grid) CalculateSkips(grid1, grid2 map[Coordinate]int, threshold int) []Hop {
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
					hop := Hop{from: coord1, to: neighbor}
					hops = append(hops, hop)
				}
			}
		}
	}
	return hops
}
