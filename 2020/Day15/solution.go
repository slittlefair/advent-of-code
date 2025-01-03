package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/regex"
	"Advent-of-Code/slice"
	"fmt"
)

type NumsSaid struct {
	lastSaid        int
	penultimateSaid int
}

type WhenNumsSaid map[int]NumsSaid

var lastNumSaid int

func (wns WhenNumsSaid) parseInput(numStrings []int) int {
	lastNumSaid := 0
	for i, num := range numStrings {
		wns[num] = NumsSaid{
			lastSaid: i + 1,
		}
		lastNumSaid = num
	}
	return lastNumSaid
}

func (wns WhenNumsSaid) playGame(startingIndex int) (int, int) {
	var part1Sol int
	for i := startingIndex; i <= 30000000; i++ {
		if wns[lastNumSaid].penultimateSaid == 0 {
			lastNumSaid = 0
		} else {
			lastNumSaid = wns[lastNumSaid].lastSaid - wns[lastNumSaid].penultimateSaid
		}
		if _, ok := wns[lastNumSaid]; !ok {
			wns[lastNumSaid] = NumsSaid{
				lastSaid: i,
			}
		} else {
			wns[lastNumSaid] = NumsSaid{
				penultimateSaid: wns[lastNumSaid].lastSaid,
				lastSaid:        i,
			}
		}
		if i == 2020 {
			part1Sol = lastNumSaid
		}
	}
	return part1Sol, lastNumSaid
}

func main() {
	numList := file.Read()[0]
	numStrings, err := slice.StringSliceToIntSlice(regex.MatchNums.FindAllString(numList, -1))
	if err != nil {
		fmt.Println(err)
		return
	}
	whenNumsSaid := WhenNumsSaid{}
	lastNumSaid = whenNumsSaid.parseInput(numStrings)
	part1Sol, part2Sol := whenNumsSaid.playGame(len(numStrings) + 1)
	fmt.Println("Part 1:", part1Sol)
	fmt.Println("Part 2:", part2Sol)
}
