package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"Advent-of-Code/regex"
	"Advent-of-Code/slice"
	"fmt"
	"slices"
	"sort"
	"strconv"
)

type Playground struct {
	junctions       []graph.Co
	connections     []map[graph.Co]struct{}
	distances       map[float64][]graph.Co
	sortedDistances []float64
}

func parseInput(input []string) *Playground {
	p := &Playground{
		junctions:   make([]graph.Co, len(input)),
		connections: make([]map[graph.Co]struct{}, 0),
		distances:   map[float64][]graph.Co{},
	}
	for i, line := range input {
		matches := regex.MatchNums.FindAllString(line, 3)
		x, _ := strconv.Atoi(matches[0])
		y, _ := strconv.Atoi(matches[1])
		z, _ := strconv.Atoi(matches[2])
		co := graph.Co{X: x, Y: y, Z: z}
		p.junctions[i] = co

		for j := range i {
			co2 := p.junctions[j]
			p.distances[graph.CalculateEuclideanDistance(co, co2)] = []graph.Co{co, co2}
		}
	}

	sortedDistances := make([]float64, 0, len(p.distances))
	for k := range p.distances {
		sortedDistances = append(sortedDistances, k)
	}
	slices.Sort(sortedDistances)
	p.sortedDistances = sortedDistances

	return p
}

func (pg *Playground) findJunctionCircuit(co graph.Co) func(mp map[graph.Co]struct{}) bool {
	return func(mp map[graph.Co]struct{}) bool {
		if _, ok := mp[co]; ok {
			return true
		}
		return false
	}
}

func (pg *Playground) joinJunctions() (graph.Co, graph.Co) {
	junctions := pg.distances[pg.sortedDistances[0]]
	co1 := junctions[0]
	co2 := junctions[1]
	idx1 := slices.IndexFunc(pg.connections, pg.findJunctionCircuit(co1))
	idx2 := slices.IndexFunc(pg.connections, pg.findJunctionCircuit(co2))

	if idx1 == -1 && idx2 == -1 {
		pg.connections = append(pg.connections, map[graph.Co]struct{}{
			co1: {},
			co2: {},
		})
	} else if idx1 == -1 {
		pg.connections[idx2][co1] = struct{}{}
	} else if idx2 == -1 {
		pg.connections[idx1][co2] = struct{}{}
	} else if idx1 != idx2 {
		for k := range pg.connections[idx2] {
			pg.connections[idx1][k] = struct{}{}
		}
		pg.connections = slice.RemoveByIndex(pg.connections, idx2)
	}
	pg.sortedDistances = pg.sortedDistances[1:]
	return co1, co2
}

func (pg *Playground) part1Solution() int {
	lengths := make([]int, len(pg.connections))
	for i, v := range pg.connections {
		lengths[i] = len(v)
	}
	sort.Slice(lengths, func(i, j int) bool {
		return lengths[i] > lengths[j]
	})
	return lengths[0] * lengths[1] * lengths[2]
}

func findSolutions(input []string) (part1 int, part2 int) {
	playground := parseInput(input)
	i := 1
	for {
		co1, co2 := playground.joinJunctions()
		if i == 1000 {
			part1 = playground.part1Solution()
		}
		i++
		if len(playground.connections) == 1 && len(playground.connections[0]) == len(playground.junctions) {
			return part1, co1.X * co2.X
		}
	}
}

func main() {
	input := file.Read()
	part1, part2 := findSolutions(input)
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}
