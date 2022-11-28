package main

import (
	"Advent-of-Code/graph"
	"Advent-of-Code/maths"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setupSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  solution
	}{
		{
			name:  "creates a solution struct from given input",
			input: "helloworld",
			want: solution{
				currentRoom:        graph.Co{X: 0, Y: 0},
				input:              "helloworld",
				shortestPathLength: maths.Infinity,
				directions: directions{
					{letter: "U", dir: graph.Co{Y: -1}},
					{letter: "D", dir: graph.Co{Y: 1}},
					{letter: "L", dir: graph.Co{X: -1}},
					{letter: "R", dir: graph.Co{X: 1}},
				},
				floors: floors{
					{X: 0, Y: 0}: struct{}{},
					{X: 1, Y: 0}: struct{}{},
					{X: 2, Y: 0}: struct{}{},
					{X: 3, Y: 0}: struct{}{},
					{X: 0, Y: 1}: struct{}{},
					{X: 1, Y: 1}: struct{}{},
					{X: 2, Y: 1}: struct{}{},
					{X: 3, Y: 1}: struct{}{},
					{X: 0, Y: 2}: struct{}{},
					{X: 1, Y: 2}: struct{}{},
					{X: 2, Y: 2}: struct{}{},
					{X: 3, Y: 2}: struct{}{},
					{X: 0, Y: 3}: struct{}{},
					{X: 1, Y: 3}: struct{}{},
					{X: 2, Y: 3}: struct{}{},
					{X: 3, Y: 3}: struct{}{},
				},
				shortestPath: "",
				longestPath:  "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setupSolution(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solution_reachedEnd(t *testing.T) {
	tests := []struct {
		name  string
		s     *solution
		want  bool
		want1 *solution
	}{
		{
			name: "returns false if the X coordinates of the current room and target room don't match",
			s: &solution{
				currentPath:        "UDDLRLD",
				currentRoom:        graph.Co{X: 2, Y: 3},
				longestPath:        "UUD",
				shortestPath:       "DLLDLRLDLDUUDLRUDL",
				shortestPathLength: 34,
			},
			want: false,
			want1: &solution{
				currentPath:        "UDDLRLD",
				currentRoom:        graph.Co{X: 2, Y: 3},
				longestPath:        "UUD",
				shortestPath:       "DLLDLRLDLDUUDLRUDL",
				shortestPathLength: 34,
			},
		},
		{
			name: "returns false if the Y coordinates of the current room and target room don't match",
			s: &solution{
				currentPath:        "UDDLRLD",
				currentRoom:        graph.Co{X: 3, Y: 1},
				longestPath:        "UUD",
				shortestPath:       "DLLDLRLDLDUUDLRUDL",
				shortestPathLength: 34,
			},
			want: false,
			want1: &solution{
				currentPath:        "UDDLRLD",
				currentRoom:        graph.Co{X: 3, Y: 1},
				longestPath:        "UUD",
				shortestPath:       "DLLDLRLDLDUUDLRUDL",
				shortestPathLength: 34,
			},
		},
		{
			name: "returns true, and nothing else, if the current room and target room are the same but path is not longest or shortest",
			s: &solution{
				currentPath:        "UDDLRLD",
				currentRoom:        graph.Co{X: 3, Y: 3},
				shortestPath:       "UUD",
				longestPath:        "DLLDLRLDLDUUDLRUDL",
				shortestPathLength: 3,
			},
			want: true,
			want1: &solution{
				currentPath:        "UDDLRLD",
				currentRoom:        graph.Co{X: 3, Y: 3},
				shortestPath:       "UUD",
				longestPath:        "DLLDLRLDLDUUDLRUDL",
				shortestPathLength: 3,
			},
		},
		{
			name: "returns true if the current room and target room are the same and changes solution if current path is new shortest",
			s: &solution{
				currentPath:        "UDDLRLD",
				currentRoom:        graph.Co{X: 3, Y: 3},
				shortestPath:       "URLDURLDURLDDLRUURURDUD",
				longestPath:        "DLLDLRLDLDUUDLRUDL",
				shortestPathLength: 21,
			},
			want: true,
			want1: &solution{
				currentPath:        "UDDLRLD",
				currentRoom:        graph.Co{X: 3, Y: 3},
				shortestPath:       "UDDLRLD",
				longestPath:        "DLLDLRLDLDUUDLRUDL",
				shortestPathLength: 7,
			},
		},
		{
			name: "returns true if the current room and target room are the same and changes solution if current path is new longest",
			s: &solution{
				currentPath:        "UDDLRLD",
				currentRoom:        graph.Co{X: 3, Y: 3},
				shortestPath:       "UURL",
				longestPath:        "",
				shortestPathLength: 4,
			},
			want: true,
			want1: &solution{
				currentPath:        "UDDLRLD",
				currentRoom:        graph.Co{X: 3, Y: 3},
				shortestPath:       "UURL",
				longestPath:        "UDDLRLD",
				shortestPathLength: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.s
			got := s.reachedEnd()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, s)
		})
	}
}

func Test_solution_getLockStatus(t *testing.T) {
	type fields struct {
		input       string
		currentPath string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "returns the first 4 characters of the correct hash, advent of code example 1",
			fields: fields{
				input:       "hijkl",
				currentPath: "",
			},
			want: "ced9",
		},
		{
			name: "returns the first 4 characters of the correct hash, advent of code example 2",
			fields: fields{
				input:       "hijkl",
				currentPath: "D",
			},
			want: "f2bc",
		},
		{
			name: "returns the first 4 characters of the correct hash, advent of code example 3",
			fields: fields{
				input:       "hijkl",
				currentPath: "DR",
			},
			want: "5745",
		},
		{
			name: "returns the first 4 characters of the correct hash, advent of code example 4",
			fields: fields{
				input:       "hijkl",
				currentPath: "DU",
			},
			want: "528e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := solution{
				input:       tt.fields.input,
				currentPath: tt.fields.currentPath,
			}
			got := s.getLockStatus()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_solution_isValidRoom(t *testing.T) {
	type args struct {
		char string
		diff graph.Co
	}
	tests := []struct {
		name        string
		currentRoom graph.Co
		args        args
		want        bool
	}{
		{
			name:        "returns false if given char is not valid",
			currentRoom: graph.Co{X: 2, Y: 1},
			args: args{
				char: "a",
				diff: graph.Co{X: -1},
			},
			want: false,
		},
		{
			name:        "returns false if given char is valid but the proposed new room is not valid",
			currentRoom: graph.Co{X: 3, Y: 1},
			args: args{
				char: "b",
				diff: graph.Co{X: 1},
			},
			want: false,
		},
		{
			name:        "returns true if given char is valid and the proposed new room is valid",
			currentRoom: graph.Co{X: 3, Y: 1},
			args: args{
				char: "b",
				diff: graph.Co{X: -1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := solution{
				currentRoom: tt.currentRoom,
				floors: floors{
					{X: 0, Y: 0}: struct{}{},
					{X: 1, Y: 0}: struct{}{},
					{X: 2, Y: 0}: struct{}{},
					{X: 3, Y: 0}: struct{}{},
					{X: 0, Y: 1}: struct{}{},
					{X: 1, Y: 1}: struct{}{},
					{X: 2, Y: 1}: struct{}{},
					{X: 3, Y: 1}: struct{}{},
					{X: 0, Y: 2}: struct{}{},
					{X: 1, Y: 2}: struct{}{},
					{X: 2, Y: 2}: struct{}{},
					{X: 3, Y: 2}: struct{}{},
					{X: 0, Y: 3}: struct{}{},
					{X: 1, Y: 3}: struct{}{},
					{X: 2, Y: 3}: struct{}{},
					{X: 3, Y: 3}: struct{}{},
				},
			}
			got := s.isValidRoom(tt.args.char, tt.args.diff)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_findSolutions(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
		want1 int
	}{
		{
			name:  "returns correct solutions for parts 1 and 2, advent of code example 1",
			input: "ihgpwlah",
			want:  "DDRRRD",
			want1: 370,
		},
		{
			name:  "returns correct solutions for parts 1 and 2, advent of code example 2",
			input: "kglvqrro",
			want:  "DDUDRLRRUDRD",
			want1: 492,
		},
		{
			name:  "returns correct solutions for parts 1 and 2, advent of code example 3",
			input: "ulqzkmiv",
			want:  "DRURDRUDDLLDLUURRDULRLDUUDDDRR",
			want1: 830,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findSolutions(tt.input)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
