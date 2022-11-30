package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/maths"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Spreadsheet [][]int

func parseInput(input []string) (Spreadsheet, error) {
	ss := Spreadsheet{}
	re := regexp.MustCompile(`\d+`)
	for _, line := range input {
		row := []int{}
		nums := re.FindAllString(line, -1)
		for _, n := range nums {
			i, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			row = append(row, i)
		}
		ss = append(ss, row)
	}
	return ss, nil
}

func dividesEvenly(x, y float64) bool {
	return math.Remainder(x, y) == 0 || math.Remainder(y, x) == 0
}

func findSolutions(ss Spreadsheet) (int, int) {
	part1, part2 := 0, 0
	for _, row := range ss {
		max := 0
		min := maths.Infinity
		foundDivisiblePair := false
		for i, x := range row {
			max = maths.Max(x, max)
			min = maths.Min(x, min)
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
		part1 += max - min
	}
	return part1, part2
}

func main() {
	input := file.Read()
	spreadsheet, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	part1, part2 := findSolutions(spreadsheet)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
