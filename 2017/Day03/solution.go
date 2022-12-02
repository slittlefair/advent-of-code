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

func main() {
	target := file.ReadAsInts()[0]
	layer, corner, prevCorner := getLayer(target)
	midDistance := getMiddlePointDistance(target, corner, prevCorner)
	fmt.Println("Part 1:", layer+midDistance)
}
