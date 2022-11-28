package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name               string
		input              []string
		want               int
		want1              int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "returns an error if there are less than 2 numbers in the row",
			input:              []string{"row 3, first column"},
			want:               -1,
			want1:              -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "returns an error if there are more than 2 numbers in the row",
			input:              []string{"row 3, column 4 or 5"},
			want:               -1,
			want1:              -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "returns column and row from given input",
			input:              []string{"row 3, column 4"},
			want:               3,
			want1:              4,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseInput(tt.input)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_nthNumber(t *testing.T) {
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "returns the correct number from the first row",
			args: args{
				row: 1,
				col: 4,
			},
			want: 10,
		},
		{
			name: "returns the correct number from the first column",
			args: args{
				row: 6,
				col: 1,
			},
			want: 16,
		},
		{
			name: "returns the correct number from the table",
			args: args{
				row: 4,
				col: 3,
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nthNumber(tt.args.row, tt.args.col)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getCodeAt(t *testing.T) {
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "returns the correct code at row 2, col 3",
			args: args{
				row: 2,
				col: 3,
			},
			want: 16929656,
		},
		{
			name: "returns the correct code at row 5, col 6",
			args: args{
				row: 5,
				col: 6,
			},
			want: 31663883,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getCodeAt(tt.args.row, tt.args.col)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getSolution(t *testing.T) {
	tests := []struct {
		name               string
		input              []string
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "returns an error from a dodgy input",
			input:              []string{"this won't work"},
			want:               -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "returns correct code from input line",
			input:              []string{"To continue, please consult the code grid in the manual.  Enter the code at row 6, column 4."},
			want:               24659492,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getSolution(tt.input)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
