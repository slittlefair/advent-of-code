package main

import (
	"Advent-of-Code/file"
	"fmt"
	"math"
	"strconv"
)

func deliverPresentsPart1(target int) (int, error) {
	for i := 1; i <= target; i++ {
		sum := 0
		for j := 1; j < int(math.Sqrt(float64(i)))+1; j++ {
			if i%j == 0 {
				sum += j
				sum += (i / j)
			}
		}
		if sum*10 >= target {
			return i, nil
		}
	}
	return -1, fmt.Errorf("something went wrong, could not find solution for target %d", target)
}

func deliverPresentsPart2(target int) (int, error) {
	for i := 1; i <= target; i++ {
		sum := 0
		for j := 1; j < int(math.Sqrt(float64(i)))+1; j++ {
			if i%j == 0 {
				if j <= 50 {
					sum += (i / j)
				}
				if i/j < 50 {
					sum += i
				}
			}
		}
		if sum*11 >= target {
			return i, nil
		}
	}
	return -1, fmt.Errorf("something went wrong, could not find solution for target %d", target)
}

func main() {
	input := file.Read()[0]
	target, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	house, err := deliverPresentsPart1(target)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", house)
	house, err = deliverPresentsPart2(target)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 2:", house)
}
