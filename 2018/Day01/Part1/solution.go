package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"strconv"
)

func main() {
	lines := utils.ReadFile()
	var sum int
	for _, val := range lines {
		n, err := strconv.Atoi(val)
		utils.Check(err)
		sum = sum + n
	}
	fmt.Println(sum)
}
