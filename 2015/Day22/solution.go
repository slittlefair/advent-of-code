package main

import (
	helpers "Advent-of-Code"
	"Advent-of-Code/2015/Day21/combatant"
	fight "Advent-of-Code/2015/Day21/fight"
	"fmt"
)

func runFights(input []string) (int, int, error) {
	f := &fight.Fighters{
		LowestManaSpent: helpers.Infinty,
	}
	err := f.ParseBoss(input, false)
	if err != nil {
		return -1, -1, err
	}
	bossHP := f.Boss.HitPoints
	f.Player = &combatant.Combatant{
		LowestManaSpent: helpers.Infinty,
		Spells:          combatant.PopulateSpells(),
	}
	lowestMana, err := fight.SpellFight(*f.Player, *f.Boss, bossHP, false)
	if err != nil {
		return -1, -1, err
	}
	lowestManaHardMode, err := fight.SpellFight(*f.Player, *f.Boss, bossHP, true)
	if err != nil {
		return -1, -1, err
	}
	return lowestMana, lowestManaHardMode, nil
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
