package day2

type Reports struct {
	reports []int
	changes []int
}

func GenerateReports(readings []int) Reports {
	reports := make([]int, len(readings))
	changes := make([]int, len(readings)-1)

	for i := 0; i < len(readings); i++ {
		reports[i] = readings[i]
	}
	for i := 0; i < len(readings)-1; i++ {
		changes[i] = readings[i+1] - readings[i]
	}
	return Reports{reports, changes}
}

func (r Reports) IsSafe() bool {
	delta := r.allIncreasing() || r.allDecreasing()
	if !delta {
		return false
	}
	return r.allWithinLimit()
}

func (r Reports) allIncreasing() bool {
	for _, change := range r.changes {
		if change < 0 {
			return false
		}
	}
	return true
}

func (r Reports) allDecreasing() bool {
	for _, change := range r.changes {
		if change > 0 {
			return false
		}
	}
	return true
}

func (r Reports) allWithinLimit() bool {
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
