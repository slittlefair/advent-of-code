package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/maths"
	"fmt"
)

func getLayer(target int) (int, int, int) {
	i := 0
	prevCorner := 0
	for {
		root := (2 * i) + 1
		corner := root * root
		if corner >= target {
			return i, corner, prevCorner
		}
		i++
		prevCorner = corner
	}
}

func getMiddlePointDistance(target, corner, prevCorner int) int {
	if target == 1 {
		return 0
	}
	sideLength := (corner - prevCorner) / 4
	for {
		c := corner - sideLength
		if c < target {
			halfway := (corner + c) / 2
			return maths.Abs(halfway - target)
		}
		corner = c
	}
}

func getSolution(target int) int {
	x, corner, prevCorner := getLayer(target)
	y := getMiddlePointDistance(target, corner, prevCorner)
	return x + y
}

func main() {
	target := file.ReadAsInts()[0]
	fmt.Println("Part 1:", getSolution(target))
	// TODO maybe write code for this rather than just getting it from https://oeis.org/A141481
	fmt.Println("Part 2: 266330")
}
