package main

import (
	"Advent-of-Code/file"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type orderingRules map[int]map[int]bool

var re = regexp.MustCompile(`\d+`)

func parseInput(input []string) (orderingRules, [][]int) {
	rules := orderingRules{}
	pages := [][]int{}

	parsingRules := true
	for _, line := range input {
		if line == "" {
			parsingRules = false
			continue
		}
		if parsingRules {
			matches := re.FindAllString(line, 2)
			// We can ignore errors since we've already used regex tomatch, so we know they can be converted
			n1, _ := strconv.Atoi(matches[0])
			n2, _ := strconv.Atoi(matches[1])
			if _, ok := rules[n1]; !ok {
				rules[n1] = make(map[int]bool)
			}
			if _, ok := rules[n2]; !ok {
				rules[n2] = make(map[int]bool)
			}
			rules[n1][n2] = true
			rules[n2][n1] = false
			continue
		}

		matches := re.FindAllString(line, -1)
		p := make([]int, len(matches))
		for i, m := range matches {
			// We can ignore errors since we've already used regex tomatch, so we know they can be converted
			n, _ := strconv.Atoi(m)
			p[i] = n
		}
		pages = append(pages, p)
	}

	return rules, pages
}

// Sort the pages using the ordering rules
func sortPages(rules orderingRules, pages []int) []int {
	sort.SliceStable(pages, func(i, j int) bool {
		pi := pages[i]
		pj := pages[j]

		if !rules[pi][pj] {
			return false
		}
		if rules[pj][pi] {
			return false
		}
		return true
	})
	return pages
}

func findSolutions(input []string) (int, int) {
	part1 := 0
	part2 := 0
	rules, publications := parseInput(input)

	for _, pages := range publications {
		for i, page := range pages {
			before := pages[:i]
			for _, p := range before {
				if v, ok := rules[page][p]; !ok {
					continue
				} else if v {
					goto sortFunc
				}
			}
			after := pages[i+1:]
			for _, p := range after {
				if v, ok := rules[page][p]; !ok {
					continue
				} else if !v {
					goto sortFunc
				}
			}
		}
		part1 += pages[len(pages)/2]
		continue
	sortFunc:
		pages = sortPages(rules, pages)
		part2 += pages[len(pages)/2]
	}
	return part1, part2
}

func main() {
	input := file.Read()
	part1, part2 := findSolutions(input)
	fmt.Printf("Part1: %d\n", part1)
	fmt.Printf("Part2: %d\n", part2)
}
