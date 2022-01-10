package main

import (
	utils "Advent-of-Code/utils"
	"reflect"
	"testing"
)

var keypad map[utils.Co]string = map[utils.Co]string{
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
		currentCo utils.Co
		dir       string
		want      utils.Co
	}{
		{
			name:      "sets a new coordinate moving up",
			currentCo: utils.Co{X: 1, Y: 2},
			dir:       "U",
			want:      utils.Co{X: 1, Y: 1},
		},
		{
			name:      "does not set a new coordinate moving up if off the keyboard",
			currentCo: utils.Co{X: 1, Y: 0},
			dir:       "U",
			want:      utils.Co{X: 1, Y: 0},
		},
		{
			name:      "sets a new coordinate moving down",
			currentCo: utils.Co{X: 0, Y: 0},
			dir:       "D",
			want:      utils.Co{X: 0, Y: 1},
		},
		{
			name:      "does not set a new coordinate moving down if off the keyboard",
			currentCo: utils.Co{X: 1, Y: 2},
			dir:       "D",
			want:      utils.Co{X: 1, Y: 2},
		},
		{
			name:      "sets a new coordinate moving left",
			currentCo: utils.Co{X: 1, Y: 2},
			dir:       "L",
			want:      utils.Co{X: 0, Y: 2},
		},
		{
			name:      "does not set a new coordinate moving left if off the keyboard",
			currentCo: utils.Co{X: 0, Y: 0},
			dir:       "L",
			want:      utils.Co{X: 0, Y: 0},
		},
		{
			name:      "sets a new coordinate moving right",
			currentCo: utils.Co{X: 1, Y: 2},
			dir:       "R",
			want:      utils.Co{X: 2, Y: 2},
		},
		{
			name:      "does not set a new coordinate moving right if off the keyboard",
			currentCo: utils.Co{X: 2, Y: 0},
			dir:       "R",
			want:      utils.Co{X: 2, Y: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &CodeConstructor{
				currentCo: tt.currentCo,
			}
			cc.move(tt.dir, keypad)
			if !reflect.DeepEqual(cc.currentCo, tt.want) {
				t.Errorf("CodeConstructor.move().currentCo = %v, want %v", cc.currentCo, tt.want)
			}
		})
	}
}

func TestCodeConstructor_followDirections(t *testing.T) {
	type fields struct {
		currentCo utils.Co
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
				currentCo: utils.Co{X: 1, Y: 1},
				code:      []string{},
			},
			line: "ULL",
			want: &CodeConstructor{
				currentCo: utils.Co{X: 0, Y: 0},
				code:      []string{"1"},
			},
		},
		{
			name: "appends a number to the code after following instructions, advent of code example 2",
			fields: fields{
				currentCo: utils.Co{X: 0, Y: 0},
				code:      []string{"1"},
			},
			line: "RRDDD",
			want: &CodeConstructor{
				currentCo: utils.Co{X: 2, Y: 2},
				code:      []string{"1", "9"},
			},
		},
		{
			name: "appends a number to the code after following instructions, advent of code example 3",
			fields: fields{
				currentCo: utils.Co{X: 2, Y: 2},
				code:      []string{"1", "9"},
			},
			line: "LURDL",
			want: &CodeConstructor{
				currentCo: utils.Co{X: 1, Y: 2},
				code:      []string{"1", "9", "8"},
			},
		},
		{
			name: "appends a number to the code after following instructions, advent of code example 4",
			fields: fields{
				currentCo: utils.Co{X: 1, Y: 2},
				code:      []string{"1", "9", "8"},
			},
			line: "UUUUD",
			want: &CodeConstructor{
				currentCo: utils.Co{X: 1, Y: 1},
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
			if !reflect.DeepEqual(cc, tt.want) {
				t.Errorf("CodeConstructor.followDirections() = %v, want %v", cc, tt.want)
			}
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
			if got := cc.getCode(); got != tt.want {
				t.Errorf("CodeConstructor.getCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSolution(t *testing.T) {
	type args struct {
		input      []string
		keypad     map[utils.Co]string
		startingCo utils.Co
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
				startingCo: utils.Co{X: 1, Y: 1},
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
				keypad: map[utils.Co]string{
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
				startingCo: utils.Co{X: 0, Y: 2},
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
