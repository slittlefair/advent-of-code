package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateDataStep(t *testing.T) {
	tests := []struct {
		name               string
		input              string
		want               string
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "returns an error if the input contains an invalid character",
			input:              "10010010301001001110110101110101",
			want:               "",
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "returns correct data for input, advent of code example 1",
			input:              "1",
			want:               "100",
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "returns correct data for input, advent of code example 2",
			input:              "0",
			want:               "001",
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "returns correct data for input, advent of code example 3",
			input:              "11111",
			want:               "11111000000",
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "returns correct data for input, advent of code example 4",
			input:              "111100001010",
			want:               "1111000010100101011110000",
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateDataStep(tt.input)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_generateData(t *testing.T) {
	type args struct {
		input          string
		requiredLength int
	}
	tests := []struct {
		name               string
		args               args
		want               string
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if generateDataStep returns an error",
			args: args{
				input:          "10010001111010100101010101010100101010101101101031010100101010100100101",
				requiredLength: 1000,
			},
			want:               "",
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns correct data for given input, advent of code example 1",
			args: args{
				input:          "10000",
				requiredLength: 20,
			},
			want:               "10000011110010000111",
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateData(tt.args.input, tt.args.requiredLength)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_calculateChecksum(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "returns correct checksum for given data, advent of code example 1",
			data: "110010110100",
			want: "100",
		},
		{
			name: "returns correct checksum for given data, advent of code example 2",
			data: "10000011110010000111",
			want: "01100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateChecksum(tt.data)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_findSolution(t *testing.T) {
	type args struct {
		input              string
		requiredDataLength int
	}
	tests := []struct {
		name               string
		args               args
		want               string
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if generateData returns an error",
			args: args{
				input:              "1000100010001111010110010119110100010001",
				requiredDataLength: 100,
			},
			want:               "",
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if generateData returns an error",
			args: args{
				input:              "10000",
				requiredDataLength: 20,
			},
			want:               "01100",
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findSolution(tt.args.input, tt.args.requiredDataLength)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_findSolutions(t *testing.T) {
	type args struct {
		input       string
		part1Length int
		part2Length int
	}
	tests := []struct {
		name               string
		args               args
		want               string
		want1              string
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if findSolution returns an error",
			args: args{
				input:       "10010011110000101101100010018101010101111100111",
				part1Length: 100,
				part2Length: 1000,
			},
			want:               "",
			want1:              "",
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns correct solutions to parts 1 and 2 (real input since AoC doesn't provide part2 example)",
			args: args{
				input:       "11101000110010100",
				part1Length: 272,
				part2Length: 35651584,
			},
			want:               "10100101010101101",
			want1:              "01100001101101001",
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := findSolutions(tt.args.input, tt.args.part1Length, tt.args.part2Length)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
