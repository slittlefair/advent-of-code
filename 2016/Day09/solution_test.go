package main

import (
	"regexp"
	"testing"
)

var re = regexp.MustCompile(`\d+`)

// func Test_handleMarker(t *testing.T) {
// 	tests := []struct {
// 		s     string
// 		want  string
// 		want1 int
// 	}{
// 		{
// 			s:     "(1x5)BC",
// 			want:  "BBBBB",
// 			want1: 6,
// 		},
// 		{
// 			s:     "(3x3)XYZ",
// 			want:  "XYZXYZXYZ",
// 			want1: 8,
// 		},
// 		{
// 			s:     "(2x2)BCD(2x2)EFG",
// 			want:  "BCBC",
// 			want1: 7,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.s, func(t *testing.T) {
// 			got, got1 := handleMarker(tt.s, re)
// 			if got != tt.want {
// 				t.Errorf("handleMarker() got = %v, want %v", got, tt.want)
// 			}
// 			if got1 != tt.want1 {
// 				t.Errorf("handleMarker() got1 = %v, want %v", got1, tt.want1)
// 			}
// 		})
// 	}
// }

func Test_decompress(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s:    "ADVENT",
			want: "ADVENT",
		},
		{
			s:    "A(1x5)BC",
			want: "ABBBBBC",
		},
		{
			s:    "(3x3)XYZ",
			want: "XYZXYZXYZ",
		},
		{},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := decompress(tt.s, re); got != tt.want {
				t.Errorf("decompress() = %v, want %v", got, tt.want)
			}
		})
	}
}
