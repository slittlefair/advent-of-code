package waypoint

import (
	"Advent-of-Code/2020/Day12/ship"
	"Advent-of-Code/graph"
)

type Waypoint graph.Co

func (w *Waypoint) TurnWaypointLeft(val int) {
	// Rotated around the origin 90 degrees anticlockwise point M (h, k) takes the image M' (-k, h)
	for i := 0; i < val; i += 90 {
		newWaypoint := graph.Co{
			X: -w.Y,
			Y: w.X,
		}
		w.X = newWaypoint.X
		w.Y = newWaypoint.Y
	}
}

func (w *Waypoint) TurnWaypointRight(val int) {
	// Rotated around the origin 90 degrees clockwise point M (h, k) takes the image M' (k, -h)
	for i := 0; i < val; i += 90 {
		newWaypoint := graph.Co{
			X: w.Y,
			Y: -w.X,
		}
		w.X = newWaypoint.X
		w.Y = newWaypoint.Y
	}
}

func (w *Waypoint) MoveWaypoint(s *ship.Ship, d string, val int) {
	switch d {
	case "N":
		w.Y += val
	case "E":
		w.X += val
	case "S":
		w.Y -= val
	case "W":
		w.X -= val
	case "L":
		w.TurnWaypointLeft(val)
	case "R":
		w.TurnWaypointRight(val)
	case "F":
		for i := 0; i < val; i++ {
			if w.X > 0 {
				s.MoveShip("E", w.X)
			} else {
				s.MoveShip("W", -w.X)
			}
			if w.Y > 0 {
				s.MoveShip("N", w.Y)
			} else {
				s.MoveShip("S", -w.Y)
			}
		}
	}
}
