package day23

import (
	"Advent-of-Code"
)

type Coordinate struct {
	X, Y, Z int
}

var Zero = Coordinate{X: 0, Y: 0, Z: 0}

func (c Coordinate) Distance(a Coordinate) int {
	return helpers.Abs(c.X-a.X) + helpers.Abs(c.Y-a.Y) + helpers.Abs(c.Z-a.Z)
}
