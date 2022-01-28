package main

import (
	"Advent-of-Code/graph"
	"reflect"
	"testing"
)

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		want Grid
	}{
		{
			name: "correctly parses a simple grid",
			arg: []string{
				"#.#",
				"..#",
			},
			want: Grid{
				Height: 1,
				Width:  2,
				Lights: Lights{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: "#",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_LightStaysOn(t *testing.T) {
	type fields struct {
		Lights Lights
		Height int
		Width  int
	}
	tests := []struct {
		name   string
		fields fields
		arg    graph.Co
		want   bool
	}{
		{
			name: "returns false if an on light has fewer than two on neighbours",
			fields: fields{
				Lights: Lights{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: "#",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: "#",
				},
			},
			arg:  graph.Co{X: 0, Y: 0},
			want: false,
		},
		{
			name: "returns false if an on light has more than three on neighbours",
			fields: fields{
				Lights: Lights{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: "#",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: "#",
				},
			},
			arg:  graph.Co{X: 1, Y: 1},
			want: false,
		},
		{
			name: "returns false if an on light has more than three on neighbours",
			fields: fields{
				Lights: Lights{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: "#",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: "#",
				},
			},
			arg:  graph.Co{X: 1, Y: 1},
			want: false,
		},
		{
			name: "returns true if an on light has two on neighbours",
			fields: fields{
				Lights: Lights{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: "#",
					{X: 2, Y: 1}: ".",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: ".",
				},
			},
			arg:  graph.Co{X: 1, Y: 1},
			want: true,
		},
		{
			name: "returns true if an on light has two on neighbours",
			fields: fields{
				Lights: Lights{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: "#",
					{X: 2, Y: 1}: ".",
					{X: 0, Y: 2}: "#",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: ".",
				},
			},
			arg:  graph.Co{X: 1, Y: 1},
			want: true,
		},
		{
			name: "returns false if an off light has fewer than 3 on neighbours",
			fields: fields{
				Lights: Lights{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: ".",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: ".",
				},
			},
			arg:  graph.Co{X: 1, Y: 1},
			want: false,
		},
		{
			name: "returns false if an off light has more than 3 on neighbours",
			fields: fields{
				Lights: Lights{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: ".",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: "#",
					{X: 2, Y: 2}: "#",
				},
			},
			arg:  graph.Co{X: 1, Y: 1},
			want: false,
		},
		{
			name: "returns true if an off light has 3 on neighbours",
			fields: fields{
				Lights: Lights{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: ".",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: "#",
				},
			},
			arg:  graph.Co{X: 1, Y: 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Lights: tt.fields.Lights,
				Height: tt.fields.Height,
				Width:  tt.fields.Width,
			}
			if got := g.LightStaysOn(tt.arg); got != tt.want {
				t.Errorf("Grid.LightStaysOn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_ChangeLights(t *testing.T) {
	type fields struct {
		Lights Lights
		Height int
		Width  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Grid
	}{
		{
			name: "correctly changes lights, advent of code example 1",
			fields: fields{
				Lights: Lights{
					{X: 0, Y: 0}: ".",
					{X: 1, Y: 0}: "#",
					{X: 2, Y: 0}: ".",
					{X: 3, Y: 0}: "#",
					{X: 4, Y: 0}: ".",
					{X: 5, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: ".",
					{X: 3, Y: 1}: "#",
					{X: 4, Y: 1}: "#",
					{X: 5, Y: 1}: ".",
					{X: 0, Y: 2}: "#",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: ".",
					{X: 3, Y: 2}: ".",
					{X: 4, Y: 2}: ".",
					{X: 5, Y: 2}: "#",
					{X: 0, Y: 3}: ".",
					{X: 1, Y: 3}: ".",
					{X: 2, Y: 3}: "#",
					{X: 3, Y: 3}: ".",
					{X: 4, Y: 3}: ".",
					{X: 5, Y: 3}: ".",
					{X: 0, Y: 4}: "#",
					{X: 1, Y: 4}: ".",
					{X: 2, Y: 4}: "#",
					{X: 3, Y: 4}: ".",
					{X: 4, Y: 4}: ".",
					{X: 5, Y: 4}: "#",
					{X: 0, Y: 5}: "#",
					{X: 1, Y: 5}: "#",
					{X: 2, Y: 5}: "#",
					{X: 3, Y: 5}: "#",
					{X: 4, Y: 5}: ".",
					{X: 5, Y: 5}: ".",
				},
			},
			want: &Grid{
				Lights: Lights{
					{X: 0, Y: 0}: ".",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 3, Y: 0}: "#",
					{X: 4, Y: 0}: ".",
					{X: 5, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: "#",
					{X: 3, Y: 1}: "#",
					{X: 4, Y: 1}: ".",
					{X: 5, Y: 1}: "#",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: ".",
					{X: 3, Y: 2}: "#",
					{X: 4, Y: 2}: "#",
					{X: 5, Y: 2}: ".",
					{X: 0, Y: 3}: ".",
					{X: 1, Y: 3}: ".",
					{X: 2, Y: 3}: ".",
					{X: 3, Y: 3}: ".",
					{X: 4, Y: 3}: ".",
					{X: 5, Y: 3}: ".",
					{X: 0, Y: 4}: "#",
					{X: 1, Y: 4}: ".",
					{X: 2, Y: 4}: ".",
					{X: 3, Y: 4}: ".",
					{X: 4, Y: 4}: ".",
					{X: 5, Y: 4}: ".",
					{X: 0, Y: 5}: "#",
					{X: 1, Y: 5}: ".",
					{X: 2, Y: 5}: "#",
					{X: 3, Y: 5}: "#",
					{X: 4, Y: 5}: ".",
					{X: 5, Y: 5}: ".",
				},
			},
		},
		{
			name: "correctly changes lights, advent of code example 2",
			fields: fields{
				Lights: Lights{
					{X: 0, Y: 0}: ".",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 3, Y: 0}: "#",
					{X: 4, Y: 0}: ".",
					{X: 5, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: "#",
					{X: 3, Y: 1}: "#",
					{X: 4, Y: 1}: ".",
					{X: 5, Y: 1}: "#",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: ".",
					{X: 3, Y: 2}: "#",
					{X: 4, Y: 2}: "#",
					{X: 5, Y: 2}: ".",
					{X: 0, Y: 3}: ".",
					{X: 1, Y: 3}: ".",
					{X: 2, Y: 3}: ".",
					{X: 3, Y: 3}: ".",
					{X: 4, Y: 3}: ".",
					{X: 5, Y: 3}: ".",
					{X: 0, Y: 4}: "#",
					{X: 1, Y: 4}: ".",
					{X: 2, Y: 4}: ".",
					{X: 3, Y: 4}: ".",
					{X: 4, Y: 4}: ".",
					{X: 5, Y: 4}: ".",
					{X: 0, Y: 5}: "#",
					{X: 1, Y: 5}: ".",
					{X: 2, Y: 5}: "#",
					{X: 3, Y: 5}: "#",
					{X: 4, Y: 5}: ".",
					{X: 5, Y: 5}: ".",
				},
			},
			want: &Grid{
				Lights: Lights{
					{X: 0, Y: 0}: ".",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 3, Y: 0}: "#",
					{X: 4, Y: 0}: "#",
					{X: 5, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: ".",
					{X: 3, Y: 1}: ".",
					{X: 4, Y: 1}: ".",
					{X: 5, Y: 1}: ".",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: "#",
					{X: 3, Y: 2}: "#",
					{X: 4, Y: 2}: "#",
					{X: 5, Y: 2}: ".",
					{X: 0, Y: 3}: ".",
					{X: 1, Y: 3}: ".",
					{X: 2, Y: 3}: ".",
					{X: 3, Y: 3}: ".",
					{X: 4, Y: 3}: ".",
					{X: 5, Y: 3}: ".",
					{X: 0, Y: 4}: ".",
					{X: 1, Y: 4}: "#",
					{X: 2, Y: 4}: ".",
					{X: 3, Y: 4}: ".",
					{X: 4, Y: 4}: ".",
					{X: 5, Y: 4}: ".",
					{X: 0, Y: 5}: ".",
					{X: 1, Y: 5}: "#",
					{X: 2, Y: 5}: ".",
					{X: 3, Y: 5}: ".",
					{X: 4, Y: 5}: ".",
					{X: 5, Y: 5}: ".",
				},
			},
		},
		{
			name: "correctly changes lights, advent of code example 3",
			fields: fields{
				Lights: Lights{
					{X: 0, Y: 0}: ".",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 3, Y: 0}: "#",
					{X: 4, Y: 0}: "#",
					{X: 5, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: ".",
					{X: 3, Y: 1}: ".",
					{X: 4, Y: 1}: ".",
					{X: 5, Y: 1}: ".",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: "#",
					{X: 3, Y: 2}: "#",
					{X: 4, Y: 2}: "#",
					{X: 5, Y: 2}: ".",
					{X: 0, Y: 3}: ".",
					{X: 1, Y: 3}: ".",
					{X: 2, Y: 3}: ".",
					{X: 3, Y: 3}: ".",
					{X: 4, Y: 3}: ".",
					{X: 5, Y: 3}: ".",
					{X: 0, Y: 4}: ".",
					{X: 1, Y: 4}: "#",
					{X: 2, Y: 4}: ".",
					{X: 3, Y: 4}: ".",
					{X: 4, Y: 4}: ".",
					{X: 5, Y: 4}: ".",
					{X: 0, Y: 5}: ".",
					{X: 1, Y: 5}: "#",
					{X: 2, Y: 5}: ".",
					{X: 3, Y: 5}: ".",
					{X: 4, Y: 5}: ".",
					{X: 5, Y: 5}: ".",
				},
			},
			want: &Grid{
				Lights: Lights{
					{X: 0, Y: 0}: ".",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: ".",
					{X: 3, Y: 0}: "#",
					{X: 4, Y: 0}: ".",
					{X: 5, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: ".",
					{X: 3, Y: 1}: ".",
					{X: 4, Y: 1}: ".",
					{X: 5, Y: 1}: ".",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: ".",
					{X: 3, Y: 2}: "#",
					{X: 4, Y: 2}: ".",
					{X: 5, Y: 2}: ".",
					{X: 0, Y: 3}: ".",
					{X: 1, Y: 3}: ".",
					{X: 2, Y: 3}: "#",
					{X: 3, Y: 3}: "#",
					{X: 4, Y: 3}: ".",
					{X: 5, Y: 3}: ".",
					{X: 0, Y: 4}: ".",
					{X: 1, Y: 4}: ".",
					{X: 2, Y: 4}: ".",
					{X: 3, Y: 4}: ".",
					{X: 4, Y: 4}: ".",
					{X: 5, Y: 4}: ".",
					{X: 0, Y: 5}: ".",
					{X: 1, Y: 5}: ".",
					{X: 2, Y: 5}: ".",
					{X: 3, Y: 5}: ".",
					{X: 4, Y: 5}: ".",
					{X: 5, Y: 5}: ".",
				},
			},
		},
		{
			name: "correctly changes lights, advent of code example 4",
			fields: fields{
				Lights: Lights{
					{X: 0, Y: 0}: ".",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: ".",
					{X: 3, Y: 0}: "#",
					{X: 4, Y: 0}: ".",
					{X: 5, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: ".",
					{X: 3, Y: 1}: ".",
					{X: 4, Y: 1}: ".",
					{X: 5, Y: 1}: ".",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: ".",
					{X: 3, Y: 2}: "#",
					{X: 4, Y: 2}: ".",
					{X: 5, Y: 2}: ".",
					{X: 0, Y: 3}: ".",
					{X: 1, Y: 3}: ".",
					{X: 2, Y: 3}: "#",
					{X: 3, Y: 3}: "#",
					{X: 4, Y: 3}: ".",
					{X: 5, Y: 3}: ".",
					{X: 0, Y: 4}: ".",
					{X: 1, Y: 4}: ".",
					{X: 2, Y: 4}: ".",
					{X: 3, Y: 4}: ".",
					{X: 4, Y: 4}: ".",
					{X: 5, Y: 4}: ".",
					{X: 0, Y: 5}: ".",
					{X: 1, Y: 5}: ".",
					{X: 2, Y: 5}: ".",
					{X: 3, Y: 5}: ".",
					{X: 4, Y: 5}: ".",
					{X: 5, Y: 5}: ".",
				},
			},
			want: &Grid{
				Lights: Lights{
					{X: 0, Y: 0}: ".",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: ".",
					{X: 3, Y: 0}: ".",
					{X: 4, Y: 0}: ".",
					{X: 5, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: ".",
					{X: 3, Y: 1}: ".",
					{X: 4, Y: 1}: ".",
					{X: 5, Y: 1}: ".",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: "#",
					{X: 3, Y: 2}: "#",
					{X: 4, Y: 2}: ".",
					{X: 5, Y: 2}: ".",
					{X: 0, Y: 3}: ".",
					{X: 1, Y: 3}: ".",
					{X: 2, Y: 3}: "#",
					{X: 3, Y: 3}: "#",
					{X: 4, Y: 3}: ".",
					{X: 5, Y: 3}: ".",
					{X: 0, Y: 4}: ".",
					{X: 1, Y: 4}: ".",
					{X: 2, Y: 4}: ".",
					{X: 3, Y: 4}: ".",
					{X: 4, Y: 4}: ".",
					{X: 5, Y: 4}: ".",
					{X: 0, Y: 5}: ".",
					{X: 1, Y: 5}: ".",
					{X: 2, Y: 5}: ".",
					{X: 3, Y: 5}: ".",
					{X: 4, Y: 5}: ".",
					{X: 5, Y: 5}: ".",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Lights: tt.fields.Lights,
				Height: tt.fields.Height,
				Width:  tt.fields.Width,
			}
			g.ChangeLights()
			if !reflect.DeepEqual(g, tt.want) {
				t.Errorf("Grid.ChangeLights() = %v, want %v", g, tt.want)
			}
		})
	}
}

func TestGrid_CountLightsOn(t *testing.T) {
	type fields struct {
		Lights Lights
		Height int
		Width  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "returns 0 for an empty map of lights",
			want: 0,
		},
		{
			name: "returns 0 for an empty map of lights",
			fields: fields{
				Lights: Lights{
					{X: 0, Y: 0}: ".",
					{X: 1, Y: 0}: "#",
					{X: 2, Y: 0}: ".",
					{X: 3, Y: 0}: "#",
					{X: 4, Y: 0}: ".",
					{X: 5, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: ".",
					{X: 3, Y: 1}: "#",
					{X: 4, Y: 1}: "#",
					{X: 5, Y: 1}: ".",
					{X: 0, Y: 2}: "#",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: ".",
					{X: 3, Y: 2}: ".",
					{X: 4, Y: 2}: ".",
					{X: 5, Y: 2}: "#",
					{X: 0, Y: 3}: ".",
					{X: 1, Y: 3}: ".",
					{X: 2, Y: 3}: "#",
					{X: 3, Y: 3}: ".",
					{X: 4, Y: 3}: ".",
					{X: 5, Y: 3}: ".",
					{X: 0, Y: 4}: "#",
					{X: 1, Y: 4}: ".",
					{X: 2, Y: 4}: "#",
					{X: 3, Y: 4}: ".",
					{X: 4, Y: 4}: ".",
					{X: 5, Y: 4}: "#",
					{X: 0, Y: 5}: "#",
					{X: 1, Y: 5}: "#",
					{X: 2, Y: 5}: "#",
					{X: 3, Y: 5}: "#",
					{X: 4, Y: 5}: ".",
					{X: 5, Y: 5}: ".",
				},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Grid{
				Lights: tt.fields.Lights,
				Height: tt.fields.Height,
				Width:  tt.fields.Width,
			}
			if got := g.CountLightsOn(); got != tt.want {
				t.Errorf("Grid.CountLightsOn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_TurnCornersOn(t *testing.T) {
	type fields struct {
		Lights Lights
		Height int
		Width  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Grid
	}{
		{
			name: "turns on the four corners of lights",
			fields: fields{
				Lights: Lights{
					{X: 0, Y: 0}: ".",
					{X: 1, Y: 0}: "#",
					{X: 2, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: "#",
					{X: 2, Y: 1}: ".",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: ".",
				},
				Height: 2,
				Width:  2,
			},
			want: &Grid{
				Lights: Lights{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: "#",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: "#",
					{X: 2, Y: 1}: ".",
					{X: 0, Y: 2}: "#",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: "#",
				},
				Height: 2,
				Width:  2,
			},
		},
		{
			name: "turns on the four corners of lights with some already on",
			fields: fields{
				Lights: Lights{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: "#",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: "#",
					{X: 2, Y: 1}: ".",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: ".",
				},
				Height: 2,
				Width:  2,
			},
			want: &Grid{
				Lights: Lights{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: "#",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: "#",
					{X: 2, Y: 1}: ".",
					{X: 0, Y: 2}: "#",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: "#",
				},
				Height: 2,
				Width:  2,
			},
		},
		{
			name: "turns on the four corners of lights with all already on",
			fields: fields{
				Lights: Lights{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: "#",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: "#",
					{X: 2, Y: 1}: ".",
					{X: 0, Y: 2}: "#",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: "#",
				},
				Height: 2,
				Width:  2,
			},
			want: &Grid{
				Lights: Lights{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: "#",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: "#",
					{X: 2, Y: 1}: ".",
					{X: 0, Y: 2}: "#",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: "#",
				},
				Height: 2,
				Width:  2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Lights: tt.fields.Lights,
				Height: tt.fields.Height,
				Width:  tt.fields.Width,
			}
			g.TurnCornersOn()
			if !reflect.DeepEqual(g, tt.want) {
				t.Errorf("Grid.TurnCornersOn() = %v, want %v", g, tt.want)
			}
		})
	}
}

func TestGrid_RunStepsPart1(t *testing.T) {
	type fields struct {
		Lights Lights
		Height int
		Width  int
	}
	tests := []struct {
		name   string
		fields fields
		arg    int
		want   *Grid
	}{
		{
			name: "runs steps for part 1, advent of code example",
			fields: fields{
				Lights: Lights{
					{X: 0, Y: 0}: ".",
					{X: 1, Y: 0}: "#",
					{X: 2, Y: 0}: ".",
					{X: 3, Y: 0}: "#",
					{X: 4, Y: 0}: ".",
					{X: 5, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: ".",
					{X: 3, Y: 1}: "#",
					{X: 4, Y: 1}: "#",
					{X: 5, Y: 1}: ".",
					{X: 0, Y: 2}: "#",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: ".",
					{X: 3, Y: 2}: ".",
					{X: 4, Y: 2}: ".",
					{X: 5, Y: 2}: "#",
					{X: 0, Y: 3}: ".",
					{X: 1, Y: 3}: ".",
					{X: 2, Y: 3}: "#",
					{X: 3, Y: 3}: ".",
					{X: 4, Y: 3}: ".",
					{X: 5, Y: 3}: ".",
					{X: 0, Y: 4}: "#",
					{X: 1, Y: 4}: ".",
					{X: 2, Y: 4}: "#",
					{X: 3, Y: 4}: ".",
					{X: 4, Y: 4}: ".",
					{X: 5, Y: 4}: "#",
					{X: 0, Y: 5}: "#",
					{X: 1, Y: 5}: "#",
					{X: 2, Y: 5}: "#",
					{X: 3, Y: 5}: "#",
					{X: 4, Y: 5}: ".",
					{X: 5, Y: 5}: ".",
				},
			},
			arg: 4,
			want: &Grid{
				Lights: Lights{
					{X: 0, Y: 0}: ".",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: ".",
					{X: 3, Y: 0}: ".",
					{X: 4, Y: 0}: ".",
					{X: 5, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: ".",
					{X: 3, Y: 1}: ".",
					{X: 4, Y: 1}: ".",
					{X: 5, Y: 1}: ".",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: "#",
					{X: 3, Y: 2}: "#",
					{X: 4, Y: 2}: ".",
					{X: 5, Y: 2}: ".",
					{X: 0, Y: 3}: ".",
					{X: 1, Y: 3}: ".",
					{X: 2, Y: 3}: "#",
					{X: 3, Y: 3}: "#",
					{X: 4, Y: 3}: ".",
					{X: 5, Y: 3}: ".",
					{X: 0, Y: 4}: ".",
					{X: 1, Y: 4}: ".",
					{X: 2, Y: 4}: ".",
					{X: 3, Y: 4}: ".",
					{X: 4, Y: 4}: ".",
					{X: 5, Y: 4}: ".",
					{X: 0, Y: 5}: ".",
					{X: 1, Y: 5}: ".",
					{X: 2, Y: 5}: ".",
					{X: 3, Y: 5}: ".",
					{X: 4, Y: 5}: ".",
					{X: 5, Y: 5}: ".",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Lights: tt.fields.Lights,
				Height: tt.fields.Height,
				Width:  tt.fields.Width,
			}
			g.RunStepsPart1(tt.arg)
			if !reflect.DeepEqual(g, tt.want) {
				t.Errorf("Grid.RunStepsPart1() = %v, want %v", g, tt.want)
			}
		})
	}
}

func TestGrid_RunStepsPart2(t *testing.T) {
	type fields struct {
		Lights Lights
		Height int
		Width  int
	}
	tests := []struct {
		name   string
		fields fields
		arg    int
		want   *Grid
	}{
		{
			name: "runs steps for part 2, advent of code example",
			fields: fields{
				Lights: Lights{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: "#",
					{X: 2, Y: 0}: ".",
					{X: 3, Y: 0}: "#",
					{X: 4, Y: 0}: ".",
					{X: 5, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: ".",
					{X: 3, Y: 1}: "#",
					{X: 4, Y: 1}: "#",
					{X: 5, Y: 1}: ".",
					{X: 0, Y: 2}: "#",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: ".",
					{X: 3, Y: 2}: ".",
					{X: 4, Y: 2}: ".",
					{X: 5, Y: 2}: "#",
					{X: 0, Y: 3}: ".",
					{X: 1, Y: 3}: ".",
					{X: 2, Y: 3}: "#",
					{X: 3, Y: 3}: ".",
					{X: 4, Y: 3}: ".",
					{X: 5, Y: 3}: ".",
					{X: 0, Y: 4}: "#",
					{X: 1, Y: 4}: ".",
					{X: 2, Y: 4}: "#",
					{X: 3, Y: 4}: ".",
					{X: 4, Y: 4}: ".",
					{X: 5, Y: 4}: "#",
					{X: 0, Y: 5}: "#",
					{X: 1, Y: 5}: "#",
					{X: 2, Y: 5}: "#",
					{X: 3, Y: 5}: "#",
					{X: 4, Y: 5}: ".",
					{X: 5, Y: 5}: "#",
				},
				Height: 5,
				Width:  5,
			},
			arg: 5,
			want: &Grid{
				Lights: Lights{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: "#",
					{X: 2, Y: 0}: ".",
					{X: 3, Y: 0}: "#",
					{X: 4, Y: 0}: "#",
					{X: 5, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: "#",
					{X: 2, Y: 1}: "#",
					{X: 3, Y: 1}: ".",
					{X: 4, Y: 1}: ".",
					{X: 5, Y: 1}: "#",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: "#",
					{X: 2, Y: 2}: "#",
					{X: 3, Y: 2}: ".",
					{X: 4, Y: 2}: ".",
					{X: 5, Y: 2}: ".",
					{X: 0, Y: 3}: ".",
					{X: 1, Y: 3}: "#",
					{X: 2, Y: 3}: "#",
					{X: 3, Y: 3}: ".",
					{X: 4, Y: 3}: ".",
					{X: 5, Y: 3}: ".",
					{X: 0, Y: 4}: "#",
					{X: 1, Y: 4}: ".",
					{X: 2, Y: 4}: "#",
					{X: 3, Y: 4}: ".",
					{X: 4, Y: 4}: ".",
					{X: 5, Y: 4}: ".",
					{X: 0, Y: 5}: "#",
					{X: 1, Y: 5}: "#",
					{X: 2, Y: 5}: ".",
					{X: 3, Y: 5}: ".",
					{X: 4, Y: 5}: ".",
					{X: 5, Y: 5}: "#",
				},
				Height: 5,
				Width:  5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				Lights: tt.fields.Lights,
				Height: tt.fields.Height,
				Width:  tt.fields.Width,
			}
			g.RunStepsPart2(tt.arg)
			if !reflect.DeepEqual(g, tt.want) {
				t.Errorf("Grid.RunStepsPart2() = %v, want %v", g, tt.want)
			}
		})
	}
}

func Test_runAndCountLightsPart1(t *testing.T) {
	type args struct {
		input []string
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "runs steps and counts lights for part 1, advent of code example",
			args: args{
				input: []string{
					".#.#.#",
					"...##.",
					"#....#",
					"..#...",
					"#.#..#",
					"####..",
				},
				steps: 4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runAndCountLightsPart1(tt.args.input, tt.args.steps); got != tt.want {
				t.Errorf("runAndCountLightsPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_runAndCountLightsPart2(t *testing.T) {
	type args struct {
		input []string
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "runs steps and counts lights for part 2, advent of code example",
			args: args{
				input: []string{
					".#.#.#",
					"...##.",
					"#....#",
					"..#...",
					"#.#..#",
					"####..",
				},
				steps: 5,
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runAndCountLightsPart2(tt.args.input, tt.args.steps); got != tt.want {
				t.Errorf("runAndCountLightsPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
