package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"sort"
)

type HeightMap map[utils.Co]int

type LowPoints map[utils.Co]int

type Basin map[utils.Co]struct{}
type Basins []int

func parseInput(input []string) HeightMap {
	hm := HeightMap{}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			hm[utils.Co{X: x, Y: y}] = int(input[y][x] - '0')
		}
	}
	return hm
}

func (hm HeightMap) findLowPoints() LowPoints {
	lowPoints := LowPoints{}
	for co, v := range hm {
		for _, adjCo := range utils.AdjacentCos(co, false) {
			if val, ok := hm[adjCo]; ok && val <= v {
				goto out
			}
		}
		lowPoints[co] = v
	out:
	}
	return lowPoints
}

func calculateRiskLevels(lowPoints LowPoints) int {
	risk := 0
	for _, v := range lowPoints {
		risk += v + 1
	}
	return risk
}

func (hm HeightMap) coIsPartOfBasin(b Basin, co utils.Co) bool {
	if _, ok := b[co]; !ok {
		if v, ok := hm[co]; ok && v != 9 {
			return true
		}
	}
	return false
}

func (hm HeightMap) calculateBasin(co utils.Co) int {
	b := Basin{
		co: {},
	}
	for {
		newCos := []utils.Co{}
		for co := range b {
			for _, newCo := range utils.AdjacentCos(co, false) {
				if hm.coIsPartOfBasin(b, newCo) {
					newCos = append(newCos, newCo)
				}
			}
		}
		if len(newCos) == 0 {
			return len(b)
		}
		for _, co := range newCos {
			b[co] = struct{}{}
		}
	}
}

func (bs Basins) multiplyLargestBasinSizes() int {
	sort.Ints(bs)
	return bs[len(bs)-1] * bs[len(bs)-2] * bs[len(bs)-3]
}

func findSolutions(input []string) (int, int) {
	hm := parseInput(input)
	lp := hm.findLowPoints()
	basins := Basins{}
	for co := range lp {
		basins = append(basins, hm.calculateBasin(co))
	}
	return calculateRiskLevels(lp), basins.multiplyLargestBasinSizes()
}

func main() {
	input := utils.ReadFile()
	part1, part2 := findSolutions(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
