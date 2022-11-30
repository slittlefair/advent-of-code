package main

import (
	"crypto/md5"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hashContainsTriple(t *testing.T) {
	tests := []struct {
		name  string
		hash  string
		want  bool
		want1 byte
	}{
		{
			name: "returns 0 if the given hash does not contain a triple, advent of code example 1",
			hash: "abc0",
			want: false,
		},
		{
			name: "returns 0 if the given hash does not contain a triple, advent of code example 2",
			hash: "abc7",
			want: false,
		},
		{
			name:  "returns the triple character of the given hash if one exists, advent of code example 3",
			hash:  "abc18",
			want:  true,
			want1: '8',
		},
		{
			name:  "returns the triple character of the given hash if one exists, advent of code example 4",
			hash:  "abc39",
			want:  true,
			want1: 'e',
		},
		{
			name:  "returns the triple character of the given hash if one exists, advent of code example 5",
			hash:  "abc92",
			want:  true,
			want1: '9',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := hashContainsTriple(fmt.Sprintf("%x", md5.Sum([]byte(tt.hash))))
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_hashContainsQuintuple(t *testing.T) {
	type args struct {
		hash string
		s    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns false if the given hash does not contain a quintuple of the given character, advent of code example 1",
			args: args{
				hash: "abc19",
				s:    "8",
			},
			want: false,
		},
		{
			name: "returns false if the given hash does not contain a quintuple of the given character, advent of code example 2",
			args: args{
				hash: "abc815",
				s:    "e",
			},
			want: false,
		},
		{
			name: "returns false if the given hash does not contain a quintuple of the given character, advent of code example 3",
			args: args{
				hash: "abc1018",
				s:    "8",
			},
			want: false,
		},
		{
			name: "returns true if the given hash does contain a quintuple of the given character, advent of code example 4",
			args: args{
				hash: "abc816",
				s:    "e",
			},
			want: true,
		},
		{
			name: "returns true if the given hash does contain a quintuple of the given character, advent of code example 5",
			args: args{
				hash: "abc200",
				s:    "9",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hashContainsQuintuple(fmt.Sprintf("%x", md5.Sum([]byte(tt.args.hash))), tt.args.s[0])
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_findKeys(t *testing.T) {
	type args struct {
		salt  string
		part2 bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "returns 64th key for part 1, advent of code example",
			args: args{
				salt:  "abc",
				part2: false,
			},
			want: 22728,
		},
		{
			name: "returns 64th key for part 2, advent of code example",
			args: args{
				salt:  "abc",
				part2: true,
			},
			want: 22551,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findKeys(tt.args.salt, tt.args.part2)
			assert.Equal(t, tt.want, got)
		})
	}
}
