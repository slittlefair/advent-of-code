package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	t.Run("returns an error if a character cannot be converted to an int", func(t *testing.T) {
		input := "23498a765"
		output, err := parseInput(input)
		assert.Nil(t, output)
		assert.Error(t, err)
	})

	t.Run("returns a list of converted ints from the input", func(t *testing.T) {
		input := "234989765"
		output, err := parseInput(input)
		expectedOutput := []int{2, 3, 4, 9, 8, 9, 7, 6, 5}
		assert.Equal(t, expectedOutput, output)
		assert.NoError(t, err)
	})
}

func TestFindSolutions(t *testing.T) {
	tests := []struct {
		part          int
		exampleNumber int
		input         []int
		expectedPart1 int
		expectedPart2 int
	}{
		{
			part:          1,
			exampleNumber: 1,
			input:         []int{1, 1, 2, 2},
			expectedPart1: 3,
		},
		{
			part:          1,
			exampleNumber: 2,
			input:         []int{1, 1, 1, 1},
			expectedPart1: 4,
		},
		{
			part:          1,
			exampleNumber: 3,
			input:         []int{1, 2, 3, 4},
			expectedPart1: 0,
		},
		{
			part:          1,
			exampleNumber: 4,
			input:         []int{9, 1, 2, 1, 2, 1, 2, 9},
			expectedPart1: 9,
		},
		{
			part:          2,
			exampleNumber: 1,
			input:         []int{1, 2, 1, 2},
			expectedPart2: 6,
		},
		{
			part:          2,
			exampleNumber: 2,
			input:         []int{1, 2, 2, 1},
			expectedPart2: 0,
		},
		{
			part:          2,
			exampleNumber: 3,
			input:         []int{1, 2, 3, 4, 2, 5},
			expectedPart2: 4,
		},
		{
			part:          2,
			exampleNumber: 4,
			input:         []int{1, 2, 3, 1, 2, 3},
			expectedPart2: 12,
		},
		{
			part:          2,
			exampleNumber: 5,
			input:         []int{1, 2, 1, 3, 1, 4, 1, 5},
			expectedPart2: 4,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("returns correct part%d solution, advent of code example %d", tt.part, tt.exampleNumber), func(t *testing.T) {
			part1, part2 := findSolution(tt.input)
			if tt.part == 1 {
				assert.Equal(t, tt.expectedPart1, part1)
			} else {
				assert.Equal(t, tt.expectedPart2, part2)
			}
		})
	}
}
