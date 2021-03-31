package main

import (
	"reflect"
	"testing"
)

var basicNumbers = Numbers{20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21, 22, 23, 24, 25}
var complexNumbers = Numbers{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}

func TestNumbers_cyclePrevNumbers(t *testing.T) {
	type args struct {
		preambleLength int
		i              int
	}
	tests := []struct {
		name  string
		n     Numbers
		args  args
		want  bool
		want1 int
	}{
		{
			name: "advent of code example 1",
			n:    append(basicNumbers, 26),
			args: args{
				preambleLength: 25,
				i:              25,
			},
			want:  false,
			want1: -1,
		},
		{
			name: "advent of code example 2",
			n:    append(basicNumbers, 49),
			args: args{
				preambleLength: 25,
				i:              25,
			},
			want:  false,
			want1: -1,
		},
		{
			name: "advent of code example 3",
			n:    append(basicNumbers, 100),
			args: args{
				preambleLength: 25,
				i:              25,
			},
			want:  true,
			want1: 100,
		},
		{
			name: "advent of code example 4",
			n:    append(basicNumbers, 50),
			args: args{
				preambleLength: 25,
				i:              25,
			},
			want:  true,
			want1: 50,
		},
		{
			name: "advent of code example 5",
			n:    append(basicNumbers, 45, 26),
			args: args{
				preambleLength: 25,
				i:              26,
			},
			want:  false,
			want1: -1,
		},
		{
			name: "advent of code example 6",
			n:    append(basicNumbers, 45, 65),
			args: args{
				preambleLength: 25,
				i:              26,
			},
			want:  true,
			want1: 65,
		},
		{
			name: "advent of code example 7",
			n:    append(basicNumbers, 45, 64),
			args: args{
				preambleLength: 25,
				i:              26,
			},
			want:  false,
			want1: -1,
		},
		{
			name: "advent of code example 8",
			n:    append(basicNumbers, 45, 66),
			args: args{
				preambleLength: 25,
				i:              26,
			},
			want:  false,
			want1: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.n.cyclePrevNumbers(tt.args.preambleLength, tt.args.i)
			if got != tt.want {
				t.Errorf("Numbers.cyclePrevNumbers() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Numbers.cyclePrevNumbers() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNumbers_part1(t *testing.T) {
	tests := []struct {
		name           string
		n              Numbers
		preambleLength int
		want           int
		wantErr        bool
	}{
		{
			name:           "advent of code example 1, returns an error",
			n:              append(basicNumbers, 26),
			preambleLength: 25,
			want:           -1,
			wantErr:        true,
		},
		{
			name:           "advent of code example 2, returns the invalid number",
			n:              append(basicNumbers, 100),
			preambleLength: 25,
			want:           100,
			wantErr:        false,
		},
		{
			name:           "advent of code example 3, returns the invalid number",
			n:              complexNumbers,
			preambleLength: 5,
			want:           127,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.part1(tt.preambleLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("Numbers.part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Numbers.part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumbers_getSumNumbers(t *testing.T) {
	tests := []struct {
		name     string
		n        Numbers
		part1Sol int
		want     []int
		wantErr  bool
	}{
		{
			name:     "returns an error if there is no solution",
			n:        complexNumbers,
			part1Sol: 99,
			want:     []int{},
			wantErr:  true,
		},
		{
			name:     "advent of code example 1",
			n:        complexNumbers,
			part1Sol: 127,
			want:     []int{15, 25, 47, 40},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.getSumNumbers(tt.part1Sol)
			if (err != nil) != tt.wantErr {
				t.Errorf("Numbers.getSumNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Numbers.getSumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumbers_part2(t *testing.T) {
	tests := []struct {
		name     string
		n        Numbers
		part1Sol int
		want     int
		wantErr  bool
	}{
		{
			name:     "returns an error if there is no solution",
			n:        complexNumbers,
			part1Sol: 99,
			want:     -1,
			wantErr:  true,
		},
		{
			name:     "advent of code example 1",
			n:        complexNumbers,
			part1Sol: 127,
			want:     62,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.part2(tt.part1Sol)
			if (err != nil) != tt.wantErr {
				t.Errorf("Numbers.part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Numbers.part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
