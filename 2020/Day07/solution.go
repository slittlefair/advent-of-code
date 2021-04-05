package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type BagMap map[string]map[string]int

func (bm *BagMap) parseBag(entries []string) {
	re := regexp.MustCompile(`(\d+) (\w+ \w+)`)
	for _, entry := range entries {
		foundBags := strings.Split(entry, " bags contain ")
		containingBag := foundBags[0]
		bagContents := make(map[string]int)
		subBags := re.FindAllStringSubmatch(foundBags[1], -1)
		for _, bag := range subBags {
			// Because of the regexp matching we know this conversion won't return an error
			s, _ := strconv.Atoi(bag[1])
			bagContents[bag[2]] = s
		}
		(*bm)[containingBag] = bagContents
	}
}

func (bm BagMap) traverseBags(bag string, bagsFound map[string]bool) {
	for colour, containingBag := range bm {
		if _, exists := containingBag[bag]; exists {
			if _, ok := bagsFound[colour]; !ok {
				bagsFound[colour] = true
				bm.traverseBags(colour, bagsFound)
			}
		}
	}
}

func (bm BagMap) getBagsFound(myBag string) int {
	bagsFound := make(map[string]bool)
	bm.traverseBags(myBag, bagsFound)
	return len(bagsFound)
}

func (bm BagMap) countBags(bag string) int {
	total := 0
	contents := bm[bag]
	for colour, numBags := range contents {
		subCount := bm.countBags(colour)
		if subCount > 0 {
			total += numBags * subCount
		}
		total += numBags
	}
	return total
}

func main() {
	entries := helpers.ReadFile()
	bagMap := &BagMap{}
	bagMap.parseBag(entries)
	myBag := "shiny gold"

	fmt.Println("Part 1:", bagMap.getBagsFound(myBag))
	fmt.Println("Part 2:", bagMap.countBags(myBag))
}
