package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/maths"
	"Advent-of-Code/slice"
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
	minDist := maths.Infinity
	for x := min; x <= max; x++ {
		dist := 0
		for _, i := range input {
			if part2 {
				dist += calculateFuelSpend(maths.Abs(i - x))
			} else {
				dist += maths.Abs(i - x)
			}
		}
		if dist < minDist {
			minDist = dist
		}
	}
	return minDist
}

func main() {
	input, err := file.ReadSingleLineAsInts()
	if err != nil {
		fmt.Println(err)
		return
	}
	min, max := slice.FindExtremities(input)
	fmt.Println("Part 1:", findMinFuelSpend(input, min, max, false))
	fmt.Println("Part 2:", findMinFuelSpend(input, min, max, true))
}
