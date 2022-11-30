package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWires_doBitwiseAND(t *testing.T) {
	type args struct {
		identifiers []string
		nums        []string
	}
	tests := []struct {
		name string
		w    Wires
		args args
		want Wires
	}{
		{
			name: "returns early if first identifier does not have a value in Wires",
			w: Wires{
				"c": 99,
				"d": 0,
			},
			args: args{
				identifiers: []string{"a", "b", "f"},
				nums:        []string{},
			},
			want: Wires{
				"c": 99,
				"d": 0,
			},
		},
		{
			name: "returns early if second identifier does not have a value in Wires",
			w: Wires{
				"c": 99,
				"d": 0,
				"a": 8,
			},
			args: args{
				identifiers: []string{"a", "b", "f"},
				nums:        []string{},
			},
			want: Wires{
				"c": 99,
				"d": 0,
				"a": 8,
			},
		},
		{
			name: "does bitwise AND of two identifiers and assigns it to a new value in Wires",
			w: Wires{
				"a": 2,
				"b": 3,
				"c": 5,
			},
			args: args{
				identifiers: []string{"a", "b", "f"},
				nums:        []string{},
			},
			want: Wires{
				"a": 2,
				"b": 3,
				"c": 5,
				"f": 2,
			},
		},
		{
			name: "does bitwise AND of two identifiers and assigns it to an existing value in Wires",
			w: Wires{
				"a": 2,
				"b": 3,
				"c": 5,
			},
			args: args{
				identifiers: []string{"a", "b", "b"},
				nums:        []string{},
			},
			want: Wires{
				"a": 2,
				"b": 2,
				"c": 5,
			},
		},
		{
			name: "does bitwise AND of two identifiers and assigns it to an existing value value in Wires",
			w: Wires{
				"a": 2,
				"b": 3,
				"c": 5,
			},
			args: args{
				identifiers: []string{"b", "ii"},
				nums:        []string{"7"},
			},
			want: Wires{
				"a":  2,
				"b":  3,
				"c":  5,
				"ii": 3,
			},
		},
		{
			name: "does bitwise AND of an identifier and a number and assigns it to an existing value in Wires",
			w: Wires{
				"a": 2,
				"b": 3,
				"c": 5,
			},
			args: args{
				identifiers: []string{"a", "c"},
				nums:        []string{"8"},
			},
			want: Wires{
				"a": 2,
				"b": 3,
				"c": 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.doBitwiseAND(tt.args.identifiers, tt.args.nums)
			assert.Equal(t, tt.want, tt.w)
		})
	}
}

func TestWires_doBitwiseOR(t *testing.T) {
	type args struct {
		identifiers []string
		nums        []string
	}
	tests := []struct {
		name string
		w    Wires
		args args
		want Wires
	}{
		{
			name: "returns early if first identifier does not have a value in Wires",
			w: Wires{
				"c": 99,
				"d": 0,
			},
			args: args{
				identifiers: []string{"a", "b", "f"},
				nums:        []string{},
			},
			want: Wires{
				"c": 99,
				"d": 0,
			},
		},
		{
			name: "returns early if second identifier does not have a value in Wires",
			w: Wires{
				"c": 99,
				"d": 0,
				"a": 8,
			},
			args: args{
				identifiers: []string{"a", "b", "f"},
				nums:        []string{},
			},
			want: Wires{
				"c": 99,
				"d": 0,
				"a": 8,
			},
		},
		{
			name: "does bitwise OR of two identifiers and assigns it to a new value in Wires",
			w: Wires{
				"a": 2,
				"b": 3,
				"c": 5,
			},
			args: args{
				identifiers: []string{"a", "b", "f"},
				nums:        []string{},
			},
			want: Wires{
				"a": 2,
				"b": 3,
				"c": 5,
				"f": 3,
			},
		},
		{
			name: "does bitwise OR of two identifiers and assigns it to an existing value in Wires",
			w: Wires{
				"a": 2,
				"b": 3,
				"c": 5,
			},
			args: args{
				identifiers: []string{"a", "b", "b"},
				nums:        []string{},
			},
			want: Wires{
				"a": 2,
				"b": 3,
				"c": 5,
			},
		},
		{
			name: "does bitwise OR of two identifiers and assigns it to a new value in Wires",
			w: Wires{
				"a": 2,
				"b": 3,
				"c": 5,
			},
			args: args{
				identifiers: []string{"b", "ii"},
				nums:        []string{"8"},
			},
			want: Wires{
				"a":  2,
				"b":  3,
				"c":  5,
				"ii": 11,
			},
		},
		{
			name: "does bitwise OR of an identifier and a number and assigns it to an existing value in Wires",
			w: Wires{
				"a": 2,
				"b": 3,
				"c": 5,
			},
			args: args{
				identifiers: []string{"a", "c"},
				nums:        []string{"8"},
			},
			want: Wires{
				"a": 2,
				"b": 3,
				"c": 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.doBitwiseOR(tt.args.identifiers, tt.args.nums)
			assert.Equal(t, tt.want, tt.w)
		})
	}
}

