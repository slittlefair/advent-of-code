package strings_test

import (
	"Advent-of-Code/strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUpper(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "returns false if the string is not all in upper case",
			s:    "HELLO WOrLD",
			want: false,
		},
		{
			name: "returns true if the string is all in upper case",
			s:    "HELLO WORLD",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strings.IsUpper(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsLower(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "returns false if the string is not all in lower case",
			s:    "hellO world",
			want: false,
		},
		{
			name: "returns true if the string is all in lower case",
			s:    "hello world",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strings.IsLower(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAreAnagrams(t *testing.T) {
	t.Run("returns false if strings are of different lengths", func(t *testing.T) {
		got := strings.AreAnagrams("angered", "enrage")
		assert.False(t, got)
	})

	t.Run("returns false if strings are same length but not anagrams", func(t *testing.T) {
		got := strings.AreAnagrams("angered", "enrages")
		assert.False(t, got)
	})

	t.Run("returns true if strings are anagrams", func(t *testing.T) {
		got := strings.AreAnagrams("dictionary", "indicatory")
		assert.True(t, got)
	})
}
