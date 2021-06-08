package main

import (
	"testing"
)

func Test_findExactMatch(t *testing.T) {
	type args struct {
		input  []string
		ticker map[string]int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "returns an error if a line has mismatched number of words and numbers",
			args: args{
				input: []string{
					"Sue 1: children: 1, cars: 8, vizslas: 7",
					"Sue 2: akitas: 10, perfumes: 10, children: about 5",
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "returns an error if no matches are found",
			args: args{
				input: []string{
					"Sue 1: children: 1, cars: 8, vizslas: 7",
					"Sue 2: akitas: 10, perfumes: 10, children: 5",
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "returns a matching sue",
			args: args{
				input: []string{
					"Sue 1: children: 1, cars: 8, vizslas: 7",
					"Sue 2: akitas: 10, perfumes: 10, children: 5",
					"Sue 3: cars: 5, pomeranians: 4, vizslas: 1",
				},
				ticker: map[string]int{
					"children": 5,
					"akitas":   10,
					"perfumes": 10,
					"cars":     9,
				},
			},
			want:    "2",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findExactMatch(tt.args.input, tt.args.ticker)
			if (err != nil) != tt.wantErr {
				t.Errorf("findExactMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findExactMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRangedMatch(t *testing.T) {
	type args struct {
		input  []string
		ticker map[string]int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "returns an error if a line has mismatched number of words and numbers",
			args: args{
				input: []string{
					"Sue 1: children: 1, cars: 8, vizslas: 7",
					"Sue 2: akitas: 10, perfumes: 10, children: about 5",
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "returns an error if no matches are found",
			args: args{
				input: []string{
					"Sue 1: children: 1, cars: 8, vizslas: 7",
					"Sue 2: akitas: 10, perfumes: 10, children: 5",
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "doesn't match on cats if sue value is equal to ticker value",
			args: args{
				input: []string{
					"Sue 1: children: 1, cats: 8, vizslas: 7",
					"Sue 2: akitas: 10, perfumes: 10, children: 5",
				},
				ticker: map[string]int{
					"children": 1,
					"akitas":   10,
					"perfumes": 10,
					"cats":     8,
					"vizslas":  7,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "doesn't match on cats if sue value is less than ticker value",
			args: args{
				input: []string{
					"Sue 1: children: 1, cats: 5, vizslas: 7",
					"Sue 2: akitas: 10, perfumes: 10, children: 5",
				},
				ticker: map[string]int{
					"children": 1,
					"akitas":   10,
					"perfumes": 10,
					"cats":     8,
					"vizslas":  7,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "does match on cats if sue value is greater than ticker value",
			args: args{
				input: []string{
					"Sue 1: children: 1, cats: 9, vizslas: 7",
					"Sue 2: akitas: 10, perfumes: 10, children: 5",
				},
				ticker: map[string]int{
					"children": 1,
					"akitas":   10,
					"perfumes": 10,
					"cats":     8,
					"vizslas":  7,
				},
			},
			want:    "1",
			wantErr: false,
		},
		{
			name: "doesn't match on trees if sue value is equal to ticker value",
			args: args{
				input: []string{
					"Sue 1: children: 1, akitas: 10, vizslas: 7",
					"Sue 2: akitas: 10, trees: 10, children: 5",
				},
				ticker: map[string]int{
					"children": 1,
					"akitas":   10,
					"trees":    10,
					"cats":     5,
					"vizslas":  6,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "doesn't match on trees if sue value is less than ticker value",
			args: args{
				input: []string{
					"Sue 1: children: 1, akitas: 10, vizslas: 7",
					"Sue 2: akitas: 10, trees: 9, children: 5",
				},
				ticker: map[string]int{
					"children": 1,
					"akitas":   10,
					"trees":    10,
					"cats":     8,
					"vizslas":  6,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "does match on trees if sue value is greater than ticker value",
			args: args{
				input: []string{
					"Sue 1: children: 1, akitas: 10, vizslas: 7",
					"Sue 2: akitas: 10, trees: 99, children: 5",
				},
				ticker: map[string]int{
					"children": 5,
					"akitas":   10,
					"trees":    10,
					"cats":     8,
					"vizslas":  6,
				},
			},
			want:    "2",
			wantErr: false,
		},
		{
			name: "doesn't match on pomeranians if sue value is equal to ticker value",
			args: args{
				input: []string{
					"Sue 1: children: 1, akitas: 10, vizslas: 7",
					"Sue 2: akitas: 10, pomeranians: 10, children: 5",
				},
				ticker: map[string]int{
					"children":    1,
					"akitas":      10,
					"pomeranians": 10,
					"cats":        5,
					"vizslas":     6,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "doesn't match on pomeranians if sue value is greater than ticker value",
			args: args{
				input: []string{
					"Sue 1: children: 1, akitas: 10, vizslas: 7",
					"Sue 2: akitas: 10, pomeranians: 11, children: 5",
				},
				ticker: map[string]int{
					"children":    1,
					"akitas":      10,
					"pomeranians": 10,
					"cats":        8,
					"vizslas":     6,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "does match on pomeranians if sue value is less than ticker value",
			args: args{
				input: []string{
					"Sue 1: children: 1, akitas: 10, vizslas: 7",
					"Sue 2: akitas: 10, pomeranians: 0, children: 5",
				},
				ticker: map[string]int{
					"children":    5,
					"akitas":      10,
					"pomeranians": 10,
					"cats":        8,
					"vizslas":     6,
				},
			},
			want:    "2",
			wantErr: false,
		},
		{
			name: "doesn't match on goldfish if sue value is equal to ticker value",
			args: args{
				input: []string{
					"Sue 1: children: 1, akitas: 10, vizslas: 7",
					"Sue 2: akitas: 10, goldfish: 10, children: 5",
				},
				ticker: map[string]int{
					"children": 1,
					"akitas":   10,
					"goldfish": 10,
					"cats":     5,
					"vizslas":  6,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "doesn't match on goldfish if sue value is greater than ticker value",
			args: args{
				input: []string{
					"Sue 1: children: 1, akitas: 10, vizslas: 7",
					"Sue 2: akitas: 10, goldfish: 11, children: 5",
				},
				ticker: map[string]int{
					"children": 1,
					"akitas":   10,
					"goldfish": 10,
					"cats":     8,
					"vizslas":  6,
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "does match on goldfish if sue value is less than ticker value",
			args: args{
				input: []string{
					"Sue 1: children: 1, akitas: 10, vizslas: 7",
					"Sue 2: akitas: 10, goldfish: 0, children: 5",
				},
				ticker: map[string]int{
					"children": 5,
					"akitas":   10,
					"goldfish": 10,
					"cats":     8,
					"vizslas":  6,
				},
			},
			want:    "2",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findRangedMatch(tt.args.input, tt.args.ticker)
			if (err != nil) != tt.wantErr {
				t.Errorf("findRangedMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findRangedMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
