package main

import (
	"Advent-of-Code/file"
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	dir string
	val int
}

type position struct {
	X   int
	Y   int
	aim int
}

func evaluateLine(line string) (*instruction, error) {
	sp := strings.Split(line, " ")
	val, err := strconv.Atoi(sp[1])
	if err != nil {
		return nil, err
	}
	return &instruction{
		dir: sp[0],
		val: val,
	}, nil
}

func validateInstruction(dir string) error {
	if dir == "forward" || dir == "up" || dir == "down" {
		return nil
	}
	return fmt.Errorf("followInstruction: invalid direction from instruction %s", dir)
}

func (co *position) followInstruction(inst instruction, part2 bool) {
	switch inst.dir {
	case "forward":
		co.X += inst.val
		if part2 {
			co.Y += (co.aim * inst.val)
		}
	case "up":
		if part2 {
			co.aim -= inst.val
		} else {
			co.Y -= inst.val
		}
	case "down":
		if part2 {
			co.aim += inst.val
		} else {
			co.Y += inst.val
		}
	}
}

func findSolution(input []string, part2 bool) (int, error) {
	co := position{}
	for _, line := range input {
		inst, err := evaluateLine(line)
		if err != nil {
			return -1, err
		}
		err = validateInstruction(inst.dir)
		if err != nil {
			return -1, err
		}
		co.followInstruction(*inst, part2)
	}
	return co.X * co.Y, nil
}

func main() {
	input := file.Read()
	part1, err := findSolution(input, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	part2, err := findSolution(input, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 2:", part2)
}
