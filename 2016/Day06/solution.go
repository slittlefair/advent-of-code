package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
)

func compileFrequencies(input []string) []map[string]int {
	f := make([]map[string]int, len(input[0]))
	for _, line := range input {
		for i, l := range line {
			if f[i] == nil {
				f[i] = map[string]int{}
			}
			f[i][string(l)]++
		}
	}
	return f
}

func getWordMostCommon(f []map[string]int) string {
	w := ""
	for _, m := range f {
		mode := ""
		modeVal := 0
		for k, v := range m {
			if v > modeVal {
				modeVal = v
				mode = k
			}
		}
		w += mode
	}
	return w
}

func getWordLeastCommon(f []map[string]int) string {
	w := ""
	for _, m := range f {
		mode := ""
		modeVal := utils.Infinity
		for k, v := range m {
			if v < modeVal {
				modeVal = v
				mode = k
			}
		}
		w += mode
	}
	return w
}

func main() {
	input := utils.ReadFile()
	f := compileFrequencies(input)
	fmt.Println(f)
	fmt.Println("Part 1:", getWordMostCommon(f))
	fmt.Println("Part 2:", getWordLeastCommon(f))
}
