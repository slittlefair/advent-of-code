package game

import (
	"Advent-of-Code/2021/Day04/card"
	utils "Advent-of-Code/utils"
	"reflect"
	"testing"
)

func Test_parseNums(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    []int
		wantErr bool
	}{
		{
			name:    "returns an error if a number on the card cannot be converted to int",
			str:     "1,23,14,two,9,21",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "returns a slice of converted ints from the given string",
			str:     "1,23,14,2,9,21",
			want:    []int{1, 23, 14, 2, 9, 21},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseNums(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseNums() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseNums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInput(t *testing.T) {
	var card0 = &card.Card{
		Numbers: map[utils.Co]*card.Number{
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
	}
	var card1 = &card.Card{
		Numbers: map[utils.Co]*card.Number{
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
	}
	var card2 = &card.Card{
		Numbers: map[utils.Co]*card.Number{
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
	}
	tests := []struct {
		name    string
		input   []string
		want    *Game
		wantErr bool
	}{
		{
			name: "returns an error if parseNums returns an error",
			input: []string{
				"1,24,two,7,12",
				"",
				"2 3 4 5",
				"1 8 7 6",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns a parsed input, advent of code example",
			input: []string{
				"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
				"",
				"22 13 17 11  0",
				"8  2 23  4 24",
				"21  9 14 16  7",
				"6 10  3 18  5",
				"1 12 20 15 19",
				"",
				"3 15  0  2 22",
				"9 18 13 17  5",
				"19  8  7 25 23",
				"20 11 10 24  4",
				"14 21 16 12  6",
				"",
				"14 21 17 24  4",
				"10 16 15  9 19",
				"18  8 23 26 20",
				"22 11 13  6  5",
				"2  0 12  3  7",
			},
			want: &Game{
				CardsNotWon: map[*card.Card]struct{}{
					card0: {},
					card1: {},
					card2: {},
				},
				Cards: []*card.Card{
					card0, card1, card2,
				},
				Nums: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInput(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if !reflect.DeepEqual(got.Nums, tt.want.Nums) {
					t.Errorf("ParseInput().Nums = %v, want %v", got.Nums, tt.want.Nums)
				}
				for i, card := range tt.want.Cards {
					for co, n := range card.Numbers {
						if !reflect.DeepEqual(got.Cards[i].Numbers[co], n) {
							t.Errorf("ParseInput().Cards[%d].Numbers[%v] = %v, want %v", i, co, got.Cards[i].Numbers[co], n)
						}
					}
				}
			}
		})
	}
}

func TestGame_PlayGame(t *testing.T) {
	var card00 = &card.Card{
		Numbers: map[utils.Co]*card.Number{
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
	}
	var card01 = &card.Card{
		Numbers: map[utils.Co]*card.Number{
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
	}
	var card02 = &card.Card{
		Numbers: map[utils.Co]*card.Number{
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
	}
	var card10 = &card.Card{
		Numbers: map[utils.Co]*card.Number{
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
	}
	var card11 = &card.Card{
		Numbers: map[utils.Co]*card.Number{
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
	}
	var card12 = &card.Card{
		Numbers: map[utils.Co]*card.Number{
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
	}
	var card20 = &card.Card{
		Numbers: map[utils.Co]*card.Number{
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
	}
	var card21 = &card.Card{
		Numbers: map[utils.Co]*card.Number{
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
	}
	var card22 = &card.Card{
		Numbers: map[utils.Co]*card.Number{
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
	}
	type fields struct {
		CardsNotWon map[*card.Card]struct{}
		Cards       []*card.Card
		Nums        []int
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		want1   int
		wantErr bool
	}{
		{
			name: "returns an error if no games can be solved",
			fields: fields{
				CardsNotWon: map[*card.Card]struct{}{
					card00: {},
					card01: {},
					card02: {},
				},
				Cards: []*card.Card{card00, card01, card02},
				Nums:  []int{7, 4, 9, 5},
			},
			want:    -1,
			want1:   -1,
			wantErr: true,
		},
		{
			name: "returns an error if not all games can be won",
			fields: fields{
				CardsNotWon: map[*card.Card]struct{}{
					card10: {},
					card11: {},
					card12: {},
				},
				Cards: []*card.Card{card10, card11, card12},
				Nums:  []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10},
			},
			want:    4512,
			want1:   -1,
			wantErr: true,
		},
		{
			name: "returns both part solutions for a game",
			fields: fields{
				CardsNotWon: map[*card.Card]struct{}{
					card20: {},
					card21: {},
					card22: {},
				},
				Cards: []*card.Card{card20, card21, card22},
				Nums:  []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
			},
			want:    4512,
			want1:   1924,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				CardsNotWon: tt.fields.CardsNotWon,
				Cards:       tt.fields.Cards,
				Nums:        tt.fields.Nums,
			}
			got, got1, err := g.PlayGame()
			if (err != nil) != tt.wantErr {
				t.Errorf("Game.PlayGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Game.PlayGame() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Game.PlayGame() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
