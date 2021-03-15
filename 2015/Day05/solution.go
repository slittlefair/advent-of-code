package main

import (
	"Advent-of-Code"
	"fmt"
)

func part1(strings []string) (numNiceStrings int) {
	vowels := [5]string{"a", "e", "i", "o", "u"}
	naughtyStrings := [4]string{"ab", "cd", "pq", "xy"}
	for _, s := range strings {
		naughty := false
		numVowels := 0
		doubleLetter := false
		for i, l := range s {
			letter := string(l)
			for _, j := range vowels {
				if j == letter {
					numVowels++
				}
			}
			if i < len(s)-1 {
				if letter == string(s[i+1]) {
					doubleLetter = true
				}
				for _, ns := range naughtyStrings {
					if letter+string(s[i+1]) == ns {
						naughty = true
					}
				}
			}
		}
		if numVowels < 3 || !doubleLetter {
			naughty = true
		}
		if !naughty {
			numNiceStrings++
		}
	}
	return numNiceStrings
}

func part2(strings []string) (numNiceStrings int) {
	for _, s := range strings {
		doubles := make(map[string]int)
		hasDoubles := false
		gapRepeats := false
		for i := range s {
			if i < len(s)-2 {
				if s[i] == s[i+2] {
					gapRepeats = true
				}
			}
		}
		for i := 0; i < len(s); i++ {
			if i <= len(s)-2 {
				if freq, ok := doubles[string(s[i])+string(s[i+1])]; !ok {
					doubles[string(s[i])+string(s[i+1])] = 1
				} else {
					doubles[string(s[i])+string(s[i+1])] = freq + 1
				}
				if i < len(s)-2 && s[i] == s[i+1] && s[i] == s[i+2] {
					i++
				}
			}

		}
		for _, freq := range doubles {
			if freq > 1 {
				hasDoubles = true
			}
		}
		if gapRepeats && hasDoubles {
			numNiceStrings++
		}
	}

	return numNiceStrings
}

func main() {
	strings := helpers.ReadFile()
	fmt.Println("Part 1:", part1(strings))
	fmt.Println("Part 2:", part2(strings))
}
