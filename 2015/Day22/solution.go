package main

import (
	helpers "Advent-of-Code"
	"Advent-of-Code/2015/Day22/spellFight"
	"fmt"
)

func main() {
	input := helpers.ReadFile()
	lowestManaSpent, lowestManaSpentHardMode, err := spellFight.RunSpellFights(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", lowestManaSpent)
	fmt.Println("Part 2:", lowestManaSpentHardMode)
}
