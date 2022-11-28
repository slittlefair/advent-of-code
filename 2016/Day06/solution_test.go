package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var f = []map[string]int{
	{"a": 1, "d": 2, "e": 3, "n": 2, "r": 2, "s": 2, "t": 2, "v": 2},
	{"a": 3, "d": 1, "e": 2, "n": 2, "r": 2, "s": 2, "t": 2, "v": 2},
	{"a": 2, "d": 2, "e": 2, "n": 2, "r": 2, "s": 3, "t": 2, "v": 1},
	{"a": 2, "d": 2, "e": 1, "n": 2, "r": 2, "s": 2, "t": 3, "v": 2},
	{"a": 2, "d": 2, "e": 3, "n": 1, "r": 2, "s": 2, "t": 2, "v": 2},
	{"a": 2, "d": 2, "e": 2, "n": 2, "r": 3, "s": 2, "t": 1, "v": 2},
}

func Test_compileFrequencies(t *testing.T) {
	t.Run("it compiles the frequencies of each letter in each column", func(t *testing.T) {
		input := []string{
			"eedadn",
			"drvtee",
			"eandsr",
			"raavrd",
			"atevrs",
			"tsrnev",
			"sdttsa",
			"rasrtv",
			"nssdts",
			"ntnada",
			"svetve",
			"tesnvt",
			"vntsnd",
			"vrdear",
			"dvrsen",
			"enarar",
		}
		got := compileFrequencies(input)
		assert.Equal(t, f, got)
	})
}

func Test_getWordMostCommon(t *testing.T) {
	t.Run("finds the word compiled with most frequent letter in each column, advent of code example 1", func(t *testing.T) {
		got := getWordMostCommon(f)
		assert.Equal(t, "easter", got)
	})
}

func Test_getWordLeastCommon(t *testing.T) {
	t.Run("finds the word compiled with least frequent letter in each column, advent of code example 2", func(t *testing.T) {
		got := getWordLeastCommon(f)
		assert.Equal(t, "advent", got)
	})
}
