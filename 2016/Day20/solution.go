package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/regex"
	"fmt"
	"sort"
	"strconv"
)

func composeBounds(input []string, upperBound int) ([]int, []int) {
	lowers := make([]int, len(input)+1)
	uppers := make([]int, len(input)+1)
	for i, line := range input {
		matches := regex.MatchNums.FindAllString(line, -1)
		lower, _ := strconv.Atoi(matches[0])
		lowers[i] = lower
		upper, _ := strconv.Atoi(matches[1])
		uppers[i] = upper
	}
	// The largest upper may be less than the max possible IP, so add an extra lower upper so we
	// include IPs up to this.
	lowers[len(lowers)-1] = upperBound + 1
	uppers[len(uppers)-1] = upperBound + 1

	sort.Ints(lowers)
	sort.Ints(uppers)
	return lowers, uppers
}

func findAllowedIPs(lowers, uppers []int) (int, int) {
	lowestAllowed := lowers[0]
	allowed := []int{}
	for i := range lowers {
		if lowers[i] > lowestAllowed {
			for j := uppers[i-1] + 1; j < lowers[i]; j++ {
				allowed = append(allowed, j)
			}
		} else {
			lowestAllowed = uppers[i] + 1
		}
	}
	return allowed[0], len(allowed)
}

func main() {
	input := file.Read()
	lowers, uppers := composeBounds(input, 4294967295)
	part1, part2 := findAllowedIPs(lowers, uppers)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
