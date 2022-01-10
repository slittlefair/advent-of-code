package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"regexp"
)

var reWord = regexp.MustCompile("[a-z]+")
var reNum = regexp.MustCompile("\\d+")

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

func main() {
	lines := utils.ReadFile()
	ipString := reNum.FindAllString(lines[0], -1)
	// fmt.Println(ip)

	for i := 1; i < len(lines); i++ {
		instText := reWord.FindAllString(lines[i], -1)
		instNums := reNum.FindAllString(lines[i], -1)
		inst := instructions{
			inst:      instText[0],
			registers: utils.StringSliceToIntSlice(instNums),
		}
		allInstructions = append(allInstructions, inst)
	}

	ip := 0
	ipRegister := utils.StringToInt(ipString[0])
	registers := []int{1, 0, 0, 0, 0, 0}
	for ip >= 0 && ip < len(allInstructions) {
		registers[ipRegister] = ip
		inst := allInstructions[ip]
		fmt.Println(ip, ipRegister, allInstructions[ip], registers)
		registers = functionsMap[inst.inst](registers, append([]int{0}, inst.registers...))
		ip = registers[ipRegister]
		ip++
	}
	fmt.Println(registers)
}
