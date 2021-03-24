package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
)

type NumsSaid struct {
	lastSaid        int
	penultimateSaid int
}

type WhenNumsSaid map[int]NumsSaid

func main() {
	numList := helpers.ReadFile()[0]
	re := regexp.MustCompile(`\d+`)
	numStrings := helpers.StringSliceToIntSlice(re.FindAllString(numList, -1))
	whenNumsSaid := WhenNumsSaid{}
	lastNumSaid := 0
	for i, num := range numStrings {
		whenNumsSaid[num] = NumsSaid{
			lastSaid: i + 1,
		}
		lastNumSaid = num
	}
	for i := len(numStrings) + 1; i <= 30000000; i++ {
		if whenNumsSaid[lastNumSaid].penultimateSaid == 0 {
			lastNumSaid = 0
		} else {
			lastNumSaid = whenNumsSaid[lastNumSaid].lastSaid - whenNumsSaid[lastNumSaid].penultimateSaid
		}
		if _, ok := whenNumsSaid[lastNumSaid]; !ok {
			whenNumsSaid[lastNumSaid] = NumsSaid{
				lastSaid: i,
			}
		} else {
			whenNumsSaid[lastNumSaid] = NumsSaid{
				penultimateSaid: whenNumsSaid[lastNumSaid].lastSaid,
				lastSaid:        i,
			}
		}
		if i == 2020 {
			fmt.Println("Part 1:", lastNumSaid)
		}
	}
	fmt.Println("Part 2:", lastNumSaid)
}
