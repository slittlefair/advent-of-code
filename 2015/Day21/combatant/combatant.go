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

func (c *Combatant) HasEffectInPlace(spell Spell) bool {
	for _, sp := range c.Spells {
		if sp.Name == spell.Name {
			return sp.Effect.Active
		}
	}
	return false
}

type Effect struct {
	Active            bool
	Duration          int
	DurationRemaining int
	Effect            func(player, boss *Combatant)
}

type Effects map[string]Effect

func (e *Effect) ApplyEffect(attacker, defender *Combatant) {
	e.Effect(attacker, defender)
	e.DurationRemaining--
	if e.DurationRemaining == 0 {
		e.Active = false
	}
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
	Effect    Effect
}

type Spells []Spell

func PopulateSpells() Spells {
	return Spells{
		{
			Name:   "Magic Missile",
			Mana:   53,
			Damage: 4,
		},
		{
			Name:      "Drain",
			Mana:      73,
			Damage:    2,
			HitPoints: 2,
		},
		{
			Name: "Shield",
			Mana: 113,
			Effect: Effect{
				Duration: 6,
				Effect:   Shield,
			},
		},
		{
			Name: "Poison",
			Mana: 173,
			Effect: Effect{
				Duration: 6,
				Effect:   Poison,
			},
		},
		{
			Name: "Recharge",
			Mana: 229,
			Effect: Effect{
				Duration: 5,
				Effect:   Recharge,
			},
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

func (c *Combatant) ValidSpells(spells Spells) Spells {
	validSpells := Spells{}
	for _, sp := range spells {
		if sp.Mana <= c.Mana {
			validSpells = append(validSpells, sp)
		}
	}
	return validSpells
}
