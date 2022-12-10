package main

import (
	"Advent-of-Code/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeRope(t *testing.T) {
	t.Run("makes a 2 knotted rope, advent of code example", func(t *testing.T) {
		want := Rope{
			visited: make(map[graph.Co]bool),
			knots: map[int]*Knot{
				0: {},
				1: {parent: &Knot{}},
			},
		}
		got := makeRope(2)
		assert.Equal(t, want, got)
	})

	t.Run("makes a 5 knotted rope", func(t *testing.T) {
		k0 := &Knot{}
		k1 := &Knot{parent: k0}
		k2 := &Knot{parent: k1}
		k3 := &Knot{parent: k2}
		k4 := &Knot{parent: k3}
		want := Rope{
			visited: make(map[graph.Co]bool),
			knots: map[int]*Knot{
				0: k0,
				1: k1,
				2: k2,
				3: k3,
				4: k4,
			},
		}
		got := makeRope(5)
		assert.Equal(t, want, got)
	})
}

func TestMoveRope(t *testing.T) {
	t.Run("returns an error if a line doesn't scan correctly", func(t *testing.T) {
		r := &Rope{}
		err := r.moveRope("U3")
		assert.Error(t, err)
	})

	t.Run("returns an error if a line doesn't provide one of the correct directions", func(t *testing.T) {
		r := &Rope{}
		err := r.moveRope("F 6")
		assert.Error(t, err)
	})

	t.Run("moves a 2 knot rope the given direction and distance, non diagonally, advent of code example 1", func(t *testing.T) {
		k0 := &Knot{}
		k1 := &Knot{parent: k0}
		rope := &Rope{
			visited: make(map[graph.Co]bool),
			knots: map[int]*Knot{
				0: k0,
				1: k1,
			},
		}
		k00 := &Knot{
			co: graph.Co{X: 4, Y: 0},
		}
		k01 := &Knot{
			co:     graph.Co{X: 3, Y: 0},
			parent: k00,
		}
		want := &Rope{
			knots: map[int]*Knot{
				0: k00,
				1: k01,
			},
			visited: map[graph.Co]bool{
				{X: 0, Y: 0}: true,
				{X: 1, Y: 0}: true,
				{X: 2, Y: 0}: true,
				{X: 3, Y: 0}: true,
			},
		}
		err := rope.moveRope("R 4")
		assert.NoError(t, err)
		assert.Equal(t, want, rope)
	})

	t.Run("moves a 2 knot rope the given direction and distance, diagonally, advent of code example 2", func(t *testing.T) {
		k0 := &Knot{co: graph.Co{X: 4, Y: 0}}
		k1 := &Knot{
			co:     graph.Co{X: 3, Y: 0},
			parent: k0,
		}
		rope := &Rope{
			visited: map[graph.Co]bool{
				{X: 0, Y: 0}: true,
				{X: 1, Y: 0}: true,
				{X: 2, Y: 0}: true,
				{X: 3, Y: 0}: true,
			},
			knots: map[int]*Knot{
				0: k0,
				1: k1,
			},
		}

		k00 := &Knot{
			co: graph.Co{X: 4, Y: -4},
		}
		k01 := &Knot{
			co:     graph.Co{X: 4, Y: -3},
			parent: k00,
		}
		want := &Rope{
			knots: map[int]*Knot{
				0: k00,
				1: k01,
			},
			visited: map[graph.Co]bool{
				{X: 0, Y: 0}:  true,
				{X: 1, Y: 0}:  true,
				{X: 2, Y: 0}:  true,
				{X: 3, Y: 0}:  true,
				{X: 4, Y: -1}: true,
				{X: 4, Y: -2}: true,
				{X: 4, Y: -3}: true,
			},
		}
		err := rope.moveRope("U 4")
		assert.NoError(t, err)
		assert.Equal(t, want, rope)
	})

	t.Run("moves a 10 knot rope the given direction and distance, diagonally and non diagonally, advent of code example 3", func(t *testing.T) {
		k0 := &Knot{co: graph.Co{X: 0, Y: 2}}
		k1 := &Knot{co: graph.Co{X: 0, Y: 1}, parent: k0}
		k2 := &Knot{co: graph.Co{X: 1, Y: 0}, parent: k1}
		k3 := &Knot{co: graph.Co{X: 2, Y: 0}, parent: k2}
		k4 := &Knot{co: graph.Co{X: 3, Y: 0}, parent: k3}
		k5 := &Knot{co: graph.Co{X: 4, Y: 0}, parent: k4}
		k6 := &Knot{co: graph.Co{X: 4, Y: 1}, parent: k5}
		k7 := &Knot{co: graph.Co{X: 4, Y: 2}, parent: k6}
		k8 := &Knot{co: graph.Co{X: 4, Y: 3}, parent: k7}
		k9 := &Knot{co: graph.Co{X: 4, Y: 4}, parent: k8}
		rope := &Rope{
			visited: map[graph.Co]bool{},
			knots: map[int]*Knot{
				0: k0,
				1: k1,
				2: k2,
				3: k3,
				4: k4,
				5: k5,
				6: k6,
				7: k7,
				8: k8,
				9: k9,
			},
		}

		k00 := &Knot{co: graph.Co{X: 17, Y: 2}}
		k01 := &Knot{co: graph.Co{X: 16, Y: 2}, parent: k00}
		k02 := &Knot{co: graph.Co{X: 15, Y: 2}, parent: k01}
		k03 := &Knot{co: graph.Co{X: 14, Y: 2}, parent: k02}
		k04 := &Knot{co: graph.Co{X: 13, Y: 2}, parent: k03}
		k05 := &Knot{co: graph.Co{X: 12, Y: 2}, parent: k04}
		k06 := &Knot{co: graph.Co{X: 11, Y: 2}, parent: k05}
		k07 := &Knot{co: graph.Co{X: 10, Y: 2}, parent: k06}
		k08 := &Knot{co: graph.Co{X: 9, Y: 2}, parent: k07}
		k09 := &Knot{co: graph.Co{X: 8, Y: 2}, parent: k08}

		want := &Rope{
			knots: map[int]*Knot{
				0: k00,
				1: k01,
				2: k02,
				3: k03,
				4: k04,
				5: k05,
				6: k06,
				7: k07,
				8: k08,
				9: k09,
			},
			visited: map[graph.Co]bool{
				{X: 4, Y: 4}: true,
				{X: 5, Y: 3}: true,
				{X: 6, Y: 2}: true,
				{X: 7, Y: 2}: true,
				{X: 8, Y: 2}: true,
			},
		}
		err := rope.moveRope("R 17")
		assert.NoError(t, err)
		assert.Equal(t, want, rope)
	})
}

