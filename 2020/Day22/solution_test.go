package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame_parseInput(t *testing.T) {
	tests := []struct {
		name               string
		input              []string
		want               *Game
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if one of the rows in the input cannot be converted to an int",
			input: []string{
				"Player 1:",
				"9",
				"2",
				"6",
				"3",
				"1",
				"",
				"Player 2:",
				"5",
				"8",
				"abc",
				"7",
				"10",
			},
			want: &Game{
				player1: Deck{9, 2, 6, 3, 1},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "advent of code example 1",
			input: []string{
				"Player 1:",
				"9",
				"2",
				"6",
				"3",
				"1",
				"",
				"Player 2:",
				"5",
				"8",
				"4",
				"7",
				"10",
			},
			want: &Game{
				player1: Deck{9, 2, 6, 3, 1},
				player2: Deck{5, 8, 4, 7, 10},
			},
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "advent of code example 2",
			input: []string{
				"Player 1:",
				"43",
				"19",
				"",
				"Player 2:",
				"2",
				"29",
				"14",
			},
			want: &Game{
				player1: Deck{43, 19},
				player2: Deck{2, 29, 14},
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{}
			err := g.parseInput(tt.input)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, g)
		})
	}
}

func TestGame_player1Wins(t *testing.T) {
	type fields struct {
		player1 Deck
		player2 Deck
	}
	tests := []struct {
		name   string
		fields fields
		want   *Game
	}{
		{
			name: "advent of code example 1",
			fields: fields{
				player1: Deck{9, 2, 3, 6, 1},
				player2: Deck{5, 8, 4, 7, 10},
			},
			want: &Game{
				player1: Deck{2, 3, 6, 1, 9, 5},
				player2: Deck{8, 4, 7, 10},
			},
		},
		{
			name: "advent of code example 2",
			fields: fields{
				player1: Deck{8, 4, 2, 3, 6, 1, 9, 5},
				player2: Deck{7, 10},
			},
			want: &Game{
				player1: Deck{4, 2, 3, 6, 1, 9, 5, 8, 7},
				player2: Deck{10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				player1: tt.fields.player1,
				player2: tt.fields.player2,
			}
			g.player1Wins()
			assert.Equal(t, tt.want, g)
		})
	}
}

func TestGame_player2Wins(t *testing.T) {
	type fields struct {
		player1 Deck
		player2 Deck
	}
	tests := []struct {
		name   string
		fields fields
		want   *Game
	}{
		{
			name: "advent of code example 3",
			fields: fields{
				player1: Deck{3, 1, 9, 5, 6, 4},
				player2: Deck{7, 10, 8, 2},
			},
			want: &Game{
				player1: Deck{1, 9, 5, 6, 4},
				player2: Deck{10, 8, 2, 7, 3},
			},
		},
		{
			name: "advent of code example 4",
			fields: fields{
				player1: Deck{1},
				player2: Deck{7, 3, 2, 10, 6, 8, 5, 9, 4},
			},
			want: &Game{
				player1: Deck{},
				player2: Deck{3, 2, 10, 6, 8, 5, 9, 4, 7, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				player1: tt.fields.player1,
				player2: tt.fields.player2,
			}
			g.player2Wins()
			assert.Equal(t, tt.want, g)
		})
	}
}

func TestGame_playNormalRound(t *testing.T) {
	type fields struct {
		player1 Deck
		player2 Deck
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  *Game
	}{
		{
			name: "advent of code example1, player 1 wins",
			fields: fields{
				player1: Deck{9, 2, 3, 6, 1},
				player2: Deck{5, 8, 4, 7, 10},
			},
			want: Player1,
			want1: &Game{
				player1: Deck{2, 3, 6, 1, 9, 5},
				player2: Deck{8, 4, 7, 10},
			},
		},
		{
			name: "advent of code example 2, player 1 wins",
			fields: fields{
				player1: Deck{8, 4, 2, 3, 6, 1, 9, 5},
				player2: Deck{7, 10},
			},
			want: Player1,
			want1: &Game{
				player1: Deck{4, 2, 3, 6, 1, 9, 5, 8, 7},
				player2: Deck{10},
			},
		},
		{
			name: "advent of code example 3, player 2 wins",
			fields: fields{
				player1: Deck{3, 1, 9, 5, 6, 4},
				player2: Deck{7, 10, 8, 2},
			},
			want: Player2,
			want1: &Game{
				player1: Deck{1, 9, 5, 6, 4},
				player2: Deck{10, 8, 2, 7, 3},
			},
		},
		{
			name: "advent of code example 4, player 2 wins",
			fields: fields{
				player1: Deck{1},
				player2: Deck{7, 3, 2, 10, 6, 8, 5, 9, 4},
			},
			want: Player2,
			want1: &Game{
				player1: Deck{},
				player2: Deck{3, 2, 10, 6, 8, 5, 9, 4, 7, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				player1: tt.fields.player1,
				player2: tt.fields.player2,
			}
			got := g.playNormalRound()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, g)
		})
	}
}

