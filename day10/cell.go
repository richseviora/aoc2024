package main

// Coordinates represents the row and column position of a cell
type Coordinates struct {
	Row, Column int
}

// Update the Cell struct to include Coordinates
type Cell struct {
	Value       int
	Grid        *Grid
	Coordinates Coordinates
}

// IterateAdjacentCells iterates over all adjacent cells (including diagonals) of the current cell.
// The provided callback function is called with the row, column, and cell for each adjacent cell.
func (c *Cell) IterateAdjacentCells(callback func(row, col int, cell *Cell)) {
	// Define the relative positions for all 8 adjacent cells
	adjacentOffsets := []struct{ dRow, dCol int }{
		{-1, 0}, // Top
		{0, -1}, // Left
		{0, 1},  // Right
		{1, 0},  // Bottom
	}

	for _, offset := range adjacentOffsets {
		adjRow := c.Coordinates.Row + offset.dRow
		adjCol := c.Coordinates.Column + offset.dCol
		if adjCell := c.Grid.GetCell(adjRow, adjCol); adjCell != nil {
			callback(adjRow, adjCol, adjCell)
		}
	}
}

// GetAdjacentCellsWithValueHigherByOne iterates over adjacent cells and
// returns a slice of cells that have a value higher by 1.
func (c *Cell) GetAdjacentCellsWithValueHigherByOne() []*Cell {
	var result []*Cell
	c.IterateAdjacentCells(func(row, col int, cell *Cell) {
		if cell.Value == c.Value+1 {
			result = append(result, cell)
		}
	})
	return result
}

// GetRoutesStartingWithHigherValue recursively finds all possible routes starting from
// adjacent cells with values higher by one. Each route is represented as a slice of *Cell.
func (c *Cell) GetRoutesStartingWithHigherValue() [][]*Cell {
	var routes [][]*Cell

	var dfs func(current *Cell, path []*Cell)
	dfs = func(current *Cell, path []*Cell) {
		path = append(path, current)

		// Get adjacent cells with value higher by one
		adjacentCells := current.GetAdjacentCellsWithValueHigherByOne()
		if len(adjacentCells) == 0 {
			// If no higher-value adjacent cells, finalize this path
			routes = append(routes, append([]*Cell(nil), path...))
			return
		}

		// Recurse for each adjacent cell
		for _, adjCell := range adjacentCells {
			dfs(adjCell, path)
		}
	}

	// Start recursion from the current cell
	dfs(c, []*Cell{})

	return routes
}

// GetRoutesEndingWithValueNine filters the routes that end in a cell with a value of 9
func (c *Cell) GetRoutesEndingWithValueNine() [][]*Cell {
	allRoutes := c.GetRoutesStartingWithHigherValue()
	var filteredRoutes [][]*Cell

	for _, route := range allRoutes {
		if len(route) > 0 && route[len(route)-1].Value == 9 {
			filteredRoutes = append(filteredRoutes, route)
		}
	}

	return filteredRoutes
}

// GetUniqueLastCells returns a list of unique cells that are the last in the routes ending with value 9
func (c *Cell) GetUniqueLastCells() []*Cell {
	lastCells := make(map[*Cell]struct{})
	routesEndingWithValueNine := c.GetRoutesEndingWithValueNine()

	for _, route := range routesEndingWithValueNine {
		if len(route) > 0 {
			lastCell := route[len(route)-1]
			lastCells[lastCell] = struct{}{}
		}
	}

	uniqueLastCells := make([]*Cell, 0, len(lastCells))
	for cell := range lastCells {
		uniqueLastCells = append(uniqueLastCells, cell)
	}

	return uniqueLastCells
}
