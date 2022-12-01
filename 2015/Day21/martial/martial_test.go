package martial_test

import (
	"Advent-of-Code/2015/Day21/martial"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBoss(t *testing.T) {
	type args struct {
		input     []string
		hasArmour bool
	}
	tests := []struct {
		name               string
		args               args
		want               *martial.Martial
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if function expects an armour line but only has 2 lines",
			args: args{
				input: []string{
					"Hit Points: 3",
					"Damage: 9",
				},
				hasArmour: true,
			},
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if function expects an armour line but has 4 lines",
			args: args{
				input: []string{
					"Hit Points: 3",
					"Damage: 9",
					"Armor: 8",
					"Age: 1000",
				},
				hasArmour: true,
			},
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if function doesn't expect an armour line but has 1 line",
			args: args{
				input: []string{
					"Hit Points: 3",
				},
				hasArmour: false,
			},
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if function doesn't expect an armour line but has 3 lines",
			args: args{
				input: []string{
					"Hit Points: 3",
					"Damage: 9",
					"Armor: 8",
				},
				hasArmour: false,
			},
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: `returns an error if the Hit Points line in input doesn't have substring "Hit Points: "`,
			args: args{
				input: []string{
					"Hit Points:3",
					"Damage: 9",
					"Armor: 8",
				},
			},
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: `returns an error if the Damage line in input doesn't have substring "Damage: "`,
			args: args{
				input: []string{
					"Hit Points: 3",
					"Dmg: 9",
					"Armor: 8",
				},
				hasArmour: true,
			},
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: `returns an error if the Armour line in input doesn't have substring "Armor: "`,
			args: args{
				input: []string{
					"Hit Points: 3",
					"Damage: 9",
					"Armour: 8",
				},
				hasArmour: true,
			},
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if the Hit Points line doesn't split correctly",
			args: args{
				input: []string{
					"Hit Points:",
					"Damage: 9",
					"Armor: 8",
				},
				hasArmour: true,
			},
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if the Damage line doesn't split correctly",
			args: args{
				input: []string{
					"Hit Points: 3",
					"",
					"Armor: 8",
				},
				hasArmour: true,
			},
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if the Armour line doesn't split correctly",
			args: args{
				input: []string{
					"Hit Points: 3",
					"Damage: 9",
					"Armor: 8 or Armor: 9",
				},
				hasArmour: true,
			},
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if the Hit Points line doesn't yield a numeric value",
			args: args{
				input: []string{
					"Hit Points: ",
					"Damage: 9",
					"Armor: 8",
				},
				hasArmour: true,
			},
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if the Damage line doesn't yield a numeric value",
			args: args{
				input: []string{
					"Hit Points: 3",
					"Damage: some",
					"Armor: 8",
				},
				hasArmour: true,
			},
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if the Armour line doesn't yield a numeric value",
			args: args{
				input: []string{
					"Hit Points: 3",
					"Damage: 9",
					"Armor: about 8",
				},
				hasArmour: true,
			},
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns a constructed Boss from valid input when expecting Armour",
			args: args{
				input: []string{
					"Hit Points: 3",
					"Damage: 9",
					"Armor: 8",
				},
				hasArmour: true,
			},
			want: &martial.Martial{
				HP: 3, Damage: 9, Armour: 8,
			},
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns a constructed Boss from valid input when not expecting Armour",
			args: args{
				input: []string{
					"Hit Points: 10",
					"Damage: 1001",
				},
				hasArmour: false,
			},
			want: &martial.Martial{
				HP: 10, Damage: 1001,
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := martial.ParseBoss(tt.args.input, tt.args.hasArmour)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
