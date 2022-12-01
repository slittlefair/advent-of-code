package graph_test

import (
	"Advent-of-Code/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdjacentCos(t *testing.T) {
	type args struct {
		co               graph.Co
		includeDiagonals bool
	}
	tests := []struct {
		name string
		args args
		want []graph.Co
	}{
		{
			name: "returns adjacent coordinates, including diagonals, of a given coordinate",
			args: args{
				co:               graph.Co{X: 4, Y: 7},
				includeDiagonals: true,
			},
			want: []graph.Co{
				{X: 3, Y: 6},
				{X: 4, Y: 6},
				{X: 5, Y: 6},
				{X: 3, Y: 7},
				{X: 5, Y: 7},
				{X: 3, Y: 8},
				{X: 4, Y: 8},
				{X: 5, Y: 8},
			},
		},
		{
			name: "returns adjacent coordinates, including diagonals, of the origin",
			args: args{
				co:               graph.Co{X: 0, Y: 0},
				includeDiagonals: true,
			},
			want: []graph.Co{
				{X: -1, Y: -1},
				{X: 0, Y: -1},
				{X: 1, Y: -1},
				{X: -1, Y: 0},
				{X: 1, Y: 0},
				{X: -1, Y: 1},
				{X: 0, Y: 1},
				{X: 1, Y: 1},
			},
		},
		{
			name: "returns adjacent coordinates, excluding diagonals, of a given coordinate",
			args: args{
				co:               graph.Co{X: 4, Y: 7},
				includeDiagonals: false,
			},
			want: []graph.Co{
				{X: 4, Y: 6},
				{X: 3, Y: 7},
				{X: 5, Y: 7},
				{X: 4, Y: 8},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := graph.AdjacentCos(tt.args.co, tt.args.includeDiagonals)
			assert.ElementsMatch(t, tt.want, got1)
		})
	}
}

func TestCalculateManhattanDistance(t *testing.T) {
	type args struct {
		co1 graph.Co
		co2 graph.Co
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "returns distance between a coordinate and origin",
			args: args{
				co1: graph.Co{X: 7, Y: 8},
				co2: graph.Co{},
			},
			want: 15,
		},
		{
			name: "returns distance between a positive and negative coordinate",
			args: args{
				co1: graph.Co{X: 9, Y: 1},
				co2: graph.Co{X: -9, Y: -7},
			},
			want: 26,
		},
		{
			name: "returns distance when where difference between the two will be negative",
			args: args{
				co1: graph.Co{X: 1, Y: 1},
				co2: graph.Co{X: 8, Y: 11},
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := graph.CalculateManhattanDistance(tt.args.co1, tt.args.co2)
			assert.Equal(t, tt.want, got)
		})
	}
}
