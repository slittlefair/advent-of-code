package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"strconv"
)

func main() {
	lines := utils.ReadFile()

	m := make(map[int]bool)
	var frequency int

	for {
		for _, val := range lines {
			n, err := strconv.Atoi(val)
			utils.Check(err)
			frequency = frequency + n
			if m[frequency] {
				fmt.Println(frequency)
				return
			}
			m[frequency] = true
		}
	}
}
