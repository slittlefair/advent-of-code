package main

import (
	"Advent-of-Code/graph"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInput(t *testing.T) {
	t.Run("simple parse of input", func(t *testing.T) {
		input := []string{
			"esenee",
			"esew",
			"nwwswee",
		}
		want := [][]string{
			{"e", "se", "ne", "e"},
			{"e", "se", "w"},
			{"nw", "w", "sw", "e", "e"},
		}
		got := parseInput(input)
		assert.Equal(t, want, got)
	})
}

var simpleMap = map[graph.Co]bool{
	{X: 0, Y: 0}: true,
	{X: 1, Y: 0}: false,
	{X: 2, Y: 0}: true,
	{X: 3, Y: 0}: false,
	{X: 0, Y: 1}: false,
	{X: 1, Y: 1}: true,
	{X: 2, Y: 1}: false,
	{X: 3, Y: 1}: true,
	{X: 0, Y: 2}: true,
	{X: 1, Y: 2}: false,
	{X: 2, Y: 2}: true,
	{X: 3, Y: 2}: false,
}

func TestTiles_getETile(t *testing.T) {
	tests := []struct {
		name  string
		co    graph.Co
		want  graph.Co
		want1 bool
	}{
		{
			name:  "returns co and value of east tile of an even row co",
			co:    graph.Co{X: 2, Y: 2},
			want:  graph.Co{X: 3, Y: 2},
			want1: false,
		},
		{
			name:  "returns co and value of east tile of an odd row co",
			co:    graph.Co{X: 2, Y: 1},
			want:  graph.Co{X: 3, Y: 1},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tiles{
				Map: simpleMap,
			}
			got, got1 := tr.getETile(tt.co)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestTiles_getSETile(t *testing.T) {
	tests := []struct {
		name  string
		co    graph.Co
		want  graph.Co
		want1 bool
	}{
		{
			name:  "returns co and value of south east tile of an even row co",
			co:    graph.Co{X: 0, Y: 0},
			want:  graph.Co{X: 0, Y: 1},
			want1: false,
		},
		{
			name:  "returns co and value of south east tile of an odd row co",
			co:    graph.Co{X: 0, Y: 1},
			want:  graph.Co{X: 1, Y: 2},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tiles{
				Map: simpleMap,
			}
			got, got1 := tr.getSETile(tt.co)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestTiles_getNETile(t *testing.T) {
	tests := []struct {
		name  string
		co    graph.Co
		want  graph.Co
		want1 bool
	}{
		{
			name:  "returns co and value of north east tile of an even row co",
			co:    graph.Co{X: 0, Y: 2},
			want:  graph.Co{X: 0, Y: 1},
			want1: false,
		},
		{
			name:  "returns co and value of north east tile of an odd row co",
			co:    graph.Co{X: 0, Y: 1},
			want:  graph.Co{X: 1, Y: 0},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tiles{
				Map: simpleMap,
			}
			got, got1 := tr.getNETile(tt.co)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestTiles_getWTile(t *testing.T) {
	tests := []struct {
		name  string
		co    graph.Co
		want  graph.Co
		want1 bool
	}{
		{
			name:  "returns co and value of west tile of an even row co",
			co:    graph.Co{X: 3, Y: 0},
			want:  graph.Co{X: 2, Y: 0},
			want1: true,
		},
		{
			name:  "returns co and value of west tile of an odd row co",
			co:    graph.Co{X: 1, Y: 1},
			want:  graph.Co{X: 0, Y: 1},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tiles{
				Map: simpleMap,
			}
			got, got1 := tr.getWTile(tt.co)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestTiles_getSWTile(t *testing.T) {
	tests := []struct {
		name  string
		co    graph.Co
		want  graph.Co
		want1 bool
	}{
		{
			name:  "returns co and value of south west tile of an even row co",
			co:    graph.Co{X: 2, Y: 0},
			want:  graph.Co{X: 1, Y: 1},
			want1: true,
		},
		{
			name:  "returns co and value of south west tile of an odd row co",
			co:    graph.Co{X: 1, Y: 1},
			want:  graph.Co{X: 1, Y: 2},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tiles{
				Map: simpleMap,
			}
			got, got1 := tr.getSWTile(tt.co)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestTiles_getNWTile(t *testing.T) {
	tests := []struct {
		name  string
		co    graph.Co
		want  graph.Co
		want1 bool
	}{
		{
			name:  "returns co and value of north west tile of an even row co",
			co:    graph.Co{X: 2, Y: 2},
			want:  graph.Co{X: 1, Y: 1},
			want1: true,
		},
		{
			name:  "returns co and value of north west tile of an odd row co",
			co:    graph.Co{X: 1, Y: 1},
			want:  graph.Co{X: 1, Y: 0},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tiles{
				Map: simpleMap,
			}
			got, got1 := tr.getNWTile(tt.co)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestTiles_moveTile(t *testing.T) {
	type fields struct {
		Map         map[graph.Co]bool
		CurrentTile graph.Co
		MaxX        int
		MaxY        int
		MinX        int
		MinY        int
	}
	tests := []struct {
		name   string
		fields fields
		dir    string
		want   *Tiles
	}{
		{
			name: "changes current tile to east tile",
			fields: fields{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 1, Y: 1},
				MaxX:        3,
				MaxY:        2,
				MinX:        0,
				MinY:        0,
			},
			dir: "e",
			want: &Tiles{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 2, Y: 1},
				MaxX:        3,
				MaxY:        2,
				MinX:        0,
				MinY:        0,
			},
		},
		{
			name: "changes current tile to south east tile",
			fields: fields{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 1, Y: 1},
				MaxX:        3,
				MaxY:        2,
				MinX:        0,
				MinY:        0,
			},
			dir: "se",
			want: &Tiles{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 2, Y: 2},
				MaxX:        3,
				MaxY:        2,
				MinX:        0,
				MinY:        0,
			},
		},
		{
			name: "changes current tile to north east tile",
			fields: fields{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 1, Y: 1},
				MaxX:        3,
				MaxY:        2,
				MinX:        0,
				MinY:        0,
			},
			dir: "ne",
			want: &Tiles{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 2, Y: 0},
				MaxX:        3,
				MaxY:        2,
				MinX:        0,
				MinY:        0,
			},
		},
		{
			name: "changes current tile to west tile",
			fields: fields{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 1, Y: 1},
				MaxX:        3,
				MaxY:        2,
				MinX:        0,
				MinY:        0,
			},
			dir: "w",
			want: &Tiles{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 0, Y: 1},
				MaxX:        3,
				MaxY:        2,
				MinX:        0,
				MinY:        0,
			},
		},
		{
			name: "changes current tile to south west tile",
			fields: fields{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 1, Y: 1},
				MaxX:        3,
				MaxY:        2,
				MinX:        0,
				MinY:        0,
			},
			dir: "sw",
			want: &Tiles{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 1, Y: 2},
				MaxX:        3,
				MaxY:        2,
				MinX:        0,
				MinY:        0,
			},
		},
		{
			name: "changes current tile to west tile",
			fields: fields{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 1, Y: 1},
				MaxX:        3,
				MaxY:        2,
				MinX:        0,
				MinY:        0,
			},
			dir: "nw",
			want: &Tiles{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 1, Y: 0},
				MaxX:        3,
				MaxY:        2,
				MinX:        0,
				MinY:        0,
			},
		},
		{
			name: "descreases minX if currentTile.X moves below it",
			fields: fields{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 0, Y: 0},
				MaxX:        3,
				MaxY:        2,
				MinX:        0,
				MinY:        0,
			},
			dir: "w",
			want: &Tiles{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: -1, Y: 0},
				MaxX:        3,
				MaxY:        2,
				MinX:        -1,
				MinY:        0,
			},
		},
		{
			name: "descreases minY if currentTile.Y moves below it",
			fields: fields{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 0, Y: 0},
				MaxX:        3,
				MaxY:        2,
				MinX:        0,
				MinY:        0,
			},
			dir: "ne",
			want: &Tiles{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 0, Y: -1},
				MaxX:        3,
				MaxY:        2,
				MinX:        0,
				MinY:        -1,
			},
		},
		{
			name: "increases maxX if currentTile.X moves above it",
			fields: fields{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 3, Y: 0},
				MaxX:        3,
				MaxY:        2,
				MinX:        0,
				MinY:        0,
			},
			dir: "e",
			want: &Tiles{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 4, Y: 0},
				MaxX:        4,
				MaxY:        2,
				MinX:        0,
				MinY:        0,
			},
		},
		{
			name: "increases maxY if currentTile.Y moves above it",
			fields: fields{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 0, Y: 2},
				MaxX:        3,
				MaxY:        2,
				MinX:        0,
				MinY:        0,
			},
			dir: "se",
			want: &Tiles{
				Map:         simpleMap,
				CurrentTile: graph.Co{X: 0, Y: 3},
				MaxX:        3,
				MaxY:        3,
				MinX:        0,
				MinY:        0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tiles{
				Map:         tt.fields.Map,
				CurrentTile: tt.fields.CurrentTile,
				MaxX:        tt.fields.MaxX,
				MaxY:        tt.fields.MaxY,
				MinX:        tt.fields.MinX,
				MinY:        tt.fields.MinY,
			}
			tr.moveTile(tt.dir)
			assert.Equal(t, tt.want, tr)
		})
	}
}

