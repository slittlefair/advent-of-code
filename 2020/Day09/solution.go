package main

import (
	helpers "Advent-of-Code"
	"fmt"
)

func cyclePrevNumbers(entries []int, preambleLength int, i int) bool {
	for j := i - preambleLength; j < i; j++ {
		for k := i - preambleLength; k < i; k++ {
			if j != k && entries[j]+entries[k] == entries[i] {
				return true
			}
		}
	}
	return false
}

func main() {
	entries := helpers.ReadFileAsInts()

	preambleLength := 25

	for i := preambleLength; i < len(entries); i++ {
		moveOn := cyclePrevNumbers(entries, preambleLength, i)
		if !moveOn {
			fmt.Println("Part 1:", entries[i])
			break
		}
	}
}
