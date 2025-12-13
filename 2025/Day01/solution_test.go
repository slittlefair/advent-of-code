package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findSolutions(t *testing.T) {
	t.Run("returns correct number when going to 0", func(t *testing.T) {
		input := []string{
			"L50",
			"L10",
			"R10",
			"R15",
			"R250",
			"R500",
			"R235",
			"L100",
			"L200",
		}
		part1, part2, err := findSolutions(input)
		assert.NoError(t, err)
		assert.Equal(t, 5, part1)
		assert.Equal(t, 15, part2)
	})
}
