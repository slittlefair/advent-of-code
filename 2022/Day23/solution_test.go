package main

import (
	"Advent-of-Code/graph"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	t.Run("parses an input", func(t *testing.T) {
		input := []string{
			".....",
			"..##.",
			"..#..",
			".....",
			"..##.",
			".....",
		}
		want := &Grove{
			minX: 2,
			maxX: 3,
			minY: 1,
			maxY: 4,
			elves: map[graph.Co]bool{
				{X: 2, Y: 1}: true,
				{X: 3, Y: 1}: true,
				{X: 2, Y: 2}: true,
				{X: 2, Y: 4}: true,
				{X: 3, Y: 4}: true,
			},
		}
		got := parseInput(input)
		assert.Equal(t, want, got)
	})
}

func TestProposeMove(t *testing.T) {
	tests := []struct {
		grove            Grove
		elf              graph.Co
		currentDirection Direction
		want             graph.Co
		want1            bool
	}{
		{
			grove: Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 1}: true,
					{X: 3, Y: 1}: true,
					{X: 2, Y: 2}: true,
					{X: 2, Y: 4}: true,
					{X: 3, Y: 4}: true,
				},
			},
			elf:              graph.Co{X: 2, Y: 1},
			currentDirection: North,
			want:             graph.Co{X: 2, Y: 0},
			want1:            true,
		},
		{
			grove: Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 1}: true,
					{X: 3, Y: 1}: true,
					{X: 2, Y: 2}: true,
					{X: 2, Y: 4}: true,
					{X: 3, Y: 4}: true,
				},
			},
			elf:              graph.Co{X: 3, Y: 1},
			currentDirection: North,
			want:             graph.Co{X: 3, Y: 0},
			want1:            true,
		},
		{
			grove: Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 1}: true,
					{X: 3, Y: 1}: true,
					{X: 2, Y: 2}: true,
					{X: 2, Y: 4}: true,
					{X: 3, Y: 4}: true,
				},
			},
			elf:              graph.Co{X: 2, Y: 2},
			currentDirection: North,
			want:             graph.Co{X: 2, Y: 3},
			want1:            true,
		},
		{
			grove: Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 1}: true,
					{X: 3, Y: 1}: true,
					{X: 2, Y: 2}: true,
					{X: 2, Y: 4}: true,
					{X: 3, Y: 4}: true,
				},
			},
			elf:              graph.Co{X: 2, Y: 4},
			currentDirection: North,
			want:             graph.Co{X: 2, Y: 3},
			want1:            true,
		},
		{
			grove: Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 1}: true,
					{X: 3, Y: 1}: true,
					{X: 2, Y: 2}: true,
					{X: 2, Y: 4}: true,
					{X: 3, Y: 4}: true,
				},
			},
			elf:              graph.Co{X: 3, Y: 4},
			currentDirection: North,
			want:             graph.Co{X: 3, Y: 3},
			want1:            true,
		},
		{
			grove: Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 0}: true,
					{X: 3, Y: 0}: true,
					{X: 2, Y: 2}: true,
					{X: 2, Y: 4}: true,
					{X: 3, Y: 3}: true,
				},
			},
			elf:              graph.Co{X: 2, Y: 0},
			currentDirection: South,
			want:             graph.Co{X: 2, Y: 1},
			want1:            true,
		},
		{
			grove: Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 0}: true,
					{X: 3, Y: 0}: true,
					{X: 2, Y: 2}: true,
					{X: 2, Y: 4}: true,
					{X: 3, Y: 3}: true,
				},
			},
			elf:              graph.Co{X: 3, Y: 0},
			currentDirection: South,
			want:             graph.Co{X: 3, Y: 1},
			want1:            true,
		},
		{
			grove: Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 0}: true,
					{X: 3, Y: 0}: true,
					{X: 2, Y: 2}: true,
					{X: 2, Y: 4}: true,
					{X: 3, Y: 3}: true,
				},
			},
			elf:              graph.Co{X: 2, Y: 2},
			currentDirection: South,
			want:             graph.Co{X: 1, Y: 2},
			want1:            true,
		},
		{
			grove: Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 0}: true,
					{X: 3, Y: 0}: true,
					{X: 2, Y: 2}: true,
					{X: 2, Y: 4}: true,
					{X: 3, Y: 3}: true,
				},
			},
			elf:              graph.Co{X: 2, Y: 4},
			currentDirection: South,
			want:             graph.Co{X: 2, Y: 5},
			want1:            true,
		},
		{
			grove: Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 0}: true,
					{X: 3, Y: 0}: true,
					{X: 2, Y: 2}: true,
					{X: 2, Y: 4}: true,
					{X: 3, Y: 3}: true,
				},
			},
			elf:              graph.Co{X: 3, Y: 3},
			currentDirection: South,
			want:             graph.Co{X: 4, Y: 3},
			want1:            true,
		},
		{
			grove: Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 1}: true,
					{X: 3, Y: 1}: true,
					{X: 1, Y: 2}: true,
					{X: 2, Y: 5}: true,
					{X: 4, Y: 3}: true,
				},
			},
			elf:              graph.Co{X: 2, Y: 1},
			currentDirection: West,
			want:             graph.Co{X: 2, Y: 0},
			want1:            true,
		},
		{
			grove: Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 1}: true,
					{X: 3, Y: 1}: true,
					{X: 1, Y: 2}: true,
					{X: 2, Y: 5}: true,
					{X: 4, Y: 3}: true,
				},
			},
			elf:              graph.Co{X: 3, Y: 1},
			currentDirection: West,
			want:             graph.Co{X: 4, Y: 1},
			want1:            true,
		},
		{
			grove: Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 1}: true,
					{X: 3, Y: 1}: true,
					{X: 1, Y: 2}: true,
					{X: 2, Y: 5}: true,
					{X: 4, Y: 3}: true,
				},
			},
			elf:              graph.Co{X: 1, Y: 2},
			currentDirection: West,
			want:             graph.Co{X: 0, Y: 2},
			want1:            true,
		},
		{
			grove: Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 1}: true,
					{X: 3, Y: 1}: true,
					{X: 1, Y: 2}: true,
					{X: 2, Y: 5}: true,
					{X: 4, Y: 3}: true,
				},
			},
			elf:              graph.Co{X: 2, Y: 5},
			currentDirection: West,
			want:             graph.Co{X: 2, Y: 5},
			want1:            false,
		},
		{
			grove: Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 1}: true,
					{X: 3, Y: 1}: true,
					{X: 1, Y: 2}: true,
					{X: 2, Y: 5}: true,
					{X: 4, Y: 3}: true,
				},
			},
			elf:              graph.Co{X: 4, Y: 3},
			currentDirection: West,
			want:             graph.Co{X: 4, Y: 3},
			want1:            false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("proposes elf moves to the correct position, advent of code example %d", i+1), func(t *testing.T) {
			got, got1 := tt.grove.proposeMove(tt.elf, tt.currentDirection)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}

	t.Run("returns original elf coordinates if it wants to move but can't", func(t *testing.T) {
		grove := Grove{
			elves: map[graph.Co]bool{
				{X: 0, Y: 1}: true,
				{X: 1, Y: 1}: true,
				{X: 2, Y: 1}: true,
				{X: 1, Y: 0}: true,
				{X: 1, Y: 2}: true,
			},
		}
		elf := graph.Co{X: 1, Y: 1}
		got, got1 := grove.proposeMove(elf, North)
		assert.Equal(t, graph.Co{X: 1, Y: 1}, got)
		assert.Equal(t, false, got1)
	})
}

