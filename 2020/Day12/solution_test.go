package main

import (
	"testing"
)

func Test_parseDirection(t *testing.T) {
	tests := []struct {
		name    string
		entry   string
		want    string
		want1   int
		wantErr bool
	}{
		{
			name:    "returns an error if the input can't be parsed into an int",
			entry:   "F2?3",
			want:    "F",
			want1:   0,
			wantErr: true,
		},
		{
			name:    "returns correctly parsed input",
			entry:   "F23",
			want:    "F",
			want1:   23,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseDirection(tt.entry)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDirection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseDirection() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseDirection() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name    string
		entries []string
		want    int
		wantErr bool
	}{
		{
			name: "returns an error if one of the entries can't be parsed successfully",
			entries: []string{
				"F10",
				"N3",
				"F7.",
				"R90",
				"F11",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "advent of code example 1",
			entries: []string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			want:    25,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.entries)
			if (err != nil) != tt.wantErr {
				t.Errorf("part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name    string
		entries []string
		want    int
		wantErr bool
	}{
		{
			name: "returns an error if one of the entries can't be parsed correctly",
			entries: []string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F!11",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "advent of code example 1",
			entries: []string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			want:    286,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.entries)
			if (err != nil) != tt.wantErr {
				t.Errorf("part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
