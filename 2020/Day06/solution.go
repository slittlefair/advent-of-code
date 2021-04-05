package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
)

func getSolution(entries []string) (int, int) {
	re := regexp.MustCompile(`\w`)
	groupPositiveResponses := map[string]int{}
	numAnyoneYes := 0
	numGroupMembers := 0
	numEveryoneYes := 0
	for i, entry := range entries {
		matches := re.FindAllString(entry, -1)
		for _, match := range matches {
			if val, ok := groupPositiveResponses[match]; !ok {
				groupPositiveResponses[match] = 1
			} else {
				groupPositiveResponses[match] = val + 1
			}
		}
		if len(matches) != 0 {
			numGroupMembers++
		}
		if len(matches) == 0 || i == len(entries)-1 {
			numAnyoneYes += len(groupPositiveResponses)
			for k := range groupPositiveResponses {
				if groupPositiveResponses[k] == numGroupMembers {
					numEveryoneYes++
				}
			}
			groupPositiveResponses = map[string]int{}
			numGroupMembers = 0
		}
	}
	return numAnyoneYes, numEveryoneYes
}

func main() {
	entries := helpers.ReadFile()

	part1Sol, part2Sol := getSolution(entries)

	fmt.Println("Part 1:", part1Sol)
	fmt.Println("Part 2:", part2Sol)
}
