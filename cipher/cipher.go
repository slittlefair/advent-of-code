package cipher

// CaesarCipher applies a Caesar Cipher to the given text, shifted shiftNum times
func CaesarCipher(text string, shiftNum int) string {
	shift, offset := rune(shiftNum%26), rune(26)

	runes := []rune(text)

	for index, char := range runes {
		if char >= 'a' && char <= 'z'-shift ||
			char >= 'A' && char <= 'Z'-shift {
			char = char + shift
		} else if char > 'z'-shift && char <= 'z' ||
			char > 'Z'-shift && char <= 'Z' {
			char = char + shift - offset
		}

		// Above handles both upper and lower case ASCII
		// characters; anything else is returned as is (includes
		// numbers, punctuation and space).
		runes[index] = char
	}

	return string(runes)
}
