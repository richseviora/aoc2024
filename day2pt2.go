package main

import (
	"aoc2024/m/v2/day2"
	"bufio"
	"fmt"
	"os"
)

func main() {
	// read file, process values,
	fmt.Println("Day 2 - with Dampener")

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	safeReports := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		report := day2.GenerateReportsFromStr(line)
		if report.IsSafe() {
			safeReports++
		}
	}

	fmt.Println("Total Count of Safe Reports:", safeReports, "")

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
