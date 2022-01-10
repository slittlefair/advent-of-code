package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
)

type VisitedHouses struct {
	Map       map[utils.Co]bool
	Santa     utils.Co
	RoboSanta utils.Co
}

func createVisitedHouses() *VisitedHouses {
	return &VisitedHouses{
		Map: map[utils.Co]bool{
			{X: 0, Y: 0}: true,
		},
	}
}

func (vh *VisitedHouses) moveSanta(dir string) {
	if dir == "<" {
		vh.Santa.X--
	} else if dir == ">" {
		vh.Santa.X++
	} else if dir == "^" {
		vh.Santa.Y--
	} else if dir == "v" {
		vh.Santa.Y++
	}
}

func (vh *VisitedHouses) moveRoboSanta(dir string) {
	if dir == "<" {
		vh.RoboSanta.X--
	} else if dir == ">" {
		vh.RoboSanta.X++
	} else if dir == "^" {
		vh.RoboSanta.Y--
	} else if dir == "v" {
		vh.RoboSanta.Y++
	}
}

func (vh *VisitedHouses) alreadyVisitedHouse(santa utils.Co) bool {
	return vh.Map[santa]
}

func (vh *VisitedHouses) countUniqueHousesVisited(input string, part1 bool) int {
	uniqueHousesVisited := 1
	for i, dir := range input {
		strDir := string(dir)
		if part1 || i%2 == 0 {
			vh.moveSanta(strDir)
			if !vh.alreadyVisitedHouse(vh.Santa) {
				uniqueHousesVisited++
				vh.Map[vh.Santa] = true
			}
		} else {
			vh.moveRoboSanta(strDir)
			if !vh.alreadyVisitedHouse(vh.RoboSanta) {
				uniqueHousesVisited++
				vh.Map[vh.RoboSanta] = true
			}
		}
	}
	return uniqueHousesVisited
}

func main() {
	input := utils.ReadFile()[0]
	visitedHouses := createVisitedHouses()
	fmt.Println("Part 1:", visitedHouses.countUniqueHousesVisited(input, true))
	visitedHouses = createVisitedHouses()
	fmt.Println("Part 2:", visitedHouses.countUniqueHousesVisited(input, false))
}
