package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInput(t *testing.T) {
	t.Run("returns a parsed set of equations from a given input", func(t *testing.T) {
		input := []string{
			"190: 10 19",
			"3267: 81 40 27",
			"83: 17 5",
			"156: 15 6",
			"7290: 6 8 6 15",
			"161011: 16 10 13",
			"192: 17 8 14",
			"21037: 9 7 18 13",
			"292: 11 6 16 20",
		}
		want := []Equation{
			{
				testValue: 190,
				equations: []int{10, 19},
			},
			{
				testValue: 3267,
				equations: []int{81, 40, 27},
			},
			{
				testValue: 83,
				equations: []int{17, 5},
			},
			{
				testValue: 156,
				equations: []int{15, 6},
			},
			{
				testValue: 7290,
				equations: []int{6, 8, 6, 15},
			},
			{
				testValue: 161011,
				equations: []int{16, 10, 13},
			},
			{
				testValue: 192,
				equations: []int{17, 8, 14},
			},
			{
				testValue: 21037,
				equations: []int{9, 7, 18, 13},
			},
			{
				testValue: 292,
				equations: []int{11, 6, 16, 20},
			},
		}
		got := parseInput(input)
		assert.Equal(t, want, got)
	})
}

func TestEquation_evaluateEquation(t *testing.T) {
	t.Run("returns true for both parts for an equation true for part 1 ops", func(t *testing.T) {
		eq := Equation{
			testValue: 190,
			equations: []int{10, 19},
		}

		part1, part2 := eq.evaluateEquation()
		assert.True(t, part1)
		assert.True(t, part2)
	})

	t.Run("returns true then false for an equation true for part 2 only", func(t *testing.T) {
		eq := Equation{
			testValue: 7290,
			equations: []int{6, 8, 6, 15},
		}

		part1, part2 := eq.evaluateEquation()
		assert.False(t, part1)
		assert.True(t, part2)
	})

	t.Run("returns false for an equation that is false for both parts", func(t *testing.T) {
		eq := Equation{
			testValue: 21037,
			equations: []int{9, 7, 18, 13},
		}

		part1, part2 := eq.evaluateEquation()
		assert.False(t, part1)
		assert.False(t, part2)
	})

	t.Run("returns false for an equation that is true before all numbers have been evaluated", func(t *testing.T) {
		eq := Equation{
			testValue: 123,
			equations: []int{1, 2, 3, 4},
		}

		part1, part2 := eq.evaluateEquation()
		assert.False(t, part1)
		assert.False(t, part2)
	})
}

func Test_findSolutions(t *testing.T) {
	t.Run("returns solutions for part1 and part2 for a given input", func(t *testing.T) {
		input := []string{
			"190: 10 19",
			"3267: 81 40 27",
			"83: 17 5",
			"156: 15 6",
			"7290: 6 8 6 15",
			"161011: 16 10 13",
			"192: 17 8 14",
			"21037: 9 7 18 13",
			"292: 11 6 16 20",
		}

		part1, part2 := findSolutions(input)
		assert.Equal(t, 3749, part1)
		assert.Equal(t, 11387, part2)
	})
}
