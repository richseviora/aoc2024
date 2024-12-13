package main

import (
	"fmt"
	"os"
	"slices"
)

var testFileName = "test.txt"
var actualFileName = "input.txt"

func ProcessChallenge(input string) {
	parameters := ParseInput(input)
	totalCost := 0
	for _, parameter := range parameters {
		solutions := parameter.Solutions()
		if len(solutions) == 0 {
			continue
		}
		slices.SortFunc(solutions, func(a, b Solution) int {
			return a.Cost() - b.Cost()
		})
		cheapest := solutions[0]
		fmt.Printf("cheapest cost: %d for solution:%+v, parameter: %+v\n", cheapest.Cost(), cheapest, parameter)
		totalCost += cheapest.Cost()
	}
	fmt.Printf("Total Cost: %d\n", totalCost)
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
	//HandleFile(test2FileName)
	HandleFile(actualFileName)
}