func TestTiles_flipTiles(t *testing.T) {
	tests := []struct {
		name string
		co   graph.Co
		Map  map[graph.Co]bool
		want *Tiles
	}{
		{
			name: "a tile not in the map is added as flipped to black",
			co:   graph.Co{X: 5, Y: 5},
			Map: map[graph.Co]bool{
				{X: 0, Y: 0}: true,
				{X: 1, Y: 0}: false,
				{X: 2, Y: 0}: true,
				{X: 3, Y: 0}: false,
				{X: 0, Y: 1}: false,
				{X: 1, Y: 1}: true,
				{X: 2, Y: 1}: false,
				{X: 3, Y: 1}: true,
				{X: 0, Y: 2}: true,
				{X: 1, Y: 2}: false,
				{X: 2, Y: 2}: true,
				{X: 3, Y: 2}: false,
			},
			want: &Tiles{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}: true,
					{X: 1, Y: 0}: false,
					{X: 2, Y: 0}: true,
					{X: 3, Y: 0}: false,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 1}: true,
					{X: 2, Y: 1}: false,
					{X: 3, Y: 1}: true,
					{X: 0, Y: 2}: true,
					{X: 1, Y: 2}: false,
					{X: 2, Y: 2}: true,
					{X: 3, Y: 2}: false,
					{X: 5, Y: 5}: true,
				},
			},
		},
		{
			name: "a white tile gets flipped to black",
			co:   graph.Co{X: 0, Y: 1},
			Map: map[graph.Co]bool{
				{X: 0, Y: 0}: true,
				{X: 1, Y: 0}: false,
				{X: 2, Y: 0}: true,
				{X: 3, Y: 0}: false,
				{X: 0, Y: 1}: false,
				{X: 1, Y: 1}: true,
				{X: 2, Y: 1}: false,
				{X: 3, Y: 1}: true,
				{X: 0, Y: 2}: true,
				{X: 1, Y: 2}: false,
				{X: 2, Y: 2}: true,
				{X: 3, Y: 2}: false,
			},
			want: &Tiles{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}: true,
					{X: 1, Y: 0}: false,
					{X: 2, Y: 0}: true,
					{X: 3, Y: 0}: false,
					{X: 0, Y: 1}: true,
					{X: 1, Y: 1}: true,
					{X: 2, Y: 1}: false,
					{X: 3, Y: 1}: true,
					{X: 0, Y: 2}: true,
					{X: 1, Y: 2}: false,
					{X: 2, Y: 2}: true,
					{X: 3, Y: 2}: false,
				},
			},
		},
		{
			name: "a black tile gets flipped to white",
			co:   graph.Co{X: 3, Y: 2},
			Map: map[graph.Co]bool{
				{X: 0, Y: 0}: true,
				{X: 1, Y: 0}: false,
				{X: 2, Y: 0}: true,
				{X: 3, Y: 0}: false,
				{X: 0, Y: 1}: false,
				{X: 1, Y: 1}: true,
				{X: 2, Y: 1}: false,
				{X: 3, Y: 1}: true,
				{X: 0, Y: 2}: true,
				{X: 1, Y: 2}: false,
				{X: 2, Y: 2}: true,
				{X: 3, Y: 2}: false,
			},
			want: &Tiles{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}: true,
					{X: 1, Y: 0}: false,
					{X: 2, Y: 0}: true,
					{X: 3, Y: 0}: false,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 1}: true,
					{X: 2, Y: 1}: false,
					{X: 3, Y: 1}: true,
					{X: 0, Y: 2}: true,
					{X: 1, Y: 2}: false,
					{X: 2, Y: 2}: true,
					{X: 3, Y: 2}: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tiles{
				Map: tt.Map,
			}
			tr.flipTiles(tt.co)
			assert.Equal(t, tt.want, tr)
		})
	}
}

