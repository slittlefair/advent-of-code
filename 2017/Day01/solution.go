package main

import (
	"Advent-of-Code/file"
	"fmt"
	"strconv"
)

func parseInput(input string) ([]int, error) {
	ipt := []int{}
	for _, n := range input {
		i, err := strconv.Atoi(string(n))
		if err != nil {
			return nil, err
		}
		ipt = append(ipt, i)
	}
	return ipt, nil
}

func findSolution(nums []int) (int, int) {
	part1, part2 := 0, 0
	length := len(nums)
	// Part1 - compare the last element to the first
	if nums[length-1] == nums[0] {
		part1 += nums[0]
	}
	// Part2 - compare the last element to the middle element-1
	if nums[length-1] == nums[length/2-1] {
		part2 += nums[length-1]
	}
	for i := 0; i < length-1; i++ {
		if nums[i] == nums[i+1] {
			part1 += nums[i]
		}
		if nums[i] == nums[(i+length/2)%length] {
			part2 += nums[i]
		}
	}
	return part1, part2
}

func main() {
	input := file.Read()
	nums, err := parseInput(input[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	part1, part2 := findSolution(nums)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
