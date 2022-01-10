package main

import (
	utils "Advent-of-Code/utils"
	"reflect"
	"testing"
)

func Test_createVisitedHouses(t *testing.T) {
	tests := []struct {
		name string
		want *VisitedHouses
	}{
		{
			name: "creates a basic VisitedHouse object",
			want: &VisitedHouses{
				Map: map[utils.Co]bool{
					{X: 0, Y: 0}: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createVisitedHouses(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createVisitedHouses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVisitedHouses_moveSanta(t *testing.T) {
	tests := []struct {
		name  string
		Santa utils.Co
		dir   string
		want  utils.Co
	}{
		{
			name:  "subtracts 1 from currentHouse.X if moving west",
			Santa: utils.Co{X: 0, Y: 0},
			dir:   "<",
			want:  utils.Co{X: -1, Y: 0},
		},
		{
			name:  "adds 1 to currentHouse.X if moving east",
			Santa: utils.Co{X: 0, Y: 0},
			dir:   ">",
			want:  utils.Co{X: 1, Y: 0},
		},
		{
			name:  "subtracts 1 from currentHouse.Y if moving north",
			Santa: utils.Co{X: 0, Y: 0},
			dir:   "^",
			want:  utils.Co{X: 0, Y: -1},
		},
		{
			name:  "adds 1 from currentHouse.Y if moving south",
			Santa: utils.Co{X: 0, Y: 0},
			dir:   "v",
			want:  utils.Co{X: 0, Y: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vh := &VisitedHouses{
				Santa: tt.Santa,
			}
			vh.moveSanta(tt.dir)
			if !reflect.DeepEqual(vh.Santa, tt.want) {
				t.Errorf("VisitedHouses.moveSanta() = %v, want %v", vh.Santa, tt.want)
			}
		})
	}
}

func TestVisitedHouses_moveRoboSanta(t *testing.T) {
	tests := []struct {
		name      string
		RoboSanta utils.Co
		dir       string
		want      utils.Co
	}{
		{
			name:      "subtracts 1 from currentHouse.X if moving west",
			RoboSanta: utils.Co{X: 0, Y: 0},
			dir:       "<",
			want:      utils.Co{X: -1, Y: 0},
		},
		{
			name:      "adds 1 to currentHouse.X if moving east",
			RoboSanta: utils.Co{X: 0, Y: 0},
			dir:       ">",
			want:      utils.Co{X: 1, Y: 0},
		},
		{
			name:      "subtracts 1 from currentHouse.Y if moving north",
			RoboSanta: utils.Co{X: 0, Y: 0},
			dir:       "^",
			want:      utils.Co{X: 0, Y: -1},
		},
		{
			name:      "adds 1 from currentHouse.Y if moving south",
			RoboSanta: utils.Co{X: 0, Y: 0},
			dir:       "v",
			want:      utils.Co{X: 0, Y: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vh := &VisitedHouses{
				RoboSanta: tt.RoboSanta,
			}
			vh.moveRoboSanta(tt.dir)
			if !reflect.DeepEqual(vh.RoboSanta, tt.want) {
				t.Errorf("VisitedHouses.moveRoboSanta() = %v, want %v", vh.RoboSanta, tt.want)
			}
		})
	}
}

func TestVisitedHouses_alreadyVisitedHouse(t *testing.T) {
	tests := []struct {
		name  string
		Map   map[utils.Co]bool
		santa utils.Co
		want  bool
	}{
		{
			name: "returns true if the current house has been seen",
			Map: map[utils.Co]bool{
				{X: 0, Y: 0}:  true,
				{X: 1, Y: 99}: true,
				{X: -2, Y: 1}: true,
				{X: 1, Y: 1}:  true,
			},
			santa: utils.Co{X: 1, Y: 99},
			want:  true,
		},
		{
			name: "returns false if the current house has not been seen",
			Map: map[utils.Co]bool{
				{X: 0, Y: 0}:  true,
				{X: 1, Y: 99}: true,
				{X: -2, Y: 1}: true,
				{X: 1, Y: 1}:  true,
			},
			santa: utils.Co{X: 2, Y: 1},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vh := &VisitedHouses{
				Map: tt.Map,
			}
			if got := vh.alreadyVisitedHouse(tt.santa); got != tt.want {
				t.Errorf("VisitedHouses.alreadyVisitedHouse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVisitedHouses_countUniqueHousesVisited(t *testing.T) {
	tests := []struct {
		name  string
		input string
		part1 bool
		want  int
	}{
		{
			name:  "advent of code example 1, part1",
			input: ">",
			part1: true,
			want:  2,
		},
		{
			name:  "advent of code example 2, part1",
			input: "^>v<",
			part1: true,
			want:  4,
		},
		{
			name:  "advent of code example 3, part1",
			input: "^v^v^v^v^v",
			part1: true,
			want:  2,
		},
		{
			name:  "advent of code example 1, part2",
			input: "^v",
			part1: false,
			want:  3,
		},
		{
			name:  "advent of code example 2, part2",
			input: "^>v<",
			part1: false,
			want:  3,
		},
		{
			name:  "advent of code example 3, part2",
			input: "^v^v^v^v^v",
			part1: false,
			want:  11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vh := &VisitedHouses{
				Map: map[utils.Co]bool{
					{X: 0, Y: 0}: true,
				},
			}
			if got := vh.countUniqueHousesVisited(tt.input, tt.part1); got != tt.want {
				t.Errorf("VisitedHouses.countUniqueHousesVisited() = %v, want %v", got, tt.want)
			}
		})
	}
}
