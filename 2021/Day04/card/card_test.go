package card_test

import (
	"Advent-of-Code/2021/Day04/card"
	"Advent-of-Code/graph"
	"Advent-of-Code/regex"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCard_ParseCard(t *testing.T) {
	type args struct {
		lines []string
		reNum *regexp.Regexp
	}
	tests := []struct {
		name               string
		card               *card.Card
		args               args
		want               *card.Card
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if match from regex can't be converted to int",
			card: &card.Card{
				Numbers: make(map[graph.Co]*card.Number),
			},
			args: args{
				lines: []string{
					"",
					"1 2 one",
					"4 5 6",
				},
				reNum: regexp.MustCompile(`\w`),
			},
			want: &card.Card{
				Numbers: map[graph.Co]*card.Number{
					{X: 0, Y: 0}: {Val: 1},
					{X: 1, Y: 0}: {Val: 2},
				},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns a parsed card from input, advent of code example 1",
			card: &card.Card{
				Numbers: make(map[graph.Co]*card.Number),
			},
			args: args{
				lines: []string{
					"",
					"22 13 17 11  0",
					"8  2 23  4 24",
					"21  9 14 16  7",
					"6 10  3 18  5",
					"1 12 20 15 19",
				},
				reNum: regex.MatchNums,
			},
			want: &card.Card{
				Numbers: map[graph.Co]*card.Number{
					{X: 0, Y: 0}: {Val: 22},
					{X: 0, Y: 1}: {Val: 8},
					{X: 0, Y: 2}: {Val: 21},
					{X: 0, Y: 3}: {Val: 6},
					{X: 0, Y: 4}: {Val: 1},
					{X: 1, Y: 0}: {Val: 13},
					{X: 1, Y: 1}: {Val: 2},
					{X: 1, Y: 2}: {Val: 9},
					{X: 1, Y: 3}: {Val: 10},
					{X: 1, Y: 4}: {Val: 12},
					{X: 2, Y: 0}: {Val: 17},
					{X: 2, Y: 1}: {Val: 23},
					{X: 2, Y: 2}: {Val: 14},
					{X: 2, Y: 3}: {Val: 3},
					{X: 2, Y: 4}: {Val: 20},
					{X: 3, Y: 0}: {Val: 11},
					{X: 3, Y: 1}: {Val: 4},
					{X: 3, Y: 2}: {Val: 16},
					{X: 3, Y: 3}: {Val: 18},
					{X: 3, Y: 4}: {Val: 15},
					{X: 4, Y: 0}: {Val: 0},
					{X: 4, Y: 1}: {Val: 24},
					{X: 4, Y: 2}: {Val: 7},
					{X: 4, Y: 3}: {Val: 5},
					{X: 4, Y: 4}: {Val: 19},
				},
			},
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns a parsed card from input, advent of code example 2",
			card: &card.Card{
				Numbers: make(map[graph.Co]*card.Number),
			},
			args: args{
				lines: []string{
					"					",
					"3 15  0  2 22",
					"9 18 13 17  5",
					"19  8  7 25 23",
					"20 11 10 24  4",
					"14 21 16 12  6",
				},
				reNum: regex.MatchNums,
			},
			want: &card.Card{
				Numbers: map[graph.Co]*card.Number{
					{X: 0, Y: 0}: {Val: 3},
					{X: 0, Y: 1}: {Val: 9},
					{X: 0, Y: 2}: {Val: 19},
					{X: 0, Y: 3}: {Val: 20},
					{X: 0, Y: 4}: {Val: 14},
					{X: 1, Y: 0}: {Val: 15},
					{X: 1, Y: 1}: {Val: 18},
					{X: 1, Y: 2}: {Val: 8},
					{X: 1, Y: 3}: {Val: 11},
					{X: 1, Y: 4}: {Val: 21},
					{X: 2, Y: 0}: {Val: 0},
					{X: 2, Y: 1}: {Val: 13},
					{X: 2, Y: 2}: {Val: 7},
					{X: 2, Y: 3}: {Val: 10},
					{X: 2, Y: 4}: {Val: 16},
					{X: 3, Y: 0}: {Val: 2},
					{X: 3, Y: 1}: {Val: 17},
					{X: 3, Y: 2}: {Val: 25},
					{X: 3, Y: 3}: {Val: 24},
					{X: 3, Y: 4}: {Val: 12},
					{X: 4, Y: 0}: {Val: 22},
					{X: 4, Y: 1}: {Val: 5},
					{X: 4, Y: 2}: {Val: 23},
					{X: 4, Y: 3}: {Val: 4},
					{X: 4, Y: 4}: {Val: 6},
				},
			},
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns a parsed card from input, advent of code example 3",
			card: &card.Card{
				Numbers: make(map[graph.Co]*card.Number),
			},
			args: args{
				lines: []string{
					"",
					"14 21 17 24  4",
					"10 16 15  9 19",
					"18  8 23 26 20",
					"22 11 13  6  5",
					"2  0 12  3  7",
				},
				reNum: regex.MatchNums,
			},
			want: &card.Card{
				Numbers: map[graph.Co]*card.Number{
					{X: 0, Y: 0}: {Val: 14},
					{X: 0, Y: 1}: {Val: 10},
					{X: 0, Y: 2}: {Val: 18},
					{X: 0, Y: 3}: {Val: 22},
					{X: 0, Y: 4}: {Val: 2},
					{X: 1, Y: 0}: {Val: 21},
					{X: 1, Y: 1}: {Val: 16},
					{X: 1, Y: 2}: {Val: 8},
					{X: 1, Y: 3}: {Val: 11},
					{X: 1, Y: 4}: {Val: 0},
					{X: 2, Y: 0}: {Val: 17},
					{X: 2, Y: 1}: {Val: 15},
					{X: 2, Y: 2}: {Val: 23},
					{X: 2, Y: 3}: {Val: 13},
					{X: 2, Y: 4}: {Val: 12},
					{X: 3, Y: 0}: {Val: 24},
					{X: 3, Y: 1}: {Val: 9},
					{X: 3, Y: 2}: {Val: 26},
					{X: 3, Y: 3}: {Val: 6},
					{X: 3, Y: 4}: {Val: 3},
					{X: 4, Y: 0}: {Val: 4},
					{X: 4, Y: 1}: {Val: 19},
					{X: 4, Y: 2}: {Val: 20},
					{X: 4, Y: 3}: {Val: 5},
					{X: 4, Y: 4}: {Val: 7},
				},
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.card
			err := c.ParseCard(tt.args.lines, tt.args.reNum)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, c)
		})
	}
}

