package main

import (
	helpers "Advent-of-Code"
	"Advent-of-Code/2021/Day04/game"
	"fmt"
)

func main() {
	input := helpers.ReadFile()
	g, err := game.ParseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	part1, part2, err := g.PlayGame()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
