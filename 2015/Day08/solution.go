package main

import (
	"Advent-of-Code"
	"fmt"
)

func main() {
	lines := helpers.ReadFile()
	part1Total := 0
	part2Total := 0
	for _, l := range lines {
		part1Total += 2
		part2Total += 2
		for i := 0; i < len(l); i++ {
			if string(l[i]) == "\\" {
				if string(l[i+1]) == "x" {
					part1Total += 3
				} else {
					part1Total++
				}
				if string(l[i+1]) == "\\" {
					i++
					part2Total++
				}
			}
			if string(l[i]) == "\"" || string(l[i]) == "\\" || (string(l[i]) == "\\" && string(l[i+1]) == "x") {
				part2Total++
			}
		}
	}
	fmt.Println("Part 1:", part1Total)
	fmt.Println("Part 2:", part2Total)
}
