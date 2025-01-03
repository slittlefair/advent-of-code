package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/regex"
	"fmt"
	"regexp"
	"strconv"
)

var reMulPart1 = regexp.MustCompile(`mul\(\d+,\d+\)`)
var reMulPart2 = regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

func getMultiplication(mul string) int {
	// mul is already regex matched to ensure we have two numbers, so we don't need to validate
	// or handle any errors
	nums := regex.MatchNums.FindAllString(mul, 2)
	n1, _ := strconv.Atoi(nums[0])
	n2, _ := strconv.Atoi(nums[1])
	return n1 * n2
}

func findSolutions(input []string) (int, int) {
	var part1, part2 int
	mulEnabled := true
	for _, line := range input {
		matches := reMulPart1.FindAllString(line, -1)
		for _, mul := range matches {
			part1 += getMultiplication(mul)
		}

		matches = reMulPart2.FindAllString(line, -1)
		for _, m := range matches {
			if m == "do()" {
				mulEnabled = true
				continue
			}
			if m == "don't()" {
				mulEnabled = false
				continue
			}
			if !mulEnabled {
				continue
			}
			part2 += getMultiplication(m)
		}
	}
	return part1, part2
}

func main() {
	input := file.Read()
	part1, part2 := findSolutions(input)
	fmt.Printf("Part1: %d\n", part1)
	fmt.Printf("Part2: %d\n", part2)
}
