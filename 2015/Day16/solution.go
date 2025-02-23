package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/regex"
	"errors"
	"fmt"
	"strconv"
)

func findExactMatch(input []string, ticker map[string]int) (string, error) {
	for _, line := range input {
		words := regex.MatchLettersLower.FindAllString(line, -1)[1:]
		nums := regex.MatchNums.FindAllString(line, -1)
		if len(words) != len(nums)-1 {
			return "", fmt.Errorf("line wasn't parsed correctly, expected one fewer elem in %v than %v", words, nums)
		}
		for i, w := range words {
			// We know this will convert without error due to regex match
			n, _ := strconv.Atoi(nums[i+1])
			if ticker[w] != n {
				goto out
			}
		}
		return nums[0], nil
	out:
	}
	return "", errors.New("could not find matching sue")
}

func findRangedMatch(input []string, ticker map[string]int) (string, error) {
	for _, line := range input {
		words := regex.MatchLettersLower.FindAllString(line, -1)[1:]
		nums := regex.MatchNums.FindAllString(line, -1)
		if len(words) != len(nums)-1 {
			return "", fmt.Errorf("line wasn't parsed correctly, expected one fewer elem in %v than %v", words, nums)
		}
		for i, w := range words {
			n, _ := strconv.Atoi(nums[i+1])
			switch w {
			case "cats", "trees":
				if ticker[w] >= n {
					goto out
				}
			case "pomeranians", "goldfish":
				if ticker[w] <= n {
					goto out
				}
			default:
				if ticker[w] != n {
					goto out
				}
			}
		}
		return nums[0], nil
	out:
	}
	return "", errors.New("could not find matching sue")
}

func main() {
	input := file.Read()
	ticker := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	sueNumber, err := findExactMatch(input, ticker)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", sueNumber)

	sueNumber, err = findRangedMatch(input, ticker)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 2:", sueNumber)
}
