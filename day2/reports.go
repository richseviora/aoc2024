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

func (r Reports) allValid() {

}
