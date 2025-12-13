package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/maths"
	"fmt"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`([L|R])(\d+)`)

func findSolutions(input []string) (part1, part2 int, err error) {
	var dial = 50
	for _, line := range input {
		matches := re.FindAllStringSubmatch(line, -1)
		if len(matches) != 1 {
			return part1, part2, fmt.Errorf("error parsing line: %v", line)
		}
		match := matches[0]
		if len(match) != 3 {
			return part1, part2, fmt.Errorf("error parsing line: %v", line)
		}
		// Can ignore error due to regex matching - we know it'll convert
		val, _ := strconv.Atoi(match[2])

		switch match[1] {
		case "R":
			end := dial + val
			i := dial + 1
			for {
				if maths.Modulo(i, 100) == 0 {
					part2++
					i += 99
				}
				i++
				if i > end {
					break
				}
			}
			dial = maths.Modulo(end, 100)
		case "L":
			end := dial - val
			i := dial - 1
			for {
				if maths.Modulo(i, 100) == 0 {
					part2++
					i -= 99
				}
				i--
				if i < end {
					break
				}
			}
			dial = maths.Modulo(end, 100)
		default:
			return part1, part2, fmt.Errorf("no valid instruction: %v", line)
		}

		if dial == 0 {
			part1++
		}
	}
	return part1, part2, nil
}

func main() {
	input := file.Read()
	part1, part2, err := findSolutions(input)
	if err != nil {
		fmt.Printf("Error finding solutions: %v\n", err)
	}
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}
