package main

import (
	helpers "Advent-of-Code"
	"fmt"
)

func readInputForPart1(input []string) int {
	// For each line rather than count the number of literal characters and subtract the number of
	// characters in memory, just count the number of "special" characters that are included in the
	// former but not the latter
	//
	// Each line starts and ends with double quotes ("), which are not stored in memory so
	// that's immediately a difference of two per line
	count := 2 * len(input)
	for _, line := range input {
		for i := 0; i < len(line); i++ {
			// If a character is a backslash escape character then there's something we're
			// interested in
			if string(line[i]) == "\\" {
				// If the following character is "x" then it's a unicode character of 4 literal
				// that represent 1 character in memory, so a difference of 3
				if string(line[i+1]) == "x" {
					count += 3
					continue
				} else {
					// If the following character is not "x" then it's either "\\" or "\"", in
					// either case it's one character difference between literal and memory
					count++
				}
				// If that following character is another backslash then skip it, as it's already
				// been evaluated and will cause the next character in the loop to add to our count,
				// which we don't want since it'll be a "normal" character that gets put into memory
				if string(line[i+1]) == "\\" {
					i++
				}
			}
		}
	}
	return count
}

func readInputForPart2(input []string) int {
	// For each line rather than count the number of encoded characters and subtract the number of
	// literal characters, just count the number of "special" characters that are included in the
	// former but not the latter
	//
	// We'll be surrounded the encoded strings with quotes anyway, so that's already a difference of
	// 2 for each line in the input
	count := 2 * len(input)
	for _, line := range input {
		for i := 0; i < len(line); i++ {
			if string(line[i]) == "\"" {
				// If a character is a quotation mark (") then it'll need a backslash added, so that
				// increases the difference by 1
				count++
				continue
			}
			if string(line[i]) == "\\" {
				// If a character is a backslash (\) then it'll need a backslash added, so that
				// increases the difference by 1
				count++
			}
		}
	}
	return count
}

func main() {
	input := helpers.ReadFile()
	fmt.Println("Part 1:", readInputForPart1(input))
	fmt.Println("Part 2:", readInputForPart2(input))
}
