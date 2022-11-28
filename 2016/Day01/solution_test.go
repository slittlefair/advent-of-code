package main

import (
	"Advent-of-Code/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPosition_turnLeft(t *testing.T) {
	type fields struct {
		direction int
		location  graph.Co
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "reduces the position by one",
			fields: fields{
				direction: 2,
			},
			want: 1,
		},
		{
			name: "sets the position to 3 if originally 0",
			fields: fields{
				direction: 0,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				direction: tt.fields.direction,
				location:  tt.fields.location,
				seen:      map[graph.Co]bool{{X: 0, Y: 0}: true},
			}
			p.turnLeft()
			assert.Equal(t, tt.want, p.direction)
		})
	}
}

func TestPosition_turnRight(t *testing.T) {
	type fields struct {
		direction int
		location  graph.Co
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "increases the position by one",
			fields: fields{
				direction: 2,
			},
			want: 3,
		},
		{
			name: "sets the position to 0 if originally 3",
			fields: fields{
				direction: 3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				direction: tt.fields.direction,
				location:  tt.fields.location,
				seen:      map[graph.Co]bool{{X: 0, Y: 0}: true},
			}
			p.turnRight()
			assert.Equal(t, tt.want, p.direction)
		})
	}
}

func TestPosition_move(t *testing.T) {
	type fields struct {
		direction int
		location  graph.Co
		seen      map[graph.Co]bool
		hq        graph.Co
	}
	tests := []struct {
		name   string
		fields fields
		want   graph.Co
		want1  graph.Co
	}{
		{
			name: "moves north",
			fields: fields{
				direction: 0,
				location:  graph.Co{X: 6, Y: 9},
				seen:      make(map[graph.Co]bool),
			},
			want: graph.Co{X: 6, Y: 10},
		},
		{
			name: "moves east",
			fields: fields{
				direction: 1,
				location:  graph.Co{X: -6, Y: 9},
				seen:      make(map[graph.Co]bool),
			},
			want: graph.Co{X: -5, Y: 9},
		},
		{
			name: "moves south",
			fields: fields{
				direction: 2,
				location:  graph.Co{X: 6, Y: 9},
				seen:      make(map[graph.Co]bool),
			},
			want: graph.Co{X: 6, Y: 8},
		},
		{
			name: "moves west",
			fields: fields{
				direction: 3,
				location:  graph.Co{X: 2, Y: 9},
				seen:      make(map[graph.Co]bool),
			},
			want: graph.Co{X: 1, Y: 9},
		},
		{
			name: "sets the hq to the location if it has been seen before",
			fields: fields{
				direction: 0,
				location:  graph.Co{X: 0, Y: 0},
				seen: map[graph.Co]bool{
					{X: 0, Y: 0}:    true,
					{X: 9, Y: 3}:    true,
					{X: -8, Y: -15}: true,
					{X: 0, Y: 1}:    true,
					{X: 0, Y: 2}:    true,
				},
			},
			want:  graph.Co{X: 0, Y: 1},
			want1: graph.Co{X: 0, Y: 1},
		},
		{
			name: "doesn't set the hq to the location if it has been set before",
			fields: fields{
				direction: 0,
				location:  graph.Co{X: 0, Y: 0},
				seen: map[graph.Co]bool{
					{X: 0, Y: 0}:    true,
					{X: 9, Y: 3}:    true,
					{X: -8, Y: -15}: true,
					{X: 0, Y: 1}:    true,
					{X: 0, Y: 3}:    true,
				},
				hq: graph.Co{X: 4, Y: 5},
			},
			want:  graph.Co{X: 0, Y: 1},
			want1: graph.Co{X: 4, Y: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				direction: tt.fields.direction,
				location:  tt.fields.location,
				seen:      tt.fields.seen,
				hq:        tt.fields.hq,
			}
			p.move()
			assert.Equal(t, tt.want, p.location)
			assert.Equal(t, tt.want1, p.hq)
		})
	}
}

func TestPosition_followInstruction(t *testing.T) {
	type fields struct {
		direction int
		location  graph.Co
		seen      map[graph.Co]bool
	}
	tests := []struct {
		name               string
		fields             fields
		inst               string
		want               *Position
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "returns an error if there is more than one letter",
			inst:               "LR1",
			want:               &Position{},
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "returns an error if the letter is not L or R",
			inst:               "F1",
			want:               &Position{},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if there is more than one number",
			inst: "2R1",
			want: &Position{
				direction: 1,
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if there is more than one number",
			inst: "2R1",
			want: &Position{
				direction: 1,
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "follows an instruction turning left",
			fields: fields{
				direction: 2,
				location:  graph.Co{X: 8, Y: 10},
				seen:      map[graph.Co]bool{{X: 0, Y: 0}: true},
			},
			inst: "L2",
			want: &Position{
				direction: 1,
				location:  graph.Co{X: 10, Y: 10},
				seen: map[graph.Co]bool{
					{X: 0, Y: 0}:   true,
					{X: 9, Y: 10}:  true,
					{X: 10, Y: 10}: true,
				},
			},
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "follows an instruction turning right",
			fields: fields{
				direction: 3,
				location:  graph.Co{X: 0, Y: 0},
				seen:      map[graph.Co]bool{{X: 0, Y: 0}: true},
			},
			inst: "R3",
			want: &Position{
				direction: 0,
				location:  graph.Co{X: 0, Y: 3},
				seen: map[graph.Co]bool{
					{X: 0, Y: 0}: true,
					{X: 0, Y: 1}: true,
					{X: 0, Y: 2}: true,
					{X: 0, Y: 3}: true,
				},
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				direction: tt.fields.direction,
				location:  tt.fields.location,
				seen:      tt.fields.seen,
			}
			err := p.followInstruction(tt.inst)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, p)
		})
	}
}

func Test_followSteps(t *testing.T) {
	tests := []struct {
		name               string
		input              []string
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "returns an error if an instruction is malformed",
			input:              []string{"L1, L2, R5, F5, L6"},
			want:               -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "returns manhattan distance of instructions, advent of code example 1",
			input:              []string{"R2, L3"},
			want:               5,
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "returns manhattan distance of instructions, advent of code example 2",
			input:              []string{"R2, R2, R2"},
			want:               2,
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "returns manhattan distance of instructions, advent of code example 2",
			input:              []string{"R5, L5, R5, R3"},
			want:               12,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Position{
				seen: map[graph.Co]bool{{X: 0, Y: 0}: true},
			}
			got, err := p.followSteps(tt.input)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
