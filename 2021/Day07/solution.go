package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
)

func calculateFuelSpend(x int) int {
	spend := 0
	for i := 1; i <= x; i++ {
		spend += i
	}
	return spend
}

func findMinFuelSpend(input []int, min, max int, part2 bool) int {
	minDist := utils.Infinity
	for x := min; x <= max; x++ {
		dist := 0
		for _, i := range input {
			if part2 {
				dist += calculateFuelSpend(utils.Abs(i - x))
			} else {
				dist += utils.Abs(i - x)
			}

		}
		if dist < minDist {
			minDist = dist
		}
	}
	return minDist
}

func main() {
	input, err := utils.ReadFileSingleLineAsInts()
	if err != nil {
		fmt.Println(err)
		return
	}
	min, max := utils.FindExtremities(input)
	fmt.Println("Part 1:", findMinFuelSpend(input, min, max, false))
	fmt.Println("Part 2:", findMinFuelSpend(input, min, max, true))
}
