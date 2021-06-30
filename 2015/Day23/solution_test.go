package main

import (
	"reflect"
	"testing"
)

func TestRegisters_hlfInstruction(t *testing.T) {
	tests := []struct {
		name string
		r    Registers
		arg  string
		want Registers
	}{
		{
			name: "it halves the value of the given register",
			r:    Registers{"a": 43, "b": 86},
			arg:  "b",
			want: Registers{"a": 43, "b": 43},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.hlfInstruction(tt.arg)
			if !reflect.DeepEqual(tt.r, tt.want) {
				t.Errorf("Registers.hlfInstruction() = %v, want %v", tt.r, tt.want)
			}
		})
	}
}

func TestRegisters_tplInstruction(t *testing.T) {
	tests := []struct {
		name string
		r    Registers
		arg  string
		want Registers
	}{
		{
			name: "it triples the value of the given register",
			r:    Registers{"a": 43, "b": 86},
			arg:  "a",
			want: Registers{"a": 129, "b": 86},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.tplInstruction(tt.arg)
			if !reflect.DeepEqual(tt.r, tt.want) {
				t.Errorf("Registers.tplInstruction() = %v, want %v", tt.r, tt.want)
			}
		})
	}
}

func TestRegisters_incInstruction(t *testing.T) {
	tests := []struct {
		name string
		r    Registers
		arg  string
		want Registers
	}{
		{
			name: "it increments the value of the given register",
			r:    Registers{"a": 43, "b": 86},
			arg:  "a",
			want: Registers{"a": 44, "b": 86},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.incInstruction(tt.arg)
			if !reflect.DeepEqual(tt.r, tt.want) {
				t.Errorf("Registers.incInstruction() = %v, want %v", tt.r, tt.want)
			}
		})
	}
}

