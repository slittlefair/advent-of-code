package main

import (
	helpers "Advent-of-Code"
	"fmt"
)

type Lanternfish []int

func createLanternfish(input []int) Lanternfish {
	lf := make(Lanternfish, 9)
	for _, i := range input {
		lf[i]++
	}
	return lf
}

func (lf Lanternfish) iterate() Lanternfish {
	nextLF := make(Lanternfish, 9)
	for i := 1; i < len(lf); i++ {
		nextLF[i-1] = lf[i]
	}
	nextLF[6] += lf[0]
	nextLF[8] = lf[0]
	return nextLF
}

func (lf Lanternfish) count() int {
	sum := 0
	for _, v := range lf {
		sum += v
	}
	return sum
}

func (lf Lanternfish) findSolution() (int, int) {
	for i := 0; i < 80; i++ {
		lf = lf.iterate()
	}
	part1 := lf.count()
	for i := 80; i < 256; i++ {
		lf = lf.iterate()
	}
	part2 := lf.count()
	return part1, part2
}

func main() {
	input, err := helpers.ReadFileSingleLineAsInts()
	if err != nil {
		fmt.Println(err)
		return
	}
	lf := createLanternfish(input)
	part1, part2 := lf.findSolution()
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
