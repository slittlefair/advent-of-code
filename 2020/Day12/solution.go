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

func main() {
	ship := Ship{
		facingDir: "E",
	}
	entries := helpers.ReadFile()

	for _, entry := range entries {
		dir := string(entry[0])
		val, err := strconv.Atoi(entry[1:])
		if err != nil {
			fmt.Println(err, entry)
			return
		}
		ship.moveShip(dir, val)
	}

	fmt.Println("Part 1:", ship.calculateDistance())
}
