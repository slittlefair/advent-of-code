package spellFight

import (
	"Advent-of-Code/2015/Day21/martial"
	"Advent-of-Code/2015/Day22/mage"
	"reflect"
	"testing"
)

func TestMartialAttack(t *testing.T) {
	type args struct {
		boss   *martial.Martial
		player *mage.Mage
	}
	tests := []struct {
		name  string
		args  args
		want  *martial.Martial
		want1 *mage.Mage
	}{
		{
			name: "reduces defender's hit points by 1 if attacker's damage is less than defender's armour",
			args: args{
				boss:   &martial.Martial{Damage: 2},
				player: &mage.Mage{Armour: 10, HP: 12},
			},
			want:  &martial.Martial{Damage: 2},
			want1: &mage.Mage{Armour: 10, HP: 11},
		},
		{
			name: "reduces defender's hit points by 1 if attacker's damage is equal to defender's armour",
			args: args{
				boss:   &martial.Martial{Damage: 10},
				player: &mage.Mage{Armour: 10, HP: 20},
			},
			want:  &martial.Martial{Damage: 10},
			want1: &mage.Mage{Armour: 10, HP: 19},
		},
		{
			name: "reduces defender's hit points by 1 if attacker's damage is one more than defender's armour",
			args: args{
				boss:   &martial.Martial{Damage: 5},
				player: &mage.Mage{Armour: 4, HP: 9},
			},
			want:  &martial.Martial{Damage: 5},
			want1: &mage.Mage{Armour: 4, HP: 8},
		},
		{
			name: "reduces defender's hit points by the difference between attackers damage and defenders armour if damage is more than one greater than armour",
			args: args{
				boss:   &martial.Martial{Damage: 16},
				player: &mage.Mage{Armour: 11, HP: 43},
			},
			want:  &martial.Martial{Damage: 16},
			want1: &mage.Mage{Armour: 11, HP: 38},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MartialAttack(tt.args.boss, tt.args.player)
			if !reflect.DeepEqual(tt.args.boss, tt.want) {
				t.Errorf("MartialAttack() = %v, want %v", tt.args.boss, tt.want)
			}
			if !reflect.DeepEqual(tt.args.player, tt.want1) {
				t.Errorf("MartialAttack() = %v, want %v", tt.args.player, tt.want1)
			}
		})
	}
}

