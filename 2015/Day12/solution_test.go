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
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{
			name:    "returns an error if the input is not valid json",
			input:   "[1,2,3]]",
			want:    -1,
			wantErr: true,
		},
		{
			name:    "counts number in simple slice",
			input:   "[1,2,3]",
			want:    6,
			wantErr: false,
		},
		{
			name:    "counts number in slice with red included in object",
			input:   `[1,{"c":"red","b":2},3]`,
			want:    4,
			wantErr: false,
		},
		{
			name:    "counts number in object with red",
			input:   `{"d":"red","e":[1,2,3,4],"f":5}`,
			want:    0,
			wantErr: false,
		},
		{
			name:    "counts number in slice with red included",
			input:   `[1,"red",5]`,
			want:    6,
			wantErr: false,
		},
		{
			name:    "counts number in embedded objects with red included",
			input:   `{"d": {"sss": {"a": [1, 2]}, "a": {"bb": ["green", 1, 2, 3, 4], "aaa": ["green", "red", {"xx": 99}], "t": "red"}}}`,
			want:    3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findNonRedNumbers(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("findNonRedNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findNonRedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
