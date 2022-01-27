package main

import (
	"reflect"
	"testing"
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
	tests := []struct {
		name  string
		input []string
		want  []map[string]int
	}{
		{
			name: "it compiles the frequencies of each letter in each column",
			input: []string{
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
			},
			want: f,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compileFrequencies(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compileFrequencies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getWordMostCommon(t *testing.T) {
	tests := []struct {
		name string
		f    []map[string]int
		want string
	}{
		{
			name: "finds the word compiled with most frequent letter in each column, advent of code example 1",
			f:    f,
			want: "easter",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWordMostCommon(tt.f); got != tt.want {
				t.Errorf("getWordMostCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getWordLeastCommon(t *testing.T) {
	tests := []struct {
		name string
		f    []map[string]int
		want string
	}{
		{
			name: "finds the word compiled with least frequent letter in each column, advent of code example 2",
			f:    f,
			want: "advent",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWordLeastCommon(tt.f); got != tt.want {
				t.Errorf("getWordLeastCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}
