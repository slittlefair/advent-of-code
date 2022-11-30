package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleInstructions = []Instructions{
	{
		instruction: "nop",
		value:       0,
	},
	{
		instruction: "acc",
		value:       1,
	},
	{
		instruction: "jmp",
		value:       4,
	},
	{
		instruction: "acc",
		value:       3,
	},
	{
		instruction: "jmp",
		value:       -3,
	},
	{
		instruction: "acc",
		value:       -99,
	},
	{
		instruction: "acc",
		value:       1,
	},
	{
		instruction: "jmp",
		value:       -4,
	},
	{
		instruction: "acc",
		value:       6,
	},
}

func Test_parseProgramme(t *testing.T) {
	tests := []struct {
		name               string
		entries            []string
		want               *Programme
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if an input value can't be converted to an int",
			entries: []string{
				"nop +0",
				"acc -9",
				"acc ???",
				"jmp +2",
			},
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "advent of code example",
			entries: []string{
				"nop +0",
				"acc +1",
				"jmp +4",
				"acc +3",
				"jmp -3",
				"acc -99",
				"acc +1",
				"jmp -4",
				"acc +6",
			},
			want: &Programme{
				instructions: exampleInstructions,
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseProgramme(tt.entries)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProgramme_runProgramme(t *testing.T) {
	type fields struct {
		foundSolution bool
		instructions  []Instructions
	}
	tests := []struct {
		name         string
		fields       fields
		tweakAtIndex int
		want         int
	}{
		{
			name: "advent of code example 1",
			fields: fields{
				foundSolution: false,
				instructions:  exampleInstructions,
			},
			tweakAtIndex: -1,
			want:         5,
		},
		{
			name: "advent of code example 2",
			fields: fields{
				foundSolution: false,
				instructions:  exampleInstructions,
			},
			tweakAtIndex: 7,
			want:         8,
		},
		{
			name: "advent of code example tweak at nop",
			fields: fields{
				foundSolution: false,
				instructions:  exampleInstructions,
			},
			tweakAtIndex: 0,
			want:         0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Programme{
				foundSolution: tt.fields.foundSolution,
				instructions:  tt.fields.instructions,
			}
			got := p.runProgramme(tt.tweakAtIndex)
			assert.Equal(t, tt.want, got)
		})
	}
}
