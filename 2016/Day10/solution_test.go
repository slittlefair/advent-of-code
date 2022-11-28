package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bot_addValue(t *testing.T) {
	tests := []struct {
		name string
		b    *bot
		v    int
		want *bot
	}{
		{
			name: "adds a value to a bot without vals",
			b: &bot{
				id:   "bot 1",
				vals: []int{},
			},
			v: 3,
			want: &bot{
				id:   "bot 1",
				vals: []int{3},
			},
		},
		{
			name: "adds a value to a bot with another, lesser, val",
			b: &bot{
				id:   "bot 1",
				vals: []int{4},
			},
			v: 9,
			want: &bot{
				id:   "bot 1",
				vals: []int{4, 9},
			},
		},
		{
			name: "adds a value to a bot with another, greater, val",
			b: &bot{
				id:   "bot 1",
				vals: []int{87},
			},
			v: 24,
			want: &bot{
				id:   "bot 1",
				vals: []int{24, 87},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.b
			b.addValue(tt.v)
			assert.Equal(t, tt.want, b)
		})
	}
}

func Test_bots_giveValue(t *testing.T) {
	type args struct {
		b         string
		recLowID  string
		recHighID string
	}
	tests := []struct {
		name string
		bs   bots
		args args
		want bots
	}{
		{
			name: "gives low and high values to bots that do exist",
			bs: bots{
				"bot 99": &bot{id: "bot 99", vals: []int{0, 46}},
				"bot 5":  &bot{id: "bot 5", vals: []int{2, 9}},
				"bot 6":  &bot{id: "bot 6", vals: []int{33}},
				"bot 0":  &bot{id: "bot 0", vals: []int{6}},
				"bot 1":  &bot{id: "bot 1", vals: []int{}},
			},
			args: args{
				b:         "bot 99",
				recLowID:  "bot 6",
				recHighID: "bot 1",
			},
			want: bots{
				"bot 99": &bot{id: "bot 99", vals: []int{}},
				"bot 5":  &bot{id: "bot 5", vals: []int{2, 9}},
				"bot 6":  &bot{id: "bot 6", vals: []int{0, 33}},
				"bot 0":  &bot{id: "bot 0", vals: []int{6}},
				"bot 1":  &bot{id: "bot 1", vals: []int{46}},
			},
		},
		{
			name: "gives low and high values to bots that don't exist",
			bs: bots{
				"bot 99": &bot{id: "bot 99", vals: []int{0, 46}},
				"bot 5":  &bot{id: "bot 5", vals: []int{3, 6}},
			},
			args: args{
				b:         "bot 5",
				recLowID:  "bot 6",
				recHighID: "bot 0",
			},
			want: bots{
				"bot 99": &bot{id: "bot 99", vals: []int{0, 46}},
				"bot 5":  &bot{id: "bot 5", vals: []int{}},
				"bot 6":  &bot{id: "bot 6", vals: []int{3}},
				"bot 0":  &bot{id: "bot 0", vals: []int{6}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := tt.bs
			bs.giveValue(bs[tt.args.b], tt.args.recLowID, tt.args.recHighID)
			assert.Len(t, bs, len(tt.want))
			assert.Equal(t, tt.want, bs)
		})
	}
}

func Test_bots_add(t *testing.T) {
	t.Run("adds a bot to bots", func(t *testing.T) {
		bs := bots{
			"bot 5": &bot{id: "bot 5", vals: []int{2, 9}},
			"bot 6": &bot{id: "bot 6", vals: []int{33, 0}},
			"bot 0": &bot{id: "bot 0", vals: []int{6}},
		}
		want := bots{
			"bot 5": &bot{id: "bot 5", vals: []int{2, 9}},
			"bot 6": &bot{id: "bot 6", vals: []int{33, 0}},
			"bot 0": &bot{id: "bot 0", vals: []int{6}},
			"bot 8": &bot{id: "bot 8", vals: []int{48}},
		}
		bs.add("bot 8", 48)
		assert.Len(t, bs, len(want))
		assert.Equal(t, want, bs)
	})
}

func Test_bots_handleValueLine(t *testing.T) {
	tests := []struct {
		name               string
		bs                 bots
		split              []string
		want               bots
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if split has fewer than 6 elements",
			bs: bots{
				"bot 1": &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2": &bot{id: "bot 2", vals: []int{89}},
			},
			split: []string{"22", "goes", "to", "bot", "9"},
			want: bots{
				"bot 1": &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2": &bot{id: "bot 2", vals: []int{89}},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if split has more than 6 elements",
			bs: bots{
				"bot 1": &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2": &bot{id: "bot 2", vals: []int{89}},
			},
			split: []string{"value", "22", "goes", "to", "bot", "9", "and", "bot", "2"},
			want: bots{
				"bot 1": &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2": &bot{id: "bot 2", vals: []int{89}},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if value cannot be converted to int",
			bs: bots{
				"bot 1": &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2": &bot{id: "bot 2", vals: []int{89}},
			},
			split: []string{"value", "twenty", "goes", "to", "bot", "9"},
			want: bots{
				"bot 1": &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2": &bot{id: "bot 2", vals: []int{89}},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "adds a new bot with the given value if it doesn't exist",
			bs: bots{
				"bot 1": &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2": &bot{id: "bot 2", vals: []int{89}},
			},
			split: []string{"value", "53", "goes", "to", "bot", "9"},
			want: bots{
				"bot 1": &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2": &bot{id: "bot 2", vals: []int{89}},
				"bot 9": &bot{id: "bot 9", vals: []int{53}},
			},
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "adds a value to the given bot if it exists",
			bs: bots{
				"bot 1": &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2": &bot{id: "bot 2", vals: []int{89}},
				"bot 9": &bot{id: "bot 9", vals: []int{53}},
			},
			split: []string{"value", "20", "goes", "to", "bot", "9"},
			want: bots{
				"bot 1": &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2": &bot{id: "bot 2", vals: []int{89}},
				"bot 9": &bot{id: "bot 9", vals: []int{20, 53}},
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := tt.bs
			err := bs.handleValueLine(tt.split)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, bs)
		})
	}
}

func Test_bots_handleGiveLine(t *testing.T) {
	tests := []struct {
		name               string
		bs                 bots
		split              []string
		want               bots
		want1              bool
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if split has fewer than 12 elements",
			bs: bots{
				"bot 1": &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2": &bot{id: "bot 2", vals: []int{89}},
				"bot 9": &bot{id: "bot 9", vals: []int{53}},
			},
			split: []string{"bot", "2", "gives", "low", "to", "output", "9"},
			want: bots{
				"bot 1": &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2": &bot{id: "bot 2", vals: []int{89}},
				"bot 9": &bot{id: "bot 9", vals: []int{53}},
			},
			want1:              false,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if split has more than 12 elements",
			bs: bots{
				"bot 1": &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2": &bot{id: "bot 2", vals: []int{89}},
				"bot 9": &bot{id: "bot 9", vals: []int{53}},
			},
			split: []string{"bot", "2", "gives", "low", "to", "output", "9", "and", "high", "to", "bot", "12", "also"},
			want: bots{
				"bot 1": &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2": &bot{id: "bot 2", vals: []int{89}},
				"bot 9": &bot{id: "bot 9", vals: []int{53}},
			},
			want1:              false,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns false and does nothing if bot doesn't exist",
			bs: bots{
				"bot 1": &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2": &bot{id: "bot 2", vals: []int{89}},
				"bot 9": &bot{id: "bot 9", vals: []int{53}},
			},
			split: []string{"bot", "23", "gives", "low", "to", "output", "9", "and", "high", "to", "bot", "12"},
			want: bots{
				"bot 1": &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2": &bot{id: "bot 2", vals: []int{89}},
				"bot 9": &bot{id: "bot 9", vals: []int{53}},
			},
			want1:              false,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns false and does nothing if bot doesn't have two values",
			bs: bots{
				"bot 1":  &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2":  &bot{id: "bot 2", vals: []int{89}},
				"bot 9":  &bot{id: "bot 9", vals: []int{53}},
				"bot 23": &bot{id: "bot 23", vals: []int{76}},
			},
			split: []string{"bot", "23", "gives", "low", "to", "output", "9", "and", "high", "to", "bot", "12"},
			want: bots{
				"bot 1":  &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2":  &bot{id: "bot 2", vals: []int{89}},
				"bot 9":  &bot{id: "bot 9", vals: []int{53}},
				"bot 23": &bot{id: "bot 23", vals: []int{76}},
			},
			want1:              false,
			errorAssertionFunc: assert.NoError,
		},
		{
			name: "returns true and gives values to given bots",
			bs: bots{
				"bot 1":  &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2":  &bot{id: "bot 2", vals: []int{89}},
				"bot 9":  &bot{id: "bot 9", vals: []int{53}},
				"bot 23": &bot{id: "bot 23", vals: []int{6, 90}},
			},
			split: []string{"bot", "23", "gives", "low", "to", "bot", "9", "and", "high", "to", "output", "12"},
			want: bots{
				"bot 1":     &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2":     &bot{id: "bot 2", vals: []int{89}},
				"bot 9":     &bot{id: "bot 9", vals: []int{6, 53}},
				"bot 23":    &bot{id: "bot 23", vals: []int{}},
				"output 12": &bot{id: "output 12", vals: []int{90}},
			},
			want1:              true,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := tt.bs
			got, err := bs.handleGiveLine(tt.split)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, bs)
			assert.Equal(t, tt.want1, got)
		})
	}
}

