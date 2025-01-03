package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/regex"
	"Advent-of-Code/strings"
	"fmt"
)

func isValidPart1(matches []string) bool {
	words := map[string]struct{}{}
	for _, word := range matches {
		if _, ok := words[word]; ok {
			return false
		}
		words[word] = struct{}{}
	}
	return true
}

func isValidPart2(matches []string) bool {
	for _, word := range matches {
		for _, word2 := range matches {
			if word != word2 && strings.AreAnagrams(word, word2) {
				return false
			}
		}
	}
	return true
}

func countValidPhrases(input []string) (int, int) {
	part1 := 0
	part2 := 0
	for _, line := range input {
		matches := regex.MatchWords.FindAllString(line, -1)
		if isValidPart1(matches) {
			part1++
			if isValidPart2(matches) {
				part2++
			}
		}
	}
	return part1, part2
}

func main() {
	input := file.Read()
	part1, part2 := countValidPhrases(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
