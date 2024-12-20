package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var examplePasswordCollection = []passwords{
	{
		minimum:  1,
		maximum:  3,
		letter:   "a",
		password: "abcde",
	},
	{
		minimum:  1,
		maximum:  3,
		letter:   "b",
		password: "cdefg",
	},
	{
		minimum:  2,
		maximum:  9,
		letter:   "c",
		password: "ccccccccc",
	},
}

func Test_populatePasswordCollection(t *testing.T) {
	tests := []struct {
		name               string
		input              []string
		want               []passwords
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "returns an empty collection if no input provided",
			input:              []string{},
			want:               []passwords{},
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns some passwords for the given input",
			input: []string{
				"1-3 a: abcde",
				"1-3 b: cdefg",
				"2-9 c: ccccccccc",
			},
			want:               examplePasswordCollection,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns an error if an input line is invalid",
			input: []string{
				"1-3 a: abcde",
				"1-somerandomtext2 a: abc some random text",
				"2-9 c: ccccccccc",
			},
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := populatePasswordCollection(tt.input)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_readPassword(t *testing.T) {
	t.Run("returns a password given a list of input matches", func(t *testing.T) {
		want := passwords{
			minimum:  1,
			maximum:  45,
			letter:   "s",
			password: "abcde",
		}
		got := readPassword(
			[]string{"1-45 a:abcde", "1", "45", "s", "abcde"},
		)
		assert.Equal(t, want, got)
	})
}

func Test_getSolutions(t *testing.T) {
	t.Run("advent of code example input", func(t *testing.T) {
		got, got1 := getSolutions(examplePasswordCollection)
		assert.Equal(t, 2, got)
		assert.Equal(t, 1, got1)
	})
}
