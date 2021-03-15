package main

import (
	"Advent-of-Code"
	"fmt"
	"regexp"
)

type relationships map[string]int

var people = make(map[string]relationships)

var allPeople []string

func calculateMaxHappiness() (maxHappiness int) {
	for _, perm := range helpers.Permutations(allPeople) {
		happiness := 0
		for i := 0; i < len(perm)-1; i++ {
			happiness += people[perm[i]][perm[i+1]]
			happiness += people[perm[i+1]][perm[i]]
		}
		happiness += people[perm[len(perm)-1]][perm[0]]
		happiness += people[perm[0]][perm[len(perm)-1]]
		if happiness > maxHappiness {
			maxHappiness = happiness
		}
	}
	return maxHappiness
}

func main() {
	lines := helpers.ReadFile()
	nameRe := regexp.MustCompile("[A-Z][a-z]+")
	numberRe := regexp.MustCompile("\\d+")
	gainRe := regexp.MustCompile("gain")
	for _, l := range lines {
		persons := nameRe.FindAllString(l, -1)
		if !helpers.StringInSlice(persons[0], allPeople) {
			allPeople = append(allPeople, persons[0])
		}
		happiness := helpers.StringToInt(numberRe.FindAllString(l, -1)[0])
		if sign := gainRe.MatchString(l); !sign {
			happiness = -happiness
		}
		if rel, ok := people[persons[0]]; !ok {
			people[persons[0]] = relationships{persons[1]: happiness}
		} else {
			rel[persons[1]] = happiness
			people[persons[0]] = rel
		}
	}

	fmt.Println("Part 1:", calculateMaxHappiness())
	people["Me"] = relationships{}
	for _, p := range allPeople {
		people["Me"][p] = 0
		people[p]["Me"] = 0
	}
	allPeople = append(allPeople, "Me")
	fmt.Println("Part 2:", calculateMaxHappiness())
}
