package main

import (
	helpers "Advent-of-Code"
	"Advent-of-Code/2015/Day21/combatant"
	fight "Advent-of-Code/2015/Day21/fight"
	"fmt"
	"math/rand"
	"time"
)

func runFights(input []string) (int, error) {
	f := &fight.Fighters{}
	err := f.ParseBoss(input, false)
	if err != nil {
		return -1, err
	}
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	f.Player = &combatant.Combatant{
		LowestManaSpent: helpers.Infinty,
		Spells:          combatant.PopulateSpells(),
		HitPoints:       50,
		Mana:            500,
	}
	for {
		fight.Fight(f.Player, f.Boss)
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
	fmt.Println("Part 1:", lowestManaSpent)
}
