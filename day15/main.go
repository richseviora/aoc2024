package main

import (
	"fmt"
	"os"
)

var testFileName = "test.txt"
var test2FileName = "test2.txt"
var test3FileName = "test3.txt"
var actualFileName = "input.txt"

func ProcessChallenge(input string, pt2 bool) {
	grid, directions := NewGridFromInput(input, pt2)
	fmt.Println("Grid:\n", grid.ToString())
	for i, direction := range directions {
		valid := grid.HandleDirection(direction)
		fmt.Printf("Step: %d in direction %d VALID: %t \nGrid:\n%s\n", i, direction, valid, grid.ToString())
	}
	fmt.Println("Final Grid Value: ", grid.GetBoxGPSTotal())
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
	for _, pt2 := range []bool{true} {
		fmt.Println(fname, pt2, len(fileContent))
		ProcessChallenge(fileContent, pt2)
	}
}

func main() {
	HandleFile(actualFileName)
}
