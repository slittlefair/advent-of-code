package main

import (
	"reflect"
	"regexp"
	"testing"
)

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name    string
		line    string
		want    SignalPatterns
		want1   []string
		wantErr bool
	}{
		{
			name:    "returns an error if there are fewer than 14 matches",
			line:    "afgcb ecbdgfa bdgecf ad dgea cadfg afd afecbd cedagf fecgd | beafdc",
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
		{
			name:    "returns an error if there are fewer than 14 matches",
			line:    "facge ecadgb fecdba dfcg fc fac eagfcd bgecfda afgeb eacdg | cdfg caefg fc bdeafcg fac",
			want:    nil,
			want1:   nil,
			wantErr: true,
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
			want1:   []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseInput(tt.line, regexp.MustCompile(`\w+`))
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseInput() got1 = %v, want %v", got1, tt.want1)
			}
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
			if got := countSimpleDigits(tt.outputValues); got != tt.want {
				t.Errorf("countSimpleDigits() = %v, want %v", got, tt.want)
			}
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
			if got := stringsShareParts(tt.args.str1, tt.args.str2, tt.args.wantEqual); got != tt.want {
				t.Errorf("stringsShareParts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValueMap_find3Letter(t *testing.T) {
	tests := []struct {
		name           string
		vm             ValueMap
		signalPatterns SignalPatterns
		want           ValueMap
		want1          SignalPatterns
		wantErr        bool
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
			wantErr: true,
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
			wantErr: true,
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := tt.vm
			signalPatterns := tt.signalPatterns
			if err := vm.find3Letter(signalPatterns); (err != nil) != tt.wantErr {
				t.Errorf("ValueMap.find3Letter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(vm, tt.want) {
				t.Errorf("ValueMap.find3Letter() vm = %v, want %v", vm, tt.want)
			}
			if !reflect.DeepEqual(signalPatterns, tt.want1) {
				t.Errorf("ValueMap.find3Letter() signalPatterns = %v, want %v", signalPatterns, tt.want1)
			}
		})
	}
}

func TestValueMap_find9Letter(t *testing.T) {
	tests := []struct {
		name           string
		vm             ValueMap
		signalPatterns SignalPatterns
		want           ValueMap
		want1          SignalPatterns
		wantErr        bool
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
			wantErr: true,
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
			wantErr: true,
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := tt.vm
			signalPatterns := tt.signalPatterns
			if err := vm.find9Letter(signalPatterns); (err != nil) != tt.wantErr {
				t.Errorf("ValueMap.find9Letter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(vm, tt.want) {
				t.Errorf("ValueMap.find9Letter() vm = %v, want %v", vm, tt.want)
			}
			if !reflect.DeepEqual(signalPatterns, tt.want1) {
				t.Errorf("ValueMap.find9Letter() signalPatterns = %v, want %v", signalPatterns, tt.want1)
			}
		})
	}
}

func TestValueMap_find5Letter(t *testing.T) {
	tests := []struct {
		name           string
		vm             ValueMap
		signalPatterns SignalPatterns
		want           ValueMap
		want1          SignalPatterns
		wantErr        bool
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
			wantErr: true,
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
			wantErr: true,
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := tt.vm
			signalPatterns := tt.signalPatterns
			if err := vm.find5Letter(signalPatterns); (err != nil) != tt.wantErr {
				t.Errorf("ValueMap.find5Letter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(vm, tt.want) {
				t.Errorf("ValueMap.find5Letter() vm = %v, want %v", vm, tt.want)
			}
			if !reflect.DeepEqual(signalPatterns, tt.want1) {
				t.Errorf("ValueMap.find5Letter() signalPatterns = %v, want %v", signalPatterns, tt.want1)
			}
		})
	}
}

func TestValueMap_find2Letter(t *testing.T) {
	tests := []struct {
		name           string
		vm             ValueMap
		signalPatterns SignalPatterns
		want           ValueMap
		want1          SignalPatterns
		wantErr        bool
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
			wantErr: true,
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
			wantErr: true,
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := tt.vm
			signalPatterns := tt.signalPatterns
			if err := vm.find2Letter(signalPatterns); (err != nil) != tt.wantErr {
				t.Errorf("ValueMap.find5Letter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(vm, tt.want) {
				t.Errorf("ValueMap.find5Letter() vm = %v, want %v", vm, tt.want)
			}
			if !reflect.DeepEqual(signalPatterns, tt.want1) {
				t.Errorf("ValueMap.find5Letter() signalPatterns = %v, want %v", signalPatterns, tt.want1)
			}
		})
	}
}

func TestValueMap_find6Letter(t *testing.T) {
	tests := []struct {
		name           string
		vm             ValueMap
		signalPatterns SignalPatterns
		want           ValueMap
		want1          SignalPatterns
		wantErr        bool
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
			wantErr: true,
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
			wantErr: true,
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := tt.vm
			signalPatterns := tt.signalPatterns
			if err := vm.find6Letter(signalPatterns); (err != nil) != tt.wantErr {
				t.Errorf("ValueMap.find6Letter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(vm, tt.want) {
				t.Errorf("ValueMap.find6Letter() vm = %v, want %v", vm, tt.want)
			}
			if !reflect.DeepEqual(signalPatterns, tt.want1) {
				t.Errorf("ValueMap.find6Letter() signalPatterns = %v, want %v", signalPatterns, tt.want1)
			}
		})
	}
}

func TestValueMap_assignValues(t *testing.T) {
	tests := []struct {
		name           string
		signalPatterns SignalPatterns
		want           ValueMap
		wantErr        bool
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
			wantErr: true,
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
			wantErr: true,
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
			wantErr: true,
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
			wantErr: true,
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
			wantErr: true,
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := make(ValueMap, 10)
			if err := vm.assignValues(tt.signalPatterns); (err != nil) != tt.wantErr {
				t.Errorf("ValueMap.assignValues() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(vm, tt.want) {
				t.Errorf("ValueMap.assignValues() vm = %v, want %v", vm, tt.want)
			}
		})
	}
}

func TestValueMap_decodeOutputValue(t *testing.T) {
	tests := []struct {
		name        string
		vMap        ValueMap
		outputValue []string
		want        int
	}{
		{
			name: "calculates output value code from value map, advent of code example",
			vMap: ValueMap{
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
			outputValue: []string{"cdfeb", "fcadb", "cdfeb", "cdbaf"},
			want:        5353,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vMap.decodeOutputValue(tt.outputValue); got != tt.want {
				t.Errorf("ValueMap.decodeOutputValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSolution(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    int
		want1   int
		wantErr bool
	}{
		{
			name: "returns an error if the input can't be passed correctly",
			input: []string{
				"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
				"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
				"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg fdcagb cbg",
			},
			want:    -1,
			want1:   -1,
			wantErr: true,
		},
		{
			name: "returns an error if values cannot be correctly assigned",
			input: []string{
				"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
				"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
				"fgaebd cg bdaec gdafbe agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
			},
			want:    -1,
			want1:   -1,
			wantErr: true,
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
			want:    26,
			want1:   61229,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := findSolution(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("findSolution() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findSolution() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findSolution() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
