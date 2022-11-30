package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutations(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		want [][]string
	}{
		{
			name: "returns all permutations of two element slice",
			arg: []string{
				"Alligator",
				"Broccoli",
			},
			want: [][]string{
				{"Alligator", "Broccoli"},
				{"Broccoli", "Alligator"},
			},
		},
		{
			name: "returns all permutations of three element slice",
			arg: []string{
				"Alligator",
				"Broccoli",
				"Calcium",
			},
			want: [][]string{
				{"Alligator", "Broccoli", "Calcium"},
				{"Alligator", "Calcium", "Broccoli"},
				{"Broccoli", "Calcium", "Alligator"},
				{"Broccoli", "Alligator", "Calcium"},
				{"Calcium", "Broccoli", "Alligator"},
				{"Calcium", "Alligator", "Broccoli"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Permutations(tt.arg)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestIntSlicesAreEqual(t *testing.T) {
	type args struct {
		slice1 []int
		slice2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns false if the given decks aren't of equal length",
			args: args{
				slice1: []int{1, 2, 3, 4},
				slice2: []int{1, 2},
			},
			want: false,
		},
		{
			name: "returns false if the given decks aren't equal",
			args: args{
				slice1: []int{1, 2, 3, 4, 6, 5, 7},
				slice2: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: false,
		},
		{
			name: "returns true if the given decks are equal",
			args: args{
				slice1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				slice2: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntSlicesAreEqual(tt.args.slice1, tt.args.slice2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		s []int
		i int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		{
			name: "it removes first element from given slice",
			args: args{
				s: []int{1, 2, 3, 4},
				i: 0,
			},
			want:  []int{2, 3, 4},
			want1: []int{1, 2, 3, 4},
		},
		{
			name: "it removes last element from given slice",
			args: args{
				s: []int{1, 2, 3, 4},
				i: 3,
			},
			want:  []int{1, 2, 3},
			want1: []int{1, 2, 3, 4},
		},
		{
			name: "it removes a middle element from given slice",
			args: args{
				s: []int{1, 2, 3, 4},
				i: 1,
			},
			want:  []int{1, 3, 4},
			want1: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Remove(tt.args.s, tt.args.i)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, tt.args.s)
		})
	}
}

func TestFindExtremities(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		want  int
		want1 int
	}{
		{
			name:  "returns max and min numbers from a slice of ints, low values",
			nums:  []int{3, 2, 5, 1, 3, 6, 7, 4, 3, 5, 6, 7},
			want:  1,
			want1: 7,
		},
		{
			name:  "returns max and min numbers from a slice of ints, include negatives",
			nums:  []int{3, -2, 5, 1, 0, 3, -6, 10, 4, 3, 5, 6, 7},
			want:  -6,
			want1: 10,
		},
		{
			name:  "returns max and min numbers from a slice of ints, high ranged values",
			nums:  []int{639, 261, 280, 7635, 38005, 72619, 9811, 375, 209, 3856, 1111, 11114, 5739},
			want:  209,
			want1: 72619,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindExtremities(tt.nums)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
