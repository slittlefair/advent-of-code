package main

import (
	"Advent-of-Code/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createVisitedHouses(t *testing.T) {
	t.Run("creates a basic VisitedHouse object", func(t *testing.T) {
		got := createVisitedHouses()
		expected := &VisitedHouses{
			Map: map[graph.Co]bool{
				{X: 0, Y: 0}: true,
			},
		}
		assert.Equal(t, got, expected)
	})
}

func TestVisitedHouses_moveSanta(t *testing.T) {
	tests := []struct {
		name  string
		Santa graph.Co
		dir   string
		want  graph.Co
	}{
		{
			name:  "subtracts 1 from currentHouse.X if moving west",
			Santa: graph.Co{X: 0, Y: 0},
			dir:   "<",
			want:  graph.Co{X: -1, Y: 0},
		},
		{
			name:  "adds 1 to currentHouse.X if moving east",
			Santa: graph.Co{X: 0, Y: 0},
			dir:   ">",
			want:  graph.Co{X: 1, Y: 0},
		},
		{
			name:  "subtracts 1 from currentHouse.Y if moving north",
			Santa: graph.Co{X: 0, Y: 0},
			dir:   "^",
			want:  graph.Co{X: 0, Y: -1},
		},
		{
			name:  "adds 1 from currentHouse.Y if moving south",
			Santa: graph.Co{X: 0, Y: 0},
			dir:   "v",
			want:  graph.Co{X: 0, Y: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vh := &VisitedHouses{
				Santa: tt.Santa,
			}
			vh.moveSanta(tt.dir)
			assert.Equal(t, tt.want, vh.Santa)
		})
	}
}

func TestVisitedHouses_moveRoboSanta(t *testing.T) {
	tests := []struct {
		name      string
		RoboSanta graph.Co
		dir       string
		want      graph.Co
	}{
		{
			name:      "subtracts 1 from currentHouse.X if moving west",
			RoboSanta: graph.Co{X: 0, Y: 0},
			dir:       "<",
			want:      graph.Co{X: -1, Y: 0},
		},
		{
			name:      "adds 1 to currentHouse.X if moving east",
			RoboSanta: graph.Co{X: 0, Y: 0},
			dir:       ">",
			want:      graph.Co{X: 1, Y: 0},
		},
		{
			name:      "subtracts 1 from currentHouse.Y if moving north",
			RoboSanta: graph.Co{X: 0, Y: 0},
			dir:       "^",
			want:      graph.Co{X: 0, Y: -1},
		},
		{
			name:      "adds 1 from currentHouse.Y if moving south",
			RoboSanta: graph.Co{X: 0, Y: 0},
			dir:       "v",
			want:      graph.Co{X: 0, Y: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vh := &VisitedHouses{
				RoboSanta: tt.RoboSanta,
			}
			vh.moveRoboSanta(tt.dir)
			assert.Equal(t, tt.want, vh.RoboSanta)
		})
	}
}

func TestVisitedHouses_alreadyVisitedHouse(t *testing.T) {
	tests := []struct {
		name  string
		Map   map[graph.Co]bool
		santa graph.Co
		want  bool
	}{
		{
			name: "returns true if the current house has been seen",
			Map: map[graph.Co]bool{
				{X: 0, Y: 0}:  true,
				{X: 1, Y: 99}: true,
				{X: -2, Y: 1}: true,
				{X: 1, Y: 1}:  true,
			},
			santa: graph.Co{X: 1, Y: 99},
			want:  true,
		},
		{
			name: "returns false if the current house has not been seen",
			Map: map[graph.Co]bool{
				{X: 0, Y: 0}:  true,
				{X: 1, Y: 99}: true,
				{X: -2, Y: 1}: true,
				{X: 1, Y: 1}:  true,
			},
			santa: graph.Co{X: 2, Y: 1},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vh := &VisitedHouses{
				Map: tt.Map,
			}
			got := vh.alreadyVisitedHouse(tt.santa)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestVisitedHouses_countUniqueHousesVisited(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part1 bool
		want  int
	}{
		{
			name:  "advent of code example 1, part1",
			input: ">",
			part1: true,
			want:  2,
		},
		{
			name:  "advent of code example 2, part1",
			input: "^>v<",
			part1: true,
			want:  4,
		},
		{
			name:  "advent of code example 3, part1",
			input: "^v^v^v^v^v",
			part1: true,
			want:  2,
		},
		{
			name:  "advent of code example 1, part2",
			input: "^v",
			part1: false,
			want:  3,
		},
		{
			name:  "advent of code example 2, part2",
			input: "^>v<",
			part1: false,
			want:  3,
		},
		{
			name:  "advent of code example 3, part2",
			input: "^v^v^v^v^v",
			part1: false,
			want:  11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vh := &VisitedHouses{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}: true,
				},
			}
			got := vh.countUniqueHousesVisited(tt.input, tt.part1)
			assert.Equal(t, tt.want, got)
		})
	}
}
