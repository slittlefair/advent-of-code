package main

import (
	"Advent-of-Code/file"
	"fmt"
	"regexp"
	"strconv"
)

type Elf struct {
	low  int
	high int
}

func findContainedAssignments(input []string) (int, int, error) {
	containedPairs := 0
	overlappingPairs := 0
	re := regexp.MustCompile(`\d+`)
	for _, line := range input {
		matches := re.FindAllString(line, -1)
		if len(matches) != 4 {
			return -1, -1, fmt.Errorf("malformed input %s", line)
		}
		nums := []int{}
		for _, m := range matches {
			// We already know all matches can be convrted to ints due to regex matching,
			// so Atoi cannot return an error.
			val, _ := strconv.Atoi(m)
			nums = append(nums, val)
		}
		e1 := Elf{
			low:  nums[0],
			high: nums[1],
		}
		e2 := Elf{
			low:  nums[2],
			high: nums[3],
		}
		if (e1.low <= e2.low && e1.high >= e2.high) ||
			(e2.low <= e1.low && e2.high >= e1.high) {
			containedPairs++
		}
		if e1.low <= e2.high && e1.high >= e2.low {
			overlappingPairs++
		}
	}
	return containedPairs, overlappingPairs, nil
}

func main() {
	input := file.Read()
	containedPairs, overlappingPairs, err := findContainedAssignments(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", containedPairs)
	fmt.Println("Part 2:", overlappingPairs)
}
