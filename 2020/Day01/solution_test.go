package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name               string
		entries            []int
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "advent of code example",
			entries:            []int{1721, 979, 366, 299, 675, 1456},
			want:               514579,
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "returns an error if there are no solutions",
			entries:            []int{123, 82, 1, 999999},
			want:               0,
			errorAssertionFunc: assert.Error,
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
		entries            []int
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "advent of code example",
			entries:            []int{1721, 979, 366, 299, 675, 1456},
			want:               241861950,
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "returns an error if there are no solutions",
			entries:            []int{1, 876, 2, 919191919, 231},
			want:               0,
			errorAssertionFunc: assert.Error,
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
