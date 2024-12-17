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

// Grid is a struct containing min and max values as well as a map of coordinates to string values
type Grid struct {
	Graph                  map[Co]string
	MinX, MaxX, MinY, MaxY int
}

func NewGrid() *Grid {
	return &Grid{
		Graph: make(map[Co]string),
	}
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
