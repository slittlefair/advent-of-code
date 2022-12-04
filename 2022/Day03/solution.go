package main

import (
	"Advent-of-Code/file"
	"fmt"
	"strings"
	"unicode"
)

type Items []rune

// Find common items for part 1
func (items *Items) findCommonItems(input []string) {
	for _, sack := range input {
		comp1 := sack[:len(sack)/2]
		comp2 := sack[len(sack)/2:]
		comp1Runes := map[rune]struct{}{}
		for _, r := range comp1 {
			comp1Runes[r] = struct{}{}
		}
		for _, r := range comp2 {
			if _, ok := comp1Runes[r]; ok {
				*items = append(*items, r)
				break
			}
		}
	}
}

// Find badges for part 2
func (items *Items) getBadges(input []string) {
	for i := 0; i < len(input); i = i + 3 {
		for _, r := range input[i] {
			if strings.ContainsRune(input[i+1], r) && strings.ContainsRune(input[i+2], r) {
				*items = append(*items, r)
				break
			}
		}
	}
}

func (items Items) sumPriorities() int {
	sum := 0
	for _, itm := range items {
		if unicode.IsLower(itm) {
			sum += int(itm-'a') + 1
		} else {
			sum += int(itm-'A') + 27
		}
	}
	return sum
}

func main() {
	input := file.Read()
	i := Items{}
	i.findCommonItems(input)
	fmt.Println("Part 1:", i.sumPriorities())
	i = Items{}
	i.getBadges(input)
	fmt.Println("Part 2:", i.sumPriorities())
}
