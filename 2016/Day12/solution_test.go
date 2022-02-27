package main

import (
	"reflect"
	"testing"
)

func Test_registers_cpy(t *testing.T) {
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		name    string
		r       registers
		args    args
		want    registers
		wantErr bool
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
			wantErr: true,
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
			wantErr: false,
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			if err := r.cpy(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("registers.cpy() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(r, tt.want) {
				t.Errorf("registers.cpy() = %v, want %v", r, tt.want)
			}
		})
	}
}

func Test_registers_inc(t *testing.T) {
	tests := []struct {
		name string
		r    registers
		reg  string
		want registers
	}{
		{
			name: "increments the value of the given register by 1",
			r: registers{
				"a": 3,
				"b": 88,
				"c": 64,
				"d": 12,
			},
			reg: "c",
			want: registers{
				"a": 3,
				"b": 88,
				"c": 65,
				"d": 12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			r.inc(tt.reg)
			if !reflect.DeepEqual(r, tt.want) {
				t.Errorf("registers.inc() = %v, want %v", r, tt.want)
			}
		})
	}
}

func Test_registers_dec(t *testing.T) {
	tests := []struct {
		name string
		r    registers
		reg  string
		want registers
	}{
		{
			name: "increments the value of the given register by 1",
			r: registers{
				"a": 3,
				"b": 88,
				"c": 64,
				"d": 12,
			},
			reg: "d",
			want: registers{
				"a": 3,
				"b": 88,
				"c": 64,
				"d": 11,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			r.dec(tt.reg)
			if !reflect.DeepEqual(r, tt.want) {
				t.Errorf("registers.dec() = %v, want %v", r, tt.want)
			}
		})
	}
}

func Test_registers_jnz(t *testing.T) {
	type args struct {
		reg  string
		jump string
		i    int
	}
	tests := []struct {
		name    string
		r       registers
		args    args
		want    registers
		want1   int
		wantErr bool
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
			want1:   9,
			wantErr: false,
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
			want1:   8,
			wantErr: true,
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
			want1:   84,
			wantErr: false,
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
			want1:   2,
			wantErr: true,
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
			want1:   97,
			wantErr: false,
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
			want1:   36,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			if err := r.jnz(tt.args.reg, tt.args.jump, &tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("registers.jnz() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.args.i != tt.want1 {
				t.Errorf("registers.jnz() i = %d, want %d", tt.args.i, tt.want1)
			}
			if !reflect.DeepEqual(r, tt.want) {
				t.Errorf("registers.jnz() = %v, want %v", r, tt.want)
			}
		})
	}
}

func Test_registers_followInstruction(t *testing.T) {
	type args struct {
		inst string
		i    int
	}
	tests := []struct {
		name    string
		r       registers
		args    args
		want    registers
		want1   int
		wantErr bool
	}{
		{
			name: "returns an error if cpy returns an error",
			r:    registers{"a": 2, "b": 4, "c": 8, "d": 0},
			args: args{
				inst: "cpy six a",
				i:    3,
			},
			want:    registers{"a": 2, "b": 4, "c": 8, "d": 0},
			want1:   4,
			wantErr: true,
		},
		{
			name: "returns an error if jnz returns an error",
			r:    registers{"a": 2, "b": 4, "c": 8, "d": 0},
			args: args{
				inst: "jnz six a",
				i:    3,
			},
			want:    registers{"a": 2, "b": 4, "c": 8, "d": 0},
			want1:   3,
			wantErr: true,
		},
		{
			name: "carries out cpy instruction",
			r:    registers{"a": 2, "b": 4, "c": 8, "d": 0},
			args: args{
				inst: "cpy 99 a",
				i:    34,
			},
			want:    registers{"a": 99, "b": 4, "c": 8, "d": 0},
			want1:   35,
			wantErr: false,
		},
		{
			name: "carries out inc instruction",
			r:    registers{"a": 2, "b": 4, "c": 8, "d": 0},
			args: args{
				inst: "inc d",
				i:    9,
			},
			want:    registers{"a": 2, "b": 4, "c": 8, "d": 1},
			want1:   10,
			wantErr: false,
		},
		{
			name: "carries out dec instruction",
			r:    registers{"a": 2, "b": 4, "c": 8, "d": 0},
			args: args{
				inst: "dec a",
				i:    8,
			},
			want:    registers{"a": 1, "b": 4, "c": 8, "d": 0},
			want1:   9,
			wantErr: false,
		},
		{
			name: "carries out jnz instruction",
			r:    registers{"a": 2, "b": 4, "c": 8, "d": 0},
			args: args{
				inst: "jnz b 24",
				i:    19,
			},
			want:    registers{"a": 2, "b": 4, "c": 8, "d": 0},
			want1:   43,
			wantErr: false,
		},
		{
			name: "returns an error if an invalid instruction is supplied",
			r:    registers{"a": 2, "b": 4, "c": 8, "d": 0},
			args: args{
				inst: "do something",
				i:    3,
			},
			want:    registers{"a": 2, "b": 4, "c": 8, "d": 0},
			want1:   3,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			if err := r.followInstruction(tt.args.inst, &tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("registers.followInstruction() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.args.i != tt.want1 {
				t.Errorf("registers.followInstruction() i = %d, want %d", tt.args.i, tt.want1)
			}
			if !reflect.DeepEqual(r, tt.want) {
				t.Errorf("registers.followInstruction() = %v, want %v", r, tt.want)
			}
		})
	}
}

func Test_registers_findSolution(t *testing.T) {
	type args struct {
		instructions []string
		i            int
	}
	tests := []struct {
		name    string
		r       registers
		args    args
		want    int
		wantErr bool
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
			want:    42,
			wantErr: false,
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
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.findSolution(tt.args.instructions, &tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("registers.findSolution() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("registers.findSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSolutions(t *testing.T) {
	tests := []struct {
		name         string
		instructions []string
		want         int
		want1        int
		wantErr      bool
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
			want:    -1,
			want1:   -1,
			wantErr: true,
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
			want:    317993,
			want1:   9227647,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := findSolutions(tt.instructions)
			if (err != nil) != tt.wantErr {
				t.Errorf("findSolutions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findSolutions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findSolutions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
