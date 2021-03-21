package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"strconv"
)

type Ship struct {
	co        helpers.Coordinate
	facingDir string
}

type Waypoint helpers.Coordinate

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
	s.facingDir = points[(pointsToIndex[s.facingDir]+turnTimes+len(points))%len(points)]
}

func (s *Ship) moveShip(d string, val int) {
	switch d {
	case "N":
		s.co.Y += val
	case "E":
		s.co.X += val
	case "S":
		s.co.Y -= val
	case "W":
		s.co.X -= val
	case "L":
		s.turnShip(d, val)
	case "R":
		s.turnShip(d, val)
	case "F":
		s.moveShip(s.facingDir, val)
	}
}

func (s *Ship) calculateDistance() int {
	return helpers.Abs(s.co.X) + helpers.Abs(s.co.Y)
}

func parseDirection(entry string) (string, int, error) {
	dir := string(entry[0])
	val, err := strconv.Atoi(entry[1:])
	return dir, val, err
}

func part1(entries []string) (int, error) {
	ship := Ship{
		facingDir: "E",
	}
	for _, entry := range entries {
		dir, val, err := parseDirection(entry)
		if err != nil {
			return 0, err
		}
		ship.moveShip(dir, val)
	}
	return ship.calculateDistance(), nil
}

func (w *Waypoint) turnWaypointLeft(val int) {
	// Rotated around the origin 90 degrees anticlockwise point M (h, k) takes the image M' (-k, h)
	for i := 0; i < val; i += 90 {
		newWaypoint := helpers.Coordinate{
			X: -w.Y,
			Y: w.X,
		}
		w.X = newWaypoint.X
		w.Y = newWaypoint.Y
	}
}

func (w *Waypoint) turnWaypointRight(val int) {
	// Rotated around the origin 90 degrees clockwise point M (h, k) takes the image M' (k, -h)
	for i := 0; i < val; i += 90 {
		newWaypoint := helpers.Coordinate{
			X: w.Y,
			Y: -w.X,
		}
		w.X = newWaypoint.X
		w.Y = newWaypoint.Y
	}
}

func (w *Waypoint) moveWaypoint(s *Ship, d string, val int) {
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
		w.turnWaypointLeft(val)
	case "R":
		w.turnWaypointRight(val)
	case "F":
		for i := 0; i < val; i++ {
			if w.X > 0 {
				s.moveShip("E", w.X)
			} else {
				s.moveShip("W", -w.X)
			}
			if w.Y > 0 {
				s.moveShip("N", w.Y)
			} else {
				s.moveShip("S", -w.Y)
			}
		}
	}
}

func part2(entries []string) (int, error) {
	ship := Ship{
		facingDir: "E",
	}
	waypoint := Waypoint{
		X: 10,
		Y: 1,
	}
	for _, entry := range entries {
		dir, val, err := parseDirection(entry)
		if err != nil {
			return 0, err
		}
		waypoint.moveWaypoint(&ship, dir, val)
	}
	return ship.calculateDistance(), nil
}

func main() {
	entries := helpers.ReadFile()
	part1Sol, err := part1(entries)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1Sol)

	part2Sol, err := part2(entries)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 2:", part2Sol)
}
