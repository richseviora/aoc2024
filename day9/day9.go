package main

import (
	"fmt"
	"os"
)

const testFileName = "test.txt"
const actualFileName = "input.txt"

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

}

func main() {
	HandleFile(testFileName)
	HandleFile(actualFileName)
}
