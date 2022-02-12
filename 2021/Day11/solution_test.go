package main

import (
	"reflect"
	"testing"
)

var adventOfCodeExampleInput = []string{
	"5483143223",
	"2745854711",
	"5264556173",
	"6141336146",
	"6357385478",
	"4167524645",
	"2176841721",
	"6882881134",
	"4846848554",
	"5283751526",
}

var adventOfCodeExampleGrid1 = Grid{
	{X: 0, Y: 0}: 5,
	{X: 1, Y: 0}: 4,
	{X: 2, Y: 0}: 8,
	{X: 3, Y: 0}: 3,
	{X: 4, Y: 0}: 1,
	{X: 5, Y: 0}: 4,
	{X: 6, Y: 0}: 3,
	{X: 7, Y: 0}: 2,
	{X: 8, Y: 0}: 2,
	{X: 9, Y: 0}: 3,
	{X: 0, Y: 1}: 2,
	{X: 1, Y: 1}: 7,
	{X: 2, Y: 1}: 4,
	{X: 3, Y: 1}: 5,
	{X: 4, Y: 1}: 8,
	{X: 5, Y: 1}: 5,
	{X: 6, Y: 1}: 4,
	{X: 7, Y: 1}: 7,
	{X: 8, Y: 1}: 1,
	{X: 9, Y: 1}: 1,
	{X: 0, Y: 2}: 5,
	{X: 1, Y: 2}: 2,
	{X: 2, Y: 2}: 6,
	{X: 3, Y: 2}: 4,
	{X: 4, Y: 2}: 5,
	{X: 5, Y: 2}: 5,
	{X: 6, Y: 2}: 6,
	{X: 7, Y: 2}: 1,
	{X: 8, Y: 2}: 7,
	{X: 9, Y: 2}: 3,
	{X: 0, Y: 3}: 6,
	{X: 1, Y: 3}: 1,
	{X: 2, Y: 3}: 4,
	{X: 3, Y: 3}: 1,
	{X: 4, Y: 3}: 3,
	{X: 5, Y: 3}: 3,
	{X: 6, Y: 3}: 6,
	{X: 7, Y: 3}: 1,
	{X: 8, Y: 3}: 4,
	{X: 9, Y: 3}: 6,
	{X: 0, Y: 4}: 6,
	{X: 1, Y: 4}: 3,
	{X: 2, Y: 4}: 5,
	{X: 3, Y: 4}: 7,
	{X: 4, Y: 4}: 3,
	{X: 5, Y: 4}: 8,
	{X: 6, Y: 4}: 5,
	{X: 7, Y: 4}: 4,
	{X: 8, Y: 4}: 7,
	{X: 9, Y: 4}: 8,
	{X: 0, Y: 5}: 4,
	{X: 1, Y: 5}: 1,
	{X: 2, Y: 5}: 6,
	{X: 3, Y: 5}: 7,
	{X: 4, Y: 5}: 5,
	{X: 5, Y: 5}: 2,
	{X: 6, Y: 5}: 4,
	{X: 7, Y: 5}: 6,
	{X: 8, Y: 5}: 4,
	{X: 9, Y: 5}: 5,
	{X: 0, Y: 6}: 2,
	{X: 1, Y: 6}: 1,
	{X: 2, Y: 6}: 7,
	{X: 3, Y: 6}: 6,
	{X: 4, Y: 6}: 8,
	{X: 5, Y: 6}: 4,
	{X: 6, Y: 6}: 1,
	{X: 7, Y: 6}: 7,
	{X: 8, Y: 6}: 2,
	{X: 9, Y: 6}: 1,
	{X: 0, Y: 7}: 6,
	{X: 1, Y: 7}: 8,
	{X: 2, Y: 7}: 8,
	{X: 3, Y: 7}: 2,
	{X: 4, Y: 7}: 8,
	{X: 5, Y: 7}: 8,
	{X: 6, Y: 7}: 1,
	{X: 7, Y: 7}: 1,
	{X: 8, Y: 7}: 3,
	{X: 9, Y: 7}: 4,
	{X: 0, Y: 8}: 4,
	{X: 1, Y: 8}: 8,
	{X: 2, Y: 8}: 4,
	{X: 3, Y: 8}: 6,
	{X: 4, Y: 8}: 8,
	{X: 5, Y: 8}: 4,
	{X: 6, Y: 8}: 8,
	{X: 7, Y: 8}: 5,
	{X: 8, Y: 8}: 5,
	{X: 9, Y: 8}: 4,
	{X: 0, Y: 9}: 5,
	{X: 1, Y: 9}: 2,
	{X: 2, Y: 9}: 8,
	{X: 3, Y: 9}: 3,
	{X: 4, Y: 9}: 7,
	{X: 5, Y: 9}: 5,
	{X: 6, Y: 9}: 1,
	{X: 7, Y: 9}: 5,
	{X: 8, Y: 9}: 2,
	{X: 9, Y: 9}: 6,
}

