package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"strconv"
)

type frequencies struct {
	zeros int
	ones  int
}

type allFrequencies []frequencies

func (f allFrequencies) compileFrequencies(input []string) {
	for _, s := range input {
		for i, j := range s {
			if str := string(j); str == "0" {
				f[i].zeros++
			} else {
				f[i].ones++
			}
		}
	}
}

func (f allFrequencies) compileRates() (string, string) {
	var gRate, eRate string
	for _, i := range f {
		if i.zeros > i.ones {
			gRate += "0"
			eRate += "1"
		} else {
			gRate += "1"
			eRate += "0"
		}
	}
	return gRate, eRate
}

func part1(input []string) (int64, error) {
	allFreqs := make(allFrequencies, len(input[0]))
	allFreqs.compileFrequencies(input)
	gRate, eRate := allFreqs.compileRates()

	// compileRates() should never return a string that can't be parsed to binary, so error should
	// never not be nil. However, we like to ensure good code quality so we handle it anyway
	g, err := strconv.ParseInt(gRate, 2, 64)
	if err != nil {
		return -1, err
	}

	e, err := strconv.ParseInt(eRate, 2, 64)
	if err != nil {
		return -1, err
	}

	return g * e, nil
}

func getRatings(input []string, og bool) (string, error) {
	validNums := []string{}
	validNums = append(validNums, input...)
	for i := 0; i < len(input[0]); i++ {
		f := frequencies{}
		for _, line := range validNums {
			if string(line[i]) == "0" {
				f.zeros++
			} else {
				f.ones++
			}
		}
		newNums := []string{}
		for _, num := range validNums {
			if str := string(num[i]); og {
				if (str == "0" && f.zeros > f.ones) || (str == "1" && f.ones >= f.zeros) {
					newNums = append(newNums, num)
				}
			} else if (str == "0" && f.zeros <= f.ones) || (str == "1" && f.ones < f.zeros) {
				newNums = append(newNums, num)
			}
		}
		validNums = newNums
		// If there's only one number left, return it
		if len(validNums) == 1 {
			return validNums[0], nil
		}
	}
	// If we've gone through all the numbers and haven't narrowed it down to 1, something went wrong
	return "", fmt.Errorf("could not narrow down nums, left with %v", validNums)
}

func part2(input []string) (int64, error) {
	ogRatingString, err := getRatings(input, true)
	if err != nil {
		return -1, err
	}
	ogRating, err := strconv.ParseInt(ogRatingString, 2, 64)
	if err != nil {
		return -1, err
	}
	c02sRatingString, err := getRatings(input, false)
	if err != nil {
		return -1, err
	}
	c02sRating, err := strconv.ParseInt(c02sRatingString, 2, 64)
	if err != nil {
		return -1, err
	}
	return ogRating * c02sRating, nil
}

func main() {
	input := helpers.ReadFile()
	part1, err := part1(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)

	part2, err := part2(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 2:", part2)
}
