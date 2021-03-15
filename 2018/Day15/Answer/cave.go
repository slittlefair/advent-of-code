package day15

import (
	"fmt"
	"sort"
	"strings"
)

type Cave struct {
	Units SortableUnits
	Map   Map
}

const (
	KindSpace = 1 << iota
	KindElf
	KindGoblin
	KindWall
	KindHighlight
)

var KindRunes = map[int]rune{
	KindSpace:     '.',
	KindElf:       'E',
	KindGoblin:    'G',
	KindWall:      '#',
	KindHighlight: '@',
}

var RuneKinds = map[rune]int{
	'.': KindSpace,
	'E': KindElf,
	'G': KindGoblin,
	'#': KindWall,
}

func IsUnit(bit int) bool {
	return (KindElf|KindGoblin)&bit != 0
}

func NewCave(input []string, elfPower int) *Cave {
	c := &Cave{}
	c.ParseMap(input, elfPower)
	return c
}

func (c *Cave) ParseMap(input []string, elfPower int) {
	m := make(Map)

	for y, row := range input {
		for x, col := range row {
			kind, ok := RuneKinds[col]
			if !ok {
				kind = KindWall
			}

			tile := &Tile{Kind: kind}
			if IsUnit(kind) {
				c.Units = append(c.Units, NewUnit(tile, kind, elfPower))
			}
			m.SetTile(tile, x, y)
		}
	}
	c.Map = m
}

func (c Cave) PrintMap(highlight *Tile) {
	for y := 0; y < len(c.Map); y++ {
		var units []string
		for x := 0; x < len(c.Map[y]); x++ {
			t := c.Map.Tile(x, y)
			if t == highlight {
				fmt.Print(string(KindRunes[KindHighlight]))
			} else {
				fmt.Print(string(KindRunes[t.Kind]))
			}

			if t.Unit != nil {
				units = append(units, fmt.Sprintf("%c(%d)", KindRunes[t.Unit.Kind], t.Unit.Hitpoints))
			}
		}
		if len(units) > 0 {
			fmt.Print("  ", strings.Join(units, ", "))
		}
		fmt.Println()
	}
}

func (c Cave) PrintDistance(t *Tile) {
	distances, _ := c.Map.FindWalkableTiles(t)
	for y := 0; y < len(c.Map); y++ {
		for x := 0; x < len(c.Map[y]); x++ {
			curT := c.Map.Tile(x, y)
			if d, ok := distances[curT]; ok && curT != t {
				fmt.Print(d)
			} else {
				fmt.Print(string(KindRunes[curT.Kind]))
			}
		}
		fmt.Println()
	}
}

// Status returns the sum of remaining hitpoints for all units and true if the fight is still ongoing.
func (c Cave) Status() (int, bool) {
	var elves, goblins bool
	var hp int

	for _, u := range c.Units {
		if u.Hitpoints <= 0 {
			continue
		}
		if u.Kind == KindElf {
			elves = true
		} else {
			goblins = true
		}
		hp += u.Hitpoints
	}

	return hp, elves && goblins
}

func (c *Cave) RemoveTheDead() {
	var newUnits SortableUnits
	for _, unit := range c.Units {
		if unit.Hitpoints > 0 {
			newUnits = append(newUnits, unit)
		}
	}
	c.Units = newUnits
}

func (c *Cave) RemoveUnit(u *Unit) {
	u.Tile.Kind = KindSpace
	u.Tile.Unit = nil
	u.Tile = nil
}

// Tick returns false if combat ended during the round, and whether or not an elf has died this round.
func (c *Cave) Tick(stopOnElfDeath bool) (bool, bool) {
	c.RemoveTheDead()
	sort.Sort(c.Units)

	for _, unit := range c.Units {
		if unit.Hitpoints <= 0 {
			continue
		}
		if !unit.Targets(c) {
			return false, false
		}
		unit.Move(c)
		if unit.Attack(c) && stopOnElfDeath {
			return false, true
		}
	}
	return true, false
}
