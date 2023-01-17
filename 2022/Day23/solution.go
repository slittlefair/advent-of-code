package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"fmt"
)

type Grove struct {
	minX, maxX, minY, maxY int
	elves                  map[graph.Co]bool
}

type Direction int

const (
	North Direction = iota
	South
	West
	East
)

func parseInput(input []string) *Grove {
	g := &Grove{
		minX: 0,
		maxX: len(input[0]) - 1,
		minY: 0,
		maxY: len(input) - 1,
	}
	elves := map[graph.Co]bool{}
	for y, line := range input {
		for x, r := range line {
			if string(r) == "#" {
				elves[graph.Co{X: x, Y: y}] = true
			}
		}
	}
	g.elves = elves
	return g
}

func (g *Grove) proposeMoves(currentDirection Direction) {
	newGrove := &Grove{
		minX:  g.minX,
		maxX:  g.maxX,
		minY:  g.minY,
		maxY:  g.maxY,
		elves: make(map[graph.Co]bool),
	}
	newPositions := map[graph.Co]graph.Co{}
	for elf := range g.elves {
		// get new co
		// newGrove.elves[co] = true
	}

}

func main() {
	input := file.Read()
	fmt.Println(input)
	grove := parseInput(input)
	fmt.Printf("%#+v\n", grove)
}
