package fight

import (
	"Advent-of-Code/2015/Day21/martial"
	"Advent-of-Code/2015/Day21/shop"
)

func MartialAttack(attacker, defender *martial.Martial) {
	damage := attacker.Damage - defender.Armour
	if damage < 1 {
		damage = 1
	}
	defender.HP -= damage
}

func Fight(player, boss *martial.Martial) bool {
	for {
		MartialAttack(player, boss)
		if boss.HP <= 0 {
			return true
		}
		MartialAttack(boss, player)
		if player.HP <= 0 {
			return false
		}
	}
}

func InitiatePlayerForFight(weapon, armour, ring1, ring2 shop.Equipment) (*martial.Martial, int) {
	return &martial.Martial{
		Armour: weapon.Armour + armour.Armour + ring1.Armour + ring2.Armour,
		Damage: weapon.Damage + armour.Damage + ring1.Damage + ring2.Damage,
		HP:     100,
	}, weapon.Cost + armour.Cost + ring1.Cost + ring2.Cost
}
