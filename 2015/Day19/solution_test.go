package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInput(t *testing.T) {
	t.Run("correctly parses input, advent of code example", func(t *testing.T) {
		arg := []string{
			"H => HO",
			"H => OH",
			"O => HH",
			"",
			"HOH",
		}
		want := &Medicine{
			Replacements: Replacements{
				"H": {"HO", "OH"},
				"O": {"HH"},
			},
			Molecule:             "HOH",
			DistinctNewMolecules: make(map[string]bool),
		}
		got := parseInput(arg)
		assert.Equal(t, want, got)
	})
}

func TestMedicine_FindIndicesOfStringInMolecule(t *testing.T) {
	type fields struct {
		Replacements         Replacements
		Molecule             string
		NewMolecules         []string
		DistinctNewMolecules map[string]bool
	}
	tests := []struct {
		name   string
		fields fields
		arg    string
		want   []int
	}{
		{
			name: "returns an empty slice if the substring does not appear",
			fields: fields{
				Molecule: "HOH",
			},
			arg:  "S",
			want: []int{},
		},
		{
			name: "returns an slice of substring indices for a single character substring",
			fields: fields{
				Molecule: "HOH",
			},
			arg:  "H",
			want: []int{0, 2},
		},
		{
			name: "returns an slice of substring indices for a multiple character substring",
			fields: fields{
				Molecule: "AbXDAgarAbabbAAAVAAbadRBa",
			},
			arg:  "Ab",
			want: []int{0, 8, 18},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Medicine{
				Replacements:         tt.fields.Replacements,
				Molecule:             tt.fields.Molecule,
				DistinctNewMolecules: tt.fields.DistinctNewMolecules,
			}
			got := m.FindIndicesOfStringInMolecule(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMedicine_ReplaceAndFindNewMolecules(t *testing.T) {
	type fields struct {
		Replacements         Replacements
		Molecule             string
		NewMolecules         []string
		DistinctNewMolecules map[string]bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Medicine
	}{
		{
			name: "finds new molecules for single character replacements, advent of code example 1",
			fields: fields{
				Replacements: Replacements{
					"H": {"HO", "OH"},
					"O": {"HH"},
				},
				Molecule:             "HOH",
				DistinctNewMolecules: make(map[string]bool),
			},
			want: &Medicine{
				Replacements: Replacements{
					"H": {"HO", "OH"},
					"O": {"HH"},
				},
				Molecule: "HOH",
				DistinctNewMolecules: map[string]bool{
					"HOOH": true,
					"HOHO": true,
					"OHOH": true,
					"HHHH": true,
				},
			},
		},
		{
			name: "finds new molecules for single character replacements, advent of code example 2",
			fields: fields{
				Replacements: Replacements{
					"H": {"HO", "OH"},
					"O": {"HH"},
				},
				Molecule:             "HOHOHO",
				DistinctNewMolecules: make(map[string]bool),
			},
			want: &Medicine{
				Replacements: Replacements{
					"H": {"HO", "OH"},
					"O": {"HH"},
				},
				Molecule: "HOHOHO",
				DistinctNewMolecules: map[string]bool{
					"HOOHOHO": true,
					"HOHOOHO": true,
					"HOHOHOO": true,
					"OHOHOHO": true,
					"HHHHOHO": true,
					"HOHHHHO": true,
					"HOHOHHH": true,
				},
			},
		},
		{
			name: "finds new molecules for multiple character replacements",
			fields: fields{
				Replacements: Replacements{
					"H":  {"HO", "OH"},
					"O":  {"HH"},
					"Ab": {"H", "AA"},
				},
				Molecule:             "HOHAAAbO",
				DistinctNewMolecules: make(map[string]bool),
			},
			want: &Medicine{
				Replacements: Replacements{
					"H":  {"HO", "OH"},
					"O":  {"HH"},
					"Ab": {"H", "AA"},
				},
				Molecule: "HOHAAAbO",
				DistinctNewMolecules: map[string]bool{
					"HOOHAAAbO": true,
					"HOHOAAAbO": true,
					"OHOHAAAbO": true,
					"HHHHAAAbO": true,
					"HOHAAAbHH": true,
					"HOHAAHO":   true,
					"HOHAAAAO":  true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Medicine{
				Replacements:         tt.fields.Replacements,
				Molecule:             tt.fields.Molecule,
				DistinctNewMolecules: tt.fields.DistinctNewMolecules,
			}
			m.ReplaceAndFindNewMolecules()
			assert.Equal(t, tt.want, m)
		})
	}
}

func Test_countUpper(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{
			name: "returns 0 if given string is empty",
			arg:  "",
			want: 0,
		},
		{
			name: "returns 0 if given string has no capitals",
			arg:  "kfuhgjkhekgjer;k:ahfgkjfhkerfl",
			want: 0,
		},
		{
			name: "counts number of upper case letters in a given string",
			arg:  "AbaJHAbahJ;k:LKKJjjkjKJjKj",
			want: 12,
		},
		{
			name: "returns length of given string if all letters are capitals",
			arg:  "GHJHGKKLKPKHL",
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countUpper(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMedicine_GetNumberOfSubs(t *testing.T) {
	type fields struct {
		Replacements         Replacements
		Molecule             string
		DistinctNewMolecules map[string]bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "returns number of steps required to make molecule, all upper characters",
			fields: fields{
				Molecule: "HOHOHO",
			},
			want: 5,
		},
		{
			name: "returns number of steps required to make molecule, complex",
			fields: fields{
				Molecule: "RnAArRNHGuKLReCaRnYDSxArYOC",
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Medicine{
				Replacements:         tt.fields.Replacements,
				Molecule:             tt.fields.Molecule,
				DistinctNewMolecules: tt.fields.DistinctNewMolecules,
			}
			got := m.GetNumberOfSubs()
			assert.Equal(t, tt.want, got)
		})
	}
}
