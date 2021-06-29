package martial

import (
	"fmt"
	"strconv"
	"strings"
)

type Martial struct {
	HP     int
	Damage int
	Armour int
}

func ParseBoss(input []string, hasArmour bool) (*Martial, error) {
	if hasArmour {
		if length := len(input); length != 3 {
			return nil, fmt.Errorf("something went wrong, expected 3 lines of input, got %d, %v", length, input)
		}
	} else {
		if length := len(input); length != 2 {
			return nil, fmt.Errorf("something went wrong, expected 2 lines of input, got %d, %v", length, input)
		}
	}

	boss := &Martial{}

	val := strings.Split(input[0], "Hit Points: ")
	if len(val) != 2 {
		return nil, fmt.Errorf("something went wrong, could not correctly split line %s", input[0])
	}
	hp, err := strconv.Atoi(val[1])
	if err != nil {
		return nil, err
	}
	boss.HP = hp

	val = strings.Split(input[1], "Damage: ")
	if len(val) != 2 {
		return nil, fmt.Errorf("something went wrong, could not correctly split line %s", input[0])
	}
	damage, err := strconv.Atoi(val[1])
	if err != nil {
		return nil, err
	}
	boss.Damage = damage

	if hasArmour {
		val = strings.Split(input[2], "Armor: ")
		if len(val) != 2 {
			return nil, fmt.Errorf("something went wrong, could not correctly split line %s", input[0])
		}
		armour, err := strconv.Atoi(val[1])
		if err != nil {
			return nil, err
		}
		boss.Armour = armour
	}

	return boss, nil
}
