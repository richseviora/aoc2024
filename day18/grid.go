package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Grid struct {
	cells         map[Coordinate]*Cell
	height, width int
}

var enableDetail = os.Getenv("ENABLE_DETAIL") == "true"

func NewGrid(height, width int) *Grid {
	cells := make(map[Coordinate]*Cell)
	grid := &Grid{
		cells:  cells,
		width:  width,
		height: height,
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			coordinate := Coordinate{x: x, y: y}
			grid.cells[coordinate] = &Cell{
				Coordinate: coordinate,
				Content:    ".",
				Grid:       grid,
			}
		}
	}
	return grid
}

func (g *Grid) PopulateGridFromInput(input string, limit int) {
	inputLines := strings.Split(input, "\n")
	for i, line := range inputLines {
		if i >= limit {
			return
		}
		coordinates := strings.Split(line, ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		coordinate := Coordinate{x: x, y: y}
		if enableDetail {
			fmt.Println("Placing Coordinate", coordinate)
		}
		g.SetCellContent(coordinate, "#")
		if enableDetail {

			g.PrintGrid([]*Cell{})
		}
	}
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

func (c *Cell) IsEnd() bool {
	return c.x == c.Grid.width-1 && c.y == c.Grid.height-1
}

func (c *Cell) CanMoveInDirection(d Direction) bool {
	cell := c.GetCellInDirection(d)
	if cell == nil {
		return false
	}
	return cell.IsEmpty()
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
	return g.cells[Coordinate{x: 0, y: 0}]
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
	//previousAction := Right
	//totalCost := 0
	return len(path) - 1
	//for i := 1; i < len(path); i++ {
	//	thisAction, err := path[i-1].GetCellDirection(path[i])
	//	if err != nil {
	//		fmt.Println("Error", err, path[i-1], path[i])
	//	}
	//	cost := GetMovementCost(previousAction, thisAction)
	//	totalCost += cost
	//	if output {
	//		fmt.Println("Cost", cost, "Total Cost", totalCost, previousAction, thisAction)
	//	}
	//	previousAction = thisAction
	//
	//}
	//return totalCost
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

func (g *Grid) SetCellContent(c Coordinate, content string) {
	g.cells[c].Content = content
}

func (g *Grid) PrintGrid(path []*Cell) {
	fmt.Println("GRID >>>")
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			cell := g.cells[Coordinate{x, y}]
			if slices.Index(path, cell) > -1 {
				fmt.Print("X")
			} else {
				fmt.Print(cell.Content)
			}
		}
		fmt.Println()
	}
	fmt.Println(">>> GRID")
}

func PathToString(path []*Cell) string {
	result := ""
	for _, cell := range path {
		result += fmt.Sprintf("|%d:%d", cell.x, cell.y)
	}
	return result
}

func GetMovementCost(d1, d2 Direction) int {
	return 1
}
