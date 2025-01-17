package main

import (
	"Advent-of-Code/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInput(t *testing.T) {
	t.Run("returns an error if any line doesn't have the right data", func(t *testing.T) {
		input := []string{
			"p=0,4 v=3,-3",
			"p=6,3 v=-1,-3",
			"p=10,3 v=-1,2",
			"p=2,0 v=2,-1",
			"p=0 v=1,3",
			"p=3,0 v=-2,-2",
			"p=7,6 v=-1,-3",
			"p=3,0 v=-1,-2",
			"p=9,3 v=2,3",
			"p=7,3 v=-1,2",
			"p=2,4 v=2,-3",
			"p=9,5 v=-3,-3",
		}
		b, err := parseInput(input, 11, 7)
		assert.Nil(t, b)
		assert.ErrorContains(t, err, "malformed input, expected 4 nums, got")
	})

	t.Run("returns a bathroom with robots from the given input", func(t *testing.T) {
		input := []string{
			"p=0,4 v=3,-3",
			"p=6,3 v=-1,-3",
			"p=10,3 v=-1,2",
			"p=2,0 v=2,-1",
			"p=0,0 v=1,3",
			"p=3,0 v=-2,-2",
			"p=7,6 v=-1,-3",
			"p=3,0 v=-1,-2",
			"p=9,3 v=2,3",
			"p=7,3 v=-1,2",
			"p=2,4 v=2,-3",
			"p=9,5 v=-3,-3",
		}

		want := &Bathroom{
			maxX: 11,
			maxY: 7,
			robots: []*Robot{
				{position: graph.Co{X: 0, Y: 4}, velocity: graph.Co{X: 3, Y: -3}},
				{position: graph.Co{X: 6, Y: 3}, velocity: graph.Co{X: -1, Y: -3}},
				{position: graph.Co{X: 10, Y: 3}, velocity: graph.Co{X: -1, Y: 2}},
				{position: graph.Co{X: 2, Y: 0}, velocity: graph.Co{X: 2, Y: -1}},
				{position: graph.Co{X: 0, Y: 0}, velocity: graph.Co{X: 1, Y: 3}},
				{position: graph.Co{X: 3, Y: 0}, velocity: graph.Co{X: -2, Y: -2}},
				{position: graph.Co{X: 7, Y: 6}, velocity: graph.Co{X: -1, Y: -3}},
				{position: graph.Co{X: 3, Y: 0}, velocity: graph.Co{X: -1, Y: -2}},
				{position: graph.Co{X: 9, Y: 3}, velocity: graph.Co{X: 2, Y: 3}},
				{position: graph.Co{X: 7, Y: 3}, velocity: graph.Co{X: -1, Y: 2}},
				{position: graph.Co{X: 2, Y: 4}, velocity: graph.Co{X: 2, Y: -3}},
				{position: graph.Co{X: 9, Y: 5}, velocity: graph.Co{X: -3, Y: -3}},
			},
		}

		got, err := parseInput(input, 11, 7)
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})
}

func TestBathroom_moveRobots(t *testing.T) {
	t.Run("moves robots according to their poition and velocity", func(t *testing.T) {
		b := &Bathroom{
			maxX: 11,
			maxY: 7,
			robots: []*Robot{
				{position: graph.Co{X: 0, Y: 4}, velocity: graph.Co{X: 3, Y: -3}},
				{position: graph.Co{X: 6, Y: 3}, velocity: graph.Co{X: -1, Y: -3}},
				{position: graph.Co{X: 10, Y: 3}, velocity: graph.Co{X: -1, Y: 2}},
				{position: graph.Co{X: 2, Y: 0}, velocity: graph.Co{X: 2, Y: -1}},
				{position: graph.Co{X: 0, Y: 0}, velocity: graph.Co{X: 1, Y: 3}},
				{position: graph.Co{X: 3, Y: 0}, velocity: graph.Co{X: -2, Y: -2}},
				{position: graph.Co{X: 7, Y: 6}, velocity: graph.Co{X: -1, Y: -3}},
				{position: graph.Co{X: 3, Y: 0}, velocity: graph.Co{X: -1, Y: -2}},
				{position: graph.Co{X: 9, Y: 3}, velocity: graph.Co{X: 2, Y: 3}},
				{position: graph.Co{X: 7, Y: 3}, velocity: graph.Co{X: -1, Y: 2}},
				{position: graph.Co{X: 2, Y: 4}, velocity: graph.Co{X: 2, Y: -3}},
				{position: graph.Co{X: 9, Y: 5}, velocity: graph.Co{X: -3, Y: -3}},
			},
		}
		b.moveRobots()
		want := []*Robot{
			{position: graph.Co{X: 3, Y: 1}, velocity: graph.Co{X: 3, Y: -3}},
			{position: graph.Co{X: 5, Y: 0}, velocity: graph.Co{X: -1, Y: -3}},
			{position: graph.Co{X: 9, Y: 5}, velocity: graph.Co{X: -1, Y: 2}},
			{position: graph.Co{X: 4, Y: 6}, velocity: graph.Co{X: 2, Y: -1}},
			{position: graph.Co{X: 1, Y: 3}, velocity: graph.Co{X: 1, Y: 3}},
			{position: graph.Co{X: 1, Y: 5}, velocity: graph.Co{X: -2, Y: -2}},
			{position: graph.Co{X: 6, Y: 3}, velocity: graph.Co{X: -1, Y: -3}},
			{position: graph.Co{X: 2, Y: 5}, velocity: graph.Co{X: -1, Y: -2}},
			{position: graph.Co{X: 0, Y: 6}, velocity: graph.Co{X: 2, Y: 3}},
			{position: graph.Co{X: 6, Y: 5}, velocity: graph.Co{X: -1, Y: 2}},
			{position: graph.Co{X: 4, Y: 1}, velocity: graph.Co{X: 2, Y: -3}},
			{position: graph.Co{X: 6, Y: 2}, velocity: graph.Co{X: -3, Y: -3}},
		}
		assert.Equal(t, want, b.robots)
	})
}

func TestBathroom_findSafetyFactor(t *testing.T) {
	b := Bathroom{
		maxX: 11,
		maxY: 7,
		robots: []*Robot{
			{position: graph.Co{X: 6, Y: 0}},
			{position: graph.Co{X: 6, Y: 0}},
			{position: graph.Co{X: 9, Y: 0}},
			{position: graph.Co{X: 0, Y: 2}},
			{position: graph.Co{X: 1, Y: 3}},
			{position: graph.Co{X: 2, Y: 3}},
			{position: graph.Co{X: 5, Y: 4}},
			{position: graph.Co{X: 3, Y: 6}},
			{position: graph.Co{X: 4, Y: 6}},
			{position: graph.Co{X: 4, Y: 6}},
			{position: graph.Co{X: 1, Y: 7}},
			{position: graph.Co{X: 6, Y: 6}},
		},
	}
	got := b.findSafetyFactor()
	assert.Equal(t, 12, got)
}
