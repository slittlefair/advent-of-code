package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/maths"
	"Advent-of-Code/timer"
	"fmt"
	"maps"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/goforj/godump"
)

type Machine struct {
	Lights   []bool
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
		lights := strings.SplitSeq(lightMatches[1], "")
		for b := range lights {
			mach.Lights = append(mach.Lights, b == "#")
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

type Cache map[string]struct{}

func makeCacheKey(lights []bool) (cacheKey string) {
	for _, l := range lights {
		if l {
			cacheKey += "t"
		} else {
			cacheKey += "f"
		}
	}
	return cacheKey
}

func (m *Machine) switchLights(lights []bool, buttons int, sequence []int, cache Cache) {
	if m.Success == 1 {
		return
	}
	lightCopy := make([]bool, len(lights))
	copy(lightCopy, lights)
	for _, b := range m.Buttons[buttons] {
		lightCopy[b] = !lightCopy[b]
	}
	if slices.Equal(lightCopy, m.Lights) {
		m.Success = len(sequence)
		godump.Dump(sequence)
		return
	}
	cacheKey := makeCacheKey(lightCopy)
	cacheCopy := maps.Clone(cache)
	if _, ok := cacheCopy[cacheKey]; ok {
		return
	}
	cacheCopy[cacheKey] = struct{}{}
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

func buttonNeeded(button []int, lights []bool) bool {
	for i, l := range lights {
		if l && slices.Contains(button, i) {
			return true
		}
	}
	return false
}

func findSolutions(input []string) (part1 int, part2 int) {
	machines := parseInput(input)
	machineTimes := make([]struct {
		t time.Duration
		i int
		s int
	}, len(machines))
	for i, m := range machines {
		t := time.Now()
		for b, btns := range m.Buttons {
			if !buttonNeeded(btns, m.Lights) {
				// fmt.Println("dont need", btns)
				continue
			}
			lights := make([]bool, len(m.Lights))
			cache := make(map[string]struct{})
			cache[makeCacheKey(m.Lights)] = struct{}{}
			m.switchLights(lights, b, []int{b}, cache)
		}
		machineTimes[i] = struct {
			t time.Duration
			i int
			s int
		}{t: time.Since(t), i: i, s: m.Success}
		part1 += m.Success
	}
	sort.Slice(machineTimes, func(i, j int) bool {
		return machineTimes[i].t > machineTimes[j].t
	})
	fmt.Println("Longest:")
	max := maths.Min(10, len(machineTimes))
	for i := range max {
		m := machineTimes[i]
		fmt.Printf("Machine %v: %v (%v)\n", m.i+1, m.t, m.s)
	}
	return part1, part2
}

func main() {
	f := time.Now()
	input := file.Read()
	part1, part2 := findSolutions(input)
	// TODO speed up part 1. Takes ~20seconds alone, machines 103 and 117 take 5s, so work out how
	// to minimise these
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
	timer.Track(f)
}
