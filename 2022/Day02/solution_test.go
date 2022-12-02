package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var shapes = map[string]Shape{
	"X": {
		ties:   "A",
		beats:  "C",
		loses:  "B",
		points: 1,
	},
	"A": {
		ties:   "A",
		beats:  "C",
		loses:  "B",
		points: 1,
	},
	"Y": {
		ties:   "B",
		beats:  "A",
		loses:  "C",
		points: 2,
	},
	"B": {
		ties:   "B",
		beats:  "A",
		loses:  "C",
		points: 2,
	},
	"Z": {
		ties:   "C",
		beats:  "B",
		loses:  "A",
		points: 3,
	},
	"C": {
		ties:   "C",
		beats:  "B",
		loses:  "A",
		points: 3,
	},
}

func TestCreateGame(t *testing.T) {
	t.Run("returns a game", func(t *testing.T) {
		got := createGame()
		game := &Game{
			shapes: shapes,
		}
		assert.Equal(t, game, got)
	})
}

func TestPlayRound(t *testing.T) {
	t.Run("returns an error if the round input is not 3 characters", func(t *testing.T) {
		g := &Game{
			shapes: shapes,
		}
		err := g.playRound("A  X")
		assert.Error(t, err)
	})

	t.Run("returns an error if the round input is malformed", func(t *testing.T) {
		g := &Game{
			shapes: shapes,
		}
		err := g.playRound("AvY")
		assert.Error(t, err)
	})

	testsPart1 := []struct {
		them  string
		me    string
		round string
		want  int
	}{
		{
			them:  "rock",
			me:    "rock",
			round: "A X",
			want:  4,
		},
		{
			them:  "rock",
			me:    "paper",
			round: "A Y",
			want:  8,
		},
		{
			them:  "rock",
			me:    "scissors",
			round: "A Z",
			want:  3,
		},
		{
			them:  "paper",
			me:    "rock",
			round: "B X",
			want:  1,
		},
		{
			them:  "paper",
			me:    "paper",
			round: "B Y",
			want:  5,
		},
		{
			them:  "paper",
			me:    "scissors",
			round: "B Z",
			want:  9,
		},
		{
			them:  "scissors",
			me:    "rock",
			round: "C X",
			want:  7,
		},
		{
			them:  "scissors",
			me:    "paper",
			round: "C Y",
			want:  2,
		},
		{
			them:  "scissors",
			me:    "scissors",
			round: "C Z",
			want:  6,
		},
	}
	for _, tt := range testsPart1 {
		t.Run(fmt.Sprintf("they play %s, I play %s in part 1 and get %d points", tt.them, tt.me, tt.want), func(t *testing.T) {
			game := &Game{
				shapes: shapes,
			}
			err := game.playRound(tt.round)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, game.part1)
		})
	}

	testsPart2 := []struct {
		them    string
		outcome string
		round   string
		want    int
	}{
		{
			them:    "rock",
			outcome: "lose",
			round:   "A X",
			want:    3,
		},
		{
			them:    "rock",
			outcome: "draw",
			round:   "A Y",
			want:    4,
		},
		{
			them:    "rock",
			outcome: "win",
			round:   "A Z",
			want:    8,
		},
		{
			them:    "paper",
			outcome: "lose",
			round:   "B X",
			want:    1,
		},
		{
			them:    "paper",
			outcome: "draw",
			round:   "B Y",
			want:    5,
		},
		{
			them:    "paper",
			outcome: "win",
			round:   "B Z",
			want:    9,
		},
		{
			them:    "scissors",
			outcome: "lose",
			round:   "C X",
			want:    2,
		},
		{
			them:    "scissors",
			outcome: "draw",
			round:   "C Y",
			want:    6,
		},
		{
			them:    "scissors",
			outcome: "win",
			round:   "C Z",
			want:    7,
		},
	}
	for _, tt := range testsPart2 {
		t.Run(fmt.Sprintf("they play %s, I need to %s in part 2, so I get %d points", tt.them, tt.outcome, tt.want), func(t *testing.T) {
			game := &Game{
				shapes: shapes,
			}
			err := game.playRound(tt.round)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, game.part2)
		})
	}
}

func TestPlayGames(t *testing.T) {
	t.Run("returns an error if an input line is not playable", func(t *testing.T) {
		game := &Game{
			shapes: shapes,
		}
		err := game.playGames([]string{"A X", "CvY", "B Z"})
		assert.Error(t, err)
	})

	t.Run("plays games from given input, advent of code example", func(t *testing.T) {
		game := &Game{
			shapes: shapes,
		}
		err := game.playGames([]string{"A Y", "B X", "C Z"})
		assert.NoError(t, err)
		assert.Equal(t, 15, game.part1)
		assert.Equal(t, 12, game.part2)
	})
}
