package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleBuses = Buses{
	Bus{
		id:     7,
		offset: 0,
	},
	Bus{
		id:     13,
		offset: 1,
	},
	Bus{
		id:     59,
		offset: 4,
	},
	Bus{
		id:     31,
		offset: 6,
	},
	Bus{
		id:     19,
		offset: 7,
	},
}

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name               string
		entries            []string
		want               int
		want1              Buses
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if the timestamp cannot be converted to an int",
			entries: []string{
				"12883vfh12",
				"1, x, 2, x, x, 3, x",
			},
			want:               0,
			want1:              nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "advent of code example 1",
			entries: []string{
				"939",
				"7,13,x,x,59,x,31,19",
			},
			want:               939,
			want1:              exampleBuses,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseInput(tt.entries)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestBuses_part1(t *testing.T) {
	t.Run("advent of code example 1", func(t *testing.T) {
		b := &exampleBuses
		got := b.part1(939)
		assert.Equal(t, 295, got)
	})
}

func TestBuses_part2(t *testing.T) {
	t.Run("advent of code example", func(t *testing.T) {
		b := &exampleBuses
		got := b.part2()
		assert.Equal(t, 1068781, got)
	})
}
