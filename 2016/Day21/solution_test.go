package main

import (
	"reflect"
	"testing"
)

func Test_swapPosition(t *testing.T) {
	type args struct {
		password []string
		i        int
		j        int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "swaps letters at given positions, advent of code example",
			args: args{
				password: []string{"a", "b", "c", "d", "e"},
				i:        0,
				j:        4,
			},
			want: []string{"e", "b", "c", "d", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPosition(tt.args.password, tt.args.i, tt.args.j); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swapLetter(t *testing.T) {
	type args struct {
		password []string
		x        string
		y        string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "swaps given letters, advent of code example",
			args: args{
				password: []string{"e", "b", "c", "d", "a"},
				x:        "b",
				y:        "d",
			},
			want: []string{"e", "d", "c", "b", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapLetter(tt.args.password, tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotateLeft(t *testing.T) {
	type args struct {
		password []string
		x        int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "rotates letters left by given number, advent of code example",
			args: args{
				password: []string{"e", "d", "c", "b", "a"},
				x:        1,
			},
			want: []string{"d", "c", "b", "a", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateLeft(tt.args.password, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotateRight(t *testing.T) {
	type args struct {
		password []string
		x        int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "rotates letters right by given number, advent of code example",
			args: args{
				password: []string{"e", "d", "c", "b", "a"},
				x:        1,
			},
			want: []string{"a", "e", "d", "c", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.password, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rotateByPosition(t *testing.T) {
	type args struct {
		password []string
		x        string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "rotates letters by index of given letter, less than 4, advent of code example",
			args: args{
				password: []string{"a", "b", "d", "e", "c"},
				x:        "b",
			},
			want: []string{"e", "c", "a", "b", "d"},
		},
		{
			name: "rotates letters by index of given letter, more than 4, advent of code example",
			args: args{
				password: []string{"e", "c", "a", "b", "d"},
				x:        "d",
			},
			want: []string{"d", "e", "c", "a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateByPosition(tt.args.password, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateByPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		password []string
		x        int
		y        int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "reverses letters through given odd positions, advent of code example",
			args: args{
				password: []string{"e", "d", "c", "b", "a"},
				x:        0,
				y:        4,
			},
			want: []string{"a", "b", "c", "d", "e"},
		},
		{
			name: "reverses letters through given even positions",
			args: args{
				password: []string{"e", "d", "c", "b", "a"},
				x:        1,
				y:        4,
			},
			want: []string{"e", "a", "b", "c", "d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.password, tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_move(t *testing.T) {
	type args struct {
		password []string
		x        int
		y        int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "moves element from given index to given index, advent of code example",
			args: args{
				password: []string{"b", "d", "e", "a", "c"},
				x:        3,
				y:        0,
			},
			want: []string{"a", "b", "d", "e", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := move(tt.args.password, tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("move() = %v, want %v", got, tt.want)
			}
		})
	}
}
