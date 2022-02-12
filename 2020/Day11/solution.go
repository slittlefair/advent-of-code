package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"fmt"
	"reflect"
)

type Grid map[graph.Co]string

// Get the 8 adjacent coordinates so we can continue in these directions for part 2
var directions = []graph.Co{
	{X: -1, Y: -1},
	{X: -1, Y: 0},
	{X: -1, Y: +1},
	{X: 0, Y: -1},
	{X: 0, Y: +1},
	{X: +1, Y: -1},
	{X: +1, Y: 0},
	{X: +1, Y: +1},
}

// Decide what state an empty seat should be in at the next stage
func (g Grid) evaluateEmptySeat(co graph.Co, part int) string {
	newVal := "#"
	for _, d := range directions {
		startingCo := co
		for {
			startingCo = graph.Co{
				X: startingCo.X + d.X,
				Y: startingCo.Y + d.Y,
			}
			if val := g[graph.Co{X: startingCo.X, Y: startingCo.Y}]; val != "." {
				if val == "#" {
					newVal = "L"
				}
				break
			} else if part == 1 || (part == 2 && val == "") {
				break
			}
		}
		if newVal == "L" {
			return newVal
		}
	}
	return newVal
}

// Decide what state an occupied seat should be in at the next stage
func (g Grid) evaluateOccupiedSeat(co graph.Co, part int) string {
	// If part 1 empty seat if 4 occupied, if part 2 empty seat if 5 occupied
	adjacentOccupied := 0
	for _, d := range directions {
		startingCo := co
		for {
			startingCo = graph.Co{
				X: startingCo.X + d.X,
				Y: startingCo.Y + d.Y,
			}
			if val := g[graph.Co{X: startingCo.X, Y: startingCo.Y}]; val != "." {
				if val == "#" {
					adjacentOccupied++
				}
				break
			} else if part == 2 && val == "" {
				break
			}
			if part == 1 {
				break
			}
		}
	}
	if adjacentOccupied >= 3+part {
		return "L"
	}
	return "#"
}

// Generate the next grid at the next stage, evaluating each seat in the grid
func (g Grid) generateNextGrid(part int) Grid {
	newGrid := Grid{}
	for co, val := range g {
		if val == "L" {
			newGrid[co] = g.evaluateEmptySeat(co, part)
		} else if val == "#" {
			newGrid[co] = g.evaluateOccupiedSeat(co, part)
		} else {
			newGrid[co] = "."
		}
	}
	return newGrid
}

// Count the number of occupied seats in a grid
func (g Grid) countOccupiedSeats() int {
	numOccupiedSeats := 0
	for _, val := range g {
		if val == "#" {
			numOccupiedSeats++
		}
	}
	return numOccupiedSeats
}

// keep generating a new grid until we reach equilibrium, in which case return the number of
// occupied seats
func (g Grid) findSolution(part int) int {
	for {
		newGrid := g.generateNextGrid(part)
		if reflect.DeepEqual(g, newGrid) {
			return g.countOccupiedSeats()
		}
		g = newGrid
	}
}

// For debugging: uncomment this and add to findSolution to see a visial representation of the grid
// at each iteration
// func (g Grid) printExampleGrid() {
// 	printGrid := [10][10]string{}
// 	for co, val := range g {
// 		printGrid[co.Y][co.X] = val
// 	}
// 	for _, row := range printGrid {
// 		fmt.Println(row)
// 	}
// 	fmt.Println()
// }

func (g Grid) parseInput(plan []string) {
	for r, row := range plan {
		for c, col := range row {
			g[graph.Co{X: c, Y: r}] = string(col)
		}
	}
}

func main() {
	plan := file.Read()
	g := Grid{}
	g.parseInput(plan)
	fmt.Println("Part 1:", g.findSolution(1))
	fmt.Println("Part 2:", g.findSolution(2))
}
