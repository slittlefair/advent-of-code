package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createLanternfish(t *testing.T) {
	t.Run("returns a parsed Lanternfish srray from input, advent of code example", func(t *testing.T) {
		got := createLanternfish([]int{3, 4, 3, 1, 2})
		assert.Equal(t, Lanternfish{0, 1, 1, 2, 1, 0, 0, 0, 0}, got)
	})
}

func TestLanternfish_iterate(t *testing.T) {
	tests := []struct {
		name string
		lf   Lanternfish
		want Lanternfish
	}{
		{
			name: "returns a correct new Lanternfish slice from a previous iteration, advent of code example 1",
			lf:   Lanternfish{0, 1, 1, 2, 1, 0, 0, 0, 0},
			want: Lanternfish{1, 1, 2, 1, 0, 0, 0, 0, 0},
		},
		{
			name: "returns a correct new Lanternfish slice from a previous iteration, advent of code example 2",
			lf:   Lanternfish{1, 1, 2, 1, 0, 0, 0, 0, 0},
			want: Lanternfish{1, 2, 1, 0, 0, 0, 1, 0, 1},
		},
		{
			name: "returns a correct new Lanternfish slice from a previous iteration, advent of code example 3",
			lf:   Lanternfish{2, 1, 0, 0, 0, 1, 1, 1, 0},
			want: Lanternfish{1, 0, 0, 0, 1, 1, 3, 0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.lf.iterate()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLanternfish_count(t *testing.T) {
	t.Run("returns sum of Lanternfish", func(t *testing.T) {
		lf := Lanternfish{1, 8, 2, 5, 19, 4, 3, 0, 7}
		got := lf.count()
		assert.Equal(t, 49, got)
	})
}

func TestLanternfish_findSolution(t *testing.T) {
	t.Run("correctly returns number of lanternfish after 80 and 256 iterations, advent of code example", func(t *testing.T) {
		lf := Lanternfish{0, 1, 1, 2, 1, 0, 0, 0, 0}
		got, got1 := lf.findSolution()
		assert.Equal(t, 5934, got)
		assert.Equal(t, 26984457539, got1)
	})
}
