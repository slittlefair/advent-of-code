package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"47|53",
	"97|13",
	"97|61",
	"97|47",
	"75|29",
	"61|13",
	"75|53",
	"29|13",
	"97|29",
	"53|29",
	"61|53",
	"97|53",
	"61|29",
	"47|13",
	"75|47",
	"97|75",
	"47|61",
	"75|61",
	"47|29",
	"75|13",
	"53|13",
	"",
	"75,47,61,53,29",
	"97,61,53,29,13",
	"75,29,13",
	"75,97,47,61,53",
	"61,13,29",
	"97,13,75,29,47",
}

var expectedOrderingRules = orderingRules{
	53: {
		47: false,
		75: false,
		29: true,
		97: false,
		13: true,
		61: false,
	},
	13: {
		97: false,
		29: false,
		61: false,
		75: false,
		47: false,
		53: false,
	},
	29: {
		13: true,
		75: false,
		97: false,
		53: false,
		61: false,
		47: false,
	},
	61: {
		13: true,
		53: true,
		29: true,
		97: false,
		47: false,
		75: false,
	},
	75: {
		29: true,
		53: true,
		47: true,
		61: true,
		13: true,
		97: false,
	},
	47: {
		53: true,
		13: true,
		61: true,
		29: true,
		97: false,
		75: false,
	},
	97: {
		13: true,
		61: true,
		47: true,
		29: true,
		53: true,
		75: true,
	},
}

func Test_parseInput(t *testing.T) {
	t.Run("returns ordering rules and pages for a given input", func(t *testing.T) {
		expectedPages := [][]int{
			{75, 47, 61, 53, 29},
			{97, 61, 53, 29, 13},
			{75, 29, 13},
			{75, 97, 47, 61, 53},
			{61, 13, 29},
			{97, 13, 75, 29, 47},
		}
		orderingRules, pages := parseInput(input)

		assert.Equal(t, expectedOrderingRules, orderingRules)
		assert.Equal(t, expectedPages, pages)
	})
}

func Test_sortPages(t *testing.T) {
	tests := []struct {
		pages    []int
		expected []int
	}{
		{
			pages:    []int{75, 97, 47, 61, 53},
			expected: []int{97, 75, 47, 61, 53},
		},
		{
			pages:    []int{61, 13, 29},
			expected: []int{61, 29, 13},
		},
		{
			pages:    []int{97, 13, 75, 29, 47},
			expected: []int{97, 75, 47, 29, 13},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("sorts pages based on ordering rules, %d", i+1), func(t *testing.T) {
			assert.Equal(t, tt.expected, sortPages(expectedOrderingRules, tt.pages))
		})
	}
}

func Test_findSolutions(t *testing.T) {
	t.Run("finds part1 and part2 slutions for a given input", func(t *testing.T) {
		part1, part2 := findSolutions(input)
		assert.Equal(t, 143, part1)
		assert.Equal(t, 123, part2)
	})
}