func TestSpellAttack(t *testing.T) {
	type args struct {
		player  *mage.Mage
		boss    *martial.Martial
		spell   *mage.Spell
		effects map[string]mage.Effect
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]mage.Effect
		want1 *mage.Mage
		want2 *martial.Martial
	}{
		{
			name: "correctly attacks with the Magic Missile spell",
			args: args{
				player: &mage.Mage{Mana: 438},
				boss:   &martial.Martial{HP: 65},
				spell:  mage.SpellList["Magic Missile"],
				effects: map[string]mage.Effect{
					"Shield": {
						Active:            true,
						Duration:          6,
						DurationRemaining: 2,
					},
				},
			},
			want: map[string]mage.Effect{
				"Shield": {
					Active:            true,
					Duration:          6,
					DurationRemaining: 2,
				},
			},
			want1: &mage.Mage{Mana: 385},
			want2: &martial.Martial{HP: 61},
		},
		{
			name: "correctly attacks with the Drain spell",
			args: args{
				player: &mage.Mage{
					HP:   7,
					Mana: 222,
				},
				boss:  &martial.Martial{HP: 98},
				spell: mage.SpellList["Drain"],
				effects: map[string]mage.Effect{
					"Shield": {
						Active:            true,
						Duration:          6,
						DurationRemaining: 2,
					},
				},
			},
			want: map[string]mage.Effect{
				"Shield": {
					Active:            true,
					Duration:          6,
					DurationRemaining: 2,
				},
			},
			want1: &mage.Mage{HP: 9, Mana: 149},
			want2: &martial.Martial{HP: 96},
		},
		{
			name: "correctly attacks with the Shield spell",
			args: args{
				player: &mage.Mage{
					HP:   7,
					Mana: 200,
				},
				boss:  &martial.Martial{HP: 98},
				spell: mage.SpellList["Shield"],
				effects: map[string]mage.Effect{
					"Poison": {
						Active:            true,
						Duration:          6,
						DurationRemaining: 4,
					},
					"Shield": {
						Duration: 6,
					},
				},
			},
			want: map[string]mage.Effect{
				"Shield": {
					Active:            true,
					Duration:          6,
					DurationRemaining: 6,
				},
				"Poison": {
					Active:            true,
					Duration:          6,
					DurationRemaining: 4,
				},
			},
			want1: &mage.Mage{
				HP:   7,
				Mana: 87,
			},
			want2: &martial.Martial{HP: 98},
		},
		{
			name: "correctly attacks with the Poison spell",
			args: args{
				player: &mage.Mage{
					HP:   10,
					Mana: 541,
				},
				boss:  &martial.Martial{HP: 54},
				spell: mage.SpellList["Poison"],
				effects: map[string]mage.Effect{
					"Poison": {
						Duration: 6,
					},
					"Shield": {
						Duration: 6,
					},
				},
			},
			want: map[string]mage.Effect{
				"Shield": {
					Duration: 6,
				},
				"Poison": {
					Active:            true,
					Duration:          6,
					DurationRemaining: 6,
				},
			},
			want1: &mage.Mage{
				HP:   10,
				Mana: 368,
			},
			want2: &martial.Martial{HP: 54},
		},
		{
			name: "correctly attacks with the Recharge spell",
			args: args{
				player: &mage.Mage{
					HP:   10,
					Mana: 250,
				},
				boss:  &martial.Martial{HP: 9},
				spell: mage.SpellList["Recharge"],
				effects: map[string]mage.Effect{
					"Poison": {
						Duration: 6,
					},
					"Shield": {
						Duration: 6,
					},
					"Recharge": {
						Duration: 5,
					},
				},
			},
			want: map[string]mage.Effect{
				"Shield": {
					Duration: 6,
				},
				"Poison": {
					Duration: 6,
				},
				"Recharge": {
					Active:            true,
					Duration:          5,
					DurationRemaining: 5,
				},
			},
			want1: &mage.Mage{
				HP:   10,
				Mana: 21,
			},
			want2: &martial.Martial{HP: 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpellAttack(tt.args.player, tt.args.boss, tt.args.spell, tt.args.effects); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpellAttack() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.args.player, tt.want1) {
				t.Errorf("SpellAttack().player = %+v, want %+v", tt.args.player, tt.want1)
			}
			if !reflect.DeepEqual(tt.args.boss, tt.want2) {
				t.Errorf("SpellAttack().boss = %v, want %v", tt.args.boss, tt.want2)
			}
		})
	}
}

func compareEffects(want, got mage.Effect, mage *mage.Mage, boss *martial.Martial) bool {
	if want.Active != got.Active {
		return false
	}
	if want.Duration != got.Duration {
		return false
	}
	if want.DurationRemaining != got.DurationRemaining {
		return false
	}
	if reflect.ValueOf(want.Effect).Pointer() != reflect.ValueOf(got.Effect).Pointer() {
		return false
	}
	return true
}

