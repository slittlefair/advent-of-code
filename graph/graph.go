package graph

import (
	"Advent-of-Code/maths"
	"fmt"
)

// Co is a simple struct for a graph coordinate with points x, y, z
type Co struct {
	X int
	Y int
	Z int
}

// Graph is a map allowing to record string values at coordinates
type Graph map[Co]string

// Grid is a struct containing min and max values as well as a map of coordinates to string values
type Grid struct {
	Graph                  Graph
	MinX, MaxX, MinY, MaxY int
}

func NewGrid() *Grid {
	return &Grid{
		Graph: make(map[Co]string),
	}
}

// OutOfBounds returns a boolean for whether the given coordinate is out of bounds of the grid. That
// is, are any points above any maximum values or below any minimum values
func (g *Grid) OutOfBounds(co Co) bool {
	return co.X < g.MinX || co.X > g.MaxX || co.Y < g.MinY || co.Y > g.MaxY
}

func (g *Grid) PrintGrid() {
	for y := 0; y <= g.MaxY; y++ {
		for x := 0; x <= g.MaxX; x++ {
			fmt.Print(g.Graph[Co{X: x, Y: y}])
		}
		fmt.Println()
	}
}

// AdjacentCos returns all adjacent coordinates for the given coordinate, including diagonals
func AdjacentCos(co Co, includeDiagonals bool) []Co {
	cos := []Co{
		{X: co.X, Y: co.Y - 1},
		{X: co.X - 1, Y: co.Y},
		{X: co.X + 1, Y: co.Y},
		{X: co.X, Y: co.Y + 1},
	}
	if !includeDiagonals {
		return cos
	}
	return append(cos,
		Co{X: co.X - 1, Y: co.Y - 1},
		Co{X: co.X + 1, Y: co.Y - 1},
		Co{X: co.X - 1, Y: co.Y + 1},
		Co{X: co.X + 1, Y: co.Y + 1},
	)
}

// CalculateManhattanDistance calculates the manhattan distance between the origin
func CalculateManhattanDistance(co1, co2 Co) int {
	x := co1.X - co2.X
	y := co1.Y - co2.Y
	return maths.Abs(x) + maths.Abs(y)
}
