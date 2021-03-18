package main

import (
	helpers "Advent-of-Code"
	"errors"
	"fmt"
)

func halfSeats(pass string, min int, max int) (string, int, int, error) {
	direction := string(pass[0])
	if direction != "F" && direction != "B" && direction != "L" && direction != "R" {
		return "", 0, 0, errors.New("invalid character found")
	}
	if direction == "F" || direction == "L" {
		max = max - ((max - min + 1) / 2)
	} else {
		min = min + ((max - min + 1) / 2)
	}
	return string(pass[1:]), min, max, nil
}

func findMyID(usedIDs map[int]bool, lowestID int, highestID int) (int, error) {
	for i := lowestID; i <= highestID; i++ {
		if _, ok := usedIDs[i]; !ok {
			return i, nil
		}
	}
	return 0, errors.New("could not find my ID")
}

func part1(entries []string) (int, int, map[int]bool, error) {
	highestID := 0
	lowestID := 10000000
	usedIDs := make(map[int]bool)
	for _, entry := range entries {
		minRow := 1
		maxRow := 128
		minCol := 1
		maxCol := 8
		var err error
		for minRow != maxRow {
			entry, minRow, maxRow, err = halfSeats(entry, minRow, maxRow)
			// fmt.Println(pass, min, max)
			if err != nil {
				return 0, 0, nil, err
			}
		}
		for minCol != maxCol {
			entry, minCol, maxCol, err = halfSeats(entry, minCol, maxCol)
			if err != nil {
				return 0, 0, nil, err
			}
		}
		id := (minRow-1)*8 + minCol - 1
		if id > highestID {
			highestID = id
		}
		if id < lowestID {
			lowestID = id
		}
		usedIDs[id] = true
	}

	return lowestID, highestID, usedIDs, nil
}

func main() {
	entries := helpers.ReadFile()
	lowestID, highestID, usedIDs, err := part1(entries)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", highestID)
	myID, err := findMyID(usedIDs, lowestID, highestID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 2:", myID)
}
