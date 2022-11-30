package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisters_hlfInstruction(t *testing.T) {
	t.Run("it halves the value of the given register", func(t *testing.T) {
		r := Registers{"a": 43, "b": 86}
		r.hlfInstruction("b")
		assert.Equal(t, Registers{"a": 43, "b": 43}, r)
	})
}

func TestRegisters_tplInstruction(t *testing.T) {
	t.Run("it triples the value of the given register", func(t *testing.T) {
		r := Registers{"a": 43, "b": 86}
		r.tplInstruction("a")
		assert.Equal(t, Registers{"a": 129, "b": 86}, r)
	})
}

func TestRegisters_incInstruction(t *testing.T) {
	t.Run("it increments the value of the given register", func(t *testing.T) {
		r := Registers{"a": 43, "b": 86}
		r.incInstruction("a")
		assert.Equal(t, Registers{"a": 44, "b": 86}, r)
	})
}

func TestRegisters_jmpInstruction(t *testing.T) {
	tests := []struct {
		name               string
		arg                string
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "it returns an error if the offset cannot be converted into an int",
			arg:                "a",
			want:               -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "it returns the given offset converted to int, without sign",
			arg:                "18",
			want:               18,
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "it returns the given offset converted to int, with positive sign",
			arg:                "+19",
			want:               19,
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "it returns the given offset converted to int, with negative sign",
			arg:                "-42",
			want:               -42,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Registers{}
			got, err := r.jmpInstruction(tt.arg)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRegisters_jieInstruction(t *testing.T) {
	type args struct {
		register string
		offset   string
	}
	tests := []struct {
		name               string
		r                  Registers
		args               args
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "returns an error if the given offset cannot be converted to an int",
			r:                  Registers{"a": 22, "b": 2},
			args:               args{register: "a", offset: "+as"},
			want:               -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "returns offset of 1 if given register value is not even",
			r:                  Registers{"a": 21, "b": 2},
			args:               args{register: "a", offset: "+9"},
			want:               1,
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "returns a converted offset if the given register value is even",
			r:                  Registers{"a": 22, "b": 2},
			args:               args{register: "a", offset: "+9"},
			want:               9,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.jieInstruction(tt.args.register, tt.args.offset)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRegisters_jioInstruction(t *testing.T) {
	type args struct {
		register string
		offset   string
	}
	tests := []struct {
		name               string
		r                  Registers
		args               args
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "returns an error if the given offset cannot be converted to an int",
			r:                  Registers{"a": 1, "b": 2},
			args:               args{register: "a", offset: "+as"},
			want:               -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "returns offset of 1 if given register value is not 1",
			r:                  Registers{"a": 21, "b": 2},
			args:               args{register: "a", offset: "+9"},
			want:               1,
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               "returns a converted offset if the given register value is 1",
			r:                  Registers{"a": 1, "b": 2},
			args:               args{register: "a", offset: "+9"},
			want:               9,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.jioInstruction(tt.args.register, tt.args.offset)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRegisters_FollowInstruction(t *testing.T) {
	tests := []struct {
		name               string
		r                  Registers
		arg                string
		want               int
		want1              Registers
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               `runs hlfInstructions when instruction starts with "hlf"`,
			r:                  Registers{"a": 34, "b": 0},
			arg:                "hlf a",
			want:               1,
			want1:              Registers{"a": 17, "b": 0},
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               `runs tplInstructions when instruction starts with "tpl"`,
			r:                  Registers{"a": 34, "b": 0},
			arg:                "tpl a",
			want:               1,
			want1:              Registers{"a": 102, "b": 0},
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               `runs incInstructions when instruction starts with "inc"`,
			r:                  Registers{"a": 34, "b": 0},
			arg:                "inc a",
			want:               1,
			want1:              Registers{"a": 35, "b": 0},
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               `runs jmpInstructions when instruction starts with "jmp"`,
			r:                  Registers{"a": 34, "b": 0},
			arg:                "jmp -76",
			want:               -76,
			want1:              Registers{"a": 34, "b": 0},
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               `runs jmpInstructions and returns error if offset can't be converted when instruction starts with "jmp"`,
			r:                  Registers{"a": 34, "b": 0},
			arg:                "jmp aa",
			want:               -1,
			want1:              Registers{"a": 34, "b": 0},
			errorAssertionFunc: assert.Error,
		},
		{
			name:               `runs jieInstructions when instruction starts with "jie" and register value is even`,
			r:                  Registers{"a": 34, "b": 0},
			arg:                "jie a, -76",
			want:               -76,
			want1:              Registers{"a": 34, "b": 0},
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               `runs jieInstructions when instruction starts with "jie" and register value is odd`,
			r:                  Registers{"a": 35, "b": 0},
			arg:                "jie a, -76",
			want:               1,
			want1:              Registers{"a": 35, "b": 0},
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               `runs jieInstructions and returns error if offset can't be converted when instruction starts with "jie"`,
			r:                  Registers{"a": 34, "b": 0},
			arg:                "jie a, aa",
			want:               -1,
			want1:              Registers{"a": 34, "b": 0},
			errorAssertionFunc: assert.Error,
		},
		{
			name:               `runs jioInstructions when instruction starts with "jio" and register value is 1`,
			r:                  Registers{"a": 34, "b": 1},
			arg:                "jio b, -76",
			want:               -76,
			want1:              Registers{"a": 34, "b": 1},
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               `runs jioInstructions when instruction starts with "jio" and register value is not 1`,
			r:                  Registers{"a": 35, "b": 0},
			arg:                "jio b, -76",
			want:               1,
			want1:              Registers{"a": 35, "b": 0},
			errorAssertionFunc: assert.NoError,
		},
		{
			name:               `runs jioInstructions and returns error if offset can't be converted when instruction starts with "jio"`,
			r:                  Registers{"a": 34, "b": 1},
			arg:                "jio b, aa",
			want:               -1,
			want1:              Registers{"a": 34, "b": 1},
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "returns an error if the input line contains no valid instruction",
			r:                  Registers{"a": 34, "b": 1},
			arg:                "something b, aa",
			want:               -1,
			want1:              Registers{"a": 34, "b": 1},
			errorAssertionFunc: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.FollowInstruction(tt.arg)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, tt.r)
		})
	}
}

func TestRegisters_RunInstructions(t *testing.T) {
	tests := []struct {
		name               string
		r                  Registers
		arg                []string
		want               Registers
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "runs a set of instructions",
			r:    Registers{"a": 0, "b": 0},
			arg: []string{
				"inc a",
				"jio a, +2",
				"tpl a",
				"inc a",
			},
			want:               Registers{"a": 2, "b": 0},
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns an error if running a set of instructions produces an error",
			r:    Registers{"a": 0, "b": 0},
			arg: []string{
				"inc a",
				"jio a, +2",
				"tpl a",
				"inc a",
				"blah blah",
			},
			want:               Registers{"a": 2, "b": 0},
			errorAssertionFunc: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.r.RunInstructions(tt.arg)
			tt.errorAssertionFunc(t, err)
		})
	}
}
