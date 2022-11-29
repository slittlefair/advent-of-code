package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sortedAdapters1 = []int{1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, 22}
var sortedAdapters2 = []int{1, 2, 3, 4, 7, 8, 9, 10, 11, 14, 17, 18, 19, 20, 23, 24, 25, 28, 31, 32, 33, 34, 35, 38, 39, 42, 45, 46, 47, 48, 49, 52}

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "advent of code example 1",
			input: []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4},
			want:  sortedAdapters1,
		},
		{
			name:  "advent of code example 2",
			input: []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3},
			want:  sortedAdapters2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseInput(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestJoltages_part1(t *testing.T) {
	tests := []struct {
		name     string
		j        Joltages
		adapters []int
		want     int
	}{
		{
			name:     "advent of code example 1",
			j:        Joltages{1: 0, 2: 0, 3: 0},
			adapters: sortedAdapters1,
			want:     35,
		},
		{
			name:     "advent of code example 2",
			j:        Joltages{1: 0, 2: 0, 3: 0},
			adapters: sortedAdapters2,
			want:     220,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.j.part1(tt.adapters)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_calculatePerms(t *testing.T) {
	type args struct {
		adapters []int
		val      int
		i        int
		cache    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "advent of code example 1",
			args: args{
				adapters: sortedAdapters1,
				val:      0,
				i:        -1,
				cache:    []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			},
			want: 8,
		},
		{
			name: "advent of code example 1",
			args: args{
				adapters: sortedAdapters2,
				val:      0,
				i:        -1,
				cache:    []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			},
			want: 19208,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculatePerms(tt.args.adapters, tt.args.val, tt.args.i, tt.args.cache)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name     string
		adapters []int
		want     int
	}{
		{
			name:     "advent of code example 1",
			adapters: sortedAdapters1,
			want:     8,
		},
		{
			name:     "advent of code example 2",
			adapters: sortedAdapters2,
			want:     19208,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := part2(tt.adapters)
			assert.Equal(t, tt.want, got)
		})
	}
}
