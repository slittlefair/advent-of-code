package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Combatant struct {
	HitPoints int
	Damage    int
	Armour    int
	Cost      int
}

type Fighters struct {
	Player            *Combatant
	Boss              *Combatant
	SuccessfulCosts   []int
	UnsuccessfulCosts []int
}

type Equipment struct {
	Name   string
	Cost   int
	Damage int
	Armour int
}

type Shop struct {
	Weapons []Equipment
	Armour  []Equipment
	Rings   []Equipment
}

func (f *Fighters) Attack(combatant *Combatant, opponent *Combatant) {
	damage := combatant.Damage - opponent.Armour
	if damage < 1 {
		damage = 1
	}
	opponent.HitPoints -= damage
}

func (c *Combatant) IsDead() bool {
	return c.HitPoints <= 0
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

	boss := &Combatant{}

	hp, err := strconv.Atoi(strings.Split(input[0], "Hit Points: ")[1])
	if err != nil {
		return err
	}
	boss.HitPoints = hp

	damage, err := strconv.Atoi(strings.Split(input[1], "Damage: ")[1])
	if err != nil {
		return err
	}
	boss.Damage = damage

	armour, err := strconv.Atoi(strings.Split(input[2], "Armor: ")[1])
	if err != nil {
		return err
	}
	boss.Armour = armour

	f.Boss = &Combatant{
		HitPoints: hp,
		Damage:    damage,
		Armour:    armour,
	}

	return nil
}

func PopulateShop() *Shop {
	return &Shop{
		Armour: []Equipment{
			{
				Name: "Optional Armour",
			},
			{
				Name:   "Leather",
				Armour: 1,
				Cost:   13,
			},
			{
				Name:   "Chainmail",
				Armour: 2,
				Cost:   31,
			},
			{
				Name:   "Splintmail",
				Armour: 3,
				Cost:   53,
			},
			{
				Name:   "Bandedmail",
				Armour: 4,
				Cost:   75,
			},
			{
				Name:   "Platemail",
				Armour: 5,
				Cost:   102,
			},
		},
		Rings: []Equipment{
			{
				Name: "Optional Ring 1",
			},
			{
				Name: "Optional Ring 2",
			},
			{
				Name:   "Damage +1",
				Damage: 1,
				Cost:   25,
			},
			{
				Name:   "Damage +2",
				Damage: 2,
				Cost:   50,
			},
			{
				Name:   "Damage +3",
				Damage: 3,
				Cost:   100,
			},
			{
				Name:   "Defence +1",
				Armour: 1,
				Cost:   20,
			},
			{
				Name:   "Defence +2",
				Armour: 2,
				Cost:   40,
			},
			{
				Name:   "Defence +3",
				Armour: 3,
				Cost:   80,
			},
		},
		Weapons: []Equipment{
			{
				Name:   "Dagger",
				Damage: 4,
				Cost:   8,
			},
			{
				Name:   "Shortsword",
				Damage: 5,
				Cost:   10,
			},
			{
				Name:   "Warhammer",
				Damage: 6,
				Cost:   25,
			},
			{
				Name:   "Longsword",
				Damage: 7,
				Cost:   40,
			},
			{
				Name:   "Greataxe",
				Damage: 8,
				Cost:   74,
			},
		},
	}
}

func (f *Fighters) AssignEquipment(weapon, armour, ring1, ring2 Equipment) {
	f.Player = &Combatant{
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
	sort.Ints(f.UnsuccessfulCosts)
	return f.UnsuccessfulCosts[len(f.UnsuccessfulCosts)-1], nil
}

func runFights(input []string) *Fighters {
	f := &Fighters{
		SuccessfulCosts: []int{},
	}
	f.ParseBoss(input)
	bossHP := f.Boss.HitPoints
	s := PopulateShop()
	for _, armour := range s.Armour {
		for _, ring1 := range s.Rings {
			for _, ring2 := range s.Rings {
				if ring1 == ring2 {
					continue
				}
				for _, weapon := range s.Weapons {
					f.AssignEquipment(weapon, armour, ring1, ring2)
					// Heal boss up before fight
					f.Boss.HitPoints = bossHP
					if f.Fight() {
						f.SuccessfulCosts = append(f.SuccessfulCosts, f.Player.Cost)
					} else {
						f.UnsuccessfulCosts = append(f.UnsuccessfulCosts, f.Player.Cost)
					}
				}
			}
		}
	}
	return f
}

func main() {
	input := helpers.ReadFile()
	fighters := runFights(input)
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
