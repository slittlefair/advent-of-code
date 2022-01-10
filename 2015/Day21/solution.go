package main

import (
	"Advent-of-Code/2015/Day21/fight"
	"Advent-of-Code/2015/Day21/martial"
	"Advent-of-Code/2015/Day21/shop"
	utils "Advent-of-Code/utils"
	"fmt"
)

func runFights(input []string) (int, int, error) {
	boss, err := martial.ParseBoss(input, true)
	if err != nil {
		return -1, -1, err
	}
	bossHP := boss.HP
	s := shop.PopulateShop()
	cheapestVictory := utils.Infinty
	dearestLoss := 0
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
					player, fightCost := fight.InitiatePlayerForFight(weapon, armour, ring1, ring2)
					if fight.Fight(player, boss) {
						if fightCost < cheapestVictory {
							cheapestVictory = fightCost
						}
					} else {
						if fightCost > dearestLoss {
							dearestLoss = fightCost
						}
					}
					// Heal boss up after fight
					boss.HP = bossHP
				}
			}
		}
	}
	return cheapestVictory, dearestLoss, nil
}

func main() {
	input := utils.ReadFile()
	cheapestVictory, dearestLoss, err := runFights(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", cheapestVictory)
	fmt.Println("Part 2:", dearestLoss)
}
