package main

import (
	helpers "Advent-of-Code"
	"errors"
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

var re = regexp.MustCompile(`^(\d+)-(\d+) (\w+): (\w+)$`)

func populatePasswordCollection(input []string) ([]passwords, error) {
	passwordCollection := []passwords{}
	for _, val := range input {
		match := re.FindStringSubmatch(val)
		if len(match) != 5 {
			return nil, errors.New("match is not 5 items long")
		}
		password := readPassword(match)
		passwordCollection = append(passwordCollection, password)
	}
	return passwordCollection, nil
}

func readPassword(match []string) passwords {
	return passwords{
		min:      helpers.StringToInt(match[1]),
		max:      helpers.StringToInt(match[2]),
		letter:   match[3],
		password: match[4],
	}
}

func getSolutions(passwordCollection []passwords) (int, int) {
	part1ValidCount, part2ValidCount := 0, 0
	for _, password := range passwordCollection {
		if count := strings.Count(password.password, password.letter); count >= password.min && count <= password.max {
			part1ValidCount++
		}

		letter1 := password.password[password.min-1]
		letter2 := password.password[password.max-1]
		if letter1 != letter2 && (string(letter1) == password.letter || string(letter2) == password.letter) {
			part2ValidCount++
		}
	}
	return part1ValidCount, part2ValidCount
}

func main() {
	input := helpers.ReadFile()
	passwordCollection, err := populatePasswordCollection(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	part1Solution, part2Solution := getSolutions(passwordCollection)
	fmt.Println("Part 1:", part1Solution)
	fmt.Println("Part 2:", part2Solution)
}
