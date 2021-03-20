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
	accumulator   int
	foundSolution bool
	instructions  []Instructions
}

func parseProgramme(entries []string) (*Programme, error) {
	prog := &Programme{}

	for _, entry := range entries {
		split := strings.Split(entry, " ")
		instruction := split[0]
		value, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		prog.instructions = append(prog.instructions, Instructions{
			instruction: instruction,
			value:       value,
		})
	}

	return prog, nil
}

func (p *Programme) runProgramme(tweakAtIndex int) {
	index := 0
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
			p.accumulator += inst.value
			index++
		} else if inst.instruction == "jmp" {
			if index == tweakAtIndex {
				index++
			} else {
				index += inst.value
			}
		}
		if _, ok := instructionsRun[index]; ok {
			return
		} else if index >= len(p.instructions) {
			p.foundSolution = true
			return
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
	prog.runProgramme(-1)
	fmt.Println("Part 1:", prog.accumulator)
	for i := range prog.instructions {
		prog.accumulator = 0
		prog.runProgramme(i)
		if prog.foundSolution {
			break
		}
	}
	fmt.Println("Part 2:", prog.accumulator)
}
