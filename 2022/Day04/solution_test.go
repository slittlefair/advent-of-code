package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindContainedAssignments(t *testing.T) {
	t.Run("returns an error if a line of input is malformed", func(t *testing.T) {
		input := []string{
			"2-4,5-9",
			"12-12,9-77",
			"1-32-4",
			"8-9,10-11",
		}
		got, got1, err := findContainedAssignments(input)
		assert.Error(t, err)
		assert.Equal(t, -1, got)
		assert.Equal(t, -1, got1)
	})

	t.Run("returns answers to parts 1 and 2, advent of code example", func(t *testing.T) {
		input := []string{
			"2-4,6-8",
			"2-3,4-5",
			"5-7,7-9",
			"2-8,3-7",
			"6-6,4-6",
			"2-6,4-8",
		}
		got, got1, err := findContainedAssignments(input)
		assert.NoError(t, err)
		assert.Equal(t, 2, got)
		assert.Equal(t, 4, got1)
	})
}
