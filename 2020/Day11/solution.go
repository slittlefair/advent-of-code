package main

import (
	helpers "Advent-of-Code"
	"fmt"
)

func main() {
	plan := helpers.ReadFile()

	for _, p := range plan {
		fmt.Println(p)
	}
}
