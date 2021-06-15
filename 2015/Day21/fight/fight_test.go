package fight

import (
	"Advent-of-Code/2015/Day21/martial"
	"Advent-of-Code/2015/Day21/shop"
	"reflect"
	"testing"
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
			if got := Fight(tt.args.player, tt.args.boss); got != tt.want {
				t.Errorf("Fight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitiatePlayerForFight(t *testing.T) {
	type args struct {
		weapon shop.Equipment
		armour shop.Equipment
		ring1  shop.Equipment
		ring2  shop.Equipment
	}
	tests := []struct {
		name  string
		args  args
		want  *martial.Martial
		want1 int
	}{
		{
			name: "correctly initiates a player's stats from given equipment",
			args: args{
				weapon: shop.Equipment{
					Armour: 1,
					Damage: 20,
					Cost:   300,
				},
				armour: shop.Equipment{
					Armour: 100,
					Damage: 2,
					Cost:   30,
				},
				ring1: shop.Equipment{
					Armour: 10,
					Damage: 200,
					Cost:   3,
				},
				ring2: shop.Equipment{},
			},
			want: &martial.Martial{
				Armour: 111,
				Damage: 222,
				HP:     100,
			},
			want1: 333,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := InitiatePlayerForFight(tt.args.weapon, tt.args.armour, tt.args.ring1, tt.args.ring2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitiatePlayerForFight() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("InitiatePlayerForFight() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
