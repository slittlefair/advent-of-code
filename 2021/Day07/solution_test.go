package main

import (
	"testing"
)

func Test_calculateFuelSpend(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "calculates fuel spend, advent of code example 1",
			x:    11,
			want: 66,
		},
		{
			name: "calculates fuel spend, advent of code example 2",
			x:    9,
			want: 45,
		},
		{
			name: "calculates fuel spend, advent of code example 3",
			x:    1,
			want: 1,
		},
		{
			name: "calculates fuel spend, advent of code example 4",
			x:    3,
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateFuelSpend(tt.x); got != tt.want {
				t.Errorf("calculateFuelSpend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMinFuelSpend(t *testing.T) {
	type args struct {
		input []int
		min   int
		max   int
		part2 bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "returns min fuel spend for part 1, advent of code example",
			args: args{
				input: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
				min:   0,
				max:   16,
				part2: false,
			},
			want: 37,
		},
		{
			name: "returns min fuel spend for part 2, advent of code example",
			args: args{
				input: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
				min:   0,
				max:   16,
				part2: true,
			},
			want: 168,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinFuelSpend(tt.args.input, tt.args.min, tt.args.max, tt.args.part2); got != tt.want {
				t.Errorf("getMinDist() = %v, want %v", got, tt.want)
			}
		})
	}
}
