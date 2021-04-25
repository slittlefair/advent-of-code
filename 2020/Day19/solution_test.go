package main

import (
	"reflect"
	"testing"
)

var adventOfCodeRules1 = map[string]Rule{
	"0": {
		subRules: [][]string{{"4", "1", "5"}},
	},
	"1": {
		subRules: [][]string{{"2", "3"}, {"3", "2"}},
	},
	"2": {
		subRules: [][]string{{"4", "4"}, {"5", "5"}},
	},
	"3": {
		subRules: [][]string{{"4", "5"}, {"5", "4"}},
	},
	"8": {
		subRules: [][]string{{"42"}},
	},
	"11": {
		subRules: [][]string{{"42", "31"}},
	},
	"4": {
		val: "a",
	},
	"5": {
		val: "b",
	},
}

var adventOfCodeMessages1 = []string{
	"ababbb",
	"bababa",
	"abbbab",
	"aaabbb",
	"aaaabbb",
}

func TestInput_parseInput(t *testing.T) {
	tests := []struct {
		name     string
		rawInput []string
		want     Input
	}{
		{
			name: "parses a simple input",
			rawInput: []string{
				"11: 42 31",
				"0: 4 1 5",
				"1: 2 3 | 3 2",
				"2: 4 4 | 5 5",
				"8: 42",
				"3: 4 5 | 5 4",
				"4: \"a\"",
				"5: \"b\"",
				"",
				"ababbb",
				"bababa",
				"abbbab",
				"aaabbb",
				"aaaabbb",
			},
			want: Input{
				Rules:    adventOfCodeRules1,
				Messages: adventOfCodeMessages1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Input{
				Rules: map[string]Rule{},
			}
			i.parseInput(tt.rawInput)
			if !reflect.DeepEqual(*i, tt.want) {
				t.Errorf("Input.parseInput() = %v, want %v", i, tt.want)
			}
		})
	}
}

