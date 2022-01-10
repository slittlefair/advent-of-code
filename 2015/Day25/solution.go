package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"regexp"
	"strconv"
)

func parseInput(input []string) (int, int, error) {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(input[0], -1)
	if len(matches) != 2 {
		return -1, -1, fmt.Errorf("got %d numbers, should have 2 from %v", len(matches), input)
	}
	// We can ignore these errors as we'll know they'll convert due to regex match
	row, _ := strconv.Atoi(matches[0])
	col, _ := strconv.Atoi(matches[1])
	return row, col, nil
}

func nthNumber(row, col int) int {
	return (row+col-1)*(row+col)/2 - row + 1
}

func getCodeAt(row, col int) int {
	code := 20151125
	for i := 1; i < nthNumber(row, col); i++ {
		code = (code * 252533) % 33554393
	}
	return code
}

func getSolution(input []string) (int, error) {
	row, col, err := parseInput(input)
	if err != nil {
		return -1, err
	}
	return getCodeAt(row, col), nil
}

func main() {
	input := utils.ReadFile()
	code, err := getSolution(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", code)
}
