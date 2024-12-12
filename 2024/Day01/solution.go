package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/maths"
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

var re = regexp.MustCompile(`\d+`)

func createSortedSlicesFromInput(input []string) ([]int, []int, error) {
	slice1 := make([]int, len(input))
	slice2 := make([]int, len(input))
	for i, line := range input {
		matches := re.FindAllString(line, -1)
		if l := len(matches); l != 2 {
			return nil, nil, fmt.Errorf("malformed input line, expected 2 numbers, got %d: %v", l, line)
		}

		// We can ignore errors from strconv.Atoi since we know they're ints due to the regex match
		m1, _ := strconv.Atoi(matches[0])
		slice1[i] = m1
		m2, _ := strconv.Atoi(matches[1])
		slice2[i] = m2
	}

	// Sort the slices before returning
	slices.Sort(slice1)
	slices.Sort(slice2)

	return slice1, slice2, nil
}

func calculateDiffInSlices(slice1, slice2 []int) (int, error) {
	var diff int

	if !slices.IsSorted(slice1) {
		return diff, fmt.Errorf("expect slice1 is already sorted")
	}

	if !slices.IsSorted(slice2) {
		return diff, fmt.Errorf("Expect slice2 is already sorted")
	}

	if len(slice1) != len(slice2) {
		return diff, fmt.Errorf("Expect slices to be of the same length: %d vs. %d", len(slice1), len(slice2))
	}

	for i := range slice1 {
		diff += maths.Abs(slice1[i] - slice2[i])
	}

	return diff, nil
}

func calculateSimilarityScore(slice1, slice2 []int) int {
	var similarityScore int
	frequencyMap := map[int]int{}
	for _, v := range slice2 {
		frequencyMap[v]++
	}

	for _, v := range slice1 {
		similarityScore += frequencyMap[v] * v
	}

	return similarityScore
}

func findSolutions(input []string) (int, int, error) {
	slice1, slice2, err := createSortedSlicesFromInput(input)
	if err != nil {
		return 0, 0, fmt.Errorf("error getting sorted slices from input: %v\n", err)
	}

	part1, err := calculateDiffInSlices(slice1, slice2)
	if err != nil {
		return 0, 0, fmt.Errorf("error calculating diff in slices: %v\n", err)
	}

	part2 := calculateSimilarityScore(slice1, slice2)

	return part1, part2, err
}

func main() {
	input := file.Read()

	part1, part2, err := findSolutions(input)
	if err != nil {
		fmt.Printf("finding solutions: %v\n", err)
	}

	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
