package main

import (
	"Advent-of-Code/file"
	"errors"
	"fmt"
	"sort"
)

type Numbers []int

func (n Numbers) cyclePrevNumbers(preambleLength int, i int) (bool, int) {
	for j := i - preambleLength; j < i; j++ {
		for k := i - preambleLength; k < i; k++ {
			if j != k && n[j]+n[k] == n[i] {
				return false, -1
			}
		}
	}
	return true, n[i]
}

func (n Numbers) part1(preambleLength int) (int, error) {
	for i := preambleLength; i < len(n); i++ {
		solved, solValue := n.cyclePrevNumbers(preambleLength, i)
		if solved {
			return solValue, nil
		}
	}
	return -1, errors.New("could not find solution to part 1")
}

func (n Numbers) getSumNumbers(part1Sol int) ([]int, error) {
	for i := 0; i < len(n); i++ {
		count := 0
		for j := i; j < len(n); j++ {
			count += n[j]
			if count > part1Sol {
				continue
			} else if count == part1Sol {
				return n[i : j+1], nil
			}
		}
	}
	return []int{}, errors.New("could not find solution to part 2")
}

func (n Numbers) part2(part1Sol int) (int, error) {
	numbers, err := n.getSumNumbers(part1Sol)
	if err != nil {
		return -1, err
	}
	sort.Ints(numbers)
	return numbers[0] + numbers[len(numbers)-1], nil
}

func main() {
	entries := file.ReadAsInts()
	numbers := Numbers{}
	for _, e := range entries {
		numbers = append(numbers, e)
	}
	preambleLength := 25

	part1Sol, err := numbers.part1(preambleLength)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1Sol)

	part2Sol, err := numbers.part2(part1Sol)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 2:", part2Sol)
}
