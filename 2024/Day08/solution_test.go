package main

import (
	"Advent-of-Code/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInput(t *testing.T) {
	t.Run("produces an AntennaMap for a given input", func(t *testing.T) {
		input := []string{
			"............",
			"........0...",
			".....0......",
			".......0....",
			"....0.......",
			"......A.....",
			"............",
			"............",
			"........A...",
			".........A..",
			"............",
			"............",
		}

		want := AntennaMap{
			Grid: graph.Grid{
				MaxX: 11,
				MaxY: 11,
				Graph: graph.Graph{
					{X: 8, Y: 1}: "0",
					{X: 5, Y: 2}: "0",
					{X: 7, Y: 3}: "0",
					{X: 4, Y: 4}: "0",
					{X: 6, Y: 5}: "A",
					{X: 8, Y: 8}: "A",
					{X: 9, Y: 9}: "A",
				},
			},
			antennas: map[string]map[graph.Co]bool{
				"0": {
					{X: 8, Y: 1}: true,
					{X: 5, Y: 2}: true,
					{X: 7, Y: 3}: true,
					{X: 4, Y: 4}: true,
				},
				"A": {
					{X: 6, Y: 5}: true,
					{X: 8, Y: 8}: true,
					{X: 9, Y: 9}: true,
				},
			},
			antinodes:      make(map[graph.Co]bool),
			antinodesPart2: make(map[graph.Co]bool),
		}

		got := parseInput(input)
		assert.Equal(t, want, got)
	})
}

func TestAntennaMap_findAntinodes(t *testing.T) {
	t.Run("returns number of antinodes for part 1 for a given input", func(t *testing.T) {
		am := AntennaMap{
			Grid: graph.Grid{
				MaxX: 9,
				MaxY: 9,
				Graph: graph.Graph{
					{X: 4, Y: 3}: "a",
					{X: 8, Y: 4}: "a",
					{X: 5, Y: 5}: "a",
				},
			},
			antennas: map[string]map[graph.Co]bool{
				"a": {
					{X: 4, Y: 3}: true,
					{X: 8, Y: 4}: true,
					{X: 5, Y: 5}: true,
				},
			},
			antinodes:      make(map[graph.Co]bool),
			antinodesPart2: make(map[graph.Co]bool),
		}

		part1, _ := am.findAntinodes()
		assert.Equal(t, 4, part1)
	})

	t.Run("returns number of antinodes for part 2 for a given input", func(t *testing.T) {
		am := AntennaMap{
			Grid: graph.Grid{
				MaxX: 9,
				MaxY: 9,
				Graph: graph.Graph{
					{X: 0, Y: 0}: "T",
					{X: 3, Y: 1}: "T",
					{X: 1, Y: 2}: "T",
				},
			},
			antennas: map[string]map[graph.Co]bool{
				"T": {
					{X: 0, Y: 0}: true,
					{X: 3, Y: 1}: true,
					{X: 1, Y: 2}: true,
				},
			},
			antinodes:      make(map[graph.Co]bool),
			antinodesPart2: make(map[graph.Co]bool),
		}

		_, part2 := am.findAntinodes()
		assert.Equal(t, 9, part2)
	})

	t.Run("returns number of antinodes for parts 1 and 2 for a given input", func(t *testing.T) {
		am := AntennaMap{
			Grid: graph.Grid{
				MaxX: 11,
				MaxY: 11,
				Graph: graph.Graph{
					{X: 8, Y: 1}: "0",
					{X: 5, Y: 2}: "0",
					{X: 7, Y: 3}: "0",
					{X: 4, Y: 4}: "0",
					{X: 6, Y: 5}: "A",
					{X: 8, Y: 8}: "A",
					{X: 9, Y: 9}: "A",
				},
			},
			antennas: map[string]map[graph.Co]bool{
				"0": {
					{X: 8, Y: 1}: true,
					{X: 5, Y: 2}: true,
					{X: 7, Y: 3}: true,
					{X: 4, Y: 4}: true,
				},
				"A": {
					{X: 6, Y: 5}: true,
					{X: 8, Y: 8}: true,
					{X: 9, Y: 9}: true,
				},
			},
			antinodes:      make(map[graph.Co]bool),
			antinodesPart2: make(map[graph.Co]bool),
		}

		part1, part2 := am.findAntinodes()
		assert.Equal(t, 14, part1)
		assert.Equal(t, 34, part2)
	})
}

func Test_findSolutions(t *testing.T) {
	t.Run("finds solutions to parts 1 and 2 for a given input", func(t *testing.T) {
		input := []string{
			"............",
			"........0...",
			".....0......",
			".......0....",
			"....0.......",
			"......A.....",
			"............",
			"............",
			"........A...",
			".........A..",
			"............",
			"............",
		}

		part1, part2 := findSolutions(input)
		assert.Equal(t, 14, part1)
		assert.Equal(t, 34, part2)
	})
}
