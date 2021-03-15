package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

var input = "ckczppom"

func main() {
	i := 1
	var part1Solution int
	var part2Solution int
	for {
		code := input + strconv.Itoa(i)
		byteConversion := md5.Sum([]byte(code))
		conversion := hex.EncodeToString(byteConversion[:])
		foundPart1Solution := true
		foundPart2Solution := true
		for _, l := range conversion[:5] {
			if string(l) != "0" {
				foundPart1Solution = false
				foundPart2Solution = false
			}
		}
		if string(conversion[5]) != "0" {
			foundPart2Solution = false
		}
		if foundPart1Solution && part1Solution == 0 {
			part1Solution = i
		}
		if foundPart2Solution {
			part2Solution = i
		}
		if foundPart1Solution && foundPart2Solution {
			break
		}
		i++
	}
	fmt.Println("Part 1:", part1Solution)
	fmt.Println("Part 2:", part2Solution)
}
