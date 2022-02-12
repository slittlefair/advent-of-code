package main

import (
	"reflect"
	"testing"
)

func Test_createLanternfish(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  Lanternfish
	}{
		{
			name:  "returns a parsed Lanternfish srray from input, advent of code example",
			input: []int{3, 4, 3, 1, 2},
			want:  Lanternfish{0, 1, 1, 2, 1, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createLanternfish(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createLanternfish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLanternfish_iterate(t *testing.T) {
	tests := []struct {
		name string
		lf   Lanternfish
		want Lanternfish
	}{
		{
			name: "returns a correct new Lanternfish slice from a previous iteration, advent of code example 1",
			lf:   Lanternfish{0, 1, 1, 2, 1, 0, 0, 0, 0},
			want: Lanternfish{1, 1, 2, 1, 0, 0, 0, 0, 0},
		},
		{
			name: "returns a correct new Lanternfish slice from a previous iteration, advent of code example 2",
			lf:   Lanternfish{1, 1, 2, 1, 0, 0, 0, 0, 0},
			want: Lanternfish{1, 2, 1, 0, 0, 0, 1, 0, 1},
		},
		{
			name: "returns a correct new Lanternfish slice from a previous iteration, advent of code example 3",
			lf:   Lanternfish{2, 1, 0, 0, 0, 1, 1, 1, 0},
			want: Lanternfish{1, 0, 0, 0, 1, 1, 3, 0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lf.iterate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lanternfish.iterate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLanternfish_count(t *testing.T) {
	tests := []struct {
		name string
		lf   Lanternfish
		want int
	}{
		{
			name: "returns sum of Lanternfish",
			lf:   Lanternfish{1, 8, 2, 5, 19, 4, 3, 0, 7},
			want: 49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lf.count(); got != tt.want {
				t.Errorf("Lanternfish.count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLanternfish_findSolution(t *testing.T) {
	tests := []struct {
		name  string
		lf    Lanternfish
		want  int
		want1 int
	}{
		{
			name:  "correctly returns number of lanternfish after 80 and 256 iterations, advent of code example",
			lf:    Lanternfish{0, 1, 1, 2, 1, 0, 0, 0, 0},
			want:  5934,
			want1: 26984457539,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.lf.findSolution()
			if got != tt.want {
				t.Errorf("Lanternfish.findSolution() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Lanternfish.findSolution() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
