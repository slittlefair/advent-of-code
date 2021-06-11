package main

import (
	helpers "Advent-of-Code"
	"Advent-of-Code/2015/Day21/combatant"
	fight "Advent-of-Code/2015/Day21/fight"
	"fmt"
)

func runFights(input []string) (int, error) {
	f := &fight.Fighters{}
	err := f.ParseBoss(input, false)
	if err != nil {
		return -1, err
	}
	f.Player = &combatant.Combatant{
		HitPoints:       50,
		Mana:            500,
		LowestManaSpent: helpers.Infinty,
		Spells:          combatant.PopulateSpells(),
	}
	for _, sp := range f.Player.Spells {
		f.SpellRound(&sp)
	}
	return f.Player.LowestManaSpent, nil
}

func main() {
	input := helpers.ReadFile()
	lowestManaSpent, err := runFights(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(helpers.Infinty - 9223372036854775807)
	fmt.Println("Part 1:", lowestManaSpent)
}
