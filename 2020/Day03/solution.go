package main

import (
	helpers "Advent-of-Code"
	"fmt"
)

func traverseSlopes(right int, down int) int {
	entries := helpers.ReadFile()
	col := 0
	treesEncountered := 0
	for i := 0; i < len(entries); i += down {
		row := entries[i]
		if string(row[col%len(row)]) == "#" {
			treesEncountered++
		}
		col += right
	}
	return treesEncountered
}

func part1() int {
	return traverseSlopes(3, 1)
}

func part2() int {
	return traverseSlopes(1, 1) * traverseSlopes(3, 1) * traverseSlopes(5, 1) * traverseSlopes(7, 1) * traverseSlopes(1, 2)
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