func Test_bots_getPart2Solution(t *testing.T) {
	tests := []struct {
		name string
		bs   bots
		want int
	}{
		{
			name: "returns products of values in outputs 0, 1 and 2",
			bs: bots{
				"bot 1":     &bot{id: "bot 1", vals: []int{7, 9}},
				"bot 2":     &bot{id: "bot 2", vals: []int{899}},
				"output 2":  &bot{id: "output 2", vals: []int{37}},
				"bot 9":     &bot{id: "bot 9", vals: []int{1, 53}},
				"bot 23":    &bot{id: "bot 23", vals: []int{}},
				"output 12": &bot{id: "output 12", vals: []int{90}},
				"bot 31":    &bot{id: "bot 31", vals: []int{6, 43}},
				"bot 22":    &bot{id: "bot 22", vals: []int{88}},
				"bot 16":    &bot{id: "bot 16", vals: []int{50}},
				"output 0":  &bot{id: "output 23", vals: []int{47}},
				"bot 12":    &bot{id: "bot 12", vals: []int{3, 9}},
				"output 1":  &bot{id: "output 1", vals: []int{89}},
			},
			want: 154771,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bs.getPart2Solution(); got != tt.want {
				t.Errorf("bots.getPart2Solution() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Run("returns products of values in outputs 0, 1 and 2", func(t *testing.T) {
		bs := bots{
			"bot 1":     &bot{id: "bot 1", vals: []int{7, 9}},
			"bot 2":     &bot{id: "bot 2", vals: []int{899}},
			"output 2":  &bot{id: "output 2", vals: []int{37}},
			"bot 9":     &bot{id: "bot 9", vals: []int{1, 53}},
			"bot 23":    &bot{id: "bot 23", vals: []int{}},
			"output 12": &bot{id: "output 12", vals: []int{90}},
			"bot 31":    &bot{id: "bot 31", vals: []int{6, 43}},
			"bot 22":    &bot{id: "bot 22", vals: []int{88}},
			"bot 16":    &bot{id: "bot 16", vals: []int{50}},
			"output 0":  &bot{id: "output 23", vals: []int{47}},
			"bot 12":    &bot{id: "bot 12", vals: []int{3, 9}},
			"output 1":  &bot{id: "output 1", vals: []int{89}},
		}
		got := bs.getPart2Solution()
		assert.Equal(t, 154771, got)
	})
}

func Test_findSolution(t *testing.T) {
	type args struct {
		input         []string
		expectedChips []int
	}
	tests := []struct {
		name               string
		args               args
		want               string
		want1              int
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if bots.handleLineValue returns an error",
			args: args{
				input: []string{
					"bot 147 gives low to bot 67 and high to bot 71",
					"bot 142 gives low to bot 128 and high to bot 164",
					"value two goes to bot 6",
					"bot 47 gives low to bot 4 and high to bot 209",
				},
				expectedChips: []int{2, 5},
			},
			want:               "",
			want1:              -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if bots.handleLineValue returns an error",
			args: args{
				input: []string{
					"bot 147 gives low to bot 67 and high to bot 71",
					"bot 142 gives low to bot 128 and high to bot 164",
					"value 2 goes to bot 6",
					"bot 47 gives low to bot 4 and high to bot 209 as well",
				},
				expectedChips: []int{2, 5},
			},
			want:               "",
			want1:              -1,
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns correct solutions to parts 1 and 2, advent of code example",
			args: args{
				input: []string{
					"value 5 goes to bot 2",
					"bot 2 gives low to bot 1 and high to bot 0",
					"value 3 goes to bot 1",
					"bot 1 gives low to output 1 and high to bot 0",
					"bot 0 gives low to output 2 and high to output 0",
					"value 2 goes to bot 2",
				},
				expectedChips: []int{2, 5},
			},
			want:               "2",
			want1:              30,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := findSolution(tt.args.input, tt.args.expectedChips)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
