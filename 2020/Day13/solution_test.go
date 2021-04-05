package main

import (
	"reflect"
	"testing"
)

var exampleBuses = Buses{
	Bus{
		id:     7,
		offset: 0,
	},
	Bus{
		id:     13,
		offset: 1,
	},
	Bus{
		id:     59,
		offset: 4,
	},
	Bus{
		id:     31,
		offset: 6,
	},
	Bus{
		id:     19,
		offset: 7,
	},
}

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name    string
		entries []string
		want    int
		want1   Buses
		wantErr bool
	}{
		{
			name: "returns an error if the timestamp cannot be converted to an int",
			entries: []string{
				"12883vfh12",
				"1, x, 2, x, x, 3, x",
			},
			want:    0,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "advent of code example 1",
			entries: []string{
				"939",
				"7,13,x,x,59,x,31,19",
			},
			want:    939,
			want1:   exampleBuses,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseInput(tt.entries)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseInput() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseInput() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBuses_part1(t *testing.T) {
	tests := []struct {
		name        string
		b           *Buses
		arrivalTime int
		want        int
	}{
		{
			name:        "advent of code example 1",
			b:           &exampleBuses,
			arrivalTime: 939,
			want:        295,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.part1(tt.arrivalTime); got != tt.want {
				t.Errorf("Buses.part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuses_part2(t *testing.T) {
	tests := []struct {
		name string
		b    *Buses
		want int
	}{
		{
			name: "advent of code example",
			b:    &exampleBuses,
			want: 1068781,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.part2(); got != tt.want {
				t.Errorf("Buses.part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
