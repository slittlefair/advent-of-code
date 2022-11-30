package strings

import (
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
			got := IsUpper(tt.s)
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
			got := IsLower(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
