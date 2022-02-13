package main

import (
	"Advent-of-Code/file"
	"fmt"
	"regexp"
	"strconv"
)

func decompress(s string, part1 bool) int {
	re := regexp.MustCompile(`\d+`)
	length := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			nums := re.FindAllString(s[i:], 2)
			n0, _ := strconv.Atoi(nums[0])
			n1, _ := strconv.Atoi(nums[1])
			for {
				if string(s[i]) == ")" {
					i++
					break
				}
				i++
			}
			if part1 {
				length += (n0 * n1)
			} else {
				length += decompress(s[i:i+n0], false) * n1
			}
			i += n0 - 1
		} else {
			length++
		}
	}
	return length
}

func main() {
	input := file.Read()[0]
	fmt.Println("Part 1:", decompress(input, true))
	fmt.Println("Part 2:", decompress(input, false))
}
