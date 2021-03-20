package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type BagMap map[string]map[string]int

func (bm *BagMap) traverseBags(bag string, bagsFound *map[string]bool) {
	for colour, containingBag := range *bm {
		if _, exists := containingBag[bag]; exists {
			if _, ok := (*bagsFound)[colour]; !ok {
				(*bagsFound)[colour] = true
				bm.traverseBags(colour, bagsFound)
			}
		}
	}
}

func parseBag(entries []string) (BagMap, error) {
	bagMap := make(map[string]map[string]int)

	for _, entry := range entries {
		foundBags := strings.Split(entry, " bags contain ")
		containingBag := foundBags[0]
		bagContents := make(map[string]int)
		re := regexp.MustCompile(`(\d+) (\w+ \w+)`)
		subBags := re.FindAllStringSubmatch(foundBags[1], -1)
		for _, bag := range subBags {
			s, err := strconv.Atoi(bag[1])
			if err != nil {
				return nil, err
			}
			bagContents[bag[2]] = s
		}
		bagMap[containingBag] = bagContents
	}

	return bagMap, nil
}

func (bm *BagMap) countBags(bag string) int {
	total := 0
	contents := (*bm)[bag]
	for colour, numBags := range contents {
		subCount := bm.countBags(colour)
		if subCount > 0 {
			total += numBags * subCount
		}
		total += numBags
	}
	return total
}

func part1(bagMap BagMap, myBag string) int {
	bagsFound := make(map[string]bool)
	bagMap.traverseBags(myBag, &bagsFound)
	return len(bagsFound)
}

func main() {
	entries := helpers.ReadFile()
	bagMap, err := parseBag(entries)
	if err != nil {
		fmt.Println(err)
		return
	}
	myBag := "shiny gold"

	fmt.Println("Part 1:", part1(bagMap, myBag))
	fmt.Println("Part 2:", bagMap.countBags(myBag))
}
