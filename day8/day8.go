package main

import (
	"fmt"
	"github.com/mowshon/iterium"
	"math"
	"os"
	"slices"
	"strings"
)

type Grid struct {
	Antennas    []Antenna
	Height      int
	Width       int
	Frequencies map[string]bool
}

type Antenna struct {
	Frequency  string
	Coordinate Coordinate
}

type Coordinate struct {
	Row int
	Col int
}

func NewTableFromString(input string) Grid {
	nodes := make([]Antenna, 0)
	frequencies := make(map[string]bool)
	rows := strings.Split(input, "\n")
	for ri, row := range rows {
		for ci, value := range strings.Split(row, "") {
			if value != "." && value != "#" {
				if _, exists := frequencies[value]; !exists {
					frequencies[value] = false
				}
				nodes = append(nodes, Antenna{
					Frequency: value,
					Coordinate: Coordinate{
						Row: ri,
						Col: ci,
					},
				})
			}
		}
	}
	return Grid{nodes, len(rows), len(rows[0]), frequencies}
}

func (g *Grid) InBounds(c Coordinate) bool {
	return c.Row >= 0 && c.Row < g.Height && c.Col >= 0 && c.Col < g.Width
}

func (g *Grid) GetAntennasForFrequency(value string) []Antenna {
	antennas := make([]Antenna, 0)
	for _, a := range g.Antennas {
		if a.Frequency == value {
			antennas = append(antennas, a)
		}
	}
	return antennas
}

func (g *Grid) GetAntiNodesForFrequency(v string) []Coordinate {
	coordinates := make([]Coordinate, 0)
	antennas := g.GetAntennasForFrequency(v)
	pairs, err := iterium.Combinations(antennas, 2).Slice()
	if err != nil {
		panic(err)
	}
	for _, pair := range pairs {
		newCoordinates := CalculateAntiNodes(pair[0].Coordinate, pair[1].Coordinate, g)
		coordinates = append(coordinates, newCoordinates...)
	}
	return coordinates
}

func (g *Grid) GetAntiNodesForAllFrequencies() []Coordinate {
	coordinates := make([]Coordinate, 0)

	for coord := range g.Frequencies {
		coordinates = append(coordinates, g.GetAntiNodesForFrequency(coord)...)
	}
	seen := make(map[Coordinate]bool)
	uniques := []Coordinate{}
	for _, value := range coordinates {
		if !seen[value] {
			uniques = append(uniques, value)
			seen[value] = true
		}
	}

	return uniques
}

func CalculateAntiNodes(a, b Coordinate, g *Grid) []Coordinate {
	deltaY := float64(b.Row - a.Row)
	deltaX := float64(b.Col - a.Col)
	behind := Coordinate{
		Row: int(float64(a.Row) - deltaY),
		Col: int(float64(a.Col) - deltaX),
	}
	after := Coordinate{
		Row: int(float64(b.Row) + deltaY),
		Col: int(float64(b.Col) + deltaX),
	}
	coordinates := []Coordinate{behind, after}
	deltaY3rd := deltaY / 3
	deltaX3rd := deltaX / 3
	if math.Mod(deltaY3rd, 1) == 0 && math.Mod(deltaX3rd, 1) == 0 {
		mid1 := Coordinate{
			Row: int(deltaY3rd) + a.Row,
			Col: int(deltaX3rd) + a.Col,
		}
		mid2 := Coordinate{
			Row: int(deltaY3rd) + (a.Row * 2),
			Col: int(deltaX3rd) + (a.Col * 2),
		}
		coordinates = append(coordinates, mid1, mid2)
	}
	validValues := slices.DeleteFunc(coordinates, func(c Coordinate) bool {
		return !g.InBounds(c)
	})
	return validValues
}

const filename = "input.txt"

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

func main() {
	contents := ReadInput(filename)
	table := NewTableFromString(contents)
	uniqueCoords := table.GetAntiNodesForAllFrequencies()
	fmt.Println(len(uniqueCoords))
}