func TestTiles_moveThroughList(t *testing.T) {
	type fields struct {
		Map  map[graph.Co]bool
		MaxX int
		MaxY int
	}
	tests := []struct {
		name   string
		fields fields
		tiles  []string
		want   *Tiles
	}{
		{
			name: "advent of code example 1",
			fields: fields{
				Map: make(map[graph.Co]bool),
			},
			tiles: []string{"e", "se", "ne", "e"},
			want: &Tiles{
				Map: map[graph.Co]bool{
					{X: 3, Y: 0}: true,
				},
				CurrentTile: graph.Co{X: 3, Y: 0},
				MaxX:        3,
				MaxY:        1,
			},
		},
		{
			name: "advent of code example 2",
			fields: fields{
				Map: map[graph.Co]bool{
					{X: 3, Y: 0}: true,
				},
				MaxX: 3,
			},
			tiles: []string{"e", "se", "w"},
			want: &Tiles{
				Map: map[graph.Co]bool{
					{X: 3, Y: 0}: true,
					{X: 0, Y: 1}: true,
				},
				CurrentTile: graph.Co{X: 0, Y: 1},
				MaxX:        3,
				MaxY:        1,
			},
		},
		{
			name: "advent of code example 3",
			fields: fields{
				Map: map[graph.Co]bool{
					{X: 3, Y: 0}: true,
					{X: 0, Y: 1}: true,
				},
				MaxX: 3,
				MaxY: 1,
			},
			tiles: []string{"nw", "w", "sw", "e", "e"},
			want: &Tiles{
				Map: map[graph.Co]bool{
					{X: 3, Y: 0}: true,
					{X: 0, Y: 1}: true,
					{X: 0, Y: 0}: true,
				},
				CurrentTile: graph.Co{X: 0, Y: 0},
				MaxX:        3,
				MaxY:        1,
				MinY:        -1,
				MinX:        -2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tiles{
				Map:  tt.fields.Map,
				MaxX: tt.fields.MaxX,
				MaxY: tt.fields.MaxY,
			}
			tr.moveThroughList(tt.tiles)
			assert.Equal(t, tt.want, tr)
		})
	}
}

