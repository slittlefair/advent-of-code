package main

import (
	"Advent-of-Code"
	"fmt"
	"strconv"
)

func main() {
	lines := helpers.ReadFile()

	m := make(map[int]bool)
	var frequency int

	for {
		for _, val := range lines {
			n, err := strconv.Atoi(val)
			helpers.Check(err)
			frequency = frequency + n
			if m[frequency] {
				fmt.Println(frequency)
				return
			}
			m[frequency] = true
		}
	}
}
