package main

import (
	"Advent-of-Code/graph"
	djk "Advent-of-Code/graph/dijkstra"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name               string
		input              []string
		want               int
		want1              graph.Co
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if first line of input is not an int",
			input: []string{
				"abc",
				"7,4",
			},
			want:               -1,
			want1:              graph.Co{},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns number and target co, advent of code example",
			input: []string{
				"10",
				"7,4",
			},
			want:               10,
			want1:              graph.Co{X: 7, Y: 4},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseInput(tt.input)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_isSpace(t *testing.T) {
	tests := []struct {
		co   graph.Co
		want bool
	}{
		{
			co:   graph.Co{X: 7, Y: 3},
			want: false,
		},
		{
			co:   graph.Co{X: 0, Y: 0},
			want: true,
		},
		{
			co:   graph.Co{X: 9, Y: 6},
			want: false,
		},
		{
			co:   graph.Co{X: 6, Y: 4},
			want: true,
		},
		{
			co:   graph.Co{X: 2, Y: 4},
			want: false,
		},
		{
			co:   graph.Co{X: 3, Y: 4},
			want: true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("returns %t for co (%d,%d), advent of code example %d", tt.want, tt.co.X, tt.co.Y, i+1), func(t *testing.T) {
			got := isSpace(tt.co, 10)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_populateOffice(t *testing.T) {
	t.Run("returns graph from given input, advent of code example", func(t *testing.T) {
		got := populateOffice(graph.Co{X: -8, Y: -8}, 10)
		want := &djk.Graph{
			MaxX: 51,
			MaxY: 51,
			Grid: map[graph.Co]int{
				{X: 0, Y: 0}: 1,
				{X: 2, Y: 0}: 1,
				{X: 0, Y: 1}: 1,
				{X: 1, Y: 1}: 1,
				{X: 1, Y: 2}: 1,
				{X: 2, Y: 2}: 1,
			},
			Nodes: map[graph.Co][]djk.Edge{
				{X: 0, Y: 0}: {{Node: graph.Co{X: 0, Y: 1}, Weight: 1}},
				{X: 0, Y: 1}: {{Node: graph.Co{X: 0, Y: 0}, Weight: 1}, {Node: graph.Co{X: 1, Y: 1}, Weight: 1}},
				{X: 1, Y: 1}: {{Node: graph.Co{X: 0, Y: 1}, Weight: 1}, {Node: graph.Co{X: 1, Y: 2}, Weight: 1}},
				{X: 1, Y: 2}: {{Node: graph.Co{X: 2, Y: 2}, Weight: 1}, {Node: graph.Co{X: 1, Y: 1}, Weight: 1}},
				{X: 2, Y: 2}: {{Node: graph.Co{X: 1, Y: 2}, Weight: 1}},
			},
		}
		assert.Equal(t, want.Grid, got.Grid)
		assert.Len(t, got.Nodes, len(want.Nodes))
		for co, edges := range got.Nodes {
			assert.ElementsMatch(t, want.Nodes[co], edges)
		}
		assert.Equal(t, want.MaxX, got.MaxX)
		assert.Equal(t, want.MaxY, got.MaxY)
	})
}

func Test_findSolutions(t *testing.T) {
	tests := []struct {
		name               string
		input              []string
		want               int
		want1              int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if parse input returns an error",
			input: []string{
				"a",
				"2,3",
			},
			want:               -1,
			want1:              -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if a path to target can't be found",
			input: []string{
				"10",
				"5,2",
			},
			want:               -1,
			want1:              -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns correct answers for part1 and part2, advent of code example",
			input: []string{
				"10",
				"7,4",
			},
			want:               11,
			want1:              65,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns correct answers for part1 and part2, real solution (since AoC doesn't provide example part2)",
			input: []string{
				"1350",
				"31,39",
			},
			want:               92,
			want1:              124,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := findSolutions(tt.input)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
