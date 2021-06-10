package main

import (
	cmb "Advent-of-Code/2015/Day21/combatant"
	shop "Advent-of-Code/2015/Day21/shop"
	"reflect"
	"sort"
	"testing"
)

func TestFighters_Attack(t *testing.T) {
	tests := []struct {
		name   string
		Player *cmb.Combatant
		Boss   *cmb.Combatant
		want   *Fighters
	}{
		{
			name:   "reduces defender's hit points by 1 if attacker's damage is less than defender's armour",
			Player: &cmb.Combatant{Damage: 2},
			Boss:   &cmb.Combatant{Armour: 10, HitPoints: 12},
			want: &Fighters{
				Player: &cmb.Combatant{Damage: 2},
				Boss:   &cmb.Combatant{Armour: 10, HitPoints: 11},
			},
		},
		{
			name:   "reduces defender's hit points by 1 if attacker's damage is equal to defender's armour",
			Player: &cmb.Combatant{Damage: 10},
			Boss:   &cmb.Combatant{Armour: 10, HitPoints: 20},
			want: &Fighters{
				Player: &cmb.Combatant{Damage: 10},
				Boss:   &cmb.Combatant{Armour: 10, HitPoints: 19},
			},
		},
		{
			name:   "reduces defender's hit points by 1 if attacker's damage is one more than defender's armour",
			Player: &cmb.Combatant{Damage: 5},
			Boss:   &cmb.Combatant{Armour: 4, HitPoints: 9},
			want: &Fighters{
				Player: &cmb.Combatant{Damage: 5},
				Boss:   &cmb.Combatant{Armour: 4, HitPoints: 8},
			},
		},
		{
			name:   "reduces defender's hit points by the difference between attackers damage and defenders armour if damage is more than one greater than armour",
			Player: &cmb.Combatant{Damage: 16},
			Boss:   &cmb.Combatant{Armour: 11, HitPoints: 43},
			want: &Fighters{
				Player: &cmb.Combatant{Damage: 16},
				Boss:   &cmb.Combatant{Armour: 11, HitPoints: 38},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Fighters{
				Player: tt.Player,
				Boss:   tt.Boss,
			}
			f.Attack(tt.Player, tt.Boss)
			if !reflect.DeepEqual(f, tt.want) {
				t.Errorf("Fighters.Attack() = %v, want %v", f, tt.want)
			}
		})
	}
}

