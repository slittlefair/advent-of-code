package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name            string
		arg             []string
		want            []int
		errorAssertFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if an input line can't be converted to an int",
			arg: []string{
				"1",
				"5",
				"2",
				"a",
				"22",
			},
			want:            nil,
			errorAssertFunc: assert.Error,
		},
		{
			name: "returns a sorted list of containers, highest to lowest",
			arg: []string{
				"1",
				"5",
				"2",
				"10",
				"22",
			},
			want:            []int{22, 10, 5, 2, 1},
			errorAssertFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInput(tt.arg)
			tt.errorAssertFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEggnogContainers_FindContainers(t *testing.T) {
	t.Run("updates EggnogContainer with permutations that sum to WantedTotal, advent of code example", func(t *testing.T) {
		ec := &EggnogContainers{
			WantedTotal: 25,
			Ways:        make(map[int]int),
		}
		want := &EggnogContainers{
			WantedTotal: 25,
			Ways: map[int]int{
				3: 1,
				2: 3,
			},
		}
		ec.FindContainers([]int{20, 15, 10, 5, 5}, 0, 0)
		assert.Equal(t, want, ec)
	})
}

func TestEggnogContainers_CountPermutations(t *testing.T) {
	type fields struct {
		WantedTotal int
		Ways        map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "returns 0 for an empty EggnogComntainers",
			want: 0,
		},
		{
			name: "returns sum of all values in Ways, advent of code example",
			fields: fields{
				Ways: map[int]int{
					3: 1,
					2: 3,
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := EggnogContainers{
				WantedTotal: tt.fields.WantedTotal,
				Ways:        tt.fields.Ways,
			}
			got := ec.CountPermutations()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEggnogContainers_CountSmallestContainersPermutations(t *testing.T) {
	type fields struct {
		WantedTotal int
		Ways        map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "returns 0 for an empty EggnogComntainers",
			want: 0,
		},
		{
			name: "returns number of values for the smallest key in Ways, advent of code example",
			fields: fields{
				Ways: map[int]int{
					3: 1,
					2: 3,
				},
			},
			want: 3,
		},
		{
			name: "returns number of values for the smallest key in Ways, more complex example",
			fields: fields{
				Ways: map[int]int{
					6:     1,
					8:     3,
					12:    628,
					9:     12,
					3:     32,
					20:    7,
					88:    17,
					88888: 1,
					42527: 3673,
					7:     9182,
				},
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := EggnogContainers{
				WantedTotal: tt.fields.WantedTotal,
				Ways:        tt.fields.Ways,
			}
			got := ec.CountSmallestContainersPermutations()
			assert.Equal(t, tt.want, got)
		})
	}
}
