package strings

import "unicode"

// IsUpper returns true if all characters in a string are upper case, false otherwise
func IsUpper(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return false
		}
	}
	return true
}

// IsLower returns true if all characters in a string are lower case, false otherwise
func IsLower(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

// AreAnagrams returns whether the given words are anagrams of each other
func AreAnagrams(x, y string) bool {
	lenX := len(x)
	lenY := len(y)

	if lenX != lenY {
		return false
	}

	lettersX := map[rune]int{}
	lettersY := map[rune]int{}

	for _, r := range x {
		lettersX[r]++
	}
	for _, r := range y {
		lettersY[r]++
	}

	for _, r := range x {
		if lettersX[r] != lettersY[r] {
			return false
		}
	}
	return true
}
