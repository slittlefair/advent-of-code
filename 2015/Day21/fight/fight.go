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
	LowestManaSpent   int
}

func MartialAttack(attacker, defender *cmb.Combatant) {
	damage := attacker.Damage - defender.Armour
	if damage < 1 {
		damage = 1
	}
	defender.HitPoints -= damage
}

func SpellAttack(attacker, defender *cmb.Combatant, spell *cmb.Spell, effects map[string]cmb.Effect) map[string]cmb.Effect {
	attacker.Mana -= spell.Mana
	attacker.HitPoints += spell.HitPoints
	defender.HitPoints -= spell.Damage
	if spell.Effect != "None" {
		eff := effects[spell.Effect]
		eff.Active = true
		eff.DurationRemaining = eff.Duration
		effects[spell.Effect] = eff
	}
	return effects
}

func ApplyEffects(attacker, defender *cmb.Combatant, effects map[string]cmb.Effect) map[string]cmb.Effect {
	effs := map[string]cmb.Effect{}
	for k, eff := range effects {
		if eff.Active {
			effs[k] = cmb.ApplyEffect(attacker, defender, eff)
		} else {
			effs[k] = eff
		}
	}
	return effs
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

type ManaSpends struct {
	Winning []int
}

func (ms *ManaSpends) SpellRound(player, boss cmb.Combatant, spell *cmb.Spell, currentMana int, spells []*cmb.Spell, level int, effects map[string]cmb.Effect, hardMode bool) {
	// Set playar armour to 0
	player.Armour = 0

	if hardMode {
		player.HitPoints--
		if player.IsDead() {
			return
		}
	}

	effects = ApplyEffects(&player, &boss, effects)

	// If the spell effects have caused the boss to die, end the fight and compare mana scores
	if boss.IsDead() {
		ms.Winning = append(ms.Winning, currentMana)
		return
	}

	if !player.SpellIsValid(spell, effects) {
		return
	}

	// Cast the spell and accumulate spent mana
	effects = SpellAttack(&player, &boss, spell, effects)
	spells = append(spells, spell)

	// fmt.Println(*player.Spells[1])
	currentMana += spell.Mana

	// If the spell causes the boss to die, end the fight and compare mana scores
	if boss.IsDead() {
		ms.Winning = append(ms.Winning, currentMana)
		return
	}

	// At the start of the boss' turn, set player armour to 0 and apply spell effects
	player.Armour = 0
	effects = ApplyEffects(&player, &boss, effects)

	// If the spell effects have caused the boss to die, end the fight and compare mana scores
	if boss.IsDead() {
		ms.Winning = append(ms.Winning, currentMana)
		return
	}

	// Boss attacks the player with a marital attack
	MartialAttack(&boss, &player)

	// If the boss attack causes the player to die, end the fight
	if player.IsDead() {
		return
	}

	for _, vs := range player.Spells {
		ms.SpellRound(player, boss, vs, currentMana, spells, level+1, effects, hardMode)
	}
}

func SpellFight(player, boss cmb.Combatant, bossHP int, hardMode bool) (int, error) {
	ms := ManaSpends{
		Winning: []int{},
	}

	spells := cmb.PopulateSpells()
	effects := cmb.PopulateEffects()

	for _, sp := range spells {
		player.HitPoints = 50
		player.Mana = 500
		player.Spells = cmb.PopulateSpells()
		boss.HitPoints = bossHP
		ms.SpellRound(player, boss, sp, 0, []*cmb.Spell{}, 1, effects, hardMode)
	}

	sort.Ints(ms.Winning)
	if len(ms.Winning) == 0 {
		return -1, fmt.Errorf("no winning game found")
	}

	return ms.Winning[0], nil
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
