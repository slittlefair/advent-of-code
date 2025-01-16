package main

import (
	"Advent-of-Code/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getNumsFromLine(t *testing.T) {
	t.Run("returns an error if the line doesn't have the expected prefix", func(t *testing.T) {
		got, err := getNumsFromLine("Button A: X+94, Y+34", "Button B")
		assert.Nil(t, got)
		assert.ErrorContains(t, err, "malformed input, expected")
	})

	t.Run("returns an error if there aren't two numbers in the line", func(t *testing.T) {
		got, err := getNumsFromLine("Button A: X+94, Y+one", "Button A")
		assert.Nil(t, got)
		assert.ErrorContains(t, err, "malformed input, expected 2 nums")
	})

	t.Run("returns the numbers int he given line", func(t *testing.T) {
		got, err := getNumsFromLine("Button A: X+94, Y+34", "Button A")
		assert.NoError(t, err)
		assert.Equal(t, []int{94, 34}, got)
	})
}

func Test_parseInput(t *testing.T) {
	t.Run("returns an error if Button A line is malformed", func(t *testing.T) {
		input := []string{
			"Button a: X+94, Y+34",
			"Button B: X+22, Y+67",
			"Prize: X=8400, Y=5400",
			"",
		}
		cm, err := parseInput(input)
		assert.Nil(t, cm)
		assert.ErrorContains(t, err, `malformed input, expected "Button A:", got`)
	})

	t.Run("returns an error if Button B line is malformed", func(t *testing.T) {
		input := []string{
			"Button A: X+94, Y+34",
			"Button B: X+, Y+67",
			"Prize: X=8400, Y=5400",
			"",
		}
		cm, err := parseInput(input)
		assert.Nil(t, cm)
		assert.ErrorContains(t, err, "malformed input, expected 2 nums")
	})

	t.Run("returns an error if Prize line is malformed", func(t *testing.T) {
		input := []string{
			"Button A: X+94, Y+34",
			"Button B: X+22, Y+67",
			"Prizee: X=8400, Y=5400",
			"",
		}
		cm, err := parseInput(input)
		assert.Nil(t, cm)
		assert.ErrorContains(t, err, `malformed input, expected "Prize:", got`)
	})

	t.Run("returns an error if there's not an empty line when we expect one", func(t *testing.T) {
		input := []string{
			"Button A: X+94, Y+34",
			"Button B: X+22, Y+67",
			"Prize: X=8400, Y=5400",
			"Button A: X+23, Y+879",
		}
		cm, err := parseInput(input)
		assert.Nil(t, cm)
		assert.ErrorContains(t, err, `malformed input, expected "", got`)
	})

	t.Run("returns a list of valid claw machines", func(t *testing.T) {
		input := []string{
			"Button A: X+94, Y+34",
			"Button B: X+22, Y+67",
			"Prize: X=8400, Y=5400",
			"",
			"Button A: X+26, Y+66",
			"Button B: X+67, Y+21",
			"Prize: X=12748, Y=12176",
			"",
			"Button A: X+17, Y+86",
			"Button B: X+84, Y+37",
			"Prize: X=7870, Y=6450",
			"",
			"Button A: X+69, Y+23",
			"Button B: X+27, Y+71",
			"Prize: X=18641, Y=10279",
		}

		want := []ClawMachine{
			{
				ButtonA: graph.Co{X: 94, Y: 34},
				ButtonB: graph.Co{X: 22, Y: 67},
				Prize:   graph.Co{X: 8400, Y: 5400},
			},
			{
				ButtonA: graph.Co{X: 26, Y: 66},
				ButtonB: graph.Co{X: 67, Y: 21},
				Prize:   graph.Co{X: 12748, Y: 12176},
			},
			{
				ButtonA: graph.Co{X: 17, Y: 86},
				ButtonB: graph.Co{X: 84, Y: 37},
				Prize:   graph.Co{X: 7870, Y: 6450},
			},
			{
				ButtonA: graph.Co{X: 69, Y: 23},
				ButtonB: graph.Co{X: 27, Y: 71},
				Prize:   graph.Co{X: 18641, Y: 10279},
			},
		}

		cm, err := parseInput(input)
		assert.NoError(t, err)
		assert.Equal(t, want, cm)
	})
}

func TestClawMachine_findTokensForWin(t *testing.T) {
	t.Run("returns correct number of tokens needed to win, 1", func(t *testing.T) {
		cm := ClawMachine{
			ButtonA: graph.Co{X: 94, Y: 34},
			ButtonB: graph.Co{X: 22, Y: 67},
			Prize:   graph.Co{X: 8400, Y: 5400},
		}
		got := cm.findTokensForWin()
		assert.Equal(t, 280, got)
	})

	t.Run("returns correct number of tokens needed to win, 2", func(t *testing.T) {
		cm := ClawMachine{
			ButtonA: graph.Co{X: 17, Y: 86},
			ButtonB: graph.Co{X: 84, Y: 37},
			Prize:   graph.Co{X: 7870, Y: 6450},
		}
		got := cm.findTokensForWin()
		assert.Equal(t, 200, got)
	})

	t.Run("returns 0 if the prize cannot be won, 1", func(t *testing.T) {
		cm := ClawMachine{
			ButtonA: graph.Co{X: 26, Y: 66},
			ButtonB: graph.Co{X: 67, Y: 21},
			Prize:   graph.Co{X: 12748, Y: 12176},
		}
		got := cm.findTokensForWin()
		assert.Zero(t, got)
	})

	t.Run("returns 0 if the prize cannot be won, 1", func(t *testing.T) {
		cm := ClawMachine{
			ButtonA: graph.Co{X: 69, Y: 23},
			ButtonB: graph.Co{X: 27, Y: 71},
			Prize:   graph.Co{X: 18641, Y: 10279},
		}
		got := cm.findTokensForWin()
		assert.Zero(t, got)
	})
}

func Test_findSolutions(t *testing.T) {
	t.Run("returns an error if the input is malformed", func(t *testing.T) {
		input := []string{
			"Button A: X+94, Y+34",
			"Button B: X+22, Y+67",
			"Prize: X=8400, Y=5400",
			"",
			"Button A: X+26, Y+66",
			"Button B: X+67, Y+21",
			"Prize: X=12748, Y=12176",
			"",
			"Button A: X+17, Y+86",
			"Button B: X+84, Y+37",
			"Prize: X=7870, Y=6450",
			"",
			"Button A: X+69, Y+23",
			"Button D: X+27, Y+71",
			"Prize: X=18641, Y=10279",
		}

		part1, part2, err := findSolutions(input)
		assert.Zero(t, part1)
		assert.Zero(t, part2)
		assert.Error(t, err)
	})

	t.Run("returns correct answers for parts 1 and 2 for a given input", func(t *testing.T) {
		input := []string{
			"Button A: X+94, Y+34",
			"Button B: X+22, Y+67",
			"Prize: X=8400, Y=5400",
			"",
			"Button A: X+26, Y+66",
			"Button B: X+67, Y+21",
			"Prize: X=12748, Y=12176",
			"",
			"Button A: X+17, Y+86",
			"Button B: X+84, Y+37",
			"Prize: X=7870, Y=6450",
			"",
			"Button A: X+69, Y+23",
			"Button B: X+27, Y+71",
			"Prize: X=18641, Y=10279",
		}

		part1, part2, err := findSolutions(input)
		assert.NoError(t, err)
		assert.Equal(t, 480, part1)
		assert.Equal(t, 875318608908, part2)
	})
}
