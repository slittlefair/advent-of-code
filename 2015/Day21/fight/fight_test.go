package fight

import (
	"Advent-of-Code/2015/Day21/martial"
	"Advent-of-Code/2015/Day21/shop"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMartialAttack(t *testing.T) {
	type args struct {
		attacker *martial.Martial
		defender *martial.Martial
	}
	tests := []struct {
		name  string
		args  args
		want  *martial.Martial
		want1 *martial.Martial
	}{
		{
			name: "reduces defender's hit points by 1 if attacker's damage is less than defender's armour",
			args: args{
				attacker: &martial.Martial{Damage: 2},
				defender: &martial.Martial{Armour: 10, HP: 12},
			},
			want:  &martial.Martial{Damage: 2},
			want1: &martial.Martial{Armour: 10, HP: 11},
		},
		{
			name: "reduces defender's hit points by 1 if attacker's damage is equal to defender's armour",
			args: args{
				attacker: &martial.Martial{Damage: 10},
				defender: &martial.Martial{Armour: 10, HP: 20},
			},
			want:  &martial.Martial{Damage: 10},
			want1: &martial.Martial{Armour: 10, HP: 19},
		},
		{
			name: "reduces defender's hit points by 1 if attacker's damage is one more than defender's armour",
			args: args{
				attacker: &martial.Martial{Damage: 5},
				defender: &martial.Martial{Armour: 4, HP: 9},
			},
			want:  &martial.Martial{Damage: 5},
			want1: &martial.Martial{Armour: 4, HP: 8},
		},
		{
			name: "reduces defender's hit points by the difference between attackers damage and defenders armour if damage is more than one greater than armour",
			args: args{
				attacker: &martial.Martial{Damage: 16},
				defender: &martial.Martial{Armour: 11, HP: 43},
			},
			want:  &martial.Martial{Damage: 16},
			want1: &martial.Martial{Armour: 11, HP: 38},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MartialAttack(tt.args.attacker, tt.args.defender)
			assert.Equal(t, tt.want, tt.args.attacker)
			assert.Equal(t, tt.want1, tt.args.defender)
		})
	}
}

func TestFight(t *testing.T) {
	type args struct {
		player *martial.Martial
		boss   *martial.Martial
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns true if the player wins the fight",
			args: args{
				player: &martial.Martial{HP: 8, Damage: 5, Armour: 5},
				boss:   &martial.Martial{HP: 12, Damage: 7, Armour: 2},
			},
			want: true,
		},
		{
			name: "returns false if the boss wins the fight",
			args: args{
				player: &martial.Martial{HP: 8, Damage: 5, Armour: 5},
				boss:   &martial.Martial{HP: 120, Damage: 7, Armour: 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Fight(tt.args.player, tt.args.boss)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInitiatePlayerForFight(t *testing.T) {
	t.Run("correctly initiates a player's stats from given equipment", func(t *testing.T) {
		weapon := shop.Equipment{
			Armour: 1,
			Damage: 20,
			Cost:   300,
		}
		armour := shop.Equipment{
			Armour: 100,
			Damage: 2,
			Cost:   30,
		}
		ring1 := shop.Equipment{
			Armour: 10,
			Damage: 200,
			Cost:   3,
		}
		ring2 := shop.Equipment{}
		want := &martial.Martial{
			Armour: 111,
			Damage: 222,
			HP:     100,
		}
		got, got1 := InitiatePlayerForFight(weapon, armour, ring1, ring2)
		assert.Equal(t, want, got)
		assert.Equal(t, 333, got1)
	})
}