func TestFollowInstructions(t *testing.T) {
	t.Run("returns an error if running an instruction returns an error", func(t *testing.T) {
		input := []string{
			"R 4",
			"U 4",
			"L 3",
			"Forward 1",
			"R 4",
			"D 1",
			"L 5",
			"R 2",
		}
		k0 := &Knot{}
		k1 := &Knot{parent: k0}
		rope := &Rope{
			visited: make(map[graph.Co]bool),
			knots: map[int]*Knot{
				0: k0,
				1: k1,
			},
		}
		err := rope.followInstructions(input)
		assert.Error(t, err)
	})

	t.Run("runs instructions for a two knot rope", func(t *testing.T) {
		input := []string{
			"R 4",
			"U 4",
			"L 3",
			"D 1",
			"R 4",
			"D 1",
			"L 5",
			"R 2",
		}
		k0 := &Knot{}
		k1 := &Knot{parent: k0}
		rope := &Rope{
			visited: make(map[graph.Co]bool),
			knots: map[int]*Knot{
				0: k0,
				1: k1,
			},
		}
		err := rope.followInstructions(input)
		assert.NoError(t, err)
		assert.Equal(t, 13, len(rope.visited))
	})

	t.Run("runs instructions for a 10 knot rope", func(t *testing.T) {
		input := []string{
			"R 5",
			"U 8",
			"L 8",
			"D 3",
			"R 17",
			"D 10",
			"L 25",
			"U 20",
		}
		k0 := &Knot{}
		k1 := &Knot{parent: k0}
		k2 := &Knot{parent: k1}
		k3 := &Knot{parent: k2}
		k4 := &Knot{parent: k3}
		k5 := &Knot{parent: k4}
		k6 := &Knot{parent: k5}
		k7 := &Knot{parent: k6}
		k8 := &Knot{parent: k7}
		k9 := &Knot{parent: k8}
		rope := &Rope{
			visited: map[graph.Co]bool{},
			knots: map[int]*Knot{
				0: k0,
				1: k1,
				2: k2,
				3: k3,
				4: k4,
				5: k5,
				6: k6,
				7: k7,
				8: k8,
				9: k9,
			},
		}
		err := rope.followInstructions(input)
		assert.NoError(t, err)
		assert.Equal(t, 36, len(rope.visited))
	})
}
