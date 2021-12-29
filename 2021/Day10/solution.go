package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"sort"
)

var delimeters = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var delimeterCost = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var autoCompleteCost = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func invalidDelimeter(line string) (string, []string) {
	chunks := []string{}
	for _, r := range line {
		l := string(r)
		if _, ok := delimeters[l]; ok {
			chunks = append(chunks, l)
		} else if l == delimeters[chunks[len(chunks)-1]] {
			chunks = chunks[:len(chunks)-1]
		} else {
			return l, nil
		}
	}
	return "", chunks
}

func lineCost(chunks []string) int {
	cost := 0
	for i := len(chunks) - 1; i >= 0; i-- {
		cost *= 5
		cost += autoCompleteCost[delimeters[chunks[i]]]
	}
	return cost
}

func findSolutions(input []string) (int, int) {
	invDel := 0
	scores := []int{}
	for _, i := range input {
		delimeter, chunks := invalidDelimeter(i)
		if delimeter != "" {
			invDel += delimeterCost[delimeter]
		} else {
			scores = append(scores, lineCost(chunks))
		}
	}
	sort.Ints(scores)
	return invDel, scores[len(scores)/2]
}

func main() {
	input := helpers.ReadFile()
	part1, part2 := findSolutions(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