func TestCard_CardIsWinner(t *testing.T) {
	tests := []struct {
		name    string
		numbers map[graph.Co]*card.Number
		want    bool
	}{
		{
			name: "returns false if card has no winning rows or columns",
			numbers: map[graph.Co]*card.Number{
				{X: 0, Y: 0}: {Val: 14, Called: true},
				{X: 0, Y: 1}: {Val: 10},
				{X: 0, Y: 2}: {Val: 18},
				{X: 0, Y: 3}: {Val: 22},
				{X: 0, Y: 4}: {Val: 2},
				{X: 1, Y: 0}: {Val: 21},
				{X: 1, Y: 1}: {Val: 16},
				{X: 1, Y: 2}: {Val: 8},
				{X: 1, Y: 3}: {Val: 11, Called: true},
				{X: 1, Y: 4}: {Val: 0},
				{X: 2, Y: 0}: {Val: 17},
				{X: 2, Y: 1}: {Val: 15},
				{X: 2, Y: 2}: {Val: 23},
				{X: 2, Y: 3}: {Val: 13},
				{X: 2, Y: 4}: {Val: 12},
				{X: 3, Y: 0}: {Val: 24},
				{X: 3, Y: 1}: {Val: 9},
				{X: 3, Y: 2}: {Val: 26},
				{X: 3, Y: 3}: {Val: 6, Called: true},
				{X: 3, Y: 4}: {Val: 3},
				{X: 4, Y: 0}: {Val: 4},
				{X: 4, Y: 1}: {Val: 19},
				{X: 4, Y: 2}: {Val: 20},
				{X: 4, Y: 3}: {Val: 5},
				{X: 4, Y: 4}: {Val: 7},
			},
			want: false,
		},
		{
			name: "returns true if card has a winning column",
			numbers: map[graph.Co]*card.Number{
				{X: 0, Y: 0}: {Val: 14, Called: true},
				{X: 0, Y: 1}: {Val: 10, Called: true},
				{X: 0, Y: 2}: {Val: 18},
				{X: 0, Y: 3}: {Val: 22},
				{X: 0, Y: 4}: {Val: 2},
				{X: 1, Y: 0}: {Val: 21, Called: true},
				{X: 1, Y: 1}: {Val: 16},
				{X: 1, Y: 2}: {Val: 8},
				{X: 1, Y: 3}: {Val: 11, Called: true},
				{X: 1, Y: 4}: {Val: 0},
				{X: 2, Y: 0}: {Val: 17},
				{X: 2, Y: 1}: {Val: 15},
				{X: 2, Y: 2}: {Val: 23},
				{X: 2, Y: 3}: {Val: 13},
				{X: 2, Y: 4}: {Val: 12},
				{X: 3, Y: 0}: {Val: 24, Called: true},
				{X: 3, Y: 1}: {Val: 9, Called: true},
				{X: 3, Y: 2}: {Val: 26, Called: true},
				{X: 3, Y: 3}: {Val: 6, Called: true},
				{X: 3, Y: 4}: {Val: 3, Called: true},
				{X: 4, Y: 0}: {Val: 4},
				{X: 4, Y: 1}: {Val: 19},
				{X: 4, Y: 2}: {Val: 20},
				{X: 4, Y: 3}: {Val: 5},
				{X: 4, Y: 4}: {Val: 7},
			},
			want: true,
		},
		{
			name: "returns true if card has a winning row",
			numbers: map[graph.Co]*card.Number{
				{X: 0, Y: 0}: {Val: 14, Called: true},
				{X: 0, Y: 1}: {Val: 10, Called: true},
				{X: 0, Y: 2}: {Val: 18},
				{X: 0, Y: 3}: {Val: 22},
				{X: 0, Y: 4}: {Val: 2, Called: true},
				{X: 1, Y: 0}: {Val: 21, Called: true},
				{X: 1, Y: 1}: {Val: 16},
				{X: 1, Y: 2}: {Val: 8},
				{X: 1, Y: 3}: {Val: 11, Called: true},
				{X: 1, Y: 4}: {Val: 0, Called: true},
				{X: 2, Y: 0}: {Val: 17},
				{X: 2, Y: 1}: {Val: 15},
				{X: 2, Y: 2}: {Val: 23},
				{X: 2, Y: 3}: {Val: 13},
				{X: 2, Y: 4}: {Val: 12, Called: true},
				{X: 3, Y: 0}: {Val: 24, Called: true},
				{X: 3, Y: 1}: {Val: 9},
				{X: 3, Y: 2}: {Val: 26, Called: true},
				{X: 3, Y: 3}: {Val: 6, Called: true},
				{X: 3, Y: 4}: {Val: 3, Called: true},
				{X: 4, Y: 0}: {Val: 4},
				{X: 4, Y: 1}: {Val: 19},
				{X: 4, Y: 2}: {Val: 20},
				{X: 4, Y: 3}: {Val: 5},
				{X: 4, Y: 4}: {Val: 7, Called: true},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &card.Card{
				Numbers: tt.numbers,
			}
			got := c.CardIsWinner()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCard_CalculateScore(t *testing.T) {
	tests := []struct {
		name    string
		numbers map[graph.Co]*card.Number
		num     int
		want    int
	}{
		{
			name: "returns the product of given num and uncalled numbers from card, advent of code example 1",
			numbers: map[graph.Co]*card.Number{
				{X: 0, Y: 0}: {Val: 14, Called: true},
				{X: 0, Y: 1}: {Val: 10},
				{X: 0, Y: 2}: {Val: 18},
				{X: 0, Y: 3}: {Val: 22},
				{X: 0, Y: 4}: {Val: 2, Called: true},
				{X: 1, Y: 0}: {Val: 21, Called: true},
				{X: 1, Y: 1}: {Val: 16},
				{X: 1, Y: 2}: {Val: 8},
				{X: 1, Y: 3}: {Val: 11, Called: true},
				{X: 1, Y: 4}: {Val: 0, Called: true},
				{X: 2, Y: 0}: {Val: 17, Called: true},
				{X: 2, Y: 1}: {Val: 15},
				{X: 2, Y: 2}: {Val: 23, Called: true},
				{X: 2, Y: 3}: {Val: 13},
				{X: 2, Y: 4}: {Val: 12},
				{X: 3, Y: 0}: {Val: 24, Called: true},
				{X: 3, Y: 1}: {Val: 9, Called: true},
				{X: 3, Y: 2}: {Val: 26},
				{X: 3, Y: 3}: {Val: 6},
				{X: 3, Y: 4}: {Val: 3},
				{X: 4, Y: 0}: {Val: 4, Called: true},
				{X: 4, Y: 1}: {Val: 19},
				{X: 4, Y: 2}: {Val: 20},
				{X: 4, Y: 3}: {Val: 5, Called: true},
				{X: 4, Y: 4}: {Val: 7, Called: true},
			},
			num:  24,
			want: 4512,
		},
		{
			name: "returns the product of given num and uncalled numbers from card, advent of code example 2",
			numbers: map[graph.Co]*card.Number{
				{X: 0, Y: 0}: {Val: 3},
				{X: 0, Y: 1}: {Val: 9, Called: true},
				{X: 0, Y: 2}: {Val: 19},
				{X: 0, Y: 3}: {Val: 20},
				{X: 0, Y: 4}: {Val: 14, Called: true},
				{X: 1, Y: 0}: {Val: 15},
				{X: 1, Y: 1}: {Val: 18},
				{X: 1, Y: 2}: {Val: 8},
				{X: 1, Y: 3}: {Val: 11, Called: true},
				{X: 1, Y: 4}: {Val: 21, Called: true},
				{X: 2, Y: 0}: {Val: 0, Called: true},
				{X: 2, Y: 1}: {Val: 13, Called: true},
				{X: 2, Y: 2}: {Val: 7, Called: true},
				{X: 2, Y: 3}: {Val: 10, Called: true},
				{X: 2, Y: 4}: {Val: 16, Called: true},
				{X: 3, Y: 0}: {Val: 2, Called: true},
				{X: 3, Y: 1}: {Val: 17, Called: true},
				{X: 3, Y: 2}: {Val: 25},
				{X: 3, Y: 3}: {Val: 24, Called: true},
				{X: 3, Y: 4}: {Val: 12},
				{X: 4, Y: 0}: {Val: 22},
				{X: 4, Y: 1}: {Val: 5, Called: true},
				{X: 4, Y: 2}: {Val: 23, Called: true},
				{X: 4, Y: 3}: {Val: 4, Called: true},
				{X: 4, Y: 4}: {Val: 6},
			},
			num:  13,
			want: 1924,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := card.Card{
				Numbers: tt.numbers,
			}
			got := c.CalculateScore(tt.num)
			assert.Equal(t, tt.want, got)
		})
	}
}
