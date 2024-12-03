package day2

type ReportWithDampener struct {
	readings []int
	changes  []int
}

func (r ReportWithDampener) Readings() []int {
	return r.readings
}

func (r ReportWithDampener) Changes() []int {
	return r.changes
}

func (r ReportWithDampener) IsSafe() bool {
	increaseErrors := r.allIncreasing()
	decreaseErrors := r.allDecreasing()
	limitErrors := r.allWithinLimit()
	decreaseErrorsToEval := append(decreaseErrors, limitErrors...)
	increaseErrorsToEval := append(increaseErrors, limitErrors...)

	unique := func(input []int) []int {
		uniqueValues := make(map[int]bool)
		result := []int{}
		for _, value := range input {
			if !uniqueValues[value] {
				uniqueValues[value] = true
				result = append(result, value)
			}
		}
		return result
	}

	increaseErrorsToEval = unique(increaseErrorsToEval)

	decreaseErrorsToEval = unique(decreaseErrorsToEval)

	return len(increaseErrorsToEval) <= 1 || len(decreaseErrorsToEval) <= 1
}

func (r ReportWithDampener) allIncreasing() []int {
	errors := make([]int, 0)
	for i, change := range r.changes {
		if change < 0 {
			errors = append(errors, i)
		}
	}
	return errors
}

func (r ReportWithDampener) allDecreasing() []int {
	errors := make([]int, 0)
	for i, change := range r.changes {
		if change > 0 {
			errors = append(errors, i)
		}
	}
	return errors
}

func (r ReportWithDampener) allWithinLimit() []int {
	errors := make([]int, 0)
	for i, change := range r.changes {
		if change == 0 {
			errors = append(errors, i)
		}
		if change > 3 || change < -3 {
			errors = append(errors, i)
		}
	}
	return errors
}