func TestFighters_Fight(t *testing.T) {
	type fields struct {
		Player            *cmb.Combatant
		Boss              *cmb.Combatant
		SuccessfulCosts   []int
		UnsuccessfulCosts []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "returns true if the player wins the fight",
			fields: fields{
				Player: &cmb.Combatant{HitPoints: 8, Damage: 5, Armour: 5},
				Boss:   &cmb.Combatant{HitPoints: 12, Damage: 7, Armour: 2},
			},
			want: true,
		},
		{
			name: "returns false if the boss wins the fight",
			fields: fields{
				Player: &cmb.Combatant{HitPoints: 8, Damage: 5, Armour: 5},
				Boss:   &cmb.Combatant{HitPoints: 120, Damage: 7, Armour: 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Fighters{
				Player:            tt.fields.Player,
				Boss:              tt.fields.Boss,
				SuccessfulCosts:   tt.fields.SuccessfulCosts,
				UnsuccessfulCosts: tt.fields.UnsuccessfulCosts,
			}
			if got := f.Fight(); got != tt.want {
				t.Errorf("Fighters.Fight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFighters_ParseBoss(t *testing.T) {
	type fields struct {
		Player            *cmb.Combatant
		Boss              *cmb.Combatant
		SuccessfulCosts   []int
		UnsuccessfulCosts []int
	}
	tests := []struct {
		name    string
		fields  fields
		arg     []string
		want    *Fighters
		wantErr bool
	}{
		{
			name: "returns an error if the input is less than 3 lines long",
			arg: []string{
				"Hit Points: 3",
				"Damage: 9",
			},
			want:    &Fighters{},
			wantErr: true,
		},
		{
			name: "returns an error if the input is more than 3 lines long",
			arg: []string{
				"Hit Points: 3",
				"Damage: 9",
				"Armor: 8",
				"Age: 1000",
			},
			want:    &Fighters{},
			wantErr: true,
		},
		{
			name: `returns an error if the Hit Points line in input doesn't have substring "Hit Points: "`,
			arg: []string{
				"Hit Points:3",
				"Damage: 9",
				"Armor: 8",
			},
			want:    &Fighters{},
			wantErr: true,
		},
		{
			name: `returns an error if the Damage line in input doesn't have substring "Damage: "`,
			arg: []string{
				"Hit Points: 3",
				"Dmg: 9",
				"Armor: 8",
			},
			want:    &Fighters{},
			wantErr: true,
		},
		{
			name: `returns an error if the Armour line in input doesn't have substring "Armor: "`,
			arg: []string{
				"Hit Points: 3",
				"Damage: 9",
				"Armour: 8",
			},
			want:    &Fighters{},
			wantErr: true,
		},
		{
			name: "returns an error if the Hit Points line doesn't yield a numeric value",
			arg: []string{
				"Hit Points: ",
				"Damage: 9",
				"Armor: 8",
			},
			want:    &Fighters{},
			wantErr: true,
		},
		{
			name: "returns an error if the Damage line doesn't yield a numeric value",
			arg: []string{
				"Hit Points: 3",
				"Damage: some",
				"Armor: 8",
			},
			want:    &Fighters{},
			wantErr: true,
		},
		{
			name: "returns an error if the Armour line doesn't yield a numeric value",
			arg: []string{
				"Hit Points: 3",
				"Damage: 9",
				"Armor: about 8",
			},
			want:    &Fighters{},
			wantErr: true,
		},
		{
			name: "returns a constructed Boss from valid input",
			arg: []string{
				"Hit Points: 3",
				"Damage: 9",
				"Armor: 8",
			},
			want: &Fighters{
				Boss: &cmb.Combatant{HitPoints: 3, Damage: 9, Armour: 8},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Fighters{
				Player:            tt.fields.Player,
				Boss:              tt.fields.Boss,
				SuccessfulCosts:   tt.fields.SuccessfulCosts,
				UnsuccessfulCosts: tt.fields.UnsuccessfulCosts,
			}
			if err := f.ParseBoss(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("Fighters.ParseBoss() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(f, tt.want) {
				t.Errorf("Fighters.ParseBoss() = %v, want %v", f, tt.want)
			}
		})
	}
}

func TestFighters_InitiatePlayerForFight(t *testing.T) {
	type args struct {
		weapon shop.Equipment
		armour shop.Equipment
		ring1  shop.Equipment
		ring2  shop.Equipment
	}
	tests := []struct {
		name string
		args args
		want *Fighters
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
			want: &Fighters{
				Player: &cmb.Combatant{
					Armour:    111,
					Damage:    222,
					Cost:      333,
					HitPoints: 100,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Fighters{}
			f.InitiatePlayerForFight(tt.args.weapon, tt.args.armour, tt.args.ring1, tt.args.ring2)
			if !reflect.DeepEqual(f, tt.want) {
				t.Errorf("Fighters.InitiatePlayerForFight() = %v, want %v", f, tt.want)
			}
		})
	}
}

func TestFighters_CheapestVictory(t *testing.T) {
	tests := []struct {
		name            string
		SuccessfulCosts []int
		want            int
		wantErr         bool
	}{
		{
			name:    "returns an error if there are no successful costs",
			want:    -1,
			wantErr: true,
		},
		{
			name:            "returns the lowest successful cost",
			SuccessfulCosts: []int{23, 99, 105, 22, 78, 888, 63},
			want:            22,
			wantErr:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Fighters{
				SuccessfulCosts: tt.SuccessfulCosts,
			}
			got, err := f.CheapestVictory()
			if (err != nil) != tt.wantErr {
				t.Errorf("Fighters.CheapestVictory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Fighters.CheapestVictory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFighters_DearestLoss(t *testing.T) {
	tests := []struct {
		name              string
		UnsuccessfulCosts []int
		want              int
		wantErr           bool
	}{
		{
			name:    "returns an error if there are no successful costs",
			want:    -1,
			wantErr: true,
		},
		{
			name:              "returns the lowest successful cost",
			UnsuccessfulCosts: []int{23, 99, 105, 22, 78, 888, 63},
			want:              888,
			wantErr:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Fighters{
				UnsuccessfulCosts: tt.UnsuccessfulCosts,
			}
			got, err := f.DearestLoss()
			if (err != nil) != tt.wantErr {
				t.Errorf("Fighters.DearestLoss() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Fighters.DearestLoss() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_runFights(t *testing.T) {
	tests := []struct {
		name                string
		arg                 []string
		wantBoss            *cmb.Combatant
		wantCheapestVictory int
		wantDearestLoss     int
		wantErr             bool
	}{
		{
			name:    "returns an error if there is an error parsing the Boss",
			wantErr: true,
		},
		// don't have a good example, so just use the real advent of code questions since we know
		// we got the correct solution. It's too much to compare everything, so just make sure the
		// Boss is parsed and we will end up with the right Cost arrays
		{
			name: "runs fights against given boss",
			arg: []string{
				"Hit Points: 109",
				"Damage: 8",
				"Armor: 2",
			},
			wantBoss: &cmb.Combatant{
				HitPoints: 109,
				Damage:    8,
				Armour:    2,
			},
			wantCheapestVictory: 111,
			wantDearestLoss:     188,
			wantErr:             false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runFights(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("runFights() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr == true {
				return
			}
			if !reflect.DeepEqual(got.Boss, tt.wantBoss) {
				t.Errorf("runFights().Boss = %v, want %v", got.Boss, tt.wantBoss)
			}
			sort.Ints(got.SuccessfulCosts)
			if cheapestVictory := got.SuccessfulCosts[0]; cheapestVictory != tt.wantCheapestVictory {
				t.Errorf("runFights().CheapestVictory = %v, want %v", cheapestVictory, tt.wantCheapestVictory)
			}
			sort.Sort(sort.Reverse(sort.IntSlice(got.UnsuccessfulCosts)))
			if dearestLoss := got.UnsuccessfulCosts[0]; dearestLoss != tt.wantDearestLoss {
				t.Errorf("runFights().CheapestVictory = %v, want %v", dearestLoss, tt.wantDearestLoss)
			}
		})
	}
}