var adventOfCodeExampleGrid2 = Grid{
	{X: 0, Y: 0}: 6,
	{X: 1, Y: 0}: 5,
	{X: 2, Y: 0}: 9,
	{X: 3, Y: 0}: 4,
	{X: 4, Y: 0}: 2,
	{X: 5, Y: 0}: 5,
	{X: 6, Y: 0}: 4,
	{X: 7, Y: 0}: 3,
	{X: 8, Y: 0}: 3,
	{X: 9, Y: 0}: 4,
	{X: 0, Y: 1}: 3,
	{X: 1, Y: 1}: 8,
	{X: 2, Y: 1}: 5,
	{X: 3, Y: 1}: 6,
	{X: 4, Y: 1}: 9,
	{X: 5, Y: 1}: 6,
	{X: 6, Y: 1}: 5,
	{X: 7, Y: 1}: 8,
	{X: 8, Y: 1}: 2,
	{X: 9, Y: 1}: 2,
	{X: 0, Y: 2}: 6,
	{X: 1, Y: 2}: 3,
	{X: 2, Y: 2}: 7,
	{X: 3, Y: 2}: 5,
	{X: 4, Y: 2}: 6,
	{X: 5, Y: 2}: 6,
	{X: 6, Y: 2}: 7,
	{X: 7, Y: 2}: 2,
	{X: 8, Y: 2}: 8,
	{X: 9, Y: 2}: 4,
	{X: 0, Y: 3}: 7,
	{X: 1, Y: 3}: 2,
	{X: 2, Y: 3}: 5,
	{X: 3, Y: 3}: 2,
	{X: 4, Y: 3}: 4,
	{X: 5, Y: 3}: 4,
	{X: 6, Y: 3}: 7,
	{X: 7, Y: 3}: 2,
	{X: 8, Y: 3}: 5,
	{X: 9, Y: 3}: 7,
	{X: 0, Y: 4}: 7,
	{X: 1, Y: 4}: 4,
	{X: 2, Y: 4}: 6,
	{X: 3, Y: 4}: 8,
	{X: 4, Y: 4}: 4,
	{X: 5, Y: 4}: 9,
	{X: 6, Y: 4}: 6,
	{X: 7, Y: 4}: 5,
	{X: 8, Y: 4}: 8,
	{X: 9, Y: 4}: 9,
	{X: 0, Y: 5}: 5,
	{X: 1, Y: 5}: 2,
	{X: 2, Y: 5}: 7,
	{X: 3, Y: 5}: 8,
	{X: 4, Y: 5}: 6,
	{X: 5, Y: 5}: 3,
	{X: 6, Y: 5}: 5,
	{X: 7, Y: 5}: 7,
	{X: 8, Y: 5}: 5,
	{X: 9, Y: 5}: 6,
	{X: 0, Y: 6}: 3,
	{X: 1, Y: 6}: 2,
	{X: 2, Y: 6}: 8,
	{X: 3, Y: 6}: 7,
	{X: 4, Y: 6}: 9,
	{X: 5, Y: 6}: 5,
	{X: 6, Y: 6}: 2,
	{X: 7, Y: 6}: 8,
	{X: 8, Y: 6}: 3,
	{X: 9, Y: 6}: 2,
	{X: 0, Y: 7}: 7,
	{X: 1, Y: 7}: 9,
	{X: 2, Y: 7}: 9,
	{X: 3, Y: 7}: 3,
	{X: 4, Y: 7}: 9,
	{X: 5, Y: 7}: 9,
	{X: 6, Y: 7}: 2,
	{X: 7, Y: 7}: 2,
	{X: 8, Y: 7}: 4,
	{X: 9, Y: 7}: 5,
	{X: 0, Y: 8}: 5,
	{X: 1, Y: 8}: 9,
	{X: 2, Y: 8}: 5,
	{X: 3, Y: 8}: 7,
	{X: 4, Y: 8}: 9,
	{X: 5, Y: 8}: 5,
	{X: 6, Y: 8}: 9,
	{X: 7, Y: 8}: 6,
	{X: 8, Y: 8}: 6,
	{X: 9, Y: 8}: 5,
	{X: 0, Y: 9}: 6,
	{X: 1, Y: 9}: 3,
	{X: 2, Y: 9}: 9,
	{X: 3, Y: 9}: 4,
	{X: 4, Y: 9}: 8,
	{X: 5, Y: 9}: 6,
	{X: 6, Y: 9}: 2,
	{X: 7, Y: 9}: 6,
	{X: 8, Y: 9}: 3,
	{X: 9, Y: 9}: 7,
}

