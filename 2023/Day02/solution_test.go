package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleLine(t *testing.T) {
	tests := []struct {
		line          string
		wantValidGame assert.BoolAssertionFunc
		wantPower     int
	}{
		{
			line:          "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			wantValidGame: assert.True,
			wantPower:     48,
		},
		{
			line:          "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			wantValidGame: assert.True,
			wantPower:     12,
		},
		{
			line:          "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			wantValidGame: assert.False,
			wantPower:     1560,
		},
		{
			line:          "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			wantValidGame: assert.False,
			wantPower:     630,
		},
		{
			line:          "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			wantValidGame: assert.True,
			wantPower:     36,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("assert line evaluates correctly for parts 1 and 2, example %d", i+1), func(t *testing.T) {
			validGame, power := handleLine(tt.line)
			tt.wantValidGame(t, validGame)
			assert.Equal(t, tt.wantPower, power)
		})
	}
}

func TestFindSolutions(t *testing.T) {
	t.Run("returns correct solutions for parts 1 and 2", func(t *testing.T) {
		input := []string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		}
		part1, part2 := findSolutions(input)
		assert.Equal(t, 8, part1)
		assert.Equal(t, 2286, part2)
	})
}
