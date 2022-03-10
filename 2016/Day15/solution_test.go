package main

import (
	"reflect"
	"testing"
)

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    allDiscs
		wantErr bool
	}{
		{
			name: "returns an error if a line has fewer than four int matches",
			input: []string{
				"Disc #1 has 17 positions; at time=0, it is at position 5.",
				"Disc #2 has 19 positions; at time=0, it is at position 8.",
				"Disc #c has 7 positions; at time=0, it is at position 1.",
				"Disc #4 has 13 positions; at time=0, it is at position 7.",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns an error if a line has fewer than four int matches",
			input: []string{
				"Disc #1 has 17 positions; at time=0, it is at position 5.",
				"Disc #2 has 19 positions; at time=0, it is at position 8.",
				"Disc #3 has 7 positions; at time=0, it is at position 1.",
				"Discs #4 and #5 have 13 positions; at time=0, they are at position 7.",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns an error if a line has fewer than four int matches",
			input: []string{
				"Disc #1 has 17 positions; at time=0, it is at position 5.",
				"Disc #2 has 19 positions; at time=0, it is at position 8.",
				"Disc #3 has 7 positions; at time=0, it is at position 1.",
				"Disc #4 has 13 positions; at time=0, it is at position 7.",
			},
			want: allDiscs{
				1: {
					id:           1,
					numPositions: 17,
					position:     5,
				},
				2: {
					id:           2,
					numPositions: 19,
					position:     8,
				},
				3: {
					id:           3,
					numPositions: 7,
					position:     1,
				},
				4: {
					id:           4,
					numPositions: 13,
					position:     7,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInput(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allDiscs_moveDiscs(t *testing.T) {
	tests := []struct {
		name string
		ad   allDiscs
		want allDiscs
	}{
		{
			name: "increases the position of all discs by 1",
			ad: allDiscs{
				1: {
					id:           1,
					numPositions: 17,
					position:     5,
				},
				2: {
					id:           2,
					numPositions: 19,
					position:     18,
				},
				3: {
					id:           3,
					numPositions: 7,
					position:     1,
				},
				4: {
					id:           4,
					numPositions: 13,
					position:     7,
				},
			},
			want: allDiscs{
				1: {
					id:           1,
					numPositions: 17,
					position:     6,
				},
				2: {
					id:           2,
					numPositions: 19,
					position:     0,
				},
				3: {
					id:           3,
					numPositions: 7,
					position:     2,
				},
				4: {
					id:           4,
					numPositions: 13,
					position:     8,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ad := tt.ad
			ad.moveDiscs()
			if !reflect.DeepEqual(ad, tt.want) {
				t.Errorf("allDiscs.moveDiscs() = %v, want %v", ad, tt.want)
			}
		})
	}
}

func Test_allDiscs_getCapsule(t *testing.T) {
	tests := []struct {
		name string
		ad   allDiscs
		want bool
	}{
		{
			name: "returns false if the capsule won't exit the machine",
			ad: allDiscs{
				1: {
					id:           1,
					numPositions: 17,
					position:     6,
				},
				2: {
					id:           2,
					numPositions: 19,
					position:     0,
				},
				3: {
					id:           3,
					numPositions: 7,
					position:     2,
				},
				4: {
					id:           4,
					numPositions: 13,
					position:     8,
				},
			},
			want: false,
		},
		{
			name: "returns true if the capsule will exit the machine",
			ad: allDiscs{
				1: {
					id:           1,
					numPositions: 17,
					position:     0,
				},
				2: {
					id:           2,
					numPositions: 19,
					position:     18,
				},
				3: {
					id:           3,
					numPositions: 6,
					position:     4,
				},
				4: {
					id:           4,
					numPositions: 2,
					position:     1,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ad.getCapsule(); got != tt.want {
				t.Errorf("allDiscs.getCapsule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allDiscs_findSuccessfulTime(t *testing.T) {
	tests := []struct {
		name string
		ad   allDiscs
		want int
	}{
		{
			name: "returns successful time for advent of code example 1",
			ad: allDiscs{
				1: {
					id:           1,
					numPositions: 5,
					position:     4,
				},
				2: {
					id:           2,
					numPositions: 2,
					position:     1,
				},
			},
			want: 5,
		},
		{
			name: "returns successful time for advent of code example 2",
			ad: allDiscs{
				1: {
					id:           1,
					numPositions: 5,
					position:     4,
				},
				2: {
					id:           2,
					numPositions: 2,
					position:     1,
				},
				3: {
					id:           3,
					numPositions: 11,
					position:     0,
				},
			},
			want: 85,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ad.findSuccessfulTime(); got != tt.want {
				t.Errorf("allDiscs.findSuccessfulTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSolutions(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    int
		want1   int
		wantErr bool
	}{
		{
			name: "returns an error if parseInput returns an error",
			input: []string{
				"Disc #1 has 17 positions; at time=0, it is at position 5.",
				"Disc #2 has 19 positions; at time=0, it is at position 8.",
				"Disc #c has 7 positions; at time=0, it is at position 1.",
				"Disc #4 has 13 positions; at time=0, it is at position 7.",
			},
			want:    -1,
			want1:   -1,
			wantErr: true,
		},
		{
			name: "returns solutions for parts 1 and 2, advent of code example",
			input: []string{
				"Disc #1 has 5 positions; at time=0, it is at position 4.",
				"Disc #2 has 2 positions; at time=0, it is at position 1.",
			},
			want:    5,
			want1:   85,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := findSolutions(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("findSolutions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findSolutions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findSolutions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
