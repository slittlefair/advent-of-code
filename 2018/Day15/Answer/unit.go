package day15

import (
	"math"
	"sort"
)

type Unit struct {
	Kind      int
	Hitpoints int
	Power     int
	Tile      *Tile
}

const (
	defaultHitpoints = 200
	defaultPower     = 3
)

type SortableUnits []*Unit

func (s SortableUnits) Len() int {
	return len(s)
}

func (s SortableUnits) Less(i, j int) bool {
	if s[i].Tile.Y == s[j].Tile.Y {
		return s[i].Tile.X < s[j].Tile.X
	}
	return s[i].Tile.Y < s[j].Tile.Y
}

func (s SortableUnits) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func NewUnit(tile *Tile, kind, elfPower int) *Unit {
	unit := &Unit{
		Kind:      kind,
		Hitpoints: defaultHitpoints,
		Power:     defaultPower,
		Tile:      tile,
	}
	tile.Unit = unit
	if unit.Kind == KindElf {
		unit.Power = elfPower
	}
	return unit
}

func (u Unit) Targets(c *Cave) bool {
	for _, unit := range c.Units {
		if unit.Kind != u.Kind && unit.Hitpoints > 0 {
			return true
		}
	}
	return false
}

// NextTile returns the next tile a unit should move to and the target tile it is moving towards, or nil if no reachable
// target has been found.
func (u *Unit) NextTile(c *Cave) (*Tile, *Tile) {
	var targets SortableTiles

	closestTargetDistance := math.MaxInt32
	distances, path := c.Map.FindWalkableTiles(u.Tile)
	enemies := u.Enemies(c)

	for _, enemy := range enemies {
		for _, target := range enemy.Tile.WalkableNeighbors() {
			if distance, ok := distances[target]; ok && distance <= closestTargetDistance {
				if distance < closestTargetDistance {
					closestTargetDistance = distance
					targets = SortableTiles{}
				}
				targets = append(targets, target)
			}
		}
	}
	sort.Sort(targets)
	if len(targets) > 0 {
		target := targets[0]
		current := target
		for {
			if path[current] == u.Tile {
				return current, target
			}
			current = path[current]
		}
	}
	return nil, nil
}

// Enemies returns a list of enemy units sorted by map position in reading order
func (u *Unit) Enemies(c *Cave) SortableUnits {
	var enemies SortableUnits
	for _, unit := range c.Units {
		if unit.Kind != u.Kind && unit.Hitpoints > 0 {
			enemies = append(enemies, unit)
		}
	}
	sort.Sort(enemies)
	return enemies
}

func (u *Unit) EnemyNeighbor(c *Cave) *Unit {
	var target *Unit
	for _, offset := range offsets {
		if t := c.Map.Tile(u.Tile.X+offset.X, u.Tile.Y+offset.Y); t != nil && t.Unit != nil && t.Unit.Kind != u.Kind && t.Unit.Hitpoints > 0 {
			if target == nil || t.Unit.Hitpoints < target.Hitpoints {
				target = t.Unit
			}
		}
	}
	return target
}

func (u *Unit) Move(c *Cave) {
	if u.EnemyNeighbor(c) != nil {
		return
	}
	if next, _ := u.NextTile(c); next != nil {
		next.Unit = u
		next.Kind = u.Kind
		u.Tile.Kind = KindSpace
		u.Tile.Unit = nil
		u.Tile = next
	}
}

func (u *Unit) Attack(c *Cave) bool {
	if enemy := u.EnemyNeighbor(c); enemy != nil {
		killed := enemy.Damage(c, u.Power)
		return killed && enemy.Kind == KindElf
	}
	return false
}

func (u *Unit) Damage(c *Cave, damage int) bool {
	u.Hitpoints -= damage
	if u.Hitpoints <= 0 {
		c.RemoveUnit(u)
		return true
	}
	return false
}
