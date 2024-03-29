package main

import (
	"Advent-of-Code/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

var keypad map[graph.Co]string = map[graph.Co]string{
	{X: 0, Y: 0}: "1",
	{X: 1, Y: 0}: "2",
	{X: 2, Y: 0}: "3",
	{X: 0, Y: 1}: "4",
	{X: 1, Y: 1}: "5",
	{X: 2, Y: 1}: "6",
	{X: 0, Y: 2}: "7",
	{X: 1, Y: 2}: "8",
	{X: 2, Y: 2}: "9",
}

func TestCodeConstructor_move(t *testing.T) {
	tests := []struct {
		name      string
		currentCo graph.Co
		dir       string
		want      graph.Co
	}{
		{
			name:      "sets a new coordinate moving up",
			currentCo: graph.Co{X: 1, Y: 2},
			dir:       "U",
			want:      graph.Co{X: 1, Y: 1},
		},
		{
			name:      "does not set a new coordinate moving up if off the keyboard",
			currentCo: graph.Co{X: 1, Y: 0},
			dir:       "U",
			want:      graph.Co{X: 1, Y: 0},
		},
		{
			name:      "sets a new coordinate moving down",
			currentCo: graph.Co{X: 0, Y: 0},
			dir:       "D",
			want:      graph.Co{X: 0, Y: 1},
		},
		{
			name:      "does not set a new coordinate moving down if off the keyboard",
			currentCo: graph.Co{X: 1, Y: 2},
			dir:       "D",
			want:      graph.Co{X: 1, Y: 2},
		},
		{
			name:      "sets a new coordinate moving left",
			currentCo: graph.Co{X: 1, Y: 2},
			dir:       "L",
			want:      graph.Co{X: 0, Y: 2},
		},
		{
			name:      "does not set a new coordinate moving left if off the keyboard",
			currentCo: graph.Co{X: 0, Y: 0},
			dir:       "L",
			want:      graph.Co{X: 0, Y: 0},
		},
		{
			name:      "sets a new coordinate moving right",
			currentCo: graph.Co{X: 1, Y: 2},
			dir:       "R",
			want:      graph.Co{X: 2, Y: 2},
		},
		{
			name:      "does not set a new coordinate moving right if off the keyboard",
			currentCo: graph.Co{X: 2, Y: 0},
			dir:       "R",
			want:      graph.Co{X: 2, Y: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &CodeConstructor{
				currentCo: tt.currentCo,
			}
			cc.move(tt.dir, keypad)
			assert.Equal(t, tt.want, cc.currentCo)
		})
	}
}

func TestCodeConstructor_followDirections(t *testing.T) {
	type fields struct {
		currentCo graph.Co
		code      []string
	}
	tests := []struct {
		name   string
		fields fields
		line   string
		want   *CodeConstructor
	}{
		{
			name: "appends a number to the code after following instructions, advent of code example 1",
			fields: fields{
				currentCo: graph.Co{X: 1, Y: 1},
				code:      []string{},
			},
			line: "ULL",
			want: &CodeConstructor{
				currentCo: graph.Co{X: 0, Y: 0},
				code:      []string{"1"},
			},
		},
		{
			name: "appends a number to the code after following instructions, advent of code example 2",
			fields: fields{
				currentCo: graph.Co{X: 0, Y: 0},
				code:      []string{"1"},
			},
			line: "RRDDD",
			want: &CodeConstructor{
				currentCo: graph.Co{X: 2, Y: 2},
				code:      []string{"1", "9"},
			},
		},
		{
			name: "appends a number to the code after following instructions, advent of code example 3",
			fields: fields{
				currentCo: graph.Co{X: 2, Y: 2},
				code:      []string{"1", "9"},
			},
			line: "LURDL",
			want: &CodeConstructor{
				currentCo: graph.Co{X: 1, Y: 2},
				code:      []string{"1", "9", "8"},
			},
		},
		{
			name: "appends a number to the code after following instructions, advent of code example 4",
			fields: fields{
				currentCo: graph.Co{X: 1, Y: 2},
				code:      []string{"1", "9", "8"},
			},
			line: "UUUUD",
			want: &CodeConstructor{
				currentCo: graph.Co{X: 1, Y: 1},
				code:      []string{"1", "9", "8", "5"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &CodeConstructor{
				currentCo: tt.fields.currentCo,
				code:      tt.fields.code,
			}
			cc.followDirections(tt.line, keypad)
			assert.Equal(t, tt.want, cc)
		})
	}
}

func TestCodeConstructor_getCode(t *testing.T) {
	tests := []struct {
		name string
		code []string
		want string
	}{
		{
			name: "returns a string code from the keys pressed, advent of code example 1",
			code: []string{"1", "9", "8", "5"},
			want: "1985",
		},
		{
			name: "returns a string code from the keys pressed, advent of code example 2",
			code: []string{"5", "D", "B", "3"},
			want: "5DB3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := CodeConstructor{
				code: tt.code,
			}
			got := cc.getCode()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getSolution(t *testing.T) {
	type args struct {
		input      []string
		keypad     map[graph.Co]string
		startingCo graph.Co
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "it gets the correct solution, advent of code example 1",
			args: args{
				input: []string{
					"ULL",
					"RRDDD",
					"LURDL",
					"UUUUD",
				},
				keypad:     keypad,
				startingCo: graph.Co{X: 1, Y: 1},
			},
			want: "1985",
		},
		{
			name: "it gets the correct solution, advent of code example 2",
			args: args{
				input: []string{
					"ULL",
					"RRDDD",
					"LURDL",
					"UUUUD",
				},
				keypad: map[graph.Co]string{
					{X: 2, Y: 0}: "1",
					{X: 1, Y: 1}: "2",
					{X: 2, Y: 1}: "3",
					{X: 3, Y: 1}: "4",
					{X: 0, Y: 2}: "5",
					{X: 1, Y: 2}: "6",
					{X: 2, Y: 2}: "7",
					{X: 3, Y: 2}: "8",
					{X: 4, Y: 2}: "9",
					{X: 1, Y: 3}: "A",
					{X: 2, Y: 3}: "B",
					{X: 3, Y: 3}: "C",
					{X: 2, Y: 4}: "D",
				},
				startingCo: graph.Co{X: 0, Y: 2},
			},
			want: "5DB3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSolution(tt.args.input, tt.args.keypad, tt.args.startingCo); got != tt.want {
				t.Errorf("getSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
