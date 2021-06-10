package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"strconv"
	"strings"
)

type Combatant struct {
	HitPoints int
	Damage    int
	Armour    int
}

type Fighters struct {
	Player *Combatant
	Boss   *Combatant
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

func main() {
	input := helpers.ReadFile()
	f := &Fighters{
		Player: &Combatant{
			HitPoints: 8,
			Damage:    5,
			Armour:    5,
		},
	}
	f.ParseBoss(input)
	fmt.Printf("%+v\n", f.Boss)
	fmt.Println(f.Fight())
}
