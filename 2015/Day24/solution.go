package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
)

type ValidCombos [][]int

func calculateQuantumEntanglement(g []int) int {
	if len(g) == 0 {
		return 0
	}
	product := 1
	for _, n := range g {
		product *= n
	}
	return product
}

func groupSum(packages []int) int {
	sum := 0
	for _, p := range packages {
		sum += p
	}
	return sum
}

func getLowestQuantumEntanglement(combos [][]int) (int, error) {
	lowestQE := utils.Infinity
	for _, c := range combos {
		if qe := calculateQuantumEntanglement(c); qe < lowestQE {
			lowestQE = qe
		}
	}
	if lowestQE == utils.Infinity {
		return -1, fmt.Errorf("could not find lowestQE of groups %v", combos)
	}
	return lowestQE, nil
}

func (vc *ValidCombos) iterate(remainingPackages, bucket []int, weight, maxLevel int) ([]int, []int) {
	var newRemainingPackages, newBucket []int
	for i, rp := range remainingPackages {
		newBucket := append(bucket, rp)
		newRemainingPackages := utils.Remove(remainingPackages, i)
		if len(newBucket) < maxLevel {
			vc.iterate(newRemainingPackages, newBucket, weight, maxLevel)
		}
		if sum := groupSum(newBucket); sum > weight {
			return newRemainingPackages, newBucket
		} else if sum == weight {
			*vc = append(*vc, newBucket)
			return newRemainingPackages, newBucket
		}
	}
	return newRemainingPackages, newBucket
}

func validPermutations(input []int, weight int) ([][]int, error) {
	validCombos := &ValidCombos{}
	for i := 1; i <= len(input); i++ {
		validCombos.iterate(input, []int{}, weight, i)
		if len(*validCombos) > 0 {
			return *validCombos, nil
		}
	}
	return nil, fmt.Errorf("could not get any valid combos for input: %v, weight: %d", input, weight)
}

func findSolutions(input []int, part1Sections, part2Sections int) (int, int, error) {
	weight1 := groupSum(input) / part1Sections
	combos1, err := validPermutations(input, weight1)
	if err != nil {
		return -1, -1, err
	}
	lowestQE1, err := getLowestQuantumEntanglement(combos1)
	if err != nil {
		return -1, -1, err
	}
	weight2 := groupSum(input) / part2Sections
	combos2, err := validPermutations(input, weight2)
	if err != nil {
		return -1, -1, err
	}
	// we can ignore this error since there would have been an error at part1 if this also errored
	lowestQE2, _ := getLowestQuantumEntanglement(combos2)
	return lowestQE1, lowestQE2, nil
}

func main() {
	input := utils.ReadFileAsInts()
	part1Sol, part2Sol, err := findSolutions(input, 3, 4)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1Sol)
	fmt.Println("Part 2:", part2Sol)
}
