package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findSolutions(t *testing.T) {
	t.Run("returns correct solutions to parts 1 and 2, 1", func(t *testing.T) {
		input := []string{
			"AAAA",
			"BBCD",
			"BBCC",
			"EEEC",
		}
		part1, part2 := findSolutions(input)
		assert.Equal(t, 140, part1)
		assert.Equal(t, 80, part2)
	})

	t.Run("returns correct solutions to parts 1 and 2, 2", func(t *testing.T) {
		input := []string{
			"OOOOO",
			"OXOXO",
			"OOOOO",
			"OXOXO",
			"OOOOO",
		}
		part1, part2 := findSolutions(input)
		assert.Equal(t, 772, part1)
		assert.Equal(t, 436, part2)
	})

	t.Run("returns correct solutions to part 1 and part2, 3", func(t *testing.T) {
		input := []string{
			"RRRRIICCFF",
			"RRRRIICCCF",
			"VVRRRCCFFF",
			"VVRCCCJFFF",
			"VVVVCJJCFE",
			"VVIVCCJJEE",
			"VVIIICJJEE",
			"MIIIIIJJEE",
			"MIIISIJEEE",
			"MMMISSJEEE",
		}
		part1, part2 := findSolutions(input)
		assert.Equal(t, 1930, part1)
		assert.Equal(t, 1206, part2)
	})

	t.Run("returns correct solutions to part 2", func(t *testing.T) {
		input := []string{
			"EEEEE",
			"EXXXX",
			"EEEEE",
			"EXXXX",
			"EEEEE",
		}
		_, part2 := findSolutions(input)
		assert.Equal(t, 236, part2)
	})

	t.Run("returns correct solutions to part 2", func(t *testing.T) {
		input := []string{
			"AAAAAA",
			"AAABBA",
			"AAABBA",
			"ABBAAA",
			"ABBAAA",
			"AAAAAA",
		}
		_, part2 := findSolutions(input)
		assert.Equal(t, 368, part2)
	})
}
