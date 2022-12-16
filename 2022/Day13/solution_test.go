package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createInterfaces(left, right string) ([]interface{}, []interface{}) {
	pkt1 := []interface{}{}
	pkt2 := []interface{}{}
	_ = json.Unmarshal([]byte(left), &pkt1)
	_ = json.Unmarshal([]byte(right), &pkt2)
	return pkt1, pkt2
}

func TestComparePackets(t *testing.T) {
	tests := []struct {
		left           string
		right          string
		comparisonFunc assert.ComparisonAssertionFunc
	}{
		{
			left:           "[1,1,3,1,1]",
			right:          "[1,1,5,1,1]",
			comparisonFunc: assert.LessOrEqual,
		},
		{
			left:           "[[1],[2,3,4]]",
			right:          "[[1],4]",
			comparisonFunc: assert.LessOrEqual,
		},
		{
			left:           "[9]",
			right:          "[[8,7,6]]",
			comparisonFunc: assert.Greater,
		},
		{
			left:           "[[4,4],4,4]",
			right:          "[[4,4],4,4,4]",
			comparisonFunc: assert.LessOrEqual,
		},
		{
			left:           "[7,7,7,7]",
			right:          "[7,7,7]",
			comparisonFunc: assert.Greater,
		},
		{
			left:           "[]",
			right:          "[3]",
			comparisonFunc: assert.LessOrEqual,
		},
		{
			left:           "[[[]]]",
			right:          "[[]]",
			comparisonFunc: assert.Greater,
		},
		{
			left:           "[1,[2,[3,[4,[5,6,7]]]],8,9]",
			right:          "[1,[2,[3,[4,[5,6,0]]]],8,9]",
			comparisonFunc: assert.Greater,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("compares packets and returns order, advent of code example %d", i), func(t *testing.T) {
			got := comparePackets(createInterfaces(tt.left, tt.right))
			tt.comparisonFunc(t, got, 0)
		})
	}
}

func TestFindSolutions(t *testing.T) {
	t.Run("returns an error if left input is invalid", func(t *testing.T) {
		input := []string{
			"[1,1,3,1,1]",
			"[1,1,5,1,1]",
			"",
			"[[1],[2,3,4]]",
			"[[1],4]",
			"",
			"[9]",
			"[[8,7,6]]",
			"",
			"[[4,4,4,4]",
			"[[4,4],4,4,4]",
			"",
			"[7,7,7,7]",
			"[7,7,7]",
			"",
			"[]",
			"[3]",
			"",
			"[[[]]]",
			"[[]]",
			"",
			"[1,[2,[3,[4,[5,6,7]]]],8,9]",
			"[1,[2,[3,[4,[5,6,0]]]],8,9]",
		}
		got, got1, err := findSolutions(input)
		assert.Equal(t, -1, got)
		assert.Equal(t, -1, got1)
		assert.Error(t, err)
	})

	t.Run("returns an error if right input is invalid", func(t *testing.T) {
		input := []string{
			"[1,1,3,1,1]",
			"[1,1,5,1,1]",
			"",
			"[[1],[2,3,4]]",
			"[[1],4]",
			"",
			"[9]",
			"[[8,7,]]",
			"",
			"[[4,4],4,4]",
			"[[4,4],4,4,4]",
			"",
			"[7,7,7,7]",
			"[7,7,7]",
			"",
			"[]",
			"[3]",
			"",
			"[[[]]]",
			"[[]]",
			"",
			"[1,[2,[3,[4,[5,6,7]]]],8,9]",
			"[1,[2,[3,[4,[5,6,0]]]],8,9]",
		}
		got, got1, err := findSolutions(input)
		assert.Equal(t, -1, got)
		assert.Equal(t, -1, got1)
		assert.Error(t, err)
	})

	t.Run("returns part1 and part2 solutions, advent of code example", func(t *testing.T) {
		input := []string{
			"[1,1,3,1,1]",
			"[1,1,5,1,1]",
			"",
			"[[1],[2,3,4]]",
			"[[1],4]",
			"",
			"[9]",
			"[[8,7,6]]",
			"",
			"[[4,4],4,4]",
			"[[4,4],4,4,4]",
			"",
			"[7,7,7,7]",
			"[7,7,7]",
			"",
			"[]",
			"[3]",
			"",
			"[[[]]]",
			"[[]]",
			"",
			"[1,[2,[3,[4,[5,6,7]]]],8,9]",
			"[1,[2,[3,[4,[5,6,0]]]],8,9]",
		}
		got, got1, err := findSolutions(input)
		assert.Equal(t, 13, got)
		assert.Equal(t, 140, got1)
		assert.NoError(t, err)
	})
}
