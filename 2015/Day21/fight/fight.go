package fight

import (
	cmb "Advent-of-Code/2015/Day21/combatant"
	"Advent-of-Code/2015/Day21/shop"
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

func (f *Fighters) MartialAttack(attacker, defender *cmb.Combatant) {
	damage := attacker.Damage - defender.Armour
	if damage < 1 {
		damage = 1
	}
	defender.HitPoints -= damage
}

func (f *Fighters) SpellAttack(spell *cmb.Spell) int {
	fmt.Println("casting", spell.Name)
	f.Player.Mana -= spell.Mana
	f.Player.ManaSpent += spell.Mana
	f.Player.HitPoints += spell.HitPoints
	f.Boss.HitPoints -= spell.Damage
	if spell.Effect.Effect != nil {
		spell.Effect.Active = true
		spell.Effect.DurationRemaining = spell.Effect.Duration
	}
	return spell.Mana
}

func ApplyEffects(attacker, defender *cmb.Combatant) {
	for _, sp := range attacker.Spells {
		if sp.Effect.Active {
			sp.Effect.ApplyEffect(attacker, defender)
		}
	}
}

func (f *Fighters) Fight() bool {
	for {
		f.MartialAttack(f.Player, f.Boss)
		if f.Boss.IsDead() {
			return true
		}
		f.MartialAttack(f.Boss, f.Player)
		if f.Player.IsDead() {
			return false
		}
	}
}

func (f *Fighters) SpellRound(spell *cmb.Spell) {
	ApplyEffects(f.Player, f.Boss)
	if f.Boss.IsDead() {
		f.CompareManaSpent()
		return
	}
	if f.Player.Mana < spell.Mana {
		return
	}
	f.SpellAttack(spell)
	if f.Boss.IsDead() {
		f.CompareManaSpent()
		return
	}
	f.MartialAttack(f.Boss, f.Player)
	if f.Player.IsDead() {
		return
	}
	for _, sp := range f.Player.Spells {
		if !sp.Effect.Active {
			f.SpellRound(&sp)
		}
	}
}

func (f *Fighters) CompareManaSpent() {
	if f.Player.ManaSpent < f.Player.LowestManaSpent {
		f.Player.LowestManaSpent = f.Player.ManaSpent
	}
}

func (f *Fighters) ParseBoss(input []string, hasArmour bool) error {
	if length := len(input); hasArmour && length != 3 {
		return fmt.Errorf("something went wrong, expected 3 lines of input, got %d, %v", length, input)
	} else if length != 2 {
		return fmt.Errorf("something went wrong, expected 2 lines of input, got %d, %v", length, input)
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

	if hasArmour {
		val = strings.Split(input[2], "Armor: ")
		if len(val) != 2 {
			return fmt.Errorf("something went wrong, could not correctly split line %s", input[0])
		}
		armour, err := strconv.Atoi(val[1])
		if err != nil {
			return err
		}
		boss.Armour = armour
	}

	f.Boss = boss

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
