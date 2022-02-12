package main

import (
	"Advent-of-Code/graph"
	djk "Advent-of-Code/graph/dijkstra"
	"reflect"
	"testing"
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
			name: "returns solutions for part 1 and part 2, advent of code example",
			input: []string{
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
			want:    40,
			want1:   315,
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
