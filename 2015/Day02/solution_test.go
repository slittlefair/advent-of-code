package main

import (
	"testing"
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
			if got := paperForPresent(tt.dimensions); got != tt.want {
				t.Errorf("paperForPresent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_totalPaperForPresents(t *testing.T) {
	tests := []struct {
		name     string
		presents []string
		want     int
		want1    int
		wantErr  bool
	}{
		{
			name: "returns an error if a present has more than 3 dimensions",
			presents: []string{
				"1x2x3",
				"10x8x999",
				"4x3x2x1",
				"5x6x7",
			},
			want:    -1,
			want1:   -1,
			wantErr: true,
		},
		{
			name: "returns the sum of all paper and ribbon needed, advent of code example",
			presents: []string{
				"4x3x2",
				"1x1x10",
			},
			want:    101,
			want1:   48,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := totalPaperForPresents(tt.presents)
			if (err != nil) != tt.wantErr {
				t.Errorf("totalPaperForPresents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("totalPaperForPresents() = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("totalPaperForPresents() = %v, want %v", got1, tt.want1)
			}
		})
	}
}
