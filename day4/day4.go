package main

import (
	"fmt"
	"os"
)

const day4Name = "input.txt"

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
	input := ReadInput(day4Name)
	table := NewTable(input)
	results := table.IterateOverTablePart2()
	fmt.Println(results)
	fmt.Println(len(results))
}
