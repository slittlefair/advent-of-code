package main

import (
	"Advent-of-Code/file"
	"fmt"
	"regexp"
	"strconv"
)

type DataStruct struct {
	min, max int
}

var reRange = regexp.MustCompile(`\d+-\d+`)
var reNum = regexp.MustCompile(`\d+`)

func parseInput(input []string) ([]DataStruct, error) {
	ds := []DataStruct{}
	for _, line := range input {
		matches := reRange.FindAllString(line, -1)
		for _, m := range matches {
			nums := reNum.FindAllString(m, -1)
			if len(nums) != 2 {
				return nil, fmt.Errorf("malformed nums: %v", m)
			}
			n0, _ := strconv.Atoi(nums[0])
			n1, _ := strconv.Atoi(nums[1])
			ds = append(ds, DataStruct{min: n0, max: n1})
		}
	}
	return ds, nil
}

func evaluateNum(s string) (int, int) {
	l := len(s)
	// Go through each number leq half the length, and see whether the number is made up of repeats
	// of that substring. Start at half length and work downwards so that we find longer instances
	// first, e.g. for 222222 we find repeats of 222 before 22 or 2, which we want for part 1.
	for j := l / 2; j >= 1; j-- {
		m := map[string]int{}
		// If the substring won't go evenly then continue
		if l%j != 0 {
			continue
		}
		// Cycle over the number getting substrings of the various lengths, adding them to a map. We
		// want to end up with a map of length 1, in which case it's the same substring repeated.
		// If the value in the map is two then it's made up of the same substring twice, in which
		// case it satisifies part 1 as well as part 2.
		for i := 0; i <= l-j; i += j {
			m[s[i:i+j]]++
		}
		if len(m) == 1 {
			n, _ := strconv.Atoi(s)
			for _, v := range m {
				if v == 2 {
					return n, n
				}
			}
			return 0, n
		}
	}
	return 0, 0
}

func numInvalid(ds DataStruct) (int, int) {
	var part1, part2 int
	for i := ds.min; i <= ds.max; i++ {
		s := strconv.Itoa(i)
		p1, p2 := evaluateNum(s)
		part1 += p1
		part2 += p2
	}
	return part1, part2
}

func findSolutions(input []string) (int, int, error) {
	part1 := 0
	part2 := 0
	ds, err := parseInput(input)
	if err != nil {
		return -1, -1, err
	}
	for _, d := range ds {
		p1, p2 := numInvalid(d)
		part1 += p1
		part2 += p2
	}
	return part1, part2, nil
}

func main() {
	input := file.Read()
	part1, part2, err := findSolutions(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}
