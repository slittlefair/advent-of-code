package spellFight

import (
	helpers "Advent-of-Code"
	"Advent-of-Code/2015/Day21/martial"
	"Advent-of-Code/2015/Day22/mage"
)

func MartialAttack(boss *martial.Martial, player *mage.Mage) {
	damage := boss.Damage - player.Armour
	if damage < 1 {
		damage = 1
	}
	player.HP -= damage
}

func SpellAttack(player *mage.Mage, boss *martial.Martial, spell *mage.Spell, effects map[string]mage.Effect) map[string]mage.Effect {
	player.Mana -= spell.Mana
	player.HP += spell.HP
	boss.HP -= spell.Damage
	if spell.Effect != "None" {
		eff := effects[spell.Effect]
		eff.Active = true
		eff.DurationRemaining = eff.Duration
		effects[spell.Effect] = eff
	}
	return effects
}

func ApplyEffects(player *mage.Mage, boss *martial.Martial, effects map[string]mage.Effect) map[string]mage.Effect {
	effs := map[string]mage.Effect{}
	for k, eff := range effects {
		if eff.Active {
			effs[k] = mage.ApplyEffect(player, boss, eff)
		} else {
			effs[k] = eff
		}
	}
	return effs
}

type ManaSpent struct {
	LowestManaSpent int
}

func (ms *ManaSpent) CompareMana(currentMana int) {
	if currentMana < ms.LowestManaSpent {
		ms.LowestManaSpent = currentMana
	}
}

func (ms *ManaSpent) SpellRound(player mage.Mage, boss martial.Martial, spell *mage.Spell, effects map[string]mage.Effect, hardMode bool) {
	// Set playar armour to 0
	player.Armour = 0

	// If we're in hard mode reduce player's HP by 1
	if hardMode {
		player.HP--
		if player.HP <= 0 {
			return
		}
	}

	// Apply the effects
	effects = ApplyEffects(&player, &boss, effects)

	// If the spell effects have caused the boss to die, end the fight and compare mana scores
	if boss.HP <= 0 {
		ms.CompareMana(player.ManaSpent)
		return
	}

	// If the spell isn't actually valid, either it's effect is already active or we don't have
	// enough mana to cast it, then the player has lost so end the fight
	if !player.SpellIsValid(spell, effects) {
		return
	}

	// Cast the spell and accumulate spent mana
	effects = SpellAttack(&player, &boss, spell, effects)
	player.ManaSpent += spell.Mana

	// If the spell causes the boss to die, end the fight and compare mana scores
	if boss.HP <= 0 {
		ms.CompareMana(player.ManaSpent)
		return
	}

	// At the start of the boss' turn, set player armour to 0 and apply spell effects
	player.Armour = 0
	effects = ApplyEffects(&player, &boss, effects)

	// If the spell effects have caused the boss to die, end the fight and compare mana scores
	if boss.HP <= 0 {
		ms.CompareMana(player.ManaSpent)
		return
	}

	// Boss attacks the player with a marital attack
	MartialAttack(&boss, &player)

	// If the boss attack causes the player to die, end the fight
	if player.HP <= 0 {
		return
	}

	// Carry on the fight with each spell
	for _, vs := range player.Spells {
		ms.SpellRound(player, boss, vs, effects, hardMode)
	}
}

func SpellFight(player mage.Mage, boss martial.Martial, bossHP int, hardMode bool) int {
	ms := ManaSpent{
		LowestManaSpent: helpers.Infinty,
	}

	spells := mage.PopulateSpells()
	effects := mage.Effects

	for _, sp := range spells {
		player.HP = 50
		player.Mana = 500
		player.ManaSpent = 0
		player.Spells = mage.PopulateSpells()
		boss.HP = bossHP
		ms.SpellRound(player, boss, sp, effects, hardMode)
	}

	return ms.LowestManaSpent
}
