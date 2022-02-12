package main

import (
	"Advent-of-Code/2021/Day04/game"
	"Advent-of-Code/file"
	"fmt"
)

func main() {
	input := file.Read()
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
