package graph

import (
	"testing"
)

func TestAdjacentCos(t *testing.T) {
	type args struct {
		co               Co
		includeDiagonals bool
	}
	tests := []struct {
		name string
		args args
		want map[Co]struct{}
	}{
		{
			name: "returns adjacent coordinates, including diagonals, of a given coordinate",
			args: args{
				co:               Co{X: 4, Y: 7},
				includeDiagonals: true,
			},
			want: map[Co]struct{}{
				{X: 3, Y: 6}: {},
				{X: 4, Y: 6}: {},
				{X: 5, Y: 6}: {},
				{X: 3, Y: 7}: {},
				{X: 5, Y: 7}: {},
				{X: 3, Y: 8}: {},
				{X: 4, Y: 8}: {},
				{X: 5, Y: 8}: {},
			},
		},
		{
			name: "returns adjacent coordinates, including diagonals, of the origin",
			args: args{
				co:               Co{X: 0, Y: 0},
				includeDiagonals: true,
			},
			want: map[Co]struct{}{
				{X: -1, Y: -1}: {},
				{X: 0, Y: -1}:  {},
				{X: 1, Y: -1}:  {},
				{X: -1, Y: 0}:  {},
				{X: 1, Y: 0}:   {},
				{X: -1, Y: 1}:  {},
				{X: 0, Y: 1}:   {},
				{X: 1, Y: 1}:   {},
			},
		},
		{
			name: "returns adjacent coordinates, excluding diagonals, of a given coordinate",
			args: args{
				co:               Co{X: 4, Y: 7},
				includeDiagonals: false,
			},
			want: map[Co]struct{}{
				{X: 4, Y: 6}: {},
				{X: 3, Y: 7}: {},
				{X: 5, Y: 7}: {},
				{X: 4, Y: 8}: {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[Co]struct{}{}
			for _, co := range AdjacentCos(tt.args.co, tt.args.includeDiagonals) {
				got[co] = struct{}{}
			}
			if len(got) != len(tt.want) {
				t.Errorf("AdjacentCos() got = %v, want %v", got, tt.want)
			}
			for co := range got {
				if _, ok := tt.want[co]; !ok {
					t.Errorf("AdjacentCos() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestCalculateManhattanDistance(t *testing.T) {
	type args struct {
		co1 Co
		co2 Co
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "returns distance between a coordinate and origin",
			args: args{
				co1: Co{X: 7, Y: 8},
				co2: Co{},
			},
			want: 15,
		},
		{
			name: "returns distance between a positive and negative coordinate",
			args: args{
				co1: Co{X: 9, Y: 1},
				co2: Co{X: -9, Y: -7},
			},
			want: 26,
		},
		{
			name: "returns distance when where difference between the two will be negative",
			args: args{
				co1: Co{X: 1, Y: 1},
				co2: Co{X: 8, Y: 11},
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateManhattanDistance(tt.args.co1, tt.args.co2); got != tt.want {
				t.Errorf("CalculateManhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
