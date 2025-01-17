package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"Advent-of-Code/maths"
	"Advent-of-Code/regex"
	"fmt"
	"slices"
	"strconv"
)

type Bathroom struct {
	robots     []*Robot
	maxX, maxY int
}

type Robot struct {
	position graph.Co
	velocity graph.Co
}

func parseInput(input []string, x, y int) (*Bathroom, error) {
	robots := make([]*Robot, len(input))
	for i, line := range input {
		// For each line, get the four numbers and create a robot with those values
		matches := regex.MatchNumsWithNegatives.FindAllString(line, -1)
		if l := len(matches); l != 4 {
			return nil, fmt.Errorf("malformed input, expected 4 nums, got %d: %v", l, line)
		}
		r := &Robot{}
		for j, m := range matches {
			// We can ignore the error as we know they'll convert due to the regex matching
			n, _ := strconv.Atoi(m)
			switch j {
			case 0:
				r.position.X = n
			case 1:
				r.position.Y = n
			case 2:
				r.velocity.X = n
			case 3:
				r.velocity.Y = n
			}
		}
		robots[i] = r
	}
	return &Bathroom{
		robots: robots,
		maxX:   x,
		maxY:   y,
	}, nil
}

// Move each robot based on its current position and velocity. If the robot would be outside of the
// bathroom (above max or below min values) then act as if it wraps the boundary
func (b *Bathroom) moveRobots() {
	for _, r := range b.robots {
		r.position.X = maths.Modulo(r.position.X+r.velocity.X, b.maxX)
		r.position.Y = maths.Modulo(r.position.Y+r.velocity.Y, b.maxY)
	}
}

// Get all robots in the bathroom, find how many are in each quadrant and then multiply those
// numbers together to return the safety factor. Ignore any robot in the central lines vertically
// and horizontally, which aren't classed as being in any quadrant
func (b Bathroom) findSafetyFactor() int {
	halfWidth := (b.maxX) / 2
	halfHeight := (b.maxY) / 2
	quadrants := [4]int{}
	for _, r := range b.robots {
		if r.position.X == halfWidth || r.position.Y == halfHeight {
			continue
		}
		if r.position.X < halfWidth && r.position.Y < halfHeight {
			quadrants[0]++
			continue
		}
		if r.position.X > halfWidth && r.position.Y < halfHeight {
			quadrants[1]++
			continue
		}
		if r.position.X > halfWidth && r.position.Y > halfHeight {
			quadrants[2]++
			continue
		}
		quadrants[3]++
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

// Try and find robots that are in the shape of a Christmas Tree. Rather than get an exact shape
// assume that being in that shape means at least 10 robots are placed together horizontally. Also
// print the configuration if we find such a time to verify we have a valid tree
func (b Bathroom) foundChristmasTree() bool {
	// For each robot, group them by their position's y coordinate, keeping record of their
	// position's x coordinate
	yFreq := map[int][]int{}
	for _, r := range b.robots {
		yFreq[r.position.Y] = append(yFreq[r.position.Y], r.position.X)
	}
	for _, x := range yFreq {
		// Ignore any lines that don't have at least 10 values
		if len(x) < 10 {
			continue
		}
		// Sort the x values
		slices.Sort(x)
		streak := 0
		// Get any streak of x values that are sequential that is greater than 10.
		for i := 0; i < len(x)-1; i++ {
			// If subsequent x values aren't sequential then start the streak again
			if x[i] != x[i+1]-1 {
				streak = 0
				continue
			}
			streak++
			// If the streak is greater than 10 then we are likely to have found the tree, so return
			// true and print it in all its glory!
			if streak > 10 {
				b.print()
				return true
			}
		}
	}
	return false
}

func findSolutions(input []string, w, h int) (int, int, error) {
	var part1 int
	bathroom, err := parseInput(input, w, h)
	if err != nil {
		return part1, 0, err
	}

	// Keep moving robots and try and find the Christmas tree, returning when we do. We will return
	// at most 100000 iterations as its unlikely the solution will be greater than that, in which
	// case we'll return an error as we assume something's gone wrong.
	i := 1
	for {
		bathroom.moveRobots()
		// Part 1 is the safety factor after 100 rounds of moving
		if i == 100 {
			part1 = bathroom.findSafetyFactor()
		}
		if bathroom.foundChristmasTree() {
			return part1, i, nil
		}
		i++
		if i > 100000 {
			return 0, 0, fmt.Errorf("couldn't find solution for part 2")
		}
	}
}

// Helper function to print the layout of the bathroom. If any space contains at least one robot
// we print the number of robots in that space, otherwise we print a full stop.
func (b Bathroom) print() {
	for y := 0; y < b.maxY; y++ {
		for x := 0; x < b.maxX; x++ {
			n := 0
			for _, r := range b.robots {
				if x == r.position.X && y == r.position.Y {
					n++
				}
			}
			if n > 0 {
				fmt.Print(n)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	input := file.Read()
	part1, part2, err := findSolutions(input, 101, 103)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}
