package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"Advent-of-Code/maths"
	"Advent-of-Code/regex"
	"Advent-of-Code/timer"
	"fmt"
	"slices"
	"strconv"
	"time"
)

type bounds struct {
	min int
	max int
}

type Theater struct {
	*graph.Grid[string]
	tiles       []graph.Co
	wallLookupY map[int]*bounds
	wallLookupX map[int]*bounds
}

func (t *Theater) populateBounds(x, y int) {
	if t.wallLookupY[y] == nil {
		t.wallLookupY[y] = &bounds{
			min: maths.Infinity,
		}
	}
	yBounds := t.wallLookupY[y]
	yBounds.min = maths.Min(yBounds.min, x)
	yBounds.max = maths.Max(yBounds.max, x)
	if t.wallLookupX[x] == nil {
		t.wallLookupX[x] = &bounds{
			min: maths.Infinity,
		}
	}
	xBounds := t.wallLookupX[x]
	xBounds.min = maths.Min(xBounds.min, y)
	xBounds.max = maths.Max(xBounds.max, y)
}

func parseInput(input []string) *Theater {
	tr := &Theater{
		Grid:        graph.NewGrid[string](),
		tiles:       make([]graph.Co, len(input)),
		wallLookupY: make(map[int]*bounds),
		wallLookupX: make(map[int]*bounds),
	}
	tr.MinX = maths.Infinity
	tr.MinY = maths.Infinity
	for i, line := range input {
		matches := regex.MatchNums.FindAllString(line, 2)
		x, _ := strconv.Atoi(matches[0])
		y, _ := strconv.Atoi(matches[1])
		tr.Graph[graph.Co{X: x, Y: y}] = "#"
		tr.tiles[i] = graph.Co{X: x, Y: y}
		tr.populateBounds(x, y)
		tr.MaxX = maths.Max(tr.MaxX, x)
		tr.MaxY = maths.Max(tr.MaxY, y)
		tr.MinX = maths.Min(tr.MinX, x)
		tr.MinY = maths.Min(tr.MinY, y)
	}

	// Populate green tiles
	for i, t1 := range tr.tiles {
		j := i + 1
		if j == len(tr.tiles) {
			j = 0
		}
		t2 := tr.tiles[j]

		if t1.X == t2.X {
			yCos := []int{t1.Y, t2.Y}
			slices.Sort(yCos)
			for y := yCos[0] + 1; y < yCos[1]; y++ {
				x := t1.X
				tr.Graph[graph.Co{X: x, Y: y}] = "X"
				tr.populateBounds(x, y)
			}
			continue
		}

		xCos := []int{t1.X, t2.X}
		slices.Sort(xCos)
		for x := xCos[0] + 1; x < xCos[1]; x++ {
			y := t1.Y
			tr.Graph[graph.Co{X: x, Y: y}] = "X"
			tr.populateBounds(x, y)
		}
	}

	tilesWalls := make([]graph.Co, 0, len(tr.Graph))
	for v := range tr.Graph {
		tilesWalls = append(tilesWalls, v)
	}

	return tr
}

func (t *Theater) coordinateInside(co graph.Co, checkX, checkY bool) bool {
	if _, ok := t.Graph[co]; ok {
		return true
	}
	if checkX {
		if co.Y < t.wallLookupX[co.X].min || co.Y > t.wallLookupX[co.X].max {
			return false
		}
	}
	if checkY {
		if co.X < t.wallLookupY[co.Y].min || co.X > t.wallLookupY[co.Y].max {
			return false
		}
	}
	return true
}

func findSolutions(input []string) (part1 int, part2 int) {
	t := parseInput(input)
	for i := range len(t.tiles) {
		if i%20 == 0 {
			fmt.Println(i)
		}
		for j := range i {
			co1 := t.tiles[i]
			co2 := t.tiles[j]
			diffX := maths.Abs(co1.X-co2.X) + 1
			diffY := maths.Abs(co1.Y-co2.Y) + 1
			part1 = maths.Max(part1, diffX*diffY)

			xCos := []int{co1.X, co2.X}
			slices.Sort(xCos)
			yCos := []int{co1.Y, co2.Y}
			slices.Sort(yCos)
			if diffX*diffY < part2 {
				goto out
			}

			for y := yCos[0] + 1; y < yCos[1]; y++ {
				if !t.coordinateInside(graph.Co{X: xCos[0], Y: y}, true, true) {
					goto out
				}
				if !t.coordinateInside(graph.Co{X: xCos[1], Y: y}, true, false) {
					goto out
				}
			}
			for x := xCos[0] + 1; x < xCos[1]; x++ {
				if !t.coordinateInside(graph.Co{X: x, Y: yCos[0]}, true, true) {
					goto out
				}
				if !t.coordinateInside(graph.Co{X: x, Y: yCos[1]}, false, true) {
					goto out
				}
			}

			part2 = maths.Max(part2, diffX*diffY)
		out:
		}
	}
	return part1, part2
}

func main() {
	// TODO optimise this - part 2 causes run to yield correct solution but takes ~1m07s.
	// Feels like current approach is optimised as much as possible so maybe total rethink and
	// refactor is needed
	t := time.Now()
	input := file.Read()
	part1, part2 := findSolutions(input)
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
	timer.Track(t)
}
