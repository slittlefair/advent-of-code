package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
)

var ticker = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

var reWord = regexp.MustCompile("[a-z]+")
var reNums = regexp.MustCompile("\\d+")

func part1(lines []string) {
	for _, l := range lines {
		properties := reWord.FindAllString(l, -1)[1:]
		nums := helpers.StringSliceToIntSlice(reNums.FindAllString(l, -1))
		sueNum := nums[0]
		values := nums[1:]
		match := true
		for i := range properties {
			if ticker[properties[i]] != values[i] {
				match = false
			}
		}
		if match == true {
			fmt.Println("Part 1:", sueNum)
			return
		}
	}
}

func part2(lines []string) {
	for _, l := range lines {
		properties := reWord.FindAllString(l, -1)[1:]
		nums := helpers.StringSliceToIntSlice(reNums.FindAllString(l, -1))
		sueNum := nums[0]
		values := nums[1:]
		match := true
		for i := range properties {
			if properties[i] == "cats" || properties[i] == "trees" {
				if ticker[properties[i]] >= values[i] {
					match = false
				}
			} else if properties[i] == "pomeranians" || properties[i] == "goldfish" {
				if ticker[properties[i]] <= values[i] {
					match = false
				}
			} else if ticker[properties[i]] != values[i] {
				match = false
			}
		}
		if match == true {
			fmt.Println("Part 2:", sueNum)
			return
		}
	}
}

func main() {
	lines := helpers.ReadFile()
	part1(lines)
	part2(lines)
}
