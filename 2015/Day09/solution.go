package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"strconv"
	"strings"
)

type Graph struct {
	Edges []Edge
	Nodes []string
	Paths [][]string
}

type Edge struct {
	Parent string
	Child  string
	Cost   int
}

func (g *Graph) addNode(n string) {
	for _, node := range g.Nodes {
		if node == n {
			return
		}
	}
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) addEdge(parent, child string, cost int) {
	edge := Edge{
		Parent: parent,
		Child:  child,
		Cost:   cost,
	}

	g.Edges = append(g.Edges, edge)
	g.addNode(parent)
	g.addNode(child)
}

func (g *Graph) parseInput(input []string) error {
	for _, line := range input {
		split := strings.Split(line, " ")
		cost, err := strconv.Atoi(split[4])
		if err != nil {
			return err
		}
		node1 := split[0]
		node2 := split[2]
		g.addEdge(node1, node2, cost)
	}
	nodes := make([]string, len(g.Nodes))
	copy(nodes, g.Nodes)
	g.Paths = utils.Permutations(nodes)
	return nil
}

func (g *Graph) getDistanceOfPath(path []string) int {
	distance := 0
	for i := 0; i < len(path)-1; i++ {
		locationA := path[i]
		locationB := path[i+1]
		for _, edge := range g.Edges {
			if (edge.Child == locationA && edge.Parent == locationB) || (edge.Child == locationB && edge.Parent == locationA) {
				distance += edge.Cost
				break
			}
		}
	}
	return distance
}

func (g *Graph) findMinimumAndMaximumPaths() (int, int) {
	minPathDistance := int(^uint(0) >> 1)
	maxPathDistance := 0
	for _, path := range g.Paths {
		dist := g.getDistanceOfPath(path)
		if dist < minPathDistance {
			minPathDistance = dist
		}
		if dist > maxPathDistance {
			maxPathDistance = dist
		}
	}
	return minPathDistance, maxPathDistance
}

func main() {
	input := utils.ReadFile()
	graph := Graph{}
	graph.parseInput(input)
	minPathDistance, maxPathDistance := graph.findMinimumAndMaximumPaths()
	fmt.Println("Part 1:", minPathDistance)
	fmt.Println("Part 2:", maxPathDistance)
}
