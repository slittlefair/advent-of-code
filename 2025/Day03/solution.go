package main

import (
	"Advent-of-Code/file"
	"fmt"
	"regexp"
	"strconv"
)

type Banks [][]int

var re = regexp.MustCompile(`\d`)

func parseInput(input []string) Banks {
	banks := make(Banks, len(input))
	for i, line := range input {
		matches := re.FindAllString(line, -1)
		bank := make([]int, len(matches))
		for j, m := range matches {
			n, _ := strconv.Atoi(m)
			bank[j] = n
		}
		banks[i] = bank
	}
	return banks
}

func evaluateLine(bank []int, digits int) int {
	num := 0
	maxIdx := -1
	for i := digits - 1; i >= 0; i-- {
		maxN := 0
		for j := maxIdx + 1; j < len(bank)-i; j++ {
			n := bank[j]
			if n > maxN {
				maxN = n
				maxIdx = j
			}
		}
		num = num*10 + maxN
	}
	return num
}

func findSolutions(input []string) (int, int) {
	part1 := 0
	part2 := 0
	banks := parseInput(input)
	for _, bank := range banks {
		part1 += evaluateLine(bank, 2)
		part2 += evaluateLine(bank, 12)
	}
	return part1, part2
}

func main() {
	input := file.Read()
	part1, part2 := findSolutions(input)
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}
