package main

import (
	"Advent-of-Code/file"
	"fmt"
	"strconv"
	"strings"
)

type registers map[string]int

func (r registers) cpy(from, to string) error {
	if val, ok := r[from]; ok {
		r[to] = val
	} else {
		conv, err := strconv.Atoi(from)
		if err != nil {
			return err
		}
		r[to] = conv
	}
	return nil
}

func (r registers) inc(reg string) {
	r[reg]++
}

func (r registers) dec(reg string) {
	r[reg]--
}

func (r registers) jnz(reg, jump string, i *int) error {
	if v, ok := r[reg]; ok {
		if v == 0 {
			*i++
			return nil
		}
	} else {
		v, err := strconv.Atoi(reg)
		if err != nil {
			return err
		}
		if v == 0 {
			*i++
			return nil
		}
	}

	v, err := strconv.Atoi(jump)
	if err != nil {
		return err
	}
	*i += v
	return nil
}

func (r registers) followInstruction(inst string, i *int) error {
	split := strings.Split(inst, " ")
	var err error
	switch split[0] {
	case "cpy":
		err = r.cpy(split[1], split[2])
	case "inc":
		r.inc(split[1])
	case "dec":
		r.dec(split[1])
	case "jnz":
		err = r.jnz(split[1], split[2], i)
		// we increment i before exiting so revert that here since we already incremented
		*i--
	default:
		return fmt.Errorf("invalid instruction: %v", inst)
	}
	*i++
	return err
}

func (r registers) findSolution(instructions []string, i *int) (int, error) {
	length := len(instructions)
	for {
		if *i > length-1 {
			return r["a"], nil
		}
		err := r.followInstruction(instructions[*i], i)
		if err != nil {
			return -1, err
		}
	}
}

func findSolutions(instructions []string) (int, int, error) {
	reg1 := registers{"a": 0, "b": 0, "c": 0, "d": 0}
	i := 0
	part1, err := reg1.findSolution(instructions, &i)
	if err != nil {
		return -1, -1, err
	}

	reg2 := registers{"a": 0, "b": 0, "c": 1, "d": 0}
	i = 0
	part2, err := reg2.findSolution(instructions, &i)
	return part1, part2, err
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
