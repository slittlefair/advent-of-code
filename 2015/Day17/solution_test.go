package main

import (
	"reflect"
	"testing"
)

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name    string
		arg     []string
		want    []int
		wantErr bool
	}{
		{
			name: "returns an error if an input line can't be converted to an int",
			arg: []string{
				"1",
				"5",
				"2",
				"a",
				"22",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns a sorted list of containers, highest to lowest",
			arg: []string{
				"1",
				"5",
				"2",
				"10",
				"22",
			},
			want:    []int{22, 10, 5, 2, 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInput(tt.arg)
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

func TestEggnogContainers_FindContainers(t *testing.T) {
	type fields struct {
		WantedTotal int
		Ways        map[int]int
	}
	type args struct {
		remainingContainers []int
		totalCapacity       int
		levels              int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *EggnogContainers
	}{
		{
			name: "updates EggnogContainer with permutations that sum to WantedTotal, advent of code example",
			fields: fields{
				WantedTotal: 25,
				Ways:        make(map[int]int),
			},
			args: args{
				remainingContainers: []int{20, 15, 10, 5, 5},
				totalCapacity:       0,
				levels:              0,
			},
			want: &EggnogContainers{
				WantedTotal: 25,
				Ways: map[int]int{
					3: 1,
					2: 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := &EggnogContainers{
				WantedTotal: tt.fields.WantedTotal,
				Ways:        tt.fields.Ways,
			}
			ec.FindContainers(tt.args.remainingContainers, tt.args.totalCapacity, tt.args.levels)
			if !reflect.DeepEqual(ec, tt.want) {
				t.Errorf("EggnogContainers.FindContainers() = %v, want %v", ec, tt.want)
			}
		})
	}
}

func TestEggnogContainers_CountPermutations(t *testing.T) {
	type fields struct {
		WantedTotal int
		Ways        map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "returns 0 for an empty EggnogComntainers",
			want: 0,
		},
		{
			name: "returns sum of all values in Ways, advent of code example",
			fields: fields{
				Ways: map[int]int{
					3: 1,
					2: 3,
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := EggnogContainers{
				WantedTotal: tt.fields.WantedTotal,
				Ways:        tt.fields.Ways,
			}
			if got := ec.CountPermutations(); got != tt.want {
				t.Errorf("EggnogContainers.CountPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEggnogContainers_CountSmallestContainersPermutations(t *testing.T) {
	type fields struct {
		WantedTotal int
		Ways        map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "returns 0 for an empty EggnogComntainers",
			want: 0,
		},
		{
			name: "returns number of values for the smallest key in Ways, advent of code example",
			fields: fields{
				Ways: map[int]int{
					3: 1,
					2: 3,
				},
			},
			want: 3,
		},
		{
			name: "returns number of values for the smallest key in Ways, more complex example",
			fields: fields{
				Ways: map[int]int{
					6:     1,
					8:     3,
					12:    628,
					9:     12,
					3:     32,
					20:    7,
					88:    17,
					88888: 1,
					42527: 3673,
					7:     9182,
				},
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := EggnogContainers{
				WantedTotal: tt.fields.WantedTotal,
				Ways:        tt.fields.Ways,
			}
			if got := ec.CountSmallestContainersPermutations(); got != tt.want {
				t.Errorf("EggnogContainers.CountSmallestContainersPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
