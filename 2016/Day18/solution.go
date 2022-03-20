package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"fmt"
)

type grid struct {
	tiles     map[graph.Co]bool
	width     int
	safeTiles int
}

func createGrid(input string) grid {
	g := grid{
		tiles:     map[graph.Co]bool{},
		width:     len(input),
		safeTiles: 0,
	}
	for i, val := range input {
		co := graph.Co{X: i, Y: 0}
		g.tiles[co] = string(val) == "^"
		if string(val) != "^" {
			g.safeTiles++
		}
	}
	return g
}

func (g *grid) isTrap(x int) bool {
	lVal, lOk := g.tiles[graph.Co{X: x - 1}]
	rVal, rOk := g.tiles[graph.Co{X: x + 1}]

	if (!lOk || !lVal) == (!rOk || !rVal) {
		g.safeTiles++
		return false
	}
	return true
}

func (g *grid) assessRow() {
	newTiles := map[graph.Co]bool{}
	for x := 0; x < g.width; x++ {
		newTiles[graph.Co{X: x}] = g.isTrap(x)
	}
	g.tiles = newTiles
}

func findSolutions(input string, part1, part2 int) (int, int) {
	g := createGrid(input)
	var part1Sol int
	for i := 0; i < part2-1; i++ {
		if i == part1-1 {
			part1Sol = g.safeTiles
		}
		g.assessRow()
	}
	return part1Sol, g.safeTiles
}

func main() {
	input := file.Read()[0]
	part1, part2 := findSolutions(input, 40, 400000)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