var adventOfCodeExampleGrid3 = Grid{
	{X: 0, Y: 0}: 8,
	{X: 1, Y: 0}: 8,
	{X: 2, Y: 0}: 0,
	{X: 3, Y: 0}: 7,
	{X: 4, Y: 0}: 4,
	{X: 5, Y: 0}: 7,
	{X: 6, Y: 0}: 6,
	{X: 7, Y: 0}: 5,
	{X: 8, Y: 0}: 5,
	{X: 9, Y: 0}: 5,
	{X: 0, Y: 1}: 5,
	{X: 1, Y: 1}: 0,
	{X: 2, Y: 1}: 8,
	{X: 3, Y: 1}: 9,
	{X: 4, Y: 1}: 0,
	{X: 5, Y: 1}: 8,
	{X: 6, Y: 1}: 7,
	{X: 7, Y: 1}: 0,
	{X: 8, Y: 1}: 5,
	{X: 9, Y: 1}: 4,
	{X: 0, Y: 2}: 8,
	{X: 1, Y: 2}: 5,
	{X: 2, Y: 2}: 9,
	{X: 3, Y: 2}: 7,
	{X: 4, Y: 2}: 8,
	{X: 5, Y: 2}: 8,
	{X: 6, Y: 2}: 9,
	{X: 7, Y: 2}: 6,
	{X: 8, Y: 2}: 0,
	{X: 9, Y: 2}: 8,
	{X: 0, Y: 3}: 8,
	{X: 1, Y: 3}: 4,
	{X: 2, Y: 3}: 8,
	{X: 3, Y: 3}: 5,
	{X: 4, Y: 3}: 7,
	{X: 5, Y: 3}: 6,
	{X: 6, Y: 3}: 9,
	{X: 7, Y: 3}: 6,
	{X: 8, Y: 3}: 0,
	{X: 9, Y: 3}: 0,
	{X: 0, Y: 4}: 8,
	{X: 1, Y: 4}: 7,
	{X: 2, Y: 4}: 0,
	{X: 3, Y: 4}: 0,
	{X: 4, Y: 4}: 9,
	{X: 5, Y: 4}: 0,
	{X: 6, Y: 4}: 8,
	{X: 7, Y: 4}: 8,
	{X: 8, Y: 4}: 0,
	{X: 9, Y: 4}: 0,
	{X: 0, Y: 5}: 6,
	{X: 1, Y: 5}: 6,
	{X: 2, Y: 5}: 0,
	{X: 3, Y: 5}: 0,
	{X: 4, Y: 5}: 0,
	{X: 5, Y: 5}: 8,
	{X: 6, Y: 5}: 8,
	{X: 7, Y: 5}: 9,
	{X: 8, Y: 5}: 8,
	{X: 9, Y: 5}: 9,
	{X: 0, Y: 6}: 6,
	{X: 1, Y: 6}: 8,
	{X: 2, Y: 6}: 0,
	{X: 3, Y: 6}: 0,
	{X: 4, Y: 6}: 0,
	{X: 5, Y: 6}: 0,
	{X: 6, Y: 6}: 5,
	{X: 7, Y: 6}: 9,
	{X: 8, Y: 6}: 4,
	{X: 9, Y: 6}: 3,
	{X: 0, Y: 7}: 0,
	{X: 1, Y: 7}: 0,
	{X: 2, Y: 7}: 0,
	{X: 3, Y: 7}: 0,
	{X: 4, Y: 7}: 0,
	{X: 5, Y: 7}: 0,
	{X: 6, Y: 7}: 7,
	{X: 7, Y: 7}: 4,
	{X: 8, Y: 7}: 5,
	{X: 9, Y: 7}: 6,
	{X: 0, Y: 8}: 9,
	{X: 1, Y: 8}: 0,
	{X: 2, Y: 8}: 0,
	{X: 3, Y: 8}: 0,
	{X: 4, Y: 8}: 0,
	{X: 5, Y: 8}: 0,
	{X: 6, Y: 8}: 0,
	{X: 7, Y: 8}: 8,
	{X: 8, Y: 8}: 7,
	{X: 9, Y: 8}: 6,
	{X: 0, Y: 9}: 8,
	{X: 1, Y: 9}: 7,
	{X: 2, Y: 9}: 0,
	{X: 3, Y: 9}: 0,
	{X: 4, Y: 9}: 0,
	{X: 5, Y: 9}: 0,
	{X: 6, Y: 9}: 6,
	{X: 7, Y: 9}: 8,
	{X: 8, Y: 9}: 4,
	{X: 9, Y: 9}: 8,
}

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  Grid
	}{
		{
			name:  "correctly parses given input to a grid",
			input: adventOfCodeExampleInput,
			want:  adventOfCodeExampleGrid1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_followStep(t *testing.T) {
	tests := []struct {
		name  string
		g     Grid
		want  int
		want1 Grid
	}{
		{
			name:  "follows step, advent of code example step 1",
			g:     adventOfCodeExampleGrid1,
			want:  0,
			want1: adventOfCodeExampleGrid2,
		},
		{
			name:  "follows step, advent of code example step 2",
			g:     adventOfCodeExampleGrid2,
			want:  35,
			want1: adventOfCodeExampleGrid3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.g
			if got := g.followStep(); got != tt.want {
				t.Errorf("Grid.followStep() = %v, want %v", got, tt.want)
			}
			for co, v := range g {
				if v != tt.want1[co] {
					t.Errorf("%v, %d, %d", co, v, tt.want1[co])
				}
			}
			if !reflect.DeepEqual(g, tt.want1) {
				t.Errorf("Grid.followStep(), %v, want %v", g, tt.want1)
			}
		})
	}
}

func TestGrid_isSynchronised(t *testing.T) {
	tests := []struct {
		name string
		g    Grid
		want bool
	}{
		{
			name: "returns false if a value in grid is not 0",
			g: Grid{
				{X: 0, Y: 0}: 0,
				{X: 1, Y: 0}: 0,
				{X: 2, Y: 0}: 0,
				{X: 0, Y: 1}: 0,
				{X: 1, Y: 1}: 0,
				{X: 2, Y: 1}: 6,
				{X: 0, Y: 2}: 7,
				{X: 1, Y: 2}: 0,
				{X: 2, Y: 2}: 3,
			},
			want: false,
		},
		{
			name: "returns true if all values in a grid are 0",
			g: Grid{
				{X: 0, Y: 0}: 0,
				{X: 1, Y: 0}: 0,
				{X: 2, Y: 0}: 0,
				{X: 0, Y: 1}: 0,
				{X: 1, Y: 1}: 0,
				{X: 2, Y: 1}: 0,
				{X: 0, Y: 2}: 0,
				{X: 1, Y: 2}: 0,
				{X: 2, Y: 2}: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.isSynchronised(); got != tt.want {
				t.Errorf("Grid.isSynchronised() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSolution(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
		want1 int
	}{
		{
			name:  "find solutions for parts 1 and 2, advent of code example",
			input: adventOfCodeExampleInput,
			want:  1656,
			want1: 195,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findSolution(tt.input)
			if got != tt.want {
				t.Errorf("findSolution() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findSolution() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
