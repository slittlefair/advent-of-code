package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
)

var initialState = ".#####.##.#.##...#.#.###..#.#..#..#.....#..####.#.##.#######..#...##.#..#.#######...#.#.#..##..#.#.#"
var rules = make(map[string]string)
var changes = make(map[int]string)
var pots []string
var buffer = 0

func countPlants() (genPlants int) {
	genPlants = 0
	for idx, p := range pots {
		if p == "#" {
			genPlants += idx - buffer
		}
	}
	return
}

func addPotBuffers() {
	if pots[0] == "#" || pots[1] == "#" {
		pots = append([]string{".", ".", "."}, pots...)
		buffer += 3
	}
	if pots[len(pots)-3] == "#" || pots[len(pots)-2] == "#" || pots[len(pots)-1] == "#" {
		pots = append(pots, []string{".", ".", "."}...)
	}
}

func main() {
	initialRules := utils.ReadFile()
	for _, v := range initialRules {
		rules[v[:5]] = v[9:]
	}
	for _, v := range initialState {
		pots = append(pots, string(v))
	}

	var gen = 0
	var total = 100
	genPlants := countPlants()
	diff := 0
	for ; gen <= total; gen++ {
		diff = countPlants() - genPlants
		genPlants = countPlants()
		fmt.Println(gen, genPlants, diff)
		addPotBuffers()
		for i := 2; i < len(pots)-2; i++ {
			matchingSequence := ""
			for j := i - 2; j <= i+2; j++ {
				matchingSequence += pots[j]
			}
			changes[i] = rules[matchingSequence]
		}
		for k, v := range changes {
			pots[k] = v
		}
	}
	fmt.Println(((50000000000 - 100) * diff) + genPlants)
}
