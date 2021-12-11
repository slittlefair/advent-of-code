package main

import (
	"reflect"
	"testing"
)

var adventOfCodeExampleInput = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func Test_allFrequencies_compileFrequencies(t *testing.T) {
	tests := []struct {
		name  string
		f     allFrequencies
		input []string
		want  allFrequencies
	}{
		{
			name:  "compiles ones and zero frequencies on one input from empty",
			f:     make(allFrequencies, 6),
			input: []string{"110010"},
			want: allFrequencies{
				{ones: 1},
				{ones: 1},
				{zeros: 1},
				{zeros: 1},
				{ones: 1},
				{zeros: 1},
			},
		},
		{
			name: "compiles ones and zero frequencies on one input from non empty",
			f: allFrequencies{
				{ones: 1},
				{ones: 1},
				{zeros: 1},
				{zeros: 1},
				{ones: 1},
				{zeros: 1},
			},
			input: []string{"101011"},
			want: allFrequencies{
				{ones: 2},
				{ones: 1, zeros: 1},
				{zeros: 1, ones: 1},
				{zeros: 2},
				{ones: 2},
				{zeros: 1, ones: 1},
			},
		},
		{
			name: "compiles ones and zero frequencies on multiple inputs",
			f:    make(allFrequencies, 4),
			input: []string{
				"1010",
				"1111",
				"1011",
				"0001",
				"0100",
				"1010",
				"1010",
				"1011",
				"0111",
			},
			want: allFrequencies{
				{ones: 6, zeros: 3},
				{ones: 3, zeros: 6},
				{zeros: 2, ones: 7},
				{zeros: 4, ones: 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.f
			f.compileFrequencies((tt.input))
			if !reflect.DeepEqual(f, tt.want) {
				t.Errorf("compileFrequencies() = %v, want %v", f, tt.want)
			}
		})
	}
}

func Test_allFrequencies_compileRates(t *testing.T) {
	tests := []struct {
		name  string
		f     allFrequencies
		want  string
		want1 string
	}{
		{
			name: "compiles eRate and gRates from allFrequencies",
			f: allFrequencies{
				{ones: 6, zeros: 3},
				{ones: 3, zeros: 6},
				{ones: 7, zeros: 2},
				{ones: 5, zeros: 4},
			},
			want:  "1011",
			want1: "0100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.f.compileRates()
			if got != tt.want {
				t.Errorf("allFrequencies.compileRates() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("allFrequencies.compileRates() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    int64
		wantErr bool
	}{
		{
			name:    "returns the correct eRate and gRate product for input, advent of code example",
			input:   adventOfCodeExampleInput,
			want:    198,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRatings(t *testing.T) {
	type args struct {
		input []string
		og    bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "returns an error if it cannot find a single valid value",
			args: args{
				input: []string{
					"1000",
					"0101",
					"0110",
					"1111",
					"1011",
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "returns a valid og rating, advent of code example",
			args: args{
				input: adventOfCodeExampleInput,
				og:    true,
			},
			want:    "10111",
			wantErr: false,
		},
		{
			name: "returns a valid c02s rating, advent of code example",
			args: args{
				input: adventOfCodeExampleInput,
				og:    false,
			},
			want:    "01010",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getRatings(tt.args.input, tt.args.og)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRatings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getRatings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    int64
		wantErr bool
	}{
		{
			name:    "returns an error if the ogRating cannot be obtained",
			input:   []string{"2"},
			want:    -1,
			wantErr: true,
		},
		{
			name:    "returns an error if the c02sRating cannot be obtained",
			input:   []string{"1", "2"},
			want:    -1,
			wantErr: true,
		},
		{
			name:    "returns product of o2Rating and c02sRating from input, advent of code example",
			input:   adventOfCodeExampleInput,
			want:    230,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
