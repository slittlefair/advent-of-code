package shop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopulateShop(t *testing.T) {
	t.Run("returns shop with equipment as per instructions", func(t *testing.T) {
		want := &Shop{
			Weapons: []Equipment{
				{Cost: 8, Damage: 4},
				{Cost: 10, Damage: 5},
				{Cost: 25, Damage: 6},
				{Cost: 40, Damage: 7},
				{Cost: 74, Damage: 8},
			},
			Armour: []Equipment{
				{},
				{Cost: 13, Armour: 1},
				{Cost: 31, Armour: 2},
				{Cost: 53, Armour: 3},
				{Cost: 75, Armour: 4},
				{Cost: 102, Armour: 5},
			},
			Rings: []Equipment{
				{},
				{},
				{Cost: 25, Damage: 1},
				{Cost: 50, Damage: 2},
				{Cost: 100, Damage: 3},
				{Cost: 20, Armour: 1},
				{Cost: 40, Armour: 2},
				{Cost: 80, Armour: 3},
			},
		}
		got := PopulateShop()
		assert.Equal(t, want, got)
	})
}
