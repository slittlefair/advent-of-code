package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"Advent-of-Code/regex"
	"fmt"
	"strconv"
	"strings"
)

type ClawMachine struct {
	ButtonA, ButtonB, Prize graph.Co
}

// For a given line of input and prefix, return the two numbers, as ints, contained in that line.
// If the line doesn't contain the correct prefix, or two numbers, then an error is returned instead
func getNumsFromLine(line, prefix string) ([]int, error) {
	if !strings.HasPrefix(line, prefix) {
		return nil, fmt.Errorf(`malformed input, expected "%s", got %v`, prefix, line)
	}
	matches := regex.MatchNums.FindAllString(line, -1)
	if len(matches) != 2 {
		return nil, fmt.Errorf(`malformed input, expected 2 nums: %v`, line)
	}
	nums := make([]int, 2)
	for i, m := range matches {
		// We don't have to handle an error a we know the conversion will work since we matched
		// using regex
		f, _ := strconv.Atoi(m)
		nums[i] = f
	}
	return nums, nil
}

func parseInput(input []string) ([]ClawMachine, error) {
	machines := []ClawMachine{}
	cm := ClawMachine{}
	for i, line := range input {
		// Lines are one of four types, each need handling in a different way
		switch i % 4 {
		case 0:
			// Line should start "Button A" and contain two numbers for x and y respectively
			nums, err := getNumsFromLine(line, "Button A:")
			if err != nil {
				return nil, err
			}
			cm.ButtonA = graph.Co{X: nums[0], Y: nums[1]}
		case 1:
			// Line should start "Button B" and contain two numbers for x and y respectively
			nums, err := getNumsFromLine(line, "Button B:")
			if err != nil {
				return nil, err
			}
			cm.ButtonB = graph.Co{X: nums[0], Y: nums[1]}
		case 2:
			// Line should start "Prize" and contain two numbers for x and y respectively
			nums, err := getNumsFromLine(line, "Prize:")
			if err != nil {
				return nil, err
			}
			cm.Prize = graph.Co{X: nums[0], Y: nums[1]}
			// Once the Prize has been added to the claw machine it's complete, so it can be added
			// to our list of all machines
			machines = append(machines, cm)
		case 3:
			// Check the next line is blank, indicating we're moving to a new claw machine. If the
			// line isn't blank then we'd potentially be skipping over important information, so
			// return an error
			if line != "" {
				return nil, fmt.Errorf(`malformed input, expected "", got %v`, line)
			}
			cm = ClawMachine{}
		}
	}
	return machines, nil
}

// Simultaneous equations means we can find values for the number of times we need to press A (a)
// and B (b) with mathematical sums based on the values of the buttons and claws.
func (cm ClawMachine) findTokensForWin() int {
	b := float64(
		(cm.ButtonA.Y*cm.Prize.X)-(cm.Prize.Y*cm.ButtonA.X),
	) / float64(
		(cm.ButtonB.X*cm.ButtonA.Y)-(cm.ButtonB.Y*cm.ButtonA.X),
	)
	// Check that b is an int, if not then a solution doesn't exist
	if b != float64(int(b)) {
		return 0
	}
	a := float64(cm.Prize.X-(int(b)*cm.ButtonB.X)) / float64(cm.ButtonA.X)
	// Check that a is an int, if not then a solution doesn't exist
	if a != float64(int(a)) {
		return 0
	}
	return int(a)*3 + int(b)
}

func findSolutions(input []string) (int, int, error) {
	var part1, part2 int
	machines, err := parseInput(input)
	for _, m := range machines {
		part1 += m.findTokensForWin()
		m.Prize.X += 10000000000000
		m.Prize.Y += 10000000000000
		part2 += m.findTokensForWin()
	}
	if err != nil {
		return part1, part2, err
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
