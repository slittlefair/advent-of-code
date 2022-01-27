package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"regexp"
)

type PolymerizationEquipment struct {
	pir map[string]string
	pf  map[string]int
	lf  map[string]int
}

func combineLetters(l1, l2 string) string {
	return fmt.Sprintf("%v%v", string(l1), string(l2))
}

func parseInput(input []string) (*PolymerizationEquipment, error) {
	pe := &PolymerizationEquipment{
		pir: make(map[string]string),
		pf:  make(map[string]int),
		lf:  make(map[string]int),
	}
	line := input[0]
	for i := 0; i < len(line); i++ {
		if i < len(line)-1 {
			pe.pf[combineLetters(string(line[i]), string(line[i+1]))]++
		}
		pe.lf[string(line[i])]++
	}
	re := regexp.MustCompile(`\w+`)
	for i := 2; i < len(input); i++ {
		matches := re.FindAllString(input[i], -1)
		if len(matches) != 2 {
			return nil, fmt.Errorf("error parsing input line, expected 2 strings for line %s", input[i])
		}
		pe.pir[matches[0]] = matches[1]
	}
	return pe, nil
}

func (pe *PolymerizationEquipment) followInstructions() {
	newPF := map[string]int{}
	for pair, freq := range pe.pf {
		newLetter := pe.pir[pair]
		newPF[combineLetters(string(pair[0]), newLetter)] += freq
		newPF[combineLetters(newLetter, string(pair[1]))] += freq
		pe.lf[newLetter] += freq
	}
	pe.pf = newPF
}

func (pe PolymerizationEquipment) getVal() int {
	minVal := utils.Infinity
	maxVal := 0
	for _, v := range pe.lf {
		minVal = utils.Min(minVal, v)
		maxVal = utils.Max(maxVal, v)
	}
	return maxVal - minVal
}

func findSolutions(input []string) (int, int, error) {
	pe, err := parseInput(input)
	if err != nil {
		return -1, -1, err
	}
	part1 := 0
	for i := 1; i <= 40; i++ {
		pe.followInstructions()
		if i == 10 {
			part1 = pe.getVal()
		}
	}
	return part1, pe.getVal(), nil
}

func main() {
	input := utils.ReadFile()
	part1, part2, err := findSolutions(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
