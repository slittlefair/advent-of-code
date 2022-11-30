package main

import (
	"Advent-of-Code/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid_evaluateEmptySeat(t *testing.T) {
	type args struct {
		co   graph.Co
		part int
	}
	tests := []struct {
		name string
		g    Grid
		args args
		want string
	}{
		{
			name: "part 1 - returns '#' if an edge empty seat is adjacent only to empty seats",
			g: Grid{
				graph.Co{X: 0, Y: 0}: "L",
				graph.Co{X: 1, Y: 0}: "L",
				graph.Co{X: 2, Y: 0}: "L",
				graph.Co{X: 0, Y: 1}: "L",
				graph.Co{X: 1, Y: 1}: "L",
				graph.Co{X: 2, Y: 1}: "#",
				graph.Co{X: 0, Y: 2}: "#",
				graph.Co{X: 1, Y: 2}: "#",
				graph.Co{X: 2, Y: 2}: "L",
			},
			args: args{co: graph.Co{X: 0, Y: 0},
				part: 1,
			},
			want: "#",
		},
		{
			name: "part 1 - returns '# if an edge empty seat is adjacent only to spaces",
			g: Grid{
				graph.Co{X: 0, Y: 0}: "L",
				graph.Co{X: 1, Y: 0}: ".",
				graph.Co{X: 2, Y: 0}: "#",
				graph.Co{X: 0, Y: 1}: ".",
				graph.Co{X: 1, Y: 1}: ".",
				graph.Co{X: 2, Y: 1}: ".",
				graph.Co{X: 0, Y: 2}: ".",
				graph.Co{X: 1, Y: 2}: "#",
				graph.Co{X: 2, Y: 2}: ".",
			},
			args: args{co: graph.Co{X: 0, Y: 0},
				part: 1,
			},
			want: "#",
		},
		{
			name: "part 1 - returns '#' if an edge empty seat is adjacent only to spaces or empty seats",
			g: Grid{
				graph.Co{X: 0, Y: 0}: "L",
				graph.Co{X: 1, Y: 0}: ".",
				graph.Co{X: 2, Y: 0}: ".",
				graph.Co{X: 0, Y: 1}: "L",
				graph.Co{X: 1, Y: 1}: "L",
				graph.Co{X: 2, Y: 1}: ".",
				graph.Co{X: 0, Y: 2}: "#",
				graph.Co{X: 1, Y: 2}: "L",
				graph.Co{X: 2, Y: 2}: "#",
			},
			args: args{co: graph.Co{X: 0, Y: 0},
				part: 1,
			},
			want: "#",
		},
		{
			name: "part 1 - returns 'L' if an edge empty seat is adjacent to one occupied seat",
			g: Grid{
				graph.Co{X: 0, Y: 0}: "L",
				graph.Co{X: 1, Y: 0}: "#",
				graph.Co{X: 2, Y: 0}: ".",
				graph.Co{X: 0, Y: 1}: "L",
				graph.Co{X: 1, Y: 1}: "L",
				graph.Co{X: 2, Y: 1}: "#",
				graph.Co{X: 0, Y: 2}: "L",
				graph.Co{X: 1, Y: 2}: "L",
				graph.Co{X: 2, Y: 2}: "#",
			},
			args: args{co: graph.Co{X: 0, Y: 0},
				part: 1,
			},
			want: "L",
		},
		{
			name: "part 1 - returns 'L' if an edge empty seat is adjacent to multiple occupied seats",
			g: Grid{
				graph.Co{X: 0, Y: 0}: "L",
				graph.Co{X: 1, Y: 0}: "#",
				graph.Co{X: 2, Y: 0}: ".",
				graph.Co{X: 0, Y: 1}: "#",
				graph.Co{X: 1, Y: 1}: "L",
				graph.Co{X: 2, Y: 1}: "#",
				graph.Co{X: 0, Y: 2}: "L",
				graph.Co{X: 1, Y: 2}: "L",
				graph.Co{X: 2, Y: 2}: "#",
			},
			args: args{co: graph.Co{X: 0, Y: 0},
				part: 1,
			},
			want: "L",
		},
		{
			name: "part 1 - returns '#' if an empty seat is adjacent only to empty seats or spaces",
			g: Grid{
				graph.Co{X: 0, Y: 0}: "L",
				graph.Co{X: 1, Y: 0}: ".",
				graph.Co{X: 2, Y: 0}: ".",
				graph.Co{X: 0, Y: 1}: "L",
				graph.Co{X: 1, Y: 1}: "L",
				graph.Co{X: 2, Y: 1}: ".",
				graph.Co{X: 0, Y: 2}: "L",
				graph.Co{X: 1, Y: 2}: "L",
				graph.Co{X: 2, Y: 2}: ".",
				graph.Co{X: 1, Y: 3}: "#",
			},
			args: args{co: graph.Co{X: 1, Y: 1},
				part: 1,
			},
			want: "#",
		},
		{
			name: "part 1 - returns 'L' if an empty seat is adjacent to one occupied seat",
			g: Grid{
				graph.Co{X: 0, Y: 0}: "L",
				graph.Co{X: 1, Y: 0}: ".",
				graph.Co{X: 2, Y: 0}: "#",
				graph.Co{X: 0, Y: 1}: "L",
				graph.Co{X: 1, Y: 1}: "L",
				graph.Co{X: 2, Y: 1}: ".",
				graph.Co{X: 0, Y: 2}: "L",
				graph.Co{X: 1, Y: 2}: "L",
				graph.Co{X: 2, Y: 2}: ".",
				graph.Co{X: 1, Y: 3}: "#",
			},
			args: args{co: graph.Co{X: 1, Y: 1},
				part: 1,
			},
			want: "L",
		},
		{
			name: "part 1 - returns 'L' if an empty seat is adjacent to multiple occupied seat",
			g: Grid{
				graph.Co{X: 0, Y: 0}: "L",
				graph.Co{X: 1, Y: 0}: "#",
				graph.Co{X: 2, Y: 0}: "#",
				graph.Co{X: 0, Y: 1}: "L",
				graph.Co{X: 1, Y: 1}: "L",
				graph.Co{X: 2, Y: 1}: ".",
				graph.Co{X: 0, Y: 2}: "L",
				graph.Co{X: 1, Y: 2}: "L",
				graph.Co{X: 2, Y: 2}: ".",
				graph.Co{X: 1, Y: 3}: "#",
			},
			args: args{co: graph.Co{X: 1, Y: 1},
				part: 1,
			},
			want: "L",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.g.evaluateEmptySeat(tt.args.co, tt.args.part)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGrid_countOccupiedSeats(t *testing.T) {
	tests := []struct {
		name string
		g    Grid
		want int
	}{
		{
			name: "returns 0 if empty grid",
			g:    Grid{},
			want: 0,
		},
		{
			name: "returns 0 if no occupied seats",
			g: Grid{
				graph.Co{X: 0, Y: 0}: ".",
				graph.Co{X: 0, Y: 1}: "L",
			},
			want: 0,
		},
		{
			name: "returns 1 if 1 occupied seat",
			g: Grid{
				graph.Co{X: 0, Y: 0}: ".",
				graph.Co{X: 0, Y: 1}: "L",
				graph.Co{X: 1, Y: 0}: "#",
			},
			want: 1,
		},
		{
			name: "returns the correct number of multiple occupied seats",
			g: Grid{
				graph.Co{X: 0, Y: 0}: ".",
				graph.Co{X: 0, Y: 1}: "L",
				graph.Co{X: 1, Y: 0}: "#",
				graph.Co{X: 2, Y: 0}: "#",
				graph.Co{X: 3, Y: 1}: "blah",
				graph.Co{X: 4, Y: 0}: "#",
				graph.Co{X: 5, Y: 0}: "#",
				graph.Co{X: 6, Y: 1}: "",
				graph.Co{X: 7, Y: 0}: "#",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.g.countOccupiedSeats()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGrid_parseInput(t *testing.T) {
	tests := []struct {
		name string
		g    Grid
		plan []string
		want Grid
	}{
		{
			name: "single line input",
			g:    Grid{},
			plan: []string{"#.L."},
			want: Grid{
				graph.Co{X: 0, Y: 0}: "#",
				graph.Co{X: 1, Y: 0}: ".",
				graph.Co{X: 3, Y: 0}: ".",
				graph.Co{X: 2, Y: 0}: "L",
			},
		},
		{
			name: "multiple line input",
			g:    Grid{},
			plan: []string{"#.L#", "...L", "12X*"},
			want: Grid{
				graph.Co{X: 0, Y: 0}: "#",
				graph.Co{X: 1, Y: 0}: ".",
				graph.Co{X: 2, Y: 0}: "L",
				graph.Co{X: 3, Y: 0}: "#",
				graph.Co{X: 0, Y: 1}: ".",
				graph.Co{X: 1, Y: 1}: ".",
				graph.Co{X: 2, Y: 1}: ".",
				graph.Co{X: 3, Y: 1}: "L",
				graph.Co{X: 0, Y: 2}: "1",
				graph.Co{X: 1, Y: 2}: "2",
				graph.Co{X: 2, Y: 2}: "X",
				graph.Co{X: 3, Y: 2}: "*",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.parseInput(tt.plan)
			assert.Equal(t, tt.want, tt.g)
		})
	}
}
