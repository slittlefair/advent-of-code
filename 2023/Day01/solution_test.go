package main

import (
	"Advent-of-Code/maths"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessLine(t *testing.T) {
	tests := []struct {
		name      string
		line      string
		wantTens  Num
		wantUnits Num
	}{
		{
			name: "returns tens and units with num values, simple",
			line: "1abc2",
			wantTens: Num{
				idx:    0,
				val:    1,
				strIdx: maths.Infinity,
			},
			wantUnits: Num{
				idx:    4,
				val:    2,
				strIdx: -1,
			},
		},
		{
			name: "returns tens and units with num values, simple",
			line: "a1b2c3d4e5f",
			wantTens: Num{
				idx:    1,
				val:    1,
				strIdx: maths.Infinity,
			},
			wantUnits: Num{
				idx:    9,
				val:    5,
				strIdx: -1,
			},
		},
		{
			name: "returns tens and units with same num values",
			line: "treb7uchet",
			wantTens: Num{
				idx:    4,
				val:    7,
				strIdx: maths.Infinity,
			},
			wantUnits: Num{
				idx:    4,
				val:    7,
				strIdx: -1,
			},
		},
		{
			name: "returns tens and units with num and string values, example 1",
			line: "two1nine",
			wantTens: Num{
				idx:    3,
				val:    1,
				strIdx: 0,
				strVal: 2,
			},
			wantUnits: Num{
				idx:    3,
				val:    1,
				strIdx: 4,
				strVal: 9,
			},
		},
		{
			name: "returns tens and units with num and string values, example 2",
			line: "7pqrstsixteen",
			wantTens: Num{
				idx:    0,
				val:    7,
				strIdx: 6,
				strVal: 6,
			},
			wantUnits: Num{
				idx:    0,
				val:    7,
				strIdx: 6,
				strVal: 6,
			},
		},
		{
			name: "returns tens and units with num and string values, example 3",
			line: "4nineeightseven2",
			wantTens: Num{
				idx:    0,
				val:    4,
				strIdx: 1,
				strVal: 9,
			},
			wantUnits: Num{
				idx:    15,
				val:    2,
				strIdx: 10,
				strVal: 7,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tens, units := processLine(tt.line)
			assert.Equal(t, tt.wantTens, tens)
			assert.Equal(t, tt.wantUnits, units)
		})
	}
}

func TestFindSolution(t *testing.T) {
	t.Run("", func(t *testing.T) {
		input := []string{
			"1abc2",
			"pqr3stu8vwx",
			"a1b2c3d4e5f",
			"treb7uchet",
		}
		part1, _ := findSolution(input)
		assert.Equal(t, 142, part1)
	})

	t.Run("", func(t *testing.T) {
		input := []string{
			"two1nine",
			"eightwothree",
			"abcone2threexyz",
			"xtwone3four",
			"4nineeightseven2",
			"zoneight234",
			"7pqrstsixteen",
		}
		_, part2 := findSolution(input)
		assert.Equal(t, 281, part2)
	})
}
