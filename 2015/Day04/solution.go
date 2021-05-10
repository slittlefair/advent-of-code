package main

import (
	helpers "Advent-of-Code"
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func hashIsValidPart1(hash string) bool {
	return strings.HasPrefix(hash, "00000")
}

func hashIsValidPart2(hash string) bool {
	return strings.HasPrefix(hash, "000000")
}

func findValidHash(input string, part1 bool) int {
	i := 0
	for {
		str := input + strconv.Itoa(i)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(str)))
		if part1 && hashIsValidPart1(hash) {
			return i
		}
		if !part1 && hashIsValidPart2(hash) {
			return i
		}
		i++
	}
}

func main() {
	input := helpers.ReadFile()[0]
	fmt.Println("Part 1:", findValidHash(input, true))
	fmt.Println("Part 2:", findValidHash(input, false))
}
