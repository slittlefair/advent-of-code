package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_registers_cpy(t *testing.T) {
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		name               string
		r                  registers
		args               args
		want               registers
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if from value can't be converted to an int",
			r: registers{
				"a": 1,
				"b": 2,
				"c": 3,
				"d": 4,
			},
			args: args{
				from: "one",
				to:   "a",
			},
			want: registers{
				"a": 1,
				"b": 2,
				"c": 3,
				"d": 4,
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "sets a register to the value of another register if provided",
			r: registers{
				"a": 1,
				"b": 2,
				"c": 3,
				"d": 4,
			},
			args: args{
				from: "b",
				to:   "a",
			},
			want: registers{
				"a": 2,
				"b": 2,
				"c": 3,
				"d": 4,
			},
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "sets a register to a value given",
			r: registers{
				"a": 1,
				"b": 2,
				"c": 3,
				"d": 4,
			},
			args: args{
				from: "9",
				to:   "a",
			},
			want: registers{
				"a": 9,
				"b": 2,
				"c": 3,
				"d": 4,
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			err := r.cpy(tt.args.from, tt.args.to)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, r)
		})
	}
}

func Test_registers_inc(t *testing.T) {
	t.Run("increments the value of the given register by 1", func(t *testing.T) {
		r := registers{
			"a": 3,
			"b": 88,
			"c": 64,
			"d": 12,
		}
		want := registers{
			"a": 3,
			"b": 88,
			"c": 65,
			"d": 12,
		}
		r.inc("c")
		assert.Equal(t, want, r)
	})
}

func Test_registers_dec(t *testing.T) {
	t.Run("increments the value of the given register by 1", func(t *testing.T) {
		r := registers{
			"a": 3,
			"b": 88,
			"c": 64,
			"d": 12,
		}
		want := registers{
			"a": 3,
			"b": 88,
			"c": 64,
			"d": 11,
		}
		r.dec("d")
		assert.Equal(t, want, r)
	})
}

func Test_registers_jnz(t *testing.T) {
	type args struct {
		reg  string
		jump string
		i    int
	}
	tests := []struct {
		name               string
		r                  registers
		args               args
		want               registers
		want1              int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "increments i and returns if the value of a given register is 0",
			r: registers{
				"a": 2,
				"b": 0,
				"c": 3,
				"d": 99,
			},
			args: args{
				reg:  "b",
				jump: "2",
				i:    8,
			},
			want: registers{
				"a": 2,
				"b": 0,
				"c": 3,
				"d": 99,
			},
			want1:              9,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns an error if the given reg is not in registers and cannot be converted to an int",
			r: registers{
				"a": 2,
				"b": 0,
				"c": 3,
				"d": 99,
			},
			args: args{
				reg:  "ttt",
				jump: "2",
				i:    8,
			},
			want: registers{
				"a": 2,
				"b": 0,
				"c": 3,
				"d": 99,
			},
			want1:              8,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "increments i by 1 and returns if given reg is 0",
			r: registers{
				"a": 2,
				"b": 0,
				"c": 3,
				"d": 99,
			},
			args: args{
				reg:  "0",
				jump: "2",
				i:    83,
			},
			want: registers{
				"a": 2,
				"b": 0,
				"c": 3,
				"d": 99,
			},
			want1:              84,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns an error if reg is in registers, not 0, and jmp cannot be converted to int",
			r: registers{
				"a": 2,
				"b": 0,
				"c": 3,
				"d": 99,
			},
			args: args{
				reg:  "a",
				jump: "two",
				i:    2,
			},
			want: registers{
				"a": 2,
				"b": 0,
				"c": 3,
				"d": 99,
			},
			want1:              2,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "increments i by given value (positive) if reg is in registers and not 0",
			r: registers{
				"a": 2,
				"b": 0,
				"c": 3,
				"d": 99,
			},
			args: args{
				reg:  "a",
				jump: "56",
				i:    41,
			},
			want: registers{
				"a": 2,
				"b": 0,
				"c": 3,
				"d": 99,
			},
			want1:              97,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "increments i by given value (negative) if reg is in registers and not 0",
			r: registers{
				"a": 2,
				"b": 0,
				"c": 3,
				"d": 99,
			},
			args: args{
				reg:  "a",
				jump: "-5",
				i:    41,
			},
			want: registers{
				"a": 2,
				"b": 0,
				"c": 3,
				"d": 99,
			},
			want1:              36,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			err := r.jnz(tt.args.reg, tt.args.jump, &tt.args.i)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want1, tt.args.i)
			assert.Equal(t, tt.want, r)
		})
	}
}

