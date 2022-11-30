package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculateIncreases(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name: "returns the correct number of increases in the input, advent of code example 1",
			input: []int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			want: 7,
		},
		{
			name: "returns the correct number of increases in the input, advent of code example 2",
			input: []int{
				607,
				618,
				618,
				617,
				647,
				716,
				769,
				792,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateIncreases(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_calculateSlidingWindows(t *testing.T) {
	t.Run("returns a slice of 3 value inputs from given slice", func(t *testing.T) {
		input := []int{
			199,
			200,
			208,
			210,
			200,
			207,
			240,
			269,
			260,
			263,
		}
		want := []int{
			607,
			618,
			618,
			617,
			647,
			716,
			769,
			792,
		}
		got := calculateSlidingWindows(input)
		assert.Equal(t, want, got)
	})
}

func Test_part1(t *testing.T) {
	t.Run("returns number of increases in a given input, advent of code example", func(t *testing.T) {
		got := part1(
			[]int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
		)
		assert.Equal(t, 7, got)
	})
}

func Test_part2(t *testing.T) {
	t.Run("returns number of increases in sliding windows of a given input", func(t *testing.T) {
		got := part2(
			[]int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
		)
		assert.Equal(t, 5, got)
	})
}
