package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createReportsFromInput(t *testing.T) {
	t.Run("returns a report from a given input", func(t *testing.T) {
		expectedReports := [][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 3, 6, 7, 9},
		}
		input := []string{
			"7 6 4 2 1",
			"1 2 7 8 9",
			"9 7 6 2 1",
			"1 3 2 4 5",
			"8 6 4 4 1",
			"1 3 6 7 9",
		}

		assert.Equal(t, expectedReports, createReportsFromInput(input))
	})
}

func Test_reportIsSafe(t *testing.T) {
	t.Run(
		"returns false and the index of the offending level if there is an increase in levels greater than 3",
		func(t *testing.T) {
			safe, index := reportIsSafe([]int{1, 2, 7, 8, 9})
			assert.False(t, safe)
			assert.Equal(t, 1, index)
		},
	)

	t.Run(
		"returns false and the index of the offending level if there is a decrease in levels greater than 3",
		func(t *testing.T) {
			safe, index := reportIsSafe([]int{9, 7, 6, 2, 1})
			assert.False(t, safe)
			assert.Equal(t, 2, index)
		},
	)

	t.Run(
		"returns false and the index of the offending level if sequential levels are the same",
		func(t *testing.T) {
			safe, index := reportIsSafe([]int{8, 6, 4, 4, 1})
			assert.False(t, safe)
			assert.Equal(t, 2, index)
		},
	)

	t.Run(
		"returns false and the index of the offending level if levels change direction",
		func(t *testing.T) {
			safe, index := reportIsSafe([]int{1, 3, 2, 4, 5})
			assert.False(t, safe)
			assert.Equal(t, 1, index)
		},
	)

	t.Run(
		"returns true and -1 if a report is safe",
		func(t *testing.T) {
			safe, index := reportIsSafe([]int{7, 6, 4, 2, 1})
			assert.True(t, safe)
			assert.Equal(t, -1, index)
		},
	)
}

func Test_findSolutions(t *testing.T) {
	t.Run("finds correct solutions for a given input", func(t *testing.T) {
		part1, part2 := findSolutions(
			[]string{
				"7 6 4 2 1",
				"1 2 7 8 9",
				"9 7 6 2 1",
				"1 3 2 4 5",
				"8 6 4 4 1",
				"1 3 6 7 9",
				"9 1 3 2 4 5",
			},
		)
		assert.Equal(t, 2, part1)
		assert.Equal(t, 4, part2)
	})
}
