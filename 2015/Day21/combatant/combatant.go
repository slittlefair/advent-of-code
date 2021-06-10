package combatant

type Combatant struct {
	HitPoints int
	Damage    int
	Armour    int
	Cost      int
}

// IsDead returns a boolean for whether the combatant's hit points are less than or equal to 0
func (c *Combatant) IsDead() bool {
	return c.HitPoints <= 0
}
