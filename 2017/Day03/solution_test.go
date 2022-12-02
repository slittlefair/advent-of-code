package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLayer(t *testing.T) {
	tests := []struct {
		target int
		want   int
		want1  int
		want2  int
	}{
		{
			target: 1,
			want:   0,
			want1:  1,
			want2:  0,
		},
		{
			target: 12,
			want:   2,
			want1:  25,
			want2:  9,
		},
		{
			target: 23,
			want:   2,
			want1:  25,
			want2:  9,
		},
		{
			target: 49,
			want:   3,
			want1:  49,
			want2:  25,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("returns correct layer for target %d", tt.target), func(t *testing.T) {
			got, got1, got2 := getLayer(tt.target)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, tt.want1, got1)
			assert.Equal(t, tt.want2, got2)
		})
	}
}

func TestGetMiddlePointDistance(t *testing.T) {
	tests := []struct {
		target     int
		corner     int
		prevCorner int
		want       int
	}{
		{
			target:     1,
			corner:     1,
			prevCorner: 0,
			want:       0,
		},
		{
			target:     12,
			corner:     25,
			prevCorner: 9,
			want:       1,
		},
		{
			target:     23,
			corner:     25,
			prevCorner: 9,
			want:       0,
		},
		{
			target:     49,
			corner:     49,
			prevCorner: 25,
			want:       3,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("returns correct layer for target %d", tt.target), func(t *testing.T) {
			got := getMiddlePointDistance(tt.target, tt.corner, tt.prevCorner)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestGetSolution(t *testing.T) {
	tests := []struct {
		target int
		want   int
	}{
		{
			target: 1,
			want:   0,
		},
		{
			target: 12,
			want:   3,
		},
		{
			target: 23,
			want:   2,
		},
		{
			target: 1024,
			want:   31,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("returns correct solution for target %d, advent of code example", tt.target), func(t *testing.T) {
			got := getSolution(tt.target)
			assert.Equal(t, got, tt.want)
		})
	}
}
