package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	t.Run("returns a Spreadsheet from given input", func(t *testing.T) {
		got := parseInput([]string{
			"5 1 9 5",
			"7 5 3",
			"2 4 6 8",
		})
		assert.Equal(t, Spreadsheet{
			{5, 1, 9, 5},
			{7, 5, 3},
			{2, 4, 6, 8},
		}, got)
	})
}

func TestDividesEvenly(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		y    float64
		want bool
	}{
		{
			name: "returns false if neither given value divides the other evenly",
			x:    5,
			y:    9,
			want: false,
		},
		{
			name: "returns true if the first number divides the second",
			x:    72,
			y:    9,
			want: true,
		},
		{
			name: "returns true if the second number divides the first",
			x:    4,
			y:    100,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dividesEvenly(tt.x, tt.y)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFindSolutions(t *testing.T) {
	t.Run("returns the correct checksum for part1, advent of code example", func(t *testing.T) {
		spreadsheet := Spreadsheet{
			{5, 1, 9, 5},
			{7, 5, 3},
			{2, 4, 6, 8},
		}
		got, _ := findSolutions(spreadsheet)
		assert.Equal(t, 18, got)
	})

	t.Run("returns the correct checksum for part1 and part2, advent of code example", func(t *testing.T) {
		spreadsheet := Spreadsheet{
			{5, 9, 2, 8},
			{9, 4, 7, 3},
			{3, 8, 6, 5},
		}
		got, got1 := findSolutions(spreadsheet)
		assert.Equal(t, 18, got)
		assert.Equal(t, 9, got1)
	})
}
