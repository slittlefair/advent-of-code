package main

import (
	helpers "Advent-of-Code"
	"fmt"
)

type Grid map[helpers.Coordinate]string

var directions = []helpers.Coordinate{
	{X: -1, Y: -1},
	{X: -1, Y: 0},
	{X: -1, Y: +1},
	{X: 0, Y: -1},
	{X: 0, Y: +1},
	{X: +1, Y: -1},
	{X: +1, Y: 0},
	{X: +1, Y: +1},
}

func (g *Grid) isSameGrid(newGrid Grid) bool {
	for co, val := range *g {
		if newGrid[co] != val {
			return false
		}
	}
	return true
}

func (g *Grid) evaluateEmptySeat(co helpers.Coordinate, part int) string {
	newVal := "#"
	for _, d := range directions {
		startingCo := co
		for {
			startingCo = helpers.Coordinate{
				X: startingCo.X + d.X,
				Y: startingCo.Y + d.Y,
			}
			if val := (*g)[helpers.Coordinate{X: startingCo.X, Y: startingCo.Y}]; val != "." {
				if val == "#" {
					newVal = "L"
				}
				break
			} else if part == 2 && val == "" {
				break
			}
			if part == 1 {
				break
			}
		}
		if newVal == "L" {
			break
		}
	}
	return newVal
}

func (g *Grid) evaluateOccupiedSeat(co helpers.Coordinate, part int) string {
	// If part 1 empty seat if 4 occupied, if part 2 empty seat if 5 occupied
	adjacentThreshold := 3 + part
	adjacentOccupied := 0
	for _, d := range directions {
		startingCo := co
		for {
			startingCo = helpers.Coordinate{
				X: startingCo.X + d.X,
				Y: startingCo.Y + d.Y,
			}
			if val := (*g)[helpers.Coordinate{X: startingCo.X, Y: startingCo.Y}]; val != "." {
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
	if adjacentOccupied >= adjacentThreshold {
		return "L"
	}
	return "#"
}

func (g *Grid) generateNextGrid(part int) Grid {
	newGrid := Grid{}
	for co, val := range *g {
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

func (g *Grid) countOccupiedSeats() int {
	numOccupiedSeats := 0
	for _, val := range *g {
		if val == "#" {
			numOccupiedSeats++
		}
	}
	return numOccupiedSeats
}

func (g *Grid) findSolution(part int) {
	for {
		newGrid := g.generateNextGrid(part)
		if g.isSameGrid(newGrid) {
			fmt.Printf("Part %d: %d\n", part, g.countOccupiedSeats())
			return
		}
		g = &newGrid
	}
}

// For debugging
// func (g *Grid) printExampleGrid() {
// 	printGrid := [10][10]string{}
// 	for co, val := range *g {
// 		printGrid[co.Y][co.X] = val
// 	}
// 	for _, row := range printGrid {
// 		fmt.Println(row)
// 	}
// 	fmt.Println()
// }

func main() {
	plan := helpers.ReadFile()
	g := Grid{}
	for r, row := range plan {
		for c, col := range row {
			g[helpers.Coordinate{X: c, Y: r}] = string(col)
		}
	}
	g1 := g
	g1.findSolution(1)
	g2 := g
	g2.findSolution(2)
}
