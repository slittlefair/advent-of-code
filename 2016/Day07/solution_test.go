package main

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var bab = [][]byte{
	[]byte("hah"),
	[]byte("vgv"),
	[]byte("hnh"),
	[]byte("sas"),
	[]byte("sds"),
	[]byte("yty"),
	[]byte("hnh"),
	[]byte("gfg"),
}

func Test_containsABBA(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "returns true if string is of the form abba",
			s:    "abba",
			want: true,
		},
		{
			name: "returns false if string does not contain abba",
			s:    "abdatgsksldiihald",
			want: false,
		},
		{
			name: "returns true if string does contain abba within it",
			s:    "abdatgskksldiihald",
			want: true,
		},
		{
			name: "returns false if string does contain abba but is in the form aaaa",
			s:    "abdatgskkkksldiihald",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsABBA(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_tlsValidation(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "returns true if abba outside brackets, advent of code example 1",
			s:    "abba[mnop]qrst",
			want: true,
		},
		{
			name: "returns true if abba outside brackets within larger string, advent of code example 4",
			s:    "ioxxoj[asdfgh]zxcvbn",
			want: true,
		},
		{
			name: "returns false if abba inside brackets, advent of code example 2",
			s:    "abcd[bddb]xyyx",
			want: false,
		},
		{
			name: "returns false if abba is actually aaaa, advent of code example 3",
			s:    "aaaa[qwer]tyui",
			want: false,
		},
	}
	re := regexp.MustCompile(`\[\w+\]`)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tlsValidation(tt.s, re)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_compileBAB(t *testing.T) {
	t.Run("returns a slice of byte slices of bab pairs to aba strings in the given slice", func(t *testing.T) {
		input := []string{
			"iygsfrrtahawqyhgvgcbnhnsss",
			"asdf",
			"asadsde",
			"tggftythsnhngggffgfssa",
		}
		got := compileBAB(input)
		assert.Equal(t, bab, got)
	})
}

func Test_hasBABMatch(t *testing.T) {
	type args struct {
		sn  []string
		bab [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns true if one of the strings provided contains one of the given bab matches",
			args: args{
				sn: []string{
					"asdrfewhc",
					"asddsyhydl",
					"irghgnfls",
					"jkhsaslfbs",
					"ssaafsjdvkpll",
				},
				bab: bab,
			},
			want: true,
		},
		{
			name: "returns false if none of the strings provided contains one of the given bab matches",
			args: args{
				sn: []string{
					"asdrfewhc",
					"asddsyhydl",
					"irghgnfls",
					"jkhtaslfbs",
					"ssaafsjdvkpll",
				},
				bab: bab,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasBABMatch(tt.args.sn, tt.args.bab)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_sslValidation(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "returns false if the string doesn't pass ssl validation",
			s:    "rfhdl[aarrewwd]skdbssu[yhythdlms]yhyjdls[aerfsqwzxds]asdvvjsolihjjkvk",
			want: false,
		},
		{
			name: "returns true if the string doesn't pass ssl validation",
			s:    "rfhdl[aarrewwd]skdbssu[yhythdlms]yhyjdls[aerfsqwzxds]asdvvjsolihyhjjkvk",
			want: true,
		},
	}
	re := regexp.MustCompile(`\[\w+\]`)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sslValidation(tt.s, re)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_countValidIPs(t *testing.T) {
	t.Run("returns the counts of ips that pass tls and ssl validations", func(t *testing.T) {
		input := []string{
			"xyx[xyx]xyx",
			"aaa[kek]eke",
			"aba[bab]xyz",
			"atgdlld[trgpo]tyhgffvcv[cvcop]",
			"ioxxoj[asdfgh]zxcvbn",
			"iuh[ldnsl]oioppos[plgghgwwgdj]ksnvsitp[ioitfsgxvcz]",
			"abba[mnop]qrst",
		}
		got, got1 := countValidIPs(input)
		assert.Equal(t, 3, got)
		assert.Equal(t, 4, got1)
	})
}
