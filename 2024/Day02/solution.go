package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/regex"
	"Advent-of-Code/slice"
	"fmt"
	"strconv"
)

func createReportsFromInput(input []string) [][]int {
	reports := [][]int{}
	for _, line := range input {
		matches := regex.MatchNums.FindAllString(line, -1)
		report := []int{}
		for _, m := range matches {
			// We can ignore the error as we know each one can be converted to an int due to the regex matching
			v, _ := strconv.Atoi(m)
			report = append(report, v)
		}
		reports = append(reports, report)
	}
	return reports
}

func reportIsSafe(report []int) (bool, int) {
	asc := report[0] < report[1]
	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]
		if diff == 0 || diff < -3 || diff > 3 {
			return false, i
		}
		if (diff*-1 < 0) == asc {
			return false, i
		}
	}
	return true, -1
}

func findSolutions(input []string) (int, int) {
	reports := createReportsFromInput(input)
	var part1, part2 int
	for _, report := range reports {
		isSafePart1, unsafeIndex := reportIsSafe(report)
		if isSafePart1 {
			part1++
			part2++
		} else {
			// Loop over the elements around the unsafe one and see if we can make a report safe by
			// removing any of them. We need to remove the one that caused the report to be made
			// unsafe, as well as the levels before and after since any of these could have caused
			// the unsafe element to be this particular level.
			for j := unsafeIndex - 1; j <= unsafeIndex+1; j++ {
				// If the unsafe level is the first element we can't remove any elements before it
				if j < 0 {
					continue
				}
				if isSafe, _ := reportIsSafe(slice.Remove(report, j)); isSafe {
					// If we're now safe having removed any element, increase the count
					part2++
					break
				}
			}
		}
	}
	return part1, part2
}

func main() {
	input := file.Read()
	part1, part2 := findSolutions(input)

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}
