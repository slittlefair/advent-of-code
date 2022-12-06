package main

import (
	"Advent-of-Code/file"
	"fmt"
)

// Find the index after the makrer that contains n distinct characters
func findStartMarker(input string, n int) (int, error) {
	// Frequency map of characters
	chars := map[string]int{}
	for i := 0; i < len(input)-1; i++ {
		chars[string(input[i])]++
		// If the length of the frequency map is n and all frequencies are 1 then we have found
		// n distinct consecutive characters
		if len(chars) == n {
			foundPacket := true
			if foundPacket {
				return i + 1, nil
			}
		}
		// If we are more than n characters through the input then we can start removing characters
		// that have come before our current one
		if i >= n-1 {
			c := string(input[i-(n-1)])
			chars[c]--
			// We are looking for the map to be n characters long so delete any that have 0 frequency
			if chars[c] == 0 {
				delete(chars, c)
			}
		}
	}

	return -1, fmt.Errorf("marker not found")
}

func main() {
	input := file.Read()[0]
	part1, err := findStartMarker(input, 4)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	part2, err := findStartMarker(input, 14)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 2:", part2)
}
