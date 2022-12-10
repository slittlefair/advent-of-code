package main

import (
	"Advent-of-Code/file"
	"fmt"
)

func followInstructions(offsets []int) int {
	steps := 0
	currentIndex := 0

	for currentIndex < len(offsets) {
		nextIndex := currentIndex + offsets[currentIndex]
		offsets[currentIndex]++
		steps++
		currentIndex = nextIndex
	}

	return steps
}

func followInstructions2(offsets []int) int {
	steps := 0
	currentIndex := 0

	for currentIndex < len(offsets) {
		nextIndex := currentIndex + offsets[currentIndex]
		if offsets[currentIndex] >= 3 {
			offsets[currentIndex]--
		} else {
			offsets[currentIndex]++
		}
		steps++
		currentIndex = nextIndex
	}

	return steps
}

func main() {
	offsets := file.ReadAsInts()
	steps := followInstructions(offsets)
	fmt.Println(steps)
	offsets2 := file.ReadAsInts()
	steps2 := followInstructions2(offsets2)
	fmt.Println(steps2)
}
