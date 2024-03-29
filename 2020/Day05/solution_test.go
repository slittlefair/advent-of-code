package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_halfSeats(t *testing.T) {
	type args struct {
		dirs string
		min  int
		max  int
	}
	tests := []struct {
		name               string
		args               args
		want               string
		want1              int
		want2              int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if the first character of dirs is an invalid letter",
			args: args{
				dirs: "XLRLRFBBF",
				min:  0,
				max:  10,
			},
			want:               "",
			want1:              0,
			want2:              0,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns the correct variables if first character of dirs is 'F'",
			args: args{
				dirs: "FRBRRFR",
				min:  32,
				max:  63,
			},
			want:               "RBRRFR",
			want1:              32,
			want2:              47,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns the correct variables if first character of dirs is 'B'",
			args: args{
				dirs: "BRBRRFR",
				min:  0,
				max:  63,
			},
			want:               "RBRRFR",
			want1:              32,
			want2:              63,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns the correct variables if first character of dirs is 'L'",
			args: args{
				dirs: "FRBRRFR",
				min:  32,
				max:  63,
			},
			want:               "RBRRFR",
			want1:              32,
			want2:              47,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns the correct variables if first character of dirs is 'L'",
			args: args{
				dirs: "LRBRRFR",
				min:  4,
				max:  7,
			},
			want:               "RBRRFR",
			want1:              4,
			want2:              5,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns the correct variables if first character of dirs is 'R'",
			args: args{
				dirs: "RRBRRFR",
				min:  0,
				max:  7,
			},
			want:               "RBRRFR",
			want1:              4,
			want2:              7,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := halfSeats(tt.args.dirs, tt.args.min, tt.args.max)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
			assert.Equal(t, tt.want2, got2)
		})
	}
}

func Test_findMyID(t *testing.T) {
	type args struct {
		usedIDs   map[int]bool
		lowestID  int
		highestID int
	}
	tests := []struct {
		name               string
		args               args
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if all IDs between lowest and highest are taken",
			args: args{
				usedIDs: map[int]bool{
					456: true,
					457: true,
					460: true,
					458: true,
					459: true,
				},
				lowestID:  456,
				highestID: 460,
			},
			want:               0,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns the correct ID between lowest and highest IDs from the usedIDs map",
			args: args{
				usedIDs: map[int]bool{
					456: true,
					457: true,
					460: true,
					458: true,
				},
				lowestID:  456,
				highestID: 460,
			},
			want:               459,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findMyID(tt.args.usedIDs, tt.args.lowestID, tt.args.highestID)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getusedIDs(t *testing.T) {
	tests := []struct {
		name               string
		entries            []string
		want               int
		want1              int
		want2              map[int]bool
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if an entry contains an invalid character instead of F or B",
			entries: []string{
				"FBFBBFFRLR",
				"BFFFBYFRRR",
				"FFFBBBFRRR",
				"BBFFBBFRLL",
			},
			want:               0,
			want1:              0,
			want2:              nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if an entry contains an invalid character instead of R or L",
			entries: []string{
				"FBFBBFFRLR",
				"BFFFBBFRRR",
				"FFFBBBFRRR",
				"BBFFBBFXLL",
			},
			want:               0,
			want1:              0,
			want2:              nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name:    "advent of code example 1",
			entries: []string{"FBFBBFFRLR"},
			want:    357,
			want1:   357,
			want2: map[int]bool{
				357: true,
			},
			errorAssertionFunc: assert.NoError,
		},
		{
			name:    "advent of code example 2",
			entries: []string{"BFFFBBFRRR"},
			want:    567,
			want1:   567,
			want2: map[int]bool{
				567: true,
			},
			errorAssertionFunc: assert.NoError,
		},
		{
			name:    "advent of code example 3",
			entries: []string{"FFFBBBFRRR"},
			want:    119,
			want1:   119,
			want2: map[int]bool{
				119: true,
			},
			errorAssertionFunc: assert.NoError,
		},
		{
			name:    "advent of code example 4",
			entries: []string{"BBFFBBFRLL"},
			want:    820,
			want1:   820,
			want2: map[int]bool{
				820: true,
			},
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "advent of code examples combined",
			entries: []string{
				"FBFBBFFRLR",
				"BFFFBBFRRR",
				"FFFBBBFRRR",
				"BBFFBBFRLL",
			},
			want:  119,
			want1: 820,
			want2: map[int]bool{
				119: true,
				357: true,
				567: true,
				820: true,
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := getusedIDs(tt.entries)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
			assert.Equal(t, tt.want2, got2)
		})
	}
}
