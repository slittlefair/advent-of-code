package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseDirection(t *testing.T) {
	tests := []struct {
		name               string
		entry              string
		want               string
		want1              int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "returns an error if the input can't be parsed into an int",
			entry:              "F2?3",
			want:               "F",
			want1:              0,
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "returns correctly parsed input",
			entry:              "F23",
			want:               "F",
			want1:              23,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseDirection(tt.entry)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name               string
		entries            []string
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if one of the entries can't be parsed successfully",
			entries: []string{
				"F10",
				"N3",
				"F7.",
				"R90",
				"F11",
			},
			want:               0,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "advent of code example 1",
			entries: []string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			want:               25,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.entries)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name               string
		entries            []string
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if one of the entries can't be parsed correctly",
			entries: []string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F!11",
			},
			want:               0,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "advent of code example 1",
			entries: []string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			want:               286,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.entries)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
