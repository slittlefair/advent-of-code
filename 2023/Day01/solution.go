package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/maths"
	"fmt"
	"strconv"
	"strings"
)

var strs = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

type Num struct {
	idx    int
	val    int
	strIdx int
	strVal int
}

func processLine(line string, tens, units *Num) {
	for i, num := range strs {
		conv := strconv.Itoa(i)

		firstIdxNum := strings.Index(line, conv)
		if firstIdxNum != -1 && firstIdxNum < tens.idx {
			tens.idx = firstIdxNum
			tens.val = i
		}

		firstIdxString := strings.Index(line, num)
		if firstIdxString != -1 && firstIdxString < tens.strIdx {
			tens.strIdx = firstIdxString
			tens.strVal = i
		}

		lastIdxNum := strings.LastIndex(line, conv)
		if lastIdxNum != -1 && lastIdxNum > units.idx {
			units.idx = lastIdxNum
			units.val = i
		}

		lastIdxString := strings.LastIndex(line, num)
		if lastIdxString != -1 && lastIdxString > units.strIdx {
			units.strIdx = lastIdxString
			units.strVal = i
		}
	}
}

func findSolution(input []string) (int, int) {
	part1 := 0
	part2 := 0
	for _, line := range input {
		tens := &Num{
			idx:    maths.Infinity,
			strIdx: maths.Infinity,
		}
		units := &Num{
			idx:    -1,
			strIdx: -1,
		}

		processLine(line, tens, units)

		part1 += tens.val*10 + units.val

		if tens.idx < tens.strIdx {
			tens.strVal = tens.val
		}
		if units.idx > units.strIdx {
			units.strVal = units.val
		}
		part2 += tens.strVal*10 + units.strVal
	}
	return part1, part2
}

func main() {
	input := file.Read()
	part1, part2 := findSolution(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
