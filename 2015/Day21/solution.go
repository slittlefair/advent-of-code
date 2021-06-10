package main

import (
	helpers "Advent-of-Code"
	cmb "Advent-of-Code/2015/Day21/combatant"
	shop "Advent-of-Code/2015/Day21/shop"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Fighters struct {
	Player            *cmb.Combatant
	Boss              *cmb.Combatant
	SuccessfulCosts   []int
	UnsuccessfulCosts []int
}

func (f *Fighters) Attack(attacker *cmb.Combatant, defender *cmb.Combatant) {
	damage := attacker.Damage - defender.Armour
	if damage < 1 {
		damage = 1
	}
	defender.HitPoints -= damage
}

func (f *Fighters) Fight() bool {
	for {
		f.Attack(f.Player, f.Boss)
		if f.Boss.IsDead() {
			return true
		}
		f.Attack(f.Boss, f.Player)
		if f.Player.IsDead() {
			return false
		}
	}
}

func (f *Fighters) ParseBoss(input []string) error {
	if length := len(input); length != 3 {
		return fmt.Errorf("something went wrong, expected 3 lines of input, got %d", length)
	}

	boss := &cmb.Combatant{}

	val := strings.Split(input[0], "Hit Points: ")
	if len(val) != 2 {
		return fmt.Errorf("something went wrong, could not correctly split line %s", input[0])
	}
	hp, err := strconv.Atoi(val[1])
	if err != nil {
		return err
	}
	boss.HitPoints = hp

	val = strings.Split(input[1], "Damage: ")
	if len(val) != 2 {
		return fmt.Errorf("something went wrong, could not correctly split line %s", input[0])
	}
	damage, err := strconv.Atoi(val[1])
	if err != nil {
		return err
	}
	boss.Damage = damage

	val = strings.Split(input[2], "Armor: ")
	if len(val) != 2 {
		return fmt.Errorf("something went wrong, could not correctly split line %s", input[0])
	}
	armour, err := strconv.Atoi(val[1])
	if err != nil {
		return err
	}
	boss.Armour = armour

	f.Boss = &cmb.Combatant{
		HitPoints: hp,
		Damage:    damage,
		Armour:    armour,
	}

	return nil
}

func (f *Fighters) InitiatePlayerForFight(weapon, armour, ring1, ring2 shop.Equipment) {
	f.Player = &cmb.Combatant{
		Armour:    weapon.Armour + armour.Armour + ring1.Armour + ring2.Armour,
		Damage:    weapon.Damage + armour.Damage + ring1.Damage + ring2.Damage,
		Cost:      weapon.Cost + armour.Cost + ring1.Cost + ring2.Cost,
		HitPoints: 100,
	}
}

func (f *Fighters) CheapestVictory() (int, error) {
	if len(f.SuccessfulCosts) == 0 {
		return -1, fmt.Errorf("something went wrong - no player victories")
	}
	sort.Ints(f.SuccessfulCosts)
	return f.SuccessfulCosts[0], nil
}

func (f *Fighters) DearestLoss() (int, error) {
	if len(f.UnsuccessfulCosts) == 0 {
		return -1, fmt.Errorf("something went wrong - no player losses")
	}
	sort.Sort(sort.Reverse(sort.IntSlice(f.UnsuccessfulCosts)))
	return f.UnsuccessfulCosts[0], nil
}

func runFights(input []string) (*Fighters, error) {
	f := &Fighters{
		SuccessfulCosts: []int{},
	}
	err := f.ParseBoss(input)
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
					if f.Fight() {
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
