package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenNumsSaid_parseInput(t *testing.T) {
	t.Run("advent of code example", func(t *testing.T) {
		wns := WhenNumsSaid{}
		want := WhenNumsSaid{
			0: NumsSaid{
				lastSaid: 1,
			},
			3: NumsSaid{
				lastSaid: 2,
			},
			6: NumsSaid{
				lastSaid: 3,
			},
		}
		got := wns.parseInput([]int{0, 3, 6})
		assert.Equal(t, 6, got)
		assert.Equal(t, want, wns)
	})
}

func TestWhenNumsSaid_playGame(t *testing.T) {
	tests := []struct {
		name          string
		wns           WhenNumsSaid
		startingIndex int
		want          int
		want1         int
	}{
		{
			name: "advent of code example 1",
			wns: WhenNumsSaid{
				0: NumsSaid{lastSaid: 1},
				3: NumsSaid{lastSaid: 2},
				6: NumsSaid{lastSaid: 3},
			},
			startingIndex: 4,
			want:          436,
			want1:         175594,
		},
		{
			name: "advent of code example 2",
			wns: WhenNumsSaid{
				1: NumsSaid{lastSaid: 1},
				3: NumsSaid{lastSaid: 2},
				2: NumsSaid{lastSaid: 3},
			},
			startingIndex: 4,
			want:          1,
			want1:         2578,
		},
		{
			name: "advent of code example 3",
			wns: WhenNumsSaid{
				2: NumsSaid{lastSaid: 1},
				1: NumsSaid{lastSaid: 2},
				3: NumsSaid{lastSaid: 3},
			},
			startingIndex: 4,
			want:          10,
			want1:         3544142,
		},
		{
			name: "advent of code example 4",
			wns: WhenNumsSaid{
				1: NumsSaid{lastSaid: 1},
				2: NumsSaid{lastSaid: 2},
				3: NumsSaid{lastSaid: 3},
			},
			startingIndex: 4,
			want:          27,
			want1:         261214,
		},
		{
			name: "advent of code example 5",
			wns: WhenNumsSaid{
				2: NumsSaid{lastSaid: 1},
				3: NumsSaid{lastSaid: 2},
				1: NumsSaid{lastSaid: 3},
			},
			startingIndex: 4,
			want:          78,
			want1:         6895259,
		},
		{
			name: "advent of code example 6",
			wns: WhenNumsSaid{
				3: NumsSaid{lastSaid: 1},
				2: NumsSaid{lastSaid: 2},
				1: NumsSaid{lastSaid: 3},
			},
			startingIndex: 4,
			want:          438,
			want1:         18,
		},
		{
			name: "advent of code example 7",
			wns: WhenNumsSaid{
				3: NumsSaid{lastSaid: 1},
				1: NumsSaid{lastSaid: 2},
				2: NumsSaid{lastSaid: 3},
			},
			startingIndex: 4,
			want:          1836,
			want1:         362,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.wns.playGame(tt.startingIndex)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
