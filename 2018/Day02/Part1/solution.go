package main

import (
	"Advent-of-Code"
	"fmt"
)

func main() {
	var twoOfLetterSum, threeOfLetterSum int
	lines := helpers.ReadFile()
	for _, id := range lines {
		m := make(map[string]int)
		var twoOfLetter, threeOfLetter = false, false

		// Loop through the id and count how many times each letter appears
		for _, l := range id {
			l := string(l)
			if _, ok := m[l]; ok {
				m[l]++
			} else {
				m[l] = 1
			}
		}

		// Loop through the map of frequencies and see if the id contributes
		for _, key := range m {
			switch key {
			case 2:
				twoOfLetter = true
			case 3:
				threeOfLetter = true
			}
		}

		if twoOfLetter {
			twoOfLetterSum++
		}
		if threeOfLetter {
			threeOfLetterSum++
		}
	}
	fmt.Print(twoOfLetterSum * threeOfLetterSum)
}
