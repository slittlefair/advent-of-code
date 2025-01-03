package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"fmt"
)

type AntennaMap struct {
	graph.Grid

	antennas       map[string]map[graph.Co]bool
	antinodes      map[graph.Co]bool
	antinodesPart2 map[graph.Co]bool
}

// Print is a function for printing a grid of antennas and their antinodes. Where an antenna and
// antinode occupy the same space, we print the antenna.
// Print takes a part too, so we print either the part 1 or part 2 antinodes. If the given part
// isn't one of 1 or 2 we return an error without printing.
func (am AntennaMap) Print(part int) error {
	if part != 1 && part != 2 {
		return fmt.Errorf("expected part 1 or 2, got %d", part)
	}
	g := am.Grid
	for y := 0; y <= g.MaxY; y++ {
		for x := 0; x <= g.MaxX; x++ {
			co := graph.Co{X: x, Y: y}
			if v, ok := g.Graph[co]; ok {
				fmt.Print(v)
			} else if part == 1 {
				if _, ok := am.antinodes[co]; ok {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
				continue
			} else if _, ok := am.antinodesPart2[co]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	return nil
}

// Parse the input to create a grid of antennas
func parseInput(input []string) AntennaMap {
	am := AntennaMap{
		Grid: graph.Grid{
			Graph: make(graph.Graph),
			MaxY:  len(input) - 1,
			MaxX:  len(input[0]) - 1,
		},
		antennas:       make(map[string]map[graph.Co]bool),
		antinodes:      make(map[graph.Co]bool),
		antinodesPart2: make(map[graph.Co]bool),
	}
	for y, line := range input {
		for x, char := range line {
			s := string(char)
			if s == "." {
				continue
			}
			co := graph.Co{X: x, Y: y}
			am.Graph[co] = s

			if _, ok := am.antennas[s]; !ok {
				am.antennas[s] = make(map[graph.Co]bool)
			}
			am.antennas[s][co] = true
		}
	}
	return am
}

// Find the antinodes for each pair of same-type antennas and store them in maps for parts 1 and 2,
// returning the respective lengths of the maps. Using a map rather than count ensures for each part
// we are taking unique locations, in the case antennas occupy the same coordinate.
func (am AntennaMap) findAntinodes() (int, int) {
	for _, locs := range am.antennas {
		for l := range locs {
			for otherLoc := range locs {
				// We can't evaluate an antenna against itself
				if l == otherLoc {
					continue
				}
				xDiff := l.X - otherLoc.X
				yDiff := l.Y - otherLoc.Y
				i := 0
				// Whilst we're in the given grid, keep finding antinodes. Once we break out of the
				// given range we can break out and stop looking.
				for {
					co := graph.Co{
						X: l.X + (i * xDiff),
						Y: l.Y + (i * yDiff),
					}
					if co.X < am.MinX || co.X > am.MaxX || co.Y < am.MinY || co.Y > am.MaxY {
						break
					}
					am.antinodesPart2[co] = true
					// Part 1 only cares about the antinode being twice the distance from the
					// antenna as the other, so part 1 is only interested where i is 1.
					if i == 1 {
						am.antinodes[co] = true
					}
					i++
				}
			}
		}
	}
	return len(am.antinodes), len(am.antinodesPart2)
}

func findSolutions(input []string) (int, int) {
	am := parseInput(input)
	return am.findAntinodes()
}

func main() {
	input := file.Read()
	part1, part2 := findSolutions(input)
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}
