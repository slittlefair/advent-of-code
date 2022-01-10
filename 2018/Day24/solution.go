package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type group struct {
	id             string
	units          int
	hp             int
	weaknesses     []string
	immunities     []string
	attackDamage   int
	attackType     string
	initiative     int
	effectivePower int
	army           string
}

var allGroupsOriginal = make(map[string]group)
var allGroups = make(map[string]group)

func readData(lines []string) {
	reStats := regexp.MustCompile("\\d+")
	reWeaknessesImmunities := regexp.MustCompile("\\((.*?)\\)")
	reDamage := regexp.MustCompile("(?s)does (\\d+ )(.*) damage")

	var activeArmy = "immuneSystem"
	i := 1
	for _, line := range lines {
		if line == "Infection:" {
			activeArmy = "infection"
		} else if line == "" || line == "Immune System:" {
			continue
		} else {
			stats := utils.StringSliceToIntSlice(reStats.FindAllString(line, -1))
			var weaknesses []string
			var immunities []string
			wi := reWeaknessesImmunities.FindStringSubmatch(line)
			if len(wi) > 1 {
				str := wi[1]
				str = strings.Replace(str, ",", "", -1)
				str = strings.Replace(str, ";", "", -1)
				str = strings.Replace(str, "to ", "", -1)
				doingImmunities := false
				fields := strings.Fields(str)
				for _, word := range fields {
					if word == "weak" {
						doingImmunities = false
					} else if word == "immune" {
						doingImmunities = true
					} else {
						if doingImmunities {
							immunities = append(immunities, word)
						} else {
							weaknesses = append(weaknesses, word)
						}
					}
				}
			}
			attackType := reDamage.FindStringSubmatch(line)
			var id string
			if activeArmy == "immuneSystem" {
				id = "System_" + strconv.Itoa(i)
			} else {
				id = "Infection_" + strconv.Itoa(i-10)
			}
			allGroupsOriginal[id] = group{
				id:             id,
				units:          stats[0],
				hp:             stats[1],
				attackDamage:   stats[2],
				initiative:     stats[3],
				weaknesses:     weaknesses,
				immunities:     immunities,
				attackType:     attackType[2],
				effectivePower: stats[0] * stats[2],
				army:           activeArmy,
			}
			i++
		}
	}
}

func calculateDamage(attack group, defence group) int {
	attackType := attack.attackType
	for _, imm := range defence.immunities {
		if imm == attackType {
			return 0
		}
	}
	for _, wkns := range defence.weaknesses {
		if wkns == attackType {
			return 2 * attack.effectivePower
		}
	}
	return attack.effectivePower
}

func bothArmiesHaveUnits() (bool, string) {
	numImmuneSystem := 0
	numInfections := 0
	for _, g := range allGroups {
		if g.army == "immuneSystem" {
			numImmuneSystem++
		} else {
			numInfections++
		}
	}
	if numImmuneSystem == 0 {
		return false, "infection"
	} else if numInfections == 0 {
		return false, "immuneSystem"
	} else {
		return true, ""
	}
}

func opponent(army string) (opponent string) {
	if army == "immuneSystem" {
		return "infection"
	}
	return "immuneSystem"
}

func killUnits(attack group, defence group, damage int) (group, int) {
	killed := damage / defence.hp
	if killed > defence.units {
		killed = defence.units
	}
	defence.units -= killed
	return defence, killed
}

func decideOrderOfSelection() []string {
	orderOfSelection := make([]string, len(allGroups))
	alreadySelected := make(map[string]bool)
	for i := 0; i < len(orderOfSelection); i++ {
		var selectedGroup group
		var selectedEP int
		for id, group := range allGroups {
			if _, ok := alreadySelected[id]; !ok {
				if group.effectivePower > selectedEP {
					selectedGroup = group
					selectedEP = group.effectivePower
				} else if group.effectivePower == selectedEP {
					if group.initiative > selectedGroup.initiative {
						selectedGroup = group
					}
				}
			}
		}
		alreadySelected[selectedGroup.id] = true
		orderOfSelection[i] = selectedGroup.id
	}
	return orderOfSelection
}

