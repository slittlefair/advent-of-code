package main

import (
	"Advent-of-Code"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func react(polymer string) int {
	for {
		reacted := false
		for i := 0; i < len(polymer)-1; {
			unitRune := rune(polymer[i])
			unitPlusRune := rune(polymer[i+1])
			if unicode.ToLower(unitRune) == unicode.ToLower(unitPlusRune) {
				if unicode.IsLower(unitRune) != unicode.IsLower(unitPlusRune) {
					polymer = polymer[:i] + polymer[i+2:]
					reacted = true
				}
			}
			i++
		}
		if !reacted {
			return utf8.RuneCountInString(polymer)
		}
	}
}

func removeCharacters(input string, characters string) string {
	filter := func(r rune) rune {
		if strings.IndexRune(characters, r) < 0 {
			return r
		}
		return -1
	}
	return strings.Map(filter, input)
}

var alphabet = strings.Split("abcdefghijklmnopqrstuvwxyz", "")

func main() {
	polymer := helpers.ReadFile()[0]
	length := react(polymer)
	fmt.Println("Part 1:", length)
	var minLength = 1000000
	for _, letter := range alphabet {
		newPoly := removeCharacters(polymer, letter+strings.ToUpper(letter))
		length := react(newPoly)
		fmt.Println(letter, length)
		if length < minLength {
			minLength = length
		}
	}
	fmt.Println("Part 2:", minLength)
}
