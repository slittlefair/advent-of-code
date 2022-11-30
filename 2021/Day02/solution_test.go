package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_evaluateLine(t *testing.T) {
	tests := []struct {
		name               string
		line               string
		want               *instruction
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "returns an error if the second part of the line cannot be converted to an int",
			line:               "forward one",
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an instruction from a line of input",
			line: "forward 4",
			want: &instruction{
				dir: "forward",
				val: 4,
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := evaluateLine(tt.line)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_validateInstruction(t *testing.T) {
	tests := []struct {
		name               string
		dir                string
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               `does not return an error if the direction is "forward"`,
			dir:                "forward",
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               `does not return an error if the direction is "up"`,
			dir:                "up",
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               `does not return an error if the direction is "down"`,
			dir:                "down",
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               `returns an error if the direction is not "forward", "up" or "down"`,
			dir:                "diagonally",
			errorAssertionFunc: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateInstruction(tt.dir)
			tt.errorAssertionFunc(t, err)
		})
	}
}

func Test_position_followInstruction(t *testing.T) {
	type args struct {
		inst  instruction
		part2 bool
	}
	tests := []struct {
		name string
		co   *position
		args args
		want *position
	}{
		{
			name: "follows a part 1 forward instruction",
			co: &position{
				X: 5,
				Y: 5,
			},
			args: args{
				inst: instruction{
					dir: "forward",
					val: 8,
				},
				part2: false,
			},
			want: &position{
				X: 13,
				Y: 5,
			},
		},
		{
			name: "follows a part 1 up instruction",
			co: &position{
				X: 13,
				Y: 5,
			},
			args: args{
				inst: instruction{
					dir: "up",
					val: 3,
				},
				part2: false,
			},
			want: &position{
				X: 13,
				Y: 2,
			},
		},
		{
			name: "follows a part 1 down instruction",
			co: &position{
				X: 13,
				Y: 2,
			},
			args: args{
				inst: instruction{
					dir: "down",
					val: 8,
				},
				part2: false,
			},
			want: &position{
				X: 13,
				Y: 10,
			},
		},
		{
			name: "follows a part 2 forward instruction",
			co: &position{
				X:   5,
				Y:   0,
				aim: 5,
			},
			args: args{
				inst: instruction{
					dir: "forward",
					val: 8,
				},
				part2: true,
			},
			want: &position{
				X:   13,
				Y:   40,
				aim: 5,
			},
		},
		{
			name: "follows a part 2 up instruction",
			co: &position{
				X:   13,
				Y:   40,
				aim: 5,
			},
			args: args{
				inst: instruction{
					dir: "up",
					val: 3,
				},
				part2: true,
			},
			want: &position{
				X:   13,
				Y:   40,
				aim: 2,
			},
		},
		{
			name: "follows a part 2 down instruction",
			co: &position{
				X:   13,
				Y:   40,
				aim: 2,
			},
			args: args{
				inst: instruction{
					dir: "down",
					val: 8,
				},
				part2: true,
			},
			want: &position{
				X:   13,
				Y:   40,
				aim: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			co := tt.co
			co.followInstruction(tt.args.inst, tt.args.part2)
			assert.Equal(t, tt.want, co)
		})
	}
}

func Test_findSolution(t *testing.T) {
	type args struct {
		input []string
		part2 bool
	}
	tests := []struct {
		name               string
		args               args
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if an input line is invalid part 1, evaluateLine",
			args: args{
				input: []string{
					"forward 5",
					"down 5",
					"forward 8",
					"up trois",
					"down 8",
					"forward 2",
				},
				part2: false,
			},
			want:               -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if an input line is invalid part 1, validateInstruction",
			args: args{
				input: []string{
					"forward 5",
					"down 5",
					"forward 8",
					"up 3",
					"across 8",
					"forward 2",
				},
				part2: false,
			},
			want:               -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "calculates product of horizontal position and depth part 1, advent of code example",
			args: args{
				input: []string{
					"forward 5",
					"down 5",
					"forward 8",
					"up 3",
					"down 8",
					"forward 2",
				},
				part2: false,
			},
			want:               150,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns an error if an input line is invalid part 2, evaluateLine",
			args: args{
				input: []string{
					"forward 5",
					"down 5",
					"forward 8",
					"up trois",
					"down 8",
					"forward 2",
				},
				part2: true,
			},
			want:               -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if an input line is invalid part 2, validateInstruction",
			args: args{
				input: []string{
					"forward 5",
					"down 5",
					"forward 8",
					"up 3",
					"across 8",
					"forward 2",
				},
				part2: true,
			},
			want:               -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "calculates product of horizontal position and depth part 1, advent of code example",
			args: args{
				input: []string{
					"forward 5",
					"down 5",
					"forward 8",
					"up 3",
					"down 8",
					"forward 2",
				},
				part2: true,
			},
			want:               900,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findSolution(tt.args.input, tt.args.part2)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
