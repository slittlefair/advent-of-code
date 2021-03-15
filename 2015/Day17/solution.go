package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"strconv"
)

var litres = 150

var idToContainer = make(map[string]int)
var successfulCombinations [][]string
var allIDs []string

func main() {
	containers := helpers.ReadFile()
	for i, c := range helpers.StringSliceToIntSlice(containers) {
		idToContainer[strconv.Itoa(i)] = c
		allIDs = append(allIDs, strconv.Itoa(i))
	}
	allPerms := helpers.Permutations(allIDs)
	fmt.Println(len(allPerms))
	for _, perm := range allPerms {
		total := 0
		idSlice := []string{}
		for _, c := range perm {
			total += idToContainer[c]
			idSlice = append(idSlice, c)
			if total == litres {
				alreadyPicked := false
				for _, comb := range successfulCombinations {
					if helpers.StringSlicesEqual(idSlice, comb) {
						alreadyPicked = true
					}
				}
				if !alreadyPicked {
					successfulCombinations = append(successfulCombinations, idSlice)
				}
			} else if total > litres {
				continue
			}
		}
	}
	fmt.Println("Part 1:", len(successfulCombinations))
}
