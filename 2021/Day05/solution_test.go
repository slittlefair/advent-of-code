package main

import (
	"Advent-of-Code/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convertToCos(t *testing.T) {
	type args struct {
		minX int
		maxX int
		minY int
		maxY int
	}
	tests := []struct {
		name string
		args args
		want []graph.Co
	}{
		{
			name: "converts a horizontal line to coordinates, advent of code example",
			args: args{
				minX: 0,
				maxX: 2,
				minY: 9,
				maxY: 9,
			},
			want: []graph.Co{
				{X: 0, Y: 9},
				{X: 1, Y: 9},
				{X: 2, Y: 9},
			},
		},
		{
			name: "converts a vertical line to coordinates, advent of code example",
			args: args{
				minX: 7,
				maxX: 7,
				minY: 0,
				maxY: 4,
			},
			want: []graph.Co{
				{X: 7, Y: 0},
				{X: 7, Y: 1},
				{X: 7, Y: 2},
				{X: 7, Y: 3},
				{X: 7, Y: 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertToCos(tt.args.minX, tt.args.maxX, tt.args.minY, tt.args.maxY)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_convertDiagonalToCos(t *testing.T) {
	type args struct {
		startX int
		endX   int
		startY int
		endY   int
	}
	tests := []struct {
		name string
		args args
		want []graph.Co
	}{
		{
			name: "converts a diagonal line (increasing X, increasing Y) to coordinates, advent of code example",
			args: args{
				startX: 1,
				endX:   3,
				startY: 1,
				endY:   3,
			},
			want: []graph.Co{
				{X: 1, Y: 1},
				{X: 2, Y: 2},
				{X: 3, Y: 3},
			},
		},
		{
			name: "converts a diagonal line (increasing X, decreasing Y) to coordinates, advent of code example",
			args: args{
				startX: 5,
				endX:   8,
				startY: 5,
				endY:   2,
			},
			want: []graph.Co{
				{X: 5, Y: 5},
				{X: 6, Y: 4},
				{X: 7, Y: 3},
				{X: 8, Y: 2},
			},
		},
		{
			name: "converts a diagonal line (decreasing X, increasing Y) to coordinates, advent of code example",
			args: args{
				startX: 8,
				endX:   0,
				startY: 0,
				endY:   8,
			},
			want: []graph.Co{
				{X: 8, Y: 0},
				{X: 7, Y: 1},
				{X: 6, Y: 2},
				{X: 5, Y: 3},
				{X: 4, Y: 4},
				{X: 3, Y: 5},
				{X: 2, Y: 6},
				{X: 1, Y: 7},
				{X: 0, Y: 8},
			},
		},
		{
			name: "converts a diagonal line (decreasing X, decreasing Y) to coordinates, advent of code example",
			args: args{
				startX: 6,
				endX:   2,
				startY: 4,
				endY:   0,
			},
			want: []graph.Co{
				{X: 6, Y: 4},
				{X: 5, Y: 3},
				{X: 4, Y: 2},
				{X: 3, Y: 1},
				{X: 2, Y: 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertDiagonalToCos(tt.args.startX, tt.args.endX, tt.args.startY, tt.args.endY)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_validCos(t *testing.T) {
	type args struct {
		input []string
		part2 bool
	}
	tests := []struct {
		name               string
		args               args
		want               []graph.Co
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if there are fewer than 4 numbers in an input line",
			args: args{
				input: []string{
					"1,2 -> 2,3",
					"8,0 -> 8,8",
					"1,1 -> 2,2 -> 3,3",
					"7,8 -> 6,9",
				},
			},
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns coordinates of lines in input, part 1",
			args: args{
				input: []string{
					"0,9 -> 5,9",
					"8,0 -> 0,8",
					"9,4 -> 3,4",
					"6,4 -> 2,0",
				},
				part2: false,
			},
			want: []graph.Co{
				{X: 0, Y: 9},
				{X: 1, Y: 9},
				{X: 2, Y: 9},
				{X: 3, Y: 9},
				{X: 4, Y: 9},
				{X: 5, Y: 9},
				{X: 3, Y: 4},
				{X: 4, Y: 4},
				{X: 5, Y: 4},
				{X: 6, Y: 4},
				{X: 7, Y: 4},
				{X: 8, Y: 4},
				{X: 9, Y: 4},
			},
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns coordinates of lines in input, part 2",
			args: args{
				input: []string{
					"0,9 -> 5,9",
					"8,0 -> 0,8",
					"9,4 -> 3,4",
					"6,4 -> 2,0",
				},
				part2: true,
			},
			want: []graph.Co{
				{X: 0, Y: 9},
				{X: 1, Y: 9},
				{X: 2, Y: 9},
				{X: 3, Y: 9},
				{X: 4, Y: 9},
				{X: 5, Y: 9},
				{X: 8, Y: 0},
				{X: 7, Y: 1},
				{X: 6, Y: 2},
				{X: 5, Y: 3},
				{X: 4, Y: 4},
				{X: 3, Y: 5},
				{X: 2, Y: 6},
				{X: 1, Y: 7},
				{X: 0, Y: 8},
				{X: 3, Y: 4},
				{X: 4, Y: 4},
				{X: 5, Y: 4},
				{X: 6, Y: 4},
				{X: 7, Y: 4},
				{X: 8, Y: 4},
				{X: 9, Y: 4},
				{X: 6, Y: 4},
				{X: 5, Y: 3},
				{X: 4, Y: 2},
				{X: 3, Y: 1},
				{X: 2, Y: 0},
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validCos(tt.args.input, tt.args.part2)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_populateGrid(t *testing.T) {
	t.Run("populates grid from list of coordinates", func(t *testing.T) {
		input := []graph.Co{
			{X: 0, Y: 9},
			{X: 1, Y: 9},
			{X: 2, Y: 9},
			{X: 3, Y: 9},
			{X: 4, Y: 9},
			{X: 5, Y: 9},
			{X: 3, Y: 4},
			{X: 4, Y: 4},
			{X: 5, Y: 4},
			{X: 6, Y: 4},
			{X: 7, Y: 4},
			{X: 8, Y: 4},
			{X: 9, Y: 4},
			{X: 2, Y: 2},
			{X: 2, Y: 1},
			{X: 7, Y: 0},
			{X: 7, Y: 1},
			{X: 7, Y: 2},
			{X: 7, Y: 3},
			{X: 7, Y: 4},
			{X: 0, Y: 9},
			{X: 1, Y: 9},
			{X: 2, Y: 9},
			{X: 3, Y: 4},
			{X: 2, Y: 4},
			{X: 1, Y: 4},
		}
		want := Grid{
			{X: 0, Y: 9}: 2,
			{X: 1, Y: 9}: 2,
			{X: 2, Y: 9}: 2,
			{X: 3, Y: 9}: 1,
			{X: 4, Y: 9}: 1,
			{X: 5, Y: 9}: 1,
			{X: 3, Y: 4}: 2,
			{X: 4, Y: 4}: 1,
			{X: 5, Y: 4}: 1,
			{X: 6, Y: 4}: 1,
			{X: 8, Y: 4}: 1,
			{X: 9, Y: 4}: 1,
			{X: 2, Y: 2}: 1,
			{X: 2, Y: 1}: 1,
			{X: 7, Y: 0}: 1,
			{X: 7, Y: 1}: 1,
			{X: 7, Y: 2}: 1,
			{X: 7, Y: 3}: 1,
			{X: 7, Y: 4}: 2,
			{X: 2, Y: 4}: 1,
			{X: 1, Y: 4}: 1,
		}
		got := populateGrid(input)
		assert.Equal(t, want, got)
	})
}

func TestGrid_countOverlaps(t *testing.T) {
	t.Run("counts number of coordinates in a grid that have a value greater than 1", func(t *testing.T) {
		g := Grid{
			{X: 0, Y: 9}: 2,
			{X: 1, Y: 9}: 2,
			{X: 2, Y: 9}: 2,
			{X: 3, Y: 9}: 1,
			{X: 4, Y: 9}: 1,
			{X: 5, Y: 9}: 1,
			{X: 3, Y: 4}: 2,
			{X: 4, Y: 4}: 1,
			{X: 5, Y: 4}: 1,
			{X: 6, Y: 4}: 1,
			{X: 8, Y: 4}: 1,
			{X: 9, Y: 4}: 1,
			{X: 2, Y: 2}: 1,
			{X: 2, Y: 1}: 1,
			{X: 7, Y: 0}: 1,
			{X: 7, Y: 1}: 1,
			{X: 7, Y: 2}: 1,
			{X: 7, Y: 3}: 1,
			{X: 7, Y: 4}: 2,
			{X: 2, Y: 4}: 1,
			{X: 1, Y: 4}: 1,
		}
		got := g.countOverlaps()
		assert.Equal(t, 5, got)
	})
}

func Test_findSolution(t *testing.T) {
	type args struct {
		input []string
		part2 bool
	}
	tests := []struct {
		name               string
		args               args
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if an input line does not contain 4 numbers",
			args: args{
				input: []string{
					"1,2 -> 2,3",
					"8,0 -> 8,8",
					"1,1 -> 2,2 -> 3,3",
					"7,8 -> 6,9",
				},
			},
			want:               -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns number of overlapping coordinates from input, advent of code example part 1",
			args: args{
				input: []string{
					"0,9 -> 5,9",
					"8,0 -> 0,8",
					"9,4 -> 3,4",
					"2,2 -> 2,1",
					"7,0 -> 7,4",
					"6,4 -> 2,0",
					"0,9 -> 2,9",
					"3,4 -> 1,4",
					"0,0 -> 8,8",
					"5,5 -> 8,2",
				},
				part2: false,
			},
			want:               5,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns number of overlapping coordinates from input, advent of code example part 1",
			args: args{
				input: []string{
					"0,9 -> 5,9",
					"8,0 -> 0,8",
					"9,4 -> 3,4",
					"2,2 -> 2,1",
					"7,0 -> 7,4",
					"6,4 -> 2,0",
					"0,9 -> 2,9",
					"3,4 -> 1,4",
					"0,0 -> 8,8",
					"5,5 -> 8,2",
				},
				part2: true,
			},
			want:               12,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findSolution(tt.args.input, tt.args.part2)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