func TestProposeMoves(t *testing.T) {
	tests := []struct {
		grove            *Grove
		currentDirection Direction
		want             bool
		want1            *Grove
	}{
		{
			grove: &Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 1}: true,
					{X: 3, Y: 1}: true,
					{X: 2, Y: 2}: true,
					{X: 2, Y: 4}: true,
					{X: 3, Y: 4}: true,
				},
				minX: 2,
				maxX: 3,
				minY: 1,
				maxY: 4,
			},
			currentDirection: North,
			want:             false,
			want1: &Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 0}: true,
					{X: 3, Y: 0}: true,
					{X: 2, Y: 2}: true,
					{X: 2, Y: 4}: true,
					{X: 3, Y: 3}: true,
				},
				minX: 2,
				maxX: 3,
				minY: 0,
				maxY: 4,
			},
		},
		{
			grove: &Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 0}: true,
					{X: 3, Y: 0}: true,
					{X: 2, Y: 2}: true,
					{X: 2, Y: 4}: true,
					{X: 3, Y: 3}: true,
				},
				minX: 2,
				maxX: 3,
				minY: 0,
				maxY: 4,
			},
			currentDirection: South,
			want:             false,
			want1: &Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 1}: true,
					{X: 3, Y: 1}: true,
					{X: 1, Y: 2}: true,
					{X: 2, Y: 5}: true,
					{X: 4, Y: 3}: true,
				},
				minX: 1,
				maxX: 4,
				minY: 1,
				maxY: 5,
			},
		},
		{
			grove: &Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 1}: true,
					{X: 3, Y: 1}: true,
					{X: 1, Y: 2}: true,
					{X: 2, Y: 5}: true,
					{X: 4, Y: 3}: true,
				},
				minX: 1,
				maxX: 4,
				minY: 1,
				maxY: 5,
			},
			currentDirection: West,
			want:             false,
			want1: &Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 0}: true,
					{X: 4, Y: 1}: true,
					{X: 0, Y: 2}: true,
					{X: 4, Y: 3}: true,
					{X: 2, Y: 5}: true,
				},
				minX: 0,
				maxX: 4,
				minY: 0,
				maxY: 5,
			},
		},
		{
			grove: &Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 0}: true,
					{X: 4, Y: 1}: true,
					{X: 0, Y: 2}: true,
					{X: 4, Y: 3}: true,
					{X: 2, Y: 5}: true,
				},
				minX: 0,
				maxX: 4,
				minY: 0,
				maxY: 5,
			},
			currentDirection: East,
			want:             true,
			want1: &Grove{
				elves: map[graph.Co]bool{
					{X: 2, Y: 0}: true,
					{X: 4, Y: 1}: true,
					{X: 0, Y: 2}: true,
					{X: 4, Y: 3}: true,
					{X: 2, Y: 5}: true,
				},
				minX: 0,
				maxX: 4,
				minY: 0,
				maxY: 5,
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("updates grid with elf moves, advent of code example %d", i+1), func(t *testing.T) {
			grove := tt.grove
			got := grove.proposeMoves(tt.currentDirection)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, grove)
		})
	}
}

func TestFindNumEmptyTiles(t *testing.T) {
	t.Run("find the number of empty tiles from the given grid", func(t *testing.T) {
		grove := Grove{
			elves: map[graph.Co]bool{
				{X: 2, Y: 0}: true,
				{X: 4, Y: 1}: true,
				{X: 0, Y: 2}: true,
				{X: 4, Y: 3}: true,
				{X: 2, Y: 5}: true,
			},
			minX: 0,
			maxX: 4,
			minY: 0,
			maxY: 5,
		}
		got := grove.findNumEmptyTiles()
		assert.Equal(t, 25, got)
	})
}

func TestFindSolutions(t *testing.T) {
	t.Run("finds solutions for parts 1 and 2, advent of code example", func(t *testing.T) {
		input := []string{
			"..............",
			"..............",
			".......#......",
			".....###.#....",
			"...#...#.#....",
			"....#...##....",
			"...#.###......",
			"...##.#.##....",
			"....#..#......",
			"..............",
			"..............",
			"..............",
		}
		got, got1 := findSolutions(input)
		assert.Equal(t, 110, got)
		assert.Equal(t, 20, got1)
	})
}
