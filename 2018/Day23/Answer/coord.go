package day23

import (
	"Advent-of-Code/maths"
)

type Coordinate struct {
	X, Y, Z int
}

var Zero = Coordinate{X: 0, Y: 0, Z: 0}

func (c Coordinate) Distance(a Coordinate) int {
	return maths.Abs(c.X-a.X) + maths.Abs(c.Y-a.Y) + maths.Abs(c.Z-a.Z)
}
