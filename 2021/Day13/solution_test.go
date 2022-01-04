package main

import (
	"reflect"
	"testing"
)

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    *Paper
		wantErr bool
	}{
		{
			name: "returns an error if input has no blank line",
			input: []string{
				"6,10",
				"9,4",
				"3,0",
				"fold along x=9",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns an error if a coordinate line does not have 2 values",
			input: []string{
				"6,10",
				"9,4",
				"3,0,8",
				"",
				"fold along x=9",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns an error if an instruction line does not have 3 regex matches",
			input: []string{
				"6,10",
				"9,4",
				"3,0",
				"",
				"fold along x=9",
				"fold along y",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns an error if an instruction line direction is not x or y",
			input: []string{
				"6,10",
				"9,4",
				"3,0",
				"",
				"fold along x=9",
				"fold along z=2",
				"fold along y=8",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns an input parsed into a Paper, advent of code example",
			input: []string{
				"6,10",
				"0,14",
				"9,10",
				"0,3",
				"10,4",
				"4,11",
				"6,0",
				"6,12",
				"4,1",
				"0,13",
				"10,12",
				"3,4",
				"3,0",
				"8,4",
				"1,10",
				"2,14",
				"8,10",
				"9,0",
				"",
				"fold along y=7",
				"fold along x=5",
			},
			want: &Paper{
				Dots: Dots{
					{X: 6, Y: 10}:  struct{}{},
					{X: 0, Y: 14}:  struct{}{},
					{X: 9, Y: 10}:  struct{}{},
					{X: 0, Y: 3}:   struct{}{},
					{X: 10, Y: 4}:  struct{}{},
					{X: 4, Y: 11}:  struct{}{},
					{X: 6, Y: 0}:   struct{}{},
					{X: 6, Y: 12}:  struct{}{},
					{X: 4, Y: 1}:   struct{}{},
					{X: 0, Y: 13}:  struct{}{},
					{X: 10, Y: 12}: struct{}{},
					{X: 3, Y: 4}:   struct{}{},
					{X: 3, Y: 0}:   struct{}{},
					{X: 8, Y: 4}:   struct{}{},
					{X: 1, Y: 10}:  struct{}{},
					{X: 2, Y: 14}:  struct{}{},
					{X: 8, Y: 10}:  struct{}{},
					{X: 9, Y: 0}:   struct{}{},
				},
				Instructions: []Instruction{
					{Dir: "y", Val: 7},
					{Dir: "x", Val: 5},
				},
				MaxX: 10,
				MaxY: 14,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInput(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaper_doFold(t *testing.T) {
	tests := []struct {
		name string
		p    *Paper
		idx  int
		want *Paper
	}{
		{
			name: "does a fold up",
			p: &Paper{
				Dots: Dots{
					{X: 6, Y: 10}:  struct{}{},
					{X: 0, Y: 14}:  struct{}{},
					{X: 9, Y: 10}:  struct{}{},
					{X: 0, Y: 3}:   struct{}{},
					{X: 10, Y: 4}:  struct{}{},
					{X: 4, Y: 11}:  struct{}{},
					{X: 6, Y: 0}:   struct{}{},
					{X: 6, Y: 12}:  struct{}{},
					{X: 4, Y: 1}:   struct{}{},
					{X: 0, Y: 13}:  struct{}{},
					{X: 10, Y: 12}: struct{}{},
					{X: 3, Y: 4}:   struct{}{},
					{X: 3, Y: 0}:   struct{}{},
					{X: 8, Y: 4}:   struct{}{},
					{X: 1, Y: 10}:  struct{}{},
					{X: 2, Y: 14}:  struct{}{},
					{X: 8, Y: 10}:  struct{}{},
					{X: 9, Y: 0}:   struct{}{},
				},
				Instructions: []Instruction{
					{Dir: "y", Val: 7},
					{Dir: "x", Val: 5},
				},
				MaxX: 10,
				MaxY: 14,
			},
			idx: 0,
			want: &Paper{
				Dots: Dots{
					{X: 6, Y: 4}:  struct{}{},
					{X: 0, Y: 0}:  struct{}{},
					{X: 9, Y: 4}:  struct{}{},
					{X: 0, Y: 3}:  struct{}{},
					{X: 10, Y: 4}: struct{}{},
					{X: 4, Y: 3}:  struct{}{},
					{X: 6, Y: 0}:  struct{}{},
					{X: 6, Y: 2}:  struct{}{},
					{X: 4, Y: 1}:  struct{}{},
					{X: 0, Y: 1}:  struct{}{},
					{X: 10, Y: 2}: struct{}{},
					{X: 3, Y: 4}:  struct{}{},
					{X: 3, Y: 0}:  struct{}{},
					{X: 8, Y: 4}:  struct{}{},
					{X: 1, Y: 4}:  struct{}{},
					{X: 2, Y: 0}:  struct{}{},
					{X: 8, Y: 4}:  struct{}{},
					{X: 9, Y: 0}:  struct{}{},
				},
				Instructions: []Instruction{
					{Dir: "y", Val: 7},
					{Dir: "x", Val: 5},
				},
				MaxX: 10,
				MaxY: 6,
			},
		},
		{
			name: "does a fold left",
			p: &Paper{
				Dots: Dots{
					{X: 6, Y: 4}:  struct{}{},
					{X: 0, Y: 0}:  struct{}{},
					{X: 9, Y: 4}:  struct{}{},
					{X: 0, Y: 3}:  struct{}{},
					{X: 10, Y: 4}: struct{}{},
					{X: 4, Y: 3}:  struct{}{},
					{X: 6, Y: 0}:  struct{}{},
					{X: 6, Y: 2}:  struct{}{},
					{X: 4, Y: 1}:  struct{}{},
					{X: 0, Y: 1}:  struct{}{},
					{X: 10, Y: 2}: struct{}{},
					{X: 3, Y: 4}:  struct{}{},
					{X: 3, Y: 0}:  struct{}{},
					{X: 8, Y: 4}:  struct{}{},
					{X: 1, Y: 4}:  struct{}{},
					{X: 2, Y: 0}:  struct{}{},
					{X: 8, Y: 4}:  struct{}{},
					{X: 9, Y: 0}:  struct{}{},
				},
				Instructions: []Instruction{
					{Dir: "y", Val: 7},
					{Dir: "x", Val: 5},
				},
				MaxX: 10,
				MaxY: 6,
			},
			idx: 1,
			want: &Paper{
				Dots: Dots{
					{X: 4, Y: 4}: struct{}{},
					{X: 0, Y: 0}: struct{}{},
					{X: 1, Y: 4}: struct{}{},
					{X: 0, Y: 3}: struct{}{},
					{X: 0, Y: 4}: struct{}{},
					{X: 4, Y: 3}: struct{}{},
					{X: 4, Y: 0}: struct{}{},
					{X: 4, Y: 2}: struct{}{},
					{X: 4, Y: 1}: struct{}{},
					{X: 0, Y: 1}: struct{}{},
					{X: 0, Y: 2}: struct{}{},
					{X: 3, Y: 4}: struct{}{},
					{X: 3, Y: 0}: struct{}{},
					{X: 2, Y: 4}: struct{}{},
					{X: 1, Y: 4}: struct{}{},
					{X: 2, Y: 0}: struct{}{},
					{X: 2, Y: 4}: struct{}{},
					{X: 1, Y: 0}: struct{}{},
				},
				Instructions: []Instruction{
					{Dir: "y", Val: 7},
					{Dir: "x", Val: 5},
				},
				MaxX: 4,
				MaxY: 6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.p
			p.doFold(tt.idx)
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Paper.doFold() = %v, want %v", p, tt.want)
			}
		})
	}
}

func Test_findSolutions(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    int
		want1   *Paper
		wantErr bool
	}{
		{
			name: "returns an error if parseInput returns an error",
			input: []string{
				"6,7",
				"8,9",
				"fold along x=8",
			},
			want:    -1,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "finds solutions for parts 1 and 2, advent of code example",
			input: []string{
				"6,10",
				"0,14",
				"9,10",
				"0,3",
				"10,4",
				"4,11",
				"6,0",
				"6,12",
				"4,1",
				"0,13",
				"10,12",
				"3,4",
				"3,0",
				"8,4",
				"1,10",
				"2,14",
				"8,10",
				"9,0",
				"",
				"fold along y=7",
				"fold along x=5",
			},
			want: 17,
			want1: &Paper{
				Dots: Dots{
					{X: 4, Y: 4}: struct{}{},
					{X: 0, Y: 0}: struct{}{},
					{X: 1, Y: 4}: struct{}{},
					{X: 0, Y: 3}: struct{}{},
					{X: 0, Y: 4}: struct{}{},
					{X: 4, Y: 3}: struct{}{},
					{X: 4, Y: 0}: struct{}{},
					{X: 4, Y: 2}: struct{}{},
					{X: 4, Y: 1}: struct{}{},
					{X: 0, Y: 1}: struct{}{},
					{X: 0, Y: 2}: struct{}{},
					{X: 3, Y: 4}: struct{}{},
					{X: 3, Y: 0}: struct{}{},
					{X: 2, Y: 4}: struct{}{},
					{X: 1, Y: 4}: struct{}{},
					{X: 2, Y: 0}: struct{}{},
					{X: 2, Y: 4}: struct{}{},
					{X: 1, Y: 0}: struct{}{},
				},
				Instructions: []Instruction{
					{Dir: "y", Val: 7},
					{Dir: "x", Val: 5},
				},
				MaxX: 4,
				MaxY: 6,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := findSolutions(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("findSolutions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findSolutions() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("findSolutions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
