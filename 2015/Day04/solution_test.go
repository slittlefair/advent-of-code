package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hashIsValidPart1(t *testing.T) {
	tests := []struct {
		name string
		hash string
		want bool
	}{
		{
			name: "returns true if the hash starts with five zeroes",
			hash: "00000abcd",
			want: true,
		},
		{
			name: "returns true if the hash starts with five zeroes, advent of code example 1",
			hash: "000001dbbfa",
			want: true,
		},
		{
			name: "returns true if the hash starts with five zeroes, advent of code example 2",
			hash: "000006136ef",
			want: true,
		},
		{
			name: "returns false if the hash doesn't start with five zeroes",
			hash: "0000abcd",
			want: false,
		},
		{
			name: "returns false if the hash doesn't start with any zeroes",
			hash: "124abcd",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hashIsValidPart1(tt.hash)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_hashIsValidPart2(t *testing.T) {
	tests := []struct {
		name string
		hash string
		want bool
	}{
		{
			name: "returns true if the hash starts with six zeroes",
			hash: "000000abcd",
			want: true,
		},
		{
			name: "returns false if the hash doesn't start with six zeroes",
			hash: "00000abcd",
			want: false,
		},
		{
			name: "returns false if the hash doesn't start with any zeroes",
			hash: "124abcd",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hashIsValidPart2(tt.hash)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_findValidHash(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part1 bool
		want  int
	}{
		{
			name:  "returns solution to code abcdef, advent of code example 1, part 1",
			input: "abcdef",
			part1: true,
			want:  609043,
		},
		{
			name:  "returns solution to code abcdef, advent of code example 2, part 1",
			input: "pqrstuv",
			part1: true,
			want:  1048970,
		},
		{
			name:  "returns solution to code abcdef, advent of code example 1, part 2",
			input: "abcdef",
			part1: false,
			want:  6742839,
		},
		{
			name:  "returns solution to code abcdef, advent of code example 2, part 2",
			input: "pqrstuv",
			part1: false,
			want:  5714438,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findValidHash(tt.input, tt.part1)
			assert.Equal(t, tt.want, got)
		})
	}
}
