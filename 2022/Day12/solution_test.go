package main

import (
	"Advent-of-Code/graph"
	"Advent-of-Code/graph/dijkstra"
	"testing"

	"github.com/stretchr/testify/assert"
)

var aocExampleInput = []string{
	"Sabqponm",
	"abcryxxl",
	"accszExk",
	"acctuvwj",
	"abdefghi",
}

var aocExampleTerrain = dijkstra.Graph{
	Grid: map[graph.Co]int{
		{X: 0, Y: 0}: 1,
		{X: 1, Y: 0}: 1,
		{X: 2, Y: 0}: 2,
		{X: 3, Y: 0}: 17,
		{X: 4, Y: 0}: 16,
		{X: 5, Y: 0}: 15,
		{X: 6, Y: 0}: 14,
		{X: 7, Y: 0}: 13,
		{X: 0, Y: 1}: 1,
		{X: 1, Y: 1}: 2,
		{X: 2, Y: 1}: 3,
		{X: 3, Y: 1}: 18,
		{X: 4, Y: 1}: 25,
		{X: 5, Y: 1}: 24,
		{X: 6, Y: 1}: 24,
		{X: 7, Y: 1}: 12,
		{X: 0, Y: 2}: 1,
		{X: 1, Y: 2}: 3,
		{X: 2, Y: 2}: 3,
		{X: 3, Y: 2}: 19,
		{X: 4, Y: 2}: 26,
		{X: 5, Y: 2}: 26,
		{X: 6, Y: 2}: 24,
		{X: 7, Y: 2}: 11,
		{X: 0, Y: 3}: 1,
		{X: 1, Y: 3}: 3,
		{X: 2, Y: 3}: 3,
		{X: 3, Y: 3}: 20,
		{X: 4, Y: 3}: 21,
		{X: 5, Y: 3}: 22,
		{X: 6, Y: 3}: 23,
		{X: 7, Y: 3}: 10,
		{X: 0, Y: 4}: 1,
		{X: 1, Y: 4}: 2,
		{X: 2, Y: 4}: 4,
		{X: 3, Y: 4}: 5,
		{X: 4, Y: 4}: 6,
		{X: 5, Y: 4}: 7,
		{X: 6, Y: 4}: 8,
		{X: 7, Y: 4}: 9,
	},
	Nodes: map[graph.Co][]dijkstra.Edge{
		{X: 0, Y: 0}: {{Node: graph.Co{X: 0, Y: 1}, Weight: 1}, {Node: graph.Co{X: 1, Y: 0}, Weight: 1}},
		{X: 1, Y: 0}: {{Node: graph.Co{X: 0, Y: 0}, Weight: 1}, {Node: graph.Co{X: 2, Y: 0}, Weight: 1}, {Node: graph.Co{X: 1, Y: 1}, Weight: 1}},
		{X: 2, Y: 0}: {{Node: graph.Co{X: 1, Y: 0}, Weight: 1}, {Node: graph.Co{X: 2, Y: 1}, Weight: 1}},
		{X: 3, Y: 0}: {{Node: graph.Co{X: 2, Y: 0}, Weight: 1}, {Node: graph.Co{X: 3, Y: 1}, Weight: 1}, {Node: graph.Co{X: 4, Y: 0}, Weight: 1}},
		{X: 4, Y: 0}: {{Node: graph.Co{X: 3, Y: 0}, Weight: 1}, {Node: graph.Co{X: 5, Y: 0}, Weight: 1}},
		{X: 5, Y: 0}: {{Node: graph.Co{X: 4, Y: 0}, Weight: 1}, {Node: graph.Co{X: 6, Y: 0}, Weight: 1}},
		{X: 6, Y: 0}: {{Node: graph.Co{X: 5, Y: 0}, Weight: 1}, {Node: graph.Co{X: 7, Y: 0}, Weight: 1}},
		{X: 7, Y: 0}: {{Node: graph.Co{X: 6, Y: 0}, Weight: 1}, {Node: graph.Co{X: 7, Y: 1}, Weight: 1}},

		{X: 0, Y: 1}: {{Node: graph.Co{X: 0, Y: 0}, Weight: 1}, {Node: graph.Co{X: 1, Y: 1}, Weight: 1}, {Node: graph.Co{X: 0, Y: 2}, Weight: 1}},
		{X: 1, Y: 1}: {{Node: graph.Co{X: 0, Y: 1}, Weight: 1}, {Node: graph.Co{X: 1, Y: 0}, Weight: 1}, {Node: graph.Co{X: 2, Y: 1}, Weight: 1}, {Node: graph.Co{X: 1, Y: 2}, Weight: 1}},
		{X: 2, Y: 1}: {{Node: graph.Co{X: 2, Y: 0}, Weight: 1}, {Node: graph.Co{X: 2, Y: 2}, Weight: 1}, {Node: graph.Co{X: 1, Y: 1}, Weight: 1}},
		{X: 3, Y: 1}: {{Node: graph.Co{X: 3, Y: 0}, Weight: 1}, {Node: graph.Co{X: 3, Y: 2}, Weight: 1}, {Node: graph.Co{X: 2, Y: 1}, Weight: 1}},
		{X: 4, Y: 1}: {{Node: graph.Co{X: 4, Y: 0}, Weight: 1}, {Node: graph.Co{X: 4, Y: 2}, Weight: 1}, {Node: graph.Co{X: 3, Y: 1}, Weight: 1}, {Node: graph.Co{X: 5, Y: 1}, Weight: 1}},
		{X: 5, Y: 1}: {{Node: graph.Co{X: 4, Y: 1}, Weight: 1}, {Node: graph.Co{X: 6, Y: 1}, Weight: 1}, {Node: graph.Co{X: 5, Y: 0}, Weight: 1}},
		{X: 6, Y: 1}: {{Node: graph.Co{X: 6, Y: 0}, Weight: 1}, {Node: graph.Co{X: 6, Y: 2}, Weight: 1}, {Node: graph.Co{X: 5, Y: 1}, Weight: 1}, {Node: graph.Co{X: 7, Y: 1}, Weight: 1}},
		{X: 7, Y: 1}: {{Node: graph.Co{X: 7, Y: 0}, Weight: 1}, {Node: graph.Co{X: 7, Y: 2}, Weight: 1}},

		{X: 0, Y: 2}: {{Node: graph.Co{X: 0, Y: 1}, Weight: 1}, {Node: graph.Co{X: 0, Y: 3}, Weight: 1}},
		{X: 1, Y: 2}: {{Node: graph.Co{X: 1, Y: 1}, Weight: 1}, {Node: graph.Co{X: 1, Y: 3}, Weight: 1}, {Node: graph.Co{X: 0, Y: 2}, Weight: 1}, {Node: graph.Co{X: 2, Y: 2}, Weight: 1}},
		{X: 2, Y: 2}: {{Node: graph.Co{X: 2, Y: 1}, Weight: 1}, {Node: graph.Co{X: 2, Y: 3}, Weight: 1}, {Node: graph.Co{X: 1, Y: 2}, Weight: 1}},
		{X: 3, Y: 2}: {{Node: graph.Co{X: 3, Y: 1}, Weight: 1}, {Node: graph.Co{X: 3, Y: 3}, Weight: 1}, {Node: graph.Co{X: 2, Y: 2}, Weight: 1}},
		{X: 4, Y: 2}: {{Node: graph.Co{X: 4, Y: 1}, Weight: 1}, {Node: graph.Co{X: 4, Y: 3}, Weight: 1}, {Node: graph.Co{X: 3, Y: 2}, Weight: 1}, {Node: graph.Co{X: 5, Y: 2}, Weight: 1}},
		{X: 5, Y: 2}: {{Node: graph.Co{X: 5, Y: 1}, Weight: 1}, {Node: graph.Co{X: 5, Y: 3}, Weight: 1}, {Node: graph.Co{X: 4, Y: 2}, Weight: 1}, {Node: graph.Co{X: 6, Y: 2}, Weight: 1}},
		{X: 6, Y: 2}: {{Node: graph.Co{X: 6, Y: 1}, Weight: 1}, {Node: graph.Co{X: 6, Y: 3}, Weight: 1}, {Node: graph.Co{X: 7, Y: 2}, Weight: 1}},
		{X: 7, Y: 2}: {{Node: graph.Co{X: 7, Y: 1}, Weight: 1}, {Node: graph.Co{X: 7, Y: 3}, Weight: 1}},

		{X: 0, Y: 3}: {{Node: graph.Co{X: 0, Y: 2}, Weight: 1}, {Node: graph.Co{X: 0, Y: 4}, Weight: 1}},
		{X: 1, Y: 3}: {{Node: graph.Co{X: 1, Y: 2}, Weight: 1}, {Node: graph.Co{X: 1, Y: 4}, Weight: 1}, {Node: graph.Co{X: 0, Y: 3}, Weight: 1}, {Node: graph.Co{X: 2, Y: 3}, Weight: 1}},
		{X: 2, Y: 3}: {{Node: graph.Co{X: 2, Y: 2}, Weight: 1}, {Node: graph.Co{X: 2, Y: 4}, Weight: 1}, {Node: graph.Co{X: 1, Y: 3}, Weight: 1}},
		{X: 3, Y: 3}: {{Node: graph.Co{X: 3, Y: 2}, Weight: 1}, {Node: graph.Co{X: 3, Y: 4}, Weight: 1}, {Node: graph.Co{X: 2, Y: 3}, Weight: 1}, {Node: graph.Co{X: 4, Y: 3}, Weight: 1}},
		{X: 4, Y: 3}: {{Node: graph.Co{X: 3, Y: 3}, Weight: 1}, {Node: graph.Co{X: 5, Y: 3}, Weight: 1}, {Node: graph.Co{X: 4, Y: 4}, Weight: 1}},
		{X: 5, Y: 3}: {{Node: graph.Co{X: 4, Y: 3}, Weight: 1}, {Node: graph.Co{X: 6, Y: 3}, Weight: 1}, {Node: graph.Co{X: 5, Y: 4}, Weight: 1}},
		{X: 6, Y: 3}: {{Node: graph.Co{X: 6, Y: 2}, Weight: 1}, {Node: graph.Co{X: 6, Y: 4}, Weight: 1}, {Node: graph.Co{X: 5, Y: 3}, Weight: 1}, {Node: graph.Co{X: 7, Y: 3}, Weight: 1}},
		{X: 7, Y: 3}: {{Node: graph.Co{X: 7, Y: 2}, Weight: 1}, {Node: graph.Co{X: 7, Y: 4}, Weight: 1}},

		{X: 0, Y: 4}: {{Node: graph.Co{X: 0, Y: 3}, Weight: 1}, {Node: graph.Co{X: 1, Y: 4}, Weight: 1}},
		{X: 1, Y: 4}: {{Node: graph.Co{X: 0, Y: 4}, Weight: 1}, {Node: graph.Co{X: 1, Y: 3}, Weight: 1}},
		{X: 2, Y: 4}: {{Node: graph.Co{X: 1, Y: 4}, Weight: 1}, {Node: graph.Co{X: 3, Y: 4}, Weight: 1}, {Node: graph.Co{X: 2, Y: 3}, Weight: 1}},
		{X: 3, Y: 4}: {{Node: graph.Co{X: 2, Y: 4}, Weight: 1}, {Node: graph.Co{X: 4, Y: 4}, Weight: 1}},
		{X: 4, Y: 4}: {{Node: graph.Co{X: 3, Y: 4}, Weight: 1}, {Node: graph.Co{X: 5, Y: 4}, Weight: 1}},
		{X: 5, Y: 4}: {{Node: graph.Co{X: 4, Y: 4}, Weight: 1}, {Node: graph.Co{X: 6, Y: 4}, Weight: 1}},
		{X: 6, Y: 4}: {{Node: graph.Co{X: 5, Y: 4}, Weight: 1}, {Node: graph.Co{X: 7, Y: 4}, Weight: 1}},
		{X: 7, Y: 4}: {{Node: graph.Co{X: 6, Y: 4}, Weight: 1}, {Node: graph.Co{X: 7, Y: 3}, Weight: 1}},
	},
	MaxX: 7,
	MaxY: 4,
}

