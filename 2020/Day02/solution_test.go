package main

import (
	"reflect"
	"testing"
)

var examplePasswordCollection = []passwords{
	{
		min:      1,
		max:      3,
		letter:   "a",
		password: "abcde",
	},
	{
		min:      1,
		max:      3,
		letter:   "b",
		password: "cdefg",
	},
	{
		min:      2,
		max:      9,
		letter:   "c",
		password: "ccccccccc",
	},
}

func Test_populatePasswordCollection(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		args    args
		want    []passwords
		wantErr bool
	}{
		{
			name: "returns an empty collection if no input provided",
			args: args{
				input: []string{},
			},
			want:    []passwords{},
			wantErr: false,
		},
		{
			name: "returns some passwords for the given input",
			args: args{
				input: []string{
					"1-3 a: abcde",
					"1-3 b: cdefg",
					"2-9 c: ccccccccc",
				},
			},
			want:    examplePasswordCollection,
			wantErr: false,
		},
		{
			name: "returns an error if an input line is invalid",
			args: args{
				input: []string{
					"1-3 a: abcde",
					"1-somerandomtext2 a: abc some random text",
					"2-9 c: ccccccccc",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := populatePasswordCollection(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("populatePasswordCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("populatePasswordCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readPassword(t *testing.T) {
	type args struct {
		match []string
	}
	tests := []struct {
		name string
		args args
		want passwords
	}{
		{
			name: "returns a password given a list of input matches",
			args: args{
				match: []string{"1-45 a:abcde", "1", "45", "s", "abcde"},
			},
			want: passwords{
				min:      1,
				max:      45,
				letter:   "s",
				password: "abcde",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readPassword(tt.args.match); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSolutions(t *testing.T) {
	type args struct {
		passwordCollection []passwords
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "advent of code example input",
			args: args{
				passwordCollection: examplePasswordCollection,
			},
			want:  2,
			want1: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getSolutions(tt.args.passwordCollection)
			if got != tt.want {
				t.Errorf("getSolutions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getSolutions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
