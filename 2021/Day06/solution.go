package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"strconv"
	"strings"
)

type Lanternfish []int

func createLanternfish(input []string) (Lanternfish, error) {
	if l := len(input); l != 1 {
		return nil, fmt.Errorf("error getting input, expected 1 line, got %d: %v", l, input)
	}
	lf := make(Lanternfish, 9)
	ints := strings.Split(input[0], ",")
	for _, i := range ints {
		n, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		lf[n]++
	}
	return lf, nil
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
	input := helpers.ReadFile()
	lf, err := createLanternfish(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	part1, part2 := lf.findSolution()
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
