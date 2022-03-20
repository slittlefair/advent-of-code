package main

import (
	"Advent-of-Code/graph"
	"reflect"
	"testing"
)

func Test_createGrid(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  grid
	}{
		{
			name:  "returns grid from given input, advent of code example 1",
			input: "..^^.",
			want: grid{
				tiles: map[graph.Co]bool{
					{X: 0}: false,
					{X: 1}: false,
					{X: 2}: true,
					{X: 3}: true,
					{X: 4}: false,
				},
				width:     5,
				safeTiles: 3,
			},
		},
		{
			name:  "returns grid from given input, advent of code example 2",
			input: ".^^.^.^^^^",
			want: grid{
				tiles: map[graph.Co]bool{
					{X: 0}: false,
					{X: 1}: true,
					{X: 2}: true,
					{X: 3}: false,
					{X: 4}: true,
					{X: 5}: false,
					{X: 6}: true,
					{X: 7}: true,
					{X: 8}: true,
					{X: 9}: true,
				},
				width:     10,
				safeTiles: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createGrid(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grid_isTrap(t *testing.T) {
	tests := []struct {
		name  string
		g     *grid
		x     int
		want  bool
		want1 int
	}{
		{
			name: "returns false and increments safeTiles if given tile is safe, advent of code example 1",
			g: &grid{
				tiles: map[graph.Co]bool{
					{X: 0}: false,
					{X: 1}: false,
					{X: 2}: true,
					{X: 3}: true,
					{X: 4}: false,
				},
				width:     5,
				safeTiles: 3,
			},
			x:     0,
			want:  false,
			want1: 4,
		},
		{
			name: "returns true and does not increment safeTiles if given tile is a trap, advent of code example 2",
			g: &grid{
				tiles: map[graph.Co]bool{
					{X: 0}: false,
					{X: 1}: false,
					{X: 2}: true,
					{X: 3}: true,
					{X: 4}: false,
				},
				width:     5,
				safeTiles: 4,
			},
			x:     2,
			want:  true,
			want1: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.g
			if got := g.isTrap(tt.x); got != tt.want {
				t.Errorf("grid.isTrap() = %v, want %v", got, tt.want)
			}
			if g.safeTiles != tt.want1 {
				t.Errorf("grid.isTrap().safeTiles = %d, want %d", g.safeTiles, tt.want1)
			}
		})
	}
}

func Test_grid_assessRow(t *testing.T) {
	tests := []struct {
		name string
		g    *grid
		want *grid
	}{
		{
			name: "correctly assesses row of the grid, advent of code example 1",
			g: &grid{
				tiles: map[graph.Co]bool{
					{X: 0}: false,
					{X: 1}: false,
					{X: 2}: true,
					{X: 3}: true,
					{X: 4}: false,
				},
				width:     5,
				safeTiles: 3,
			},
			want: &grid{
				tiles: map[graph.Co]bool{
					{X: 0}: false,
					{X: 1}: true,
					{X: 2}: true,
					{X: 3}: true,
					{X: 4}: true,
				},
				width:     5,
				safeTiles: 4,
			},
		},
		{
			name: "correctly assesses row of the grid, advent of code example 2",
			g: &grid{
				tiles: map[graph.Co]bool{
					{X: 0}: false,
					{X: 1}: true,
					{X: 2}: true,
					{X: 3}: true,
					{X: 4}: true,
				},
				width:     5,
				safeTiles: 4,
			},
			want: &grid{
				tiles: map[graph.Co]bool{
					{X: 0}: true,
					{X: 1}: true,
					{X: 2}: false,
					{X: 3}: false,
					{X: 4}: true,
				},
				width:     5,
				safeTiles: 6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.g
			g.assessRow()
			if !reflect.DeepEqual(g, tt.want) {
				t.Errorf("grid.assessRow() = %v, want %v", g, tt.want)
			}
		})
	}
}

func Test_findSolutions(t *testing.T) {
	type args struct {
		input string
		part1 int
		part2 int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "returns correct solution for parts 1 and 2, advent of code example",
			args: args{
				input: ".^^.^.^^^^",
				part1: 9,
				part2: 10,
			},
			want:  35,
			want1: 38,
		},
		{
			name: "returns correct solution for parts 1 and 2, real input (since AoC doesn't provide part2 example)",
			args: args{
				input: "^..^^.^^^..^^.^...^^^^^....^.^..^^^.^.^.^^...^.^.^.^.^^.....^.^^.^.^.^.^.^.^^..^^^^^...^.....^....^.",
				part1: 40,
				part2: 400000,
			},
			want:  2016,
			want1: 19998750,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findSolutions(tt.args.input, tt.args.part1, tt.args.part2)
			if got != tt.want {
				t.Errorf("findSolutions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findSolutions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
