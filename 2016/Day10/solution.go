package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/slice"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type bot struct {
	id   string
	vals []int
}

type bots map[string]*bot

func (b *bot) addValue(v int) {
	b.vals = append(b.vals, v)
	sort.Ints(b.vals)
}

func (bs bots) giveValue(b *bot, recLowID, recHighID string) {
	if recLow, ok := bs[recLowID]; !ok {
		bs[recLowID] = &bot{
			id:   recLowID,
			vals: []int{b.vals[0]},
		}
	} else {
		recLow.addValue(b.vals[0])
	}
	if recHigh, ok := bs[recHighID]; !ok {
		bs[recHighID] = &bot{
			id:   recHighID,
			vals: []int{b.vals[1]},
		}
	} else {
		recHigh.addValue(b.vals[1])
	}
	b.vals = []int{}
}

func (bs bots) add(id string, val int) {
	bs[id] = &bot{
		id:   id,
		vals: []int{val},
	}
}

func (bs bots) handleValueLine(split []string) error {
	if l := len(split); l != 6 {
		return fmt.Errorf("expected 6 length value line, got %d: %s", l, split)
	}
	id := fmt.Sprintf("%s %s", split[4], split[5])
	val, err := strconv.Atoi(split[1])
	if err != nil {
		return err
	}
	if b, ok := bs[id]; !ok {
		bs.add(id, val)
	} else {
		b.addValue(val)
	}
	return nil
}

func (bs bots) handleGiveLine(split []string) (bool, error) {
	if l := len(split); l != 12 {
		return false, fmt.Errorf("expected 12 length value line, got %d: %s", l, split)
	}
	b := bs[fmt.Sprintf("%s %s", split[0], split[1])]
	if b == nil || len(b.vals) != 2 {
		return false, nil
	}
	recLowID := fmt.Sprintf("%s %s", split[5], split[6])
	recHighID := fmt.Sprintf("%s %s", split[10], split[11])
	bs.giveValue(b, recLowID, recHighID)
	return true, nil
}

func (bs bots) getPart2Solution() int {
	out0 := bs["output 0"].vals[0]
	out1 := bs["output 1"].vals[0]
	out2 := bs["output 2"].vals[0]
	return out0 * out1 * out2
}

func findSolution(input []string, expectedChips []int) (string, int, error) {
	bots := bots{}
	i := 0
	var part1 string
	// Keep running until we have run all of the instructions
	for {
		for _, bot := range bots {
			if slice.IntSlicesAreEqual(bot.vals, expectedChips) {
				part1 = strings.Split(bot.id, " ")[1]
			}
		}
		// If all instructions are run then there must be values in outputs 0, 1 and 2
		if len(input) == 0 {
			return part1, bots.getPart2Solution(), nil
		}
		line := input[i]
		if split := strings.Split(line, " "); split[0] == "value" {
			err := bots.handleValueLine(split)
			if err != nil {
				return "", -1, err
			}
			// Once the value has been allocated we remove this instruction and start again
			input = append(input[:i], input[i+1:]...)
			i = 0
		} else {
			didGiveValues, err := bots.handleGiveLine(split)
			if err != nil {
				return "", -1, err
			}
			if didGiveValues {
				// Once the value has been allocated we remove this instruction and start again,
				// otherwise we move onto the next instruction
				input = append(input[:i], input[i+1:]...)
				i = 0
			} else {
				i++
			}
		}
	}
}

func main() {
	input := file.Read()
	part1, part2, err := findSolution(input, []int{17, 61})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