func decideWhichOpponentToAttack(attack group, potentialTargets []string) group {
	maxDamage := 0
	var toAttack group
	for _, id := range potentialTargets {
		candidate := allGroups[id]
		if candidate.army != attack.army {
			damage := calculateDamage(attack, candidate)
			if damage > maxDamage {
				maxDamage = damage
				toAttack = candidate
			} else if damage == maxDamage {
				if candidate.effectivePower > toAttack.effectivePower {
					toAttack = candidate
				} else if candidate.effectivePower == toAttack.effectivePower {
					if candidate.initiative > toAttack.initiative {
						toAttack = candidate
					}
				}
			}
		}
	}
	return toAttack
}

func battle(boost int) (int, string) {
	fighting := true
	survivingArmy := ""
	for k, v := range allGroupsOriginal {
		if v.army == "immuneSystem" {
			v.attackDamage += boost
			v.effectivePower = v.attackDamage * v.units
		}
		allGroups[k] = v
	}

	for fighting {
		orderOfSelection := decideOrderOfSelection()

		// targets maps a group to the opposing group it has selected to attack
		targets := make(map[string]string) // attack.id -> defence.id

		var groupIDtoEffectivePower = make(map[string]int)
		for _, g := range allGroups {
			groupIDtoEffectivePower[g.id] = g.effectivePower
		}

		for _, id := range orderOfSelection {
			attackingGroup := allGroups[id]

			// Get a slice of all targets so far, so we don't target a group more than once
			alreadyTargets := make(map[string]bool)
			for _, val := range targets {
				alreadyTargets[val] = true
			}

			// Now we know which group we're using, decide who they're targetting
			potentialTargets := []string{}
			opponent := opponent(attackingGroup.army)

			// Work out how much damage would be dealt to all opponents
			for _, grp := range allGroups {
				if _, ok := alreadyTargets[grp.id]; grp.army == opponent && calculateDamage(attackingGroup, grp) != 0 && !ok {
					potentialTargets = append(potentialTargets, grp.id)
				}
			}
			toAttack := decideWhichOpponentToAttack(attackingGroup, potentialTargets)

			// Now we know which opponent is being attacked and how much damage they'll receive
			if toAttack.id != "" {
				targets[attackingGroup.id] = toAttack.id
			}
		}

		// Attacking Phase
		groupIDToInitiative := make(map[string]int)
		for _, g := range allGroups {
			groupIDToInitiative[g.id] = g.initiative
		}

		orderOfAttack := []string{}
		for len(groupIDToInitiative) > 0 {
			maxInitiative := 0
			var maxInitiativeGroupID string
			for id, init := range groupIDToInitiative {
				if init > maxInitiative {
					maxInitiative = init
					maxInitiativeGroupID = id
				}
			}
			if _, ok := targets[maxInitiativeGroupID]; ok {
				orderOfAttack = append(orderOfAttack, maxInitiativeGroupID)
			}
			delete(groupIDToInitiative, maxInitiativeGroupID)
		}

		unitsKilled := 0
		for _, g := range orderOfAttack {
			if _, ok := allGroups[g]; ok {
				attack := allGroups[g]
				defence := allGroups[targets[g]]
				damage := calculateDamage(attack, defence)
				defence, killed := killUnits(attack, defence, damage)
				unitsKilled += killed
				defence.effectivePower = defence.units * defence.attackDamage
				// Remove the group if units are zero
				if defence.units <= 0 {
					delete(allGroups, defence.id)
				} else {
					allGroups[defence.id] = defence
				}
			}
			fighting, survivingArmy = bothArmiesHaveUnits()
			if !fighting {
				break
			}
		}
		if unitsKilled == 0 {
			fighting = false
			survivingArmy = "infection"
		}
	}
	unitsLeft := 0
	for _, g := range allGroups {
		unitsLeft += g.units
	}
	return unitsLeft, survivingArmy
}

func main() {
	lines := utils.ReadFile()
	readData(lines)
	boost := 0
	unitsLeft, survivingArmy := battle(boost)
	fmt.Println("Part 1:", unitsLeft)
	for survivingArmy != "immuneSystem" {
		boost++
		unitsLeft, survivingArmy = battle(boost)
	}
	fmt.Println("Part 2:", boost, unitsLeft)
}
