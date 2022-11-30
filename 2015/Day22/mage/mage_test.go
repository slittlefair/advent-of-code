package mage

import (
	"Advent-of-Code/2015/Day21/martial"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMage_SpellIsValid(t *testing.T) {
	type args struct {
		spell   *Spell
		effects map[string]Effect
	}
	tests := []struct {
		name string
		mana int
		args args
		want bool
	}{
		{
			name: "returns false if mage does not have mana for the spell",
			mana: 99,
			args: args{
				spell: &Spell{
					Effect: "Fire Bolt",
					Mana:   100,
				},
				effects: map[string]Effect{
					"Fire Bolt": {Active: false},
				},
			},
			want: false,
		},
		{
			name: "returns false if mage has mana for the spell but it is already active",
			mana: 100,
			args: args{
				spell: &Spell{
					Effect: "Ray of Frost",
					Mana:   100,
				},
				effects: map[string]Effect{
					"Fire Bolt":    {Active: true},
					"Ray of Frost": {Active: true},
					"Chill Touch":  {Active: false},
				},
			},
			want: false,
		},
		{
			name: "returns true if mage has mana for the spell and is not active",
			mana: 100,
			args: args{
				spell: &Spell{
					Effect: "Ray of Frost",
					Mana:   100,
				},
				effects: map[string]Effect{
					"Fire Bolt":    {Active: true},
					"Ray of Frost": {Active: false},
					"Chill Touch":  {Active: false},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mage{
				Mana: tt.mana,
			}
			got := m.SpellIsValid(tt.args.spell, tt.args.effects)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestApplyEffect(t *testing.T) {
	type args struct {
		mage *Mage
		boss *martial.Martial
		e    Effect
	}
	tests := []struct {
		name string
		args args
		want Effect
	}{
		{
			name: "applies an effect and returns it with reduced DurationRemaining",
			args: args{
				mage: &Mage{},
				boss: &martial.Martial{},
				e: Effect{
					Active:            true,
					DurationRemaining: 3,
					Effect:            Shield,
				},
			},
			want: Effect{
				Active:            true,
				DurationRemaining: 2,
			},
		},
		{
			name: "applies an effect and returns it deactivated if duration has run out",
			args: args{
				mage: &Mage{},
				boss: &martial.Martial{},
				e: Effect{
					Active:            true,
					DurationRemaining: 1,
					Effect:            Shield,
				},
			},
			want: Effect{
				Active:            false,
				DurationRemaining: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ApplyEffect(tt.args.mage, tt.args.boss, tt.args.e)
			assert.Equal(t, tt.want.Active, got.Active)
			assert.Equal(t, tt.want.DurationRemaining, got.DurationRemaining)
		})
	}
}

func TestShield(t *testing.T) {
	t.Run("applies shield effect to mage, setting armour to 7", func(t *testing.T) {
		mage := &Mage{}
		Shield(mage, &martial.Martial{})
		assert.Equal(t, &Mage{Armour: 7}, mage)
	})
}

func TestPoison(t *testing.T) {
	t.Run("applies poison effect, reducing boss HP by 3", func(t *testing.T) {
		boss := &martial.Martial{HP: 32}
		Poison(&Mage{}, boss)
		assert.Equal(t, &martial.Martial{HP: 29}, boss)
	})
}

func TestRecharge(t *testing.T) {
	t.Run("applies recharge effect to mage, increasing mana by 101", func(t *testing.T) {
		mage := &Mage{Mana: 274}
		Recharge(mage, &martial.Martial{})
		assert.Equal(t, &Mage{Mana: 375}, mage)
	})
}
