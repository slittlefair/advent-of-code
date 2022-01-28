package main

import (
	"Advent-of-Code/file"
	"fmt"
	"strconv"
)

func main() {
	lines := file.Read()

	m := make(map[int]bool)
	var frequency int

	for {
		for _, val := range lines {
			n, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println(err)
				return
			}
			frequency = frequency + n
			if m[frequency] {
				fmt.Println(frequency)
				return
			}
			m[frequency] = true
		}
	}
}
