package cipher_test

import (
	"Advent-of-Code/cipher"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaesarCipher(t *testing.T) {
	t.Run("applies Caesar Cipher to the given text shifted number of supplied times", func(t *testing.T) {
		got := cipher.CaesarCipher("qZmt-zixMtkozy-Ivhz-343", 343)
		assert.Equal(t, "vEry-encRypted-Name-343", got)
	})
}
