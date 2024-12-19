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
var printGuard = []string{"^", ">", "v", "<"}

func parseInput(input []string) *Floor {
	g := graph.NewGrid()
	g.MaxY = len(input)
	g.MaxX = len(input[0])
	guard := &Guard{
		dir: 0,
	}
	for y, line := range input {
		for x, r := range line {
			str := string(r)
			if str == "^" {
				guard.co = graph.Co{X: x, Y: y}
				str = "."
			}
			g.Graph[graph.Co{X: x, Y: y}] = str
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

func (f *Floor) printFloor() {
	numObstacles := 0
	for y := f.grid.MinY; y <= f.grid.MaxY; y++ {
		for x := f.grid.MinX; x <= f.grid.MaxX; x++ {
			co := graph.Co{X: x, Y: y}
			if co == f.guard.co {
				fmt.Print(printGuard[f.guard.dir])
			} else if _, ok := f.visitedSteps[co]; ok {
				fmt.Printf("X")
			} else {
				fmt.Print(f.grid.Graph[co])
				if f.grid.Graph[co] == "#" {
					numObstacles++
				}
			}
		}
		fmt.Println()
	}
	fmt.Printf("maxX: %d, maxY: %d, numObstacles: %d\n", f.grid.MaxX, f.grid.MaxY, numObstacles)
}

func findSolutions(input []string) (int, int) {
	originalFloor := parseInput(input)
	// TODO nice way of keeping the original
	floor := &Floor{
		guard: &Guard{
			co:  originalFloor.guard.co,
			dir: originalFloor.guard.dir,
		},
		grid: &graph.Grid{
			MaxX:  originalFloor.grid.MaxX,
			MaxY:  originalFloor.grid.MaxY,
			Graph: originalFloor.grid.Graph,
		},
		visitedSteps: originalFloor.visitedSteps,
	}
	floor.printFloor()
	fmt.Println("len1", len(originalFloor.visitedSteps))

	// Part 1
	inBounds := true
	for inBounds {
		inBounds, _ = floor.step()
		if _, ok := floor.grid.Graph[floor.guard.co]; !ok {
			break
		}
	}
	part1 := len(floor.visitedSteps)
	floor.printFloor()
	fmt.Println("len1", len(originalFloor.visitedSteps))

	originalFloor.printFloor()

	for y := floor.grid.MinY; y <= floor.grid.MaxY; y++ {
		for x := floor.grid.MaxX; x <= floor.grid.MaxX; x++ {

		}
	}

	return part1, 0
}

func (f *Floor) step() (bool, bool) {
	guardDir := dirs[f.guard.dir]
	newGuardCo := graph.Co{X: f.guard.co.X + guardDir.X, Y: f.guard.co.Y + guardDir.Y}
	if f.grid.Graph[newGuardCo] == "#" {
		f.guard.dir = (f.guard.dir + 1) % len(dirs)
		return true, false
	} else {
		f.guard.co = newGuardCo
		if f.visitedBefore() {
			fmt.Printf("visitedbefore %v\n", f.guard)
			f.printFloor()
			return true, true
		}
	}
	if _, ok := f.grid.Graph[newGuardCo]; !ok {
		return false, false
	}
	if _, ok := f.visitedSteps[newGuardCo][f.guard.dir]; !ok {
		f.visitedSteps[newGuardCo] = map[int]bool{}
	}
	f.visitedSteps[newGuardCo][f.guard.dir] = true
	return true, false
}

// Must be called before visitedSteps populated with current guard
func (f *Floor) visitedBefore() bool {
	return f.visitedSteps[f.guard.co][f.guard.dir]
}

func main() {
	input := file.Read()
	part1, part2 := findSolutions(input)
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %d\n", part2)
}
