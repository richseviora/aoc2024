package main

type Direction int64

const (
	Up Direction = iota
	Left
	Right
	Down
)

type Contents int64

const (
	Empty Contents = iota
	Barrier
)

type Cell struct {
	Row    int
	Column int
	Contents
}
