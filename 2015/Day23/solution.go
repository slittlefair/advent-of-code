package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"strconv"
	"strings"
)

type Registers map[string]int

func (r Registers) hlfInstruction(register string) {
	r[register] /= 2
}

func (r Registers) tplInstruction(register string) {
	r[register] *= 3
}

func (r Registers) incInstruction(register string) {
	r[register]++
}

func (r Registers) jmpInstruction(offset string) (int, error) {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return -1, err
	}
	return offsetInt, nil
}

func (r Registers) jieInstruction(register, offset string) (int, error) {
	if r[string(register[0])]%2 != 0 {
		return 1, nil
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return -1, err
	}
	return offsetInt, nil
}

func (r Registers) jioInstruction(register, offset string) (int, error) {
	if r[string(register[0])] != 1 {
		return 1, nil
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return -1, err
	}
	return offsetInt, nil
}

func (r Registers) FollowInstruction(instruction string) (int, error) {
	split := strings.Split(instruction, " ")
	switch split[0] {
	case "hlf":
		r.hlfInstruction(split[1])
		return 1, nil
	case "tpl":
		r.tplInstruction(split[1])
		return 1, nil
	case "inc":
		r.incInstruction(split[1])
		return 1, nil
	case "jmp":
		return r.jmpInstruction(split[1])
	case "jie":
		return r.jieInstruction(split[1], split[2])
	case "jio":
		return r.jioInstruction(split[1], split[2])
	default:
		return -1, fmt.Errorf("could not find valid instruction for %s", instruction)
	}
}

func (r Registers) RunInstructions(instructions []string) error {
	i := 0
	for i < len(instructions) {
		offset, err := r.FollowInstruction(instructions[i])
		if err != nil {
			return err
		}
		i += offset
	}
	return nil
}

func main() {
	input := helpers.ReadFile()
	registers := Registers{
		"a": 0,
		"b": 0,
	}
	err := registers.RunInstructions(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", registers["b"])
	registers["a"] = 1
	registers["b"] = 0
	err = registers.RunInstructions(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 2:", registers["b"])
}
