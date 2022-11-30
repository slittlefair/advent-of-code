package main

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_containsAtLeast3Vowels(t *testing.T) {
	re := regexp.MustCompile(`a|e|i|o|u`)
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "returns true if string contains 3 'a's",
			str:  "bcvgayhgfhtfhtfahtfhtfahtfhft",
			want: true,
		},
		{
			name: "returns true if string contains 3 'e's",
			str:  "bcevgytgfhgfembvmhehvjhvh",
			want: true,
		},
		{
			name: "returns true if string contains 3 'i's",
			str:  "bcvgghgnvggvningvnvgninvgvngvdghdrtsiyt",
			want: true,
		},
		{
			name: "returns true if string contains 3 'o's",
			str:  "ngvngooobcvgyt",
			want: true,
		},
		{
			name: "returns true if string contains 3 'u's",
			str:  "uubcvgytu",
			want: true,
		},
		{
			name: "returns true if string contains multiple vowels",
			str:  "bcvgythfhitfthrdea",
			want: true,
		},
		{
			name: "returns false if string contains fewer than 3 vowels",
			str:  "bcvabvnbhmhvnvmhvmhvgvngvmhmhvinhvnhvjhvvngvngvgyt",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsAtLeast3Vowels(re, tt.str)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_containsDoubleLetter(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "returns true if string contains a double letter",
			str:  "fgfhfjgkgbgnndhdjcfnghcgcn",
			want: true,
		},
		{
			name: "returns false if string does not contain a double letter",
			str:  "fgfhfjgkgbgndhdjcfnghcgcn",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsDoubleLetter(tt.str)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_doesNotContainBadString(t *testing.T) {
	type args struct {
		re  *regexp.Regexp
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns true if string doesn't contain \"ab\" string",
			args: args{
				re:  regexp.MustCompile(`ab`),
				str: "ghgfhtfchcbbaajytj",
			},
			want: true,
		},
		{
			name: "returns true if string doesn't contain \"cd\" string",
			args: args{
				re:  regexp.MustCompile(`cd`),
				str: "ghgfhtfchcbbaajytj",
			},
			want: true,
		},
		{
			name: "returns true if string doesn't contain \"pq\" string",
			args: args{
				re:  regexp.MustCompile(`pq`),
				str: "ghgfhtfchcbbaajytj",
			},
			want: true,
		},
		{
			name: "returns true if string doesn't contain \"xy\" string",
			args: args{
				re:  regexp.MustCompile(`xy`),
				str: "ghgfhtfchcbbaajytj",
			},
			want: true,
		},
		{
			name: "returns false if string does contain \"ab\" string",
			args: args{
				re:  regexp.MustCompile(`ab`),
				str: "ghgfhabtfchcbbaajytj",
			},
			want: false,
		},
		{
			name: "returns false if string does contain \"cd\" string",
			args: args{
				re:  regexp.MustCompile(`cd`),
				str: "ghgfhtfchcbbaajycdtj",
			},
			want: false,
		},
		{
			name: "returns false if string does contain \"pq\" string",
			args: args{
				re:  regexp.MustCompile(`pq`),
				str: "ghgfhtfchcbbaajytjpq",
			},
			want: false,
		},
		{
			name: "returns false if string does contain \"xy\" string",
			args: args{
				re:  regexp.MustCompile(`xy`),
				str: "ghgfhtfchcbbaajytjxhhhxyhhhh",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := doesNotContainBadString(tt.args.re, tt.args.str)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_doesNotContainAnyBadStrings(t *testing.T) {
	reGroup := []*regexp.Regexp{
		regexp.MustCompile(`ab`),
		regexp.MustCompile(`cd`),
		regexp.MustCompile(`pq`),
		regexp.MustCompile(`xy`),
	}
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "returns false if str contains \"ab\"",
			str:  "gjfhgjshgfjdhabfkjghdkfgh",
			want: false,
		},
		{
			name: "returns false if str contains \"cd\"",
			str:  "gjfhgjshgfjdfghgfhgfhgfdcdfkjghdkfgh",
			want: false,
		},
		{
			name: "returns false if str contains \"pq\"",
			str:  "gjfhgjshgfjdoioioioiiooioioioiiooiioioiooiioqpiiioiopqkjghdkfgh",
			want: false,
		},
		{
			name: "returns false if str contains \"xy\"",
			str:  "gjfhgjshgfxyghdkfgh",
			want: false,
		},
		{
			name: "returns false if str doesn't contain bad strings ab, cd, pq, xy",
			str:  "gjfhgjshgfjdhfkjghdkfgh",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := doesNotContainAnyBadStrings(reGroup, tt.str)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_isNicePart1(t *testing.T) {
	vowelRe := regexp.MustCompile(`a|e|i|o|u`)
	badStringRe := []*regexp.Regexp{
		regexp.MustCompile(`ab`),
		regexp.MustCompile(`cd`),
		regexp.MustCompile(`pq`),
		regexp.MustCompile(`xy`),
	}
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "returns false if string doesn't contain 3 vowels",
			str:  "mmbmvjvjbbmhvmhvinhvhvjgvmhvhehvmhvhvngvnv",
			want: false,
		},
		{
			name: "returns false if string doesn't contain double letter",
			str:  "vbvbvbvbvbvbvbvbvbvivbvbvbvbvbvbvbvovbvbvbvbvbvbvbvbvbvbvavbvbvbvbvbvbvb",
			want: false,
		},
		{
			name: "returns false if string contains bad string ab",
			str:  "mmbmvjvjbbmhvmhvinhvhvjgvabmhvhehvmhvhvuungvnv",
			want: false,
		},
		{
			name: "returns false if string contains bad string cd",
			str:  "mmbmvjvjbbmhvmhvicdnhvhvjhvhehvmhvhvuungvnv",
			want: false,
		},
		{
			name: "returns false if string contains bad string pq",
			str:  "mmbmvjvjbbmhvmhvinhvhvjgvhvmhvpqhvuungvnv",
			want: false,
		},
		{
			name: "returns false if string contains bad string xy",
			str:  "mmbmvjvjbbmhvmhvinhvhvjgvxymhvhehvmhvhvuungvnv",
			want: false,
		},
		{
			name: "returns true if string contains more than two vowels, double letter and no bad strings",
			str:  "mnbmnbmnbmnbmnbmnbmnbmnbmnbmnbmnbmnbmnamnbmnbmnbmnbaamnbmnbmnb",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isNicePart1(vowelRe, badStringRe, tt.str)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_containsRepeatedPairOfLetters(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "returns true when there is a repeated pair, advent of code example 1",
			str:  "xyxy",
			want: true,
		},
		{
			name: "returns true when there is a repeated pair, advent of code example 2",
			str:  "aabcdefgaa",
			want: true,
		},
		{
			name: "returns false when there is no repeated pair, advent of code example 1",
			str:  "aaa",
			want: false,
		},
		{
			name: "returns false when there is no repeated pair",
			str:  "ddccbbaabcd",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsRepeatedPairOfLetters(tt.str)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_repeatsLetterWithGap(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "returns true if there is a repeated letter with gap, advent of code example 1",
			str:  "xyx",
			want: true,
		},
		{
			name: "returns true if there is a repeated letter with gap, advent of code example 2",
			str:  "abcdefeghi",
			want: true,
		},
		{
			name: "returns true if there is a repeated letter with gap, advent of code example 3",
			str:  "aaa",
			want: true,
		},
		{
			name: "returns false if there is a repeated letter with gap",
			str:  "abcabcabc",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := repeatsLetterWithGap(tt.str)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_isNicePart2(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "returns false if string doesn't contain repeated pairs of letters",
			str:  "aaa",
			want: false,
		},
		{
			name: "returns false if string doesn't contain repeated letters with gaps",
			str:  "abcabcabc",
			want: false,
		},
		{
			name: "returns true if string contains repeated pairs of letters and repeated letters with gaps",
			str:  "aabbaabbaba",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isNicePart2(tt.str)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getNiceStringCount(t *testing.T) {
	t.Run("returns count of nice strings from given input", func(t *testing.T) {
		input := []string{
			"ugknbfddgicrmopn",
			"haegwjzuvuyypxyu",
			"qjhvhtzxzqqjkmpb",
			"xxyxx",
			"aaa",
			"jchzalrnumimnmhp",
			"dvszwmarrgswjxmb",
			"aabbaabbababa",
		}
		got, got1 := getNiceStringCount(input)
		assert.Equal(t, 2, got)
		assert.Equal(t, 3, got1)
	})
}
