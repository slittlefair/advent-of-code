package main

import (
	helpers "Advent-of-Code"
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
	minDist := helpers.Infinty
	for x := min; x <= max; x++ {
		dist := 0
		for _, i := range input {
			if part2 {
				dist += calculateFuelSpend(helpers.Abs(i - x))
			} else {
				dist += helpers.Abs(i - x)
			}

		}
		if dist < minDist {
			minDist = dist
		}
	}
	return minDist
}

func main() {
	input, err := helpers.ReadFileSingleLineAsInts()
	if err != nil {
		fmt.Println(err)
		return
	}
	min, max := helpers.FindExtremities(input)
	fmt.Println("Part 1:", findMinFuelSpend(input, min, max, false))
	fmt.Println("Part 2:", findMinFuelSpend(input, min, max, true))
}
