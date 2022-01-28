package strings

import "testing"

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
			if got := IsUpper(tt.s); got != tt.want {
				t.Errorf("IsUpper() = %v, want %v", got, tt.want)
			}
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
			if got := IsLower(tt.s); got != tt.want {
				t.Errorf("IsLower() = %v, want %v", got, tt.want)
			}
		})
	}
}
