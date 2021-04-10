package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"reflect"
)

type Coord4D struct {
	X int
	Y int
	Z int
	W int
}
type Grid map[Coord4D]string

// Compare the given co to its neighbour and increment the number of neighboursActive if necessary
func (g Grid) evaluateAdjacentCo(co Coord4D, adjacentCo Coord4D, neighboursActive int) int {
	if val := g[adjacentCo]; val == "#" && !reflect.DeepEqual(co, adjacentCo) {
		return neighboursActive + 1
	}
	return neighboursActive
}

// Evaluate the given coordinate and return the new value
func (g Grid) evaluateCo(is4D bool, co Coord4D) string {
	neighboursActive := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				adjacentCo := Coord4D{
					X: co.X + x,
					Y: co.Y + y,
					Z: co.Z + z,
				}
				if is4D {
					for w := -1; w <= 1; w++ {
						adjacentCo.W = co.W + w
						neighboursActive = g.evaluateAdjacentCo(co, adjacentCo, neighboursActive)
					}
				} else {
					neighboursActive = g.evaluateAdjacentCo(co, adjacentCo, neighboursActive)
				}
			}
		}
	}
	if neighboursActive == 3 || (g[co] == "#" && neighboursActive == 2) {
		return "#"
	}
	return "."
}

func (g Grid) parseInput(pocketDimension []string, iterations int) {
	// Pad out the grid so that all future cubes are already considered. We can at most add one
	// cube on either side of the dimension each iteration, so the edges of our initial grid plus
	// the number of iterations each side is the maximum our grid can get to.
	width := len(pocketDimension[0]) - 1
	height := len(pocketDimension) - 1
	for x := -iterations; x <= width+iterations; x++ {
		for y := -iterations; y <= height+iterations; y++ {
			for z := -iterations; z <= iterations; z++ {
				for w := -iterations; w <= iterations; w++ {
					g[Coord4D{X: x, Y: y, Z: z, W: w}] = string(".")
				}
			}
		}
	}
	for r, row := range pocketDimension {
		for c, col := range row {
			g[Coord4D{X: c, Y: r, Z: 0, W: 0}] = string(col)
		}
	}
}

// Generate the next grid at the next stage, evaluating each seat in the grid
func (g Grid) generateNextGrid(is4D bool) Grid {
	newGrid := Grid{}
	for co := range g {
		if is4D || co.W == 0 {
			newGrid[co] = g.evaluateCo(is4D, co)
		}
	}
	return newGrid
}

func (g Grid) countActiveCubes() int {
	numActiveCubes := 0
	for _, val := range g {
		if val == "#" {
			numActiveCubes++
		}
	}
	return numActiveCubes
}

// A runner for the solution of each part
func (g Grid) findSolution(is4D bool, iterations int) int {
	for i := 0; i < iterations; i++ {
		newGrid := g.generateNextGrid(is4D)
		g = newGrid
	}
	return g.countActiveCubes()
}

func main() {
	pocketDimension := helpers.ReadFile()
	iterations := 6
	g := Grid{}
	g.parseInput(pocketDimension, iterations)
	fmt.Println("Part 1:", g.findSolution(false, iterations))
	fmt.Println("Part 2:", g.findSolution(true, iterations))
}
