package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"math"
	"regexp"
)

type SignalPatterns map[string]struct{}

func parseInput(line string, re *regexp.Regexp) (SignalPatterns, []string, error) {
	matches := re.FindAllString(line, -1)
	if l := len(matches); l != 14 {
		return nil, nil, fmt.Errorf("wanted 14 matches, got %d: %v", l, line)
	}
	sp := make(SignalPatterns)
	for _, m := range matches[:10] {
		sp[m] = struct{}{}
	}
	return sp, matches[10:], nil
}

func countSimpleDigits(outputValues []string) int {
	sum := 0
	for _, v := range outputValues {
		if l := len(v); l == 2 || l == 3 || l == 4 || l == 7 {
			sum++
		}
	}
	return sum
}

func stringsShareParts(str1, str2 string, wantEqual bool) bool {
	if wantEqual && len(str1) != len(str2) {
		return false
	}
	equal := make(map[rune]struct{})
	for _, a := range str1 {
		equal[a] = struct{}{}
	}
	for _, b := range str2 {
		if _, ok := equal[b]; !ok {
			return false
		}
	}
	return true
}

type ConversionMap map[string]string

type ValueMap []string

func (vm ValueMap) find3Letter(signalPatterns SignalPatterns) error {
	if vm[1] == "" {
		return fmt.Errorf("find3Letter: expected value at 1 of %v", vm)
	}
	for sp := range signalPatterns {
		if len(sp) == 5 && stringsShareParts(sp, vm[1], false) {
			vm[3] = sp
			delete(signalPatterns, sp)
			return nil
		}
	}
	return fmt.Errorf("find3Letter: could not find 3 letter")
}

func (vm ValueMap) find9Letter(signalPatterns SignalPatterns) error {
	if vm[3] == "" {
		return fmt.Errorf("find9Letter: expected value at 3 of %v", vm)
	}
	for sp := range signalPatterns {
		if len(sp) == 6 && stringsShareParts(sp, vm[3], false) {
			vm[9] = sp
			delete(signalPatterns, sp)
			return nil
		}
	}
	return fmt.Errorf("find9Letter: could not find 9 letter")
}

func (vm ValueMap) find5Letter(signalPatterns SignalPatterns) error {
	if vm[9] == "" {
		return fmt.Errorf("find5Letter: expected value at 9 of %v", vm)
	}
	for sp := range signalPatterns {
		if len(sp) == 5 && stringsShareParts(vm[9], sp, false) {
			vm[5] = sp
			delete(signalPatterns, sp)
			return nil
		}
	}
	return fmt.Errorf("find9Letter: could not find 9 letter")
}

func (vm ValueMap) find2Letter(signalPatterns SignalPatterns) error {
	len5sp := []string{}
	for sp := range signalPatterns {
		if len(sp) == 5 {
			len5sp = append(len5sp, sp)
		}
	}
	if len(len5sp) != 1 {
		return fmt.Errorf("find2Letter: expected 1 remaining 5 length sp, got %v", len5sp)
	}
	vm[2] = len5sp[0]
	delete(signalPatterns, len5sp[0])
	return nil
}

func (vm ValueMap) find6Letter(signalPatterns SignalPatterns) error {
	if vm[5] == "" {
		return fmt.Errorf("find6Letter: expected value at 5 of %v", vm)
	}
	for sp := range signalPatterns {
		if len(sp) == 6 && stringsShareParts(sp, vm[5], false) {
			vm[6] = sp
			delete(signalPatterns, sp)
			return nil
		}
	}
	return fmt.Errorf("find6Letter: could not find 6 letter")
}

func (vMap ValueMap) assignValues(signalPatterns SignalPatterns) error {
	// Get simple values
	for i := range signalPatterns {
		if l := len(i); l == 2 {
			vMap[1] = i
			delete(signalPatterns, i)
		} else if l == 3 {
			vMap[7] = i
			delete(signalPatterns, i)
		} else if l == 4 {
			vMap[4] = i
			delete(signalPatterns, i)
		} else if l == 7 {
			vMap[8] = i
			delete(signalPatterns, i)
		}
	}
	err := vMap.find3Letter(signalPatterns)
	if err != nil {
		return err
	}
	err = vMap.find9Letter(signalPatterns)
	if err != nil {
		return err
	}
	err = vMap.find5Letter(signalPatterns)
	if err != nil {
		return err
	}
	err = vMap.find2Letter(signalPatterns)
	if err != nil {
		return err
	}
	err = vMap.find6Letter(signalPatterns)
	if err != nil {
		return err
	}
	for sp := range signalPatterns {
		vMap[0] = sp
	}
	return err
}

func (vMap ValueMap) decodeOutputValue(outputValue []string) int {
	value := 0
	for i, v := range outputValue {
		for j, vv := range vMap {
			if stringsShareParts(v, vv, true) {
				value += (j * int(math.Pow(10, float64(3-i))))
				break
			}
		}
	}
	return value
}

func findSolution(input []string) (int, int, error) {
	re := regexp.MustCompile(`\w+`)
	part1 := 0
	part2 := 0
	for _, i := range input {
		signalPatterns, outputValue, err := parseInput(i, re)
		if err != nil {
			return -1, -1, err
		}
		part1 += countSimpleDigits(outputValue)
		vMap := make(ValueMap, 10)
		if err := vMap.assignValues(signalPatterns); err != nil {
			return -1, -1, err
		}
		part2 += vMap.decodeOutputValue(outputValue)
	}
	return part1, part2, nil
}

func main() {
	input := helpers.ReadFile()
	part1, part2, err := findSolution(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