func TestTiles_countTiles(t *testing.T) {
	tests := []struct {
		name string
		Map  map[graph.Co]bool
		want int
	}{
		{
			name: "returns 0 if there are no tiles in the map",
			Map:  map[graph.Co]bool{},
			want: 0,
		},
		{
			name: "returns 0 if there are no black tiles in the map",
			Map: map[graph.Co]bool{
				{X: 0, Y: 0}: false,
				{X: 1, Y: 0}: false,
				{X: 2, Y: 0}: false,
				{X: 3, Y: 0}: false,
			},
			want: 0,
		},
		{
			name: "returns the number of tiles if all tiles in the map are black",
			Map: map[graph.Co]bool{
				{X: 0, Y: 0}: true,
				{X: 1, Y: 0}: true,
				{X: 2, Y: 0}: true,
				{X: 3, Y: 0}: true,
			},
			want: 4,
		},
		{
			name: "returns the number of black tiles if some of the tiles in the map are black",
			Map: map[graph.Co]bool{
				{X: 0, Y: 0}: true,
				{X: 1, Y: 0}: false,
				{X: 2, Y: 0}: true,
				{X: 3, Y: 0}: true,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Tiles{
				Map: tt.Map,
			}
			got := tr.countTiles()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTiles_populateMissingTiles(t *testing.T) {
	type fields struct {
		Map         map[graph.Co]bool
		CurrentTile graph.Co
		MaxX        int
		MaxY        int
		MinX        int
		MinY        int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Tiles
	}{
		{
			name: "adds surrounding tiles to map",
			fields: fields{
				MaxX: 1,
				MaxY: 1,
				MinX: 0,
				MinY: -1,
				Map: map[graph.Co]bool{
					{X: 0, Y: -1}: true,
					{X: 0, Y: 0}:  true,
					{X: 0, Y: 1}:  false,
					{X: 1, Y: -1}: true,
					{X: 1, Y: 0}:  true,
					{X: 1, Y: 1}:  false,
				},
			},
			want: &Tiles{
				MaxX: 2,
				MaxY: 2,
				MinX: -1,
				MinY: -2,
				Map: map[graph.Co]bool{
					{X: -1, Y: -2}: false,
					{X: -1, Y: -1}: false,
					{X: -1, Y: 0}:  false,
					{X: -1, Y: 1}:  false,
					{X: -1, Y: 2}:  false,
					{X: 0, Y: -2}:  false,
					{X: 0, Y: -1}:  true,
					{X: 0, Y: 0}:   true,
					{X: 0, Y: 1}:   false,
					{X: 0, Y: 2}:   false,
					{X: 1, Y: -2}:  false,
					{X: 1, Y: -1}:  true,
					{X: 1, Y: 0}:   true,
					{X: 1, Y: 1}:   false,
					{X: 1, Y: 2}:   false,
					{X: 2, Y: -2}:  false,
					{X: 2, Y: -1}:  false,
					{X: 2, Y: 0}:   false,
					{X: 2, Y: 1}:   false,
					{X: 2, Y: 2}:   false,
				},
			},
		},
		{
			name: "adds surrounding and missing tiles to map",
			fields: fields{
				MaxX: 1,
				MaxY: 1,
				MinX: 0,
				MinY: -1,
				Map: map[graph.Co]bool{
					{X: 0, Y: -1}: true,
					{X: 1, Y: -1}: true,
					{X: 1, Y: 0}:  true,
					{X: 1, Y: 1}:  false,
				},
			},
			want: &Tiles{
				MaxX: 2,
				MaxY: 2,
				MinX: -1,
				MinY: -2,
				Map: map[graph.Co]bool{
					{X: -1, Y: -2}: false,
					{X: -1, Y: -1}: false,
					{X: -1, Y: 0}:  false,
					{X: -1, Y: 1}:  false,
					{X: -1, Y: 2}:  false,
					{X: 0, Y: -2}:  false,
					{X: 0, Y: -1}:  true,
					{X: 0, Y: 0}:   false,
					{X: 0, Y: 1}:   false,
					{X: 0, Y: 2}:   false,
					{X: 1, Y: -2}:  false,
					{X: 1, Y: -1}:  true,
					{X: 1, Y: 0}:   true,
					{X: 1, Y: 1}:   false,
					{X: 1, Y: 2}:   false,
					{X: 2, Y: -2}:  false,
					{X: 2, Y: -1}:  false,
					{X: 2, Y: 0}:   false,
					{X: 2, Y: 1}:   false,
					{X: 2, Y: 2}:   false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tiles{
				Map:         tt.fields.Map,
				CurrentTile: tt.fields.CurrentTile,
				MaxX:        tt.fields.MaxX,
				MaxY:        tt.fields.MaxY,
				MinX:        tt.fields.MinX,
				MinY:        tt.fields.MinY,
			}
			tr.populateMissingTiles()
			assert.Equal(t, tt.want, tr)
		})
	}
}

func TestTiles_shouldFlip(t *testing.T) {
	type fields struct {
		Map         map[graph.Co]bool
		CurrentTile graph.Co
		MaxX        int
		MaxY        int
		MinX        int
		MinY        int
	}
	tests := []struct {
		name   string
		fields fields
		co     graph.Co
		want   bool
	}{
		{
			name: "returns true if a white tile is adjacent to two black tiles, east and south east",
			fields: fields{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}: false,
					{X: 1, Y: 0}: true,
					{X: 0, Y: 1}: true,
				},
			},
			co:   graph.Co{X: 0, Y: 0},
			want: true,
		},
		{
			name: "returns true if a white tile is adjacent to two black tiles, north east and west",
			fields: fields{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}:  false,
					{X: 0, Y: -1}: true,
					{X: -1, Y: 0}: true,
				},
			},
			co:   graph.Co{X: 0, Y: 0},
			want: true,
		},
		{
			name: "returns true if a white tile is adjacent to two black tiles, south west and north west",
			fields: fields{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}:   false,
					{X: -1, Y: 1}:  true,
					{X: -1, Y: -1}: true,
				},
			},
			co:   graph.Co{X: 0, Y: 0},
			want: true,
		},
		{
			name: "returns false if a white tile is adjacent to no black tiles",
			fields: fields{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}:   false,
					{X: 1, Y: 0}:   false,
					{X: 0, Y: 1}:   false,
					{X: 0, Y: -1}:  false,
					{X: -1, Y: 0}:  false,
					{X: -1, Y: 1}:  false,
					{X: -1, Y: -1}: false,
				},
			},
			co:   graph.Co{X: 0, Y: 0},
			want: false,
		},
		{
			name: "returns false if a white tile is adjacent to one black tile",
			fields: fields{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}:   false,
					{X: 1, Y: 0}:   true,
					{X: 0, Y: 1}:   false,
					{X: 0, Y: -1}:  false,
					{X: -1, Y: 0}:  false,
					{X: -1, Y: 1}:  false,
					{X: -1, Y: -1}: false,
				},
			},
			co:   graph.Co{X: 0, Y: 0},
			want: false,
		},
		{
			name: "returns false if a white tile is adjacent to three black tiles",
			fields: fields{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}:   false,
					{X: 1, Y: 0}:   true,
					{X: 0, Y: 1}:   true,
					{X: 0, Y: -1}:  true,
					{X: -1, Y: 0}:  false,
					{X: -1, Y: 1}:  false,
					{X: -1, Y: -1}: false,
				},
			},
			co:   graph.Co{X: 0, Y: 0},
			want: false,
		},
		{
			name: "returns false if a white tile is adjacent to four black tiles",
			fields: fields{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}:   false,
					{X: 1, Y: 0}:   true,
					{X: 0, Y: 1}:   false,
					{X: 0, Y: -1}:  true,
					{X: -1, Y: 0}:  false,
					{X: -1, Y: 1}:  true,
					{X: -1, Y: -1}: true,
				},
			},
			co:   graph.Co{X: 0, Y: 0},
			want: false,
		},
		{
			name: "returns false if a white tile is adjacent to five black tiles",
			fields: fields{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}:   false,
					{X: 1, Y: 0}:   true,
					{X: 0, Y: 1}:   true,
					{X: 0, Y: -1}:  true,
					{X: -1, Y: 0}:  false,
					{X: -1, Y: 1}:  true,
					{X: -1, Y: -1}: true,
				},
			},
			co:   graph.Co{X: 0, Y: 0},
			want: false,
		},
		{
			name: "returns false if a white tile is adjacent to six black tiles",
			fields: fields{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}:   false,
					{X: 1, Y: 0}:   true,
					{X: 0, Y: 1}:   false,
					{X: 0, Y: -1}:  true,
					{X: -1, Y: 0}:  false,
					{X: -1, Y: 1}:  true,
					{X: -1, Y: -1}: true,
				},
			},
			co:   graph.Co{X: 0, Y: 0},
			want: false,
		},
		{
			name: "returns true if a black tile is adjacent to no black tiles",
			fields: fields{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}:   true,
					{X: 1, Y: 0}:   false,
					{X: 0, Y: 1}:   false,
					{X: 0, Y: -1}:  false,
					{X: -1, Y: 0}:  false,
					{X: -1, Y: 1}:  false,
					{X: -1, Y: -1}: false,
				},
			},
			co:   graph.Co{X: 0, Y: 0},
			want: true,
		},
		{
			name: "returns true if a black tile is adjacent to three black tiles",
			fields: fields{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}:   true,
					{X: 1, Y: 0}:   true,
					{X: 0, Y: 1}:   true,
					{X: 0, Y: -1}:  true,
					{X: -1, Y: 0}:  false,
					{X: -1, Y: 1}:  false,
					{X: -1, Y: -1}: false,
				},
			},
			co:   graph.Co{X: 0, Y: 0},
			want: true,
		},
		{
			name: "returns true if a black tile is adjacent to four black tiles",
			fields: fields{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}:   true,
					{X: 1, Y: 0}:   true,
					{X: 0, Y: 1}:   true,
					{X: 0, Y: -1}:  true,
					{X: -1, Y: 0}:  true,
					{X: -1, Y: 1}:  false,
					{X: -1, Y: -1}: false,
				},
			},
			co:   graph.Co{X: 0, Y: 0},
			want: true,
		},
		{
			name: "returns true if a black tile is adjacent to five black tiles",
			fields: fields{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}:   true,
					{X: 1, Y: 0}:   true,
					{X: 0, Y: 1}:   true,
					{X: 0, Y: -1}:  true,
					{X: -1, Y: 0}:  true,
					{X: -1, Y: 1}:  true,
					{X: -1, Y: -1}: false,
				},
			},
			co:   graph.Co{X: 0, Y: 0},
			want: true,
		},
		{
			name: "returns true if a black tile is adjacent to six black tiles",
			fields: fields{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}:   true,
					{X: 1, Y: 0}:   true,
					{X: 0, Y: 1}:   true,
					{X: 0, Y: -1}:  true,
					{X: -1, Y: 0}:  true,
					{X: -1, Y: 1}:  true,
					{X: -1, Y: -1}: true,
				},
			},
			co:   graph.Co{X: 0, Y: 0},
			want: true,
		},
		{
			name: "returns false if a black tile is adjacent to one black tile",
			fields: fields{
				Map: map[graph.Co]bool{
					{X: 0, Y: 0}:   true,
					{X: 1, Y: 0}:   false,
					{X: 0, Y: 1}:   true,
					{X: 0, Y: -1}:  false,
					{X: -1, Y: 0}:  false,
					{X: -1, Y: 1}:  false,
					{X: -1, Y: -1}: false,
				},
			},
			co:   graph.Co{X: 0, Y: 0},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tiles{
				Map:         tt.fields.Map,
				CurrentTile: tt.fields.CurrentTile,
				MaxX:        tt.fields.MaxX,
				MaxY:        tt.fields.MaxY,
				MinX:        tt.fields.MinX,
				MinY:        tt.fields.MinY,
			}
			got := tr.shouldFlip(tt.co)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTiles_decideWhichTilesToFlip(t *testing.T) {
	t.Run("returns list of tiles to be flipped, advent of code example", func(t *testing.T) {
		tr := &Tiles{
			Map: map[graph.Co]bool{
				{X: 0, Y: 0}: true,
				{X: 1, Y: 0}: false,
				{X: 2, Y: 0}: false,
				{X: 3, Y: 0}: true,
				{X: 4, Y: 0}: false,
				{X: 0, Y: 1}: false,
				{X: 1, Y: 1}: false,
				{X: 2, Y: 1}: true,
				{X: 3, Y: 1}: false,
				{X: 4, Y: 1}: true,
				{X: 0, Y: 2}: true,
				{X: 1, Y: 2}: true,
				{X: 2, Y: 2}: false,
				{X: 3, Y: 2}: false,
				{X: 4, Y: 2}: true,
				{X: 0, Y: 3}: false,
				{X: 1, Y: 3}: false,
				{X: 2, Y: 3}: false,
				{X: 3, Y: 3}: true,
				{X: 4, Y: 3}: false,
				{X: 0, Y: 4}: false,
				{X: 1, Y: 4}: false,
				{X: 2, Y: 4}: true,
				{X: 3, Y: 4}: true,
				{X: 4, Y: 4}: false,
			},
		}
		want := []graph.Co{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 4, Y: 0},
			{X: 1, Y: 1},
			{X: 2, Y: 2},
			{X: 0, Y: 3},
			{X: 1, Y: 3},
			{X: 4, Y: 3},
			{X: 4, Y: 4},
		}
		got := tr.decideWhichTilesToFlip()
		assert.Len(t, want, len(got))
		assert.ElementsMatch(t, want, got)
	})
}

