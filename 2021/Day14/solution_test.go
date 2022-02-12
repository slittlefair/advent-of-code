package main

import (
	"reflect"
	"testing"
)

var adventOfCodeExampleInput = []string{
	"NNCB",
	"",
	"CH -> B",
	"HH -> N",
	"CB -> H",
	"NH -> C",
	"HB -> C",
	"HC -> B",
	"HN -> C",
	"NN -> C",
	"BH -> H",
	"NC -> B",
	"NB -> B",
	"BN -> B",
	"BB -> N",
	"BC -> B",
	"CC -> N",
	"CN -> C",
}

func Test_combineLetters(t *testing.T) {
	type args struct {
		l1 string
		l2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "combines the two given strings into one",
			args: args{
				l1: "hello",
				l2: "world",
			},
			want: "helloworld",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combineLetters(tt.args.l1, tt.args.l2); got != tt.want {
				t.Errorf("combineLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    *PolymerizationEquipment
		wantErr bool
	}{
		{
			name: "returns an error if an input line has fewer than two matching strings",
			input: []string{
				"NB -> B",
				"HC -> H",
				"CH ->",
				"BB -> B",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns an error if an input line has more than two matching strings",
			input: []string{
				"NB -> B",
				"HC -> H",
				"CH -> C",
				"BB -> B -> H",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:  "returns a correctly formed PolymerizationEquipment from input",
			input: adventOfCodeExampleInput,
			want: &PolymerizationEquipment{
				pir: map[string]string{
					"CH": "B",
					"HH": "N",
					"CB": "H",
					"NH": "C",
					"HB": "C",
					"HC": "B",
					"HN": "C",
					"NN": "C",
					"BH": "H",
					"NC": "B",
					"NB": "B",
					"BN": "B",
					"BB": "N",
					"BC": "B",
					"CC": "N",
					"CN": "C",
				},
				pf: map[string]int{
					"NN": 1,
					"NC": 1,
					"CB": 1,
				},
				lf: map[string]int{
					"N": 2,
					"C": 1,
					"B": 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInput(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolymerizationEquipment_followInstructions(t *testing.T) {
	tests := []struct {
		name string
		pe   *PolymerizationEquipment
		want *PolymerizationEquipment
	}{
		{
			name: "follows instruction, advent of code example 1",
			pe: &PolymerizationEquipment{
				pir: map[string]string{
					"CH": "B",
					"HH": "N",
					"CB": "H",
					"NH": "C",
					"HB": "C",
					"HC": "B",
					"HN": "C",
					"NN": "C",
					"BH": "H",
					"NC": "B",
					"NB": "B",
					"BN": "B",
					"BB": "N",
					"BC": "B",
					"CC": "N",
					"CN": "C",
				},
				pf: map[string]int{
					"NN": 1,
					"NC": 1,
					"CB": 1,
				},
				lf: map[string]int{
					"N": 2,
					"C": 1,
					"B": 1,
				},
			},
			want: &PolymerizationEquipment{
				pir: map[string]string{
					"CH": "B",
					"HH": "N",
					"CB": "H",
					"NH": "C",
					"HB": "C",
					"HC": "B",
					"HN": "C",
					"NN": "C",
					"BH": "H",
					"NC": "B",
					"NB": "B",
					"BN": "B",
					"BB": "N",
					"BC": "B",
					"CC": "N",
					"CN": "C",
				},
				pf: map[string]int{
					"NC": 1,
					"CN": 1,
					"NB": 1,
					"BC": 1,
					"CH": 1,
					"HB": 1,
				},
				lf: map[string]int{
					"N": 2,
					"C": 2,
					"B": 2,
					"H": 1,
				},
			},
		},
		{
			name: "follows instruction, advent of code example 2",
			pe: &PolymerizationEquipment{
				pir: map[string]string{
					"CH": "B",
					"HH": "N",
					"CB": "H",
					"NH": "C",
					"HB": "C",
					"HC": "B",
					"HN": "C",
					"NN": "C",
					"BH": "H",
					"NC": "B",
					"NB": "B",
					"BN": "B",
					"BB": "N",
					"BC": "B",
					"CC": "N",
					"CN": "C",
				},
				pf: map[string]int{
					"NC": 1,
					"CN": 1,
					"NB": 1,
					"BC": 1,
					"CH": 1,
					"HB": 1,
				},
				lf: map[string]int{
					"N": 2,
					"C": 2,
					"B": 2,
					"H": 1,
				},
			},
			want: &PolymerizationEquipment{
				pir: map[string]string{
					"CH": "B",
					"HH": "N",
					"CB": "H",
					"NH": "C",
					"HB": "C",
					"HC": "B",
					"HN": "C",
					"NN": "C",
					"BH": "H",
					"NC": "B",
					"NB": "B",
					"BN": "B",
					"BB": "N",
					"BC": "B",
					"CC": "N",
					"CN": "C",
				},
				pf: map[string]int{
					"NB": 2,
					"BC": 2,
					"CC": 1,
					"CN": 1,
					"BB": 2,
					"CB": 2,
					"BH": 1,
					"HC": 1,
				},
				lf: map[string]int{
					"N": 2,
					"C": 4,
					"B": 6,
					"H": 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pe := tt.pe
			pe.followInstructions()
			if !reflect.DeepEqual(pe, tt.want) {
				t.Errorf("pe.followInstructions() = %v, want %v", pe, tt.want)
			}
		})
	}
}

func TestPolymerizationEquipment_getVal(t *testing.T) {
	tests := []struct {
		name string
		lf   map[string]int
		want int
	}{
		{
			name: "returns max minus min value, advent of code example 1",
			lf: map[string]int{
				"B": 1749,
				"C": 298,
				"H": 161,
				"N": 865,
			},
			want: 1588,
		},
		{
			name: "returns max minus min value, advent of code example 2",
			lf: map[string]int{
				"B": 2192039569602,
				"C": 29862009173,
				"H": 3849876073,
				"N": 2192039569601,
			},
			want: 2188189693529,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pe := PolymerizationEquipment{
				lf: tt.lf,
			}
			if got := pe.getVal(); got != tt.want {
				t.Errorf("PolymerizationEquipment.getVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSolutions(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    int
		want1   int
		wantErr bool
	}{
		{
			name: "returns an error if input cannot be parsed",
			input: []string{
				"ABC -> J",
				"TGV -> IKH",
				"GF -> X",
				"CC -> C",
				"AG -> YU -> B",
				"HH -> H",
			},
			want:    -1,
			want1:   -1,
			wantErr: true,
		},
		{
			name:    "returns correct solutions for parts 1 and 2 from input, advent of code example",
			input:   adventOfCodeExampleInput,
			want:    1588,
			want1:   2188189693529,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := findSolutions(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("findSolutions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findSolutions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findSolutions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
