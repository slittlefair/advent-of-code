package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getVal(t *testing.T) {
	type args struct {
		val  int
		mask string
	}
	tests := []struct {
		name               string
		args               args
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if the mask can't be converted into a binary",
			args: args{
				val:  0,
				mask: "aXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			},
			want:               0,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "advent of code example 1",
			args: args{
				val:  11,
				mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			},
			want:               73,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "advent of code example 2",
			args: args{
				val:  101,
				mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			},
			want:               101,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "advent of code example 3",
			args: args{
				val:  0,
				mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			},
			want:               64,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getVal(tt.args.val, tt.args.mask)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_findSolutions(t *testing.T) {
	tests := []struct {
		name               string
		entries            []string
		want               int
		want1              int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if a mem line address can't be converted to an int",
			entries: []string{
				"mask = 000000000000000000000000000001XXXX0X",
				"mem[8] = 11",
				"mem[aaa] = 101",
				"mem[8] = 0",
			},
			want:               0,
			want1:              0,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if a mem line value can't be converted to an int",
			entries: []string{
				"mask = 000000000000000000000000000001XXXX0X",
				"mem[8] = 11",
				"mem[7] = 101",
				"mem[8] = xyz",
			},
			want:               0,
			want1:              0,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if the mask can't be converted to an int",
			entries: []string{
				"mask = 0000000000y0000000000000000001XXXX0X",
				"mem[8] = 11",
				"mem[7] = 101",
				"mem[8] = 0",
			},
			want:               0,
			want1:              0,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "advent of code example 1",
			entries: []string{
				"mask = 000000000000000000000000000001XXXX0X",
				"mem[8] = 11",
				"mem[7] = 101",
				"mem[8] = 0",
			},
			want:               165,
			want1:              3232,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "advent of code example 2",
			entries: []string{
				"mask = 000000000000000000000000000000X1001X",
				"mem[42] = 100",
				"mask = 00000000000000000000000000000000X0XX",
				"mem[26] = 1",
			},
			want:               51,
			want1:              208,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := findSolutions(tt.entries)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
