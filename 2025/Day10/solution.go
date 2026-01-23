package main

import (
	"Advent-of-Code/file"
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Machine struct {
	Lights   uint16
	Buttons  [][]int
	Joltages []int
	part1    int
}

var reLights = regexp.MustCompile(`\[([.#]+)\]`)
var reButtons = regexp.MustCompile(`\((\d+(?:,\d+)*)\)`)
var reJoltages = regexp.MustCompile(`\{(\d+(?:,\d+)*)\}`)

func parseInput(input []string) []*Machine {
	machines := make([]*Machine, len(input))
	for i, line := range input {
		lightMatches := reLights.FindStringSubmatch(line)
		buttonMatches := reButtons.FindAllStringSubmatch(line, -1)
		joltageMatches := reJoltages.FindStringSubmatch(line)

		mach := &Machine{}
		lights := lightMatches[1]
		for i := len(lights) - 1; i >= 0; i-- {
			mach.Lights <<= 1
			if lights[i] == '#' {
				mach.Lights |= 1
			}
		}

		for _, m := range buttonMatches {
			btn := strings.Split(m[1], ",")
			buttons := []int{}
			for _, b := range btn {
				n, _ := strconv.Atoi(b)
				buttons = append(buttons, n)
			}
			mach.Buttons = append(mach.Buttons, buttons)
		}

		joltages := strings.SplitSeq(joltageMatches[1], ",")
		for j := range joltages {
			n, _ := strconv.Atoi(j)
			mach.Joltages = append(mach.Joltages, n)
		}

		machines[i] = mach
	}
	return machines
}

type Cache map[uint16]struct{}

func (m *Machine) switchLights(lights uint16, buttons int, sequence []int, cache Cache, maxLevel int) {
	if len(sequence) >= maxLevel {
		return
	}
	if m.part1 == 1 {
		return
	}
	for _, v := range m.Buttons[buttons] {
		lights ^= 1 << v
	}
	if lights == m.Lights {
		m.part1 = len(sequence)
		return
	}
	cacheCopy := maps.Clone(cache)
	if _, ok := cacheCopy[lights]; ok {
		return
	}
	cacheCopy[lights] = struct{}{}
	for b := range m.Buttons {
		if b == buttons {
			continue
		}
		if slices.Contains(sequence, b) {
			continue
		}
		m.switchLights(lights, b, append(sequence, b), cacheCopy, maxLevel)
	}
}

func findSolutions(input []string) (part1 int, part2 int) {
	machines := parseInput(input)
	for _, m := range machines {
		for i := 1; i <= len(m.Buttons); i++ {
			for b := range m.Buttons {
				if m.part1 != 0 {
					goto out
				}
				cache := make(map[uint16]struct{})
				cache[m.Lights] = struct{}{}
				m.switchLights(uint16(0), b, []int{b}, cache, i)
			}
		}
	out:
		part1 += m.part1
	}
	return part1, part2
}

func main() {
	input := file.Read()
	part1, part2 := findSolutions(input)
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}
