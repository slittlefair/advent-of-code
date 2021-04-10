package main

import (
	"reflect"
	"testing"
)

var basicGrid = Grid{
	Coord4D{X: 0, Y: 0, Z: 0, W: 0}:   ".",
	Coord4D{X: 1, Y: 0, Z: 0, W: 0}:   ".",
	Coord4D{X: 2, Y: 0, Z: 0, W: 0}:   "#",
	Coord4D{X: 0, Y: 10, Z: 0, W: 0}:  "#",
	Coord4D{X: 1, Y: 10, Z: 0, W: 0}:  "#",
	Coord4D{X: 2, Y: 10, Z: 0, W: 0}:  "#",
	Coord4D{X: 1, Y: 11, Z: 0, W: 0}:  ".",
	Coord4D{X: 1, Y: 9, Z: 0, W: 0}:   "#",
	Coord4D{X: 0, Y: 20, Z: 0, W: 0}:  "#",
	Coord4D{X: 1, Y: 20, Z: 0, W: 0}:  "#",
	Coord4D{X: 2, Y: 20, Z: 0, W: 0}:  "#",
	Coord4D{X: 0, Y: 20, Z: -1, W: 0}: "#",
	Coord4D{X: 0, Y: 20, Z: 1, W: 0}:  ".",
	Coord4D{X: 0, Y: 20, Z: 1, W: 1}:  "#",
	Coord4D{X: 0, Y: 30, Z: 0, W: 0}:  ".",
	Coord4D{X: 1, Y: 30, Z: 0, W: 0}:  "#",
	Coord4D{X: 1, Y: 30, Z: 0, W: -1}: "#",
	Coord4D{X: 1, Y: 30, Z: 0, W: 1}:  "#",
	Coord4D{X: 0, Y: 40, Z: 0, W: 0}:  "#",
	Coord4D{X: 1, Y: 40, Z: 0, W: 0}:  "#",
	Coord4D{X: 2, Y: 40, Z: 0, W: 0}:  "#",
	Coord4D{X: 0, Y: 41, Z: 0, W: 0}:  "#",
	Coord4D{X: 1, Y: 40, Z: 1, W: 0}:  "#",
	Coord4D{X: 0, Y: 40, Z: -1, W: 0}: ".",
	Coord4D{X: 0, Y: 50, Z: -1, W: 0}: "#",
	Coord4D{X: 0, Y: 50, Z: -2, W: 0}: ".",
	Coord4D{X: 0, Y: 50, Z: -3, W: 0}: "#",
	Coord4D{X: 0, Y: 60, Z: 0, W: 0}:  "#",
	Coord4D{X: 0, Y: 60, Z: 1, W: 0}:  "#",
	Coord4D{X: 0, Y: 60, Z: 0, W: 1}:  "#",
	Coord4D{X: 0, Y: 60, Z: 0, W: -1}: "#",
	Coord4D{X: 0, Y: 70, Z: 0, W: -1}: ".",
	Coord4D{X: 0, Y: 70, Z: 0, W: 0}:  "#",
	Coord4D{X: 0, Y: 80, Z: 0, W: -1}: "#",
	Coord4D{X: 0, Y: 80, Z: 0, W: 0}:  "#",
	Coord4D{X: 0, Y: 80, Z: 0, W: 1}:  "#",
	Coord4D{X: 0, Y: 80, Z: 1, W: 0}:  "#",
	Coord4D{X: 0, Y: 80, Z: -1, W: 0}: "#",
	Coord4D{X: 1, Y: 80, Z: -1, W: 0}: ".",
}

