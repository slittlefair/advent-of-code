package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		entries []int
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantErr   bool
		errString string
	}{
		{
			name: "advent of code example",
			args: args{
				entries: []int{1721, 979, 366, 299, 675, 1456},
			},
			want:    514579,
			wantErr: false,
		},
		{
			name: "returns an error if there are no solutions",
			args: args{
				entries: []int{123, 82, 1, 999999},
			},
			want:      0,
			wantErr:   true,
			errString: "part1: could not find entries that sum to 2020",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.args.entries)
			if (err != nil) != tt.wantErr {
				t.Errorf("part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errString != err.Error() {
				t.Errorf("part1() error = %v, want err %v", err, tt.errString)
				return
			}
			if got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		entries []int
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantErr   bool
		errString string
	}{
		{
			name: "advent of code example",
			args: args{
				entries: []int{1721, 979, 366, 299, 675, 1456},
			},
			want:    241861950,
			wantErr: false,
		},
		{
			name: "returns an error if there are no solutions",
			args: args{
				entries: []int{1, 876, 2, 919191919, 231},
			},
			want:      0,
			wantErr:   true,
			errString: "part2: could not find entries that sum to 2020",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.args.entries)
			if (err != nil) != tt.wantErr {
				t.Errorf("part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errString != err.Error() {
				t.Errorf("part2() error = %v, want err %v", err, tt.errString)
				return
			}
			if got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
