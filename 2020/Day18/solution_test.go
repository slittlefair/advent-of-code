package main

import (
	"testing"
)

func Test_evaluateCoreSum(t *testing.T) {
	tests := []struct {
		name string
		sum  string
		want string
	}{
		{
			name: "advent of code example 1",
			sum:  "1+2*3+4*5+6",
			want: "71",
		},
		{
			name: "advent of code example 2",
			sum:  "8*3+9+3*4*3",
			want: "432",
		},
		{
			name: "advent of code example 3",
			sum:  "5+6*8",
			want: "88",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluateCoreSum(tt.sum); got != tt.want {
				t.Errorf("evaluateCoreSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_evaluateCoreSumAdditionFirst(t *testing.T) {
	tests := []struct {
		name string
		sum  string
		want string
	}{
		{
			name: "advent of code example 1",
			sum:  "1+2*3+4*5+6",
			want: "231",
		},
		{
			name: "advent of code example 2",
			sum:  "8*3+9+3*4*3",
			want: "1440",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluateCoreSumAdditionFirst(tt.sum); got != tt.want {
				t.Errorf("evaluateCoreSumAdditionFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_evaluateSum(t *testing.T) {
	type args struct {
		sum  string
		part int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "advent of code part 1 example 1",
			args: args{
				sum:  "1+2*3+4*5+6",
				part: 1,
			},
			want: "71",
		},
		{
			name: "advent of code part 1 example 2",
			args: args{
				sum:  "1+(2*3)+(4*(5+6))",
				part: 1,
			},
			want: "51",
		},
		{
			name: "advent of code part 1 example 3",
			args: args{
				sum:  "2*3+(4*5)",
				part: 1,
			},
			want: "26",
		},
		{
			name: "advent of code part 1 example 4",
			args: args{
				sum:  "5+(8*3+9+3*4*3)",
				part: 1,
			},
			want: "437",
		},
		{
			name: "advent of code part 1 example 5",
			args: args{
				sum:  "5*9*(7*3*3+9*3+(8+6*4))",
				part: 1,
			},
			want: "12240",
		},
		{
			name: "advent of code part 1 example 6",
			args: args{
				sum:  "((2+4*9)*(6+9*8+6)+6)+2+4*2",
				part: 1,
			},
			want: "13632",
		},
		{
			name: "advent of code part 2 example 1",
			args: args{
				sum:  "1+2*3+4*5+6",
				part: 2,
			},
			want: "231",
		},
		{
			name: "advent of code part 2 example 2",
			args: args{
				sum:  "1+(2*3)+(4*(5+6))",
				part: 2,
			},
			want: "51",
		},
		{
			name: "advent of code part 2 example 3",
			args: args{
				sum:  "2*3+(4*5)",
				part: 2,
			},
			want: "46",
		},
		{
			name: "advent of code part 2 example 4",
			args: args{
				sum:  "5+(8*3+9+3*4*3)",
				part: 2,
			},
			want: "1445",
		},
		{
			name: "advent of code part 2 example 5",
			args: args{
				sum:  "5*9*(7*3*3+9*3+(8+6*4))",
				part: 2,
			},
			want: "669060",
		},
		{
			name: "advent of code part 2 example 6",
			args: args{
				sum:  "((2+4*9)*(6+9*8+6)+6)+2+4*2",
				part: 2,
			},
			want: "23340",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluateSum(tt.args.sum, tt.args.part); got != tt.want {
				t.Errorf("evaluateSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSolutions(t *testing.T) {
	tests := []struct {
		name  string
		sums  []string
		want  int
		want1 int
	}{
		{
			name: "advent of code example",
			sums: []string{
				"1 + 2 * 3 + 4 * 5 + 6",
				"1 + (2 * 3) + (4 * (5 + 6))",
				"2 * 3 + (4 * 5)",
				"5 + (8 * 3 + 9 + 3 * 4 * 3)",
				"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
				"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			},
			want:  26457,
			want1: 694173,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findSolutions(tt.sums)
			if got != tt.want {
				t.Errorf("findSolutions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findSolutions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
