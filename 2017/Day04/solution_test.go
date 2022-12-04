package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidPart1(t *testing.T) {
	tests := []struct {
		name             string
		matches          []string
		boolAssertonFunc assert.BoolAssertionFunc
	}{
		{
			name: "returns valid for all different words, advent of code example 1",
			matches: []string{
				"aa", "bb", "cc", "dd", "ee",
			},
			boolAssertonFunc: assert.True,
		},
		{
			name: "returns invalid if any words are repeated, advent of code example 2",
			matches: []string{
				"aa", "bb", "cc", "dd", "aa",
			},
			boolAssertonFunc: assert.False,
		},
		{
			name: "returns valid for all different words if one is a substring, advent of code example 3",
			matches: []string{
				"aa", "bb", "cc", "dd", "aaa",
			},
			boolAssertonFunc: assert.True,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidPart1(tt.matches)
			tt.boolAssertonFunc(t, got)
		})
	}
}

func TestIsValidPart2(t *testing.T) {
	tests := []struct {
		name             string
		matches          []string
		boolAssertonFunc assert.BoolAssertionFunc
	}{
		{
			name: "returns valid for no anagram words, advent of code example 1",
			matches: []string{
				"abcde", "fghij",
			},
			boolAssertonFunc: assert.True,
		},
		{
			name: "returns invalid if any words are anagrams of each other, advent of code example 2",
			matches: []string{
				"abcde", "xyz", "ecdab",
			},
			boolAssertonFunc: assert.False,
		},
		{
			name: "returns valid for all different anagram words if one is a substring, advent of code example 3",
			matches: []string{
				"a", "ab", "abc", "abd", "abf", "abj",
			},
			boolAssertonFunc: assert.True,
		},
		{
			name: "returns valid for no anagram words, advent of code example 4",
			matches: []string{
				"iiii", "oiii", "ooii", "oooi", "oooo",
			},
			boolAssertonFunc: assert.True,
		},
		{
			name: "returns invalid if any words are anagrams of each other, advent of code example 5",
			matches: []string{
				"oiii", "ioii", "iioi", "iiio",
			},
			boolAssertonFunc: assert.False,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidPart2(tt.matches)
			tt.boolAssertonFunc(t, got)
		})
	}
}

func TestCoundValidPhrases(t *testing.T) {
	t.Run("returns count of part1 and part2 valid phrases", func(t *testing.T) {
		input := []string{
			"aa bb cc dd ee",
			"aa bb cc dd aa",
			"aa bb cc dd aaa",
			"abcde fghij",
			"abcde xyz ecdab",
			"a ab abc abd abf abj",
			"iiii oiii ooii oooi oooo",
			"oiii ioii iioi iiio",
		}
		got, got1 := countValidPhrases(input)
		assert.Equal(t, 7, got)
		assert.Equal(t, 5, got1)
	})
}
