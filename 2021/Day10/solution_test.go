package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
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
			got := lineCost(tt.chunks)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_findSolutions(t *testing.T) {
	t.Run("returns correct solutions, advent of code example", func(t *testing.T) {
		input := []string{
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
		}
		got, got1 := findSolutions(input)
		assert.Equal(t, 26397, got)
		assert.Equal(t, 288957, got1)
	})
}
