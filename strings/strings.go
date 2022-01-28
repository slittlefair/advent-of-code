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
