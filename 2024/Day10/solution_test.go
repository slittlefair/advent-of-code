package main

import (
	"Advent-of-Code/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleGrid = graph.Grid[int]{
	MaxX: 7,
	MaxY: 7,
	Graph: graph.Graph[int]{
		{X: 0, Y: 0}: 8,
		{X: 1, Y: 0}: 9,
		{X: 2, Y: 0}: 0,
		{X: 3, Y: 0}: 1,
		{X: 4, Y: 0}: 0,
		{X: 5, Y: 0}: 1,
		{X: 6, Y: 0}: 2,
		{X: 7, Y: 0}: 3,
		{X: 0, Y: 1}: 7,
		{X: 1, Y: 1}: 8,
		{X: 2, Y: 1}: 1,
		{X: 3, Y: 1}: 2,
		{X: 4, Y: 1}: 1,
		{X: 5, Y: 1}: 8,
		{X: 6, Y: 1}: 7,
		{X: 7, Y: 1}: 4,
		{X: 0, Y: 2}: 8,
		{X: 1, Y: 2}: 7,
		{X: 2, Y: 2}: 4,
		{X: 3, Y: 2}: 3,
		{X: 4, Y: 2}: 0,
		{X: 5, Y: 2}: 9,
		{X: 6, Y: 2}: 6,
		{X: 7, Y: 2}: 5,
		{X: 0, Y: 3}: 9,
		{X: 1, Y: 3}: 6,
		{X: 2, Y: 3}: 5,
		{X: 3, Y: 3}: 4,
		{X: 4, Y: 3}: 9,
		{X: 5, Y: 3}: 8,
		{X: 6, Y: 3}: 7,
		{X: 7, Y: 3}: 4,
		{X: 0, Y: 4}: 4,
		{X: 1, Y: 4}: 5,
		{X: 2, Y: 4}: 6,
		{X: 3, Y: 4}: 7,
		{X: 4, Y: 4}: 8,
		{X: 5, Y: 4}: 9,
		{X: 6, Y: 4}: 0,
		{X: 7, Y: 4}: 3,
		{X: 0, Y: 5}: 3,
		{X: 1, Y: 5}: 2,
		{X: 2, Y: 5}: 0,
		{X: 3, Y: 5}: 1,
		{X: 4, Y: 5}: 9,
		{X: 5, Y: 5}: 0,
		{X: 6, Y: 5}: 1,
		{X: 7, Y: 5}: 2,
		{X: 0, Y: 6}: 0,
		{X: 1, Y: 6}: 1,
		{X: 2, Y: 6}: 3,
		{X: 3, Y: 6}: 2,
		{X: 4, Y: 6}: 9,
		{X: 5, Y: 6}: 8,
		{X: 6, Y: 6}: 0,
		{X: 7, Y: 6}: 1,
		{X: 0, Y: 7}: 1,
		{X: 1, Y: 7}: 0,
		{X: 2, Y: 7}: 4,
		{X: 3, Y: 7}: 5,
		{X: 4, Y: 7}: 6,
		{X: 5, Y: 7}: 7,
		{X: 6, Y: 7}: 3,
		{X: 7, Y: 7}: 2,
	},
}

func Test_parseInput(t *testing.T) {
	t.Run("returns an error if the input contains a non-numeric character", func(t *testing.T) {
		input := []string{
			"89010123",
			"78121874",
			"87430965",
			"96549874",
			"45678l03",
			"32019012",
			"01329801",
			"10456732",
		}

		tm, err := parseInput(input)
		assert.Error(t, err)
		assert.Nil(t, tm)
	})

	t.Run("returns a TopMap for a given input", func(t *testing.T) {
		input := []string{
			"89010123",
			"78121874",
			"87430965",
			"96549874",
			"45678903",
			"32019012",
			"01329801",
			"10456732",
		}

		want := &TopMap{
			Grid:   exampleGrid,
			trails: map[graph.Co]map[graph.Co]map[string]bool{},
		}

		tm, err := parseInput(input)
		assert.NoError(t, err)
		assert.Equal(t, want, tm)
	})
}

func test_makePathString(t *testing.T) {
	t.Run("returns the correct path string for a given slice of coordinates", func(t *testing.T) {
		path := []graph.Co{
			{X: 1, Y: 0},
			{X: 4, Y: 4},
			{X: 999, Y: -1},
			{X: 2, Y: 3},
			{X: 0, Y: 0},
			{X: 8, Y: 9876},
		}
		want := "X:1,Y:0X:4,Y:4X:999,Y:-1X:2,Y:3X:0,Y:0X:8,Y:9876"
		pathString := makePathString(path)
		assert.Equal(t, want, pathString)
	})
}

func TestTopMap_findSolutions(t *testing.T) {
	t.Run("returns an error if the input contains a non numeric character", func(t *testing.T) {
		input := []string{
			"89010123",
			"78121874",
			"87430965",
			"96549874",
			"45678903",
			"32019012",
			"0.329801",
			"10456732",
		}
		_, _, err := findSolutions(input)
		assert.Error(t, err)
	})

	t.Run("finds solutions for parts 1 and 2 for a given input", func(t *testing.T) {
		input := []string{
			"89010123",
			"78121874",
			"87430965",
			"96549874",
			"45678903",
			"32019012",
			"01329801",
			"10456732",
		}

		part1, part2, err := findSolutions(input)
		assert.NoError(t, err)
		assert.Equal(t, 36, part1)
		assert.Equal(t, 81, part2)
	})
}
