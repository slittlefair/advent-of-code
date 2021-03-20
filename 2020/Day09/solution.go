package main

import (
	helpers "Advent-of-Code"
	"errors"
	"fmt"
)

func cyclePrevNumbers(entries []int, preambleLength int, i int) (bool, int, int) {
	for j := i - preambleLength; j < i; j++ {
		for k := i - preambleLength; k < i; k++ {
			if j != k && entries[j]+entries[k] == entries[i] {
				return false, -1, -1
			}
		}
	}
	return true, i, entries[i]
}

func part1(entries []int, preambleLength int) (int, int, error) {
	for i := preambleLength; i < len(entries); i++ {
		solved, solIndex, solValue := cyclePrevNumbers(entries, preambleLength, i)
		if solved {
			return solIndex, solValue, nil
		}
	}
	return -1, -1, errors.New("could not find solution to part 1")
}

func main() {
	entries := helpers.ReadFileAsInts()
	preambleLength := 25

	_, part1SolValue, err := part1(entries, preambleLength)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Part 1:", part1SolValue)
}
