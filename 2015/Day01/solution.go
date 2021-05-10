package main

import (
	"Advent-of-Code"
	"fmt"
)

func main() {
	instructions := helpers.ReadFile()[0]
	currentFloor := 0
	part2Sol := 0
	for i, r := range instructions {
		char := string(r)
		if char == "(" {
			currentFloor++
		} else {
			currentFloor--
		}
		if currentFloor == -1 && part2Sol == 0 {
			part2Sol = i + 1
		}
	}
	fmt.Println("Part 1:", currentFloor)
	fmt.Println("Part 2:", part2Sol)
}
