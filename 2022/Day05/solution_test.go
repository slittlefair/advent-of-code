package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	t.Run("returns an error if there is an error in an instruction line", func(t *testing.T) {
		input := []string{
			"    [D]    ",
			"[N] [C]    ",
			"[Z] [M] [P]",
			" 1   2   3 ",
			"",
			"move 1 from 2 to 1",
			"move 3 from 1 to 3",
			"move 2 from2 to 1",
			"move 1 from 1 to 2",
		}
		_, err := parseInput(input)
		assert.Error(t, err)
	})

	t.Run("returns a parsed crates struct from input, advent of code example", func(t *testing.T) {
		input := []string{
			"    [D]    ",
			"[N] [C]    ",
			"[Z] [M] [P]",
			" 1   2   3 ",
			"",
			"move 1 from 2 to 1",
			"move 3 from 1 to 3",
			"move 2 from 2 to 1",
			"move 1 from 1 to 2",
		}
		want := Crates{
			arrangements: [][]string{{"Z", "N"}, {"M", "C", "D"}, {"P"}},
			instructions: [][]int{{1, 1, 0}, {3, 0, 2}, {2, 1, 0}, {1, 0, 1}},
		}
		crates, err := parseInput(input)
		assert.NoError(t, err)
		assert.Equal(t, want, crates)
	})
}

func TestRunCrateMover9000(t *testing.T) {
	crates := Crates{
		arrangements: [][]string{{"Z", "N"}, {"M", "C", "D"}, {"P"}},
		instructions: [][]int{{1, 1, 0}, {3, 0, 2}, {2, 1, 0}, {1, 0, 1}},
	}
	crates.runCrateMover9000()
	want := [][]string{
		{"C"},
		{"M"},
		{"P", "D", "N", "Z"},
	}
	assert.Equal(t, crates.arrangements, want)
}

func TestRunCrateMover9001(t *testing.T) {
	crates := Crates{
		arrangements: [][]string{{"Z", "N"}, {"M", "C", "D"}, {"P"}},
		instructions: [][]int{{1, 1, 0}, {3, 0, 2}, {2, 1, 0}, {1, 0, 1}},
	}
	crates.runCrateMover9001()
	want := [][]string{
		{"M"},
		{"C"},
		{"P", "Z", "N", "D"},
	}
	assert.Equal(t, crates.arrangements, want)
}

func TestFinalToppers(t *testing.T) {
	tests := []struct {
		arrangement [][]string
		want        string
	}{
		{
			arrangement: [][]string{
				{"C"},
				{"M"},
				{"P", "D", "N", "Z"},
			},
			want: "CMZ",
		},
		{
			arrangement: [][]string{
				{"M"},
				{"C"},
				{"P", "Z", "N", "D"},
			},
			want: "MCD",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("finds the top crates in the arrangement, advent of code example %d", i+1), func(t *testing.T) {
			crates := Crates{arrangements: tt.arrangement}
			got := crates.finalToppers()
			assert.Equal(t, tt.want, got)
		})
	}
}
