package main

import (
	helpers "Advent-of-Code"
	"reflect"
	"testing"
)

func TestPosition_turnLeft(t *testing.T) {
	type fields struct {
		direction int
		location  helpers.Co
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
				seen:      map[helpers.Co]bool{{X: 0, Y: 0}: true},
			}
			p.turnLeft()
			if p.direction != tt.want {
				t.Errorf("turnLeft() == %d, want %d", p.direction, tt.want)
			}
		})
	}
}

func TestPosition_turnRight(t *testing.T) {
	type fields struct {
		direction int
		location  helpers.Co
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
				seen:      map[helpers.Co]bool{{X: 0, Y: 0}: true},
			}
			p.turnRight()
			if p.direction != tt.want {
				t.Errorf("turnRight() == %d, want %d", p.direction, tt.want)
			}
		})
	}
}

func TestPosition_move(t *testing.T) {
	type fields struct {
		direction int
		location  helpers.Co
		seen      map[helpers.Co]bool
		hq        helpers.Co
	}
	tests := []struct {
		name   string
		fields fields
		want   helpers.Co
		want1  helpers.Co
	}{
		{
			name: "moves north",
			fields: fields{
				direction: 0,
				location:  helpers.Co{X: 6, Y: 9},
				seen:      make(map[helpers.Co]bool),
			},
			want: helpers.Co{X: 6, Y: 10},
		},
		{
			name: "moves east",
			fields: fields{
				direction: 1,
				location:  helpers.Co{X: -6, Y: 9},
				seen:      make(map[helpers.Co]bool),
			},
			want: helpers.Co{X: -5, Y: 9},
		},
		{
			name: "moves south",
			fields: fields{
				direction: 2,
				location:  helpers.Co{X: 6, Y: 9},
				seen:      make(map[helpers.Co]bool),
			},
			want: helpers.Co{X: 6, Y: 8},
		},
		{
			name: "moves west",
			fields: fields{
				direction: 3,
				location:  helpers.Co{X: 2, Y: 9},
				seen:      make(map[helpers.Co]bool),
			},
			want: helpers.Co{X: 1, Y: 9},
		},
		{
			name: "sets the hq to the location if it has been seen before",
			fields: fields{
				direction: 0,
				location:  helpers.Co{X: 0, Y: 0},
				seen: map[helpers.Co]bool{
					{X: 0, Y: 0}:    true,
					{X: 9, Y: 3}:    true,
					{X: -8, Y: -15}: true,
					{X: 0, Y: 1}:    true,
					{X: 0, Y: 2}:    true,
				},
			},
			want:  helpers.Co{X: 0, Y: 1},
			want1: helpers.Co{X: 0, Y: 1},
		},
		{
			name: "doesn't set the hq to the location if it has been set before",
			fields: fields{
				direction: 0,
				location:  helpers.Co{X: 0, Y: 0},
				seen: map[helpers.Co]bool{
					{X: 0, Y: 0}:    true,
					{X: 9, Y: 3}:    true,
					{X: -8, Y: -15}: true,
					{X: 0, Y: 1}:    true,
					{X: 0, Y: 3}:    true,
				},
				hq: helpers.Co{X: 4, Y: 5},
			},
			want:  helpers.Co{X: 0, Y: 1},
			want1: helpers.Co{X: 4, Y: 5},
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
			if p.location != tt.want {
				t.Errorf("move() = %v, want %v", p.location, tt.want)
			}
			if p.hq != tt.want1 {
				t.Errorf("move() = %v, want %v", p.hq, tt.want1)
			}
		})
	}
}

func TestPosition_followInstruction(t *testing.T) {
	type fields struct {
		direction int
		location  helpers.Co
		seen      map[helpers.Co]bool
	}
	tests := []struct {
		name    string
		fields  fields
		inst    string
		want    *Position
		wantErr bool
	}{
		{
			name:    "returns an error if there is more than one letter",
			inst:    "LR1",
			want:    &Position{},
			wantErr: true,
		},
		{
			name:    "returns an error if the letter is not L or R",
			inst:    "F1",
			want:    &Position{},
			wantErr: true,
		},
		{
			name: "returns an error if there is more than one number",
			inst: "2R1",
			want: &Position{
				direction: 1,
			},
			wantErr: true,
		},
		{
			name: "returns an error if there is more than one number",
			inst: "2R1",
			want: &Position{
				direction: 1,
			},
			wantErr: true,
		},
		{
			name: "follows an instruction turning left",
			fields: fields{
				direction: 2,
				location:  helpers.Co{X: 8, Y: 10},
				seen:      map[helpers.Co]bool{{X: 0, Y: 0}: true},
			},
			inst: "L2",
			want: &Position{
				direction: 1,
				location:  helpers.Co{X: 10, Y: 10},
				seen: map[helpers.Co]bool{
					{X: 0, Y: 0}:   true,
					{X: 9, Y: 10}:  true,
					{X: 10, Y: 10}: true,
				},
			},
			wantErr: false,
		},
		{
			name: "follows an instruction turning right",
			fields: fields{
				direction: 3,
				location:  helpers.Co{X: 0, Y: 0},
				seen:      map[helpers.Co]bool{{X: 0, Y: 0}: true},
			},
			inst: "R3",
			want: &Position{
				direction: 0,
				location:  helpers.Co{X: 0, Y: 3},
				seen: map[helpers.Co]bool{
					{X: 0, Y: 0}: true,
					{X: 0, Y: 1}: true,
					{X: 0, Y: 2}: true,
					{X: 0, Y: 3}: true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				direction: tt.fields.direction,
				location:  tt.fields.location,
				seen:      tt.fields.seen,
			}
			if err := p.followInstruction(tt.inst); (err != nil) != tt.wantErr {
				t.Errorf("Position.followInstruction() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Position.followInstruction() = %v, want %v", p, tt.want)
			}
		})
	}
}

func Test_followSteps(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    int
		wantErr bool
	}{
		{
			name:    "returns an error if an instruction is malformed",
			input:   []string{"L1, L2, R5, F5, L6"},
			want:    -1,
			wantErr: true,
		},
		{
			name:    "returns manhattan distance of instructions, advent of code example 1",
			input:   []string{"R2, L3"},
			want:    5,
			wantErr: false,
		},
		{
			name:    "returns manhattan distance of instructions, advent of code example 2",
			input:   []string{"R2, R2, R2"},
			want:    2,
			wantErr: false,
		},
		{
			name:    "returns manhattan distance of instructions, advent of code example 2",
			input:   []string{"R5, L5, R5, R3"},
			want:    12,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Position{
				seen: map[helpers.Co]bool{{X: 0, Y: 0}: true},
			}
			got, err := p.followSteps(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Position.followSteps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Position.followSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
