package graph

import (
	"Advent-of-Code/maths"
	"fmt"
	"math"
)

// Co is a simple struct for a graph coordinate with points x, y, z
type Co struct {
	X int
	Y int
	Z int
}

// Graph is a map allowing to record string values at coordinates
type Graph[T any] map[Co]T

// Grid is a struct containing min and max values as well as a map of coordinates to string values
type Grid[T any] struct {
	Graph                  Graph[T]
	MinX, MaxX, MinY, MaxY int
}

func NewGrid[T any]() *Grid[T] {
	return &Grid[T]{
		Graph: make(map[Co]T),
	}
}

// OutOfBounds returns a boolean for whether the given coordinate is out of bounds of the grid. That
// is, are any points above any maximum values or below any minimum values
func (g *Grid[T]) OutOfBounds(co Co) bool {
	return co.X < g.MinX || co.X > g.MaxX || co.Y < g.MinY || co.Y > g.MaxY
}

func (g *Grid[T]) PrintGrid() {
	for y := g.MinY; y <= g.MaxY; y++ {
		for x := g.MinX; x <= g.MaxX; x++ {
			if v, ok := g.Graph[Co{X: x, Y: y}]; !ok {
				fmt.Print(".")
			} else {
				fmt.Print(v)
			}
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
func CalculateManhattanDistance(p, q Co) int {
	x := p.X - q.X
	y := p.Y - q.Y
	return maths.Abs(x) + maths.Abs(y)
}

func CalculateEuclideanDistance(p, q Co) float64 {
	x := float64(p.X - q.X)
	y := float64(p.Y - q.Y)
	z := float64(p.Z - q.Z)
	return math.Sqrt(x*x + y*y + z*z)
}
