package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var classParams = FieldParams{
	vn: ValidNumbers{
		1: true,
		2: true,
		3: true,
		5: true,
		6: true,
		7: true,
	},
	pvi: PossibleValueIndices{
		0: true,
		1: true,
		2: true,
	},
}
var rowParams = FieldParams{
	vn: ValidNumbers{
		6:  true,
		7:  true,
		8:  true,
		9:  true,
		10: true,
		11: true,
		33: true,
		34: true,
		35: true,
		36: true,
		37: true,
		38: true,
		39: true,
		40: true,
		41: true,
		42: true,
		43: true,
		44: true,
	},
	pvi: PossibleValueIndices{
		0: true,
		1: true,
		2: true,
	},
}
var seatParams = FieldParams{
	vn: ValidNumbers{
		13: true,
		14: true,
		15: true,
		16: true,
		17: true,
		18: true,
		19: true,
		20: true,
		21: true,
		22: true,
		23: true,
		24: true,
		25: true,
		26: true,
		27: true,
		28: true,
		29: true,
		30: true,
		31: true,
		32: true,
		33: true,
		34: true,
		35: true,
		36: true,
		37: true,
		38: true,
		39: true,
		40: true,
		45: true,
		46: true,
		47: true,
		48: true,
		49: true,
		50: true,
	},
	pvi: PossibleValueIndices{
		0: true,
		1: true,
		2: true,
	},
}

func TestTicketFields_populateField(t *testing.T) {
	tests := []struct {
		name               string
		tf                 TicketFields
		field              []string
		want               TicketFields
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name:               "returns an error if the line is not parsed correctly",
			tf:                 TicketFields{},
			field:              []string{"class", "1-2 or a-b"},
			want:               TicketFields{},
			errorAssertionFunc: assert.Error,
		},
		{
			name:  "advent of code example 1",
			tf:    TicketFields{},
			field: []string{"class", "1-3 or 5-7"},
			want: TicketFields{
				"class": {vn: classParams.vn},
			},
			errorAssertionFunc: assert.NoError,
		},
		{
			name:  "advent of code example 2",
			tf:    TicketFields{},
			field: []string{"row", "6-11 or 33-44"},
			want: TicketFields{
				"row": {vn: rowParams.vn},
			},
			errorAssertionFunc: assert.NoError,
		},
		{
			name:  "advent of code example 3",
			tf:    TicketFields{},
			field: []string{"seat", "13-40 or 45-50"},
			want: TicketFields{
				"seat": {vn: seatParams.vn},
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.tf.populateField(tt.field)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, tt.tf)
		})
	}
}

