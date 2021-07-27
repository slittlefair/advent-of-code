package main

import (
	"testing"
)

func Test_passwordComplete(t *testing.T) {
	tests := []struct {
		name     string
		password [8]string
		want     bool
	}{
		{
			name:     "returns false if one of the items in the password is empty string",
			password: [8]string{"a", "b", "c", "1", "2", "", "0", ""},
			want:     false,
		},
		{
			name:     "returns true if none of the items in the password is empty string",
			password: [8]string{"a", "b", "c", "1", "2", "3", "0", "d"},
			want:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := passwordComplete(tt.password); got != tt.want {
				t.Errorf("passwordComplete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_composePassword(t *testing.T) {
	tests := []struct {
		name     string
		password [8]string
		want     string
	}{
		{
			name:     "it constructs a string password from the given array",
			password: [8]string{"a", "b", "c", "1", "2", "3", "0", "d"},
			want:     "abc1230d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := composePassword(tt.password); got != tt.want {
				t.Errorf("composePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printPassword(t *testing.T) {
	type args struct {
		part     int
		password [8]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "it returns password print string with part, some blank",
			args: args{
				part:     1,
				password: [8]string{"a", "b", "c", "1", "2", "", "0", ""},
			},
			want: "\rPart 1: abc12_0_",
		},
		{
			name: "it returns password print string with part, complete",
			args: args{
				part:     2,
				password: [8]string{"a", "b", "c", "1", "2", "3", "0", "d"},
			},
			want: "\rPart 2: abc1230d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printPassword(tt.args.part, tt.args.password); got != tt.want {
				t.Errorf("printPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findEasyPassword(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want string
	}{
		{
			name: "finds and returns easy password, advent of code example 1",
			id:   "abc",
			want: "18f47a30",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findEasyPassword(tt.id); got != tt.want {
				t.Errorf("findEasyPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDifficultPassword(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want string
	}{
		{
			name: "finds and returns difficult password, advent of code example 1",
			id:   "abc",
			want: "05ace8e3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDifficultPassword(tt.id); got != tt.want {
				t.Errorf("findDifficultPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
