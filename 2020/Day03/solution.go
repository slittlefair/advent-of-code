package main

import (
	"Advent-of-Code/file"
	"fmt"
)

type TreeMap []string

func (tm TreeMap) traverseSlopes(right int, down int) int {
	col := 0
	treesEncountered := 0
	for i := 0; i < len(tm); i += down {
		row := tm[i]
		if string(row[col%len(row)]) == "#" {
			treesEncountered++
		}
		col += right
	}
	return treesEncountered
}

func (tm TreeMap) part1() int {
	return tm.traverseSlopes(3, 1)
}

func (tm TreeMap) part2() int {
	return tm.traverseSlopes(1, 1) * tm.traverseSlopes(3, 1) * tm.traverseSlopes(5, 1) * tm.traverseSlopes(7, 1) * tm.traverseSlopes(1, 2)
}

func main() {
	var tm TreeMap = file.Read()
	fmt.Println("Part 1:", tm.part1())
	fmt.Println("Part 2:", tm.part2())
}
