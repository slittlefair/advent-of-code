package mage_test

import (
	"Advent-of-Code/2015/Day21/martial"
	"Advent-of-Code/2015/Day22/mage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMage_SpellIsValid(t *testing.T) {
	type args struct {
		spell   *mage.Spell
		effects map[string]mage.Effect
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
				spell: &mage.Spell{
					Effect: "Fire Bolt",
					Mana:   100,
				},
				effects: map[string]mage.Effect{
					"Fire Bolt": {Active: false},
				},
			},
			want: false,
		},
		{
			name: "returns false if mage has mana for the spell but it is already active",
			mana: 100,
			args: args{
				spell: &mage.Spell{
					Effect: "Ray of Frost",
					Mana:   100,
				},
				effects: map[string]mage.Effect{
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
				spell: &mage.Spell{
					Effect: "Ray of Frost",
					Mana:   100,
				},
				effects: map[string]mage.Effect{
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
			m := &mage.Mage{
				Mana: tt.mana,
			}
			got := m.SpellIsValid(tt.args.spell, tt.args.effects)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestApply_Effect(t *testing.T) {
	type args struct {
		mage *mage.Mage
		boss *martial.Martial
		e    mage.Effect
	}
	tests := []struct {
		name string
		args args
		want mage.Effect
	}{
		{
			name: "applies an effect and returns it with reduced DurationRemaining",
			args: args{
				mage: &mage.Mage{},
				boss: &martial.Martial{},
				e: mage.Effect{
					Active:            true,
					DurationRemaining: 3,
					Effect:            mage.Shield,
				},
			},
			want: mage.Effect{
				Active:            true,
				DurationRemaining: 2,
			},
		},
		{
			name: "applies an effect and returns it deactivated if duration has run out",
			args: args{
				mage: &mage.Mage{},
				boss: &martial.Martial{},
				e: mage.Effect{
					Active:            true,
					DurationRemaining: 1,
					Effect:            mage.Shield,
				},
			},
			want: mage.Effect{
				Active:            false,
				DurationRemaining: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mage.ApplyEffect(tt.args.mage, tt.args.boss, tt.args.e)
			assert.Equal(t, tt.want.Active, got.Active)
			assert.Equal(t, tt.want.DurationRemaining, got.DurationRemaining)
		})
	}
}

func TestShield(t *testing.T) {
	t.Run("applies shield effect to mage, setting armour to 7", func(t *testing.T) {
		m := &mage.Mage{}
		mage.Shield(m, &martial.Martial{})
		assert.Equal(t, &mage.Mage{Armour: 7}, m)
	})
}

func TestPoison(t *testing.T) {
	t.Run("applies poison effect, reducing boss HP by 3", func(t *testing.T) {
		boss := &martial.Martial{HP: 32}
		mage.Poison(&mage.Mage{}, boss)
		assert.Equal(t, &martial.Martial{HP: 29}, boss)
	})
}

func TestRecharge(t *testing.T) {
	t.Run("applies recharge effect to mage, increasing mana by 101", func(t *testing.T) {
		m := &mage.Mage{Mana: 274}
		mage.Recharge(m, &martial.Martial{})
		assert.Equal(t, &mage.Mage{Mana: 375}, m)
	})
}
