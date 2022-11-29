package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_transformNumber(t *testing.T) {
	type args struct {
		value         int
		subjectNumber int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "returns a transformed number, value 1, subjectNumber 7",
			args: args{
				value:         1,
				subjectNumber: 7,
			},
			want: 7,
		},
		{
			name: "returns a transformed number, value 1, subjectNumber 7",
			args: args{
				value:         823543,
				subjectNumber: 7,
			},
			want: 5764801,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := transformNumber(tt.args.value, tt.args.subjectNumber)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getLoopSize(t *testing.T) {
	tests := []struct {
		name string
		key  int
		want int
	}{
		{
			name: "advent of code example 1",
			key:  5764801,
			want: 8,
		},
		{
			name: "advent of code example 2",
			key:  17807724,
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getLoopSize(tt.key)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getEncryptionKey(t *testing.T) {
	type args struct {
		subjectNumber int
		loopSize      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "advent of code example 1",
			args: args{
				subjectNumber: 17807724,
				loopSize:      8,
			},
			want: 14897079,
		},
		{
			name: "advent of code example 2",
			args: args{
				subjectNumber: 5764801,
				loopSize:      11,
			},
			want: 14897079,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getEncryptionKey(tt.args.subjectNumber, tt.args.loopSize)
			assert.Equal(t, tt.want, got)
		})
	}
}
