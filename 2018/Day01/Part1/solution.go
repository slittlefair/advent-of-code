package main

import (
	"Advent-of-Code"
	"fmt"
	"strconv"
)

func main() {
	lines := helpers.ReadFile()
	var sum int
	for _, val := range lines {
		n, err := strconv.Atoi(val)
		helpers.Check(err)
		sum = sum + n
	}
	fmt.Println(sum)
}
