package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInput(t *testing.T) {
	t.Run("returns a Stones map for the given input, 1", func(t *testing.T) {
		input := []string{"0 1 10 99 999"}
		got := parseInput(input)
		want := Stones{
			"0":   1,
			"1":   1,
			"10":  1,
			"99":  1,
			"999": 1,
		}

		assert.Equal(t, want, got)
	})

	t.Run("returns a Stones map for the given input, 2", func(t *testing.T) {
		input := []string{"125 17"}
		got := parseInput(input)
		want := Stones{
			"125": 1,
			"17":  1,
		}

		assert.Equal(t, want, got)
	})

	t.Run("returns a Stones map for the given input, 3", func(t *testing.T) {
		input := []string{"0 1 10 99 999 6 2024 1 0 10 3 1"}
		got := parseInput(input)
		want := Stones{
			"0":    2,
			"1":    3,
			"10":   2,
			"99":   1,
			"999":  1,
			"6":    1,
			"2024": 1,
			"3":    1,
		}

		assert.Equal(t, want, got)
	})
}

func Test_split(t *testing.T) {
	t.Run("returns false if the given string has an odd length", func(t *testing.T) {
		s1, s2, success := split("123")
		assert.Zero(t, s1)
		assert.Zero(t, s2)
		assert.False(t, success)
	})

	t.Run("returns the split of a string of even length", func(t *testing.T) {
		s1, s2, success := split("2024")
		assert.Equal(t, "20", s1)
		assert.Equal(t, "24", s2)
		assert.True(t, success)
	})

	t.Run("returns the split of a string of even length, handling leading zeroes", func(t *testing.T) {
		s1, s2, success := split("224001")
		assert.Equal(t, "224", s1)
		assert.Equal(t, "1", s2)
		assert.True(t, success)
	})

	t.Run("returns the split of a string of even length, handling all zeroes", func(t *testing.T) {
		s1, s2, success := split("224000")
		assert.Equal(t, "224", s1)
		assert.Equal(t, "0", s2)
		assert.True(t, success)
	})
}

func TestStones_blink(t *testing.T) {
	t.Run("correctly turns 0s to 1s", func(t *testing.T) {
		s := Stones{"0": 1}
		want := Stones{"1": 1}

		got := s.blink()
		assert.Equal(t, want, got)
	})

	t.Run("correctly splits even length numbers", func(t *testing.T) {
		s := Stones{"23": 1, "2024": 2, "2006": 3, "1000": 1}
		want := Stones{"2": 1, "3": 1, "20": 5, "24": 2, "6": 3, "10": 1, "0": 1}

		got := s.blink()
		assert.Equal(t, want, got)
	})

	t.Run("correctly multiplies other numbers by 2024", func(t *testing.T) {
		s := Stones{"1": 1, "9": 1, "123": 2}
		want := Stones{"2024": 1, "18216": 1, "248952": 2}

		got := s.blink()
		assert.Equal(t, want, got)
	})

	t.Run("correctly changes all stones on a blink", func(t *testing.T) {
		s := Stones{"0": 2, "11": 2, "1234": 1, "5": 1012, "3003": 2}
		want := Stones{"1": 6, "12": 1, "34": 1, "10120": 1012, "30": 2, "3": 2}

		got := s.blink()
		assert.Equal(t, want, got)
	})
}

func TestStones_findLength(t *testing.T) {
	t.Run("returns length of frequencies in provided Stones", func(t *testing.T) {
		s := Stones{"1": 6, "12": 1, "34": 1, "10120": 1012, "30": 2, "3": 2}
		got := s.findLength()
		assert.Equal(t, 1024, got)
	})
}

func Test_findSolutions(t *testing.T) {
	t.Run("returns answers to parts 1 and 2 for a given input", func(t *testing.T) {
		input := []string{"125 17"}
		part1, part2 := findSolutions(input)
		assert.Equal(t, 55312, part1)
		assert.Equal(t, 65601038650482, part2)
	})
}
