package main

import (
	"Advent-of-Code"
	"fmt"
	"regexp"
)

type leg struct {
	from     string
	to       string
	distance int
}

var allLegs []leg

var allDestinations []string

func routeDistance(perm []string) (totalDistance int) {
	for i := 0; i < len(perm)-1; i++ {
		for _, lg := range allLegs {
			if (lg.from == perm[i] && lg.to == perm[i+1]) || (lg.to == perm[i] && lg.from == perm[i+1]) {
				totalDistance += lg.distance
			}
		}
	}
	return totalDistance
}

func main() {
	lines := helpers.ReadFile()
	wordRe := regexp.MustCompile("\\w+")
	for _, l := range lines {
		words := wordRe.FindAllString(l, -1)
		allLegs = append(allLegs, leg{
			to:       words[0],
			from:     words[2],
			distance: helpers.StringToInt(words[3]),
		})
		inArray1 := false
		inArray2 := false
		for _, dest := range allDestinations {
			if words[0] == dest {
				inArray1 = true
			}
			if words[2] == dest {
				inArray2 = true
			}
		}
		if !inArray1 {
			allDestinations = append(allDestinations, words[0])
		}
		if !inArray2 {
			allDestinations = append(allDestinations, words[2])
		}
	}
	allPerms := helpers.Permutations(allDestinations)
	shortestDistance := 1000000000
	longestDistance := 0
	for _, perm := range allPerms {
		dist := routeDistance(perm)
		if dist < shortestDistance {
			shortestDistance = dist
		}
		if dist > longestDistance {
			longestDistance = dist
		}
	}
	fmt.Println("Part 1:", shortestDistance)
	fmt.Println("Part 2:", longestDistance)
}