func TestApplyEffects(t *testing.T) {
	type args struct {
		player  *mage.Mage
		boss    *martial.Martial
		effects map[string]mage.Effect
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]mage.Effect
		want1 *mage.Mage
		want2 *martial.Martial
	}{
		{
			name: "returns the supplied effects if none are active",
			args: args{
				player: &mage.Mage{HP: 5},
				boss:   &martial.Martial{HP: 3},
				effects: map[string]mage.Effect{
					"None": {},
					"Shield": {
						Duration: 6,
						Effect:   mage.Shield,
					},
					"Poison": {
						Duration: 6,
						Effect:   mage.Poison,
					},
					"Recharge": {
						Duration: 5,
						Effect:   mage.Recharge,
					},
				},
			},
			want: map[string]mage.Effect{
				"None": {},
				"Shield": {
					Duration: 6,
					Effect:   mage.Shield,
				},
				"Poison": {
					Duration: 6,
					Effect:   mage.Poison,
				},
				"Recharge": {
					Duration: 5,
					Effect:   mage.Recharge,
				},
			},
			want1: &mage.Mage{HP: 5},
			want2: &martial.Martial{HP: 3},
		},
		{
			name: "returns effects after Shield effect is applied",
			args: args{
				player: &mage.Mage{HP: 5},
				boss:   &martial.Martial{HP: 3},
				effects: map[string]mage.Effect{
					"None": {},
					"Shield": {
						Active:            true,
						Duration:          6,
						DurationRemaining: 2,
						Effect:            mage.Shield,
					},
					"Poison": {
						Duration: 6,
						Effect:   mage.Poison,
					},
					"Recharge": {
						Duration: 5,
						Effect:   mage.Recharge,
					},
				},
			},
			want: map[string]mage.Effect{
				"None": {},
				"Shield": {
					Active:            true,
					Duration:          6,
					DurationRemaining: 1,
					Effect:            mage.Shield,
				},
				"Poison": {
					Duration: 6,
					Effect:   mage.Poison,
				},
				"Recharge": {
					Duration: 5,
					Effect:   mage.Recharge,
				},
			},
			want1: &mage.Mage{Armour: 7, HP: 5},
			want2: &martial.Martial{HP: 3},
		},
		{
			name: "returns effects after Poison effect is applied",
			args: args{
				player: &mage.Mage{HP: 5},
				boss:   &martial.Martial{HP: 39},
				effects: map[string]mage.Effect{
					"None": {},
					"Shield": {
						Duration: 6,
						Effect:   mage.Shield,
					},
					"Poison": {
						Active:            true,
						Duration:          6,
						DurationRemaining: 5,
						Effect:            mage.Poison,
					},
					"Recharge": {
						Duration: 5,
						Effect:   mage.Recharge,
					},
				},
			},
			want: map[string]mage.Effect{
				"None": {},
				"Shield": {
					Duration: 6,
					Effect:   mage.Shield,
				},
				"Poison": {
					Active:            true,
					Duration:          6,
					DurationRemaining: 4,
					Effect:            mage.Poison,
				},
				"Recharge": {
					Duration: 5,
					Effect:   mage.Recharge,
				},
			},
			want1: &mage.Mage{HP: 5},
			want2: &martial.Martial{HP: 36},
		},
		{
			name: "returns effects after Recharge effect is applied",
			args: args{
				player: &mage.Mage{HP: 5, Mana: 654},
				boss:   &martial.Martial{HP: 39},
				effects: map[string]mage.Effect{
					"None": {},
					"Shield": {
						Duration: 6,
						Effect:   mage.Shield,
					},
					"Poison": {
						Duration: 6,
						Effect:   mage.Poison,
					},
					"Recharge": {
						Active:            true,
						Duration:          5,
						DurationRemaining: 3,
						Effect:            mage.Recharge,
					},
				},
			},
			want: map[string]mage.Effect{
				"None": {},
				"Shield": {
					Duration: 6,
					Effect:   mage.Shield,
				},
				"Poison": {
					Duration: 6,
					Effect:   mage.Poison,
				},
				"Recharge": {
					Active:            true,
					Duration:          5,
					DurationRemaining: 2,
					Effect:            mage.Recharge,
				},
			},
			want1: &mage.Mage{HP: 5, Mana: 755},
			want2: &martial.Martial{HP: 39},
		},
		{
			name: "returns effects after all effects are applied",
			args: args{
				player: &mage.Mage{HP: 5, Mana: 654},
				boss:   &martial.Martial{HP: 39},
				effects: map[string]mage.Effect{
					"None": {},
					"Shield": {
						Active:            true,
						Duration:          6,
						DurationRemaining: 2,
						Effect:            mage.Shield,
					},
					"Poison": {
						Active:            true,
						Duration:          6,
						DurationRemaining: 6,
						Effect:            mage.Poison,
					},
					"Recharge": {
						Active:            true,
						Duration:          5,
						DurationRemaining: 3,
						Effect:            mage.Recharge,
					},
				},
			},
			want: map[string]mage.Effect{
				"None": {},
				"Shield": {
					Active:            true,
					Duration:          6,
					DurationRemaining: 1,
					Effect:            mage.Shield,
				},
				"Poison": {
					Active:            true,
					Duration:          6,
					DurationRemaining: 5,
					Effect:            mage.Poison,
				},
				"Recharge": {
					Active:            true,
					Duration:          5,
					DurationRemaining: 2,
					Effect:            mage.Recharge,
				},
			},
			want1: &mage.Mage{Armour: 7, HP: 5, Mana: 755},
			want2: &martial.Martial{HP: 36},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ApplyEffects(tt.args.player, tt.args.boss, tt.args.effects)
			for k, e := range got {
				if !compareEffects(tt.want[k], e, tt.args.player, tt.args.boss) {
					t.Errorf("ApplyEffects() = %v, want %v", e, tt.want[k])
				}
			}
			if !reflect.DeepEqual(tt.args.player, tt.want1) {
				t.Errorf("ApplyEffects() player = %v, want %v", tt.args.player, tt.want1)
			}
			if !reflect.DeepEqual(tt.args.boss, tt.want2) {
				t.Errorf("ApplyEffects() boss = %v, want %v", tt.args.boss, tt.want2)
			}
		})
	}
}

