package main

import (
	helpers "Advent-of-Code"
	"errors"
	"fmt"
)

func part1(entries []int) (int, error) {
	otherHalves := make(map[int]int)

	for _, entry := range entries {
		if val, ok := otherHalves[entry]; ok {
			return val * entry, nil
		}
		otherHalves[2020-entry] = entry
	}

	return 0, errors.New("part1: could not find entries that sum to 2020")
}

// Not as efficient but the input txt.file isn't huge so it'll do for now
// TODO improve this
func part2(entries []int) (int, error) {
	for _, i := range entries {
		for _, j := range entries {
			for _, k := range entries {
				if i+j+k == 2020 {
					return i * j * k, nil
				}
			}
		}
	}
	return 0, errors.New(("part2: could not find entries that sum to 2020"))
}

func main() {
	entries := helpers.ReadFileAsInts()
	answer1, err := part1(entries)
	if err != nil {
		return
	}

	fmt.Println("Part 1:", answer1)

	answer2, err := part2(entries)
	if err != nil {
		return
	}
	fmt.Println("Part 2:", answer2)
}
