package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
)

func initialiseFields() map[string]bool {
	return map[string]bool{
		"byr": false,
		"iyr": false,
		"eyr": false,
		"hgt": false,
		"hcl": false,
		"ecl": false,
		"pid": false,
	}
}

func part1(entries []string) (int, error) {
	fields := initialiseFields()
	validPassports := 0
	for i, entry := range entries {
		for field, matched := range fields {
			if !matched {
				matched, err := regexp.MatchString(field, entry)
				if err != nil {
					fmt.Println("Error matching row")
					return 0, err
				}
				fields[field] = matched
			}
		}
		if len(entry) == 0 || i == len(entries)-1 {
			matches := 0
			for _, matched := range fields {
				if matched {
					matches++
				}
			}
			if matches == len(fields) {
				validPassports++
			}
			fmt.Println(i+1, matches, fields)
			fields = initialiseFields()
		}
	}
	return validPassports, nil
}

func part2() {

}

func main() {
	entries := helpers.ReadFile()
	validPassports, err := part1(entries)
	if err != nil {
		return
	}
	fmt.Println("Part 1:", validPassports)
}
