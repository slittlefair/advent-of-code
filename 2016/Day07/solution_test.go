package main

import (
	"reflect"
	"regexp"
	"testing"
)

var bab [][]byte = [][]byte{
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
			if got := containsABBA(tt.s); got != tt.want {
				t.Errorf("containsABBA() = %v, want %v", got, tt.want)
			}
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tlsValidation(tt.s, regexp.MustCompile(`\[\w+\]`)); got != tt.want {
				t.Errorf("tlsValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compileBAB(t *testing.T) {
	tests := []struct {
		name string
		ip   []string
		want [][]byte
	}{
		{
			name: "returns a slice of byte slices of bab pairs to aba strings in the given slice",
			ip: []string{
				"iygsfrrtahawqyhgvgcbnhnsss",
				"asdf",
				"asadsde",
				"tggftythsnhngggffgfssa",
			},
			want: bab,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compileBAB(tt.ip); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compileBAB() = %v, want %v", got, tt.want)
			}
		})
	}
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
			if got := hasBABMatch(tt.args.sn, tt.args.bab); got != tt.want {
				t.Errorf("hasBABMatch() = %v, want %v", got, tt.want)
			}
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sslValidation(tt.s, regexp.MustCompile(`\[\w+\]`)); got != tt.want {
				t.Errorf("sslValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countValidIPs(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
		want1 int
	}{
		{
			name: "returns the counts of ips that pass tls and ssl validations",
			input: []string{
				"xyx[xyx]xyx",
				"aaa[kek]eke",
				"aba[bab]xyz",
				"atgdlld[trgpo]tyhgffvcv[cvcop]",
				"ioxxoj[asdfgh]zxcvbn",
				"iuh[ldnsl]oioppos[plgghgwwgdj]ksnvsitp[ioitfsgxvcz]",
				"abba[mnop]qrst",
			},
			want:  3,
			want1: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := countValidIPs(tt.input)
			if got != tt.want {
				t.Errorf("countValidIPs() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("countValidIPs() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
