package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"fmt"
	"strconv"
)

type TopMap struct {
	graph.Grid[int]
	// Trails is a map of trailHead coordinates (value 0), to a map of trailEnd coordinates (value
	// 9) to a map of string representation of the path taken to get there. This way we are able to
	// record distinct trails, and know the amount of trails that have those coordinates.
	trails map[graph.Co]map[graph.Co]map[string]bool
}

func parseInput(input []string) (*TopMap, error) {
	tm := &TopMap{
		Grid: graph.Grid[int]{
			MaxY:  len(input) - 1,
			MaxX:  len(input[0]) - 1,
			Graph: make(map[graph.Co]int),
		},
		trails: make(map[graph.Co]map[graph.Co]map[string]bool),
	}
	for y, line := range input {
		for x, char := range line {
			n, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}
			tm.Graph[graph.Co{X: x, Y: y}] = n
		}
	}
	return tm, nil
}

// move to an adjacent coordinate in the TopMap and evaluate whetehr we're still on track for a
// valid trail or not. If we are, but not yet finished, we recursively call this function to move
// again.
func (tm *TopMap) move(trailHead graph.Co, currentVal int, newCo graph.Co, currentPath []graph.Co) {
	// If the new coordinate is outside of the grid, it's not a valid trail so return
	if tm.OutOfBounds(newCo) {
		return
	}

	newVal := tm.Graph[newCo]

	// If the next value in the trail isn't one greater than the current value, it's not a valid
	// trail so return
	if newVal-currentVal != 1 {
		return
	}

	// If the next value in the trail is 9, we've come to the end of a valid trail. In this case,
	// populate the TopMap.trails cache with the trailHead, trailEnd and path.
	if newVal == 9 {
		if _, ok := tm.trails[trailHead]; !ok {
			tm.trails[trailHead] = map[graph.Co]map[string]bool{}
		}
		if _, ok := tm.trails[trailHead][newCo]; !ok {
			tm.trails[trailHead][newCo] = map[string]bool{}
		}
		pathString := makePathString(currentPath)
		tm.trails[trailHead][newCo][pathString] = true
		return
	}

	// If we're here we're still travelling up a potentially valid path, so get the four possible
	// next coordinates and call move for each of them
	adjCos := graph.AdjacentCos(newCo, false)
	for _, co := range adjCos {
		tm.move(trailHead, newVal, co, append(currentPath, co))
	}
}

// run through each coordinate in the grid and start moving from every one that has 0 value
func (tm *TopMap) findTrails() {
	for co, v := range tm.Graph {
		if v == 0 {
			adjCos := graph.AdjacentCos(co, false)
			for _, aCo := range adjCos {
				tm.move(co, 0, aCo, []graph.Co{aCo})
			}
		}
	}
}

// get the scores for all the trails that were valid
func (tm TopMap) getScores() (int, int) {
	part1 := 0
	part2 := 0
	for _, trails := range tm.trails {
		// For part 1 we only care about whether a trail exists from a value of 0 to a value of 9.
		// If it does then that combination of the two has a score of 1.
		part1 += len(trails)
		for _, uniqueTrails := range trails {
			// For part 2 we care about how many valid trails there are from a value of 0 to a value
			// of 9. For this we increase the part2 score by how many such trails there are.
			part2 += len(uniqueTrails)
		}
	}
	return part1, part2
}

// Turn a slice of coordinates to a unique string representation to be used as a key in a map. This
// way we can keep a map of unique trail paths.
func makePathString(path []graph.Co) string {
	var pathString string
	for _, co := range path {
		pathString = fmt.Sprintf("%sX:%dY:%d", pathString, co.X, co.Y)
	}
	return pathString
}

// find the solutions for parts 1 and 2 for a given input. If, as we're parsing the input, we happen
// upon an error we'll return this instead.
func findSolutions(input []string) (int, int, error) {
	topMap, err := parseInput(input)
	if err != nil {
		return 0, 0, err
	}
	topMap.findTrails()
	part1, part2 := topMap.getScores()
	return part1, part2, nil
}

func main() {
	input := file.Read()
	part1, part2, err := findSolutions(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}
