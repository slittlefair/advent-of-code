package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"sort"
)

type HeightMap map[helpers.Co]int

type LowPoints map[helpers.Co]int

type Basin map[helpers.Co]struct{}
type Basins []int

func parseInput(input []string) HeightMap {
	hm := HeightMap{}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			hm[helpers.Co{X: x, Y: y}] = int(input[y][x] - '0')
		}
	}
	return hm
}

func (hm HeightMap) findLowPoints() LowPoints {
	lowPoints := LowPoints{}
	for co, v := range hm {
		if val, ok := hm[helpers.Co{X: co.X - 1, Y: co.Y}]; ok && val <= v {
			continue
		}
		if val, ok := hm[helpers.Co{X: co.X + 1, Y: co.Y}]; ok && val <= v {
			continue
		}
		if val, ok := hm[helpers.Co{X: co.X, Y: co.Y - 1}]; ok && val <= v {
			continue
		}
		if val, ok := hm[helpers.Co{X: co.X, Y: co.Y + 1}]; ok && val <= v {
			continue
		}
		lowPoints[co] = v
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

func (hm HeightMap) coIsPartOfBasin(b Basin, co helpers.Co) bool {
	if _, ok := b[co]; !ok {
		if v, ok := hm[co]; ok && v != 9 {
			return true
		}
	}
	return false
}

func (hm HeightMap) calculateBasin(co helpers.Co) int {
	b := Basin{
		co: {},
	}
	for {
		newCos := []helpers.Co{}
		for co := range b {
			newCo := helpers.Co{X: co.X - 1, Y: co.Y}
			if hm.coIsPartOfBasin(b, newCo) {
				newCos = append(newCos, newCo)
			}
			newCo = helpers.Co{X: co.X + 1, Y: co.Y}
			if hm.coIsPartOfBasin(b, newCo) {
				newCos = append(newCos, newCo)
			}
			newCo = helpers.Co{X: co.X, Y: co.Y - 1}
			if hm.coIsPartOfBasin(b, newCo) {
				newCos = append(newCos, newCo)
			}
			newCo = helpers.Co{X: co.X, Y: co.Y + 1}
			if hm.coIsPartOfBasin(b, newCo) {
				newCos = append(newCos, newCo)
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
	input := helpers.ReadFile()
	part1, part2 := findSolutions(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
