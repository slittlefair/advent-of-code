package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"sort"
	"strconv"
)

type EggnogContainers struct {
	WantedTotal int
	Ways        map[int]int
}

func parseInput(input []string) ([]int, error) {
	inputInt := make([]int, len(input))
	for i, v := range input {
		conv, err := strconv.Atoi(v)
		if err != nil {
			return nil, nil
		}
		inputInt[i] = conv
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inputInt)))
	return inputInt, nil
}

func (ec *EggnogContainers) FindContainers(remainingContainers []int, totalCapacity int, levels int) {
	if totalCapacity == ec.WantedTotal {
		if val, ok := ec.Ways[levels]; !ok {
			ec.Ways[levels] = 1
		} else {
			ec.Ways[levels] = val + 1
		}
	}
	for i := 0; i < len(remainingContainers); i++ {
		ec.FindContainers(remainingContainers[i+1:], totalCapacity+remainingContainers[i], levels+1)
	}
}

func (ec EggnogContainers) CountPermutations() int {
	count := 0
	for _, v := range ec.Ways {
		count += v
	}
	return count
}

func (ec EggnogContainers) CountSmallestContainersPermutations() int {
	smallestPermutations := helpers.Infinty
	count := 0
	for numContainers, freq := range ec.Ways {
		if numContainers < smallestPermutations {
			smallestPermutations = numContainers
			count = freq
		}
	}
	return count
}

func main() {
	input := helpers.ReadFile()
	containers, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	ec := &EggnogContainers{
		WantedTotal: 150,
		Ways:        make(map[int]int),
	}
	ec.FindContainers(containers, 0, 0)
	fmt.Println("Part 1:", ec.CountPermutations())
	fmt.Println("Part 2:", ec.CountSmallestContainersPermutations())
}
