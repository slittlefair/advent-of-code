package ship

import (
	helpers "Advent-of-Code"
)

type Ship struct {
	Co        helpers.Coordinate
	FacingDir string
}

var points = []string{"N", "E", "S", "W"}

var pointsToIndex = map[string]int{
	"N": 0,
	"E": 1,
	"S": 2,
	"W": 3,
}

func (s *Ship) turnShip(d string, val int) {
	turnTimes := val / 90 % 360
	if d == "L" {
		turnTimes *= -1
	}
	s.FacingDir = points[(pointsToIndex[s.FacingDir]+(turnTimes%len(points))+len(points))%len(points)]
}

func (s *Ship) MoveShip(d string, val int) {
	switch d {
	case "N":
		s.Co.Y += val
	case "E":
		s.Co.X += val
	case "S":
		s.Co.Y -= val
	case "W":
		s.Co.X -= val
	case "L":
		s.turnShip(d, val)
	case "R":
		s.turnShip(d, val)
	case "F":
		s.MoveShip(s.FacingDir, val)
	}
}

func (s *Ship) CalculateDistance() int {
	return helpers.Abs(s.Co.X) + helpers.Abs(s.Co.Y)
}
