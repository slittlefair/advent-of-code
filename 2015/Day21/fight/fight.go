package fight

import (
	cmb "Advent-of-Code/2015/Day21/combatant"
	"Advent-of-Code/2015/Day21/shop"
	"fmt"
	"math/rand"
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

func MartialAttack(attacker, defender *cmb.Combatant) {
	damage := attacker.Damage - defender.Armour
	if damage < 1 {
		damage = 1
	}
	defender.HitPoints -= damage
}

func SpellAttack(attacker, defender *cmb.Combatant, spell *cmb.Spell) {
	fmt.Println("casting", spell.Name)
	attacker.Mana -= spell.Mana
	attacker.ManaSpent += spell.Mana
	attacker.HitPoints += spell.HitPoints
	defender.HitPoints -= spell.Damage
	if spell.Effect.Effect != nil {
		spell.Effect.Active = true
		spell.Effect.DurationRemaining = spell.Effect.Duration
	}
}

func ApplyEffects(attacker, defender *cmb.Combatant) {
	for _, sp := range attacker.Spells {
		if sp.Effect.Active {
			sp.Effect.ApplyEffect(attacker, defender)
		}
	}
}

func Fight(player, boss *cmb.Combatant) bool {
	for {
		MartialAttack(player, boss)
		if boss.IsDead() {
			return true
		}
		MartialAttack(boss, player)
		if player.IsDead() {
			return false
		}
	}
}

func SpellRound(player, boss *cmb.Combatant) bool {
	ApplyEffects(player, boss)
	if boss.IsDead() {
		player.CompareManaSpent()
		fmt.Println("fight over, boss dead")
		return false
	}
	validSpells := player.ValidSpells()
	spell := validSpells[rand.Intn(len(validSpells))]
	if player.Mana < spell.Mana {
		fmt.Println("fight over, out of mana")
		return false
	}
	SpellAttack(player, boss, spell)
	if boss.IsDead() {
		player.CompareManaSpent()
		fmt.Println("fight over, boss dead")
		return false
	}
	MartialAttack(boss, player)
	if player.IsDead() {
		fmt.Println("fight over, player dead")
		return false
	}
	return true
}

func SpellFight(player, boss *cmb.Combatant) {
	fighting := true
	for fighting {
		fighting = SpellRound(player, boss)
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
