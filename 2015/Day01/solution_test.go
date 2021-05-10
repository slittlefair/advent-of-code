package main

import (
	"testing"
)

func Test_getFloorFromInstructions(t *testing.T) {
	tests := []struct {
		name         string
		instructions string
		want         int
	}{
		{
			name:         "advent of code example 1",
			instructions: "(())",
			want:         0,
		},
		{
			name:         "advent of code example 2",
			instructions: "()()",
			want:         0,
		},
		{
			name:         "advent of code example 3",
			instructions: "(((",
			want:         3,
		},
		{
			name:         "advent of code example 4",
			instructions: "(()(()(",
			want:         3,
		},
		{
			name:         "advent of code example 5",
			instructions: "))(((((",
			want:         3,
		},
		{
			name:         "advent of code example 6",
			instructions: "())",
			want:         -1,
		},
		{
			name:         "advent of code example 7",
			instructions: "))(",
			want:         -1,
		},
		{
			name:         "advent of code example 8",
			instructions: ")))",
			want:         -3,
		},
		{
			name:         "advent of code example 9",
			instructions: ")())())",
			want:         -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFloorFromInstructions(tt.instructions); got != tt.want {
				t.Errorf("getFloorFromInstructions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFirstInstanceOfBasement(t *testing.T) {
	tests := []struct {
		name         string
		instructions string
		want         int
		wantErr      bool
	}{
		{
			name:         "returns an error if basement is never reached",
			instructions: "(((",
			want:         -1,
			wantErr:      true,
		},
		{
			name:         "advent of code example 1",
			instructions: ")",
			want:         1,
			wantErr:      false,
		},
		{
			name:         "advent of code example 2",
			instructions: "()())",
			want:         5,
			wantErr:      false,
		},
		{
			name:         "returns early when basement is reached",
			instructions: "()()()()()((()()))))(())((((",
			want:         19,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFirstInstanceOfBasement(tt.instructions)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFirstInstanceOfBasement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getFirstInstanceOfBasement() = %v, want %v", got, tt.want)
			}
		})
	}
}
