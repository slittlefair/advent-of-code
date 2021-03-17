package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
	"strings"
)

type passwords struct {
	min      int
	max      int
	letter   string
	password string
}

func readPassword(match []string) *passwords {
	return &passwords{
		min:      helpers.StringToInt(match[1]),
		max:      helpers.StringToInt(match[2]),
		letter:   match[3],
		password: match[4],
	}
}

func part1(entries []string, re *regexp.Regexp) int {
	valid := 0

	for _, entry := range entries {
		match := re.FindStringSubmatch(entry)
		password := readPassword(match)
		if count := strings.Count(password.password, password.letter); count >= password.min && count <= password.max {
			valid++
		}
	}

	return valid
}

func part2(entries []string, re *regexp.Regexp) int {
	valid := 0

	for _, entry := range entries {
		match := re.FindStringSubmatch(entry)
		password := readPassword(match)
		letter1 := password.password[password.min-1]
		letter2 := password.password[password.max-1]
		if letter1 != letter2 && (string(letter1) == password.letter || string(letter2) == password.letter) {
			valid++
		}
	}

	return valid
}

func main() {
	entries := helpers.ReadFile()
	re := regexp.MustCompile(`^(\d+)-(\d+) (\w+): (\w+)$`)

	fmt.Println("Part 1:", part1(entries, re))
	fmt.Println("Part 2:", part2(entries, re))
}
