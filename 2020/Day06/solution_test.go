package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getSolution(t *testing.T) {
	t.Run("advent of code example", func(t *testing.T) {
		input := []string{
			"abc",
			"",
			"a",
			"b",
			"c",
			"",
			"ab",
			"ac",
			"",
			"a",
			"a",
			"a",
			"a",
			"",
			"b",
		}
		got, got1 := getSolution(input)
		assert.Equal(t, 11, got)
		assert.Equal(t, 6, got1)
	})
}
