package combatant

type Combatant struct {
	HitPoints       int
	Damage          int
	Armour          int
	Cost            int
	Mana            int
	Spells          Spells
	ManaSpent       int
	LowestManaSpent int
}

// IsDead returns a boolean for whether the combatant's hit points are less than or equal to 0
func (c *Combatant) IsDead() bool {
	return c.HitPoints <= 0
}

type Effect struct {
	Active            bool
	Duration          int
	DurationRemaining int
	Effect            func(player, boss *Combatant)
}

func ApplyEffect(attacker, defender *Combatant, e Effect) Effect {
	e.Effect(attacker, defender)
	e.DurationRemaining--
	if e.DurationRemaining == 0 {
		e.Active = false
	}
	return e
}

func Shield(player, boss *Combatant) {
	player.Armour = 7
}

func Poison(player, boss *Combatant) {
	boss.HitPoints -= 3
}

func Recharge(player, boss *Combatant) {
	player.Mana += 101
}

type Spell struct {
	Name      string
	Mana      int
	Damage    int
	HitPoints int
	Effect    string
}

func PopulateEffects() map[string]Effect {
	return map[string]Effect{
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
}

type Spells map[string]*Spell

func PopulateSpells() Spells {
	return Spells{
		"Magic Missile": {
			Name:   "Magic Missile",
			Mana:   53,
			Damage: 4,
			Effect: "None",
		},
		"Drain": {
			Name:      "Drain",
			Mana:      73,
			Damage:    2,
			HitPoints: 2,
			Effect:    "None",
		},
		"Shield": {
			Name:   "Shield",
			Mana:   113,
			Effect: "Shield",
		},
		"Poison": {
			Name:   "Poison",
			Mana:   173,
			Effect: "Poison",
		},
		"Recharge": {
			Name:   "Recharge",
			Mana:   229,
			Effect: "Recharge",
		},
	}
}

func (c *Combatant) CanCastSpells(spells Spells) bool {
	for _, spell := range spells {
		if c.Mana >= spell.Mana {
			return true
		}
	}
	return false
}

func (c *Combatant) SpellIsValid(spell *Spell, effects map[string]Effect) bool {
	if c.Mana < spell.Mana {
		return false
	}
	return !effects[spell.Effect].Active
}
