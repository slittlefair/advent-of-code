package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createSortedSlicesFromInput(t *testing.T) {
	t.Run("returns an error if the line has fewer than two numbers in it", func(t *testing.T) {
		input := []string{
			"1   2",
			"3   4",
			"5   six",
		}
		s1, s2, err := createSortedSlicesFromInput(input)
		assert.Error(t, err)
		assert.Nil(t, s1)
		assert.Nil(t, s2)
	})

	t.Run("returns an error if the line has more than two numbers in it", func(t *testing.T) {
		input := []string{
			"1   2",
			"3   4   4.5",
			"5   6",
		}
		s1, s2, err := createSortedSlicesFromInput(input)
		assert.Error(t, err)
		assert.Nil(t, s1)
		assert.Nil(t, s2)
	})

	t.Run("returns two sorted lists of numbers for the expected columns", func(t *testing.T) {
		input := []string{
			"3   4",
			"4   3",
			"2   5",
			"1   3",
			"3   9",
			"3   3",
		}
		s1, s2, err := createSortedSlicesFromInput(input)
		assert.Equal(t, []int{1, 2, 3, 3, 3, 4}, s1)
		assert.Equal(t, []int{3, 3, 3, 4, 5, 9}, s2)
		assert.Nil(t, err)
	})
}

func Test_calculateDiffInSlices(t *testing.T) {
	t.Run("returns an error if first slice isn't sorted", func(t *testing.T) {
		slice1 := []int{1, 2, 3, 5, 4}
		slice2 := []int{7, 8, 9, 10, 11}
		diff, err := calculateDiffInSlices(slice1, slice2)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "slice1")
		assert.Zero(t, diff)
	})

	t.Run("returns an error if second slice isn't sorted", func(t *testing.T) {
		slice1 := []int{1, 2, 3, 4, 5}
		slice2 := []int{7, 88, 9, 10, 11}
		diff, err := calculateDiffInSlices(slice1, slice2)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "slice2")
		assert.Zero(t, diff)
	})

	t.Run("returns an error if the slices aren't equal in length", func(t *testing.T) {
		slice1 := []int{1, 2, 3, 4, 5, 6}
		slice2 := []int{7, 8, 9, 10, 11}
		diff, err := calculateDiffInSlices(slice1, slice2)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "length")
		assert.Zero(t, diff)
	})

	t.Run("returns the correct diff for two equal length, sorted slices", func(t *testing.T) {
		slice1 := []int{1, 2, 3, 3, 3, 4}
		slice2 := []int{3, 3, 3, 4, 5, 9}
		diff, err := calculateDiffInSlices(slice1, slice2)
		assert.Equal(t, diff, 11)
		assert.Zero(t, err)
	})
}

func Test_calculateSimilarityScore(t *testing.T) {
	t.Run("returns the correct similarity score for two slices", func(t *testing.T) {
		slice1 := []int{1, 2, 3, 3, 3, 4}
		slice2 := []int{3, 3, 3, 4, 5, 9}
		score := calculateSimilarityScore(slice1, slice2)
		assert.Equal(t, score, 31)
	})
}

func Test_findSolutions(t *testing.T) {
	t.Run("returns an error if sorted slices can't be created from the input", func(t *testing.T) {
		input := []string{
			"1   2",
			"3   4",
			"5   six",
		}
		part1, part2, err := findSolutions(input)
		assert.Error(t, err)
		assert.Zero(t, part1)
		assert.Zero(t, part2)
	})

	t.Run("returns correct solutions to parts 1 and 2", func(t *testing.T) {
		input := []string{
			"3   4",
			"4   3",
			"2   5",
			"1   3",
			"3   9",
			"3   3",
		}
		part1, part2, err := findSolutions(input)
		assert.Equal(t, 11, part1)
		assert.Equal(t, 31, part2)
		assert.Nil(t, err)
	})
}