func TestGrid_evaluateAdjacentCo(t *testing.T) {
	type args struct {
		co               Coord4D
		adjacentCo       Coord4D
		neighboursActive int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "doesn't increment neightboursActive if the co and adjacentCo are the same",
			args: args{
				adjacentCo:       Coord4D{X: 0, Y: 0, Z: 0, W: 0},
				neighboursActive: 67,
			},
			want: 67,
		},
		{
			name: "doesn't increment neightboursActive if the cube at adjacentCo is inactive",
			args: args{
				adjacentCo:       Coord4D{X: 1, Y: 0, Z: 0, W: 0},
				neighboursActive: 67,
			},
			want: 67,
		},
		{
			name: "does increment neightboursActive if the cube at adjacentCo is active",
			args: args{
				adjacentCo:       Coord4D{X: 2, Y: 0, Z: 0, W: 0},
				neighboursActive: 67,
			},
			want: 68,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basicGrid.evaluateAdjacentCo(tt.args.co, tt.args.adjacentCo, tt.args.neighboursActive); got != tt.want {
				t.Errorf("Grid.evaluateAdjacentCo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_evaluateCo(t *testing.T) {
	type args struct {
		is4D bool
		co   Coord4D
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "returns '#' if the given co is '.' and has 3 active neighbours for 3D solution",
			args: args{
				is4D: false,
				co:   Coord4D{X: 1, Y: 11, Z: 0, W: 0},
			},
			want: "#",
		},
		{
			name: "returns '#' if the given co is '#' and has 3 active neighbours for 3D solution",
			args: args{
				is4D: false,
				co:   Coord4D{X: 1, Y: 9, Z: 0, W: 0},
			},
			want: "#",
		},
		{
			name: "returns '#' if the given co is '#' and has 2 active neighbours for 3D solution",
			args: args{
				is4D: false,
				co:   Coord4D{X: 0, Y: 20, Z: -1, W: 0},
			},
			want: "#",
		},
		{
			name: "returns '.' if the given co is '.' and has 2 active neighbours for 3D solution",
			args: args{
				is4D: false,
				co:   Coord4D{X: 0, Y: 20, Z: 1, W: 0},
			},
			want: ".",
		},
		{
			name: "returns '.' if the given co is '#' and has less than 2 active neighbours for 3D solution",
			args: args{
				is4D: false,
				co:   Coord4D{X: 1, Y: 30, Z: 0, W: 0},
			},
			want: ".",
		},
		{
			name: "returns '.' if the given co is '.' and has less than 2 active neighbours for 3D solution",
			args: args{
				is4D: false,
				co:   Coord4D{X: 0, Y: 30, Z: 0, W: 0},
			},
			want: ".",
		},
		{
			name: "returns '.' if the given co is '#' and has more than 3 active neighbours for 3D solution",
			args: args{
				is4D: false,
				co:   Coord4D{X: 1, Y: 40, Z: 1, W: 0},
			},
			want: ".",
		},
		{
			name: "returns '.' if the given co is '.' and has more than 3 active neighbours for 3D solution",
			args: args{
				is4D: false,
				co:   Coord4D{X: 1, Y: 40, Z: -1, W: 0},
			},
			want: ".",
		},
		{
			name: "returns '#' if the given co is '.' and has 3 active neighbours for 4D solution",
			args: args{
				is4D: true,
				co:   Coord4D{X: 0, Y: 20, Z: 1, W: 0},
			},
			want: "#",
		},
		{
			name: "returns '#' if the given co is '#' and has 2 active neighbours for 4D solution",
			args: args{
				is4D: true,
				co:   Coord4D{X: 1, Y: 30, Z: 0, W: 0},
			},
			want: "#",
		},
		{
			name: "returns '.' if the given co is '.' and has 2 active neighbours for 4D solution",
			args: args{
				is4D: true,
				co:   Coord4D{X: 0, Y: 50, Z: -2, W: 0},
			},
			want: ".",
		},
		{
			name: "returns '#' if the given co is '#' and has 3 active neighbours for 4D solution",
			args: args{
				is4D: true,
				co:   Coord4D{X: 0, Y: 60, Z: 0, W: 0},
			},
			want: "#",
		},
		{
			name: "returns '.' if the given co is '#' and has less than 2 active neighbours for 4D solution",
			args: args{
				is4D: true,
				co:   Coord4D{X: 0, Y: 70, Z: 0, W: 0},
			},
			want: ".",
		},
		{
			name: "returns '.' if the given co is '.' and has less than 2 active neighbours for 4D solution",
			args: args{
				is4D: true,
				co:   Coord4D{X: 0, Y: 70, Z: 0, W: -1},
			},
			want: ".",
		},
		{
			name: "returns '.' if the given co is '#' and has more than 3 active neighbours for 4D solution",
			args: args{
				is4D: true,
				co:   Coord4D{X: 0, Y: 80, Z: 0, W: 0},
			},
			want: ".",
		},
		{
			name: "returns '.' if the given co is '.' and has more than 3 active neighbours for 4D solution",
			args: args{
				is4D: true,
				co:   Coord4D{X: 1, Y: 80, Z: -1, W: 0},
			},
			want: ".",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basicGrid.evaluateCo(tt.args.is4D, tt.args.co); got != tt.want {
				t.Errorf("Grid.evaluateCo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_parseInput(t *testing.T) {
	type args struct {
		pocketDimension []string
		iterations      int
	}
	tests := []struct {
		name string
		g    Grid
		args args
		want Grid
	}{
		{
			name: "correctly populates the grid",
			g:    Grid{},
			args: args{
				pocketDimension: []string{"##", ".#"},
				iterations:      1,
			},
			want: Grid{
				Coord4D{X: -1, Y: -1, Z: -1, W: -1}: ".",
				Coord4D{X: -1, Y: -1, Z: -1, W: 0}:  ".",
				Coord4D{X: -1, Y: -1, Z: -1, W: 1}:  ".",
				Coord4D{X: -1, Y: -1, Z: 0, W: -1}:  ".",
				Coord4D{X: -1, Y: -1, Z: 0, W: 0}:   ".",
				Coord4D{X: -1, Y: -1, Z: 0, W: 1}:   ".",
				Coord4D{X: -1, Y: -1, Z: 1, W: -1}:  ".",
				Coord4D{X: -1, Y: -1, Z: 1, W: 0}:   ".",
				Coord4D{X: -1, Y: -1, Z: 1, W: 1}:   ".",
				Coord4D{X: -1, Y: 0, Z: -1, W: -1}:  ".",
				Coord4D{X: -1, Y: 0, Z: -1, W: 0}:   ".",
				Coord4D{X: -1, Y: 0, Z: -1, W: 1}:   ".",
				Coord4D{X: -1, Y: 0, Z: 0, W: -1}:   ".",
				Coord4D{X: -1, Y: 0, Z: 0, W: 0}:    ".",
				Coord4D{X: -1, Y: 0, Z: 0, W: 1}:    ".",
				Coord4D{X: -1, Y: 0, Z: 1, W: -1}:   ".",
				Coord4D{X: -1, Y: 0, Z: 1, W: 0}:    ".",
				Coord4D{X: -1, Y: 0, Z: 1, W: 1}:    ".",
				Coord4D{X: -1, Y: 1, Z: -1, W: -1}:  ".",
				Coord4D{X: -1, Y: 1, Z: -1, W: 0}:   ".",
				Coord4D{X: -1, Y: 1, Z: -1, W: 1}:   ".",
				Coord4D{X: -1, Y: 1, Z: 0, W: -1}:   ".",
				Coord4D{X: -1, Y: 1, Z: 0, W: 0}:    ".",
				Coord4D{X: -1, Y: 1, Z: 0, W: 1}:    ".",
				Coord4D{X: -1, Y: 1, Z: 1, W: -1}:   ".",
				Coord4D{X: -1, Y: 1, Z: 1, W: 0}:    ".",
				Coord4D{X: -1, Y: 1, Z: 1, W: 1}:    ".",
				Coord4D{X: -1, Y: 2, Z: -1, W: -1}:  ".",
				Coord4D{X: -1, Y: 2, Z: -1, W: 0}:   ".",
				Coord4D{X: -1, Y: 2, Z: -1, W: 1}:   ".",
				Coord4D{X: -1, Y: 2, Z: 0, W: -1}:   ".",
				Coord4D{X: -1, Y: 2, Z: 0, W: 0}:    ".",
				Coord4D{X: -1, Y: 2, Z: 0, W: 1}:    ".",
				Coord4D{X: -1, Y: 2, Z: 1, W: -1}:   ".",
				Coord4D{X: -1, Y: 2, Z: 1, W: 0}:    ".",
				Coord4D{X: -1, Y: 2, Z: 1, W: 1}:    ".",
				Coord4D{X: 0, Y: -1, Z: -1, W: -1}:  ".",
				Coord4D{X: 1, Y: -1, Z: -1, W: -1}:  ".",
				Coord4D{X: 2, Y: -1, Z: -1, W: -1}:  ".",
				Coord4D{X: 0, Y: -1, Z: -1, W: 0}:   ".",
				Coord4D{X: 1, Y: -1, Z: -1, W: 0}:   ".",
				Coord4D{X: 2, Y: -1, Z: -1, W: 0}:   ".",
				Coord4D{X: 0, Y: -1, Z: -1, W: 1}:   ".",
				Coord4D{X: 1, Y: -1, Z: -1, W: 1}:   ".",
				Coord4D{X: 2, Y: -1, Z: -1, W: 1}:   ".",
				Coord4D{X: 0, Y: -1, Z: 0, W: -1}:   ".",
				Coord4D{X: 1, Y: -1, Z: 0, W: -1}:   ".",
				Coord4D{X: 2, Y: -1, Z: 0, W: -1}:   ".",
				Coord4D{X: 0, Y: -1, Z: 0, W: 0}:    ".",
				Coord4D{X: 1, Y: -1, Z: 0, W: 0}:    ".",
				Coord4D{X: 2, Y: -1, Z: 0, W: 0}:    ".",
				Coord4D{X: 0, Y: -1, Z: 0, W: 1}:    ".",
				Coord4D{X: 1, Y: -1, Z: 0, W: 1}:    ".",
				Coord4D{X: 2, Y: -1, Z: 0, W: 1}:    ".",
				Coord4D{X: 0, Y: -1, Z: 1, W: -1}:   ".",
				Coord4D{X: 1, Y: -1, Z: 1, W: -1}:   ".",
				Coord4D{X: 2, Y: -1, Z: 1, W: -1}:   ".",
				Coord4D{X: 0, Y: -1, Z: 1, W: 0}:    ".",
				Coord4D{X: 1, Y: -1, Z: 1, W: 0}:    ".",
				Coord4D{X: 2, Y: -1, Z: 1, W: 0}:    ".",
				Coord4D{X: 0, Y: -1, Z: 1, W: 1}:    ".",
				Coord4D{X: 1, Y: -1, Z: 1, W: 1}:    ".",
				Coord4D{X: 2, Y: -1, Z: 1, W: 1}:    ".",
				Coord4D{X: 0, Y: 0, Z: -1, W: -1}:   ".",
				Coord4D{X: 1, Y: 0, Z: -1, W: -1}:   ".",
				Coord4D{X: 2, Y: 0, Z: -1, W: -1}:   ".",
				Coord4D{X: 0, Y: 0, Z: -1, W: 0}:    ".",
				Coord4D{X: 1, Y: 0, Z: -1, W: 0}:    ".",
				Coord4D{X: 2, Y: 0, Z: -1, W: 0}:    ".",
				Coord4D{X: 0, Y: 0, Z: -1, W: 1}:    ".",
				Coord4D{X: 1, Y: 0, Z: -1, W: 1}:    ".",
				Coord4D{X: 2, Y: 0, Z: -1, W: 1}:    ".",
				Coord4D{X: 0, Y: 0, Z: 0, W: -1}:    ".",
				Coord4D{X: 1, Y: 0, Z: 0, W: -1}:    ".",
				Coord4D{X: 2, Y: 0, Z: 0, W: -1}:    ".",
				Coord4D{X: 0, Y: 0, Z: 0, W: 0}:     "#",
				Coord4D{X: 1, Y: 0, Z: 0, W: 0}:     "#",
				Coord4D{X: 2, Y: 0, Z: 0, W: 0}:     ".",
				Coord4D{X: 0, Y: 0, Z: 0, W: 1}:     ".",
				Coord4D{X: 1, Y: 0, Z: 0, W: 1}:     ".",
				Coord4D{X: 2, Y: 0, Z: 0, W: 1}:     ".",
				Coord4D{X: 0, Y: 0, Z: 1, W: -1}:    ".",
				Coord4D{X: 1, Y: 0, Z: 1, W: -1}:    ".",
				Coord4D{X: 2, Y: 0, Z: 1, W: -1}:    ".",
				Coord4D{X: 0, Y: 0, Z: 1, W: 0}:     ".",
				Coord4D{X: 1, Y: 0, Z: 1, W: 0}:     ".",
				Coord4D{X: 2, Y: 0, Z: 1, W: 0}:     ".",
				Coord4D{X: 0, Y: 0, Z: 1, W: 1}:     ".",
				Coord4D{X: 1, Y: 0, Z: 1, W: 1}:     ".",
				Coord4D{X: 2, Y: 0, Z: 1, W: 1}:     ".",
				Coord4D{X: 0, Y: 1, Z: -1, W: -1}:   ".",
				Coord4D{X: 1, Y: 1, Z: -1, W: -1}:   ".",
				Coord4D{X: 2, Y: 1, Z: -1, W: -1}:   ".",
				Coord4D{X: 0, Y: 1, Z: -1, W: 0}:    ".",
				Coord4D{X: 1, Y: 1, Z: -1, W: 0}:    ".",
				Coord4D{X: 2, Y: 1, Z: -1, W: 0}:    ".",
				Coord4D{X: 0, Y: 1, Z: -1, W: 1}:    ".",
				Coord4D{X: 1, Y: 1, Z: -1, W: 1}:    ".",
				Coord4D{X: 2, Y: 1, Z: -1, W: 1}:    ".",
				Coord4D{X: 0, Y: 1, Z: 0, W: -1}:    ".",
				Coord4D{X: 1, Y: 1, Z: 0, W: -1}:    ".",
				Coord4D{X: 2, Y: 1, Z: 0, W: -1}:    ".",
				Coord4D{X: 0, Y: 1, Z: 0, W: 0}:     ".",
				Coord4D{X: 1, Y: 1, Z: 0, W: 0}:     "#",
				Coord4D{X: 2, Y: 1, Z: 0, W: 0}:     ".",
				Coord4D{X: 0, Y: 1, Z: 0, W: 1}:     ".",
				Coord4D{X: 1, Y: 1, Z: 0, W: 1}:     ".",
				Coord4D{X: 2, Y: 1, Z: 0, W: 1}:     ".",
				Coord4D{X: 0, Y: 1, Z: 1, W: -1}:    ".",
				Coord4D{X: 1, Y: 1, Z: 1, W: -1}:    ".",
				Coord4D{X: 2, Y: 1, Z: 1, W: -1}:    ".",
				Coord4D{X: 0, Y: 1, Z: 1, W: 0}:     ".",
				Coord4D{X: 1, Y: 1, Z: 1, W: 0}:     ".",
				Coord4D{X: 2, Y: 1, Z: 1, W: 0}:     ".",
				Coord4D{X: 0, Y: 1, Z: 1, W: 1}:     ".",
				Coord4D{X: 1, Y: 1, Z: 1, W: 1}:     ".",
				Coord4D{X: 2, Y: 1, Z: 1, W: 1}:     ".",
				Coord4D{X: 0, Y: 2, Z: -1, W: -1}:   ".",
				Coord4D{X: 1, Y: 2, Z: -1, W: -1}:   ".",
				Coord4D{X: 2, Y: 2, Z: -1, W: -1}:   ".",
				Coord4D{X: 0, Y: 2, Z: -1, W: 0}:    ".",
				Coord4D{X: 1, Y: 2, Z: -1, W: 0}:    ".",
				Coord4D{X: 2, Y: 2, Z: -1, W: 0}:    ".",
				Coord4D{X: 0, Y: 2, Z: -1, W: 1}:    ".",
				Coord4D{X: 1, Y: 2, Z: -1, W: 1}:    ".",
				Coord4D{X: 2, Y: 2, Z: -1, W: 1}:    ".",
				Coord4D{X: 0, Y: 2, Z: 0, W: -1}:    ".",
				Coord4D{X: 1, Y: 2, Z: 0, W: -1}:    ".",
				Coord4D{X: 2, Y: 2, Z: 0, W: -1}:    ".",
				Coord4D{X: 0, Y: 2, Z: 0, W: 0}:     ".",
				Coord4D{X: 1, Y: 2, Z: 0, W: 0}:     ".",
				Coord4D{X: 2, Y: 2, Z: 0, W: 0}:     ".",
				Coord4D{X: 0, Y: 2, Z: 0, W: 1}:     ".",
				Coord4D{X: 1, Y: 2, Z: 0, W: 1}:     ".",
				Coord4D{X: 2, Y: 2, Z: 0, W: 1}:     ".",
				Coord4D{X: 0, Y: 2, Z: 1, W: -1}:    ".",
				Coord4D{X: 1, Y: 2, Z: 1, W: -1}:    ".",
				Coord4D{X: 2, Y: 2, Z: 1, W: -1}:    ".",
				Coord4D{X: 0, Y: 2, Z: 1, W: 0}:     ".",
				Coord4D{X: 1, Y: 2, Z: 1, W: 0}:     ".",
				Coord4D{X: 2, Y: 2, Z: 1, W: 0}:     ".",
				Coord4D{X: 0, Y: 2, Z: 1, W: 1}:     ".",
				Coord4D{X: 1, Y: 2, Z: 1, W: 1}:     ".",
				Coord4D{X: 2, Y: 2, Z: 1, W: 1}:     ".",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.parseInput(tt.args.pocketDimension, tt.args.iterations)
			if !reflect.DeepEqual(tt.g, tt.want) {
				t.Errorf("Grid.parseInput() = %v, want %v", tt.g, tt.want)
			}
		})
	}
}

func TestGrid_generateNextGrid(t *testing.T) {
	tests := []struct {
		name string
		g    Grid
		is4D bool
		want Grid
	}{
		{
			name: "correctly generates a new 3D grid",
			g: Grid{
				Coord4D{X: 0, Y: 0, Z: -1, W: 0}: ".",
				Coord4D{X: 1, Y: 0, Z: -1, W: 0}: ".",
				Coord4D{X: 2, Y: 0, Z: -1, W: 0}: ".",
				Coord4D{X: 0, Y: 1, Z: -1, W: 0}: ".",
				Coord4D{X: 1, Y: 1, Z: -1, W: 0}: ".",
				Coord4D{X: 2, Y: 1, Z: -1, W: 0}: ".",
				Coord4D{X: 0, Y: 2, Z: -1, W: 0}: ".",
				Coord4D{X: 1, Y: 2, Z: -1, W: 0}: ".",
				Coord4D{X: 2, Y: 2, Z: -1, W: 0}: ".",
				Coord4D{X: 0, Y: 0, Z: 0, W: 0}:  ".",
				Coord4D{X: 1, Y: 0, Z: 0, W: 0}:  ".",
				Coord4D{X: 2, Y: 0, Z: 0, W: 0}:  ".",
				Coord4D{X: 0, Y: 1, Z: 0, W: 0}:  "#",
				Coord4D{X: 1, Y: 1, Z: 0, W: 0}:  "#",
				Coord4D{X: 2, Y: 1, Z: 0, W: 0}:  "#",
				Coord4D{X: 0, Y: 2, Z: 0, W: 0}:  ".",
				Coord4D{X: 1, Y: 2, Z: 0, W: 0}:  ".",
				Coord4D{X: 2, Y: 2, Z: 0, W: 0}:  ".",
				Coord4D{X: 0, Y: 0, Z: 1, W: 0}:  ".",
				Coord4D{X: 1, Y: 0, Z: 1, W: 0}:  ".",
				Coord4D{X: 2, Y: 0, Z: 1, W: 0}:  ".",
				Coord4D{X: 0, Y: 1, Z: 1, W: 0}:  ".",
				Coord4D{X: 1, Y: 1, Z: 1, W: 0}:  ".",
				Coord4D{X: 2, Y: 1, Z: 1, W: 0}:  ".",
				Coord4D{X: 0, Y: 2, Z: 1, W: 0}:  ".",
				Coord4D{X: 1, Y: 2, Z: 1, W: 0}:  ".",
				Coord4D{X: 2, Y: 2, Z: 1, W: 0}:  ".",
			},
			is4D: false,
			want: Grid{
				Coord4D{X: 0, Y: 0, Z: -1, W: 0}: ".",
				Coord4D{X: 1, Y: 0, Z: -1, W: 0}: "#",
				Coord4D{X: 2, Y: 0, Z: -1, W: 0}: ".",
				Coord4D{X: 0, Y: 1, Z: -1, W: 0}: ".",
				Coord4D{X: 1, Y: 1, Z: -1, W: 0}: "#",
				Coord4D{X: 2, Y: 1, Z: -1, W: 0}: ".",
				Coord4D{X: 0, Y: 2, Z: -1, W: 0}: ".",
				Coord4D{X: 1, Y: 2, Z: -1, W: 0}: "#",
				Coord4D{X: 2, Y: 2, Z: -1, W: 0}: ".",
				Coord4D{X: 0, Y: 0, Z: 0, W: 0}:  ".",
				Coord4D{X: 1, Y: 0, Z: 0, W: 0}:  "#",
				Coord4D{X: 2, Y: 0, Z: 0, W: 0}:  ".",
				Coord4D{X: 0, Y: 1, Z: 0, W: 0}:  ".",
				Coord4D{X: 1, Y: 1, Z: 0, W: 0}:  "#",
				Coord4D{X: 2, Y: 1, Z: 0, W: 0}:  ".",
				Coord4D{X: 0, Y: 2, Z: 0, W: 0}:  ".",
				Coord4D{X: 1, Y: 2, Z: 0, W: 0}:  "#",
				Coord4D{X: 2, Y: 2, Z: 0, W: 0}:  ".",
				Coord4D{X: 0, Y: 0, Z: 1, W: 0}:  ".",
				Coord4D{X: 1, Y: 0, Z: 1, W: 0}:  "#",
				Coord4D{X: 2, Y: 0, Z: 1, W: 0}:  ".",
				Coord4D{X: 0, Y: 1, Z: 1, W: 0}:  ".",
				Coord4D{X: 1, Y: 1, Z: 1, W: 0}:  "#",
				Coord4D{X: 2, Y: 1, Z: 1, W: 0}:  ".",
				Coord4D{X: 0, Y: 2, Z: 1, W: 0}:  ".",
				Coord4D{X: 1, Y: 2, Z: 1, W: 0}:  "#",
				Coord4D{X: 2, Y: 2, Z: 1, W: 0}:  ".",
			},
		},
		{
			name: "correctly generates a new 4D grid",
			g: Grid{
				Coord4D{X: 0, Y: 0, Z: -1, W: -1}: ".",
				Coord4D{X: 1, Y: 0, Z: -1, W: -1}: ".",
				Coord4D{X: 2, Y: 0, Z: -1, W: -1}: ".",
				Coord4D{X: 0, Y: 1, Z: -1, W: -1}: ".",
				Coord4D{X: 1, Y: 1, Z: -1, W: -1}: ".",
				Coord4D{X: 2, Y: 1, Z: -1, W: -1}: ".",
				Coord4D{X: 0, Y: 2, Z: -1, W: -1}: ".",
				Coord4D{X: 1, Y: 2, Z: -1, W: -1}: ".",
				Coord4D{X: 2, Y: 2, Z: -1, W: -1}: ".",
				Coord4D{X: 0, Y: 0, Z: 0, W: -1}:  ".",
				Coord4D{X: 1, Y: 0, Z: 0, W: -1}:  ".",
				Coord4D{X: 2, Y: 0, Z: 0, W: -1}:  ".",
				Coord4D{X: 0, Y: 1, Z: 0, W: -1}:  ".",
				Coord4D{X: 1, Y: 1, Z: 0, W: -1}:  ".",
				Coord4D{X: 2, Y: 1, Z: 0, W: -1}:  ".",
				Coord4D{X: 0, Y: 2, Z: 0, W: -1}:  ".",
				Coord4D{X: 1, Y: 2, Z: 0, W: -1}:  ".",
				Coord4D{X: 2, Y: 2, Z: 0, W: -1}:  ".",
				Coord4D{X: 0, Y: 0, Z: 1, W: -1}:  ".",
				Coord4D{X: 1, Y: 0, Z: 1, W: -1}:  ".",
				Coord4D{X: 2, Y: 0, Z: 1, W: -1}:  ".",
				Coord4D{X: 0, Y: 1, Z: 1, W: -1}:  ".",
				Coord4D{X: 1, Y: 1, Z: 1, W: -1}:  ".",
				Coord4D{X: 2, Y: 1, Z: 1, W: -1}:  ".",
				Coord4D{X: 0, Y: 2, Z: 1, W: -1}:  ".",
				Coord4D{X: 1, Y: 2, Z: 1, W: -1}:  ".",
				Coord4D{X: 2, Y: 2, Z: 1, W: -1}:  ".",
				Coord4D{X: 0, Y: 0, Z: -1, W: 0}:  ".",
				Coord4D{X: 1, Y: 0, Z: -1, W: 0}:  ".",
				Coord4D{X: 2, Y: 0, Z: -1, W: 0}:  ".",
				Coord4D{X: 0, Y: 1, Z: -1, W: 0}:  ".",
				Coord4D{X: 1, Y: 1, Z: -1, W: 0}:  ".",
				Coord4D{X: 2, Y: 1, Z: -1, W: 0}:  ".",
				Coord4D{X: 0, Y: 2, Z: -1, W: 0}:  ".",
				Coord4D{X: 1, Y: 2, Z: -1, W: 0}:  ".",
				Coord4D{X: 2, Y: 2, Z: -1, W: 0}:  ".",
				Coord4D{X: 0, Y: 0, Z: 0, W: 0}:   ".",
				Coord4D{X: 1, Y: 0, Z: 0, W: 0}:   ".",
				Coord4D{X: 2, Y: 0, Z: 0, W: 0}:   ".",
				Coord4D{X: 0, Y: 1, Z: 0, W: 0}:   "#",
				Coord4D{X: 1, Y: 1, Z: 0, W: 0}:   "#",
				Coord4D{X: 2, Y: 1, Z: 0, W: 0}:   "#",
				Coord4D{X: 0, Y: 2, Z: 0, W: 0}:   ".",
				Coord4D{X: 1, Y: 2, Z: 0, W: 0}:   ".",
				Coord4D{X: 2, Y: 2, Z: 0, W: 0}:   ".",
				Coord4D{X: 0, Y: 0, Z: 1, W: 0}:   ".",
				Coord4D{X: 1, Y: 0, Z: 1, W: 0}:   ".",
				Coord4D{X: 2, Y: 0, Z: 1, W: 0}:   ".",
				Coord4D{X: 0, Y: 1, Z: 1, W: 0}:   ".",
				Coord4D{X: 1, Y: 1, Z: 1, W: 0}:   ".",
				Coord4D{X: 2, Y: 1, Z: 1, W: 0}:   ".",
				Coord4D{X: 0, Y: 2, Z: 1, W: 0}:   ".",
				Coord4D{X: 1, Y: 2, Z: 1, W: 0}:   ".",
				Coord4D{X: 2, Y: 2, Z: 1, W: 0}:   ".",
				Coord4D{X: 0, Y: 0, Z: -1, W: 1}:  ".",
				Coord4D{X: 1, Y: 0, Z: -1, W: 1}:  ".",
				Coord4D{X: 2, Y: 0, Z: -1, W: 1}:  ".",
				Coord4D{X: 0, Y: 1, Z: -1, W: 1}:  ".",
				Coord4D{X: 1, Y: 1, Z: -1, W: 1}:  ".",
				Coord4D{X: 2, Y: 1, Z: -1, W: 1}:  ".",
				Coord4D{X: 0, Y: 2, Z: -1, W: 1}:  ".",
				Coord4D{X: 1, Y: 2, Z: -1, W: 1}:  ".",
				Coord4D{X: 2, Y: 2, Z: -1, W: 1}:  ".",
				Coord4D{X: 0, Y: 0, Z: 0, W: 1}:   ".",
				Coord4D{X: 1, Y: 0, Z: 0, W: 1}:   ".",
				Coord4D{X: 2, Y: 0, Z: 0, W: 1}:   ".",
				Coord4D{X: 0, Y: 1, Z: 0, W: 1}:   ".",
				Coord4D{X: 1, Y: 1, Z: 0, W: 1}:   ".",
				Coord4D{X: 2, Y: 1, Z: 0, W: 1}:   ".",
				Coord4D{X: 0, Y: 2, Z: 0, W: 1}:   ".",
				Coord4D{X: 1, Y: 2, Z: 0, W: 1}:   ".",
				Coord4D{X: 2, Y: 2, Z: 0, W: 1}:   ".",
				Coord4D{X: 0, Y: 0, Z: 1, W: 1}:   ".",
				Coord4D{X: 1, Y: 0, Z: 1, W: 1}:   ".",
				Coord4D{X: 2, Y: 0, Z: 1, W: 1}:   ".",
				Coord4D{X: 0, Y: 1, Z: 1, W: 1}:   ".",
				Coord4D{X: 1, Y: 1, Z: 1, W: 1}:   ".",
				Coord4D{X: 2, Y: 1, Z: 1, W: 1}:   ".",
				Coord4D{X: 0, Y: 2, Z: 1, W: 1}:   ".",
				Coord4D{X: 1, Y: 2, Z: 1, W: 1}:   ".",
				Coord4D{X: 2, Y: 2, Z: 1, W: 1}:   ".",
			},
			is4D: true,
			want: Grid{
				Coord4D{X: 0, Y: 0, Z: -1, W: -1}: ".",
				Coord4D{X: 1, Y: 0, Z: -1, W: -1}: "#",
				Coord4D{X: 2, Y: 0, Z: -1, W: -1}: ".",
				Coord4D{X: 0, Y: 1, Z: -1, W: -1}: ".",
				Coord4D{X: 1, Y: 1, Z: -1, W: -1}: "#",
				Coord4D{X: 2, Y: 1, Z: -1, W: -1}: ".",
				Coord4D{X: 0, Y: 2, Z: -1, W: -1}: ".",
				Coord4D{X: 1, Y: 2, Z: -1, W: -1}: "#",
				Coord4D{X: 2, Y: 2, Z: -1, W: -1}: ".",
				Coord4D{X: 0, Y: 0, Z: 0, W: -1}:  ".",
				Coord4D{X: 1, Y: 0, Z: 0, W: -1}:  "#",
				Coord4D{X: 2, Y: 0, Z: 0, W: -1}:  ".",
				Coord4D{X: 0, Y: 1, Z: 0, W: -1}:  ".",
				Coord4D{X: 1, Y: 1, Z: 0, W: -1}:  "#",
				Coord4D{X: 2, Y: 1, Z: 0, W: -1}:  ".",
				Coord4D{X: 0, Y: 2, Z: 0, W: -1}:  ".",
				Coord4D{X: 1, Y: 2, Z: 0, W: -1}:  "#",
				Coord4D{X: 2, Y: 2, Z: 0, W: -1}:  ".",
				Coord4D{X: 0, Y: 0, Z: 1, W: -1}:  ".",
				Coord4D{X: 1, Y: 0, Z: 1, W: -1}:  "#",
				Coord4D{X: 2, Y: 0, Z: 1, W: -1}:  ".",
				Coord4D{X: 0, Y: 1, Z: 1, W: -1}:  ".",
				Coord4D{X: 1, Y: 1, Z: 1, W: -1}:  "#",
				Coord4D{X: 2, Y: 1, Z: 1, W: -1}:  ".",
				Coord4D{X: 0, Y: 2, Z: 1, W: -1}:  ".",
				Coord4D{X: 1, Y: 2, Z: 1, W: -1}:  "#",
				Coord4D{X: 2, Y: 2, Z: 1, W: -1}:  ".",
				Coord4D{X: 0, Y: 0, Z: -1, W: 0}:  ".",
				Coord4D{X: 1, Y: 0, Z: -1, W: 0}:  "#",
				Coord4D{X: 2, Y: 0, Z: -1, W: 0}:  ".",
				Coord4D{X: 0, Y: 1, Z: -1, W: 0}:  ".",
				Coord4D{X: 1, Y: 1, Z: -1, W: 0}:  "#",
				Coord4D{X: 2, Y: 1, Z: -1, W: 0}:  ".",
				Coord4D{X: 0, Y: 2, Z: -1, W: 0}:  ".",
				Coord4D{X: 1, Y: 2, Z: -1, W: 0}:  "#",
				Coord4D{X: 2, Y: 2, Z: -1, W: 0}:  ".",
				Coord4D{X: 0, Y: 0, Z: 0, W: 0}:   ".",
				Coord4D{X: 1, Y: 0, Z: 0, W: 0}:   "#",
				Coord4D{X: 2, Y: 0, Z: 0, W: 0}:   ".",
				Coord4D{X: 0, Y: 1, Z: 0, W: 0}:   ".",
				Coord4D{X: 1, Y: 1, Z: 0, W: 0}:   "#",
				Coord4D{X: 2, Y: 1, Z: 0, W: 0}:   ".",
				Coord4D{X: 0, Y: 2, Z: 0, W: 0}:   ".",
				Coord4D{X: 1, Y: 2, Z: 0, W: 0}:   "#",
				Coord4D{X: 2, Y: 2, Z: 0, W: 0}:   ".",
				Coord4D{X: 0, Y: 0, Z: 1, W: 0}:   ".",
				Coord4D{X: 1, Y: 0, Z: 1, W: 0}:   "#",
				Coord4D{X: 2, Y: 0, Z: 1, W: 0}:   ".",
				Coord4D{X: 0, Y: 1, Z: 1, W: 0}:   ".",
				Coord4D{X: 1, Y: 1, Z: 1, W: 0}:   "#",
				Coord4D{X: 2, Y: 1, Z: 1, W: 0}:   ".",
				Coord4D{X: 0, Y: 2, Z: 1, W: 0}:   ".",
				Coord4D{X: 1, Y: 2, Z: 1, W: 0}:   "#",
				Coord4D{X: 2, Y: 2, Z: 1, W: 0}:   ".",
				Coord4D{X: 0, Y: 0, Z: -1, W: 1}:  ".",
				Coord4D{X: 1, Y: 0, Z: -1, W: 1}:  "#",
				Coord4D{X: 2, Y: 0, Z: -1, W: 1}:  ".",
				Coord4D{X: 0, Y: 1, Z: -1, W: 1}:  ".",
				Coord4D{X: 1, Y: 1, Z: -1, W: 1}:  "#",
				Coord4D{X: 2, Y: 1, Z: -1, W: 1}:  ".",
				Coord4D{X: 0, Y: 2, Z: -1, W: 1}:  ".",
				Coord4D{X: 1, Y: 2, Z: -1, W: 1}:  "#",
				Coord4D{X: 2, Y: 2, Z: -1, W: 1}:  ".",
				Coord4D{X: 0, Y: 0, Z: 0, W: 1}:   ".",
				Coord4D{X: 1, Y: 0, Z: 0, W: 1}:   "#",
				Coord4D{X: 2, Y: 0, Z: 0, W: 1}:   ".",
				Coord4D{X: 0, Y: 1, Z: 0, W: 1}:   ".",
				Coord4D{X: 1, Y: 1, Z: 0, W: 1}:   "#",
				Coord4D{X: 2, Y: 1, Z: 0, W: 1}:   ".",
				Coord4D{X: 0, Y: 2, Z: 0, W: 1}:   ".",
				Coord4D{X: 1, Y: 2, Z: 0, W: 1}:   "#",
				Coord4D{X: 2, Y: 2, Z: 0, W: 1}:   ".",
				Coord4D{X: 0, Y: 0, Z: 1, W: 1}:   ".",
				Coord4D{X: 1, Y: 0, Z: 1, W: 1}:   "#",
				Coord4D{X: 2, Y: 0, Z: 1, W: 1}:   ".",
				Coord4D{X: 0, Y: 1, Z: 1, W: 1}:   ".",
				Coord4D{X: 1, Y: 1, Z: 1, W: 1}:   "#",
				Coord4D{X: 2, Y: 1, Z: 1, W: 1}:   ".",
				Coord4D{X: 0, Y: 2, Z: 1, W: 1}:   ".",
				Coord4D{X: 1, Y: 2, Z: 1, W: 1}:   "#",
				Coord4D{X: 2, Y: 2, Z: 1, W: 1}:   ".",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.generateNextGrid(tt.is4D); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.generateNextGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_countActiveCubes(t *testing.T) {
	tests := []struct {
		name string
		g    Grid
		want int
	}{
		{
			name: "returns the correct number of active cubes",
			g: Grid{
				Coord4D{X: 0, Y: 0, Z: 0, W: 0}:   "#",
				Coord4D{X: 1, Y: 0, Z: 0, W: 0}:   "#",
				Coord4D{X: 2, Y: 0, Z: 0, W: 0}:   "#",
				Coord4D{X: 0, Y: 1, Z: 0, W: 0}:   "#",
				Coord4D{X: 0, Y: 2, Z: 0, W: 0}:   ".",
				Coord4D{X: 0, Y: 0, Z: 1, W: 0}:   ".",
				Coord4D{X: 0, Y: 0, Z: 2, W: 0}:   "#",
				Coord4D{X: 0, Y: 0, Z: 0, W: 1}:   ".",
				Coord4D{X: 0, Y: 0, Z: 0, W: 2}:   "#",
				Coord4D{X: -1, Y: 0, Z: 0, W: 2}:  ".",
				Coord4D{X: -1, Y: 0, Z: 10, W: 2}: ".",
				Coord4D{X: 0, Y: 0, Z: 20, W: 2}:  "#",
				Coord4D{X: 0, Y: 0, Z: 0, W: 25}:  "#",
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.countActiveCubes(); got != tt.want {
				t.Errorf("Grid.countActiveCubes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_findSolution(t *testing.T) {
	type args struct {
		is4D       bool
		iterations int
	}
	tests := []struct {
		name string
		g    Grid
		args args
		want int
	}{
		{
			name: "advent of code example 1",
			g:    Grid{},
			args: args{
				is4D:       false,
				iterations: 6,
			},
			want: 112,
		},
		{
			name: "advent of code example 2",
			g:    Grid{},
			args: args{
				is4D:       true,
				iterations: 6,
			},
			want: 848,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.parseInput([]string{".#.", "..#", "###"}, 6)
			if got := tt.g.findSolution(tt.args.is4D, tt.args.iterations); got != tt.want {
				t.Errorf("Grid.findSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
