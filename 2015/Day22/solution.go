package main

import (
	sp "Advent-of-Code/2015/Day22/spellfight"
	"Advent-of-Code/file"
	"fmt"
)

func main() {
	input := file.Read()
	lowestManaSpent, lowestManaSpentHardMode, err := sp.RunSpellFights(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", lowestManaSpent)
	fmt.Println("Part 2:", lowestManaSpentHardMode)
}
