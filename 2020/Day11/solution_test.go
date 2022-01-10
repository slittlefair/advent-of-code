package main

import (
	utils "Advent-of-Code/utils"
	"reflect"
	"testing"
)

func TestGrid_evaluateEmptySeat(t *testing.T) {
	type args struct {
		co   utils.Co
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
				utils.Co{X: 0, Y: 0}: "L",
				utils.Co{X: 1, Y: 0}: "L",
				utils.Co{X: 2, Y: 0}: "L",
				utils.Co{X: 0, Y: 1}: "L",
				utils.Co{X: 1, Y: 1}: "L",
				utils.Co{X: 2, Y: 1}: "#",
				utils.Co{X: 0, Y: 2}: "#",
				utils.Co{X: 1, Y: 2}: "#",
				utils.Co{X: 2, Y: 2}: "L",
			},
			args: args{co: utils.Co{X: 0, Y: 0},
				part: 1,
			},
			want: "#",
		},
		{
			name: "part 1 - returns '# if an edge empty seat is adjacent only to spaces",
			g: Grid{
				utils.Co{X: 0, Y: 0}: "L",
				utils.Co{X: 1, Y: 0}: ".",
				utils.Co{X: 2, Y: 0}: "#",
				utils.Co{X: 0, Y: 1}: ".",
				utils.Co{X: 1, Y: 1}: ".",
				utils.Co{X: 2, Y: 1}: ".",
				utils.Co{X: 0, Y: 2}: ".",
				utils.Co{X: 1, Y: 2}: "#",
				utils.Co{X: 2, Y: 2}: ".",
			},
			args: args{co: utils.Co{X: 0, Y: 0},
				part: 1,
			},
			want: "#",
		},
		{
			name: "part 1 - returns '#' if an edge empty seat is adjacent only to spaces or empty seats",
			g: Grid{
				utils.Co{X: 0, Y: 0}: "L",
				utils.Co{X: 1, Y: 0}: ".",
				utils.Co{X: 2, Y: 0}: ".",
				utils.Co{X: 0, Y: 1}: "L",
				utils.Co{X: 1, Y: 1}: "L",
				utils.Co{X: 2, Y: 1}: ".",
				utils.Co{X: 0, Y: 2}: "#",
				utils.Co{X: 1, Y: 2}: "L",
				utils.Co{X: 2, Y: 2}: "#",
			},
			args: args{co: utils.Co{X: 0, Y: 0},
				part: 1,
			},
			want: "#",
		},
		{
			name: "part 1 - returns 'L' if an edge empty seat is adjacent to one occupied seat",
			g: Grid{
				utils.Co{X: 0, Y: 0}: "L",
				utils.Co{X: 1, Y: 0}: "#",
				utils.Co{X: 2, Y: 0}: ".",
				utils.Co{X: 0, Y: 1}: "L",
				utils.Co{X: 1, Y: 1}: "L",
				utils.Co{X: 2, Y: 1}: "#",
				utils.Co{X: 0, Y: 2}: "L",
				utils.Co{X: 1, Y: 2}: "L",
				utils.Co{X: 2, Y: 2}: "#",
			},
			args: args{co: utils.Co{X: 0, Y: 0},
				part: 1,
			},
			want: "L",
		},
		{
			name: "part 1 - returns 'L' if an edge empty seat is adjacent to multiple occupied seats",
			g: Grid{
				utils.Co{X: 0, Y: 0}: "L",
				utils.Co{X: 1, Y: 0}: "#",
				utils.Co{X: 2, Y: 0}: ".",
				utils.Co{X: 0, Y: 1}: "#",
				utils.Co{X: 1, Y: 1}: "L",
				utils.Co{X: 2, Y: 1}: "#",
				utils.Co{X: 0, Y: 2}: "L",
				utils.Co{X: 1, Y: 2}: "L",
				utils.Co{X: 2, Y: 2}: "#",
			},
			args: args{co: utils.Co{X: 0, Y: 0},
				part: 1,
			},
			want: "L",
		},
		{
			name: "part 1 - returns '#' if an empty seat is adjacent only to empty seats or spaces",
			g: Grid{
				utils.Co{X: 0, Y: 0}: "L",
				utils.Co{X: 1, Y: 0}: ".",
				utils.Co{X: 2, Y: 0}: ".",
				utils.Co{X: 0, Y: 1}: "L",
				utils.Co{X: 1, Y: 1}: "L",
				utils.Co{X: 2, Y: 1}: ".",
				utils.Co{X: 0, Y: 2}: "L",
				utils.Co{X: 1, Y: 2}: "L",
				utils.Co{X: 2, Y: 2}: ".",
				utils.Co{X: 1, Y: 3}: "#",
			},
			args: args{co: utils.Co{X: 1, Y: 1},
				part: 1,
			},
			want: "#",
		},
		{
			name: "part 1 - returns 'L' if an empty seat is adjacent to one occupied seat",
			g: Grid{
				utils.Co{X: 0, Y: 0}: "L",
				utils.Co{X: 1, Y: 0}: ".",
				utils.Co{X: 2, Y: 0}: "#",
				utils.Co{X: 0, Y: 1}: "L",
				utils.Co{X: 1, Y: 1}: "L",
				utils.Co{X: 2, Y: 1}: ".",
				utils.Co{X: 0, Y: 2}: "L",
				utils.Co{X: 1, Y: 2}: "L",
				utils.Co{X: 2, Y: 2}: ".",
				utils.Co{X: 1, Y: 3}: "#",
			},
			args: args{co: utils.Co{X: 1, Y: 1},
				part: 1,
			},
			want: "L",
		},
		{
			name: "part 1 - returns 'L' if an empty seat is adjacent to multiple occupied seat",
			g: Grid{
				utils.Co{X: 0, Y: 0}: "L",
				utils.Co{X: 1, Y: 0}: "#",
				utils.Co{X: 2, Y: 0}: "#",
				utils.Co{X: 0, Y: 1}: "L",
				utils.Co{X: 1, Y: 1}: "L",
				utils.Co{X: 2, Y: 1}: ".",
				utils.Co{X: 0, Y: 2}: "L",
				utils.Co{X: 1, Y: 2}: "L",
				utils.Co{X: 2, Y: 2}: ".",
				utils.Co{X: 1, Y: 3}: "#",
			},
			args: args{co: utils.Co{X: 1, Y: 1},
				part: 1,
			},
			want: "L",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.evaluateEmptySeat(tt.args.co, tt.args.part); got != tt.want {
				t.Errorf("Grid.evaluateEmptySeat() = %v, want %v", got, tt.want)
			}
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
				utils.Co{X: 0, Y: 0}: ".",
				utils.Co{X: 0, Y: 1}: "L",
			},
			want: 0,
		},
		{
			name: "returns 1 if 1 occupied seat",
			g: Grid{
				utils.Co{X: 0, Y: 0}: ".",
				utils.Co{X: 0, Y: 1}: "L",
				utils.Co{X: 1, Y: 0}: "#",
			},
			want: 1,
		},
		{
			name: "returns the correct number of multiple occupied seats",
			g: Grid{
				utils.Co{X: 0, Y: 0}: ".",
				utils.Co{X: 0, Y: 1}: "L",
				utils.Co{X: 1, Y: 0}: "#",
				utils.Co{X: 2, Y: 0}: "#",
				utils.Co{X: 3, Y: 1}: "blah",
				utils.Co{X: 4, Y: 0}: "#",
				utils.Co{X: 5, Y: 0}: "#",
				utils.Co{X: 6, Y: 1}: "",
				utils.Co{X: 7, Y: 0}: "#",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.countOccupiedSeats(); got != tt.want {
				t.Errorf("Grid.countOccupiedSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_parseInput(t *testing.T) {
	type args struct {
		plan []string
	}
	tests := []struct {
		name string
		g    Grid
		args args
		want Grid
	}{
		{
			name: "single line input",
			g:    Grid{},
			args: args{
				plan: []string{"#.L."},
			},
			want: Grid{
				utils.Co{X: 0, Y: 0}: "#",
				utils.Co{X: 1, Y: 0}: ".",
				utils.Co{X: 3, Y: 0}: ".",
				utils.Co{X: 2, Y: 0}: "L",
			},
		},
		{
			name: "multiple line input",
			g:    Grid{},
			args: args{
				plan: []string{"#.L#", "...L", "12X*"},
			},
			want: Grid{
				utils.Co{X: 0, Y: 0}: "#",
				utils.Co{X: 1, Y: 0}: ".",
				utils.Co{X: 2, Y: 0}: "L",
				utils.Co{X: 3, Y: 0}: "#",
				utils.Co{X: 0, Y: 1}: ".",
				utils.Co{X: 1, Y: 1}: ".",
				utils.Co{X: 2, Y: 1}: ".",
				utils.Co{X: 3, Y: 1}: "L",
				utils.Co{X: 0, Y: 2}: "1",
				utils.Co{X: 1, Y: 2}: "2",
				utils.Co{X: 2, Y: 2}: "X",
				utils.Co{X: 3, Y: 2}: "*",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.parseInput(tt.args.plan)
		})
		if ok := reflect.DeepEqual(tt.g, tt.want); !ok {
			t.Errorf("Grid.parseInput() = %v, want %v", tt.g, tt.want)
		}
	}
}
