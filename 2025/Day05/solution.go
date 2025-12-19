package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/maths"
	"Advent-of-Code/regex"
	"fmt"
	"strconv"
)

type FreshRange struct {
	min, max int
}

type Ingredients struct {
	fresh     []FreshRange
	available []int
}

func (ing *Ingredients) parseFreshLine(line string) error {
	matches := regex.MatchNums.FindAllString(line, -1)
	if len(matches) != 2 {
		return fmt.Errorf("invalid fresh line: %v", line)
	}
	minFresh, _ := strconv.Atoi(matches[0])
	maxFresh, _ := strconv.Atoi(matches[1])
	ing.fresh = append(ing.fresh, FreshRange{min: minFresh, max: maxFresh})
	return nil
}

func (ing *Ingredients) parseAvailableLine(line string) error {
	matches := regex.MatchNums.FindAllString(line, -1)
	if len(matches) != 1 {
		return fmt.Errorf("invalid available line: %v", line)
	}
	n, _ := strconv.Atoi(matches[0])
	ing.available = append(ing.available, n)
	return nil
}

func parseInput(input []string) (*Ingredients, error) {
	ing := &Ingredients{}
	f := ing.parseFreshLine
	for _, line := range input {
		if line == "" {
			f = ing.parseAvailableLine
			continue
		}
		err := f(line)
		if err != nil {
			return nil, err
		}
	}
	ing.pruneFreshRanges()
	return ing, nil
}

func rangesOverlap(a, b FreshRange) bool {
	if a.max < b.min {
		return false
	}
	if a.min > b.max {
		return false
	}
	return true
}

func compressRanges(a, b FreshRange) FreshRange {
	return FreshRange{min: maths.Min(a.min, b.min), max: maths.Max(a.max, b.max)}
}

func (ing *Ingredients) pruneFreshRanges() {
	freshMap := make(map[FreshRange]struct{})
	for _, fr := range ing.fresh {
		freshMap[fr] = struct{}{}
	}
out:
	for a := range freshMap {
		for b := range freshMap {
			if a == b {
				continue
			}
			if rangesOverlap(a, b) {
				delete(freshMap, a)
				delete(freshMap, b)
				freshMap[compressRanges(a, b)] = struct{}{}
				goto out
			}
		}
	}
	freshSlice := []FreshRange{}
	for fr := range freshMap {
		freshSlice = append(freshSlice, fr)
	}
	ing.fresh = freshSlice

}

func findSolutions(input []string) (int, int, error) {
	part1 := 0
	part2 := 0
	ing, err := parseInput(input)
	if err != nil {
		return part1, part2, err
	}
	for _, i := range ing.available {
		for _, fresh := range ing.fresh {
			if i >= fresh.min && i <= fresh.max {
				part1++
				break
			}
		}
	}
	for _, fr := range ing.fresh {
		part2 += (fr.max - fr.min) + 1
	}
	return part1, part2, nil
}

func main() {
	input := file.Read()
	part1, part2, err := findSolutions(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}
