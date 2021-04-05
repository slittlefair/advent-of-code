package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"strconv"
	"strings"
)

type Instructions struct {
	instruction string
	value       int
}

type Programme struct {
	foundSolution bool
	instructions  []Instructions
}

func parseProgramme(entries []string) (*Programme, error) {
	prog := &Programme{}

	for _, entry := range entries {
		split := strings.Split(entry, " ")
		value, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		prog.instructions = append(prog.instructions, Instructions{
			instruction: split[0],
			value:       value,
		})
	}

	return prog, nil
}

func (p *Programme) runProgramme(tweakAtIndex int) int {
	index := 0
	accumulator := 0
	instructionsRun := map[int]bool{}
	for {
		inst := p.instructions[index]
		if inst.instruction == "nop" {
			if index == tweakAtIndex {
				index += inst.value
			} else {
				index++
			}
		} else if inst.instruction == "acc" {
			accumulator += inst.value
			index++
		} else if inst.instruction == "jmp" {
			if index == tweakAtIndex {
				index++
			} else {
				index += inst.value
			}
		}
		if _, ok := instructionsRun[index]; ok {
			return accumulator
		} else if index >= len(p.instructions) {
			p.foundSolution = true
			return accumulator
		}
		instructionsRun[index] = true
	}
}

func main() {
	entries := helpers.ReadFile()
	prog, err := parseProgramme(entries)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", prog.runProgramme(-1))
	part2Accumulator := 0
	for i := range prog.instructions {
		part2Accumulator = prog.runProgramme(i)
		if prog.foundSolution {
			break
		}
	}
	fmt.Println("Part 2:", part2Accumulator)
}
