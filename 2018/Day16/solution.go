package main

import (
	"Advent-of-Code"
	"fmt"
	"regexp"
)

var (
	inputs  = make(map[int][]int)
	process = make(map[int][]int)
	outputs = make(map[int][]int)
)

var re = regexp.MustCompile("\\d+")

type sample struct {
	input        []int
	instructions []int
	output       [4]int
}

func populateSamplesPart1(lines []string) (samples []sample) {
	for i := 0; i < len(lines); {
		o := helpers.StringSliceToIntSlice(re.FindAllString(lines[i+2], -1))
		var output [4]int
		copy(output[:], o)
		s := sample{
			input:        helpers.StringSliceToIntSlice(re.FindAllString(lines[i], -1)),
			instructions: helpers.StringSliceToIntSlice(re.FindAllString(lines[i+1], -1)),
			output:       output,
		}
		samples = append(samples, s)
		i += 4
	}
	return samples
}

func addr(input []int, instructions []int) (output [4]int) {
	copy(output[:], input)
	output[instructions[3]] = input[instructions[1]] + input[instructions[2]]
	return
}

func addi(input []int, instructions []int) (output [4]int) {
	copy(output[:], input)
	output[instructions[3]] = input[instructions[1]] + instructions[2]
	return
}

func mulr(input []int, instructions []int) (output [4]int) {
	copy(output[:], input)
	output[instructions[3]] = input[instructions[1]] * input[instructions[2]]
	return
}

func muli(input []int, instructions []int) (output [4]int) {
	copy(output[:], input)
	output[instructions[3]] = input[instructions[1]] * instructions[2]
	return
}

func banr(input []int, instructions []int) (output [4]int) {
	copy(output[:], input)
	output[instructions[3]] = input[instructions[1]] & input[instructions[2]]
	return
}

func bani(input []int, instructions []int) (output [4]int) {
	copy(output[:], input)
	output[instructions[3]] = input[instructions[1]] & instructions[2]
	return
}

func borr(input []int, instructions []int) (output [4]int) {
	copy(output[:], input)
	output[instructions[3]] = input[instructions[1]] | input[instructions[2]]
	return
}

func bori(input []int, instructions []int) (output [4]int) {
	copy(output[:], input)
	output[instructions[3]] = input[instructions[1]] | instructions[2]
	return
}

func setr(input []int, instructions []int) (output [4]int) {
	copy(output[:], input)
	output[instructions[3]] = input[instructions[1]]
	return
}

func seti(input []int, instructions []int) (output [4]int) {
	copy(output[:], input)
	output[instructions[3]] = instructions[1]
	return
}

func gtir(input []int, instructions []int) (output [4]int) {
	copy(output[:], input)
	output[instructions[3]] = 0
	if instructions[1] > input[instructions[2]] {
		output[instructions[3]] = 1
	}
	return
}

func gtri(input []int, instructions []int) (output [4]int) {
	copy(output[:], input)
	output[instructions[3]] = 0
	if input[instructions[1]] > instructions[2] {
		output[instructions[3]] = 1
	}
	return
}

func gtrr(input []int, instructions []int) (output [4]int) {
	copy(output[:], input)
	output[instructions[3]] = 0
	if input[instructions[1]] > input[instructions[2]] {
		output[instructions[3]] = 1
	}
	return
}

func eqir(input []int, instructions []int) (output [4]int) {
	copy(output[:], input)
	output[instructions[3]] = 0
	if instructions[1] == input[instructions[2]] {
		output[instructions[3]] = 1
	}
	return
}

func eqri(input []int, instructions []int) (output [4]int) {
	copy(output[:], input)
	output[instructions[3]] = 0
	if input[instructions[1]] == instructions[2] {
		output[instructions[3]] = 1
	}
	return
}

func eqrr(input []int, instructions []int) (output [4]int) {
	copy(output[:], input)
	output[instructions[3]] = 0
	if input[instructions[1]] == input[instructions[2]] {
		output[instructions[3]] = 1
	}
	return
}

var opcodes = [16]func([]int, []int) [4]int{
	addr,
	addi,
	mulr,
	banr,
	muli,
	bani,
	borr,
	bori,
	setr,
	seti,
	gtir,
	gtri,
	gtrr,
	eqir,
	eqri,
	eqrr,
}

func compareFuncOutputs(s sample, f func([]int, []int) [4]int) bool {
	return f(s.input, s.instructions) == s.output
}

func assignOpcodeNumbers(samples []sample) (numbersToOpcodes [16]func([]int, []int) [4]int) {
	for {
		for _, f := range opcodes {
			var matches []int
			for i := 0; i < 16; i++ {
				for _, s := range samples {
					if s.instructions[0] == i && numbersToOpcodes[i] == nil {
						if compareFuncOutputs(s, f) {
							inMatches := false
							for _, num := range matches {
								if num == i {
									inMatches = true
								}
							}
							if !inMatches {
								matches = append(matches, i)
							}
						}
					}
				}
			}
			if len(matches) == 1 {
				numbersToOpcodes[matches[0]] = f
			}
		}
		numbersToOpcodesFull := true
		for _, v := range numbersToOpcodes {
			if v == nil {
				numbersToOpcodesFull = false
			}
		}
		if numbersToOpcodesFull {
			return numbersToOpcodes
		}
	}
}

func populatePrograms(lines []string) (programs [][]int) {
	for _, line := range lines {
		o := helpers.StringSliceToIntSlice(re.FindAllString(line, -1))
		output := make([]int, 4)
		copy(output[:], o)
		programs = append(programs, output)
	}
	return programs
}

func main() {
	lines := helpers.ReadFile()
	samples := populateSamplesPart1(lines[:3352])
	var total int
	for _, s := range samples {
		var matchingOutputs int
		for _, f := range opcodes {
			if compareFuncOutputs(s, f) {
				matchingOutputs++
			}
		}
		if matchingOutputs >= 3 {
			total++
		}
	}
	fmt.Println("Part1:", total)
	numbersToOpcodes := assignOpcodeNumbers(samples)
	fmt.Println(numbersToOpcodes)
	programs := populatePrograms(lines[3354:])
	input := []int{0, 0, 0, 0}
	var output [4]int
	for _, p := range programs {
		output := numbersToOpcodes[p[0]](input, p)
		fmt.Println(p, output)
		input = output[:]
	}
	fmt.Println(output)
}
