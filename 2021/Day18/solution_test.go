package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPair_leftPopulated(t *testing.T) {
	type fields struct {
		parent    *Pair
		leftPair  *Pair
		rightPair *Pair
		leftVal   *int
		rightVal  *int
	}
	val := 3
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "returns true if leftPair is not nil",
			fields: fields{
				leftPair: &Pair{},
			},
			want: true,
		},
		{
			name: "returns true if leftVal is not nil",
			fields: fields{
				leftVal: &val,
			},
			want: true,
		},
		{
			name: "returns true if leftPair and leftVal are not nil",
			fields: fields{
				leftPair: &Pair{},
				leftVal:  &val,
			},
			want: true,
		},
		{
			name: "returns false if leftPair and leftVal are not nil",
			fields: fields{
				rightPair: &Pair{
					parent:   &Pair{},
					leftPair: &Pair{},
				},
				rightVal: &val,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pair{
				parent:    tt.fields.parent,
				leftPair:  tt.fields.leftPair,
				rightPair: tt.fields.rightPair,
				leftVal:   tt.fields.leftVal,
				rightVal:  tt.fields.rightVal,
			}
			got := p.leftPopulated()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_parseLine(t *testing.T) {
	tests := []struct {
		name               string
		line               string
		want               *Pair
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "returns an error if the line does not end in a closing bracket",
			line:               "[[2,[2,[3,4]]],[2,3",
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "returns an error if the line contains a character that cannot be converted to int",
			line:               "[[2,[2,[3,4]]],[[4,8],[9,w]]]",
			want:               nil,
			errorAssertionFunc: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := 0
			got, err := parseLine(tt.line, &i)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name               string
		input              []string
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if parseLine returns an error",
			input: []string{
				"[[[[4,3],4],4],[7,[[8,4],9]]]",
				"[1,1]",
				"[9,1]",
				"[1,9]",
				"[0,a]",
				"[9,1]",
				"[1,9]",
			},
			want:               -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns correct magnitude of snailfish sum, advent of code example 1",
			input: []string{
				"[[[[4,3],4],4],[7,[[8,4],9]]]",
				"[1,1]",
			},
			want:               1384,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns correct magnitude of snailfish sum, advent of code example 2",
			input: []string{
				"[9,1]",
				"[1,9]",
			},
			want:               129,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns correct magnitude of snailfish sum, advent of code example 3",
			input: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
			},
			want:               445,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns correct magnitude of snailfish sum, advent of code example 4",
			input: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
				"[5,5]",
			},
			want:               791,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns correct magnitude of snailfish sum, advent of code example 5",
			input: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
				"[5,5]",
				"[6,6]",
			},
			want:               1137,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns correct magnitude of snailfish sum, advent of code example 6",
			input: []string{
				"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
				"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
				"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
				"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
				"[7,[5,[[3,8],[1,4]]]]",
				"[[2,[2,2]],[8,[8,1]]]",
				"[2,9]",
				"[1,[[[9,3],9],[[9,0],[0,7]]]]",
				"[[[5,[7,4]],7],1]",
				"[[[[4,2],2],6],[8,7]]",
			},
			want:               3488,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns correct magnitude of snailfish sum, advent of code example 7",
			input: []string{
				"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
				"[[[5,[2,8]],4],[5,[[9,9],0]]]",
				"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
				"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
				"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
				"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
				"[[[[5,4],[7,7]],8],[[8,3],8]]",
				"[[9,3],[[9,9],[6,[4,9]]]]",
				"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
				"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
			},
			want:               4140,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name               string
		input              []string
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if parseLine returns an error for first line of two considered",
			input: []string{
				"[[[0,[5,w]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
				"[[[5,[2,8]],4],[5,[[9,9],0]]]",
				"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
				"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
				"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
				"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
				"[[[[5,4],[7,7]],8],[[8,3],8]]",
				"[[9,3],[[9,9],[6,[4,9]]]]",
				"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
				"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
			},
			want:               -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if parseLine returns an error for second line of two considered",
			input: []string{
				"[[[0,[5,2]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
				"[[[5,[2,8]],4],[5,[[9,9],0]]]",
				"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
				"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
				"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
				"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
				"[[[[5,4],[7,7]],8],[[8,3],8]]",
				"[[9,3],[[9,9],[6,[4,9]]]",
				"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
				"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
			},
			want:               -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "finds greatest magnitude from sum of two snailfish numbers, advent of code example",
			input: []string{
				"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
				"[[[5,[2,8]],4],[5,[[9,9],0]]]",
				"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
				"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
				"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
				"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
				"[[[[5,4],[7,7]],8],[[8,3],8]]",
				"[[9,3],[[9,9],[6,[4,9]]]]",
				"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
				"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
			},
			want:               3993,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_findSolutions(t *testing.T) {
	tests := []struct {
		name               string
		input              []string
		want               int
		want1              int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if part1 returns an error, which happens when an input line cannot be parsed",
			input: []string{
				"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
				"[[[5,[2,8]],4],[5,[[9,9],0]]]",
				"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
				"[[[6,[0,7]],[0,9]],[4,[9,[9,x]]]]",
				"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
				"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
			},
			want:               -1,
			want1:              -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns part 1 and part 2 answers for given input, advent of code example",
			input: []string{
				"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
				"[[[5,[2,8]],4],[5,[[9,9],0]]]",
				"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
				"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
				"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
				"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
				"[[[[5,4],[7,7]],8],[[8,3],8]]",
				"[[9,3],[[9,9],[6,[4,9]]]]",
				"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
				"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
			},
			want:               4140,
			want1:              3993,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := findSolutions(tt.input)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
