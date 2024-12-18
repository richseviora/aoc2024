package main

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

type Grid struct {
	cells         map[Coordinate]*Cell
	height, width int
}

func NewGridFromInput(input string) Grid {
	lines := strings.Split(input, "\n")
	cells := make(map[Coordinate]*Cell)
	grid := &Grid{
		cells: cells,
	}

	for y, line := range lines {
		for x, content := range line {
			if x >= grid.width {
				grid.width = x
			}
			if y >= grid.height {
				grid.height = y
			}
			coordinate := Coordinate{x: x, y: y}
			grid.cells[coordinate] = &Cell{
				Coordinate: coordinate,
				Content:    string(content),
				Grid:       grid,
			}
		}
	}
	return Grid{cells: cells}
}

type Coordinate struct {
	x, y int
}

type Cell struct {
	Coordinate
	Content string
	Grid    *Grid
}

type Direction int64

const (
	Up Direction = iota
	Left
	Right
	Down
)

func (c *Cell) IsEmpty() bool {
	return c.Content == "."
}

func (c *Cell) IsWall() bool {
	return c.Content == "#"
}

func (c *Cell) IsStart() bool {
	return c.Content == "S"
}

func (c *Cell) IsEnd() bool {
	return c.Content == "E"
}

func (c *Cell) CanMoveInDirection(d Direction) bool {
	cell := c.GetCellInDirection(d)
	if cell == nil {
		return false
	}
	return cell.IsEmpty() || cell.IsEnd()
}

func (c *Cell) GetPossibleMoves() []*Cell {
	possibleMoves := []*Cell{}
	for _, d := range []Direction{Up, Left, Right, Down} {
		if c.CanMoveInDirection(d) {
			possibleMoves = append(possibleMoves, c.GetCellInDirection(d))
		}

	}
	return possibleMoves
}

func (c *Cell) GetCellInDirection(d Direction) *Cell {
	switch d {
	case Up:
		return c.Grid.cells[Coordinate{
			x: c.x,
			y: c.y - 1,
		}]
	case Left:
		return c.Grid.cells[Coordinate{
			x: c.x - 1,
			y: c.y,
		}]
	case Right:
		return c.Grid.cells[Coordinate{
			x: c.x + 1,
			y: c.y,
		}]
	case Down:
		return c.Grid.cells[Coordinate{
			x: c.x,
			y: c.y + 1,
		}]
	}
	panic("invalid direction")
}

func (c *Cell) GetCellDirection(od *Cell) (Direction, error) {
	if c.x == od.x && c.y == od.y-1 {
		return Up, nil
	} else if c.x == od.x-1 && c.y == od.y {
		return Left, nil
	} else if c.x == od.x+1 && c.y == od.y {
		return Right, nil
	} else if c.x == od.x && c.y == od.y+1 {
		return Down, nil
	}
	return 0, errors.New(
		"not adjacent")
}

func (g *Grid) FindStartCell() *Cell {
	for _, cell := range g.cells {
		if cell.IsStart() {
			return cell
		}
	}
	panic("no start cell found")
}

func (g *Grid) GetPathsToEnd() [][]*Cell {
	var paths [][]*Cell
	startCell := g.FindStartCell()
	var dfs func(cell *Cell, cells []*Cell)
	dfs = func(cell *Cell, cells []*Cell) {
		if cell.IsEnd() {
			paths = append(paths, cells)
			return
		}
		grid := g
		if grid == nil {
			panic("grid is nil")
		}
		for _, possibleMove := range cell.GetPossibleMoves() {

			alreadyVisited := slices.Index(cells, possibleMove) > -1

			if alreadyVisited {
				continue
			}
			newPath := make([]*Cell, len(cells))
			copy(newPath, cells)
			newPath = append(newPath, possibleMove)
			dfs(possibleMove, newPath)
		}
	}
	dfs(startCell, []*Cell{startCell})
	return paths
}

func CalculatePathCost(path []*Cell, output bool) int {
	previousAction := Right
	totalCost := 0
	for i := 1; i < len(path); i++ {
		thisAction, err := path[i-1].GetCellDirection(path[i])
		if err != nil {
			fmt.Println("Error", err, path[i-1], path[i])
		}
		cost := GetMovementCost(previousAction, thisAction)
		totalCost += cost
		if output {
			fmt.Println("Cost", cost, "Total Cost", totalCost, previousAction, thisAction)
		}
		previousAction = thisAction

	}
	return totalCost
}

func (g *Grid) GetCheapestPath(output bool) int {
	paths := g.GetPathsToEnd()
	if len(paths) == 0 {
		panic("no paths found")
	}
	slices.SortFunc(paths, func(a, b []*Cell) int {
		return CalculatePathCost(a, false) - CalculatePathCost(b, false)
	})
	if output {

		for i, path := range paths {
			fmt.Printf("Path %d COST %d >>> %s\n", i, CalculatePathCost(path, false), PathToString(path))

		}
	}
	g.PrintGrid(paths[0])
	return CalculatePathCost(paths[0], output)
}

func (g *Grid) PrintGrid(path []*Cell) {
	for x := 0; x <= g.width; x++ {
		for y := 0; y <= g.height; y++ {
			cell := g.cells[Coordinate{x, y}]
			if slices.Index(path, cell) > -1 {
				fmt.Print("X")
			} else {
				fmt.Print(cell.Content)
			}
		}
		fmt.Println()
	}
}

func PathToString(path []*Cell) string {
	result := ""
	for _, cell := range path {
		result += fmt.Sprintf("|%d:%d", cell.x, cell.y)
	}
	return result
}

func GetMovementCost(d1, d2 Direction) int {
	if d1 == d2 {
		return 1
	}
	return 1001
}