func TestTiles_doFlips(t *testing.T) {
	t.Run("flips the correct tiles", func(t *testing.T) {
		tr := &Tiles{
			Map: map[graph.Co]bool{
				{X: 0, Y: 0}: true,
				{X: 1, Y: 0}: false,
				{X: 2, Y: 0}: false,
				{X: 3, Y: 0}: true,
				{X: 4, Y: 0}: false,
				{X: 0, Y: 1}: false,
				{X: 1, Y: 1}: false,
				{X: 2, Y: 1}: true,
				{X: 3, Y: 1}: false,
				{X: 4, Y: 1}: true,
				{X: 0, Y: 2}: true,
				{X: 1, Y: 2}: true,
				{X: 2, Y: 2}: false,
				{X: 3, Y: 2}: false,
				{X: 4, Y: 2}: true,
				{X: 0, Y: 3}: false,
				{X: 1, Y: 3}: false,
				{X: 2, Y: 3}: false,
				{X: 3, Y: 3}: true,
				{X: 4, Y: 3}: false,
				{X: 0, Y: 4}: false,
				{X: 1, Y: 4}: false,
				{X: 2, Y: 4}: true,
				{X: 3, Y: 4}: true,
				{X: 4, Y: 4}: false,
			},
		}
		want := &Tiles{
			Map: map[graph.Co]bool{
				{X: 0, Y: 0}: false,
				{X: 1, Y: 0}: false,
				{X: 2, Y: 0}: true,
				{X: 3, Y: 0}: true,
				{X: 4, Y: 0}: true,
				{X: 0, Y: 1}: false,
				{X: 1, Y: 1}: true,
				{X: 2, Y: 1}: true,
				{X: 3, Y: 1}: false,
				{X: 4, Y: 1}: true,
				{X: 0, Y: 2}: true,
				{X: 1, Y: 2}: true,
				{X: 2, Y: 2}: true,
				{X: 3, Y: 2}: false,
				{X: 4, Y: 2}: true,
				{X: 0, Y: 3}: true,
				{X: 1, Y: 3}: true,
				{X: 2, Y: 3}: false,
				{X: 3, Y: 3}: true,
				{X: 4, Y: 3}: true,
				{X: 0, Y: 4}: false,
				{X: 1, Y: 4}: false,
				{X: 2, Y: 4}: true,
				{X: 3, Y: 4}: true,
				{X: 4, Y: 4}: true,
			},
		}
		tr.doFlips()
		assert.Equal(t, want, tr)
	})
}

