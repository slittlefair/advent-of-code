package main

import (
	"Advent-of-Code/utils"
	"reflect"
	"testing"
)

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *TargetArea
		wantErr bool
	}{
		{
			name:    "returns an error if there are fewer than 4 numbers",
			input:   "target area: x=562, y=-98..613",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "returns an error if there are more than 4 numbers",
			input:   "target area: x=562..872, y=-98..613..614",
			want:    nil,
			wantErr: true,
		},
		{
			name:  "returns the correct target area from inout, advent of code example",
			input: "target area: x=20..30, y=-10..-5",
			want: &TargetArea{
				MinX:                20,
				MaxX:                30,
				MinY:                -10,
				MaxY:                -5,
				GreatestSuccessfulY: -5,
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

func TestTargetArea_isInTargetArea(t *testing.T) {
	tests := []struct {
		name string
		co   utils.Co
		want bool
	}{
		{
			name: "returns false if co X value too low for target area",
			co:   utils.Co{X: 10, Y: -8},
			want: false,
		},
		{
			name: "returns false if co XX value too high for target area",
			co:   utils.Co{X: 100, Y: -8},
			want: false,
		},
		{
			name: "returns false if co Y value too low for target area",
			co:   utils.Co{X: 23, Y: 9},
			want: false,
		},
		{
			name: "returns false if co Y value too high for target area",
			co:   utils.Co{X: 24, Y: -89},
			want: false,
		},
		{
			name: "returns true if co is in target area",
			co:   utils.Co{X: 20, Y: -8},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := TargetArea{
				MinX:                20,
				MaxX:                30,
				MinY:                -10,
				MaxY:                -5,
				GreatestSuccessfulY: -5,
			}
			if got := ta.isInTargetArea(tt.co); got != tt.want {
				t.Errorf("TargetArea.isInTargetArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTargetArea_evaluateTrajectory(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		ta   *TargetArea
		args args
		want *TargetArea
	}{
		{
			name: "returns an unchanged target area if the given coordinate won't reach target x values",
			ta: &TargetArea{
				MinX:                20,
				MaxX:                30,
				MinY:                -10,
				MaxY:                -5,
				GreatestSuccessfulY: -5,
			},
			args: args{x: 5, y: 3},
			want: &TargetArea{
				MinX:                20,
				MaxX:                30,
				MinY:                -10,
				MaxY:                -5,
				GreatestSuccessfulY: -5,
			},
		},
		{
			name: "returns an unchanged target area if the given coordinate overshoots target x values",
			ta: &TargetArea{
				MinX:                20,
				MaxX:                30,
				MinY:                -10,
				MaxY:                -5,
				GreatestSuccessfulY: -5,
			},
			args: args{x: 12, y: 3},
			want: &TargetArea{
				MinX:                20,
				MaxX:                30,
				MinY:                -10,
				MaxY:                -5,
				GreatestSuccessfulY: -5,
			},
		},
		{
			name: "returns an unchanged target area if the given coordinate overshoots target x values",
			ta: &TargetArea{
				MinX:                20,
				MaxX:                30,
				MinY:                -10,
				MaxY:                -5,
				GreatestSuccessfulY: -5,
			},
			args: args{x: 6, y: 10},
			want: &TargetArea{
				MinX:                20,
				MaxX:                30,
				MinY:                -10,
				MaxY:                -5,
				GreatestSuccessfulY: -5,
			},
		},
		{
			name: "returns a changed target area if the given velocity ends up landing in the target area",
			ta: &TargetArea{
				MinX:                   20,
				MaxX:                   30,
				MinY:                   -10,
				MaxY:                   -5,
				GreatestSuccessfulY:    -5,
				SuccessfulTrajectories: 23,
			},
			args: args{x: 7, y: 8},
			want: &TargetArea{
				MinX:                   20,
				MaxX:                   30,
				MinY:                   -10,
				MaxY:                   -5,
				GreatestSuccessfulY:    36,
				SuccessfulTrajectories: 24,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta := tt.ta
			ta.evaluateTrajectory(tt.args.x, tt.args.y)
			if !reflect.DeepEqual(ta, tt.want) {
				t.Errorf("TargetArea.evaluateTrajectory() = %v, want %v", ta, tt.want)
			}
		})
	}
}

func Test_findTrajectories(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    int
		want1   int
		wantErr bool
	}{
		{
			name:    "returns an error if input is less than one string long",
			input:   []string{},
			want:    -1,
			want1:   -1,
			wantErr: true,
		},
		{
			name: "returns an error if input is greater than one string long",
			input: []string{
				"target area: x=20..30, y=-10..-5",
				"target area: x=20..30, y=-10..-6",
			},
			want:    -1,
			want1:   -1,
			wantErr: true,
		},
		{
			name: "returns an error if input cannot be parsed",
			input: []string{
				"target area: x=20..30, y=-10..-5..-1",
			},
			want:    -1,
			want1:   -1,
			wantErr: true,
		},
		{
			name: "returns correct part 1 and part 2 solutions for given input, advent of code example",
			input: []string{
				"target area: x=20..30, y=-10..-5",
			},
			want:    45,
			want1:   112,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := findTrajectories(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("findTrajectories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findTrajectories() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findTrajectories() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
