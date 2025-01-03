package main

import (
	"Advent-of-Code/regex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name               string
		line               string
		want               SignalPatterns
		want1              []string
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "returns an error if there are fewer than 14 matches",
			line:               "afgcb ecbdgfa bdgecf ad dgea cadfg afd afecbd cedagf fecgd | beafdc",
			want:               nil,
			want1:              nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name:               "returns an error if there are fewer than 14 matches",
			line:               "facge ecadgb fecdba dfcg fc fac eagfcd bgecfda afgeb eacdg | cdfg caefg fc bdeafcg fac",
			want:               nil,
			want1:              nil,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns a correct SignalPatterns and output value from a given line",
			line: "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
			want: SignalPatterns{
				"be":      {},
				"cfbegad": {},
				"cbdgef":  {},
				"fgaecd":  {},
				"cgeb":    {},
				"fdcge":   {},
				"agebfd":  {},
				"fecdb":   {},
				"fabcd":   {},
				"edb":     {},
			},
			want1:              []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseInput(tt.line, regex.MatchWords)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_countSimpleDigits(t *testing.T) {
	tests := []struct {
		name         string
		outputValues []string
		want         int
	}{
		{
			name:         "returns 0 if no output values are the correct lengths",
			outputValues: []string{"fdgabe", "cefdb", "cefbgd", "gacbe"},
			want:         0,
		},
		{
			name:         "returns 1 if 1 output value is 2 characters long",
			outputValues: []string{"fdgabe", "ce", "cefbgd", "gacbe"},
			want:         1,
		},
		{
			name:         "returns 1 if 1 output values is 3 characters long",
			outputValues: []string{"fdgabe", "cefdb", "gba", "gacbe"},
			want:         1,
		},
		{
			name:         "returns 1 if 1 output values is 4 characters long",
			outputValues: []string{"fdgabe", "cefdb", "gba", "gacbe"},
			want:         1,
		},
		{
			name:         "returns 1 if 1 output values is 7 characters long",
			outputValues: []string{"fdgbe", "cefdb", "gcbafe", "gacbdfe"},
			want:         1,
		},
		{
			name:         "returns correct number of output values at the correct lengths",
			outputValues: []string{"fdgb", "cefdb", "gcb", "gacbdfe"},
			want:         3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countSimpleDigits(tt.outputValues)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_stringsShareParts(t *testing.T) {
	type args struct {
		str1      string
		str2      string
		wantEqual bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns false if strings are expected to be equal length but are not",
			args: args{
				str1:      "abc",
				str2:      "cbda",
				wantEqual: true,
			},
			want: false,
		},
		{
			name: "returns false if str2 contains characters not in str1",
			args: args{
				str1:      "abce",
				str2:      "cbda",
				wantEqual: false,
			},
			want: false,
		},
		{
			name: "returns true if all characters in str2 are in str1",
			args: args{
				str1:      "abced",
				str2:      "cbda",
				wantEqual: false,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringsShareParts(tt.args.str1, tt.args.str2, tt.args.wantEqual)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestValueMap_find3Letter(t *testing.T) {
	tests := []struct {
		name               string
		vm                 ValueMap
		signalPatterns     SignalPatterns
		want               ValueMap
		want1              SignalPatterns
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if valueMap has no string at index 1",
			vm: ValueMap{
				"",
				"",
				"",
				"",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"",
			},
			signalPatterns: SignalPatterns{
				"cdfbe":  {},
				"gcdfa":  {},
				"fbcad":  {},
				"cefabd": {},
				"cdfgeb": {},
				"cagedb": {},
			},
			want: ValueMap{
				"",
				"",
				"",
				"",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"",
			},
			want1: SignalPatterns{
				"cdfbe":  {},
				"gcdfa":  {},
				"fbcad":  {},
				"cefabd": {},
				"cdfgeb": {},
				"cagedb": {},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if no value for 3 can be found",
			vm: ValueMap{
				"",
				"ab",
				"",
				"",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"",
			},
			signalPatterns: SignalPatterns{
				"cdfbe":  {},
				"gcdfa":  {},
				"fgcae":  {},
				"cefabd": {},
				"cdfgeb": {},
				"cagedb": {},
			},
			want: ValueMap{
				"",
				"ab",
				"",
				"",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"",
			},
			want1: SignalPatterns{
				"cdfbe":  {},
				"gcdfa":  {},
				"fgcae":  {},
				"cefabd": {},
				"cdfgeb": {},
				"cagedb": {},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "correctly assigns and deletes matching 3 letter string",
			vm: ValueMap{
				"",
				"ab",
				"",
				"",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"",
			},
			signalPatterns: SignalPatterns{
				"cdfbe":  {},
				"gcdfa":  {},
				"fbcad":  {},
				"cefabd": {},
				"cdfgeb": {},
				"cagedb": {},
			},
			want: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"",
			},
			want1: SignalPatterns{
				"cdfbe":  {},
				"gcdfa":  {},
				"cefabd": {},
				"cdfgeb": {},
				"cagedb": {},
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := tt.vm
			signalPatterns := tt.signalPatterns
			err := vm.find3Letter(signalPatterns)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, vm)
			assert.Equal(t, tt.want1, signalPatterns)
		})
	}
}

func TestValueMap_find9Letter(t *testing.T) {
	tests := []struct {
		name               string
		vm                 ValueMap
		signalPatterns     SignalPatterns
		want               ValueMap
		want1              SignalPatterns
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if valueMap has no string at index 3",
			vm: ValueMap{
				"",
				"ab",
				"",
				"",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"",
			},
			signalPatterns: SignalPatterns{
				"cdfbe":  {},
				"gcdfa":  {},
				"cefabd": {},
				"cdfgeb": {},
				"cagedb": {},
			},
			want: ValueMap{
				"",
				"ab",
				"",
				"",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"",
			},
			want1: SignalPatterns{
				"cdfbe":  {},
				"gcdfa":  {},
				"cefabd": {},
				"cdfgeb": {},
				"cagedb": {},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if no value for 9 can be found",
			vm: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"",
			},
			signalPatterns: SignalPatterns{
				"cdfbe":  {},
				"gcdfa":  {},
				"cefabg": {},
				"cdfgeb": {},
				"cagedb": {},
			},
			want: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"",
			},
			want1: SignalPatterns{
				"cdfbe":  {},
				"gcdfa":  {},
				"cefabg": {},
				"cdfgeb": {},
				"cagedb": {},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "correctly assigns and deletes matching 9 letter string",
			vm: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"",
			},
			signalPatterns: SignalPatterns{
				"cdfbe":  {},
				"gcdfa":  {},
				"cefabd": {},
				"cdfgeb": {},
				"cagedb": {},
			},
			want: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			want1: SignalPatterns{
				"cdfbe":  {},
				"gcdfa":  {},
				"cdfgeb": {},
				"cagedb": {},
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := tt.vm
			signalPatterns := tt.signalPatterns
			err := vm.find9Letter(signalPatterns)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, vm)
			assert.Equal(t, tt.want1, signalPatterns)
		})
	}
}

func TestValueMap_find5Letter(t *testing.T) {
	tests := []struct {
		name               string
		vm                 ValueMap
		signalPatterns     SignalPatterns
		want               ValueMap
		want1              SignalPatterns
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if valueMap has no string at index 9",
			vm: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"",
			},
			signalPatterns: SignalPatterns{
				"cdfbe":  {},
				"gcdfa":  {},
				"cefabd": {},
				"cdfgeb": {},
				"cagedb": {},
			},
			want: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"",
			},
			want1: SignalPatterns{
				"cdfbe":  {},
				"gcdfa":  {},
				"cefabd": {},
				"cdfgeb": {},
				"cagedb": {},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if no value for 5 can be found",
			vm: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			signalPatterns: SignalPatterns{
				"cgfba":  {},
				"gcdfa":  {},
				"cdfgeb": {},
				"cagedb": {},
			},
			want: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			want1: SignalPatterns{
				"cgfba":  {},
				"gcdfa":  {},
				"cdfgeb": {},
				"cagedb": {},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "correctly assigns and deletes matching 5 letter string",
			vm: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			signalPatterns: SignalPatterns{
				"cdfbe":  {},
				"gcdfa":  {},
				"cdfgeb": {},
				"cagedb": {},
			},
			want: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"cdfbe",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			want1: SignalPatterns{
				"gcdfa":  {},
				"cdfgeb": {},
				"cagedb": {},
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := tt.vm
			signalPatterns := tt.signalPatterns
			err := vm.find5Letter(signalPatterns)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, vm)
			assert.Equal(t, tt.want1, signalPatterns)
		})
	}
}

func TestValueMap_find2Letter(t *testing.T) {
	tests := []struct {
		name               string
		vm                 ValueMap
		signalPatterns     SignalPatterns
		want               ValueMap
		want1              SignalPatterns
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if there are no remaining 5 letter signal patterns",
			vm: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"cdfbe",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			signalPatterns: SignalPatterns{
				"gcdfab": {},
				"cdfgeb": {},
				"cagedb": {},
			},
			want: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"cdfbe",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			want1: SignalPatterns{
				"gcdfab": {},
				"cdfgeb": {},
				"cagedb": {},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if there is more than 1 remaining 5 letter signal patterns",
			vm: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"cdfbe",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			signalPatterns: SignalPatterns{
				"gcdfa":  {},
				"cdfgb":  {},
				"cagedb": {},
			},
			want: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"cdfbe",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			want1: SignalPatterns{
				"gcdfa":  {},
				"cdfgb":  {},
				"cagedb": {},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "correctly assigns and deletes matching 2 letter string",
			vm: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"cdfbe",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			signalPatterns: SignalPatterns{
				"gcdfa":  {},
				"cdfgeb": {},
				"cagedb": {},
			},
			want: ValueMap{
				"",
				"ab",
				"gcdfa",
				"fbcad",
				"eafb",
				"cdfbe",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			want1: SignalPatterns{
				"cdfgeb": {},
				"cagedb": {},
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := tt.vm
			signalPatterns := tt.signalPatterns
			err := vm.find2Letter(signalPatterns)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, vm)
			assert.Equal(t, tt.want1, signalPatterns)
		})
	}
}

func TestValueMap_find6Letter(t *testing.T) {
	tests := []struct {
		name               string
		vm                 ValueMap
		signalPatterns     SignalPatterns
		want               ValueMap
		want1              SignalPatterns
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if valueMap has no string at index 5",
			vm: ValueMap{
				"",
				"ab",
				"gcdfa",
				"fbcad",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			signalPatterns: SignalPatterns{
				"cdafeb": {},
				"cagedb": {},
			},
			want: ValueMap{
				"",
				"ab",
				"gcdfa",
				"fbcad",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			want1: SignalPatterns{
				"cdafeb": {},
				"cagedb": {},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if no value for 6 can be found",
			vm: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"cdfbe",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			signalPatterns: SignalPatterns{
				"cdafeg": {},
				"cagedb": {},
			},
			want: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"cdfbe",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			want1: SignalPatterns{
				"cdafeg": {},
				"cagedb": {},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "correctly assigns and deletes matching 2 letter string",
			vm: ValueMap{
				"",
				"ab",
				"gcdfa",
				"fbcad",
				"eafb",
				"cdfbe",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			signalPatterns: SignalPatterns{
				"cdfgeb": {},
				"cagedb": {},
			},
			want: ValueMap{
				"",
				"ab",
				"gcdfa",
				"fbcad",
				"eafb",
				"cdfbe",
				"cdfgeb",
				"dab",
				"acedgfb",
				"cefabd",
			},
			want1: SignalPatterns{
				"cagedb": {},
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := tt.vm
			signalPatterns := tt.signalPatterns
			err := vm.find6Letter(signalPatterns)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, vm)
			assert.Equal(t, tt.want1, signalPatterns)
		})
	}
}

func TestValueMap_assignValues(t *testing.T) {
	tests := []struct {
		name               string
		signalPatterns     SignalPatterns
		want               ValueMap
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if find3Letter returns an error",
			signalPatterns: SignalPatterns{
				"acedgfb": {},
				"cdfbe":   {},
				"gcdfa":   {},
				"fbcge":   {},
				"dab":     {},
				"cefabd":  {},
				"cdfgeb":  {},
				"eafb":    {},
				"cagedb":  {},
				"ab":      {},
			},
			want: ValueMap{
				"",
				"ab",
				"",
				"",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"",
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if find9Letter returns an error",
			signalPatterns: SignalPatterns{
				"acedgfb": {},
				"cdfbe":   {},
				"gcdfa":   {},
				"fbcad":   {},
				"dab":     {},
				"cefagg":  {},
				"cdfgeb":  {},
				"eafb":    {},
				"cagedb":  {},
				"ab":      {},
			},
			want: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"",
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if find5Letter returns an error",
			signalPatterns: SignalPatterns{
				"acedgfb": {},
				"cdfga":   {},
				"gcdfa":   {},
				"fbcad":   {},
				"dab":     {},
				"cefabd":  {},
				"cdfgeb":  {},
				"eafb":    {},
				"cagedb":  {},
				"ab":      {},
			},
			want: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if find2Letter returns an error",
			signalPatterns: SignalPatterns{
				"acedgfb": {},
				"cdfbe":   {},
				"ebgdfa":  {},
				"fbcad":   {},
				"dab":     {},
				"cefabd":  {},
				"cdfgeb":  {},
				"eafb":    {},
				"cagedb":  {},
				"ab":      {},
			},
			want: ValueMap{
				"",
				"ab",
				"",
				"fbcad",
				"eafb",
				"cdfbe",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if find6Letter returns an error",
			signalPatterns: SignalPatterns{
				"acedgfb": {},
				"cdfbe":   {},
				"gcdfa":   {},
				"fbcad":   {},
				"dab":     {},
				"cefabd":  {},
				"cdfgad":  {},
				"eafb":    {},
				"cagedb":  {},
				"ab":      {},
			},
			want: ValueMap{
				"",
				"ab",
				"gcdfa",
				"fbcad",
				"eafb",
				"cdfbe",
				"",
				"dab",
				"acedgfb",
				"cefabd",
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "correctly assigns values from signal patterns, advent of code example",
			signalPatterns: SignalPatterns{
				"acedgfb": {},
				"cdfbe":   {},
				"gcdfa":   {},
				"fbcad":   {},
				"dab":     {},
				"cefabd":  {},
				"cdfgeb":  {},
				"eafb":    {},
				"cagedb":  {},
				"ab":      {},
			},
			want: ValueMap{
				"cagedb",
				"ab",
				"gcdfa",
				"fbcad",
				"eafb",
				"cdfbe",
				"cdfgeb",
				"dab",
				"acedgfb",
				"cefabd",
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := make(ValueMap, 10)
			err := vm.assignValues(tt.signalPatterns)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, vm)
		})
	}
}

func TestValueMap_decodeOutputValue(t *testing.T) {
	t.Run("calculates output value code from value map, advent of code example", func(t *testing.T) {
		vMap := ValueMap{
			"cagedb",
			"ab",
			"gcdfa",
			"fbcad",
			"eafb",
			"cdfbe",
			"cdfgeb",
			"dab",
			"acedgfb",
			"cefabd",
		}
		got := vMap.decodeOutputValue([]string{"cdfeb", "fcadb", "cdfeb", "cdbaf"})
		assert.Equal(t, 5353, got)
	})
}

func Test_findSolution(t *testing.T) {
	tests := []struct {
		name               string
		input              []string
		want               int
		want1              int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if the input can't be passed correctly",
			input: []string{
				"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
				"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
				"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg fdcagb cbg",
			},
			want:               -1,
			want1:              -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if values cannot be correctly assigned",
			input: []string{
				"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
				"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
				"fgaebd cg bdaec gdafbe agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
			},
			want:               -1,
			want1:              -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns correct part 1 and part 2 solutions, advent of code example",
			input: []string{
				"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
				"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
				"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
				"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
				"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
				"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
				"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
				"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
				"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
				"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
			},
			want:               26,
			want1:              61229,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := findSolution(tt.input)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
