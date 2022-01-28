package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/maths"
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
		modeVal := maths.Infinity
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
	input := file.Read()
	f := compileFrequencies(input)
	fmt.Println(f)
	fmt.Println("Part 1:", getWordMostCommon(f))
	fmt.Println("Part 2:", getWordLeastCommon(f))
}
