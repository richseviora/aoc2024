package day2

import (
	"strconv"
	"strings"
)

type Report struct {
	readings []int
	changes  []int
}

func GenerateReportsFromStr(input string) Report {
	inputsSlice := strings.Fields(input) // Split input string into slice of strings
	inputs := make([]int, len(inputsSlice))
	for i, val := range inputsSlice {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			// Handle error
			panic(err)
		}
		inputs[i] = intVal
	}
	return GenerateReports(inputs)
}

func GenerateReports(inputs []int) Report {
	readings := make([]int, len(inputs))
	changes := make([]int, len(inputs)-1)

	for i := 0; i < len(inputs); i++ {
		readings[i] = inputs[i]
	}
	for i := 0; i < len(inputs)-1; i++ {
		changes[i] = inputs[i+1] - inputs[i]
	}
	return Report{readings, changes}
}

func (r Report) IsSafe() bool {
	delta := r.allIncreasing() || r.allDecreasing()
	if !delta {
		return false
	}
	return r.allWithinLimit()
}

func (r Report) allIncreasing() bool {
	for _, change := range r.changes {
		if change < 0 {
			return false
		}
	}
	return true
}

func (r Report) allDecreasing() bool {
	for _, change := range r.changes {
		if change > 0 {
			return false
		}
	}
	return true
}

func (r Report) allWithinLimit() bool {
	for _, change := range r.changes {
		if change == 0 {
			return false
		}
		if change > 3 || change < -3 {
			return false
		}
	}
	return true
}
