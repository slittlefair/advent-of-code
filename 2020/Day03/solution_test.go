package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			got := tt.tm.traverseSlopes(tt.args.right, tt.args.down)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTreeMap_part1(t *testing.T) {
	t.Run("advent of code example map", func(t *testing.T) {
		got := exampleMap.part1()
		assert.Equal(t, 7, got)
	})
}

func TestTreeMap_part2(t *testing.T) {
	t.Run("advent of code example map", func(t *testing.T) {
		got := exampleMap.part2()
		assert.Equal(t, 336, got)
	})
}
