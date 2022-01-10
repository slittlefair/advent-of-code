package main

import (
	"Advent-of-Code/2015/Day22/spellFight"
	utils "Advent-of-Code/utils"
	"fmt"
)

func main() {
	input := utils.ReadFile()
	lowestManaSpent, lowestManaSpentHardMode, err := spellFight.RunSpellFights(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", lowestManaSpent)
	fmt.Println("Part 2:", lowestManaSpentHardMode)
}
