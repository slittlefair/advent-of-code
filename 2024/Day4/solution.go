package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"fmt"
)

func parseInput(input []string) *graph.Grid {
	g := graph.NewGrid()
	g.MaxY = len(input)
	g.MaxX = len(input[0])

	for y, line := range input {
		for x, char := range line {
			g.Graph[graph.Co{X: x, Y: y}] = string(char)
		}
	}

	return g
}

var expectedSequence = []string{"M", "A", "S"}

// Recursively look for coordinates in a grid from a certain point and check we find a sequence. If
// at any point we fail to follow the sequence we immediately return false all the way back to the
// original call
func traverseCos(g *graph.Grid, currentCo, dir graph.Co, currentIndex int) bool {
	adjCo := graph.Co{X: currentCo.X + dir.X, Y: currentCo.Y + dir.Y}
	v, ok := g.Graph[adjCo]

	// If we've gone outside of the grid, return false
	if !ok {
		return false
	}

	// If the value isn't what we expect it to be, return false
	if v != expectedSequence[currentIndex] {
		return false
	}

	// If we've found another expected match, and we're at the end of expected sequence, we've
	// found the word
	if currentIndex == len(expectedSequence) {
		return true
	}

	return traverseCos(g, adjCo, dir, currentIndex+1)
}

// Part 1
func findXmas(g *graph.Grid, co graph.Co) int {
	count := 0
	if g.Graph[co] != "X" {
		return count
	}

	// For each "direction" from the current coordinate, traverse the grid and see if we find the
	// expected sequence. If we fail before we reach the end we'll return false
	adjCos := graph.AdjacentCos(graph.Co{X: 0, Y: 0}, true)
	for _, dir := range adjCos {
		found := traverseCos(g, co, dir, 0)
		if found {
			count++
		}
	}

	return count
}

var diagonals = []graph.Co{
	{X: -1, Y: -1},
	{X: -1, Y: 1},
	{X: 1, Y: -1},
	{X: 1, Y: 1},
}

// Part 2
func findMasInX(g *graph.Grid, co graph.Co) bool {
	// If the center character isn't A, move on
	if g.Graph[co] != "A" {
		return false
	}

	// Find frequency of characters in the four diagonals.
	charFreq := map[string]int{}
	for _, diag := range diagonals {
		co := graph.Co{X: co.X + diag.X, Y: co.Y + diag.Y}
		charFreq[g.Graph[co]]++
	}

	// For it to be a valid X-MAS we need two Ms and 2 Ss...
	if charFreq["M"] != 2 || charFreq["S"] != 2 {
		return false
	}

	// ... and for opposite corners to be different from each other
	co1 := graph.Co{X: co.X - 1, Y: co.Y - 1}
	co2 := graph.Co{X: co.X + 1, Y: co.Y + 1}
	return g.Graph[co1] != g.Graph[co2]
}

func traverseGrid(g *graph.Grid) (int, int) {
	part1, part2 := 0, 0
	for y := g.MinY; y <= g.MaxY; y++ {
		for x := g.MinX; x <= g.MaxX; x++ {
			currentCo := graph.Co{X: x, Y: y}
			// Part 1
			part1 += findXmas(g, currentCo)

			if findMasInX(g, currentCo) {
				part2++
			}
		}
	}
	return part1, part2
}

func findSolutions(input []string) (int, int) {
	grid := parseInput(input)
	return traverseGrid(grid)
}

func main() {
	input := file.Read()

	part1, part2 := findSolutions(input)
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}
