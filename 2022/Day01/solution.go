package main

import (
	"Advent-of-Code/file"
	"fmt"
	"sort"
	"strconv"
)

type Elf struct {
	id            int
	snacks        []int
	totalCalories int
}

func (e *Elf) calculateTotalCalories() {
	tc := 0
	for _, snk := range e.snacks {
		tc += snk
	}
	e.totalCalories = tc
}

func parseInput(input []string) ([]*Elf, error) {
	elves := []*Elf{}
	elf := &Elf{
		id:     len(elves),
		snacks: []int{},
	}
	for _, line := range input {
		if line == "" {
			elf.calculateTotalCalories()
			elves = append(elves, elf)
			elf = &Elf{
				id:     len(elves),
				snacks: []int{},
			}
			continue
		}
		cal, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		elf.snacks = append(elf.snacks, cal)
	}
	elf.calculateTotalCalories()
	elves = append(elves, elf)
	return elves, nil
}

func sortElvesByDescendingCalories(elves []*Elf) []*Elf {
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].totalCalories > elves[j].totalCalories
	})
	return elves
}

func findLargestCalories(elves []*Elf, n int) int {
	calories := 0
	for i := 0; i < n; i++ {
		calories += elves[i].totalCalories
	}
	return calories
}

func main() {
	input := file.Read()
	elves, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	elves = sortElvesByDescendingCalories(elves)
	fmt.Println("Part 1:", findLargestCalories(elves, 1))
	fmt.Println("Part 2:", findLargestCalories(elves, 3))
}
