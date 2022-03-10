package main

import (
	"Advent-of-Code/file"
	"fmt"
	"regexp"
	"strconv"
)

type disc struct {
	id           int
	numPositions int
	position     int
}

type allDiscs map[int]*disc

func parseInput(input []string) (allDiscs, error) {
	ad := allDiscs{}
	re := regexp.MustCompile(`\d+`)
	for _, line := range input {
		matches := re.FindAllString(line, -1)
		if len(matches) != 4 {
			return nil, fmt.Errorf("did not get expected 4 ints from %v", line)
		}
		// We know all these can be converted due to regex match, so we won't get errors
		id, _ := strconv.Atoi(matches[0])
		numPosition, _ := strconv.Atoi(matches[1])
		pos, _ := strconv.Atoi(matches[3])
		ad[id] = &disc{
			id:           id,
			numPositions: numPosition,
			position:     pos,
		}
	}
	return ad, nil
}

func (ad allDiscs) moveDiscs() {
	for _, disc := range ad {
		disc.position = (disc.position + 1) % disc.numPositions
	}
}

func (ad allDiscs) getCapsule() bool {
	for _, disc := range ad {
		if (disc.position+disc.id-1)%disc.numPositions != 0 {
			return false
		}
	}
	return true
}

func (ad allDiscs) findSuccessfulTime() int {
	i := 0
	for {
		got := ad.getCapsule()
		if got {
			return i - 1
		}
		ad.moveDiscs()
		i++
	}
}

func findSolutions(input []string) (int, int, error) {
	ad, err := parseInput(input)
	if err != nil {
		return -1, -1, err
	}

	// We already parsed successfully so we know we won't error here
	ad2, _ := parseInput(input)

	// Add the new disc
	ad2[len(ad)+1] = &disc{
		id:           len(ad) + 1,
		numPositions: 11,
		position:     0,
	}

	return ad.findSuccessfulTime(), ad2.findSuccessfulTime(), nil
}

func main() {
	input := file.Read()
	part1, part2, err := findSolutions(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