func TestWires_doBitwiseNOT(t *testing.T) {
	type args struct {
		identifiers []string
	}
	tests := []struct {
		name string
		w    Wires
		args args
		want Wires
	}{
		{
			name: "returns early if first identifier does not have a value in Wires",
			w: Wires{
				"c": 99,
				"d": 0,
			},
			args: args{
				identifiers: []string{"a", "b"},
			},
			want: Wires{
				"c": 99,
				"d": 0,
			},
		},
		{
			name: "does bitwise NOT of an identifier and assigns it to a new value in Wires",
			w: Wires{
				"a": 123,
				"b": 3,
				"c": 5,
			},
			args: args{
				identifiers: []string{"a", "t"},
			},
			want: Wires{
				"a": 123,
				"b": 3,
				"c": 5,
				"t": 65412,
			},
		},
		{
			name: "does bitwise NOT of an identifier and assigns it to an existing value in Wires",
			w: Wires{
				"a": 123,
				"b": 456,
				"c": 5,
			},
			args: args{
				identifiers: []string{"b", "c"},
			},
			want: Wires{
				"a": 123,
				"b": 456,
				"c": 65079,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.doBitwiseNOT(tt.args.identifiers)
			assert.Equal(t, tt.want, tt.w)
		})
	}
}

func TestWires_doBitwiseLSHIFT(t *testing.T) {
	type args struct {
		identifiers []string
		nums        []string
	}
	tests := []struct {
		name string
		w    Wires
		args args
		want Wires
	}{
		{
			name: "returns early if first identifier does not have a value in Wires",
			w: Wires{
				"c": 99,
				"d": 0,
			},
			args: args{
				identifiers: []string{"a", "b"},
			},
			want: Wires{
				"c": 99,
				"d": 0,
			},
		},
		{
			name: "does bitwise LSHIFT of an identifier and assigns it to a new value in Wires",
			w: Wires{
				"a": 123,
				"b": 3,
				"c": 5,
			},
			args: args{
				identifiers: []string{"a", "t"},
				nums:        []string{"2"},
			},
			want: Wires{
				"a": 123,
				"b": 3,
				"c": 5,
				"t": 492,
			},
		},
		{
			name: "does bitwise LSHIFT of an identifier and assigns it to an existing value in Wires",
			w: Wires{
				"a": 123,
				"b": 53,
				"c": 5,
			},
			args: args{
				identifiers: []string{"b", "c"},
				nums:        []string{"2"},
			},
			want: Wires{
				"a": 123,
				"b": 53,
				"c": 212,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.doBitwiseLSHIFT(tt.args.identifiers, tt.args.nums)
			assert.Equal(t, tt.want, tt.w)
		})
	}
}

func TestWires_doBitwiseRSHIFT(t *testing.T) {
	type args struct {
		identifiers []string
		nums        []string
	}
	tests := []struct {
		name string
		w    Wires
		args args
		want Wires
	}{
		{
			name: "returns early if first identifier does not have a value in Wires",
			w: Wires{
				"c": 99,
				"d": 0,
			},
			args: args{
				identifiers: []string{"a", "b"},
			},
			want: Wires{
				"c": 99,
				"d": 0,
			},
		},
		{
			name: "does bitwise RSHIFT of an identifier and assigns it to a new value in Wires",
			w: Wires{
				"a": 212,
				"b": 3,
				"c": 5,
			},
			args: args{
				identifiers: []string{"a", "t"},
				nums:        []string{"2"},
			},
			want: Wires{
				"a": 212,
				"b": 3,
				"c": 5,
				"t": 53,
			},
		},
		{
			name: "does bitwise RSHIFT of an identifier and assigns it to an existing value in Wires",
			w: Wires{
				"a": 123,
				"b": 16,
				"c": 16,
			},
			args: args{
				identifiers: []string{"b", "c"},
				nums:        []string{"3"},
			},
			want: Wires{
				"a": 123,
				"b": 16,
				"c": 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.doBitwiseRSHIFT(tt.args.identifiers, tt.args.nums)
			assert.Equal(t, tt.want, tt.w)
		})
	}
}

