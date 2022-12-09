package main

import (
	"Advent-of-Code/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

var aocTreeMap = TreeMap{
	maxX: 4,
	maxY: 4,
	graph: map[graph.Co]int{
		{X: 0, Y: 0}: 3,
		{X: 1, Y: 0}: 0,
		{X: 2, Y: 0}: 3,
		{X: 3, Y: 0}: 7,
		{X: 4, Y: 0}: 3,
		{X: 0, Y: 1}: 2,
		{X: 1, Y: 1}: 5,
		{X: 2, Y: 1}: 5,
		{X: 3, Y: 1}: 1,
		{X: 4, Y: 1}: 2,
		{X: 0, Y: 2}: 6,
		{X: 1, Y: 2}: 5,
		{X: 2, Y: 2}: 3,
		{X: 3, Y: 2}: 3,
		{X: 4, Y: 2}: 2,
		{X: 0, Y: 3}: 3,
		{X: 1, Y: 3}: 3,
		{X: 2, Y: 3}: 5,
		{X: 3, Y: 3}: 4,
		{X: 4, Y: 3}: 9,
		{X: 0, Y: 4}: 3,
		{X: 1, Y: 4}: 5,
		{X: 2, Y: 4}: 3,
		{X: 3, Y: 4}: 9,
		{X: 4, Y: 4}: 0,
	},
}

func TestParseInput(t *testing.T) {
	t.Run("parses input into a TreeMap, advent of code example", func(t *testing.T) {
		input := []string{
			"30373",
			"25512",
			"65332",
			"33549",
			"35390",
		}
		got := parseInput(input)
		assert.Equal(t, aocTreeMap, got)
	})
}

func TestTravel(t *testing.T) {
	tests := []struct {
		name      string
		from      int
		to        int
		change    int
		direction Direction
		want      bool
		want1     int
	}{
		{
			name:      "returns visible status and viewing distance from {X: 1, Y: 2} moving left",
			from:      0,
			to:        -1,
			change:    -1,
			direction: Horizontally,
			want:      false,
			want1:     1,
		},
		{
			name:      "returns visible status and viewing distance from {X: 1, Y: 2} moving right",
			from:      2,
			to:        5,
			change:    1,
			direction: Horizontally,
			want:      true,
			want1:     3,
		},
		{
			name:      "returns visible status and viewing distance from {X: 1, Y: 2} moving up",
			from:      1,
			to:        -1,
			change:    -1,
			direction: Vertically,
			want:      false,
			want1:     1,
		},
		{
			name:      "returns visible status and viewing distance from {X: 1, Y: 2} moving down",
			from:      3,
			to:        5,
			change:    1,
			direction: Vertically,
			want:      false,
			want1:     2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := aocTreeMap.travel(graph.Co{X: 1, Y: 2}, tt.from, tt.to, tt.change, tt.direction)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestOptimiseTreehouseLocation(t *testing.T) {
	t.Run("finds visible tree count and highest scenic score, advent of code example", func(t *testing.T) {
		got, got1 := aocTreeMap.optimiseTreehouseLocation()
		assert.Equal(t, 21, got)
		assert.Equal(t, 8, got1)
	})
}
