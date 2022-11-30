package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createGame(t *testing.T) {
	type args struct {
		input  string
		maxNum int
	}
	tests := []struct {
		name string
		args args
		want Game
	}{
		{
			name: "creates game, advent of code example",
			args: args{
				input:  "389125467",
				maxNum: 9,
			},
			want: Game{
				CurrentCup: 3,
				Max:        9,
				Cups:       []int{0, 2, 5, 8, 6, 4, 7, 3, 9, 1},
			},
		},
		{
			name: "creates larger game, snippet of advent of code part 2 example",
			args: args{
				input:  "389125467",
				maxNum: 20,
			},
			want: Game{
				CurrentCup: 3,
				Max:        20,
				Cups:       []int{0, 2, 5, 8, 6, 4, 7, 10, 9, 1, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createGame(tt.args.input, tt.args.maxNum)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGame_doMove(t *testing.T) {
	type fields struct {
		CurrentCup int
		Max        int
		Cups       []int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Game
	}{
		{
			name: "move 1, advent of code example",
			fields: fields{
				CurrentCup: 3,
				Max:        9,
				Cups:       []int{0, 2, 5, 8, 6, 4, 7, 3, 9, 1},
			},
			want: &Game{
				CurrentCup: 2,
				Max:        9,
				Cups:       []int{0, 5, 8, 2, 6, 4, 7, 3, 9, 1},
			},
		},
		{
			name: "loop round destination, advent of code example",
			fields: fields{
				CurrentCup: 1,
				Max:        9,
				Cups:       []int{0, 3, 5, 6, 1, 8, 7, 9, 4, 2},
			},
			want: &Game{
				CurrentCup: 9,
				Max:        9,
				Cups:       []int{0, 9, 5, 6, 1, 8, 7, 2, 4, 3},
			},
		},
		{
			name: "skip destination from picked up cups, advent of code example",
			fields: fields{
				CurrentCup: 2,
				Max:        9,
				Cups:       []int{0, 5, 8, 2, 6, 4, 7, 3, 9, 1},
			},
			want: &Game{
				CurrentCup: 5,
				Max:        9,
				Cups:       []int{0, 3, 5, 2, 6, 4, 7, 8, 9, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				CurrentCup: tt.fields.CurrentCup,
				Max:        tt.fields.Max,
				Cups:       tt.fields.Cups,
			}
			g.doMove()
			assert.Equal(t, tt.want, g)
		})
	}
}

func TestGame_playGame(t *testing.T) {
	type fields struct {
		CurrentCup int
		Max        int
		Cups       []int
	}
	tests := []struct {
		name   string
		fields fields
		rounds int
		want   *Game
	}{
		{
			name: "play one round advent of code example",
			fields: fields{
				CurrentCup: 3,
				Max:        9,
				Cups:       []int{0, 2, 5, 8, 6, 4, 7, 3, 9, 1},
			},
			rounds: 1,
			want: &Game{
				CurrentCup: 2,
				Max:        9,
				Cups:       []int{0, 5, 8, 2, 6, 4, 7, 3, 9, 1},
			},
		},
		{
			name: "play ten rounds advent of code example",
			fields: fields{
				CurrentCup: 3,
				Max:        9,
				Cups:       []int{0, 2, 5, 8, 6, 4, 7, 3, 9, 1},
			},
			rounds: 10,
			want: &Game{
				CurrentCup: 8,
				Max:        9,
				Cups:       []int{0, 9, 6, 7, 1, 8, 5, 4, 3, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				CurrentCup: tt.fields.CurrentCup,
				Max:        tt.fields.Max,
				Cups:       tt.fields.Cups,
			}
			g.playGame(tt.rounds)
			assert.Equal(t, tt.want, g)
		})
	}
}

func TestGame_getOrderString(t *testing.T) {
	tests := []struct {
		name string
		cups []int
		want string
	}{
		{
			name: "advent of code example, 10 rounds",
			cups: []int{0, 9, 6, 7, 1, 8, 5, 4, 3, 2},
			want: "92658374",
		},
		{
			name: "advent of code example, 100 rounds",
			cups: []int{0, 6, 9, 8, 5, 2, 7, 3, 4, 1},
			want: "67384529",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Game{
				Cups: tt.cups,
			}
			got := g.getOrderString()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGame_getProductOfLabels(t *testing.T) {
	tests := []struct {
		name string
		cups []int
		want int
	}{
		{
			name: "returns the product of the two cups to the right of 1, advent of code example 1",
			cups: []int{0, 9, 6, 7, 1, 8, 5, 4, 3, 2},
			want: 18,
		},
		{
			name: "returns the product of the two cups to the right of 1, advent of code example 2",
			cups: []int{0, 6, 9, 8, 5, 2, 7, 3, 4, 1},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Game{
				Cups: tt.cups,
			}
			got := g.getProductOfLabels()
			assert.Equal(t, tt.want, got)
		})
	}
}
