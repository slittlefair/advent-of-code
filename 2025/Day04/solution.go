package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"fmt"
)

func parseInput(input []string) *graph.Grid[string] {
	g := graph.NewGrid[string]()
	g.MaxY = len(input) - 1
	g.MaxX = len(input[0]) - 1
	for y, line := range input {
		for x, char := range line {
			if c := string(char); c == "@" {
				g.Graph[graph.Co{X: x, Y: y}] = c
			}
		}
	}
	return g
}

func removePaper(g *graph.Grid[string]) int {
	rollsToRemove := []graph.Co{}
	for co := range g.Graph {
		adjCount := 0
		for _, adjCo := range graph.AdjacentCos(co, true) {
			if _, ok := g.Graph[adjCo]; ok {
				adjCount++
			}
		}
		if adjCount < 4 {
			rollsToRemove = append(rollsToRemove, co)
		}
	}
	for _, co := range rollsToRemove {
		delete(g.Graph, co)
	}
	return len(rollsToRemove)
}

func findSolutions(input []string) (int, int) {
	part1 := 0
	part2 := 0
	g := parseInput(input)
	for {
		removedCount := removePaper(g)
		if part1 == 0 {
			part1 = removedCount
		}
		if removedCount == 0 {
			break
		}
		part2 += removedCount
	}
	return part1, part2
}

func main() {
	input := file.Read()
	part1, part2 := findSolutions(input)
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}
