package main

import (
	"Advent-of-Code/file"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

func composeBounds(input []string, upperBound int) ([]int, []int) {
	lowers := []int{}
	uppers := []int{}
	re := regexp.MustCompile(`\d+`)
	for _, line := range input {
		matches := re.FindAllString(line, -1)
		lower, _ := strconv.Atoi(matches[0])
		lowers = append(lowers, lower)
		upper, _ := strconv.Atoi(matches[1])
		uppers = append(uppers, upper)
	}
	sort.Ints(lowers)
	sort.Ints(uppers)
	// The largest upper may be less than the max possible IP, so add ean extra lower upper so we
	// include IPs up to this.
	lowers = append(lowers, upperBound+1)
	uppers = append(uppers, upperBound+1)
	return lowers, uppers
}

func findAllowedIPs(lowers, uppers []int) (int, int) {
	lowestAllowed := lowers[0]
	allowed := []int{}
	for i := 0; i < len(lowers); i++ {
		if lowers[i] > lowestAllowed {
			for j := uppers[i-1] + 1; j < lowers[i]; j++ {
				allowed = append(allowed, j)
			}
		} else {
			lowestAllowed = uppers[i] + 1
		}
	}
	fmt.Println(allowed)
	return allowed[0], len(allowed)
}

func main() {
	input := file.Read()
	lowers, uppers := composeBounds(input, 4294967295)
	part1, part2 := findAllowedIPs(lowers, uppers)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
