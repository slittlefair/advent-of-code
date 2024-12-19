package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getMultiplication(t *testing.T) {
	t.Run("returns multiplication for a given input string", func(t *testing.T) {
		assert.Equal(t, 2024, getMultiplication("mul(44,46)"))
	})

	t.Run("returns multiplication fir a given input when one of the values is 0", func(t *testing.T) {
		assert.Equal(t, 0, getMultiplication("mul(5,0)"))
	})
}

func Test_findSolutions(t *testing.T) {
	t.Run("returns correct part 1 for a given input", func(t *testing.T) {
		part1, _ := findSolutions([]string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"})
		assert.Equal(t, 161, part1)
	})

	t.Run("returns correct part 2 for a given input", func(t *testing.T) {
		_, part2 := findSolutions([]string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"})
		assert.Equal(t, 48, part2)
	})
}
