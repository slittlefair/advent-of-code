package main

import (
	"Advent-of-Code/file"
	"fmt"
	"strconv"
	"strings"
)

func lookAndSay(num []string) []string {
	say := []string{}
	currentNum := num[0]
	freq := 1
	for i := 1; i < len(num); i++ {
		char := num[i]
		if char == currentNum {
			freq++
		} else {
			say = append(say, strconv.Itoa(freq), currentNum)
			currentNum = char
			freq = 1
		}
	}
	say = append(say, strconv.Itoa(freq), currentNum)
	return say
}

func main() {
	input := file.Read()[0]
	nums := strings.Split(input, "")
	for i := 1; i <= 50; i++ {
		nums = lookAndSay(nums)
		if i == 40 {
			fmt.Println("Part 1:", len(nums))
		}
		if i == 50 {
			fmt.Println("Part 2:", len(nums))
		}
	}
}