func TestWires_doASSIGN(t *testing.T) {
	type args struct {
		identifiers []string
		nums        []string
	}
	tests := []struct {
		name string
		w    Wires
		args args
		want Wires
	}{
		{
			name: "returns early if first identifier does not have a value in Wires",
			w: Wires{
				"c": 99,
				"d": 0,
			},
			args: args{
				identifiers: []string{"a", "b"},
			},
			want: Wires{
				"c": 99,
				"d": 0,
			},
		},
		{
			name: "returns early if identifier to ASSIGN a direct value to already has a value in Wires",
			w: Wires{
				"c":  99,
				"d":  0,
				"ii": 1,
			},
			args: args{
				identifiers: []string{"ii"},
				nums:        []string{"8"},
			},
			want: Wires{
				"c":  99,
				"d":  0,
				"ii": 1,
			},
		},
		{
			name: "does bitwise ASSIGN of one identifier to another and assigns it to a new value in Wires",
			w: Wires{
				"a": 2,
				"b": 3,
				"c": 5,
			},
			args: args{
				identifiers: []string{"a", "f"},
				nums:        []string{},
			},
			want: Wires{
				"a": 2,
				"b": 3,
				"c": 5,
				"f": 2,
			},
		},
		{
			name: "does bitwise ASSIGN of one identifier to another and assigns it to an existing value in Wires",
			w: Wires{
				"a": 2,
				"b": 3,
				"c": 5,
			},
			args: args{
				identifiers: []string{"a", "b"},
				nums:        []string{},
			},
			want: Wires{
				"a": 2,
				"b": 2,
				"c": 5,
			},
		},
		{
			name: "does bitwise ASSIGN of one identifier to a number and assigns it to a new value in Wires",
			w: Wires{
				"a": 2,
				"b": 3,
				"c": 5,
			},
			args: args{
				identifiers: []string{"ii"},
				nums:        []string{"8"},
			},
			want: Wires{
				"a":  2,
				"b":  3,
				"c":  5,
				"ii": 8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.doASSIGN(tt.args.identifiers, tt.args.nums)
			assert.Equal(t, tt.want, tt.w)
		})
	}
}

func TestWires_followInstructions(t *testing.T) {
	tests := []struct {
		name         string
		w            *Wires
		instructions []string
		want         *Wires
	}{
		{
			name: "follows a simple set of instructions, advent of code example",
			w:    &Wires{},
			instructions: []string{
				"123 -> x",
				"456 -> y",
				"x AND y -> d",
				"x OR y -> e",
				"x LSHIFT 2 -> f",
				"y RSHIFT 2 -> g",
				"NOT x -> h",
				"NOT y -> i",
			},
			want: &Wires{
				"d": 72,
				"e": 507,
				"f": 492,
				"g": 114,
				"h": 65412,
				"i": 65079,
				"x": 123,
				"y": 456,
			},
		},
		{
			name: "follows a simple set of instructions but having to loop over, advent of code example",
			w:    &Wires{},
			instructions: []string{
				"x OR y -> e",
				"x LSHIFT 2 -> f",
				"y RSHIFT 2 -> g",
				"NOT x -> h",
				"NOT y -> i",
				"123 -> x",
				"456 -> y",
				"x AND y -> d",
			},
			want: &Wires{
				"d": 72,
				"e": 507,
				"f": 492,
				"g": 114,
				"h": 65412,
				"i": 65079,
				"x": 123,
				"y": 456,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.followInstructions(tt.instructions)
			assert.Equal(t, tt.want, tt.w)
		})
	}
}
