package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"Advent-of-Code/maths"
	"fmt"
)

// Sensors is a map of sensors to their nearest beacon
type Sensors map[graph.Co]graph.Co

func parseInput(input []string) (Sensors, error) {
	s := Sensors{}
	for _, line := range input {
		sns := graph.Co{}
		bcn := graph.Co{}
		_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sns.X, &sns.Y, &bcn.X, &bcn.Y)
		if err != nil {
			return nil, err
		}
		s[sns] = bcn
	}
	return s, nil
}

func (s Sensors) findTakenSpaces(n int) int {
	// Keep a map of taken spaces which we'll get the length of when it's filled. We keep a map
	// rather than a count as a coordinate could be "taken" by multiple sensors
	taken := map[graph.Co]bool{}
	for sns, bcn := range s {
		// Get the manhattan distace from the sensor to its nearest beacon
		dist := graph.CalculateManhattanDistance(sns, bcn)
		// Run from the x value of the leftmost taken coordinate to the rightmost
		for x := sns.X - dist; x <= sns.X+dist; x++ {
			co := graph.Co{X: x, Y: n}
			// If the manhattan distance from the sensor to the new coordinate is less than the
			// manhattan distance to its beacon then the coordinate is "taken" by the sensor and it
			// is added to our map.
			if graph.CalculateManhattanDistance(sns, co) <= dist {
				taken[co] = true
			}
		}
	}
	// A coordinate cannot be considered taken if it contains a beacon or a sensor so remove any
	// instances of this.
	for sns, bcn := range s {
		delete(taken, sns)
		delete(taken, bcn)
	}
	return len(taken)
}

func (s Sensors) findBeacon(n int) (int, error) {
	// Keep a map of sensors to the manhattan distance to their nearest beacon
	distances := map[graph.Co]int{}
	for sns, bcn := range s {
		distances[sns] = graph.CalculateManhattanDistance(sns, bcn)
	}

	// Walk around the uotside of each sensor's boundaries and see if any of them are within the boundary of
	// another sensor. If not, we've found the beacon
	for sns, bcn := range s {
		dist := graph.CalculateManhattanDistance(sns, bcn)
		for y := maths.Max(sns.Y-dist-1, 0); y <= maths.Min(sns.Y+dist+1, n); y++ {
			offset := dist + 1 - maths.Abs(sns.Y-y)
		out:
			for _, x := range []int{sns.X + offset, sns.X - offset} {
				if x < 0 || x > n {
					continue
				}
				co := graph.Co{X: x, Y: y}
				for otherSns := range s {
					if otherSns != sns {
						if graph.CalculateManhattanDistance(otherSns, co) <= distances[otherSns] {
							continue out
						}
					}
				}
				return (4000000 * x) + y, nil
			}
		}
	}
	return -1, fmt.Errorf("could not find beacon")
}

func findSolutions(input []string, n int) (int, int, error) {
	arrangement, err := parseInput(input)
	if err != nil {
		return -1, -1, err
	}
	part1 := arrangement.findTakenSpaces(n)
	part2, err := arrangement.findBeacon(2 * n)
	if err != nil {
		return -1, -1, err
	}
	return part1, part2, err
}

func main() {
	input := file.Read()
	part1, part2, err := findSolutions(input, 2000000)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
