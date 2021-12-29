package main

import (
	helpers "Advent-of-Code"
	"reflect"
	"testing"
)

var adventOfCodeExampleHeightMap = HeightMap{
	{X: 0, Y: 0}: 2,
	{X: 1, Y: 0}: 1,
	{X: 2, Y: 0}: 9,
	{X: 3, Y: 0}: 9,
	{X: 4, Y: 0}: 9,
	{X: 5, Y: 0}: 4,
	{X: 6, Y: 0}: 3,
	{X: 7, Y: 0}: 2,
	{X: 8, Y: 0}: 1,
	{X: 9, Y: 0}: 0,
	{X: 0, Y: 1}: 3,
	{X: 1, Y: 1}: 9,
	{X: 2, Y: 1}: 8,
	{X: 3, Y: 1}: 7,
	{X: 4, Y: 1}: 8,
	{X: 5, Y: 1}: 9,
	{X: 6, Y: 1}: 4,
	{X: 7, Y: 1}: 9,
	{X: 8, Y: 1}: 2,
	{X: 9, Y: 1}: 1,
	{X: 0, Y: 2}: 9,
	{X: 1, Y: 2}: 8,
	{X: 2, Y: 2}: 5,
	{X: 3, Y: 2}: 6,
	{X: 4, Y: 2}: 7,
	{X: 5, Y: 2}: 8,
	{X: 6, Y: 2}: 9,
	{X: 7, Y: 2}: 8,
	{X: 8, Y: 2}: 9,
	{X: 9, Y: 2}: 2,
	{X: 0, Y: 3}: 8,
	{X: 1, Y: 3}: 7,
	{X: 2, Y: 3}: 6,
	{X: 3, Y: 3}: 7,
	{X: 4, Y: 3}: 8,
	{X: 5, Y: 3}: 9,
	{X: 6, Y: 3}: 6,
	{X: 7, Y: 3}: 7,
	{X: 8, Y: 3}: 8,
	{X: 9, Y: 3}: 9,
	{X: 0, Y: 4}: 9,
	{X: 1, Y: 4}: 8,
	{X: 2, Y: 4}: 9,
	{X: 3, Y: 4}: 9,
	{X: 4, Y: 4}: 9,
	{X: 5, Y: 4}: 6,
	{X: 6, Y: 4}: 5,
	{X: 7, Y: 4}: 6,
	{X: 8, Y: 4}: 7,
	{X: 9, Y: 4}: 8,
}

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  HeightMap
	}{
		{
			name: "correctly parses input into a HeightMap",
			input: []string{
				"2199943210",
				"3987894921",
				"9856789892",
				"8767896789",
				"9899965678",
			},
			want: adventOfCodeExampleHeightMap,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeightMap_findLowPoints(t *testing.T) {
	tests := []struct {
		name string
		hm   HeightMap
		want LowPoints
	}{
		{
			name: "it finds the correct low points from the given height map, advent of code example",
			hm:   adventOfCodeExampleHeightMap,
			want: LowPoints{
				{X: 1, Y: 0}: 1,
				{X: 9, Y: 0}: 0,
				{X: 2, Y: 2}: 5,
				{X: 6, Y: 4}: 5,
			},
		},
		{
			name: "it finds the correct low points from the given height map with extra adjacent points",
			hm: HeightMap{
				{X: 0, Y: 0}: 2,
				{X: 1, Y: 0}: 1,
				{X: 2, Y: 0}: 9,
				{X: 3, Y: 0}: 9,
				{X: 4, Y: 0}: 9,
				{X: 5, Y: 0}: 4,
				{X: 6, Y: 0}: 3,
				{X: 7, Y: 0}: 2,
				{X: 8, Y: 0}: 1,
				{X: 9, Y: 0}: 0,
				{X: 0, Y: 1}: 3,
				{X: 1, Y: 1}: 9,
				{X: 2, Y: 1}: 8,
				{X: 3, Y: 1}: 7,
				{X: 4, Y: 1}: 8,
				{X: 5, Y: 1}: 9,
				{X: 6, Y: 1}: 4,
				{X: 7, Y: 1}: 9,
				{X: 8, Y: 1}: 2,
				{X: 9, Y: 1}: 1,
				{X: 0, Y: 2}: 9,
				{X: 1, Y: 2}: 8,
				{X: 2, Y: 2}: 5,
				{X: 3, Y: 2}: 6,
				{X: 4, Y: 2}: 7,
				{X: 5, Y: 2}: 8,
				{X: 6, Y: 2}: 9,
				{X: 7, Y: 2}: 8,
				{X: 8, Y: 2}: 9,
				{X: 9, Y: 2}: 2,
				{X: 0, Y: 3}: 9,
				{X: 1, Y: 3}: 7,
				{X: 2, Y: 3}: 6,
				{X: 3, Y: 3}: 7,
				{X: 4, Y: 3}: 8,
				{X: 5, Y: 3}: 9,
				{X: 6, Y: 3}: 6,
				{X: 7, Y: 3}: 7,
				{X: 8, Y: 3}: 8,
				{X: 9, Y: 3}: 9,
				{X: 0, Y: 4}: 9,
				{X: 1, Y: 4}: 9,
				{X: 2, Y: 4}: 9,
				{X: 3, Y: 4}: 9,
				{X: 4, Y: 4}: 9,
				{X: 5, Y: 4}: 6,
				{X: 6, Y: 4}: 5,
				{X: 7, Y: 4}: 6,
				{X: 8, Y: 4}: 7,
				{X: 9, Y: 4}: 8,
			},
			want: LowPoints{
				{X: 1, Y: 0}: 1,
				{X: 9, Y: 0}: 0,
				{X: 2, Y: 2}: 5,
				{X: 6, Y: 4}: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hm.findLowPoints(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeightMap.findLowPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateRiskLevels(t *testing.T) {
	tests := []struct {
		name      string
		lowPoints LowPoints
		want      int
	}{
		{
			name: "returns sum of a low points + 1",
			lowPoints: LowPoints{
				{X: 1, Y: 0}: 1,
				{X: 9, Y: 0}: 0,
				{X: 2, Y: 2}: 5,
				{X: 6, Y: 4}: 5,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateRiskLevels(tt.lowPoints); got != tt.want {
				t.Errorf("calculateRiskLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeightMap_coIsPartOfBasin(t *testing.T) {
	type args struct {
		b  Basin
		co helpers.Co
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns false if given co is already part of the basin",
			args: args{
				b: Basin{
					{X: 0, Y: 0}: {},
					{X: 4, Y: 2}: {},
					{X: 1, Y: 2}: {},
					{X: 3, Y: 3}: {},
					{X: 0, Y: 9}: {},
					{X: 4, Y: 5}: {},
				},
				co: helpers.Co{X: 0, Y: 9},
			},
			want: false,
		},
		{
			name: "returns false if given co is not part of the height map",
			args: args{
				b: Basin{
					{X: 0, Y: 0}: {},
					{X: 4, Y: 2}: {},
					{X: 1, Y: 2}: {},
					{X: 3, Y: 3}: {},
					{X: 0, Y: 9}: {},
					{X: 4, Y: 5}: {},
				},
				co: helpers.Co{X: 0, Y: 10},
			},
			want: false,
		},
		{
			name: "returns false if given co is not part of the height map",
			args: args{
				b: Basin{
					{X: 0, Y: 0}: {},
					{X: 4, Y: 2}: {},
					{X: 1, Y: 2}: {},
					{X: 3, Y: 3}: {},
					{X: 0, Y: 9}: {},
					{X: 4, Y: 5}: {},
				},
				co: helpers.Co{X: 0, Y: 10},
			},
			want: false,
		},
		{
			name: "returns false if given co has a value of 9 in the height map",
			args: args{
				b: Basin{
					{X: 0, Y: 0}: {},
					{X: 4, Y: 2}: {},
					{X: 1, Y: 2}: {},
					{X: 3, Y: 3}: {},
					{X: 0, Y: 9}: {},
					{X: 4, Y: 5}: {},
				},
				co: helpers.Co{X: 3, Y: 4},
			},
			want: false,
		},
		{
			name: "returns true if given co does not fulfill the above criteria",
			args: args{
				b: Basin{
					{X: 0, Y: 0}: {},
					{X: 5, Y: 2}: {},
					{X: 1, Y: 2}: {},
					{X: 3, Y: 3}: {},
					{X: 0, Y: 9}: {},
					{X: 4, Y: 5}: {},
				},
				co: helpers.Co{X: 8, Y: 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adventOfCodeExampleHeightMap.coIsPartOfBasin(tt.args.b, tt.args.co); got != tt.want {
				t.Errorf("HeightMap.coIsPartOfBasin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeightMap_calculateBasin(t *testing.T) {
	tests := []struct {
		name string
		co   helpers.Co
		want int
	}{
		{
			name: "calculates correct basin size, advent of code example 1",
			co:   helpers.Co{X: 0, Y: 0},
			want: 3,
		},
		{
			name: "calculates correct basin size, advent of code example 2",
			co:   helpers.Co{X: 9, Y: 0},
			want: 9,
		},
		{
			name: "calculates correct basin size, advent of code example 3",
			co:   helpers.Co{X: 0, Y: 3},
			want: 14,
		},
		{
			name: "calculates correct basin size, advent of code example 4",
			co:   helpers.Co{X: 9, Y: 4},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adventOfCodeExampleHeightMap.calculateBasin(tt.co); got != tt.want {
				t.Errorf("HeightMap.calculateBasin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasins_multiplyLargestBasinSizes(t *testing.T) {
	tests := []struct {
		name string
		bs   Basins
		want int
	}{
		{
			name: "returns product of largest 3 basins, advent of code example",
			bs:   Basins{9, 14, 3, 9},
			want: 1134,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bs.multiplyLargestBasinSizes(); got != tt.want {
				t.Errorf("Basins.multiplyLargestBasinSizes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSolutions(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
		want1 int
	}{
		{
			name: "returns correct part1 and part2 solutions from input, advent of code example",
			input: []string{
				"2199943210",
				"3987894921",
				"9856789892",
				"8767896789",
				"9899965678",
			},
			want:  15,
			want1: 1134,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findSolutions(tt.input)
			if got != tt.want {
				t.Errorf("findSolutions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findSolutions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
