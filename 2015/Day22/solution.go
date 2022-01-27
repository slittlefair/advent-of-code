package main

import (
	spellfight "Advent-of-Code/2015/Day22/spells"
	utils "Advent-of-Code/utils"
	"fmt"
)

func main() {
	input := utils.ReadFile()
	lowestManaSpent, lowestManaSpentHardMode, err := spellfight.RunSpellFights(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", lowestManaSpent)
	fmt.Println("Part 2:", lowestManaSpentHardMode)
}
