package main

import (
	"reflect"
	"regexp"
	"testing"
)

func TestPairList_Len(t *testing.T) {
	tests := []struct {
		name string
		p    PairList
		want int
	}{
		{
			name: "returns length of p",
			p: PairList{
				{Key: "a", Value: 1},
				{Key: "a", Value: 1},
				{Key: "a", Value: 1},
				{Key: "a", Value: 1},
				{Key: "a", Value: 1},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Len(); got != tt.want {
				t.Errorf("PairList.Len() = %v, want %v", got, tt.want)
			}
		})
	}
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
			if !reflect.DeepEqual(tt.p, tt.want) {
				t.Errorf("PairList.Swap() = %v, want %v", tt.p, tt.want)
			}
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
			if got := tt.p.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("PairList.Less() = %v, want %v", got, tt.want)
			}
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roomIsValid(tt.line, regexp.MustCompile(`[a-z]`)); got != tt.want {
				t.Errorf("roomIsValid() = %v, want %v", got, tt.want)
			}
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
			got, got1 := getValidRooms(tt.input, regexp.MustCompile(`\d+`))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getValidRooms() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getValidRooms() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_validateRooms(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    string
		wantErr bool
	}{
		{
			name:    `returns an error if no decrypted rooms contain the string "northpole"`,
			input:   []string{"some-random-string-1"},
			want:    "",
			wantErr: true,
		},
		{
			name: `returns room idof the first decrypted room to contain the string "northpole"`,
			input: []string{
				"some-random-string-1",
				"another-random-string-12",
				"northpole-26",
				"south-pole-25",
			},
			want:    "26",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateRooms(tt.input, regexp.MustCompile(`\d+`))
			if (err != nil) != tt.wantErr {
				t.Errorf("validateRooms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("validateRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
