package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"sort"
)

type Joltages map[int]int

func parseInput(input []int) []int {
	sort.Ints(input)
	return append(input, input[len(input)-1]+3)
}

func (j Joltages) part1(adapters []int) int {
	jolt := 0
	for _, adapter := range adapters {
		j[adapter-jolt]++
		jolt = adapter
	}
	return j[1] * j[3]
}

func calculatePerms(adapters []int, val, i int, cache []int) int {
	accPerms := 0
	if i > 0 && i < len(adapters) && cache[i] > 0 {
		return cache[i]
	}
	for j := i + 1; j <= j+3 && j < len(adapters); j++ {
		nextVal := adapters[j]
		if nextVal-val <= 3 {
			permValue := calculatePerms(adapters, nextVal, j, cache)
			cache[j] = permValue
			accPerms += permValue
		}
	}
	return accPerms
}

func part2(adapters []int) int {
	cache := make([]int, len(adapters))
	cache[len(adapters)-1] = 1
	return calculatePerms(adapters, 0, -1, cache)
}

func main() {
	input := helpers.ReadFileAsInts()
	joltages := Joltages{
		1: 0,
		2: 0,
		3: 0,
	}
	adapters := parseInput(input)

	fmt.Println("Part 1:", joltages.part1(adapters))
	fmt.Println("Part 2:", part2(adapters))
}
