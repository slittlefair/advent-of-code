package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
)

func main() {
	entries := helpers.ReadFile()

	fmt.Println(entries)
	re := regexp.MustCompile(`\w+ \w+ bag?|contain`)
	for _, entry := range entries {
		matches := re.FindAllString(entry, -1)
		for _, match := range matches {
			fmt.Println(match, match == "shiny gold bag")
		}
	}
}
