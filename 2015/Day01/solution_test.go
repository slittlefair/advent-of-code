package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			got := getFloorFromInstructions(tt.instructions)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getFirstInstanceOfBasement(t *testing.T) {
	tests := []struct {
		name               string
		instructions       string
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "returns an error if basement is never reached",
			instructions:       "(((",
			want:               -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "advent of code example 1",
			instructions:       ")",
			want:               1,
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "advent of code example 2",
			instructions:       "()())",
			want:               5,
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "returns early when basement is reached",
			instructions:       "()()()()()((()()))))(())((((",
			want:               19,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFirstInstanceOfBasement(tt.instructions)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
