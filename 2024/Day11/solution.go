package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/regex"
	"fmt"
	"strconv"
	"strings"
)

type Stones map[string]int

func parseInput(input []string) Stones {
	stones := Stones{}
	matches := regex.MatchNums.FindAllString(input[0], -1)
	for _, m := range matches {
		stones[m]++
	}
	return stones
}

func split(n string) (string, string, bool) {
	l := len(n)
	if l%2 != 0 {
		return "", "", false
	}
	s1 := n[:l/2]
	s2 := n[l/2:]
	s2 = strings.TrimLeft(s2, "0")
	if s2 == "" {
		s2 = "0"
	}
	return s1, s2, true
}

func (s Stones) blink() Stones {
	newStones := Stones{}
	for v, freq := range s {
		if v == "0" {
			newStones["1"] += freq
			continue
		}
		s1, s2, success := split(v)
		if success {
			newStones[s1] += freq
			newStones[s2] += freq
			continue
		}
		// We've already checked all entries can be converted to ints in parseInput, so we can
		// ignore the error
		n, _ := strconv.Atoi(v)
		newStones[strconv.Itoa(n*2024)] += freq
	}
	return newStones
}

func (s Stones) findLength() int {
	sum := 0
	for _, freq := range s {
		sum += freq
	}
	return sum
}

func findSolutions(input []string) (int, int) {
	stones := parseInput(input)
	var part1 int
	for i := 0; i < 75; i++ {
		stones = stones.blink()
		if i == 24 {
			part1 = stones.findLength()
		}
	}
	return part1, stones.findLength()
}

func main() {
	input := file.Read()
	part1, part2 := findSolutions(input)
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}
