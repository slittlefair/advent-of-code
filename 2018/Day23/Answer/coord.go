package day23

import (
	utils "Advent-of-Code/utils"
)

type Coordinate struct {
	X, Y, Z int
}

var Zero = Coordinate{X: 0, Y: 0, Z: 0}

func (c Coordinate) Distance(a Coordinate) int {
	return utils.Abs(c.X-a.X) + utils.Abs(c.Y-a.Y) + utils.Abs(c.Z-a.Z)
}
