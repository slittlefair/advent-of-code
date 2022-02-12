package main

import (
	"Advent-of-Code/file"
	"fmt"
)

func getFloorFromInstructions(instructions string) int {
	var currentFloor int
	for _, str := range instructions {
		if string(str) == "(" {
			currentFloor++
		} else {
			currentFloor--
		}
	}
	return currentFloor
}

func getFirstInstanceOfBasement(instructions string) (int, error) {
	var currentFloor int
	for i, str := range instructions {
		if string(str) == "(" {
			currentFloor++
		} else {
			currentFloor--
		}
		if currentFloor == -1 {
			return i + 1, nil
		}
	}
	return -1, fmt.Errorf("did not get to basement")
}

func main() {
	instructions := file.Read()[0]
	fmt.Println("Part 1:", getFloorFromInstructions(instructions))
	basementIndex, err := getFirstInstanceOfBasement(instructions)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 2:", basementIndex)
}
