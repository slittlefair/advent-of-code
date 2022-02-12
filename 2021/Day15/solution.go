package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	djk "Advent-of-Code/graph/dijkstra"
	"fmt"
	"regexp"
	"strconv"
)

func parseInput(input []string, factor int) *djk.Graph {
	cave := djk.NewGraph(len(input[0])-1, len(input)-1)
	re := regexp.MustCompile(`\d`)
	for y, line := range input {
		matches := re.FindAllString(line, -1)
		for x, m := range matches {
			// All matches can be converted to an int due to regex matching, so we know we won't get an error
			n, _ := strconv.Atoi(m)
			cave.Grid[graph.Co{X: x, Y: y}] = n
		}
	}
	cave.ExtendGrid(factor)
	for co, risk := range cave.Grid {
		for _, adjCo := range graph.AdjacentCos(co, false) {
			if _, ok := cave.Grid[adjCo]; ok {
				cave.AddEdge(adjCo, co, risk)
			}
		}
	}
	return cave
}

func findSolutions(input []string) (int, int, error) {
	// Part 1
	cave := parseInput(input, 1)
	origin := graph.Co{X: 0, Y: 0}
	destination := graph.Co{X: cave.MaxX, Y: cave.MaxY}
	path1, err := cave.GetPath(origin, destination)
	if err != nil {
		return -1, -1, err
	}

	// Part 2
	cave = parseInput(input, 5)
	origin = graph.Co{X: 0, Y: 0}
	destination = graph.Co{X: cave.MaxX, Y: cave.MaxY}
	path2, err := cave.GetPath(origin, destination)
	if err != nil {
		return -1, -1, err
	}
	return path1.Value, path2.Value, nil
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
