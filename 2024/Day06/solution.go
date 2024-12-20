package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"fmt"
)

type Guard struct {
	co  graph.Co
	dir int
}

type Floor struct {
	guard        *Guard
	grid         *graph.Grid
	visitedSteps map[graph.Co]map[int]bool
}

var dirs = []graph.Co{{Y: -1}, {X: 1}, {Y: 1}, {X: -1}}

// utility function for printing the floor, showing the positions of any obstacles as well as the
// position and direction of the guard
// func (f *Floor) print() {
// 	fmt.Println()
// 	for y := f.grid.MinY; y <= f.grid.MaxY; y++ {
// 		for x := f.grid.MinX; x <= f.grid.MaxX; x++ {
// 			co := graph.Co{X: x, Y: y}
// 			if co == f.guard.co {
// 				fmt.Print([]string{"^", ">", "v", "<"}[f.guard.dir])
// 			} else if _, ok := f.visitedSteps[co]; ok {
// 				fmt.Print("X")
// 			} else if v, ok := f.grid.Graph[co]; ok {
// 				fmt.Print(v)
// 			} else {
// 				fmt.Print(".")
// 			}
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println()
// }

func parseInput(input []string) *Floor {
	g := graph.NewGrid()
	g.MaxY = len(input) - 1
	g.MaxX = len(input[0]) - 1
	guard := &Guard{
		dir: 0,
	}
	for y, line := range input {
		for x, r := range line {
			str := string(r)
			if str == "^" {
				guard.co = graph.Co{X: x, Y: y}
			} else if str == "#" {
				g.Graph[graph.Co{X: x, Y: y}] = str
			}
		}
	}
	return &Floor{
		guard: guard,
		grid:  g,
		visitedSteps: map[graph.Co]map[int]bool{
			guard.co: {guard.dir: true},
		},
	}
}

// Attempt to move the guard. If facing an obstacle they instead turn
func (f *Floor) step() (bool, bool) {
	guardDir := dirs[f.guard.dir]
	newGuardCo := graph.Co{X: f.guard.co.X + guardDir.X, Y: f.guard.co.Y + guardDir.Y}
	// If the guard's new position would see them in an obstacle, turn them 90 degress clockwise.
	// Otherwise move them forward
	if f.grid.Graph[newGuardCo] == "#" {
		f.guard.dir = (f.guard.dir + 1) % len(dirs)
		return true, false
	} else {
		f.guard.co = newGuardCo
	}
	// If the guard havs visited that space before (same space and same direction), they're in a
	// loop, so exit
	if f.visitedSteps[f.guard.co][f.guard.dir] {
		return true, true
	}
	// If we're out of bounds then the guard has moved off the floor, so return
	if f.grid.OutOfBounds(newGuardCo) {
		return false, false
	}
	// Add the guard's new space to the list of visited steps
	if _, ok := f.visitedSteps[newGuardCo]; !ok {
		f.visitedSteps[newGuardCo] = map[int]bool{}
	}
	f.visitedSteps[newGuardCo][f.guard.dir] = true
	return true, false
}

// Move the guard on the floor until they are either stuck in a loop or move out of bounds. Return
// whether they are stuck in a loop so we can keep a count of these instances.
func (f *Floor) runPatrol() bool {
	inBounds := true
	visitedBefore := false
	for inBounds && !visitedBefore {
		inBounds, visitedBefore = f.step()
	}
	return visitedBefore
}

// Reset a floor to the given values. This prevents us having to run parseInput each time which is
// slightly less efficient
func (f *Floor) resetFloor(obstacle graph.Co, originalGuard Guard) {
	f.guard = &Guard{
		co: graph.Co{
			X: originalGuard.co.X,
			Y: originalGuard.co.Y,
		},
		dir: 0,
	}
	f.visitedSteps = map[graph.Co]map[int]bool{
		originalGuard.co: {
			originalGuard.dir: true,
		},
	}
	delete(f.grid.Graph, obstacle)
}

func findSolutions(input []string) (int, int) {
	floor := parseInput(input)

	// Record the original graph and guard values so we can reset the floor after each part2 iteration
	originalGraph := graph.Graph{}
	for k, v := range floor.grid.Graph {
		originalGraph[k] = v
	}
	originalGuard := Guard{
		co:  floor.guard.co,
		dir: floor.guard.dir,
	}

	// Part 1
	floor.runPatrol()
	part1 := len(floor.visitedSteps)
	visitedSteps := map[graph.Co]map[int]bool{}
	for k, v := range floor.visitedSteps {
		visitedSteps[k] = v
	}
	floor.resetFloor(graph.Co{}, originalGuard)

	part2 := 0

	// For any placed objects to have an impact they must be placed in a space the guard would have
	// visited in part1, otherwise placing the object would have no affect on the guard's movement.
	// This lets us narrow down the number of coordinates we have to try for part2 from every space
	// on the floor to just those visited previously
	for k := range visitedSteps {
		// Don't put an obstacle in the guard's original position
		if k == originalGuard.co {
			continue
		}
		// Place the obstacle before running the patrol
		floor.grid.Graph[k] = "#"
		if floor.runPatrol() {
			part2++
		}
		// Reset the floor after running the patrol
		floor.resetFloor(k, originalGuard)
	}

	return part1, part2
}

func main() {
	input := file.Read()
	part1, part2 := findSolutions(input)
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %d\n", part2)
}
