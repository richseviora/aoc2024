package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

func HandleFile(fname string) {
	fileContent := ReadInput(fname)
	for _, pt2 := range []bool{false} {
		fmt.Println(fname, pt2, iterations, len(fileContent))
		ProcessChallenge(fileContent, iterations)
	}
}

func main() {
	HandleFile(testFileName)
	HandleFile(actualFileName)
}
