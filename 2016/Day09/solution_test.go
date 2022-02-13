package main

import "testing"

func Test_decompress(t *testing.T) {
	type args struct {
		s     string
		part1 bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "gives the correct length of a part 1 decompressed file, advent of code example 1",
			args: args{
				s:     "ADVENT",
				part1: true,
			},
			want: 6,
		},
		{
			name: "gives the correct length of a part 1 decompressed file, advent of code example 2",
			args: args{
				s:     "A(1x5)BC",
				part1: true,
			},
			want: 7,
		},
		{
			name: "gives the correct length of a part 1 decompressed file, advent of code example 3",
			args: args{
				s:     "(3x3)XYZ",
				part1: true,
			},
			want: 9,
		},
		{
			name: "gives the correct length of a part 1 decompressed file, advent of code example 4",
			args: args{
				s:     "A(2x2)BCD(2x2)EFG",
				part1: true,
			},
			want: 11,
		},
		{
			name: "gives the correct length of a part 1 decompressed file, advent of code example 5",
			args: args{
				s:     "(6x1)(1x3)A",
				part1: true,
			},
			want: 6,
		},
		{
			name: "gives the correct length of a part 1 decompressed file, advent of code example 6",
			args: args{
				s:     "X(8x2)(3x3)ABCY",
				part1: true,
			},
			want: 18,
		},
		{
			name: "gives the correct length of a part 2 decompressed file, advent of code example 1",
			args: args{
				s:     "(3x3)XYZ",
				part1: false,
			},
			want: 9,
		},
		{
			name: "gives the correct length of a part 2 decompressed file, advent of code example 2",
			args: args{
				s:     "X(8x2)(3x3)ABCY",
				part1: false,
			},
			want: 20,
		},
		{
			name: "gives the correct length of a part 2 decompressed file, advent of code example 3",
			args: args{
				s:     "(27x12)(20x12)(13x14)(7x10)(1x12)A",
				part1: false,
			},
			want: 241920,
		},
		{
			name: "gives the correct length of a part 2 decompressed file, advent of code example 4",
			args: args{
				s:     "(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN",
				part1: false,
			},
			want: 445,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decompress(tt.args.s, tt.args.part1); got != tt.want {
				t.Errorf("decompress() = %v, want %v", got, tt.want)
			}
		})
	}
}