func TestTicketFields_numIsValid(t *testing.T) {
	tests := []struct {
		name string
		tf   TicketFields
		num  int
		want bool
	}{
		{
			name: "returns true if the number is valid for multiple fields",
			tf: TicketFields{
				"row":   rowParams,
				"class": classParams,
				"seat":  seatParams,
			},
			num:  6,
			want: true,
		},
		{
			name: "returns true if the number is valid for only one field",
			tf: TicketFields{
				"row":   {vn: rowParams.vn},
				"class": {vn: classParams.vn},
				"seat":  {vn: seatParams.vn},
			},
			num:  50,
			want: true,
		},
		{
			name: "returns false if the number is valid for no fields",
			tf: TicketFields{
				"row":   rowParams,
				"class": classParams,
				"seat":  seatParams,
			},
			num:  500,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tf.numIsValid(tt.num)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTicketFields_populatePossibleValueIndices(t *testing.T) {
	t.Run("correctly populates PossibleValueIndices", func(t *testing.T) {
		tf := TicketFields{
			"class": FieldParams{
				vn: ValidNumbers{
					1: true,
					2: true,
					7: true,
				},
			},
			"seat": FieldParams{
				vn: ValidNumbers{
					8: true,
					9: true,
				},
			},
		}
		want := TicketFields{
			"class": FieldParams{
				vn: ValidNumbers{
					1: true,
					2: true,
					7: true,
				},
				pvi: PossibleValueIndices{
					0: true,
					1: true,
					2: true,
				},
			},
			"seat": FieldParams{
				vn: ValidNumbers{
					8: true,
					9: true,
				},
				pvi: PossibleValueIndices{
					0: true,
					1: true,
					2: true,
				},
			},
		}
		tf.populatePossibleValueIndices(3)
		assert.Equal(t, want, tf)
	})
}

func TestTicketFields_part1(t *testing.T) {
	type args struct {
		nums            []int
		allValidTickets TicketCollection
	}
	tests := []struct {
		name  string
		tf    TicketFields
		args  args
		want  int
		want1 TicketCollection
	}{
		{
			name: "advent of code example 1",
			tf: TicketFields{
				"row":   rowParams,
				"class": classParams,
				"seat":  seatParams,
			},
			args: args{
				nums:            []int{7, 3, 47},
				allValidTickets: TicketCollection{},
			},
			want:  0,
			want1: TicketCollection{{7, 3, 47}},
		},
		{
			name: "advent of code example 2",
			tf: TicketFields{
				"row":   rowParams,
				"class": classParams,
				"seat":  seatParams,
			},
			args: args{
				nums:            []int{40, 4, 50},
				allValidTickets: TicketCollection{{7, 3, 47}},
			},
			want:  4,
			want1: TicketCollection{{7, 3, 47}},
		},
		{
			name: "advent of code example 3",
			tf: TicketFields{
				"row":   rowParams,
				"class": classParams,
				"seat":  seatParams,
			},
			args: args{
				nums:            []int{55, 2, 20},
				allValidTickets: TicketCollection{{7, 3, 47}},
			},
			want:  55,
			want1: TicketCollection{{7, 3, 47}},
		},
		{
			name: "advent of code example 4",
			tf: TicketFields{
				"row":   rowParams,
				"class": classParams,
				"seat":  seatParams,
			},
			args: args{
				nums:            []int{38, 6, 12},
				allValidTickets: TicketCollection{{7, 3, 47}},
			},
			want:  12,
			want1: TicketCollection{{7, 3, 47}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.tf.validateTicket(tt.args.nums, tt.args.allValidTickets)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestTicketFields_part2(t *testing.T) {
	t.Run("advent of code example", func(t *testing.T) {
		tf := TicketFields{
			"departure class": FieldParams{
				vn: ValidNumbers{
					0:  true,
					1:  true,
					4:  true,
					5:  true,
					6:  true,
					7:  true,
					8:  true,
					9:  true,
					10: true,
					11: true,
					12: true,
					13: true,
					14: true,
					15: true,
					16: true,
					17: true,
					18: true,
					19: true,
				},
				pvi: PossibleValueIndices{
					0: true,
					1: true,
					2: true,
				},
			},
			"departure row": FieldParams{
				vn: ValidNumbers{
					0:  true,
					1:  true,
					2:  true,
					3:  true,
					4:  true,
					5:  true,
					8:  true,
					9:  true,
					10: true,
					11: true,
					12: true,
					13: true,
					14: true,
					15: true,
					16: true,
					17: true,
					18: true,
					19: true,
				},
				pvi: PossibleValueIndices{
					0: true,
					1: true,
					2: true,
				},
			},
			"seat": FieldParams{
				vn: ValidNumbers{
					0:  true,
					1:  true,
					2:  true,
					3:  true,
					4:  true,
					5:  true,
					6:  true,
					7:  true,
					8:  true,
					9:  true,
					10: true,
					11: true,
					12: true,
					13: true,
					16: true,
					17: true,
					18: true,
					19: true,
				},
				pvi: PossibleValueIndices{
					0: true,
					1: true,
					2: true,
				},
			},
		}
		myTicket := []int{11, 12, 13}
		allValidTickets := TicketCollection{{3, 9, 18}, {15, 1, 5}, {5, 14, 9}}
		got := tf.part2(myTicket, allValidTickets)
		assert.Equal(t, 132, got)
	})
}

func TestTicketFields_runSolution(t *testing.T) {
	tests := []struct {
		name               string
		tf                 TicketFields
		entries            []string
		want               []int
		want1              int
		want2              TicketCollection
		want3              TicketFields
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if the input cannot be parsed correctly",
			tf:   TicketFields{},
			entries: []string{
				"class: 1-3 or 5-7",
				"row: 6-11 or 33-44",
				"seat: 13-40 or 45-5bbb0",
				"",
				"your ticket:",
				"7,1,14",
				"",
				"nearby tickets:",
				"7,3,47",
				"40,4,50",
				"55,2,20",
				"38,6,12",
			},
			want:  nil,
			want1: 0,
			want2: nil,
			want3: TicketFields{
				"class": {vn: classParams.vn},
				"row":   {vn: rowParams.vn},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "advent of code example",
			tf:   TicketFields{},
			entries: []string{
				"class: 1-3 or 5-7",
				"row: 6-11 or 33-44",
				"seat: 13-40 or 45-50",
				"",
				"your ticket:",
				"7,1,14",
				"",
				"nearby tickets:",
				"7,3,47",
				"40,4,50",
				"55,2,20",
				"38,6,12",
			},
			want:  []int{7, 1, 14},
			want1: 71,
			want2: TicketCollection{{7, 3, 47}},
			want3: TicketFields{
				"row":   rowParams,
				"class": classParams,
				"seat":  seatParams,
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := tt.tf.runSolution(tt.entries)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
			assert.Equal(t, tt.want2, got2)
		})
	}
}
