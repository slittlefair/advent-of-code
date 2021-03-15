package main

import (
	"Advent-of-Code"
	"Advent-of-Code-2018/Day23/Answer"
	"fmt"
)

func main() {
	input := helpers.ReadFile()

	nanobots := day23.NewBots(input)

	fmt.Printf("StrongestReachable: %d\n", day23.StrongestReachable(nanobots))
	fmt.Printf("ClosestSuccess: %d\n", day23.ClosestSuccess(nanobots))
}