func TestTiles_countTilesAfterDays(t *testing.T) {
	tests := []struct {
		name string
		days int
		want int
	}{
		{
			days: 1,
			want: 15,
		},
		{
			days: 2,
			want: 12,
		},
		{
			days: 3,
			want: 25,
		},
		{
			days: 4,
			want: 14,
		},
		{
			days: 5,
			want: 23,
		},
		{
			days: 6,
			want: 28,
		},
		{
			days: 7,
			want: 41,
		},
		{
			days: 8,
			want: 37,
		},
		{
			days: 9,
			want: 49,
		},
		{
			days: 10,
			want: 37,
		},
		{
			days: 20,
			want: 132,
		},
		{
			days: 30,
			want: 259,
		},
		{
			days: 40,
			want: 406,
		},
		{
			days: 50,
			want: 566,
		},
		{
			days: 60,
			want: 788,
		},
		{
			days: 70,
			want: 1106,
		},
		{
			days: 80,
			want: 1373,
		},
		{
			days: 90,
			want: 1844,
		},
		{
			days: 100,
			want: 2208,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("advent of code example, %d days", tt.days), func(t *testing.T) {
			tr := &Tiles{
				Map: map[graph.Co]bool{
					{X: -2, Y: -1}: true,
					{X: -2, Y: 0}:  true,
					{X: -2, Y: 1}:  true,
					{X: -2, Y: 2}:  true,
					{X: -2, Y: 3}:  true,
					{X: -1, Y: -3}: false,
					{X: -1, Y: -2}: false,
					{X: -1, Y: -1}: true,
					{X: -1, Y: 0}:  false,
					{X: 0, Y: -2}:  false,
					{X: 0, Y: 0}:   true,
					{X: 1, Y: -3}:  true,
					{X: 1, Y: -2}:  false,
					{X: 1, Y: 2}:   true,
					{X: 2, Y: 0}:   true,
				},
				MaxX: 3,
				MaxY: 4,
				MinX: -3,
				MinY: -6,
			}
			got := tr.countTilesAfterDays(tt.days)
			assert.Equal(t, tt.want, got)
		})
	}
}
