package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"fmt"
)

type Beams struct {
	graph.Grid[string]
	beams map[graph.Co]int
	part1 int
}

func parseInput(input []string) *Beams {
	b := &Beams{
		Grid: graph.Grid[string]{
			Graph: graph.Graph[string]{},
			MinX:  0,
			MinY:  0,
			MaxY:  len(input) - 1,
			MaxX:  len(input[0]) - 1,
		},
		beams: map[graph.Co]int{},
	}
	for y := range len(input) {
		line := input[y]
		for x, r := range line {
			co := graph.Co{X: x, Y: y}
			b.Graph[co] = string(r)
			if string(r) == "S" {
				b.beams[co] = 1
			}
		}
	}
	return b
}

func (b *Beams) stepBeam(beam graph.Co) []graph.Co {
	newBeam := graph.Co{X: beam.X, Y: beam.Y + 1}
	if b.Graph[newBeam] == "^" {
		return []graph.Co{
			{X: newBeam.X + 1, Y: newBeam.Y},
			{X: newBeam.X - 1, Y: newBeam.Y},
		}
	}
	return []graph.Co{newBeam}
}

func (b *Beams) step() {
	newBeams := make(map[graph.Co]int)
	for beam, freq := range b.beams {
		new := b.stepBeam(beam)
		for _, n := range new {
			newBeams[n] += freq
		}
		if len(new) > 1 {
			b.part1++
		}
	}
	b.beams = newBeams
}

func findSolutions(input []string) (int, int) {
	b := parseInput(input)
	for range b.MaxY {
		b.step()
	}
	part2 := 0
	for _, v := range b.beams {
		part2 += v
	}
	return b.part1, part2
}

func main() {
	input := file.Read()
	part1, part2 := findSolutions(input)
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}
