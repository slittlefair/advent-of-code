package main

import (
	helpers "Advent-of-Code"
	"reflect"
	"testing"
)

func Test_calculateQuantumEntanglement(t *testing.T) {
	tests := []struct {
		name string
		g    []int
		want int
	}{
		{
			name: "returns 0 if there are no packages in the group",
			g:    []int{},
			want: 0,
		},
		{
			name: "returns the product of packages in the group",
			g:    []int{3, 7, 11},
			want: 231,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateQuantumEntanglement(tt.g); got != tt.want {
				t.Errorf("calculateQuantumEntanglement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupSum(t *testing.T) {
	tests := []struct {
		name     string
		packages []int
		want     int
	}{
		{
			name:     "it returns 0 if provided with an empty slice",
			packages: []int{},
			want:     0,
		},
		{
			name:     "it returns the sum of elements in the given slice",
			packages: []int{1, 2, 3, 4, 5},
			want:     15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupSum(tt.packages); got != tt.want {
				t.Errorf("groupSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLowestQuantumEntanglement(t *testing.T) {
	tests := []struct {
		name    string
		combos  [][]int
		want    int
		wantErr bool
	}{
		{
			name:    "it returns an error if no quantum entanglement can be found",
			combos:  [][]int{},
			want:    -1,
			wantErr: true,
		},
		{
			name: "it returns the lowest quantum entanglement of the given combos",
			combos: [][]int{
				{11, 9},
				{10, 8},
				{11, 7, 1},
				{9, 5, 3},
			},
			want:    77,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getLowestQuantumEntanglement(tt.combos)
			if (err != nil) != tt.wantErr {
				t.Errorf("getLowestQuantumEntanglement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getLowestQuantumEntanglement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidCombos_iterate(t *testing.T) {
	type args struct {
		remainingPackages []int
		bucket            []int
		weight            int
		maxLevel          int
	}
	tests := []struct {
		name string
		vc   *ValidCombos
		args args
		want *ValidCombos
	}{
		{
			name: "returns no combinations if no valid ones can be found",
			vc:   &ValidCombos{},
			args: args{
				remainingPackages: []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
				bucket:            []int{},
				weight:            20,
				maxLevel:          1,
			},
			want: &ValidCombos{},
		},
		{
			name: "returns combinations of the lowest length 1",
			vc:   &ValidCombos{},
			args: args{
				remainingPackages: []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
				bucket:            []int{},
				weight:            20,
				maxLevel:          2,
			},
			want: &ValidCombos{{9, 11}, {11, 9}},
		},
		{
			name: "returns combinations of the lowest length 2",
			vc:   &ValidCombos{},
			args: args{
				remainingPackages: []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
				bucket:            []int{},
				weight:            15,
				maxLevel:          2,
			},
			want: &ValidCombos{{4, 11}, {5, 10}, {7, 8}, {8, 7}, {10, 5}, {11, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vc := tt.vc
			vc.iterate(tt.args.remainingPackages, tt.args.bucket, tt.args.weight, tt.args.maxLevel)
			if !reflect.DeepEqual(vc, tt.want) {
				t.Errorf("ValidCombos.iterate() got = %v, want %v", vc, tt.want)
			}
		})
	}
}

func Test_validPermutations(t *testing.T) {
	type args struct {
		input  []int
		weight int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			name: "it returns an error if no valid perms can be found",
			args: args{
				input:  []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
				weight: 100,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "it returns correct combos 1",
			args: args{
				input:  []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
				weight: 20,
			},
			want:    ValidCombos{{9, 11}, {11, 9}},
			wantErr: false,
		},
		{
			name: "it returns correct combos 2",
			args: args{
				input:  []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
				weight: 15,
			},
			want:    ValidCombos{{4, 11}, {5, 10}, {7, 8}, {8, 7}, {10, 5}, {11, 4}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validPermutations(tt.args.input, tt.args.weight)
			if (err != nil) != tt.wantErr {
				t.Errorf("validPermutations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSolution(t *testing.T) {
	type args struct {
		input         []int
		part1Sections int
		part2Sections int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		wantErr bool
	}{
		{
			name: "returns an error if validPermutations returns an error for part1",
			args: args{
				input:         []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
				part1Sections: 100,
				part2Sections: 100,
			},
			want:    -1,
			want1:   -1,
			wantErr: true,
		},
		{
			name: "returns an error if validPermutations returns an error for part2",
			args: args{
				input:         []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
				part1Sections: 2,
				part2Sections: 100,
			},
			want:    -1,
			want1:   -1,
			wantErr: true,
		},
		{
			name: "returns an error if getLowestQuantumEntanglement returns an error for part1",
			args: args{
				input:         []int{helpers.Infinty},
				part1Sections: 1,
				part2Sections: 1,
			},
			want:    -1,
			want1:   -1,
			wantErr: true,
		},
		{
			name: "returns the lowest quantum entanglements for the given input and number of sections",
			args: args{
				input:         []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
				part1Sections: 3,
				part2Sections: 4,
			},
			want:    99,
			want1:   44,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := findSolutions(tt.args.input, tt.args.part1Sections, tt.args.part2Sections)
			if (err != nil) != tt.wantErr {
				t.Errorf("findSolution() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findSolution() = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findSolution() = %v, want %v", got1, tt.want1)
			}
		})
	}
}
