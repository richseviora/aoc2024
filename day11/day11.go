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

var pt2Memo = make(map[int]map[int][]int)
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

func MemoizedPt2Calculation(values []int, iterations int) []int {
	fmt.Printf("CALLED %v %d\n", values, iterations)
	var returnValue []int
	defer func() { fmt.Printf("RETURNED  %v %d FOR %v\n", returnValue, iterations, values) }()
	if iterations == 0 {
		returnValue = values
		return returnValue
	}
	if iterations == 1 {
		newValues := make([]int, 0)
		for _, v := range values {
			newValues = append(newValues, Calculation(v)...)
		}
		returnValue = newValues
		return returnValue
	}
	newValues := make([]int, 0)
	for _, v := range values {
		if pt2Memo[v] != nil && pt2Memo[v][iterations] != nil {
			newValues = append(newValues, pt2Memo[v][iterations]...)
			continue
		}
		if pt2Memo[v] == nil {
			pt2Memo[v] = make(map[int][]int)
			pt2Memo[v][0] = []int{v}
		}
		pt2Memo[v][iterations] = MemoizedPt2Calculation([]int{v}, iterations-1)
		//for i := 1; i <= iterations; i++ {
		//	if pt2Memo[v] == nil {
		//		pt2Memo[v] = make(map[int][]int)
		//		pt2Memo[v][0] = []int{v}
		//	}
		//	if pt2Memo[v][i] == nil {
		//		previousValue, _ := pt2Memo[v][i-1]
		//		pt2Memo[v][i] = MemoizedPt2Calculation(previousValue, 1)
		//	}
		//}
		newValues = append(newValues, pt2Memo[v][iterations]...)
	}
	returnValue = newValues
	return newValues
}

func Calculation(n int) []int {
	if memo[n] != nil {
		fmt.Println("RET Memoized", n)
		return memo[n]
	}
	newStones := MemoizedCalculation(n)
	memo[n] = newStones
	fmt.Println("SET Memoized", n, newStones)
	return newStones
}

func HandleBlink(stones []int, iterations int) []int {
	newStones := make([]int, 0)
	for _, s := range stones {
		newStones = append(newStones, MemoizedPt2Calculation([]int{s}, iterations)...)
	}
	return newStones
}

func ProcessChallenge(input string, iterations int) {
	stones := ProcessInput(input)
	fmt.Printf("Stones Before %d, Count: %d\n", iterations, len(stones))
	stones = HandleBlink(stones, iterations)
	fmt.Printf("Stones On %d, Count: %d\n", len(stones))
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
	//HandleFile(testFileName, 6)
	HandleFile(actualFileName, 75)
}
