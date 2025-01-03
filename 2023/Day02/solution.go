package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/regex"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Cube struct {
	re  *regexp.Regexp
	max int
}

var colours = []*Cube{
	{
		re:  regexp.MustCompile(`\d+ red`),
		max: 12,
	},
	{
		re:  regexp.MustCompile(`\d+ green`),
		max: 13,
	},
	{
		re:  regexp.MustCompile(`\d+ blue`),
		max: 14,
	},
}

func findSolutions(input []string) (int, int) {
	part1 := 0
	part2 := 0

	for i, line := range input {
		validGame, power := handleLine(line)
		if validGame {
			part1 += i + 1
		}
		part2 += power
	}
	return part1, part2
}

func handleLine(line string) (bool, int) {
	validGame := true
	split := strings.Split(line, "; ")
	powerMap := map[int]int{}
	for _, s := range split {
		for i, col := range colours {
			if match := col.re.FindString(s); match != "" {
				num := regex.MatchNums.FindString(match)
				v, _ := strconv.Atoi(num)
				if v > col.max {
					validGame = false
				}
				if v > powerMap[i] {
					powerMap[i] = v
				}
			}
		}
	}
	power := 1
	for _, pwr := range powerMap {
		power *= pwr
	}
	return validGame, power
}

func main() {
	input := file.Read()
	part1, part2 := findSolutions(input)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
