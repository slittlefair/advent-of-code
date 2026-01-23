package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/timer"
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Machine struct {
	Lights   uint16
	Buttons  [][]int
	Joltages []int
	Success  int
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

func (m *Machine) switchLights(lights uint16, buttons int, sequence []int, cache Cache) {
	if m.Success == 1 {
		return
	}
	lightCopy := lights
	for _, v := range m.Buttons[buttons] {
		lightCopy ^= 1 << v
	}
	if lightCopy == m.Lights {
		m.Success = len(sequence)
		return
	}
	cacheCopy := maps.Clone(cache)
	if _, ok := cacheCopy[lightCopy]; ok {
		return
	}
	cacheCopy[lightCopy] = struct{}{}
	for b := range m.Buttons {
		if b == buttons {
			continue
		}
		if slices.Contains(sequence, b) {
			continue
		}
		if m.Success != 0 && len(sequence) == m.Success-1 {
			return
		}
		m.switchLights(lightCopy, b, append(sequence, b), cacheCopy)
	}
}

func findSolutions(input []string) (part1 int, part2 int) {
	machines := parseInput(input)
	for _, m := range machines {
		for b := range m.Buttons {
			cache := make(map[uint16]struct{})
			cache[m.Lights] = struct{}{}
			m.switchLights(uint16(0), b, []int{b}, cache)
		}
		part1 += m.Success
	}
	return part1, part2
}

func main() {
	f := time.Now()
	input := file.Read()
	part1, part2 := findSolutions(input)
	// TODO speed up part 1
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
	timer.Track(f)
}
