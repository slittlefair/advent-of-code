package main

import (
	"testing"
)

var exampleMap = TreeMap{
	"..##.......",
	"#...#...#..",
	".#....#..#.",
	"..#.#...#.#",
	".#...##..#.",
	"..#.##.....",
	".#.#.#....#",
	".#........#",
	"#.##...#...",
	"#...##....#",
	".#..#...#.#",
}

func TestTreeMap_traverseSlopes(t *testing.T) {
	type args struct {
		right int
		down  int
	}
	tests := []struct {
		name string
		tm   TreeMap
		args args
		want int
	}{
		{
			name: "advent of code example map, right 1 down 1",
			tm:   exampleMap,
			args: args{
				right: 1,
				down:  1,
			},
			want: 2,
		},
		{
			name: "advent of code example map, right 3 down 1",
			tm:   exampleMap,
			args: args{
				right: 3,
				down:  1,
			},
			want: 7,
		},
		{
			name: "advent of code example map, right 5 down 1",
			tm:   exampleMap,
			args: args{
				right: 5,
				down:  1,
			},
			want: 3,
		},
		{
			name: "advent of code example map, right 7 down 1",
			tm:   exampleMap,
			args: args{
				right: 7,
				down:  1,
			},
			want: 4,
		},
		{
			name: "advent of code example map, right 1 down 2",
			tm:   exampleMap,
			args: args{
				right: 1,
				down:  2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tm.traverseSlopes(tt.args.right, tt.args.down); got != tt.want {
				t.Errorf("TreeMap.traverseSlopes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeMap_part1(t *testing.T) {
	tests := []struct {
		name string
		tm   TreeMap
		want int
	}{
		{
			name: "advent of code example map",
			tm:   exampleMap,
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tm.part1(); got != tt.want {
				t.Errorf("TreeMap.part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeMap_part2(t *testing.T) {
	tests := []struct {
		name string
		tm   TreeMap
		want int
	}{
		{
			name: "advent of code example map",
			tm:   exampleMap,
			want: 336,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tm.part2(); got != tt.want {
				t.Errorf("TreeMap.part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
