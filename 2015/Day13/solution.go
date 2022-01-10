package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"strconv"
	"strings"
)

type Graph struct {
	Edges []*Edge
	Nodes []string
	Paths [][]string
}

type Edge struct {
	Parent string
	Child  string
	Cost   int
}

func (g *Graph) AddNode(n string) {
	for _, node := range g.Nodes {
		if node == n {
			return
		}
	}
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(parent, child string, cost int) {
	for _, edge := range g.Edges {
		if edge.Parent == child && edge.Child == parent {
			edge.Cost += cost
			return
		}
	}
	edge := Edge{
		Parent: parent,
		Child:  child,
		Cost:   cost,
	}

	g.Edges = append(g.Edges, &edge)
	g.AddNode(parent)
	g.AddNode(child)
}

func (g *Graph) AddMe() {
	for _, node := range g.Nodes {
		g.AddEdge(node, "Me", 0)
	}
	nodes := make([]string, len(g.Nodes))
	copy(nodes, g.Nodes)
	g.Paths = utils.Permutations(nodes)
}

func (g *Graph) ParseInput(input []string) error {
	for _, line := range input {
		split := strings.Split(line, " ")
		cost, err := strconv.Atoi(split[3])
		if err != nil {
			return err
		}
		personA := split[0]
		personB := strings.TrimSuffix(split[10], ".")
		gain := split[2] == "gain"
		if !gain {
			cost *= -1
		}
		g.AddEdge(personA, personB, cost)
	}
	nodes := make([]string, len(g.Nodes))
	copy(nodes, g.Nodes)
	g.Paths = utils.Permutations(nodes)
	return nil
}

func (g *Graph) GetDistanceOfPath(path []string) int {
	distance := 0
	for i := 0; i < len(path); i++ {
		personA := path[i%len(path)]
		personB := path[(i+1)%len(path)]
		for _, edge := range g.Edges {
			if (edge.Child == personA && edge.Parent == personB) || (edge.Child == personB && edge.Parent == personA) {
				distance += edge.Cost
				break
			}
		}
	}
	return distance
}

func (g Graph) FindGreatestHappiness() int {
	greatestHappiness := 0
	for _, path := range g.Paths {
		dist := g.GetDistanceOfPath(path)
		if dist > greatestHappiness {
			greatestHappiness = dist
		}
	}
	return greatestHappiness
}

func main() {
	input := utils.ReadFile()
	graph := Graph{}
	graph.ParseInput(input)
	fmt.Println("Part 1:", graph.FindGreatestHappiness())
	graph.AddMe()
	fmt.Println("Part 2:", graph.FindGreatestHappiness())
}
