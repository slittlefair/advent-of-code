package main

import (
	"Advent-of-Code/file"
	"fmt"
)

type Crates struct {
	arrangements [][]string
	instructions [][]int
}

func parseInput(input []string) (Crates, error) {
	crates := Crates{
		arrangements: [][]string{},
		instructions: [][]int{},
	}

	// Run through input to find which stack is the largest and how many stacks there are
	var largestStack = 0
	var numStacks = 0
	for i, line := range input {
		if line[1] == '1' {
			largestStack = i
			// assume there are more than 1 columns, in which case this equation wouldn't work
			numStacks = (len(line) + 1) / 4
			break
		}
	}

	// Initialise arrangement so we can append later
	crates.arrangements = make([][]string, numStacks)

	// Run through instructions
	for i := largestStack + 2; i < len(input); i++ {
		var a, b, c int
		_, err := fmt.Sscanf(input[i], "move %d from %d to %d", &a, &b, &c)
		if err != nil {
			return crates, err
		}
		// Our stacks are 0-indexed but referred to by 1-index in input, so subtract 1 from "from" and "to"
		crates.instructions = append(crates.instructions, []int{a, b - 1, c - 1})
	}

	// Run through stacks from bottom to top so slices run from bottom to top
	for i := numStacks - 1; i >= 0; i-- {
		line := input[i]
		stackIndex := 0
		// Crate contents are every 4 characters starting at index 1 in line
		for j := 1; j < len(line); j = j + 4 {
			// If a crate exists (character is not " ") then add it to the stack
			if str := string(line[j]); str != " " {
				crates.arrangements[stackIndex] = append(crates.arrangements[stackIndex], str)
			}
			stackIndex++
		}
	}

	return crates, nil
}

// Follow instructions for part 1
func (c *Crates) runCrateMover9000() {
	for _, inst := range c.instructions {
		fromColumn := c.arrangements[inst[1]]
		toColumn := c.arrangements[inst[2]]
		// For each item in the top inst[0] crates, add them to the new stack and remove them
		// from the old stack one at a time
		for i := 0; i < inst[0]; i++ {
			toColumn = append(toColumn, fromColumn[len(fromColumn)-1])
			fromColumn = fromColumn[:len(fromColumn)-1]
		}
		c.arrangements[inst[1]] = fromColumn
		c.arrangements[inst[2]] = toColumn
	}
}

// Follow instructions for part 2
func (c *Crates) runCrateMover9001() {
	for _, inst := range c.instructions {
		fromColumn := c.arrangements[inst[1]]
		toColumn := c.arrangements[inst[2]]
		// Take the whole lot if items to move in one go, keeping the order
		toColumn = append(toColumn, fromColumn[len(fromColumn)-inst[0]:]...)
		fromColumn = fromColumn[:len(fromColumn)-inst[0]]
		c.arrangements[inst[1]] = fromColumn
		c.arrangements[inst[2]] = toColumn
	}
}

// create a string of the top crates in each stack
func (c Crates) finalToppers() string {
	toppers := ""
	for _, stack := range c.arrangements {
		toppers += stack[len(stack)-1]
	}
	return toppers
}

func main() {
	input := file.Read()
	crates, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	crates.runCrateMover9000()
	fmt.Println("Part 1:", crates.finalToppers())

	// TODO is there a better way to get another crates struct?
	crates, err = parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	crates.runCrateMover9001()
	fmt.Println("Part 2:", crates.finalToppers())
}
