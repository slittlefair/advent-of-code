package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayTetris(t *testing.T) {
	t.Run("returns an error if the input contains an invalid instruction", func(t *testing.T) {
		instructions := ">>><<><>><<<>><>>><<<>>^<<<><<<>><>><<>>"
		got, got1, err := playTetris(instructions, 2022, 1000000000000)
		assert.Error(t, err)
		assert.Equal(t, -1, got)
		assert.Equal(t, -1, got1)
	})

	t.Run("returns correct solutions for parts 1 and 2, advent of code example", func(t *testing.T) {
		instructions := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
		got, got1, err := playTetris(instructions, 2022, 1000000000000)
		assert.NoError(t, err)
		assert.Equal(t, 3068, got)
		assert.Equal(t, 1514285714288, got1)
	})
}
