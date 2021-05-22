package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"strings"
)

func lookAndSay(num string) string {
	split := strings.Split(num, "")
	say := ""
	currentNum := split[0]
	freq := 1
	for i := 1; i < len(split); i++ {
		char := split[i]
		if char == currentNum {
			freq++
		} else {
			say = fmt.Sprintf("%s%d%s", say, freq, currentNum)
			currentNum = char
			freq = 1
		}
	}
	say = fmt.Sprintf("%s%d%s", say, freq, currentNum)
	return say
}

func main() {
	num := helpers.ReadFile()[0]
	for i := 1; i <= 50; i++ {
		num = lookAndSay(num)
		if i == 40 {
			fmt.Println("Part 1:", len(num))
		}
		if i == 50 {
			fmt.Println("Part 2:", len(num))
		}
	}
}
