package main

import (
	"fmt"
	"os"
)

var testFileName = "test.txt"
var actualFileName = "input.txt"

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

func HandleInput(input string) {
	grid := NewGridFromString(input)
	trailheads := grid.FindTrailheads()
	totalCount := 0
	totalRoutes := 0
	for _, t := range trailheads {
		routes := t.GetRoutesEndingWithValueNine()
		uniqueLastCells := t.GetUniqueLastCells()
		totalCount += len(uniqueLastCells)
		totalRoutes += len(routes)
		fmt.Printf("Cell: %+v, UniqueCellCount: %d Unique Routes: %d\n", t, len(uniqueLastCells), len(routes))
	}
	fmt.Println("TotalCount", totalCount, "Routes", totalRoutes)
}

func HandleFile(fname string) {
	fileContent := ReadInput(fname)
	for _, pt2 := range []bool{false} {
		fmt.Println("result", fname, pt2, len(fileContent))
		HandleInput(fileContent)
		fmt.Println("----------------------------------------")
	}

}

func main() {
	HandleFile(testFileName)
	HandleFile(actualFileName)
}
