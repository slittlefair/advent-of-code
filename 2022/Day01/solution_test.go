package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	t.Run("returns an error if an item in the input cannot be converted into an int", func(t *testing.T) {
		input := []string{
			"1000",
			"2000",
			"two",
			"3000",
		}
		got, err := parseInput(input)
		assert.Error(t, err)
		assert.Nil(t, got)
	})

	t.Run("parses input into a slice of elf/calorie amounts, advent of code example", func(t *testing.T) {
		input := []string{
			"1000",
			"2000",
			"3000",
			"",
			"4000",
			"",
			"5000",
			"6000",
			"",
			"7000",
			"8000",
			"9000",
			"",
			"10000",
		}
		got, err := parseInput(input)
		assert.NoError(t, err)
		assert.Equal(t, []int{6000, 4000, 11000, 24000, 10000}, got)
	})
}

func TestFindLargestCalories(t *testing.T) {
	tests := []struct {
		name               string
		elves              []int
		n                  int
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "returns 0 if provided elves is empty",
			elves:              []int{},
			n:                  3,
			want:               -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "returns the largest calorie value if only one asked for, advent of code example 1",
			elves:              []int{6000, 4000, 11000, 24000, 10000},
			n:                  1,
			want:               24000,
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "returns the sum of the three largest calorie values, advent of code example 2",
			elves:              []int{6000, 4000, 11000, 24000, 10000},
			n:                  3,
			want:               45000,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findLargestCalories(tt.elves, tt.n)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