func Test_registers_followInstruction(t *testing.T) {
	type args struct {
		inst string
		i    int
	}
	tests := []struct {
		name               string
		r                  registers
		args               args
		want               registers
		want1              int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if cpy returns an error",
			r:    registers{"a": 2, "b": 4, "c": 8, "d": 0},
			args: args{
				inst: "cpy six a",
				i:    3,
			},
			want:               registers{"a": 2, "b": 4, "c": 8, "d": 0},
			want1:              4,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if jnz returns an error",
			r:    registers{"a": 2, "b": 4, "c": 8, "d": 0},
			args: args{
				inst: "jnz six a",
				i:    3,
			},
			want:               registers{"a": 2, "b": 4, "c": 8, "d": 0},
			want1:              3,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "carries out cpy instruction",
			r:    registers{"a": 2, "b": 4, "c": 8, "d": 0},
			args: args{
				inst: "cpy 99 a",
				i:    34,
			},
			want:               registers{"a": 99, "b": 4, "c": 8, "d": 0},
			want1:              35,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "carries out inc instruction",
			r:    registers{"a": 2, "b": 4, "c": 8, "d": 0},
			args: args{
				inst: "inc d",
				i:    9,
			},
			want:               registers{"a": 2, "b": 4, "c": 8, "d": 1},
			want1:              10,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "carries out dec instruction",
			r:    registers{"a": 2, "b": 4, "c": 8, "d": 0},
			args: args{
				inst: "dec a",
				i:    8,
			},
			want:               registers{"a": 1, "b": 4, "c": 8, "d": 0},
			want1:              9,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "carries out jnz instruction",
			r:    registers{"a": 2, "b": 4, "c": 8, "d": 0},
			args: args{
				inst: "jnz b 24",
				i:    19,
			},
			want:               registers{"a": 2, "b": 4, "c": 8, "d": 0},
			want1:              43,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns an error if an invalid instruction is supplied",
			r:    registers{"a": 2, "b": 4, "c": 8, "d": 0},
			args: args{
				inst: "do something",
				i:    3,
			},
			want:               registers{"a": 2, "b": 4, "c": 8, "d": 0},
			want1:              3,
			errorAssertionFunc: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			err := r.followInstruction(tt.args.inst, &tt.args.i)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want1, tt.args.i)
			assert.Equal(t, tt.want, r)
		})
	}
}

func Test_registers_findSolution(t *testing.T) {
	type args struct {
		instructions []string
		i            int
	}
	tests := []struct {
		name               string
		r                  registers
		args               args
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if any instruction fails",
			r:    registers{"a": 0, "b": 0, "c": 0, "d": 0},
			args: args{
				instructions: []string{
					"cpy 41 a",
					"inc a",
					"inc a",
					"dec a",
					"jnz a 2",
					"dec a",
				},
				i: 0,
			},
			want:               42,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns the value of a after all instructions followed",
			r:    registers{"a": 0, "b": 0, "c": 0, "d": 0},
			args: args{
				instructions: []string{
					"cpy 41 a",
					"inc a",
					"inc a",
					"blah a",
					"jnz a 2",
					"dec a",
				},
				i: 0,
			},
			want:               -1,
			errorAssertionFunc: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.findSolution(tt.args.instructions, &tt.args.i)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_findSolutions(t *testing.T) {
	tests := []struct {
		name               string
		instructions       []string
		want               int
		want1              int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if an instruction can't be followed",
			instructions: []string{
				"cpy 41 a",
				"inc a",
				"inc a",
				"blah a",
				"jnz a 2",
				"dec a",
			},
			want:               -1,
			want1:              -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns correct parts 1 and 2, real solution (since AoC doesn't give a part 2 example)",
			instructions: []string{
				"cpy 1 a",
				"cpy 1 b",
				"cpy 26 d",
				"jnz c 2",
				"jnz 1 5",
				"cpy 7 c",
				"inc d",
				"dec c",
				"jnz c -2",
				"cpy a c",
				"inc a",
				"dec b",
				"jnz b -2",
				"cpy c b",
				"dec d",
				"jnz d -6",
				"cpy 13 c",
				"cpy 14 d",
				"inc a",
				"dec d",
				"jnz d -2",
				"dec c",
				"jnz c -5",
			},
			want:               317993,
			want1:              9227647,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := findSolutions(tt.instructions)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
