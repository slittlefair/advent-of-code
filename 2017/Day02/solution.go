package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/maths"
	"Advent-of-Code/regex"
	"fmt"
	"math"
	"strconv"
)

type Spreadsheet [][]int

func parseInput(input []string) Spreadsheet {
	ss := make(Spreadsheet, len(input))
	for i, line := range input {
		nums := regex.MatchNums.FindAllString(line, -1)
		row := make([]int, len(nums))
		for j, n := range nums {
			// We already match using regex so we know all numbers can be converted to ints
			c, _ := strconv.Atoi(n)
			row[j] = c
		}
		ss[i] = row
	}
	return ss
}

func dividesEvenly(x, y float64) bool {
	return math.Remainder(x, y) == 0 || math.Remainder(y, x) == 0
}

func findSolutions(ss Spreadsheet) (int, int) {
	part1, part2 := 0, 0
	for _, row := range ss {
		maximum := 0
		minimum := maths.Infinity
		foundDivisiblePair := false
		for i, x := range row {
			maximum = maths.Max(x, maximum)
			minimum = maths.Min(x, minimum)
			for j, y := range row {
				if i != j && !foundDivisiblePair && dividesEvenly(float64(x), float64(y)) {
					part2 += maths.Max(x, y) / maths.Min(x, y)
					// Once we've found the divisible pair in the row we can stop checking it,
					// othewise we'll find the same pair, reversed, and add this to the checksum
					foundDivisiblePair = true
					break
				}
			}
		}
		part1 += maximum - minimum
	}
	return part1, part2
}

func main() {
	input := file.Read()
	spreadsheet := parseInput(input)
	part1, part2 := findSolutions(spreadsheet)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
