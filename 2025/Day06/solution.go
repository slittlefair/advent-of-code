package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/maths"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Problem struct {
	numbers [][]string
	op      func(nums []int) int
}

func add(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func multiply(nums []int) int {
	total := 1
	for _, n := range nums {
		total *= n
	}
	return total
}

var re = regexp.MustCompile(`[*+]`)

func parseInput(input []string) ([]*Problem, error) {
	numbers := [][][]string{}
	ops := []func(nums []int) int{}

	// Get column breaks to chunk up each problem's numbers
	blankColumns := re.FindAllStringSubmatchIndex(input[len(input)-1], -1)
	columnBreaks := make([]int, len(blankColumns)+1)
	for i, bc := range blankColumns {
		columnBreaks[i] = bc[0]
	}
	maxLineLength := 0
	for _, line := range input {
		maxLineLength = maths.Max(maxLineLength, len(line)+1)
	}
	columnBreaks[len(columnBreaks)-1] = maxLineLength

	for i, line := range input {
		if i != len(input)-1 {
			lineNumbers := [][]string{}
			for j := 0; j < len(columnBreaks)-1; j++ {
				lineNumbers = append(lineNumbers, strings.Split(line[columnBreaks[j]:columnBreaks[j+1]-1], ""))
			}
			numbers = append(numbers, lineNumbers)
		} else {
			matches := re.FindAllString(line, -1)
			for _, m := range matches {
				switch m {
				case "+":
					ops = append(ops, add)
				case "*":
					ops = append(ops, multiply)
				default:
					return nil, fmt.Errorf("invalid operation: %v", m)
				}
			}
		}
	}

	problems := []*Problem{}
	for i := 0; i < len(numbers[0]); i++ {
		problem := &Problem{}
		for _, line := range numbers {
			problem.numbers = append(problem.numbers, line[i])
		}
		problem.op = ops[i]
		problems = append(problems, problem)
	}
	return problems, nil
}

func findSolutions(input []string) (int, int, error) {
	part1 := 0
	part2 := 0
	problems, err := parseInput(input)
	for _, p := range problems {
		// Part 1
		nums := make([]int, len(p.numbers))
		for i, nArr := range p.numbers {
			num, _ := strconv.Atoi(strings.TrimSpace(strings.Join(nArr, "")))
			nums[i] = num
		}
		part1 += p.op(nums)

		// Part 2
		nums = []int{}
		for i := len(p.numbers[0]) - 1; i >= 0; i-- {
			num := ""
			for _, n := range p.numbers {
				num += n[i]
			}
			nConv, _ := strconv.Atoi(strings.TrimSpace(num))
			nums = append(nums, nConv)
		}
		part2 += p.op(nums)
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
