package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_composeBounds(t *testing.T) {
	type args struct {
		input      []string
		upperBound int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		{
			name: "correctly composes upper and lower bounds",
			args: args{
				input: []string{
					"4-8",
					"16-18",
					"17-19",
					"0-2",
					"14-15",
					"5-9",
					"10-12",
				},
				upperBound: 20,
			},
			want:  []int{0, 4, 5, 10, 14, 16, 17, 21},
			want1: []int{2, 8, 9, 12, 15, 18, 19, 21},
		},
		{
			name: "correctly composes upper and lower bounds, advent of code example",
			args: args{
				input: []string{
					"5-8",
					"0-2",
					"4-7",
				},
				upperBound: 9,
			},
			want:  []int{0, 4, 5, 10},
			want1: []int{2, 7, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := composeBounds(tt.args.input, tt.args.upperBound)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_findAllowedIPs(t *testing.T) {
	type args struct {
		lowers []int
		uppers []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "correctly returns answers for parts 1 and 2",
			args: args{
				lowers: []int{0, 4, 5, 10, 14, 16, 17, 21},
				uppers: []int{2, 8, 9, 12, 15, 18, 19, 21},
			},
			want:  3,
			want1: 3,
		},
		{
			name: "correctly returns answers for parts 1 and 2, advent of code example",
			args: args{
				lowers: []int{0, 4, 5, 10},
				uppers: []int{2, 7, 8, 10},
			},
			want:  3,
			want1: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findAllowedIPs(tt.args.lowers, tt.args.uppers)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
