package main

import (
	"Advent-of-Code/file"
	"fmt"
	"strconv"
)

func main() {
	lines := file.Read()
	var sum int
	for _, val := range lines {
		n, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println(err)
			return
		}
		sum = sum + n
	}
	fmt.Println(sum)
}
