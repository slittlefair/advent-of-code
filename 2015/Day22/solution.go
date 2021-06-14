package main

import (
	helpers "Advent-of-Code"
	"Advent-of-Code/2015/Day21/martial"
	"Advent-of-Code/2015/Day22/mage"
	"Advent-of-Code/2015/Day22/spellFight"
	"fmt"
)

func runFights(input []string) (int, int, error) {
	boss, err := martial.ParseBoss(input, false)
	if err != nil {
		return -1, -1, err
	}
	bossHP := boss.HP
	player := &mage.Mage{
		Spells: mage.PopulateSpells(),
	}
	lowestManaSpent := spellFight.SpellFight(*player, *boss, bossHP, false)
	lowestManaHardModeSpent := spellFight.SpellFight(*player, *boss, bossHP, true)
	return lowestManaSpent, lowestManaHardModeSpent, nil
}

func main() {
	input := helpers.ReadFile()
	lowestManaSpent, lowestManaSpentHardMode, err := runFights(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", lowestManaSpent)
	fmt.Println("Part 2:", lowestManaSpentHardMode)
}