func TestManaSpent_CompareMana(t *testing.T) {
	tests := []struct {
		name            string
		LowestManaSpent int
		currentMana     int
		want            int
	}{
		{
			name:            "does not set LowestManaSpent if currentMana is higher than it",
			LowestManaSpent: 452,
			currentMana:     786,
			want:            452,
		},
		{
			name:            "does not set LowestManaSpent if currentMana is equal to it",
			LowestManaSpent: 452,
			currentMana:     452,
			want:            452,
		},
		{
			name:            "does set LowestManaSpent if currentMana is less than it",
			LowestManaSpent: 452,
			currentMana:     213,
			want:            213,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &ManaSpent{
				LowestManaSpent: tt.LowestManaSpent,
			}
			ms.CompareMana(tt.currentMana)
			if ms.LowestManaSpent != tt.want {
				t.Errorf("ManaSpent.CompareMana().LowestManaSpent = %d, want %d", ms.LowestManaSpent, tt.want)
			}
		})
	}
}

func TestManaSpent_SpellRound(t *testing.T) {
	type args struct {
		player   mage.Mage
		boss     martial.Martial
		spell    *mage.Spell
		effects  map[string]mage.Effect
		hardMode bool
	}
	tests := []struct {
		name            string
		LowestManaSpent int
		args            args
		want            int
	}{
		{
			name:            "returns and doesn't set lowest mana if player dies due to hard mode",
			LowestManaSpent: 1000,
			args: args{
				player: mage.Mage{HP: 1, Mana: 652, ManaSpent: 451},
				boss:   martial.Martial{Damage: 8, HP: 10},
				spell:  mage.SpellList["Magic Missile"],
				effects: map[string]mage.Effect{
					"Poison": {
						Active:            true,
						Duration:          6,
						DurationRemaining: 3,
						Effect:            mage.Poison,
					},
				},
				hardMode: true,
			},
			want: 1000,
		},
		{
			name:            "returns and sets lowest mana if boss dies from an effect at the start of player's turn",
			LowestManaSpent: 1000,
			args: args{
				player: mage.Mage{HP: 32, Mana: 652, ManaSpent: 451},
				boss:   martial.Martial{Damage: 8, HP: 1},
				spell:  mage.SpellList["Magic Missile"],
				effects: map[string]mage.Effect{
					"Poison": {
						Active:            true,
						Duration:          6,
						DurationRemaining: 3,
						Effect:            mage.Poison,
					},
				},
				hardMode: false,
			},
			want: 451,
		},
		{
			name:            "returns and doesn't set lowest mana if spell is not valid",
			LowestManaSpent: 1000,
			args: args{
				player: mage.Mage{HP: 32, Mana: 1, ManaSpent: 451},
				boss:   martial.Martial{Damage: 8, HP: 10},
				spell:  mage.SpellList["Magic Missile"],
				effects: map[string]mage.Effect{
					"Poison": {
						Active:            true,
						Duration:          6,
						DurationRemaining: 3,
						Effect:            mage.Poison,
					},
				},
				hardMode: false,
			},
			want: 1000,
		},
		{
			name:            "returns and sets lowest mana if boss dies from spell attack",
			LowestManaSpent: 1000,
			args: args{
				player:   mage.Mage{HP: 32, Mana: 652, ManaSpent: 451},
				boss:     martial.Martial{Damage: 8, HP: 1},
				spell:    mage.SpellList["Magic Missile"],
				effects:  map[string]mage.Effect{},
				hardMode: false,
			},
			want: 504,
		},
		{
			name:            "returns and sets lowest mana if boss dies from an effect at the start of boss's turn",
			LowestManaSpent: 1000,
			args: args{
				player: mage.Mage{HP: 32, Mana: 652, ManaSpent: 451},
				boss:   martial.Martial{Damage: 8, HP: 9},
				spell:  mage.SpellList["Magic Missile"],
				effects: map[string]mage.Effect{
					"Poison": {
						Active:            true,
						Duration:          6,
						DurationRemaining: 3,
						Effect:            mage.Poison,
					},
				},
				hardMode: false,
			},
			want: 504,
		},
		{
			name:            "returns and doesn't set lowest mana if player dies from boss attack",
			LowestManaSpent: 1000,
			args: args{
				player: mage.Mage{HP: 1, Mana: 1000, ManaSpent: 451},
				boss:   martial.Martial{Damage: 8, HP: 100},
				spell:  mage.SpellList["Magic Missile"],
				effects: map[string]mage.Effect{
					"Poison": {
						Active:            true,
						Duration:          6,
						DurationRemaining: 3,
						Effect:            mage.Poison,
					},
				},
				hardMode: false,
			},
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &ManaSpent{
				LowestManaSpent: tt.LowestManaSpent,
			}
			ms.SpellRound(tt.args.player, tt.args.boss, tt.args.spell, tt.args.effects, tt.args.hardMode)
			if ms.LowestManaSpent != tt.want {
				t.Errorf("ManaSpent.SpellRound() = %d, want %d", ms.LowestManaSpent, tt.want)
			}
		})
	}
}

