package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"fmt"
)

type region struct {
	graph.Grid[string]
	value string
	edges []edge
}

type Farm struct {
	graph.Grid[string]
	regions   []*region
	regionCos map[graph.Co]bool
}

type edge struct {
	inside, outside graph.Co
	horizontal      bool
}

func parseInput(input []string) Farm {
	farm := Farm{
		Grid: graph.Grid[string]{
			MaxY:  len(input) - 1,
			MaxX:  len(input[0]) - 1,
			Graph: make(graph.Graph[string]),
		},
		regions:   []*region{},
		regionCos: make(map[graph.Co]bool),
	}
	for y, line := range input {
		for x, char := range line {
			s := string(char)
			farm.Graph[graph.Co{X: x, Y: y}] = s
		}
	}
	return farm
}

// Cycle through the plots and create regions of similar plots.
func (f *Farm) evaluateRegions() {
	for co, v := range f.Graph {
		if _, ok := f.regionCos[co]; ok {
			continue
		}
		r := &region{
			value: v,
			Grid: graph.Grid[string]{
				Graph: graph.Graph[string]{co: v},
			},
			edges: []edge{},
		}
		f.regionCos[co] = true
		adjCos := graph.AdjacentCos(co, false)
		for _, aCo := range adjCos {
			if f.Graph[aCo] != r.value {
				r.addEdge(co, aCo)
			}
			f.populateRegion(r, aCo)
		}
		f.regions = append(f.regions, r)
	}
}

// Recursive function to populate a region.
func (f *Farm) populateRegion(r *region, co graph.Co) {
	// If the given plot doesn't match the value of the region, return
	if f.Graph[co] != r.value {
		return
	}
	// If the given plot doesn't exist, return
	if _, ok := r.Graph[co]; ok {
		return
	}
	r.Graph[co] = r.value
	f.regionCos[co] = true
	// For adjacent plots, add an edge to the region if it's a boundary, that is if the adjacent
	// plot doesn't have the same value as the region. Otherwise, call the same function
	// recursively for each of the adjacent plots.
	adjCos := graph.AdjacentCos(co, false)
	for _, aCo := range adjCos {
		if f.Graph[aCo] != r.value {
			r.addEdge(co, aCo)
		} else {
			f.populateRegion(r, aCo)
		}
	}
}

// Calculate the number of sides of a region by collating edges into slices - each slice represents
// a side, that is a series of edges that are adjacent to each other. We keep combining sides if we
// find an element from each that are adjacent to each other, until we go through all sides without
// doing so, at which point we return the number of sides we found.
func (r *region) calculateNumSides() int {
	var horizontalMoves = []graph.Co{{X: 1}, {X: -1}}
	var verticalMoves = []graph.Co{{Y: -1}, {Y: 1}}
	var sides = [][]edge{}

	// Start by putting each edge into its own side
	for _, e := range r.edges {
		sides = append(sides, []edge{e})
	}

out:
	for i, s := range sides {
		for _, sideEdge := range s {
			// For each edge in a side, for its relevant moveset find the two edges that would be
			// adjacent to it
			adjEdges := []edge{}
			moveSet := verticalMoves
			if sideEdge.horizontal {
				moveSet = horizontalMoves
			}
			for _, m := range moveSet {
				adjEdges = append(adjEdges, edge{
					inside:     graph.Co{X: sideEdge.inside.X + m.X, Y: sideEdge.inside.Y + m.Y},
					outside:    graph.Co{X: sideEdge.outside.X + m.X, Y: sideEdge.outside.Y + m.Y},
					horizontal: sideEdge.horizontal,
				})
			}
			// Loop through each side again and see if any edges in those sides match any of the
			// current edge's adjacent edges.
			for j, ss := range sides {
				// Don't compare an edge to itself
				if i == j {
					continue
				}
				for _, sideEdge := range ss {
					for _, adjE := range adjEdges {
						// If we find an edge in a different side that is adjacent to an edge in the
						// current side, append the other side to the current one, then remove the
						// other one from the slice of sides.
						if adjE == sideEdge {
							sides[i] = append(s, ss...)
							sides = append(sides[:j], sides[j+1:]...)
							// Once we've combined two sides repeat the whole loop again
							goto out
						}
					}
				}
			}
		}
	}
	return len(sides)
}

func (r *region) addEdge(co, adjCo graph.Co) {
	r.edges = append(r.edges, edge{
		inside:     co,
		outside:    adjCo,
		horizontal: co.Y != adjCo.Y,
	})
}

func (f Farm) calculateCosts() (int, int) {
	part1 := 0
	part2 := 0
	for _, r := range f.regions {
		part1 += len(r.edges) * len(r.Graph)
		part2 += r.calculateNumSides() * len(r.Graph)
	}
	return part1, part2
}

func findSolutions(input []string) (int, int) {
	part1 := 0
	part2 := 0
	farm := parseInput(input)
	farm.evaluateRegions()
	part1, part2 = farm.calculateCosts()
	return part1, part2
}

func main() {
	input := file.Read()
	part1, part2 := findSolutions(input)
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}
