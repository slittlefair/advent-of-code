package main

import (
	"Advent-of-Code/graph"
	djk "Advent-of-Code/graph/dijkstra"
	"fmt"
	"reflect"
	"testing"
)

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    int
		want1   graph.Co
		wantErr bool
	}{
		{
			name: "returns an error if first line of input is not an int",
			input: []string{
				"abc",
				"7,4",
			},
			want:    -1,
			want1:   graph.Co{},
			wantErr: true,
		},
		{
			name: "returns number and target co, advent of code example",
			input: []string{
				"10",
				"7,4",
			},
			want:    10,
			want1:   graph.Co{X: 7, Y: 4},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseInput(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseInput() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseInput() got1 = %v, want %v", got1, tt.want1)
			}
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
			if got := isSpace(tt.co, 10); got != tt.want {
				t.Errorf("isSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_populateOffice(t *testing.T) {
	type args struct {
		target graph.Co
		num    int
	}
	tests := []struct {
		name string
		args args
		want *djk.Graph
	}{
		{
			name: "returns graph from given input, advent of code example",
			args: args{
				target: graph.Co{X: -8, Y: -8},
				num:    10,
			},
			want: &djk.Graph{
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := populateOffice(tt.args.target, tt.args.num)
			if !reflect.DeepEqual(got.Grid, tt.want.Grid) {
				t.Errorf("parseInput() = %v, want %v", got.Grid, tt.want.Grid)
			}
			for co, edge := range got.Nodes {
				if len(edge) != len(tt.want.Nodes[co]) {
					t.Errorf("parseInput().Nodes[%v] = %v, want %v", co, edge, tt.want.Nodes[co])
				}
				edges := map[djk.Edge]struct{}{}
				for _, e := range edge {
					edges[e] = struct{}{}
				}
				for _, e := range tt.want.Nodes[co] {
					if _, ok := edges[e]; !ok {
						t.Errorf("parseInput().Nodes[%v] = %v, want %v", co, edge, tt.want.Nodes[co])
					}
				}
			}
			if got.MaxX != tt.want.MaxX {
				t.Errorf("parseInput().maxX = %d, want %d", got.MaxX, tt.want.MaxX)
			}
			if got.MaxY != tt.want.MaxY {
				t.Errorf("parseInput().MaxY = %d, want %d", got.MaxY, tt.want.MaxY)
			}
		})
	}
}

func Test_findSolutions(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    int
		want1   int
		wantErr bool
	}{
		{
			name: "returns an error if parse input returns an error",
			input: []string{
				"a",
				"2,3",
			},
			want:    -1,
			want1:   -1,
			wantErr: true,
		},
		{
			name: "returns an error if a path to target can't be found",
			input: []string{
				"10",
				"5,2",
			},
			want:    -1,
			want1:   -1,
			wantErr: true,
		},
		{
			name: "returns correct answers for part1 and part2, advent of code example",
			input: []string{
				"10",
				"7,4",
			},
			want:    11,
			want1:   65,
			wantErr: false,
		},
		{
			name: "returns correct answers for part1 and part2, real solution (since AoC doesn't provide example part2)",
			input: []string{
				"1350",
				"31,39",
			},
			want:    92,
			want1:   124,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := findSolutions(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("findSolutions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findSolutions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findSolutions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
