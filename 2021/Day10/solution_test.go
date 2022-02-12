package main

import (
	"reflect"
	"testing"
)

func Test_invalidDelimeter(t *testing.T) {
	tests := []struct {
		name  string
		line  string
		want  string
		want1 []string
	}{
		{
			name:  "returns an invalid delimeter if one exists, advent of code example 1",
			line:  "{([(<{}[<>[]}>{[]{[(<()>",
			want:  "}",
			want1: nil,
		},
		{
			name:  "returns an invalid delimeter if one exists, advent of code example 2",
			line:  "[[<[([]))<([[{}[[()]]]",
			want:  ")",
			want1: nil,
		},
		{
			name:  "returns an invalid delimeter if one exists, advent of code example 3",
			line:  "[{[{({}]{}}([{[{{{}}([]",
			want:  "]",
			want1: nil,
		},
		{
			name:  "returns an invalid delimeter if one exists, advent of code example 4",
			line:  "[<(<(<(<{}))><([]([]()",
			want:  ")",
			want1: nil,
		},
		{
			name:  "returns an invalid delimeter if one exists, advent of code example 5",
			line:  "<{([([[(<>()){}]>(<<{{",
			want:  ">",
			want1: nil,
		},
		{
			name:  "returns the delimeters of an incomplete line, advent of code example 1",
			line:  "[({(<(())[]>[[{[]{<()<>>",
			want:  "",
			want1: []string{"[", "(", "{", "(", "[", "[", "{", "{"},
		},
		{
			name:  "returns the delimeters of an incomplete line, advent of code example 2",
			line:  "[(()[<>])]({[<{<<[]>>(",
			want:  "",
			want1: []string{"(", "{", "[", "<", "{", "("},
		},
		{
			name:  "returns the delimeters of an incomplete line, advent of code example 3",
			line:  "(((({<>}<{<{<>}{[]{[]{}",
			want:  "",
			want1: []string{"(", "(", "(", "(", "<", "{", "<", "{", "{"},
		},
		{
			name:  "returns the delimeters of an incomplete line, advent of code example 4",
			line:  "{<[[]]>}<{[{[{[]{()[[[]",
			want:  "",
			want1: []string{"<", "{", "[", "{", "[", "{", "{", "[", "["},
		},
		{
			name:  "returns the delimeters of an incomplete line, advent of code example 5",
			line:  "<{([{{}}[<[[[<>{}]]]>[]]",
			want:  "",
			want1: []string{"<", "{", "(", "["},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := invalidDelimeter(tt.line)
			if got != tt.want {
				t.Errorf("invalidDelimeter() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("invalidDelimeter() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_lineCost(t *testing.T) {
	tests := []struct {
		name   string
		chunks []string
		want   int
	}{
		{
			name:   "calculates correct auto complete cost of a line, advent of code example 1",
			chunks: []string{"[", "(", "{", "(", "[", "[", "{", "{"},
			want:   288957,
		},
		{
			name:   "calculates correct auto complete cost of a line, advent of code example 2",
			chunks: []string{"(", "{", "[", "<", "{", "("},
			want:   5566,
		},
		{
			name:   "calculates correct auto complete cost of a line, advent of code example 3",
			chunks: []string{"(", "(", "(", "(", "<", "{", "<", "{", "{"},
			want:   1480781,
		},
		{
			name:   "calculates correct auto complete cost of a line, advent of code example 4",
			chunks: []string{"<", "{", "[", "{", "[", "{", "{", "[", "["},
			want:   995444,
		},
		{
			name:   "calculates correct auto complete cost of a line, advent of code example 5",
			chunks: []string{"<", "{", "(", "["},
			want:   294,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lineCost(tt.chunks); got != tt.want {
				t.Errorf("lineCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSolutions(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
		want1 int
	}{
		{
			name: "returns correct solutions, advent of code example",
			input: []string{
				"[({(<(())[]>[[{[]{<()<>>",
				"[(()[<>])]({[<{<<[]>>(",
				"{([(<{}[<>[]}>{[]{[(<()>",
				"(((({<>}<{<{<>}{[]{[]{}",
				"[[<[([]))<([[{}[[()]]]",
				"[{[{({}]{}}([{[{{{}}([]",
				"{<[[]]>}<{[{[{[]{()[[[]",
				"[<(<(<(<{}))><([]([]()",
				"<{([([[(<>()){}]>(<<{{",
				"<{([{{}}[<[[[<>{}]]]>[]]",
			},
			want:  26397,
			want1: 288957,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findSolutions(tt.input)
			if got != tt.want {
				t.Errorf("findSolutions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findSolutions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
