package main

import (
	helpers "Advent-of-Code"
	"fmt"
)

type Grid map[helpers.Coordinate]string

func (g *Grid) isSameGrid(newGrid Grid) bool {
	for co, val := range *g {
		if newGrid[co] != val {
			return false
		}
	}
	return true
}

func (g *Grid) evaluateEmptySeatPart1(co helpers.Coordinate) string {
	for x := co.X - 1; x <= co.X+1; x++ {
		for y := co.Y - 1; y <= co.Y+1; y++ {
			if !(x == co.X && y == co.Y) && (*g)[helpers.Coordinate{X: x, Y: y}] == "#" {
				return "L"
			}
		}
	}
	return "#"
}

func (g *Grid) evaluateOccupiedSeatPart1(co helpers.Coordinate) string {
	adjacentOccupied := 0
	for x := co.X - 1; x <= co.X+1; x++ {
		for y := co.Y - 1; y <= co.Y+1; y++ {
			if !(x == co.X && y == co.Y) && (*g)[helpers.Coordinate{X: x, Y: y}] == "#" {
				adjacentOccupied++
			}
		}
	}
	if adjacentOccupied >= 4 {
		return "L"
	}
	return "#"
}

func (g *Grid) generateNextGrid(part int) Grid {
	newGrid := Grid{}
	for co, val := range *g {
		if val == "L" {
			if part == 1 {
				newGrid[co] = g.evaluateEmptySeatPart1(co)
			}
		} else if val == "#" {
			if part == 1 {
				newGrid[co] = g.evaluateOccupiedSeatPart1(co)
			}
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
			fmt.Printf("Part %d:%d\n", part, g.countOccupiedSeats())
			return
		}
		g = &newGrid
	}
}

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
}