func TestSpellFight(t *testing.T) {
	// we don't really have any useful examples to copmpare, so just use the actual puzzle inputs
	// since we already know it works.
	type args struct {
		boss     martial.Martial
		bossHP   int
		hardMode bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "returns the expected lowest mana spent, non-hard mode",
			args: args{
				boss:     martial.Martial{Damage: 10, HP: 71},
				bossHP:   71,
				hardMode: false,
			},
			want: 1824,
		},
		{
			name: "returns the expected lowest mana spent, non-hard mode",
			args: args{
				boss:     martial.Martial{Damage: 10, HP: 71},
				bossHP:   71,
				hardMode: true,
			},
			want: 1937,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpellFight(tt.args.boss, tt.args.bossHP, tt.args.hardMode); got != tt.want {
				t.Errorf("SpellFight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunSpellFights(t *testing.T) {
	// we don't really have any useful examples to copmpare, so just use the actual puzzle inputs
	// since we already know it works.
	tests := []struct {
		name    string
		input   []string
		want    int
		want1   int
		wantErr bool
	}{
		{
			name:    "returns an error if boss cannot be parsed",
			want:    -1,
			want1:   -1,
			wantErr: true,
		},
		{
			name: "returns expected lowest mana for both difficulty modes for a given boss input",
			input: []string{
				"Hit Points: 71",
				"Damage: 10",
			},
			want:    1824,
			want1:   1937,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := RunSpellFights(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunSpellFights() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RunSpellFights() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("RunSpellFights() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