func TestRegisters_jmpInstruction(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    int
		wantErr bool
	}{
		{
			name:    "it returns an error if the offset cannot be converted into an int",
			arg:     "a",
			want:    -1,
			wantErr: true,
		},
		{
			name:    "it returns the given offset converted to int, without sign",
			arg:     "18",
			want:    18,
			wantErr: false,
		},
		{
			name:    "it returns the given offset converted to int, with positive sign",
			arg:     "+19",
			want:    19,
			wantErr: false,
		},
		{
			name:    "it returns the given offset converted to int, with negative sign",
			arg:     "-42",
			want:    -42,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Registers{}
			got, err := r.jmpInstruction(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Registers.jmpInstruction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Registers.jmpInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisters_jieInstruction(t *testing.T) {
	type args struct {
		register string
		offset   string
	}
	tests := []struct {
		name    string
		r       Registers
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "returns an error if the given offset cannot be converted to an int",
			r:       Registers{"a": 22, "b": 2},
			args:    args{register: "a", offset: "+as"},
			want:    -1,
			wantErr: true,
		},
		{
			name:    "returns offset of 1 if given register value is not even",
			r:       Registers{"a": 21, "b": 2},
			args:    args{register: "a", offset: "+9"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "returns a converted offset if the given register value is even",
			r:       Registers{"a": 22, "b": 2},
			args:    args{register: "a", offset: "+9"},
			want:    9,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.jieInstruction(tt.args.register, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Registers.jieInstruction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Registers.jieInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisters_jioInstruction(t *testing.T) {
	type args struct {
		register string
		offset   string
	}
	tests := []struct {
		name    string
		r       Registers
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "returns an error if the given offset cannot be converted to an int",
			r:       Registers{"a": 1, "b": 2},
			args:    args{register: "a", offset: "+as"},
			want:    -1,
			wantErr: true,
		},
		{
			name:    "returns offset of 1 if given register value is not 1",
			r:       Registers{"a": 21, "b": 2},
			args:    args{register: "a", offset: "+9"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "returns a converted offset if the given register value is 1",
			r:       Registers{"a": 1, "b": 2},
			args:    args{register: "a", offset: "+9"},
			want:    9,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.jioInstruction(tt.args.register, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Registers.jioInstruction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Registers.jioInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisters_FollowInstruction(t *testing.T) {
	tests := []struct {
		name    string
		r       Registers
		arg     string
		want    int
		want1   Registers
		wantErr bool
	}{
		{
			name:    `runs hlfInstructions when instruction starts with "hlf"`,
			r:       Registers{"a": 34, "b": 0},
			arg:     "hlf a",
			want:    1,
			want1:   Registers{"a": 17, "b": 0},
			wantErr: false,
		},
		{
			name:    `runs tplInstructions when instruction starts with "tpl"`,
			r:       Registers{"a": 34, "b": 0},
			arg:     "tpl a",
			want:    1,
			want1:   Registers{"a": 102, "b": 0},
			wantErr: false,
		},
		{
			name:    `runs incInstructions when instruction starts with "inc"`,
			r:       Registers{"a": 34, "b": 0},
			arg:     "inc a",
			want:    1,
			want1:   Registers{"a": 35, "b": 0},
			wantErr: false,
		},
		{
			name:    `runs jmpInstructions when instruction starts with "jmp"`,
			r:       Registers{"a": 34, "b": 0},
			arg:     "jmp -76",
			want:    -76,
			want1:   Registers{"a": 34, "b": 0},
			wantErr: false,
		},
		{
			name:    `runs jmpInstructions and returns error if offset can't be converted when instruction starts with "jmp"`,
			r:       Registers{"a": 34, "b": 0},
			arg:     "jmp aa",
			want:    -1,
			want1:   Registers{"a": 34, "b": 0},
			wantErr: true,
		},
		{
			name:    `runs jieInstructions when instruction starts with "jie" and register value is even`,
			r:       Registers{"a": 34, "b": 0},
			arg:     "jie a, -76",
			want:    -76,
			want1:   Registers{"a": 34, "b": 0},
			wantErr: false,
		},
		{
			name:    `runs jieInstructions when instruction starts with "jie" and register value is odd`,
			r:       Registers{"a": 35, "b": 0},
			arg:     "jie a, -76",
			want:    1,
			want1:   Registers{"a": 35, "b": 0},
			wantErr: false,
		},
		{
			name:    `runs jieInstructions and returns error if offset can't be converted when instruction starts with "jie"`,
			r:       Registers{"a": 34, "b": 0},
			arg:     "jie a, aa",
			want:    -1,
			want1:   Registers{"a": 34, "b": 0},
			wantErr: true,
		},
		{
			name:    `runs jioInstructions when instruction starts with "jio" and register value is 1`,
			r:       Registers{"a": 34, "b": 1},
			arg:     "jio b, -76",
			want:    -76,
			want1:   Registers{"a": 34, "b": 1},
			wantErr: false,
		},
		{
			name:    `runs jioInstructions when instruction starts with "jio" and register value is not 1`,
			r:       Registers{"a": 35, "b": 0},
			arg:     "jio b, -76",
			want:    1,
			want1:   Registers{"a": 35, "b": 0},
			wantErr: false,
		},
		{
			name:    `runs jioInstructions and returns error if offset can't be converted when instruction starts with "jio"`,
			r:       Registers{"a": 34, "b": 1},
			arg:     "jio b, aa",
			want:    -1,
			want1:   Registers{"a": 34, "b": 1},
			wantErr: true,
		},
		{
			name:    "returns an error if there input line contains no valid instruction",
			r:       Registers{"a": 34, "b": 1},
			arg:     "something b, aa",
			want:    -1,
			want1:   Registers{"a": 34, "b": 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.FollowInstruction(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Registers.FollowInstruction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Registers.FollowInstruction() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.r, tt.want1) {
				t.Errorf("Registers.FollowInstruction() r = %v, want %v", tt.r, tt.want1)
			}
		})
	}
}

func TestRegisters_RunInstructions(t *testing.T) {
	tests := []struct {
		name    string
		r       Registers
		arg     []string
		want    Registers
		wantErr bool
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
			want:    Registers{"a": 2, "b": 0},
			wantErr: false,
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
			want:    Registers{"a": 2, "b": 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.RunInstructions(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("Registers.RunInstructions() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.r, tt.want) {
				t.Errorf("Registers.RunInstructions() error = %v, wantErr %v", tt.r, tt.want)
			}
		})
	}
}
