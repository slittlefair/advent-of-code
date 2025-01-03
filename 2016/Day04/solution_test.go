package main

import (
	"Advent-of-Code/regex"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairList_Len(t *testing.T) {
	t.Run("returns length of p", func(t *testing.T) {
		p := PairList{
			{Key: "a", Value: 1},
			{Key: "a", Value: 1},
			{Key: "a", Value: 1},
			{Key: "a", Value: 1},
			{Key: "a", Value: 1},
		}
		got := p.Len()
		assert.Equal(t, 5, got)
	})
}

func TestPairList_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    PairList
		args args
		want PairList
	}{
		{
			name: "swaps elements at given indices",
			p: PairList{
				{Key: "a", Value: 1},
				{Key: "b", Value: 2},
				{Key: "c", Value: 3},
				{Key: "d", Value: 4},
				{Key: "e", Value: 5},
			},
			args: args{i: 1, j: 3},
			want: PairList{
				{Key: "a", Value: 1},
				{Key: "d", Value: 4},
				{Key: "c", Value: 3},
				{Key: "b", Value: 2},
				{Key: "e", Value: 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Swap(tt.args.i, tt.args.j)
			assert.Equal(t, tt.want, tt.p)
		})
	}
}

func TestPairList_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    PairList
		args args
		want bool
	}{
		{
			name: "returns true if both values are equal and the first's key is alpabetically before the second's",
			p: PairList{
				{Key: "a", Value: 1},
				{Key: "c", Value: 3},
				{Key: "b", Value: 3},
				{Key: "d", Value: 4},
				{Key: "e", Value: 5},
			},
			args: args{
				i: 2,
				j: 1,
			},
			want: true,
		},
		{
			name: "returns false if both values are equal and the second's key is alpabetically before the first's",
			p: PairList{
				{Key: "a", Value: 1},
				{Key: "c", Value: 3},
				{Key: "b", Value: 3},
				{Key: "d", Value: 4},
				{Key: "e", Value: 5},
			},
			args: args{
				i: 1,
				j: 2,
			},
			want: false,
		},
		{
			name: "returns true if the first's value is greater than the second's",
			p: PairList{
				{Key: "a", Value: 1},
				{Key: "c", Value: 3},
				{Key: "b", Value: 3},
				{Key: "d", Value: 4},
				{Key: "e", Value: 5},
			},
			args: args{
				i: 4,
				j: 2,
			},
			want: true,
		},
		{
			name: "returns false if the first's value is less than the second's",
			p: PairList{
				{Key: "a", Value: 1},
				{Key: "c", Value: 3},
				{Key: "b", Value: 3},
				{Key: "d", Value: 4},
				{Key: "e", Value: 5},
			},
			args: args{
				i: 0,
				j: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.Less(tt.args.i, tt.args.j)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_roomIsValid(t *testing.T) {
	tests := []struct {
		name string
		line string
		want bool
	}{
		{
			name: "returns true if the room is valid, advent of code example 1",
			line: "aaaaa-bbb-z-y-x-123[abxyz]",
			want: true,
		},
		{
			name: "returns true if the room is valid, advent of code example 2",
			line: "a-b-c-d-e-f-g-h-987[abcde]",
			want: true,
		},
		{
			name: "returns true if the room is valid, advent of code example 3",
			line: "not-a-real-room-404[oarel]",
			want: true,
		},
		{
			name: "returns false if the room is not valid, advent of code example 4",
			line: "totally-real-room-200[decoy]",
			want: false,
		},
	}
	re := regexp.MustCompile(`[a-z]`)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := roomIsValid(tt.line, re)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getValidRooms(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
		want1 int
	}{
		{
			name: "returns an array of valid rooms, minus checksum, and the sum of their ids",
			input: []string{
				"aaaaa-bbb-z-y-x-123[abxyz]",
				"a-b-c-d-e-f-g-h-987[abcde]",
				"not-a-real-room-404[oarel]",
				"totally-real-room-200[decoy]",
			},
			want: []string{
				"aaaaa-bbb-z-y-x-123",
				"a-b-c-d-e-f-g-h-987",
				"not-a-real-room-404",
			},
			want1: 1514,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getValidRooms(tt.input, regex.MatchNums)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_validateRooms(t *testing.T) {
	tests := []struct {
		name               string
		input              []string
		want               string
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               `returns an error if no decrypted rooms contain the string "northpole"`,
			input:              []string{"some-random-string-1"},
			want:               "",
			errorAssertionFunc: assert.Error,
		},
		{
			name: `returns room idof the first decrypted room to contain the string "northpole"`,
			input: []string{
				"some-random-string-1",
				"another-random-string-12",
				"northpole-26",
				"south-pole-25",
			},
			want:               "26",
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateRooms(tt.input, regex.MatchNums)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
