package main

import (
	"reflect"
	"testing"
)

func Test_makeCave(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want *Cave
	}{
		{
			name: "returns a constructed small cave from a given id",
			id:   "start",
			want: &Cave{
				id:         "start",
				neighbours: map[string]*Cave{},
				small:      true,
			},
		},
		{
			name: "returns a constructed large cave from a given id",
			id:   "ABC",
			want: &Cave{
				id:         "ABC",
				neighbours: map[string]*Cave{},
				small:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeCave(tt.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeCave() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createTestSystem() *System {
	s := System{
		caves: map[string]*Cave{
			"start": {
				id:    "start",
				small: true,
			},
			"A": {
				id:    "A",
				small: false,
			},
			"b": {
				id:    "b",
				small: true,
			},
			"c": {
				id:    "c",
				small: true,
			},
			"d": {
				id:    "d",
				small: true,
			},
			"end": {
				id:    "end",
				small: true,
			},
		},
	}
	s.caves["start"].neighbours = map[string]*Cave{
		"A": s.caves["A"],
		"b": s.caves["b"],
	}
	s.caves["A"].neighbours = map[string]*Cave{
		"start": s.caves["start"],
		"b":     s.caves["b"],
		"c":     s.caves["c"],
		"end":   s.caves["end"],
	}
	s.caves["b"].neighbours = map[string]*Cave{
		"start": s.caves["start"],
		"A":     s.caves["A"],
		"d":     s.caves["d"],
		"end":   s.caves["end"],
	}
	s.caves["c"].neighbours = map[string]*Cave{
		"A": s.caves["A"],
	}
	s.caves["d"].neighbours = map[string]*Cave{
		"b": s.caves["b"],
	}
	s.caves["end"].neighbours = map[string]*Cave{
		"A": s.caves["A"],
		"b": s.caves["b"],
	}
	return &s
}

var aocSystem = createTestSystem()

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    *System
		wantErr bool
	}{
		{
			name: "returns an error if a line doesn't consist of two neighbours",
			input: []string{
				"start-A",
				"start-b",
				"A-c",
				"A-b",
				"b-d-c",
				"A-end",
				"b-end",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "compiles a system from the given input, advent of code example",
			input: []string{
				"start-A",
				"start-b",
				"A-c",
				"A-b",
				"b-d",
				"A-end",
				"b-end",
			},
			want:    aocSystem,
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
			if tt.wantErr {
				return
			}
			if len(got.caves) != len(tt.want.caves) {
				t.Errorf("parseInput() caves = %v, want %v", got.caves, tt.want.caves)
			}
			for id, gCave := range got.caves {
				wCave := tt.want.caves[id]
				if gCave.id != wCave.id {
					t.Errorf("parseInput() caves[%s] = %v, want %v", id, gCave, wCave)
				}
				if gCave.small != wCave.small {
					t.Errorf("parseInput() caves[%s] = %v, want %v", id, gCave, wCave)
				}
				if len(gCave.neighbours) != len(wCave.neighbours) {
					t.Errorf("parseInput() caves[%s] = %v, want %v", id, gCave, wCave)
				}
				for gID, gNghbr := range gCave.neighbours {
					if wNghbr, ok := wCave.neighbours[gID]; !ok {
						t.Errorf("parseInput() caves[%s] = %v, want %v", id, gCave, wCave)
					} else if gNghbr.id != wNghbr.id {
						t.Errorf("parseInput() caves[%s] = %v, want %v", id, gCave, wCave)
					}
				}
			}
		})
	}
}

func TestSystem_canVisitPart1(t *testing.T) {
	type args struct {
		cave *Cave
		path Path
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns true if cave is not small",
			args: args{
				cave: aocSystem.caves["A"],
			},
			want: true,
		},
		{
			name: "returns false if cave is small and already appears in the path",
			args: args{
				cave: aocSystem.caves["b"],
				path: Path{
					aocSystem.caves["start"],
					aocSystem.caves["A"],
					aocSystem.caves["b"],
					aocSystem.caves["d"],
				},
			},
			want: false,
		},
		{
			name: "returns true if cave is small and does not already appears in the path",
			args: args{
				cave: aocSystem.caves["c"],
				path: Path{
					aocSystem.caves["start"],
					aocSystem.caves["A"],
					aocSystem.caves["b"],
					aocSystem.caves["A"],
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canVisitPart1(tt.args.cave, tt.args.path); got != tt.want {
				t.Errorf("System.canVisitPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canVisitPart2(t *testing.T) {
	type args struct {
		cave *Cave
		path Path
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `returns false if cave id is "start"`,
			args: args{
				cave: aocSystem.caves["start"],
				path: Path{
					aocSystem.caves["start"],
					aocSystem.caves["A"],
					aocSystem.caves["b"],
				},
			},
			want: false,
		},
		{
			name: "returns true if cave is not small",
			args: args{
				cave: aocSystem.caves["A"],
				path: Path{
					aocSystem.caves["start"],
					aocSystem.caves["A"],
					aocSystem.caves["b"],
				},
			},
			want: true,
		},
		{
			name: "returns false if cave is small and has already been traversed to twice",
			args: args{
				cave: aocSystem.caves["b"],
				path: Path{
					aocSystem.caves["start"],
					aocSystem.caves["A"],
					aocSystem.caves["c"],
					aocSystem.caves["A"],
					aocSystem.caves["b"],
					aocSystem.caves["d"],
					aocSystem.caves["b"],
					aocSystem.caves["A"],
				},
			},
			want: false,
		},
		{
			name: "returns false if cave is small, has already been traversed to and a different cave has already been traversed to twice",
			args: args{
				cave: aocSystem.caves["c"],
				path: Path{
					aocSystem.caves["start"],
					aocSystem.caves["A"],
					aocSystem.caves["c"],
					aocSystem.caves["A"],
					aocSystem.caves["b"],
					aocSystem.caves["d"],
					aocSystem.caves["b"],
					aocSystem.caves["A"],
				},
			},
			want: false,
		},
		{
			name: "returns true if cave is small, has already been traversed to and no other cave has already been traversed to twice",
			args: args{
				cave: aocSystem.caves["c"],
				path: Path{
					aocSystem.caves["start"],
					aocSystem.caves["A"],
					aocSystem.caves["c"],
					aocSystem.caves["A"],
					aocSystem.caves["d"],
					aocSystem.caves["b"],
					aocSystem.caves["A"],
				},
			},
			want: true,
		},
		{
			name: "returns true if cave is small, has not already been traversed to and another cave has already been traversed to twice",
			args: args{
				cave: aocSystem.caves["c"],
				path: Path{
					aocSystem.caves["start"],
					aocSystem.caves["A"],
					aocSystem.caves["b"],
					aocSystem.caves["A"],
					aocSystem.caves["d"],
					aocSystem.caves["b"],
					aocSystem.caves["A"],
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canVisitPart2(tt.args.cave, tt.args.path); got != tt.want {
				t.Errorf("canVisitPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystem_findNumberOfPaths(t *testing.T) {
	tests := []struct {
		name     string
		canVisit canVisit
		want     int
	}{
		{
			name:     "finds number of paths part 1, advent of code example",
			canVisit: canVisitPart1,
			want:     10,
		},
		{
			name:     "finds number of paths part 2, advent of code example",
			canVisit: canVisitPart2,
			want:     36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &System{
				caves: aocSystem.caves,
				paths: []Path{},
			}
			if got := s.findNumberOfPaths(tt.canVisit); got != tt.want {
				t.Errorf("System.findNumberOfPaths() = %v, want %v", got, tt.want)
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
			name: "returns an error if an input line cannot be parsed",
			input: []string{
				"start-A",
				"b-A",
				"c-A",
				"b-A-d",
				"end-A",
				"end-b",
			},
			want:    -1,
			want1:   -1,
			wantErr: true,
		},
		{
			name: "returns solutions for part1 and part2, advent of code example 1",
			input: []string{
				"start-A",
				"start-b",
				"A-c",
				"A-b",
				"b-d",
				"A-end",
				"b-end",
			},
			want:    10,
			want1:   36,
			wantErr: false,
		},
		{
			name: "returns solutions for part1 and part2, advent of code example 2",
			input: []string{
				"dc-end",
				"HN-start",
				"start-kj",
				"dc-start",
				"dc-HN",
				"LN-dc",
				"HN-end",
				"kj-sa",
				"kj-HN",
				"kj-dc",
			},
			want:    19,
			want1:   103,
			wantErr: false,
		},
		{
			name: "returns solutions for part1 and part2, advent of code example 3",
			input: []string{
				"fs-end",
				"he-DX",
				"fs-he",
				"start-DX",
				"pj-DX",
				"end-zg",
				"zg-sl",
				"zg-pj",
				"pj-he",
				"RW-he",
				"fs-DX",
				"pj-RW",
				"zg-RW",
				"start-pj",
				"he-WI",
				"zg-he",
				"pj-fs",
				"start-RW",
			},
			want:    226,
			want1:   3509,
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
