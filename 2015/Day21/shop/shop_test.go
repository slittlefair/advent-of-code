package shop

import (
	"reflect"
	"testing"
)

func TestPopulateShop(t *testing.T) {
	tests := []struct {
		name string
		want *Shop
	}{
		{
			name: "returns shop with equipment as per instructions",
			want: &Shop{
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopulateShop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopulateShop() = %v, want %v", got, tt.want)
			}
		})
	}
}
