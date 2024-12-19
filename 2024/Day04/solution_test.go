package main

import (
	"Advent-of-Code/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

var g = map[graph.Co]string{
	{X: 0, Y: 0}: "M",
	{X: 1, Y: 0}: "M",
	{X: 2, Y: 0}: "M",
	{X: 3, Y: 0}: "S",
	{X: 4, Y: 0}: "X",
	{X: 5, Y: 0}: "X",
	{X: 6, Y: 0}: "M",
	{X: 7, Y: 0}: "A",
	{X: 8, Y: 0}: "S",
	{X: 9, Y: 0}: "M",
	{X: 0, Y: 1}: "M",
	{X: 1, Y: 1}: "S",
	{X: 2, Y: 1}: "A",
	{X: 3, Y: 1}: "M",
	{X: 4, Y: 1}: "X",
	{X: 5, Y: 1}: "M",
	{X: 6, Y: 1}: "S",
	{X: 7, Y: 1}: "M",
	{X: 8, Y: 1}: "S",
	{X: 9, Y: 1}: "A",
	{X: 0, Y: 2}: "A",
	{X: 1, Y: 2}: "M",
	{X: 2, Y: 2}: "X",
	{X: 3, Y: 2}: "S",
	{X: 4, Y: 2}: "X",
	{X: 5, Y: 2}: "M",
	{X: 6, Y: 2}: "A",
	{X: 7, Y: 2}: "A",
	{X: 8, Y: 2}: "M",
	{X: 9, Y: 2}: "M",
	{X: 0, Y: 3}: "M",
	{X: 1, Y: 3}: "S",
	{X: 2, Y: 3}: "A",
	{X: 3, Y: 3}: "M",
	{X: 4, Y: 3}: "A",
	{X: 5, Y: 3}: "S",
	{X: 6, Y: 3}: "M",
	{X: 7, Y: 3}: "S",
	{X: 8, Y: 3}: "M",
	{X: 9, Y: 3}: "X",
	{X: 0, Y: 4}: "X",
	{X: 1, Y: 4}: "M",
	{X: 2, Y: 4}: "A",
	{X: 3, Y: 4}: "S",
	{X: 4, Y: 4}: "A",
	{X: 5, Y: 4}: "M",
	{X: 6, Y: 4}: "X",
	{X: 7, Y: 4}: "A",
	{X: 8, Y: 4}: "M",
	{X: 9, Y: 4}: "M",
	{X: 0, Y: 5}: "X",
	{X: 1, Y: 5}: "X",
	{X: 2, Y: 5}: "A",
	{X: 3, Y: 5}: "M",
	{X: 4, Y: 5}: "M",
	{X: 5, Y: 5}: "X",
	{X: 6, Y: 5}: "X",
	{X: 7, Y: 5}: "A",
	{X: 8, Y: 5}: "M",
	{X: 9, Y: 5}: "A",
	{X: 0, Y: 6}: "S",
	{X: 1, Y: 6}: "M",
	{X: 2, Y: 6}: "S",
	{X: 3, Y: 6}: "M",
	{X: 4, Y: 6}: "S",
	{X: 5, Y: 6}: "A",
	{X: 6, Y: 6}: "S",
	{X: 7, Y: 6}: "X",
	{X: 8, Y: 6}: "S",
	{X: 9, Y: 6}: "S",
	{X: 0, Y: 7}: "S",
	{X: 1, Y: 7}: "A",
	{X: 2, Y: 7}: "M",
	{X: 3, Y: 7}: "A",
	{X: 4, Y: 7}: "M",
	{X: 5, Y: 7}: "A",
	{X: 6, Y: 7}: "S",
	{X: 7, Y: 7}: "A",
	{X: 8, Y: 7}: "A",
	{X: 9, Y: 7}: "A",
	{X: 0, Y: 8}: "M",
	{X: 1, Y: 8}: "A",
	{X: 2, Y: 8}: "M",
	{X: 3, Y: 8}: "M",
	{X: 4, Y: 8}: "M",
	{X: 5, Y: 8}: "X",
	{X: 6, Y: 8}: "M",
	{X: 7, Y: 8}: "M",
	{X: 8, Y: 8}: "M",
	{X: 9, Y: 8}: "M",
	{X: 0, Y: 9}: "M",
	{X: 1, Y: 9}: "X",
	{X: 2, Y: 9}: "S",
	{X: 3, Y: 9}: "X",
	{X: 4, Y: 9}: "A",
	{X: 5, Y: 9}: "X",
	{X: 6, Y: 9}: "M",
	{X: 7, Y: 9}: "A",
	{X: 8, Y: 9}: "S",
	{X: 9, Y: 9}: "X",
}

var grid = &graph.Grid{
	MaxX:  9,
	MaxY:  9,
	Graph: g,
}

func Test_parseInput(t *testing.T) {
	t.Run("returns a parsed grid for a given input", func(t *testing.T) {
		input := []string{
			"XMAS",
			"XMSA",
			"SXAM",
			"MAXS",
		}
		g := parseInput(input)
		assert.Equal(t, &graph.Grid{
			MaxX: 4,
			MaxY: 4,
			Graph: map[graph.Co]string{
				{X: 0, Y: 0}: "X",
				{X: 1, Y: 0}: "M",
				{X: 2, Y: 0}: "A",
				{X: 3, Y: 0}: "S",
				{X: 0, Y: 1}: "X",
				{X: 1, Y: 1}: "M",
				{X: 2, Y: 1}: "S",
				{X: 3, Y: 1}: "A",
				{X: 0, Y: 2}: "S",
				{X: 1, Y: 2}: "X",
				{X: 2, Y: 2}: "A",
				{X: 3, Y: 2}: "M",
				{X: 0, Y: 3}: "M",
				{X: 1, Y: 3}: "A",
				{X: 2, Y: 3}: "X",
				{X: 3, Y: 3}: "S",
			},
		}, g)
	})

}

func Test_traverseCos(t *testing.T) {
	t.Run("returns false when we run outside of the grid", func(t *testing.T) {
		assert.False(t, traverseCos(grid, graph.Co{X: 0, Y: 0}, graph.Co{X: -1, Y: -1}, 1))
	})

	t.Run("returns false if the value at the current coordinate isn't what we expect", func(t *testing.T) {
		assert.False(t, traverseCos(grid, graph.Co{X: 4, Y: 1}, graph.Co{X: 0, Y: 1}, 1))
	})

	t.Run(
		"returns false when calling for the first time but will fail on subsequent recursive calls",
		func(t *testing.T) {
			assert.False(t, traverseCos(grid, graph.Co{X: 4, Y: 2}, graph.Co{X: 1, Y: 0}, 0))
		},
	)

	t.Run("returns true when calling for the first time and finding a match", func(t *testing.T) {
		assert.True(t, traverseCos(grid, graph.Co{X: 6, Y: 5}, graph.Co{X: -1, Y: -1}, 0))
	})
}

func Test_findXmas(t *testing.T) {
	t.Run("returns 0 when the given coordinate isn't `X`", func(t *testing.T) {
		assert.Equal(t, 0, findXmas(grid, graph.Co{X: 0, Y: 0}))
	})

	t.Run("returns 0 when the given coordinate doesn't yeld any matches in any direction", func(t *testing.T) {
		assert.Equal(t, 0, findXmas(grid, graph.Co{X: 2, Y: 2}))
	})

	t.Run("returns 1 when the given coordinate yields a match in one direction", func(t *testing.T) {
		assert.Equal(t, 1, findXmas(grid, graph.Co{X: 4, Y: 0}))
	})

	t.Run(
		"returns the number of matches found froma given coordinate when the number is greater than 1",
		func(t *testing.T) {
			assert.Equal(t, 3, findXmas(grid, graph.Co{X: 5, Y: 9}))
		},
	)
}

func Test_findMasInX(t *testing.T) {
	t.Run("returns false when the given coordinate is not `A`", func(t *testing.T) {
		assert.False(t, findMasInX(grid, graph.Co{X: 5, Y: 9}))
	})

	t.Run("returns false when the given coordinate doesn't have correct values in its diagonals", func(t *testing.T) {
		assert.False(t, findMasInX(grid, graph.Co{X: 2, Y: 4}))
	})

	t.Run(
		"returns false when the given coordinate doesn't have the correct diagonals in the correct positions",
		func(t *testing.T) {
			assert.False(t, findMasInX(grid, graph.Co{X: 1, Y: 8}))
		},
	)

	t.Run("returns true when the given coordinate yields a valid X-MAS", func(t *testing.T) {
		assert.True(t, findMasInX(grid, graph.Co{X: 2, Y: 1}))
	})
}

func Test_calculateXmasTotals(t *testing.T) {
	t.Run("returns the correct totals for parts 1 and 2 for a given grid", func(t *testing.T) {
		part1, part2 := calculateXmasTotals(grid)
		assert.Equal(t, 18, part1)
		assert.Equal(t, 9, part2)
	})
}

func Test_findSolutions(t *testing.T) {
	t.Run("returns the correct totals for parts 1 and 2 for a given input", func(t *testing.T) {
		input := []string{
			"MMMSXXMASM",
			"MSAMXMSMSA",
			"AMXSXMAAMM",
			"MSAMASMSMX",
			"XMASAMXAMM",
			"XXAMMXXAMA",
			"SMSMSASXSS",
			"SAMAMASAAA",
			"MAMMMXMMMM",
			"MXSXAXMASX",
		}
		part1, part2 := findSolutions(input)
		assert.Equal(t, 18, part1)
		assert.Equal(t, 9, part2)
	})
}
