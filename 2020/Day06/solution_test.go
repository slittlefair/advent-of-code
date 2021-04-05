package main

import "testing"

func Test_getSolution(t *testing.T) {
	tests := []struct {
		name    string
		entries []string
		want    int
		want1   int
	}{
		{
			name: "advent of code example",
			entries: []string{
				"abc",
				"",
				"a",
				"b",
				"c",
				"",
				"ab",
				"ac",
				"",
				"a",
				"a",
				"a",
				"a",
				"",
				"b",
			},
			want:  11,
			want1: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getSolution(tt.entries)
			if got != tt.want {
				t.Errorf("getSolution() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getSolution() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
