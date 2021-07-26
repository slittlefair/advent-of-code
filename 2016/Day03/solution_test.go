package main

import (
	"testing"
)

func Test_triangleIsValid(t *testing.T) {
	tests := []struct {
		name     string
		triangle []int
		want     bool
	}{
		{
			name:     "returns false if the 3rd length is less than the sum of the 1st and 2nd",
			triangle: []int{2, 2, 5},
			want:     false,
		},
		{
			name:     "returns false if the 3rd length is equal to the sum of the 1st and 2nd",
			triangle: []int{2, 5, 7},
			want:     false,
		},
		{
			name:     "returns false if the 2nd length is less than the sum of the 1st and 3rd",
			triangle: []int{2, 3, 10},
			want:     false,
		},
		{
			name:     "returns false if the 2nd length is equal to the sum of the 1st and 3rd",
			triangle: []int{2, 6, 4},
			want:     false,
		},
		{
			name:     "returns false if the 1st length is less than the sum of the 2nd and 3rd",
			triangle: []int{2, 6, 9},
			want:     false,
		},
		{
			name:     "returns false if the 1st length is equal to the sum of the 2nd and 3rd",
			triangle: []int{12, 6, 6},
			want:     false,
		},
		{
			name:     "returns true if all edges are greater than the sum of the other two",
			triangle: []int{2, 2, 2},
			want:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleIsValid(tt.triangle); got != tt.want {
				t.Errorf("triangleIsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkTriangles(t *testing.T) {
	tests := []struct {
		name      string
		triangles [][]int
		want      int
	}{
		{
			name:      "returns 0 if no triangles are supplied",
			triangles: [][]int{},
			want:      0,
		},
		{
			name: "returns 0 if no triangles are valid",
			triangles: [][]int{
				{2, 2, 5},
				{2, 6, 4},
				{2, 6, 9},
			},
			want: 0,
		},
		{
			name: "returns the length of the supplied triangles if all are valid",
			triangles: [][]int{
				{23, 41, 26},
				{2, 2, 2},
				{8, 7, 9},
			},
			want: 3,
		},
		{
			name: "returns the number of valid supplied triangles",
			triangles: [][]int{
				{23, 41, 26},
				{3, 4, 7},
				{8, 7, 9},
				{1, 2, 6},
				{6, 8, 1},
				{9, 10, 14},
				{32, 52, 40},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkTriangles(tt.triangles); got != tt.want {
				t.Errorf("checkTriangles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateHorizontalTriangles(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name: "returns number of valid horizontal traingles in input",
			input: []string{
				"2  4  3",
				"56  62  99",
				"54  33  109",
				"5  10  25",
				"6  6  7",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateHorizontalTriangles(tt.input); got != tt.want {
				t.Errorf("validateHorizontalTriangles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateVerticalTriangles(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name: "returns number of valid vertical traingles in input",
			input: []string{
				"2  4  11",
				"56  62  99",
				"54  33  109",
				"5  10  25",
				"6  6  7",
				"7  8  19",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateVerticalTriangles(tt.input); got != tt.want {
				t.Errorf("validateVerticalTriangles() = %v, want %v", got, tt.want)
			}
		})
	}
}
