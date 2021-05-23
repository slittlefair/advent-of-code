package main

import (
	"testing"
)

func Test_countNumbers(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{
			name: "counts number in simple slice",
			arg:  "[1,2,3]",
			want: 6,
		},
		{
			name: "counts number in simple object",
			arg:  `{"a":2,"b":4}`,
			want: 6,
		},
		{
			name: "counts number in embedded slice",
			arg:  "[[[3]]]",
			want: 3,
		},
		{
			name: "counts number in embedded object",
			arg:  `{"a":{"b":4},"c":-1}`,
			want: 3,
		},
		{
			name: "counts number in object with slices",
			arg:  `{"a":[-1,1]}`,
			want: 0,
		},
		{
			name: "counts number in slice with objects",
			arg:  `[-1,{"a":1}]`,
			want: 0,
		},
		{
			name: "counts number in empty slice",
			arg:  "[]",
			want: 0,
		},
		{
			name: "counts number in empty object",
			arg:  "{}",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNumbers(tt.arg); got != tt.want {
				t.Errorf("countNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findNonRedNumbers(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{
			name: "counts number in simple slice",
			arg:  "[1,2,3]",
			want: 6,
		},
		{
			name: "counts number in slice with red included in object",
			arg:  `[1,{"c":"red","b":2},3]`,
			want: 4,
		},
		{
			name: "counts number in object with red",
			arg:  `{"d":"red","e":[1,2,3,4],"f":5}`,
			want: 0,
		},
		{
			name: "counts number in slice with red included",
			arg:  `[1,"red",5]`,
			want: 6,
		},
		{
			name: "counts number in embedded objects with red included",
			arg:  `{"d": {"sss": {"a": [1, 2]}, "a": {"bb": ["green", 1, 2, 3, 4], "aaa": ["green", "red", {"xx": 99}], "t": "red"}}}`,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNonRedNumbers(tt.arg); got != tt.want {
				t.Errorf("findNonRedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
