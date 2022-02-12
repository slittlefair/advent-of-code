package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/slice"
	"fmt"
	"regexp"
	"strconv"
)

var reWord = regexp.MustCompile("[a-z]+")
var reNum = regexp.MustCompile(`\d+`)

type instructions struct {
	inst      string
	registers []int
}

var allInstructions []instructions

func addr(input []int, instructions []int) (output []int) {
	output = input
	output[instructions[3]] = input[instructions[1]] + input[instructions[2]]
	return
}

func addi(input []int, instructions []int) (output []int) {
	output = input
	output[instructions[3]] = input[instructions[1]] + instructions[2]
	return
}

func mulr(input []int, instructions []int) (output []int) {
	output = input
	output[instructions[3]] = input[instructions[1]] * input[instructions[2]]
	return
}

func muli(input []int, instructions []int) (output []int) {
	output = input
	output[instructions[3]] = input[instructions[1]] * instructions[2]
	return
}

func banr(input []int, instructions []int) (output []int) {
	output = input
	output[instructions[3]] = input[instructions[1]] & input[instructions[2]]
	return
}

func bani(input []int, instructions []int) (output []int) {
	output = input
	output[instructions[3]] = input[instructions[1]] & instructions[2]
	return
}

func borr(input []int, instructions []int) (output []int) {
	output = input
	output[instructions[3]] = input[instructions[1]] | input[instructions[2]]
	return
}

func bori(input []int, instructions []int) (output []int) {
	output = input
	output[instructions[3]] = input[instructions[1]] | instructions[2]
	return
}

func setr(input []int, instructions []int) (output []int) {
	output = input
	output[instructions[3]] = input[instructions[1]]
	return
}

func seti(input []int, instructions []int) (output []int) {
	output = input
	output[instructions[3]] = instructions[1]
	return
}

func gtir(input []int, instructions []int) (output []int) {
	output = input
	if instructions[1] > input[instructions[2]] {
		output[instructions[3]] = 1
	} else {
		output[instructions[3]] = 0
	}
	return
}

func gtri(input []int, instructions []int) (output []int) {
	output = input
	if input[instructions[1]] > instructions[2] {
		output[instructions[3]] = 1
	} else {
		output[instructions[3]] = 0
	}
	return
}

func gtrr(input []int, instructions []int) (output []int) {
	output = input
	if input[instructions[1]] > input[instructions[2]] {
		output[instructions[3]] = 1
	} else {
		output[instructions[3]] = 0
	}
	return
}

func eqir(input []int, instructions []int) (output []int) {
	output = input
	if instructions[1] == input[instructions[2]] {
		output[instructions[3]] = 1
	} else {
		output[instructions[3]] = 0
	}
	return
}

func eqri(input []int, instructions []int) (output []int) {
	output = input
	if input[instructions[1]] == instructions[2] {
		output[instructions[3]] = 1
	} else {
		output[instructions[3]] = 0
	}
	return
}

func eqrr(input []int, instructions []int) (output []int) {
	output = input
	if input[instructions[1]] == input[instructions[2]] {
		output[instructions[3]] = 1
	} else {
		output[instructions[3]] = 0
	}
	return
}

var functionsMap = map[string]func([]int, []int) []int{
	"addi": addi,
	"addr": addr,
	"mulr": mulr,
	"muli": muli,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,
	"eqri": eqri,
	"eqir": eqir,
	"eqrr": eqrr,
}

var winners = make(map[int]bool)

func getValues(ipString []string) error {
	ip := 0
	ipRegister, err := strconv.Atoi(ipString[0])
	if err != nil {
		return err
	}
	registers := []int{10, 0, 0, 0, 0, 0}
	for ip >= 0 && ip < len(allInstructions) {
		registers[ipRegister] = ip
		inst := allInstructions[ip]
		if ip == 28 {
			if _, ok := winners[registers[2]]; !ok {
				winners[registers[2]] = true
				fmt.Println(registers[2])
				if len(winners)%500 == 0 {
					fmt.Println(len(winners))
				}
			} else {
				return nil
			}
		}
		registers = functionsMap[inst.inst](registers, append([]int{0}, inst.registers...))
		ip = registers[ipRegister]
		ip++
	}
	return nil
}

func main() {
	lines := file.Read()
	ipString := reNum.FindAllString(lines[0], -1)

	for i := 1; i < len(lines); i++ {
		instText := reWord.FindAllString(lines[i], -1)
		instNums := reNum.FindAllString(lines[i], -1)
		registers, err := slice.StringSliceToIntSlice(instNums)
		if err != nil {
			fmt.Println(err)
			return
		}
		inst := instructions{
			inst:      instText[0],
			registers: registers,
		}
		allInstructions = append(allInstructions, inst)
	}

	if err := getValues(ipString); err != nil {
		fmt.Println(err)
	}
}
