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

// func (tm TreeMap) traverseDirection() {

// }

func (tm TreeMap) countVisibleTrees() (int, int) {
	sum := 0
	highestScenicScore := 0
	for co, v := range tm.graph {
		// if a tree is on the edge then it is visible
		if co.X == 0 || co.X == tm.maxX || co.Y == 0 || co.Y == tm.maxY {
			sum++
		} else {
			visibleDirections := [4]bool{true, true, true, true}
			scenicScore := 1

			// Go left
			dist := 0
			blockedAt := 0
			for x := co.X - 1; x >= 0; x-- {
				dist++
				if tm.graph[graph.Co{X: x, Y: co.Y}] >= v {
					visibleDirections[0] = false
					if blockedAt == 0 {
						blockedAt = dist
					}
				}
			}
			if blockedAt == 0 {
				blockedAt = dist
			}
			scenicScore *= blockedAt

			// Go right
			dist = 0
			blockedAt = 0
			for x := co.X + 1; x <= tm.maxX; x++ {
				dist++
				if tm.graph[graph.Co{X: x, Y: co.Y}] >= v {
					visibleDirections[1] = false
					if blockedAt == 0 {
						blockedAt = dist
					}
				}
			}
			if blockedAt == 0 {
				blockedAt = dist
			}
			scenicScore *= blockedAt

			// Go up
			dist = 0
			blockedAt = 0
			for y := co.Y - 1; y >= 0; y-- {
				dist++
				if tm.graph[graph.Co{X: co.X, Y: y}] >= v {
					visibleDirections[2] = false
					if blockedAt == 0 {
						blockedAt = dist
					}
				}
			}
			if blockedAt == 0 {
				blockedAt = dist
			}
			scenicScore *= blockedAt

			// Go down
			dist = 0
			blockedAt = 0
			for y := co.Y + 1; y <= tm.maxY; y++ {
				dist++
				if tm.graph[graph.Co{X: co.X, Y: y}] >= v {
					visibleDirections[3] = false
					if blockedAt == 0 {
						blockedAt = dist
					}
				}
			}
			if blockedAt == 0 {
				blockedAt = dist
			}
			scenicScore *= blockedAt

			for _, b := range visibleDirections {
				if b {
					sum++
					break
				}
			}
			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}
	return sum, highestScenicScore
}

func main() {
	input := file.Read()
	tm := parseInput(input)
	part1, part2 := tm.countVisibleTrees()
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
