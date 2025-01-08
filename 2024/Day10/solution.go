package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"fmt"
	"strconv"
)

type TopMap struct {
	graph.Grid[int]
	trails map[graph.Co]map[graph.Co]bool
}

func parseInput(input []string) (*TopMap, error) {
	tm := &TopMap{
		Grid: graph.Grid[int]{
			MaxY:  len(input) - 1,
			MaxX:  len(input[0]) - 1,
			Graph: make(map[graph.Co]int),
		},
		trails: make(map[graph.Co]map[graph.Co]bool),
	}
	for y, line := range input {
		for x, char := range line {
			n, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}
			tm.Graph[graph.Co{X: x, Y: y}] = n
		}
	}
	return tm, nil
}

func (tm *TopMap) move(trailHead graph.Co, currentVal int, newCo graph.Co) {
	if tm.OutOfBounds(newCo) {
		return
	}

	newVal := tm.Graph[newCo]

	if newVal-currentVal != 1 {
		return
	}

	if newVal == 9 {
		if _, ok := tm.trails[trailHead]; !ok {
			tm.trails[trailHead] = map[graph.Co]bool{}
		}
		tm.trails[trailHead][newCo] = true
		return
	}

	adjCos := graph.AdjacentCos(newCo, false)
	for _, co := range adjCos {
		tm.move(trailHead, newVal, co)
	}
}

func (tm *TopMap) findTrails() {
	for co, v := range tm.Graph {
		if v == 0 {
			adjCos := graph.AdjacentCos(co, false)
			for _, aCo := range adjCos {
				tm.move(co, 0, aCo)
			}
		}
	}
}

func main() {
	input := file.Read()
	topMap, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	topMap.Grid.PrintGrid()

	topMap.findTrails()
	count := 0
	for _, v := range topMap.trails {
		count += len(v)
	}
	fmt.Printf("Part1: %v\n", count)
}