func TestGame_playNormalGame(t *testing.T) {
	type fields struct {
		player1 Deck
		player2 Deck
	}
	tests := []struct {
		name   string
		fields fields
		want   Deck
	}{
		{
			name: "advent of code example, player1 win",
			fields: fields{
				player1: Deck{5, 8, 4, 7, 10},
				player2: Deck{9, 2, 6, 3, 1},
			},
			want: Deck{3, 2, 10, 6, 8, 5, 9, 4, 7, 1},
		},
		{
			name: "advent of code example, player2 win",
			fields: fields{
				player1: Deck{9, 8, 7, 6, 5},
				player2: Deck{10, 1, 2, 3, 4},
			},
			want: Deck{9, 6, 10, 5, 7, 4, 8, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				player1: tt.fields.player1,
				player2: tt.fields.player2,
			}
			got := g.playNormalGame()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGame_deckSeen(t *testing.T) {
	type fields struct {
		player1 Deck
		player2 Deck
	}
	tests := []struct {
		name   string
		fields fields
		seen   []Game
		want   bool
	}{
		{
			name: "returns false if there are no seen hands the same as the current one",
			fields: fields{
				player1: Deck{1, 2},
				player2: Deck{3, 4},
			},
			seen: []Game{
				{
					player1: Deck{2, 3},
					player2: Deck{4, 1},
				},
				{
					player1: Deck{4, 3},
					player2: Deck{2, 1},
				},
			},
			want: false,
		},
		{
			name: "returns false if there are seen hands the same as the current one, but not at the same time",
			fields: fields{
				player1: Deck{1, 2},
				player2: Deck{3, 4},
			},
			seen: []Game{
				{
					player1: Deck{2, 1},
					player2: Deck{3, 4},
				},
				{
					player1: Deck{4, 3},
					player2: Deck{1, 2},
				},
			},
			want: false,
		},
		{
			name: "returns false if the current hands have been seen at the same time before but by the other player",
			fields: fields{
				player1: Deck{1, 2},
				player2: Deck{3, 4},
			},
			seen: []Game{
				{
					player1: Deck{2, 1},
					player2: Deck{3, 4},
				},
				{
					player1: Deck{4, 3},
					player2: Deck{1, 2},
				},
				{
					player1: Deck{3, 4},
					player2: Deck{1, 2},
				},
				{
					player1: Deck{2, 1},
					player2: Deck{3, 4},
				},
				{
					player1: Deck{4, 3},
					player2: Deck{1, 2},
				},
			},
			want: false,
		},
		{
			name: "returns true if the current hands have been seen at the same time before",
			fields: fields{
				player1: Deck{1, 2},
				player2: Deck{3, 4},
			},
			seen: []Game{
				{
					player1: Deck{2, 1},
					player2: Deck{3, 4},
				},
				{
					player1: Deck{4, 3},
					player2: Deck{1, 2},
				},
				{
					player1: Deck{2, 1},
					player2: Deck{3, 4},
				},
				{
					player1: Deck{1, 2},
					player2: Deck{3, 4},
				},
				{
					player1: Deck{4, 3},
					player2: Deck{1, 2},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Game{
				player1: tt.fields.player1,
				player2: tt.fields.player2,
			}
			got := g.deckSeen(tt.seen)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGame_playRecursiveRound(t *testing.T) {
	type fields struct {
		player1 Deck
		player2 Deck
	}
	tests := []struct {
		name   string
		fields fields
		seen   []Game
		want   []Game
		want1  *Game
	}{
		{
			name: "advent of code example 1, don't start a new game",
			fields: fields{
				player1: Deck{9, 2, 6, 3, 1},
				player2: Deck{5, 8, 4, 7, 10},
			},
			seen: []Game{},
			want: []Game{
				{
					player1: Deck{9, 2, 6, 3, 1},
					player2: Deck{5, 8, 4, 7, 10},
				},
			},
			want1: &Game{
				player1: Deck{2, 6, 3, 1, 9, 5},
				player2: Deck{8, 4, 7, 10},
			},
		},
		{
			name: "advent of code example 2, start a new game that player 2 wins",
			fields: fields{
				player1: Deck{4, 9, 8, 5, 2},
				player2: Deck{3, 10, 1, 7, 6},
			},
			seen: []Game{
				{
					player1: Deck{9, 2, 6, 3, 1},
					player2: Deck{5, 8, 4, 7, 10},
				},
				{
					player1: Deck{2, 6, 3, 1, 9, 5},
					player2: Deck{8, 4, 7, 10},
				},
				{
					player1: Deck{6, 3, 1, 9, 5},
					player2: Deck{4, 7, 10, 8, 2},
				},
				{
					player1: Deck{3, 1, 9, 5, 6, 4},
					player2: Deck{7, 10, 8, 2},
				},
				{
					player1: Deck{1, 9, 5, 6, 4},
					player2: Deck{10, 8, 2, 7, 3},
				},
				{
					player1: Deck{9, 5, 6, 4},
					player2: Deck{8, 2, 7, 3, 10, 1},
				},
				{
					player1: Deck{5, 6, 4, 9, 8},
					player2: Deck{2, 7, 3, 10, 1},
				},
				{
					player1: Deck{6, 4, 9, 8, 5, 2},
					player2: Deck{7, 3, 10, 1},
				},
			},
			want: []Game{
				{
					player1: Deck{9, 2, 6, 3, 1},
					player2: Deck{5, 8, 4, 7, 10},
				},
				{
					player1: Deck{2, 6, 3, 1, 9, 5},
					player2: Deck{8, 4, 7, 10},
				},
				{
					player1: Deck{6, 3, 1, 9, 5},
					player2: Deck{4, 7, 10, 8, 2},
				},
				{
					player1: Deck{3, 1, 9, 5, 6, 4},
					player2: Deck{7, 10, 8, 2},
				},
				{
					player1: Deck{1, 9, 5, 6, 4},
					player2: Deck{10, 8, 2, 7, 3},
				},
				{
					player1: Deck{9, 5, 6, 4},
					player2: Deck{8, 2, 7, 3, 10, 1},
				},
				{
					player1: Deck{5, 6, 4, 9, 8},
					player2: Deck{2, 7, 3, 10, 1},
				},
				{
					player1: Deck{6, 4, 9, 8, 5, 2},
					player2: Deck{7, 3, 10, 1},
				},
				{
					player1: Deck{4, 9, 8, 5, 2},
					player2: Deck{3, 10, 1, 7, 6},
				},
			},
			want1: &Game{
				player1: Deck{9, 8, 5, 2},
				player2: Deck{10, 1, 7, 6, 3, 4},
			},
		},
		{
			name: "advent of code example 2, start a new game that player 1 wins",
			fields: fields{
				player1: Deck{4, 10, 9, 7, 5},
				player2: Deck{1, 8, 3},
			},
			seen: []Game{
				{
					player1: Deck{10, 9, 7, 5, 4, 1},
					player2: Deck{8, 3},
				},
			},
			want: []Game{
				{
					player1: Deck{10, 9, 7, 5, 4, 1},
					player2: Deck{8, 3},
				},
				{
					player1: Deck{4, 10, 9, 7, 5},
					player2: Deck{1, 8, 3},
				},
			},
			want1: &Game{
				player1: Deck{10, 9, 7, 5, 4, 1},
				player2: Deck{8, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				player1: tt.fields.player1,
				player2: tt.fields.player2,
			}
			got := g.playRecursiveRound(tt.seen)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, g)
		})
	}
}

func TestGame_playRecursiveGame(t *testing.T) {
	type fields struct {
		player1 Deck
		player2 Deck
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  Deck
	}{
		{
			name: "advent of code example 1, game ends with a full hand",
			fields: fields{
				player1: Deck{9, 2, 6, 3, 1},
				player2: Deck{5, 8, 4, 7, 10},
			},
			want:  Player2,
			want1: Deck{7, 5, 6, 2, 4, 1, 10, 8, 9, 3},
		},
		{
			name: "advent of code example 2, game ends due to recursion",
			fields: fields{
				player1: Deck{43, 19},
				player2: Deck{2, 29, 14},
			},
			want:  Player1,
			want1: Deck{43, 19},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				player1: tt.fields.player1,
				player2: tt.fields.player2,
			}
			got, got1 := g.playRecursiveGame()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestGame_calculateWinningScore(t *testing.T) {
	tests := []struct {
		name               string
		deck               Deck
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "returns an error if there are no cards in the deck",
			deck:               Deck{},
			want:               0,
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "advent of code example 1",
			deck:               Deck{3, 2, 10, 6, 8, 5, 9, 4, 7, 1},
			want:               306,
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "advent of code example 2",
			deck:               Deck{7, 5, 6, 2, 4, 1, 10, 8, 9, 3},
			want:               291,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateWinningScore(tt.deck)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
