package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"fmt"
)

type TreeMap struct {
	graph map[graph.Co]int
	maxX  int
	maxY  int
}

func parseInput(input []string) TreeMap {
	tm := TreeMap{
		graph: make(map[graph.Co]int),
		maxY:  len(input) - 1,
	}
	for y, line := range input {
		for x, r := range line {
			tm.graph[graph.Co{X: x, Y: y}] = int(r - '0')
			tm.maxX = x
		}
	}
	return tm
}

// Create an "enum" (as close as Go allows) for directions which we use to let the traverse function
// know whether to run along the x or y axis
type Direction int

const (
	Horizontally Direction = iota
	Vertically
)

// Move from a particular coordinate to a particular coordinate and calculate whether it's visible
// and what the viewing distance is
func (tm TreeMap) travel(co graph.Co, from, to, change int, direction Direction) (bool, int) {
	dist := 0
	viewingDistance := 0
	visible := true
	for i := from; i != to; i = i + change {
		dist++
		x := co.X
		y := co.Y
		if direction == Horizontally {
			x = i
		} else {
			y = i
		}
		if tm.graph[graph.Co{X: x, Y: y}] >= tm.graph[co] {
			visible = false
			if viewingDistance == 0 {
				viewingDistance = dist
				break
			}
		}
	}
	// If viewing stance is still 0 we have got to the edge without being blocked
	if viewingDistance == 0 {
		viewingDistance = dist
	}
	return visible, viewingDistance
}

func (tm TreeMap) optimiseTreehouseLocation() (int, int) {
	visibleTrees := 0
	highestScenicScore := 0
	for co := range tm.graph {
		// If a tree is on the edge then it is visible and has at least one viewing distance of 0,
		// giving it a scenic score of 0 and making it not the highest
		if co.X == 0 || co.X == tm.maxX || co.Y == 0 || co.Y == tm.maxY {
			visibleTrees++
		} else {
			var visible bool
			var viewingDistance int
			visibleDirections := [4]bool{}
			scenicScore := 1

			// Go left
			visible, viewingDistance = tm.travel(co, co.X-1, -1, -1, Horizontally)
			visibleDirections[0] = visible
			scenicScore *= viewingDistance

			// Go right
			visible, viewingDistance = tm.travel(co, co.X+1, tm.maxX+1, 1, Horizontally)
			visibleDirections[1] = visible
			scenicScore *= viewingDistance

			// Go up
			visible, viewingDistance = tm.travel(co, co.Y-1, -1, -1, Vertically)
			visibleDirections[2] = visible
			scenicScore *= viewingDistance

			// Go down
			visible, viewingDistance = tm.travel(co, co.Y+1, tm.maxY+1, 1, Vertically)
			visibleDirections[3] = visible
			scenicScore *= viewingDistance

			// Check all directions and if any of them are visible then the tree is visible
			for _, b := range visibleDirections {
				if b {
					visibleTrees++
					break
				}
			}

			// Compare the tree's scenic score with the current highest and set it if it's higher
			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}

	return visibleTrees, highestScenicScore
}

func main() {
	input := file.Read()
	tm := parseInput(input)
	part1, part2 := tm.optimiseTreehouseLocation()
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
