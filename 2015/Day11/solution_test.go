package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_letterIndex(t *testing.T) {
	tests := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("returns %d if letter is %s", i, tt), func(t *testing.T) {
			if got := letterIndex(tt); got != i {
				t.Errorf("letterIndex() = %v, want %v", got, i)
			}
		})
	}
	t.Run("returns -1 if letter is aa", func(t *testing.T) {
		if got := letterIndex("aa"); got != -1 {
			t.Errorf("letterIndex() = %v, want %v", got, -1)
		}
	})
}

func Test_incrementCharacter(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{
			name: "increments character from 0 to 1",
			arg:  0,
			want: 1,
		},
		{
			name: "increments character from 14 to 15",
			arg:  14,
			want: 15,
		},
		{
			name: "increments character from 25 to 0",
			arg:  25,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := incrementCharacter(tt.arg); got != tt.want {
				t.Errorf("incrementCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassword_RemoveIllegalCharacters(t *testing.T) {
	tests := []struct {
		name string
		p    Password
		want Password
	}{
		{
			name: "doesn't change the password if all characters are legal",
			p:    Password{0, 1, 2, 3, 4, 5, 10, 20},
			want: Password{0, 1, 2, 3, 4, 5, 10, 20},
		},
		{
			name: "increments to the next legal password if it contains 8",
			p:    Password{0, 1, 2, 3, 4, 5, 10, 8},
			want: Password{0, 1, 2, 3, 4, 5, 10, 9},
		},
		{
			name: "increments to the next legal password if it contains 11",
			p:    Password{0, 1, 2, 3, 11, 5, 10, 20},
			want: Password{0, 1, 2, 3, 12, 0, 0, 0},
		},
		{
			name: "increments to the next legal password if it contains 14",
			p:    Password{14, 1, 2, 3, 4, 5, 10, 8},
			want: Password{15, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.p
			p.RemoveIllegalCharacters()
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Password.RemoveIllegalCharacters() = %v, want %v", p, tt.want)
			}
		})
	}
}

func TestPassword_IncrementPassword(t *testing.T) {
	tests := []struct {
		name string
		p    Password
		want Password
	}{
		{
			name: "increments a password that doesn't contain an illegal character",
			p:    Password{24, 24, 24, 24, 24, 24, 23, 23},
			want: Password{24, 24, 24, 24, 24, 24, 23, 24},
		},
		{
			name: "increments a password that doesn't contain an illegal character but does loop",
			p:    Password{24, 24, 24, 24, 24, 24, 25, 25},
			want: Password{24, 24, 24, 24, 24, 25, 0, 0},
		},
		{
			name: "increments a password that doesn't contain an illegal character but does loops from max to min",
			p:    Password{25, 25, 25, 25, 25, 25, 25, 25},
			want: Password{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "increments a password that contains an illegal character",
			p:    Password{0, 3, 6, 7, 25, 25, 25, 25},
			want: Password{0, 3, 6, 8, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.p
			p.IncrementPassword()
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Password.IncrementPassword() = %v, want %v", p, tt.want)
			}
		})
	}
}

func TestPassword_HasIncreasingStraightLetters(t *testing.T) {
	tests := []struct {
		name string
		p    Password
		want bool
	}{
		{
			name: "returns true if the password contains 3 increasing consecutive letters at the start of the password",
			p:    Password{18, 19, 20, 1, 3, 5, 5, 9},
			want: true,
		},
		{
			name: "returns true if the password contains 3 increasing consecutive letters in the middle of the password",
			p:    Password{0, 1, 0, 1, 3, 4, 5, 9},
			want: true,
		},
		{
			name: "returns true if the password contains 3 increasing consecutive letters in the end of the password",
			p:    Password{20, 1, 0, 1, 9, 4, 5, 6},
			want: true,
		},
		{
			name: "returns false if the password doesn't contains 3 increasing consecutive letters in the password",
			p:    Password{1, 2, 4, 10, 19, 19, 19, 0},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.HasIncreasingStraightLetters(); got != tt.want {
				t.Errorf("Password.HasIncreasingStraightLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassword_HasDifferentPairs(t *testing.T) {
	tests := []struct {
		name string
		p    Password
		want bool
	}{
		{
			name: "returns true if the password contains 2 different pairs of letters",
			p:    Password{1, 2, 1, 2, 1, 1, 2, 2},
			want: true,
		},
		{
			name: "returns true if the password contains different more than 2 of letters",
			p:    Password{1, 1, 2, 2, 3, 3, 2, 2},
			want: true,
		},
		{
			name: "returns false if the password contains no pairs of letters",
			p:    Password{1, 2, 3, 4, 1, 2, 3, 4},
			want: false,
		},
		{
			name: "returns false if the password contains 1 different pairs of letters",
			p:    Password{1, 2, 1, 2, 1, 2, 1, 1},
			want: false,
		},
		{
			name: "returns false if the password contains 1 different pairs of letters but 1 triple",
			p:    Password{1, 2, 1, 2, 1, 1, 1, 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.HasDifferentPairs(); got != tt.want {
				t.Errorf("Password.HasDifferentPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassword_IsValid(t *testing.T) {
	tests := []struct {
		name string
		p    Password
		want bool
	}{
		{
			name: "returns false if password doesn't have increasing straight letters",
			p:    Password{0, 0, 0, 0, 0, 0, 0, 0},
			want: false,
		},
		{
			name: "returns false if password doesn't have different pairs",
			p:    Password{0, 1, 2, 3, 4, 5, 6, 7},
			want: false,
		},
		{
			name: "returns true if password has increasing straight letters and different pairs",
			p:    Password{0, 0, 1, 1, 2, 3, 0, 0},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsValid(); got != tt.want {
				t.Errorf("Password.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassword_ConvertToString(t *testing.T) {
	tests := []struct {
		name string
		p    Password
		want string
	}{
		{
			name: "converts password to string",
			p:    Password{0, 1, 2, 25, 10, 9, 20, 21},
			want: "abczkjuv",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.ConvertToString(); got != tt.want {
				t.Errorf("Password.ConvertToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPassword_GetNextValidPassword(t *testing.T) {
	tests := []struct {
		name string
		p    Password
		want Password
	}{
		{
			name: "advent of code example 1",
			p:    Password{0, 1, 2, 3, 4, 5, 6, 7},
			want: Password{0, 1, 2, 3, 5, 5, 0, 0},
		},
		{
			name: "advent of code example 2",
			p:    Password{6, 7, 9, 0, 0, 0, 0, 0},
			want: Password{6, 7, 9, 0, 0, 1, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.p
			p.GetNextValidPassword()
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Password.GetNextValidPassword() = %v, want %v", p, tt.want)
			}
		})
	}
}

func Test_makePassword(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want *Password
	}{
		{
			name: "creates a password with no illegal characters",
			arg:  "abcdefgh",
			want: &Password{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "creates a password with illegal characters",
			arg:  "ghijklmn",
			want: &Password{6, 7, 9, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makePassword(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
