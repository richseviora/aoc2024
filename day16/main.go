package main

import (
	"github.com/fatih/color"
	"os"
)

var testFileName = "test1.txt"
var test2FileName = "test2.txt"
var test3FileName = "test3.txt"
var actualFileName = "input.txt"

func ProcessChallenge(input string, pt2 bool) int {
	return 0
}

func ReadInput(fname string) string {
	file, err := os.Open(fname)
	defer file.Close()
	if err != nil {
		color.Red("error received: %e", err)
		panic(err)
	}

	content, err := os.ReadFile(fname)
	if err != nil {
		color.Red("error received: %v", err)
		panic(err)
	}
	return string(content)

}

func HandleFile(fname string, expected int) {
	fileContent := ReadInput(fname)
	for _, pt2 := range []bool{true} {
		color.Yellow("Filename: %s, Part 2: %t, File Length: %d", fname, pt2, len(fileContent))
		result := ProcessChallenge(fileContent, pt2)
		if result == expected {
			color.Green("PASSED")
		} else {
			color.Red("FAILED")
		}
	}
}

func main() {
	HandleFile(testFileName, 7036)
}
