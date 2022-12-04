package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var aocInput = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func TestFindCommonItems(t *testing.T) {
	t.Run("adds the common items to items struct, advent of code example", func(t *testing.T) {
		items := &Items{}
		want := &Items{'p', 'L', 'P', 'v', 't', 's'}
		items.findCommonItems(aocInput)
		assert.Equal(t, want, items)
	})
}

func TestGetBadges(t *testing.T) {
	t.Run("adds the common items to items struct, advent of code example", func(t *testing.T) {
		items := &Items{}
		want := &Items{'r', 'Z'}
		items.getBadges(aocInput)
		assert.Equal(t, want, items)
	})
}

func TestSumPriorities(t *testing.T) {
	tests := []struct {
		items *Items
		want  int
	}{
		{
			items: &Items{'p', 'L', 'P', 'v', 't', 's'},
			want:  157,
		},
		{
			items: &Items{'r', 'Z'},
			want:  70,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("calculates correct priority for given items, advent of code example %d", i+1), func(t *testing.T) {
			got := tt.items.sumPriorities()
			assert.Equal(t, tt.want, got)
		})
	}
}
