package main

import (
	helpers "Advent-of-Code"
	"Advent-of-Code/2015/Day21/fight"
	shop "Advent-of-Code/2015/Day21/shop"
	"fmt"
)

func runFights(input []string) (*fight.Fighters, error) {
	f := &fight.Fighters{
		SuccessfulCosts: []int{},
	}
	err := f.ParseBoss(input, true)
	if err != nil {
		return nil, err
	}
	bossHP := f.Boss.HitPoints
	s := shop.PopulateShop()
	// Loop through all combinations of armour, rings and weapons and work out their cost as well
	// as who wins the fight
	for _, armour := range s.Armour {
		for _, ring1 := range s.Rings {
			for _, ring2 := range s.Rings {
				// Can't have two of the same ring
				if ring1 == ring2 {
					continue
				}
				for _, weapon := range s.Weapons {
					f.InitiatePlayerForFight(weapon, armour, ring1, ring2)
					if fight.Fight(f.Player, f.Boss) {
						f.SuccessfulCosts = append(f.SuccessfulCosts, f.Player.Cost)
					} else {
						f.UnsuccessfulCosts = append(f.UnsuccessfulCosts, f.Player.Cost)
					}
					// Heal boss up after fight
					f.Boss.HitPoints = bossHP
				}
			}
		}
	}
	return f, nil
}

func main() {
	input := helpers.ReadFile()
	fighters, err := runFights(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	cheapestVictory, err := fighters.CheapestVictory()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", cheapestVictory)

	dearestLoss, err := fighters.DearestLoss()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 2:", dearestLoss)
}
