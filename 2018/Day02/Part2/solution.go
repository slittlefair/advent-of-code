package main

import (
	"Advent-of-Code/file"
	"fmt"
)

func main() {
	lines := file.Read()
	// Cycle through ids
	for i := range lines {
		id1 := lines[i]
		// Cycle through ids that appear later in the list so that we don't compare twice
		for j := i + 1; j < len(lines); j++ {
			id2 := lines[j]
			// Collect indexes where the ids differ in character
			var differingIndexes []int
			// Cycle through the ids and compare characters
			for k := range id1 {
				if id1[k] != id2[k] {
					differingIndexes = append(differingIndexes, k)
				}
			}
			// If there's only one differing index then we've found our ids
			if len(differingIndexes) == 1 {
				idx := differingIndexes[0]
				fmt.Println(id1[:idx] + id1[idx+1:])
				return
			}
		}
	}
}
