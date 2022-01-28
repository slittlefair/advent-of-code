package main

import (
	"Advent-of-Code/file"
	"errors"
	"fmt"
)

// halfSeats takes the string of directions, the min seat and the max seat and returns a new string,
// min and max. It takes the direction as the first character in the string and returns an error if
// it's not valid. Otherwise it partitions the seats depending on the direction and alters the min
// or max accordingly. It returns these, as well as the remainder of the directions, minus the
// direction we just evaluated.
func halfSeats(dirs string, min int, max int) (string, int, int, error) {
	direction := string(dirs[0])
	if direction != "F" && direction != "B" && direction != "L" && direction != "R" {
		return "", 0, 0, errors.New("invalid character found")
	}
	if direction == "F" || direction == "L" {
		max = max - ((max - min + 1) / 2)
	} else {
		min = min + ((max - min + 1) / 2)
	}
	return dirs[1:], min, max, nil
}

// findMyID takes a map of IDs we know are taken, the lowest and the highest IDs in the map. It then
// cycles through each ID between the two limits and returns the first ID (there is only one) not in
// the map, which is ours. We are told that the seat with the ID +1 and -1 to ours is taken, so we
// don't have an edge seat and therefore must be in the range of lowest to highest. We can therefore
// return an error if we can't find a "blank" ID
func findMyID(usedIDs map[int]bool, lowestID int, highestID int) (int, error) {
	for i := lowestID; i <= highestID; i++ {
		if _, ok := usedIDs[i]; !ok {
			return i, nil
		}
	}
	return 0, errors.New("could not find my ID")
}

func getusedIDs(entries []string) (int, int, map[int]bool, error) {
	highestID := 0
	lowestID := (127 * 8) + 7
	usedIDs := make(map[int]bool)
	for _, entry := range entries {
		minRow := 1
		maxRow := 128
		minCol := 1
		maxCol := 8
		var err error
		// Keep running halfSeats on the trimmed directions until we have partitioned to only one
		// row/column.
		for minRow != maxRow {
			entry, minRow, maxRow, err = halfSeats(entry, minRow, maxRow)
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
	entries := file.Read()
	lowestID, highestID, usedIDs, err := getusedIDs(entries)
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
