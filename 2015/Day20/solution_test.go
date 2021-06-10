package main

import (
	"testing"
)

func Test_deliverPresentsPart1(t *testing.T) {
	tests := []struct {
		name    string
		arg     int
		want    int
		wantErr bool
	}{
		{
			name:    "returns an error if no solution is found",
			arg:     0,
			want:    -1,
			wantErr: true,
		},
		{
			name:    "returns lowest house that reaches the target, advent of code example",
			arg:     150,
			want:    8,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deliverPresentsPart1(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("deliverPresentsPart1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("deliverPresentsPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deliverPresentsPart2(t *testing.T) {
	tests := []struct {
		name    string
		arg     int
		want    int
		wantErr bool
	}{
		{
			name:    "returns an error if no solution is found",
			arg:     0,
			want:    -1,
			wantErr: true,
		},
		{
			name:    "returns lowest house that reaches the target, advent of code example",
			arg:     150,
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deliverPresentsPart2(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("deliverPresentsPart2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("deliverPresentsPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
