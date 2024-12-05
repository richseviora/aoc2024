package main

type Direction int64

const (
	Up Direction = iota
	Down
	Left
	Right
	DownRight
	DownLeft
	UpLeft
	UpRight
)

type Cell struct {
	Value  string
	Table  *Table
	Row    int
	Column int
}

func (c Cell) GetCellInDirection(d Direction) *Cell {
	switch d {
	case Up:
		return c.Table.GetCellAt(c.Row-1, c.Column)
	case UpLeft:
		return c.Table.GetCellAt(c.Row-1, c.Column-1)
	case UpRight:
		return c.Table.GetCellAt(c.Row-1, c.Column+1)
	case Down:
		return c.Table.GetCellAt(c.Row+1, c.Column)
	case DownRight:
		return c.Table.GetCellAt(c.Row+1, c.Column+1)
	case DownLeft:
		return c.Table.GetCellAt(c.Row+1, c.Column-1)
	case Left:
		return c.Table.GetCellAt(c.Row, c.Column-1)
	case Right:
		return c.Table.GetCellAt(c.Row, c.Column+1)
	}
	panic("Invalid Direction")
}

func (c Cell) GetCellsInDirection(d Direction, n int) []*Cell {
	cells := []*Cell{&c}
	currentCell := &c
	for i := 1; i < n; i++ {
		currentCell = currentCell.GetCellInDirection(d)
		if currentCell == nil {
			break
		}
		cells = append(cells, currentCell)
	}
	return cells
}
