package shop

type Equipment struct {
	Name   string
	Cost   int
	Damage int
	Armour int
}

type Shop struct {
	Weapons []Equipment
	Armour  []Equipment
	Rings   []Equipment
}

func PopulateShop() *Shop {
	return &Shop{
		Armour: []Equipment{
			// Populate a dummy armour since it's optional
			{},
			{Armour: 1, Cost: 13},
			{Armour: 2, Cost: 31},
			{Armour: 3, Cost: 53},
			{Armour: 4, Cost: 75},
			{Armour: 5, Cost: 102},
		},
		Rings: []Equipment{
			// Populate two dummy rings since they're optional
			{},
			{},
			{Damage: 1, Cost: 25},
			{Damage: 2, Cost: 50},
			{Damage: 3, Cost: 100},
			{Armour: 1, Cost: 20},
			{Armour: 2, Cost: 40},
			{Armour: 3, Cost: 80},
		},
		Weapons: []Equipment{
			{Damage: 4, Cost: 8},
			{Damage: 5, Cost: 10},
			{Damage: 6, Cost: 25},
			{Damage: 7, Cost: 40},
			{Damage: 8, Cost: 74},
		},
	}
}
