package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	djk "Advent-of-Code/graph/dijkstra"
	"Advent-of-Code/maths"
	"fmt"
)

func parseInput(input []string) (*djk.Graph, graph.Co, graph.Co) {
	var origin, destination graph.Co

	// Set up terrain graph to be populated with nodes and edges
	terrain := djk.NewGraph(len(input[0])-1, len(input)-1)

	// Loop through input and create nodes for each letter, using ints as their value for easy
	// comparison ("a" = 1, "b" = 2, ... "z" = 26)
	for y, line := range input {
		for x, r := range line {
			co := graph.Co{X: x, Y: y}
			val := int(r - 96)
			if s := string(r); s == "S" {
				val = 1
				origin = co
			} else if s == "E" {
				val = 26
				destination = co
			}
			terrain.Grid[co] = val
		}
	}

	// Loop back through each node and attempt to create edges between it and all adjacent,
	// non-diagonal nodes. If node A has adjacent node B and an edge is made, the reverse edge
	// is considered when we reach node B and attempt to create an edge to adjacent node B.
	//
	// An edge is created from node A to node B only if the height (value) of B is at most one
	// higher than that of node A.
	//
	// When determining path length we only care about how many nodes we visit, and each move from
	// one node to another adds 1 to the path length, so every edge should have weight of 1.
	for co, height := range terrain.Grid {
		for _, adjCo := range graph.AdjacentCos(co, false) {
			if adjHeight, ok := terrain.Grid[adjCo]; ok && height+1 >= adjHeight {
				terrain.AddEdge(co, adjCo, 1)
			}
		}
	}

	return terrain, origin, destination
}

func findSolutions(input []string) (int, int, error) {
	// Part 1
	terrain, origin, destination := parseInput(input)
	path, err := terrain.GetPath(origin, destination)
	if err != nil {
		return -1, -1, err
	}

	// Part 2, try and find a path from every part of the grid that has minimum height (1), and if
	// a path exists compare its length to the shortest path found so far
	shortestPathLength := maths.Infinity
	for co, height := range terrain.Grid {
		if height == 1 {
			// It's possible a path doesn't exist from the given starting point, so ignore the error
			// and just check if path is not nil
			path, _ := terrain.GetPath(co, destination)
			if path != nil && path.Value < shortestPathLength {
				shortestPathLength = path.Value
			}
		}
	}

	return path.Value, shortestPathLength, nil
}

func main() {
	input := file.Read()
	part1, part2, err := findSolutions(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
