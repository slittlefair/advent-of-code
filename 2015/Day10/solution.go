package main

import (
	"Advent-of-Code"
	"fmt"
	"strconv"
)

func lookAndSay(num string) (newNum string) {
	var figure = 0
	var freq = 0
	for i := 0; i <= len(num)-1; i++ {
		if figure == 0 {
			figure = helpers.StringToInt(string(num[i]))
			freq = 1
		} else if string(num[i]) != strconv.Itoa(figure) {
			newNum += strconv.Itoa(freq) + strconv.Itoa(figure)
			figure = helpers.StringToInt(string(num[i]))
			freq = 1
		} else {
			freq++
		}
	}
	if figure != 0 {
		newNum += strconv.Itoa(freq) + strconv.Itoa(figure)
	}
	return newNum
}

func main() {
	num := helpers.ReadFile()[0]
	for i := 0; i < 50; i++ {
		if i == 40 {
			fmt.Println("Part 1:", len(num))
		}
		num = lookAndSay(num)
	}
	fmt.Println("Part 2:", len(num))
}
