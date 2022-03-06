package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	djk "Advent-of-Code/graph/dijkstra"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parseInput(input []string) (int, graph.Co, error) {
	num, err := strconv.Atoi(input[0])
	if err != nil {
		return -1, graph.Co{}, err
	}
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(input[1], -1)
	x, _ := strconv.Atoi(matches[0])
	y, _ := strconv.Atoi(matches[1])
	return num, graph.Co{X: x, Y: y}, nil
}

func isSpace(co graph.Co, num int) bool {
	x, y := co.X, co.Y
	eq := x*x + 3*x + 2*x*y + y + y*y
	bin := strconv.FormatInt(int64(eq+num), 2)
	return strings.Count(bin, "1")%2 == 0
}

func populateOffice(target graph.Co, num int) *djk.Graph {
	office := djk.NewGraph(51, 51)
	for y := 0; y <= target.Y+10; y++ {
		for x := 0; x <= target.X+10; x++ {
			co := graph.Co{X: x, Y: y}
			if isSpace(co, num) {
				office.Grid[co] = 1
			}
		}
	}
	for co, risk := range office.Grid {
		for _, adjCo := range graph.AdjacentCos(co, false) {
			if _, ok := office.Grid[adjCo]; ok {
				office.AddEdge(adjCo, co, risk)
			}
		}
	}
	return office
}

func findSolutions(input []string) (int, int, error) {
	num, target, err := parseInput(input)
	if err != nil {
		return -1, -1, err
	}
	office := populateOffice(target, num)
	origin := graph.Co{X: 1, Y: 1}
	path1, err := office.GetPath(origin, target)
	if err != nil {
		return -1, -1, err
	}
	distinctCos := map[graph.Co]struct{}{}
	for y := 0; y <= office.MaxY; y++ {
		for x := 0; x <= office.MaxX; x++ {
			path, _ := office.GetPath(origin, graph.Co{X: x, Y: y})
			if path != nil && path.Value <= 50 {
				for _, node := range path.Nodes {
					distinctCos[node] = struct{}{}
				}
			}
		}
	}
	return path1.Value, len(distinctCos), nil
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