func TestParseInput(t *testing.T) {
	t.Run("creates terrain from given input, advent of code example", func(t *testing.T) {
		terrain, origin, destination := parseInput(aocExampleInput)
		assert.Equal(t, aocExampleTerrain.Grid, terrain.Grid)
		assert.Len(t, aocExampleTerrain.Nodes, len(terrain.Nodes))
		for co, nodes := range aocExampleTerrain.Nodes {
			assert.ElementsMatch(t, nodes, terrain.Nodes[co])
		}
		assert.Equal(t, aocExampleTerrain.MaxX, terrain.MaxX)
		assert.Equal(t, aocExampleTerrain.MaxY, terrain.MaxY)
		assert.Equal(t, graph.Co{X: 0, Y: 0}, origin)
		assert.Equal(t, graph.Co{X: 5, Y: 2}, destination)
	})
}

func TestFindSolutions(t *testing.T) {
	t.Run("returns an error if a path can't be found between origin and destination", func(t *testing.T) {
		input := []string{
			"Sabqponm",
			"abcryxxl",
			"accsaExk",
			"acctuvwj",
			"abdefghi",
		}
		got, got1, err := findSolutions(input)
		assert.Equal(t, -1, got)
		assert.Equal(t, -1, got1)
		assert.Error(t, err)
	})

	t.Run("returns an error if a path can't be found between origin and destination", func(t *testing.T) {
		got, got1, err := findSolutions(aocExampleInput)
		assert.Equal(t, 31, got)
		assert.Equal(t, 29, got1)
		assert.NoError(t, err)
	})
}
