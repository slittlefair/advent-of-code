package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFollowInstructions(t *testing.T) {
	t.Run("returns number of steps taken to find an exit, advent of code example", func(t *testing.T) {
		input := []int{
			0,
			3,
			0,
			1,
			-3,
		}
		got := followInstructions(input)
		assert.Equal(t, 5, got)
	})
}

func TestFollowInstructions2(t *testing.T) {
	t.Run("returns number of steps taken to find an exit, advent of code example", func(t *testing.T) {
		input := []int{
			0,
			3,
			0,
			1,
			-3,
		}
		got := followInstructions2(input)
		assert.Equal(t, 10, got)
	})
}
