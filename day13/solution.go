package main

type Solution struct {
	a, b int
}

func (s Solution) Cost() int {
	return s.a*3 + s.b*1
}
