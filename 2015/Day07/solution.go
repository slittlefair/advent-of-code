package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"regexp"
	"strconv"
)

type Wires map[string]int

func (w Wires) doBitwiseAND(identifiers []string, nums []string) {
	if _, ok := w[identifiers[0]]; !ok {
		return
	}
	if len(nums) > 0 {
		n, _ := strconv.Atoi(nums[0])
		w[identifiers[1]] = w[identifiers[0]] & n
		return
	}
	if _, ok := w[identifiers[1]]; !ok {
		return
	}
	w[identifiers[2]] = w[identifiers[0]] & w[identifiers[1]]
}

func (w Wires) doBitwiseOR(identifiers []string, nums []string) {
	if _, ok := w[identifiers[0]]; !ok {
		return
	}
	if len(nums) > 0 {
		n, _ := strconv.Atoi(nums[0])
		w[identifiers[1]] = w[identifiers[0]] | n
		return
	}
	if _, ok := w[identifiers[1]]; !ok {
		return
	}
	w[identifiers[2]] = w[identifiers[0]] | w[identifiers[1]]
}

func (w Wires) doBitwiseNOT(identifiers []string) {
	if _, ok := w[identifiers[0]]; !ok {
		return
	}
	w[identifiers[1]] = 65535 ^ w[identifiers[0]]
}

func (w Wires) doBitwiseLSHIFT(identifiers []string, nums []string) {
	if _, ok := w[identifiers[0]]; !ok {
		return
	}
	n, _ := strconv.Atoi(nums[0])
	w[identifiers[1]] = w[identifiers[0]] << n
}

func (w Wires) doBitwiseRSHIFT(identifiers []string, nums []string) {
	if _, ok := w[identifiers[0]]; !ok {
		return
	}
	n, _ := strconv.Atoi(nums[0])
	w[identifiers[1]] = w[identifiers[0]] >> n
}

func (w Wires) doASSIGN(identifiers []string, nums []string) {
	if len(nums) > 0 {
		if _, ok := w[identifiers[0]]; ok {
			return
		}
		n, _ := strconv.Atoi(nums[0])
		w[identifiers[0]] = n
		return
	}
	if _, ok := w[identifiers[0]]; !ok {
		return
	}
	w[identifiers[1]] = w[identifiers[0]]
}

func (w Wires) followInstructions(instructions []string) {
	identifierRe := regexp.MustCompile(`[a-z]+`)
	signalRe := regexp.MustCompile(`[A-Z]+`)
	intRe := regexp.MustCompile(`\d+`)
	for len(w) != len(instructions) {
		for _, inst := range instructions {
			identifiers := identifierRe.FindAllString(inst, -1)
			signal := signalRe.FindString(inst)
			if signal == "NOT" {
				w.doBitwiseNOT(identifiers)
				continue
			}
			nums := intRe.FindAllString(inst, -1)
			if signal == "AND" {
				w.doBitwiseAND(identifiers, nums)
				continue
			}
			if signal == "OR" {
				w.doBitwiseOR(identifiers, nums)
				continue
			}
			if signal == "LSHIFT" {
				w.doBitwiseLSHIFT(identifiers, nums)
				continue
			}
			if signal == "RSHIFT" {
				w.doBitwiseRSHIFT(identifiers, nums)
				continue
			}
			w.doASSIGN(identifiers, nums)
		}
	}
}

func main() {
	input := utils.ReadFile()
	wires := Wires{}
	wires.followInstructions(input)
	fmt.Println("Part 1:", wires["a"])

	wires = Wires{
		"b": wires["a"],
	}
	wires.followInstructions(input)
	fmt.Println("Part 2:", wires["a"])
}
