package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindStartMarker(t *testing.T) {
	t.Run("returns an error if a marker cannot be found", func(t *testing.T) {
		got, err := findStartMarker("abcabcabcabc", 4)
		assert.Error(t, err)
		assert.Equal(t, -1, got)
	})

	tests := []struct {
		input string
		n     int
		want  int
	}{
		{
			input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			n:     4,
			want:  7,
		},
		{
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			n:     4,
			want:  5,
		},
		{
			input: "nppdvjthqldpwncqszvftbrmjlhg",
			n:     4,
			want:  6,
		},
		{
			input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			n:     4,
			want:  10,
		},
		{
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			n:     4,
			want:  11,
		},
		{
			input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			n:     14,
			want:  19,
		},
		{
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			n:     14,
			want:  23,
		},
		{
			input: "nppdvjthqldpwncqszvftbrmjlhg",
			n:     14,
			want:  23,
		},
		{
			input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			n:     14,
			want:  29,
		},
		{
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			n:     14,
			want:  26,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("returns correct start marker index for %d characters, advent of code example %d", tt.n, i+1), func(t *testing.T) {
			got, err := findStartMarker(tt.input, tt.n)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
