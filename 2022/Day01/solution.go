package main

import (
	"Advent-of-Code/file"
	"fmt"
	"sort"
	"strconv"
)

func parseInput(input []string) ([]int, error) {
	elves := []int{}
	elf := 0
	for _, line := range input {
		if line == "" {
			elves = append(elves, elf)
			elf = 0
			continue
		}
		cal, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		elf += cal
	}
	elves = append(elves, elf)
	return elves, nil
}

func findLargestCalories(elves []int, n int) (int, error) {
	if n > len(elves) {
		return -1, fmt.Errorf("findLargestCalories: asked for more items (%d) than are in provided slice %v", n, elves)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	calories := 0
	for i := 0; i < n; i++ {
		calories += elves[i]
	}
	return calories, nil
}

func main() {
	input := file.Read()
	elves, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	part1, err := findLargestCalories(elves, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	part2, err := findLargestCalories(elves, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
