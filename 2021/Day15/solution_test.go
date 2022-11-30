package main

import (
	"Advent-of-Code/graph"
	djk "Advent-of-Code/graph/dijkstra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInput(t *testing.T) {
	type args struct {
		input  []string
		factor int
	}
	tests := []struct {
		name string
		args args
		want *djk.Graph
	}{
		{
			name: "correctly turns given input into a grid, factor 1",
			args: args{
				input: []string{
					"27",
					"19",
				},
				factor: 1,
			},
			want: &djk.Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}: 2,
					{X: 1, Y: 0}: 7,
					{X: 0, Y: 1}: 1,
					{X: 1, Y: 1}: 9,
				},
				Nodes: map[graph.Co][]djk.Edge{
					{X: 0, Y: 0}: {{Node: graph.Co{X: 0, Y: 1}, Weight: 1}, {Node: graph.Co{X: 1, Y: 0}, Weight: 7}},
					{X: 1, Y: 0}: {{Node: graph.Co{X: 0, Y: 0}, Weight: 2}, {Node: graph.Co{X: 1, Y: 1}, Weight: 9}},
					{X: 0, Y: 1}: {{Node: graph.Co{X: 1, Y: 1}, Weight: 9}, {Node: graph.Co{X: 0, Y: 0}, Weight: 2}},
					{X: 1, Y: 1}: {{Node: graph.Co{X: 0, Y: 1}, Weight: 1}, {Node: graph.Co{X: 1, Y: 0}, Weight: 7}},
				},
				MaxX: 1,
				MaxY: 1,
			},
		},
		{
			name: "correctly turns given input into a grid, factor greater than 1",
			args: args{
				input: []string{
					"27",
					"19",
				},
				factor: 2,
			},
			want: &djk.Graph{
				Grid: map[graph.Co]int{
					{X: 0, Y: 0}: 2,
					{X: 1, Y: 0}: 7,
					{X: 2, Y: 0}: 3,
					{X: 3, Y: 0}: 8,
					{X: 0, Y: 1}: 1,
					{X: 1, Y: 1}: 9,
					{X: 2, Y: 1}: 2,
					{X: 3, Y: 1}: 1,
					{X: 0, Y: 2}: 3,
					{X: 1, Y: 2}: 8,
					{X: 2, Y: 2}: 4,
					{X: 3, Y: 2}: 9,
					{X: 0, Y: 3}: 2,
					{X: 1, Y: 3}: 1,
					{X: 2, Y: 3}: 3,
					{X: 3, Y: 3}: 2,
				},
				Nodes: map[graph.Co][]djk.Edge{
					{X: 0, Y: 0}: {{Node: graph.Co{X: 1, Y: 0}, Weight: 7}, {Node: graph.Co{X: 0, Y: 1}, Weight: 1}},
					{X: 1, Y: 0}: {{Node: graph.Co{X: 0, Y: 0}, Weight: 2}, {Node: graph.Co{X: 2, Y: 0}, Weight: 3}, {Node: graph.Co{X: 1, Y: 1}, Weight: 9}},
					{X: 2, Y: 0}: {{Node: graph.Co{X: 1, Y: 0}, Weight: 7}, {Node: graph.Co{X: 3, Y: 0}, Weight: 8}, {Node: graph.Co{X: 2, Y: 1}, Weight: 2}},
					{X: 3, Y: 0}: {{Node: graph.Co{X: 2, Y: 0}, Weight: 3}, {Node: graph.Co{X: 3, Y: 1}, Weight: 1}},

					{X: 0, Y: 1}: {{Node: graph.Co{X: 0, Y: 0}, Weight: 2}, {Node: graph.Co{X: 0, Y: 2}, Weight: 3}, {Node: graph.Co{X: 1, Y: 1}, Weight: 9}},
					{X: 1, Y: 1}: {{Node: graph.Co{X: 1, Y: 0}, Weight: 7}, {Node: graph.Co{X: 1, Y: 2}, Weight: 8}, {Node: graph.Co{X: 0, Y: 1}, Weight: 1}, {Node: graph.Co{X: 2, Y: 1}, Weight: 2}},
					{X: 2, Y: 1}: {{Node: graph.Co{X: 2, Y: 0}, Weight: 3}, {Node: graph.Co{X: 2, Y: 2}, Weight: 4}, {Node: graph.Co{X: 1, Y: 1}, Weight: 9}, {Node: graph.Co{X: 3, Y: 1}, Weight: 1}},
					{X: 3, Y: 1}: {{Node: graph.Co{X: 3, Y: 0}, Weight: 8}, {Node: graph.Co{X: 3, Y: 2}, Weight: 9}, {Node: graph.Co{X: 2, Y: 1}, Weight: 2}},

					{X: 0, Y: 2}: {{Node: graph.Co{X: 0, Y: 1}, Weight: 1}, {Node: graph.Co{X: 0, Y: 3}, Weight: 2}, {Node: graph.Co{X: 1, Y: 2}, Weight: 8}},
					{X: 1, Y: 2}: {{Node: graph.Co{X: 1, Y: 1}, Weight: 9}, {Node: graph.Co{X: 1, Y: 3}, Weight: 1}, {Node: graph.Co{X: 0, Y: 2}, Weight: 3}, {Node: graph.Co{X: 2, Y: 2}, Weight: 4}},
					{X: 2, Y: 2}: {{Node: graph.Co{X: 2, Y: 1}, Weight: 2}, {Node: graph.Co{X: 2, Y: 3}, Weight: 3}, {Node: graph.Co{X: 1, Y: 2}, Weight: 8}, {Node: graph.Co{X: 3, Y: 2}, Weight: 9}},
					{X: 3, Y: 2}: {{Node: graph.Co{X: 3, Y: 1}, Weight: 1}, {Node: graph.Co{X: 3, Y: 3}, Weight: 2}, {Node: graph.Co{X: 2, Y: 2}, Weight: 4}},

					{X: 0, Y: 3}: {{Node: graph.Co{X: 0, Y: 2}, Weight: 3}, {Node: graph.Co{X: 1, Y: 3}, Weight: 1}},
					{X: 1, Y: 3}: {{Node: graph.Co{X: 0, Y: 3}, Weight: 2}, {Node: graph.Co{X: 2, Y: 3}, Weight: 3}, {Node: graph.Co{X: 1, Y: 2}, Weight: 8}},
					{X: 2, Y: 3}: {{Node: graph.Co{X: 1, Y: 3}, Weight: 1}, {Node: graph.Co{X: 3, Y: 3}, Weight: 2}, {Node: graph.Co{X: 2, Y: 2}, Weight: 4}},
					{X: 3, Y: 3}: {{Node: graph.Co{X: 3, Y: 2}, Weight: 9}, {Node: graph.Co{X: 2, Y: 3}, Weight: 3}},
				},
				MaxX: 3,
				MaxY: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseInput(tt.args.input, tt.args.factor)
			assert.Equal(t, tt.want.Grid, got.Grid)
			assert.Equal(t, tt.want.MaxX, got.MaxX)
			assert.Equal(t, tt.want.MaxY, got.MaxY)
			for co, edges := range got.Nodes {
				assert.ElementsMatch(t, tt.want.Nodes[co], edges)
			}
		})
	}
}

func Test_findSolutions(t *testing.T) {
	t.Run("returns solutions for part 1 and part 2, advent of code example", func(t *testing.T) {
		got, got1, err := findSolutions(
			[]string{
				"1163751742",
				"1381373672",
				"2136511328",
				"3694931569",
				"7463417111",
				"1319128137",
				"1359912421",
				"3125421639",
				"1293138521",
				"2311944581",
			},
		)
		assert.NoError(t, err)
		assert.Equal(t, 40, got)
		assert.Equal(t, 315, got1)
	})
}
