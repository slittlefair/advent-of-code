package main

import (
	"testing"
)

func Test_readInputForPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name: "returns 2 for a string of just quotes, (\"\"), advent of code example 1",
			input: []string{
				"\"\"",
			},
			want: 2,
		},
		{
			name: "returns 5 for a string of 3 characters in quotes, (\"abc\"), advent of code example 2",
			input: []string{
				"\"abc\"",
			},
			want: 2,
		},
		{
			name: "returns 7 for a string of characters and esacaped quotation mark, (\"aaa\\\"aaa\", advent of code example 3",
			input: []string{
				"\"aaa\\\"aaa\"",
			},
			want: 3,
		},
		{
			name: "returns 5 for a string of hexadecimal, (\"\\x27\"), advent of code example 4",
			input: []string{
				"\"\\x27\"",
			},
			want: 5,
		},
		{
			name: "returns 3 for a string characters and escaped backslash, (\"aa\\\\aa\"), advent of code example 5",
			input: []string{
				"\"aa\\\\aa\"",
			},
			want: 3,
		},
		{
			name: "returns sum of differences of a collection of strings, advent of code example 6",
			input: []string{
				"\"\"",
				"\"abc\"",
				"\"aaa\\\"aaa\"",
				"\"\\x27\"",
				"\"aa\\\\aa\"",
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readInputForPart1(tt.input); got != tt.want {
				t.Errorf("readInputForPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readInputForPart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name: "returns 4 for a string of just quotes, (\"\"), advent of code example 1",
			input: []string{
				"\"\"",
			},
			want: 4,
		},
		{
			name: "returns 4 for a string of 3 characters in quotes, (\"abc\"), advent of code example 2",
			input: []string{
				"\"abc\"",
			},
			want: 4,
		},
		{
			name: "returns 6 for a string of characters and esacaped quotation mark, (\"aaa\\\"aaa\", advent of code example 3",
			input: []string{
				"\"aaa\\\"aaa\"",
			},
			want: 6,
		},
		{
			name: "returns 5 for a string of hexadecimal, (\"\\x27\"), advent of code example 4",
			input: []string{
				"\"\\x27\"",
			},
			want: 5,
		},
		{
			name: "returns 7 for a string characters and escaped backslash, (\"aa\\\\aa\"), advent of code example 5",
			input: []string{
				"\"aa\\\\aa\"",
			},
			want: 6,
		},
		{
			name: "returns sum of differences of a collection of strings, advent of code example 6",
			input: []string{
				"\"\"",
				"\"abc\"",
				"\"aaa\\\"aaa\"",
				"\"\\x27\"",
				"\"aa\\\\aa\"",
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readInputForPart2(tt.input); got != tt.want {
				t.Errorf("readInputForPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
