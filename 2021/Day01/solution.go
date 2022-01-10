package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
)

func calculateIncreases(input []int) int {
	numIncreases := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			numIncreases++
		}
	}
	return numIncreases
}

func calculateSlidingWindows(input []int) []int {
	windows := []int{}
	for i := 0; i < len(input)-2; i++ {
		windows = append(windows, input[i]+input[i+1]+input[i+2])
	}
	return windows
}

func part1(input []int) int {
	return calculateIncreases(input)
}

func part2(input []int) int {
	return calculateIncreases(calculateSlidingWindows((input)))
}

func main() {
	input := utils.ReadFileAsInts()
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
