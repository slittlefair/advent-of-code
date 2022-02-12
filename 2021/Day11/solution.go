package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"fmt"
)

type Grid map[graph.Co]int

func parseInput(input []string) Grid {
	g := Grid{}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			g[graph.Co{X: x, Y: y}] = int(input[y][x] - '0')
		}
	}
	return g
}

// Debugging
// func (g Grid) print(step int) {
// 	fmt.Println("step:", step)
// 	for y := 0; y < 10; y++ {
// 		for x := 0; x < 10; x++ {
// 			fmt.Print(g[graph.Co{X: x, Y: y}])
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println()
// }

func (g Grid) followStep() int {
	for co, v := range g {
		g[co] = v + 1
	}
	flashTally := map[graph.Co]struct{}{}
	flashes := 0
	for {
		bursts := []graph.Co{}
		for co, v := range g {
			if _, ok := flashTally[co]; !ok && v > 9 {
				flashTally[co] = struct{}{}
				bursts = append(bursts, co)
			}
		}
		if len(bursts) == 0 {
			break
		}
		for _, co := range bursts {
			adjCos := graph.AdjacentCos(co, true)
			for _, aCo := range adjCos {
				if _, ok := g[aCo]; ok {
					g[aCo]++
				}
			}
		}
		flashes += len(bursts)
	}
	for co, v := range g {
		if v > 9 {
			g[co] = 0
		}
	}
	return flashes
}

func (g Grid) isSynchronised() bool {
	for _, v := range g {
		if v != 0 {
			return false
		}
	}
	return true
}

func findSolution(input []string) (int, int) {
	g := parseInput(input)
	flashes := 0
	i := 1
	var part1, part2 int
	for {
		flashes += g.followStep()
		if i == 100 {
			part1 = flashes
		}
		if g.isSynchronised() {
			part2 = i
			break
		}
		i++
	}
	return part1, part2
}

func main() {
	input := file.Read()
	part1, part2 := findSolution(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
