package main

import (
	"aoc2024/m/v2/day2"
	"bufio"
	"fmt"
	"os"
)

const fileName2 = "day2/input.txt"

func main() {
	// read file, process values,
	fmt.Println("Day 2 - with Dampener")

	file, err := os.Open(fileName2)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	safeReports := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		report := day2.GenerateReportsFromStr(line, true)
		if report.IsSafe() {
			safeReports++
		}
	}

	fmt.Println("Total Count of Safe Reports:", safeReports, "")

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
