package main

import "fmt"

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

func main() {
	f := &Fighters{
		Player: &Combatant{
			HitPoints: 8,
			Damage:    5,
			Armour:    5,
		},
		Boss: &Combatant{
			HitPoints: 12,
			Damage:    7,
			Armour:    2,
		},
	}
	fmt.Println(f.Fight())
}
