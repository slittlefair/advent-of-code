package main

import (
	"Advent-of-Code"
	"Advent-of-Code-2018/Day15/Answer"
	"fmt"
)

func main() {
	cave := helpers.ReadFile()

	fmt.Printf("Combat: %d\n", day15.Combat(cave))
	fmt.Printf("CheatingElves: %d\n", day15.CheatingElves(cave))
}
