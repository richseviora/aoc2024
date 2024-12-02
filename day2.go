package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const fileName = "day2/input.txt"

func processLine(line string) {
	fields := strings.Fields(line)

	fmt.Println("Split line:", fields)
	fmt.Println("Read line:", line)
}

func main() {
	// read file, process values,
	fmt.Println("Day 2")

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		fmt.Println("Split line:", fields)
		fmt.Println("Read line:", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
