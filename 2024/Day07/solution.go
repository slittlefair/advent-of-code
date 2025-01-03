package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/regex"
	"fmt"
	"strconv"
)

type Equation struct {
	testValue int
	equations []int
}

func parseInput(input []string) []Equation {
	equations := []Equation{}
	for _, line := range input {
		matches := regex.MatchNums.FindAllString(line, -1)
		eq := Equation{}
		for i, m := range matches {
			// We know this won't error since it's a match on the regex, so it is sure to convert
			num, _ := strconv.Atoi(m)
			if i == 0 {
				eq.testValue = num
			} else {
				eq.equations = append(eq.equations, num)
			}
		}
		equations = append(equations, eq)
	}
	return equations
}

// Define an operator
type operator interface {
	// Do the operation associated with the operator, and returns the result
	do(x, y int) int
	// Utility function for printing that returns a string representation of the operator
	print() string
}

// Add operator
type add struct{}

func (a add) do(x, y int) int {
	return x + y
}

func (a add) print() string {
	return "+"
}

// Multiply operator
type multiply struct{}

func (m multiply) do(x, y int) int {
	return x * y
}

func (m multiply) print() string {
	return "*"
}

// Concat operator (part 2 only)
type concat struct{}

func (c concat) do(x, y int) int {
	p := 10
	for p <= y {
		p *= 10
	}
	return x*p + y
}

func (c concat) print() string {
	return "||"
}

var part1Operators = []operator{add{}, multiply{}}
var part2Operators = []operator{add{}, multiply{}, concat{}}

// Recursive function that calls an opertor at a specific point of an equation. It returns true if
// it's the end of the equation and we match the test value, or false if it doesn't, or can't,
// evaluate to the test value.
// If we get a total less that the test value then we continue, calling this function again on the
// next element of the equation with each operator.
// We also keep track of the operators used, so that we can return an array of them in case we want
// to print which set of operators led to the true evaluation.
func (eq Equation) doOperation(op operator, index, tot int, allOperators, correctOps []operator) (bool, []operator) {
	// For the given operation, find the new total
	newTot := op.do(tot, eq.equations[index])

	// If the new total is greater than the test value then this path can never evaluate to true, so return
	if newTot > eq.testValue {
		return false, nil
	}

	// If we've reached the last element of the equation we'll need to stop anyway, so return whether
	// or not we've matched the test value.
	if index == len(eq.equations)-1 {
		return newTot == eq.testValue, correctOps
	}

	// If we're still going we need to continue along the equation, trying out each operator again.
	// Increment the index so we move to the next element and supply the new total.
	for _, op := range allOperators {
		// If we've returned that we've found a solution to this equation, return it
		if foundSol, correctOps := eq.doOperation(op, index+1, newTot, allOperators, append(correctOps, op)); foundSol {
			return true, correctOps
		}
	}

	return false, nil
}

// Utility function for printing a correct solution for an equation. It prints elements and operators
// in turn along with the test value.
// e.g. "6 * 8 || 6 * 15 = 7290"
func (eq Equation) Print(correctOps []operator) {
	for i, v := range eq.equations {
		if i == len(eq.equations)-1 {
			fmt.Printf("%v = %v\n", v, eq.testValue)
		} else {
			fmt.Printf("%v %v ", v, correctOps[i].print())
		}
	}
}

// For a given equation, finds whether it evaluates for part 1 or part 2
func (eq Equation) evaluateEquation() (bool, bool) {
	for _, op := range part1Operators {
		if foundSol, _ := eq.doOperation(op, 1, eq.equations[0], part1Operators, []operator{op}); foundSol {
			// If we find a solution for part 1 operators then it's true for part 2 as well
			return true, true
		}
	}
	for _, op := range part2Operators {
		if foundSol, _ := eq.doOperation(op, 1, eq.equations[0], part2Operators, []operator{op}); foundSol {
			return false, true
		}
	}
	return false, false
}

func findSolutions(input []string) (int, int) {
	equations := parseInput(input)
	part1 := 0
	part2 := 0
	for _, eq := range equations {
		p1, p2 := eq.evaluateEquation()
		if p1 {
			part1 += eq.testValue
			part2 += eq.testValue
		} else if p2 {
			part2 += eq.testValue
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
