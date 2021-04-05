package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Addresses map[int]int

var reMask = regexp.MustCompile(`\w+`)

func getVal(val int, mask string) (int, error) {
	and, err := strconv.ParseUint(strings.ReplaceAll(mask, "X", "1"), 2, 0)
	if err != nil {
		return 0, err
	}
	// If the previous ParseUint produced a nil error then so will this one
	or, _ := strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 0)
	return val&int(and) | int(or), nil
}

func findSolutions(entries []string) (int, int, error) {
	part1Total, part1Addresses := 0, Addresses{}
	part2Total, part2Addresses := 0, Addresses{}
	var mask string

	for _, entry := range entries {
		matches := reMask.FindAllString(entry, -1)
		if len(matches) == 2 {
			mask = matches[1]
		} else {
			add, err := strconv.Atoi(matches[1])
			if err != nil {
				return 0, 0, err
			}
			val, err := strconv.Atoi(matches[2])
			if err != nil {
				return 0, 0, err
			}
			for i, x := 0, strings.Count(mask, "X"); i < 1<<x; i++ {
				mask := strings.NewReplacer("X", "x", "0", "X").Replace(mask)
				for _, r := range fmt.Sprintf("%0*b", x, i) {
					mask = strings.Replace(mask, "x", string(r), 1)
				}
				add, err = getVal(add, mask)
				if err != nil {
					return 0, 0, err
				}
				part2Total += val - part2Addresses[add]
				part2Addresses[add] = val
			}

			// Any error will have been returned in the previous for loop, so one shouldn't be
			// returned here
			val, _ = getVal(val, mask)
			part1Total += val - part1Addresses[add]
			part1Addresses[add] = val
		}
	}

	return part1Total, part2Total, nil
}

func main() {
	entries := helpers.ReadFile()
	part1Sol, part2Sol, err := findSolutions(entries)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1Sol)
	fmt.Println("Part 2:", part2Sol)
}
