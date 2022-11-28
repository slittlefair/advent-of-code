package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_paperForPresent(t *testing.T) {
	tests := []struct {
		name       string
		dimensions []int
		want       int
	}{
		{
			name:       "advent of code example 1",
			dimensions: []int{2, 3, 4},
			want:       58,
		},
		{
			name:       "advent of code example 2",
			dimensions: []int{1, 1, 10},
			want:       43,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := paperForPresent(tt.dimensions)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_totalPaperForPresents(t *testing.T) {
	tests := []struct {
		name               string
		presents           []string
		want               int
		want1              int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if a present has more than 3 dimensions",
			presents: []string{
				"1x2x3",
				"10x8x999",
				"4x3x2x1",
				"5x6x7",
			},
			want:               -1,
			want1:              -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns the sum of all paper and ribbon needed, advent of code example",
			presents: []string{
				"4x3x2",
				"1x1x10",
			},
			want:               101,
			want1:              48,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := totalPaperForPresents(tt.presents)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
			tt.errorAssertionFunc(t, err)
		})
	}
}
