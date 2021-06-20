package mage

import (
	"Advent-of-Code/2015/Day21/martial"
	"reflect"
	"testing"
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
			if got := m.SpellIsValid(tt.args.spell, tt.args.effects); got != tt.want {
				t.Errorf("Mage.SpellIsValid() = %v, want %v", got, tt.want)
			}
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
			if got.Active != tt.want.Active {
				t.Errorf("ApplyEffect().Active = %v, want %v", got.Active, tt.want.Active)
			}
			if got.DurationRemaining != tt.want.DurationRemaining {
				t.Errorf("ApplyEffect().DurationRemaining = %v, want %v", got.DurationRemaining, tt.want.DurationRemaining)
			}
		})
	}
}

func TestShield(t *testing.T) {
	type args struct {
		mage *Mage
		boss *martial.Martial
	}
	tests := []struct {
		name string
		args args
		want *Mage
	}{
		{
			name: "applies shield effect to mage, setting armour to 7",
			args: args{
				mage: &Mage{},
				boss: &martial.Martial{},
			},
			want: &Mage{Armour: 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Shield(tt.args.mage, tt.args.boss)
			if !reflect.DeepEqual(tt.args.mage, tt.want) {
				t.Errorf("Shield() = %v, want %v", tt.args.mage, tt.want)
			}
		})
	}
}

func TestPoison(t *testing.T) {
	type args struct {
		mage *Mage
		boss *martial.Martial
	}
	tests := []struct {
		name string
		args args
		want *martial.Martial
	}{
		{
			name: "applies poison effect, reducing boss HP by 3",
			args: args{
				mage: &Mage{},
				boss: &martial.Martial{HP: 32},
			},
			want: &martial.Martial{HP: 29},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Poison(tt.args.mage, tt.args.boss)
			if !reflect.DeepEqual(tt.args.boss, tt.want) {
				t.Errorf("Poison() = %v, want %v", tt.args.boss, tt.want)
			}
		})
	}
}

func TestRecharge(t *testing.T) {
	type args struct {
		mage *Mage
		boss *martial.Martial
	}
	tests := []struct {
		name string
		args args
		want *Mage
	}{
		{
			name: "applies recharge effect to mage, increasing mana by 101",
			args: args{
				mage: &Mage{Mana: 274},
				boss: &martial.Martial{},
			},
			want: &Mage{Mana: 375},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Recharge(tt.args.mage, tt.args.boss)
			if !reflect.DeepEqual(tt.args.mage, tt.want) {
				t.Errorf("Recharge() = %v, want %v", tt.args.mage, tt.want)
			}
		})
	}
}

func TestPopulateSpells(t *testing.T) {
	tests := []struct {
		name string
		want Spells
	}{
		{
			name: "populates set of mage spells",
			want: Spells{
				"Magic Missile": {
					Mana:   53,
					Damage: 4,
					Effect: "None",
				},
				"Drain": {
					Mana:   73,
					Damage: 2,
					HP:     2,
					Effect: "None",
				},
				"Shield": {
					Mana:   113,
					Effect: "Shield",
				},
				"Poison": {
					Mana:   173,
					Effect: "Poison",
				},
				"Recharge": {
					Mana:   229,
					Effect: "Recharge",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopulateSpells(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopulateSpells() = %v, want %v", got, tt.want)
			}
		})
	}
}
