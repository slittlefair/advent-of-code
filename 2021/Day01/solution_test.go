package main

import (
	"reflect"
	"testing"
)

func Test_calculateIncreases(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name: "returns the correct number of increases in the input, advent of code example 1",
			input: []int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			want: 7,
		},
		{
			name: "returns the correct number of increases in the input, advent of code example 2",
			input: []int{
				607,
				618,
				618,
				617,
				647,
				716,
				769,
				792,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateIncreases(tt.input); got != tt.want {
				t.Errorf("calculateIncreases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateSlidingWindows(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name: "returns a slice of 3 value inputs from given slice",
			input: []int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			want: []int{
				607,
				618,
				618,
				617,
				647,
				716,
				769,
				792,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSlidingWindows(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateSlidingWindows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name: "returns number of increases in a given input, advent of code example",
			input: []int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name: "returns number of increases in sliding windows of a given input",
			input: []int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
