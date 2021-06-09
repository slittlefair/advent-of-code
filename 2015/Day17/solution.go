package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"sort"
	"strconv"
)

func parseInput(input []string) ([]int, error) {
	inputInt := make([]int, len(input))
	for i, v := range input {
		conv, err := strconv.Atoi(v)
		if err != nil {
			return nil, nil
		}
		inputInt[i] = conv
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inputInt)))
	return inputInt, nil
}

func findContainers(remainingContainers []int, count int, totalCapacity int, wantedTotal int) ([]int, int, int, int) {
	if totalCapacity == wantedTotal {
		count++
	}
	for i := 0; i < len(remainingContainers); i++ {
		_, count, _, wantedTotal = findContainers(remainingContainers[i+1:], count, totalCapacity+remainingContainers[i], wantedTotal)
	}
	return remainingContainers, count, totalCapacity, wantedTotal
}

func getPermutations(containers []int) int {
	_, perms, _, _ := findContainers(containers, 0, 0, 150)
	return perms
}

func main() {
	input := helpers.ReadFile()
	containers, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	permutations := getPermutations(containers)
	fmt.Println("Part 1:", permutations)
	// fmt.Println(remainingContainers, count, tc, wt)
}