func TestInput_iterateMessages(t *testing.T) {
	type fields struct {
		Rules    map[string]Rule
		Messages []string
	}
	type args struct {
		key            string
		remainingRules []string
		message        string
		index          int
		seen           map[string]bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]bool
	}{
		{
			name: "advent of code example 1",
			fields: fields{
				Rules:    adventOfCodeRules1,
				Messages: adventOfCodeMessages1,
			},
			args: args{
				key:            "4",
				remainingRules: []string{"1", "5"},
				message:        "ababbb",
				index:          0,
				seen:           map[string]bool{},
			},
			want: map[string]bool{"ababbb": true},
		},
		{
			name: "advent of code example 2",
			fields: fields{
				Rules:    adventOfCodeRules1,
				Messages: adventOfCodeMessages1,
			},
			args: args{
				key:            "4",
				remainingRules: []string{"1", "5"},
				message:        "bababa",
				index:          0,
				seen:           map[string]bool{},
			},
			want: map[string]bool{},
		},
		{
			name: "advent of code example 3",
			fields: fields{
				Rules:    adventOfCodeRules1,
				Messages: adventOfCodeMessages1,
			},
			args: args{
				key:            "4",
				remainingRules: []string{"1", "5"},
				message:        "abbbab",
				index:          0,
				seen:           map[string]bool{},
			},
			want: map[string]bool{"abbbab": true},
		},
		{
			name: "advent of code example 4",
			fields: fields{
				Rules:    adventOfCodeRules1,
				Messages: adventOfCodeMessages1,
			},
			args: args{
				key:            "4",
				remainingRules: []string{"1", "5"},
				message:        "aaabbb",
				index:          0,
				seen:           map[string]bool{},
			},
			want: map[string]bool{},
		},
		{
			name: "advent of code example 5",
			fields: fields{
				Rules:    adventOfCodeRules1,
				Messages: adventOfCodeMessages1,
			},
			args: args{
				key:            "4",
				remainingRules: []string{"1", "5"},
				message:        "aaabbb",
				index:          0,
				seen:           map[string]bool{"ababbb": true, "abbbab": true},
			},
			want: map[string]bool{"ababbb": true, "abbbab": true},
		},
		{
			name: "advent of code example 6",
			fields: fields{
				Rules:    adventOfCodeRules1,
				Messages: adventOfCodeMessages1,
			},
			args: args{
				key:            "4",
				remainingRules: []string{"1", "5"},
				message:        "a",
				index:          0,
				seen:           map[string]bool{"ababbb": true, "abbbab": true},
			},
			want: map[string]bool{"ababbb": true, "abbbab": true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Input{
				Rules:    tt.fields.Rules,
				Messages: tt.fields.Messages,
			}
			if got := i.iterateMessages(tt.args.key, tt.args.remainingRules, tt.args.message, tt.args.index, tt.args.seen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Input.iterateMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInput_changeRulesForPart2(t *testing.T) {
	type fields struct {
		Rules    map[string]Rule
		Messages []string
	}
	tests := []struct {
		name   string
		fields fields
		want   Input
	}{
		{
			name: "advent of code example 1",
			fields: fields{
				Rules: adventOfCodeRules1,
			},
			want: Input{
				Rules: map[string]Rule{
					"0": {
						subRules: [][]string{{"4", "1", "5"}},
					},
					"1": {
						subRules: [][]string{{"2", "3"}, {"3", "2"}},
					},
					"2": {
						subRules: [][]string{{"4", "4"}, {"5", "5"}},
					},
					"3": {
						subRules: [][]string{{"4", "5"}, {"5", "4"}},
					},
					"8": {
						subRules: [][]string{{"42"}, {"42", "8"}},
					},
					"11": {
						subRules: [][]string{{"42", "31"}, {"42", "11", "31"}},
					},
					"4": {
						val: "a",
					},
					"5": {
						val: "b",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Input{
				Rules:    tt.fields.Rules,
				Messages: tt.fields.Messages,
			}
			i.changeRulesForPart2()
			if !reflect.DeepEqual(i, tt.want) {
				t.Errorf("Input.parseInput() = %v, want %v", i, tt.want)
			}
		})
	}
}

func TestInput_evaluateMessages(t *testing.T) {
	type fields struct {
		Rules    map[string]Rule
		Messages []string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "advent of code example 1",
			fields: fields{
				Rules:    adventOfCodeRules1,
				Messages: adventOfCodeMessages1,
			},
			want: 2,
		},
		{
			name: "advent of code example 2",
			fields: fields{
				Rules: map[string]Rule{
					"42": {
						subRules: [][]string{{"9", "14"}, {"10", "1"}},
					},
					"9": {
						subRules: [][]string{{"14", "27"}, {"1", "26"}},
					},
					"10": {
						subRules: [][]string{{"23", "14"}, {"28", "1"}},
					},
					"1": {
						val: "a",
					},
					"11": {
						subRules: [][]string{{"42", "31"}},
					},
					"5": {
						subRules: [][]string{{"1", "14"}, {"15", "1"}},
					},
					"19": {
						subRules: [][]string{{"14", "1"}, {"14", "14"}},
					},
					"12": {
						subRules: [][]string{{"24", "14"}, {"19", "1"}},
					},
					"16": {
						subRules: [][]string{{"15", "1"}, {"14", "14"}},
					},
					"31": {
						subRules: [][]string{{"14", "17"}, {"1", "13"}},
					},
					"6": {
						subRules: [][]string{{"14", "14"}, {"1", "14"}},
					},
					"2": {
						subRules: [][]string{{"1", "24"}, {"14", "4"}},
					},
					"0": {
						subRules: [][]string{{"8", "11"}},
					},
					"13": {
						subRules: [][]string{{"14", "3"}, {"1", "12"}},
					},
					"15": {
						subRules: [][]string{{"1"}, {"14"}},
					},
					"17": {
						subRules: [][]string{{"14", "2"}, {"1", "7"}},
					},
					"23": {
						subRules: [][]string{{"25", "1"}, {"22", "14"}},
					},
					"28": {
						subRules: [][]string{{"16", "1"}},
					},
					"4": {
						subRules: [][]string{{"1", "1"}},
					},
					"20": {
						subRules: [][]string{{"14", "14"}, {"1", "15"}},
					},
					"3": {
						subRules: [][]string{{"5", "14"}, {"16", "1"}},
					},
					"27": {
						subRules: [][]string{{"1", "6"}, {"14", "18"}},
					},
					"14": {
						val: "b",
					},
					"21": {
						subRules: [][]string{{"14", "1"}, {"1", "14"}},
					},
					"25": {
						subRules: [][]string{{"1", "1"}, {"1", "14"}},
					},
					"22": {
						subRules: [][]string{{"14", "14"}},
					},
					"8": {
						subRules: [][]string{{"42"}},
					},
					"26": {
						subRules: [][]string{{"14", "22"}, {"1", "20"}},
					},
					"18": {
						subRules: [][]string{{"15", "15"}},
					},
					"7": {
						subRules: [][]string{{"14", "5"}, {"1", "21"}},
					},
					"24": {
						subRules: [][]string{{"14", "1"}},
					},
				},
				Messages: []string{
					"abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa",
					"bbabbbbaabaabba",
					"babbbbaabbbbbabbbbbbaabaaabaaa",
					"aaabbbbbbaaaabaababaabababbabaaabbababababaaa",
					"bbbbbbbaaaabbbbaaabbabaaa",
					"bbbababbbbaaaaaaaabbababaaababaabab",
					"ababaaaaaabaaab",
					"ababaaaaabbbaba",
					"baabbaaaabbaaaababbaababb",
					"abbbbabbbbaaaababbbbbbaaaababb",
					"aaaaabbaabaaaaababaa",
					"aaaabbaaaabbaaa",
					"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa",
					"babaaabbbaaabaababbaabababaaab",
					"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba",
				},
			},
			want: 3,
		},
		{
			name: "advent of code example 3",
			fields: fields{
				Rules: map[string]Rule{
					"42": {
						subRules: [][]string{{"9", "14"}, {"10", "1"}},
					},
					"9": {
						subRules: [][]string{{"14", "27"}, {"1", "26"}},
					},
					"10": {
						subRules: [][]string{{"23", "14"}, {"28", "1"}},
					},
					"1": {
						val: "a",
					},
					"11": {
						subRules: [][]string{{"42", "31"}, {"42", "11", "31"}},
					},
					"5": {
						subRules: [][]string{{"1", "14"}, {"15", "1"}},
					},
					"19": {
						subRules: [][]string{{"14", "1"}, {"14", "14"}},
					},
					"12": {
						subRules: [][]string{{"24", "14"}, {"19", "1"}},
					},
					"16": {
						subRules: [][]string{{"15", "1"}, {"14", "14"}},
					},
					"31": {
						subRules: [][]string{{"14", "17"}, {"1", "13"}},
					},
					"6": {
						subRules: [][]string{{"14", "14"}, {"1", "14"}},
					},
					"2": {
						subRules: [][]string{{"1", "24"}, {"14", "4"}},
					},
					"0": {
						subRules: [][]string{{"8", "11"}},
					},
					"13": {
						subRules: [][]string{{"14", "3"}, {"1", "12"}},
					},
					"15": {
						subRules: [][]string{{"1"}, {"14"}},
					},
					"17": {
						subRules: [][]string{{"14", "2"}, {"1", "7"}},
					},
					"23": {
						subRules: [][]string{{"25", "1"}, {"22", "14"}},
					},
					"28": {
						subRules: [][]string{{"16", "1"}},
					},
					"4": {
						subRules: [][]string{{"1", "1"}},
					},
					"20": {
						subRules: [][]string{{"14", "14"}, {"1", "15"}},
					},
					"3": {
						subRules: [][]string{{"5", "14"}, {"16", "1"}},
					},
					"27": {
						subRules: [][]string{{"1", "6"}, {"14", "18"}},
					},
					"14": {
						val: "b",
					},
					"21": {
						subRules: [][]string{{"14", "1"}, {"1", "14"}},
					},
					"25": {
						subRules: [][]string{{"1", "1"}, {"1", "14"}},
					},
					"22": {
						subRules: [][]string{{"14", "14"}},
					},
					"8": {
						subRules: [][]string{{"42"}, {"42", "8"}},
					},
					"26": {
						subRules: [][]string{{"14", "22"}, {"1", "20"}},
					},
					"18": {
						subRules: [][]string{{"15", "15"}},
					},
					"7": {
						subRules: [][]string{{"14", "5"}, {"1", "21"}},
					},
					"24": {
						subRules: [][]string{{"14", "1"}},
					},
				},
				Messages: []string{
					"abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa",
					"bbabbbbaabaabba",
					"babbbbaabbbbbabbbbbbaabaaabaaa",
					"aaabbbbbbaaaabaababaabababbabaaabbababababaaa",
					"bbbbbbbaaaabbbbaaabbabaaa",
					"bbbababbbbaaaaaaaabbababaaababaabab",
					"ababaaaaaabaaab",
					"ababaaaaabbbaba",
					"baabbaaaabbaaaababbaababb",
					"abbbbabbbbaaaababbbbbbaaaababb",
					"aaaaabbaabaaaaababaa",
					"aaaabbaaaabbaaa",
					"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa",
					"babaaabbbaaabaababbaabababaaab",
					"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba",
				},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Input{
				Rules:    tt.fields.Rules,
				Messages: tt.fields.Messages,
			}
			if got := i.evaluateMessages(); got != tt.want {
				t.Errorf("Input.evaluateMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}
