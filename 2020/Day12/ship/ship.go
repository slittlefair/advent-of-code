package ship

import (
	"Advent-of-Code/graph"
	"Advent-of-Code/maths"
)

type Ship struct {
	Co        graph.Co
	FacingDir string
}

var points = []string{"N", "E", "S", "W"}

var pointsToIndex = map[string]int{
	"N": 0,
	"E": 1,
	"S": 2,
	"W": 3,
}

func (s *Ship) TurnShip(d string, val int) {
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
		s.TurnShip(d, val)
	case "R":
		s.TurnShip(d, val)
	case "F":
		s.MoveShip(s.FacingDir, val)
	}
}

func (s *Ship) CalculateDistance() int {
	return maths.Abs(s.Co.X) + maths.Abs(s.Co.Y)
}
