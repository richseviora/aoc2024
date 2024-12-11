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

func ProcessInput(input string) []int {
	strings := strings.Split(input, " ")
	ints := make([]int, len(strings))
	for i, s := range strings {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints

}

var memo = make(map[int][]int)

func MemoizedCalculation(n int) []int {
	newStones := []int{}
	digits := fmt.Sprintf("%d", n)
	if n == 0 {
		newStones = append(newStones, 1)
	} else if len(digits)%2 == 0 {
		leftStoneStr := digits[0 : len(digits)/2]
		rightStoneStr := digits[len(digits)/2:]
		leftStone, _ := strconv.Atoi(leftStoneStr)
		rightStone, _ := strconv.Atoi(rightStoneStr)
		newStones = append(newStones, leftStone)
		newStones = append(newStones, rightStone)
	} else {
		newStones = append(newStones, n*2024)
	}
	return newStones
}

func Calculation(n int) []int {
	if memo[n] != nil {
		return memo[n]
	}
	newStones := MemoizedCalculation(n)
	memo[n] = newStones
	return newStones
}

func HandleBlink(stones []int) []int {
	newStones := make([]int, 0)
	for _, s := range stones {
		newStones = append(newStones, Calculation(s)...)
	}
	return newStones
}

func ProcessChallenge(input string, iterations int) {
	stones := ProcessInput(input)
	fmt.Printf("Stones Before %d, Count: %d\n", iterations, len(stones))
	for i := 0; i < iterations; i++ {
		stones = HandleBlink(stones)
		fmt.Printf("Stones On %d, Count: %d\n", i, len(stones))
	}
	fmt.Printf("Stones After %d, Count: %d\n", iterations, len(stones))
}

func HandleFile(fname string, iterations int) {
	fileContent := ReadInput(fname)
	for _, pt2 := range []bool{false} {
		fmt.Println(fname, pt2, iterations, len(fileContent))
		ProcessChallenge(fileContent, iterations)
	}

}

func main() {
	HandleFile(testFileName, 6)
	HandleFile(actualFileName, 75)
}
