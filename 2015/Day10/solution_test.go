package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lookAndSay(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		want []string
	}{
		{
			name: "advent of code example 1",
			arg:  []string{"1"},
			want: []string{"1", "1"},
		},
		{
			name: "advent of code example 2",
			arg:  []string{"1", "1"},
			want: []string{"2", "1"},
		},
		{
			name: "advent of code example 3",
			arg:  []string{"2", "1"},
			want: []string{"1", "2", "1", "1"},
		},
		{
			name: "advent of code example 4",
			arg:  []string{"1", "2", "1", "1"},
			want: []string{"1", "1", "1", "2", "2", "1"},
		},
		{
			name: "advent of code example 5",
			arg:  []string{"1", "1", "1", "2", "2", "1"},
			want: []string{"3", "1", "2", "2", "1", "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lookAndSay(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}
