package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

func paperForPresent(dimensions []int) int {
	return 3*dimensions[0]*dimensions[1] + 2*dimensions[1]*dimensions[2] + 2*dimensions[0]*dimensions[2]
}

func ribbonForPresent(dimensions []int) int {
	return 2*dimensions[0] + 2*dimensions[1] + dimensions[0]*dimensions[1]*dimensions[2]
}

func totalPaperForPresents(presents []string) (int, int, error) {
	re := regexp.MustCompile(`\d+`)
	paper := 0
	ribbon := 0
	for _, present := range presents {
		dimensions := re.FindAllString(present, -1)
		if len(dimensions) != 3 {
			return -1, -1, fmt.Errorf("something went wrong, got dimensions %v", dimensions)
		}
		dimensionsInt := []int{}
		for _, d := range dimensions {
			// we can ignore the error as we know each dimension is an int due to the regex
			dInt, _ := strconv.Atoi(d)
			dimensionsInt = append(dimensionsInt, dInt)
		}
		sort.Ints(dimensionsInt)
		paper += paperForPresent(dimensionsInt)
		ribbon += ribbonForPresent(dimensionsInt)
	}
	return paper, ribbon, nil
}

func main() {
	input := helpers.ReadFile()
	paper, ribbon, err := totalPaperForPresents(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", paper)
	fmt.Println("Part 2:", ribbon)
}
