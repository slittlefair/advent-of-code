package mage

import (
	"Advent-of-Code/2015/Day21/martial"
)

type Mage struct {
	HP        int
	Armour    int
	Mana      int
	Spells    Spells
	ManaSpent int
}

func (m *Mage) SpellIsValid(spell *Spell, effects map[string]Effect) bool {
	if m.Mana < spell.Mana {
		return false
	}
	return !effects[spell.Effect].Active
}

type Effect struct {
	Active            bool
	Duration          int
	DurationRemaining int
	Effect            func(mage *Mage, boss *martial.Martial)
}

var Effects = map[string]Effect{
	"None": {},
	"Shield": {
		Duration: 6,
		Effect:   Shield,
	},
	"Poison": {
		Duration: 6,
		Effect:   Poison,
	},
	"Recharge": {
		Duration: 5,
		Effect:   Recharge,
	},
}

func ApplyEffect(mage *Mage, boss *martial.Martial, e Effect) Effect {
	e.Effect(mage, boss)
	e.DurationRemaining--
	if e.DurationRemaining == 0 {
		e.Active = false
	}
	return e
}

func Shield(mage *Mage, _ *martial.Martial) {
	mage.Armour = 7
}

func Poison(_ *Mage, boss *martial.Martial) {
	boss.HP -= 3
}

func Recharge(mage *Mage, _ *martial.Martial) {
	mage.Mana += 101
}

type Spell struct {
	Name   string
	Mana   int
	Damage int
	HP     int
	Effect string
}

type Spells map[string]*Spell

var SpellList = Spells{
	"Magic Missile": {
		Mana:   53,
		Damage: 4,
		Effect: "None",
	},
	"Drain": {
		Mana:   73,
		Damage: 2,
		HP:     2,
		Effect: "None",
	},
	"Shield": {
		Mana:   113,
		Effect: "Shield",
	},
	"Poison": {
		Mana:   173,
		Effect: "Poison",
	},
	"Recharge": {
		Mana:   229,
		Effect: "Recharge",
	},
}
