package tile_test

import (
	"Advent-of-Code/2020/Day20/tile"
	"Advent-of-Code/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTile_RotateTile90(t *testing.T) {
	type fields struct {
		Pixels map[graph.Co]string
		Height int
		Width  int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[graph.Co]string
	}{
		{
			name: "test1",
			fields: fields{
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
				},
				Width:  1,
				Height: 1,
			},
			want: map[graph.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
			},
		},
		{
			name: "test2",
			fields: fields{
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: "#",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: ".",
				},
				Width:  2,
				Height: 2,
			},
			want: map[graph.Co]string{
				{X: 0, Y: 0}: ".",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &tile.Tile{
				Pixels: tt.fields.Pixels,
				Height: tt.fields.Height,
				Width:  tt.fields.Width,
			}
			tr.RotateTile90()
			assert.Equal(t, tt.want, tr.Pixels)
		})
	}
}

func TestTile_FlipTile(t *testing.T) {
	type fields struct {
		Pixels map[graph.Co]string
		Width  int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[graph.Co]string
	}{
		{
			name: "test1",
			fields: fields{
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
				},
				Width: 1,
			},
			want: map[graph.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
			},
		},
		{
			name: "test2",
			fields: fields{
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: "#",
					{X: 0, Y: 2}: "#",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: ".",
				},
				Width: 2,
			},
			want: map[graph.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: "#",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &tile.Tile{
				Pixels: tt.fields.Pixels,
				Width:  tt.fields.Width,
			}
			tr.FlipTile()
			assert.Equal(t, tt.want, tr.Pixels)
		})
	}
}

func TestTile_IsAdjacentTop(t *testing.T) {
	tests := []struct {
		name string
		t    tile.Tile
		tile tile.Tile
		want bool
	}{
		{
			name: "returns true if the bottom row (y = t.Height) of t matches the top row (y = 0) of tile",
			t: tile.Tile{
				Height: 1,
				Width:  4,
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: ".",
					{X: 1, Y: 0}: "#",
					{X: 2, Y: 0}: "#",
					{X: 3, Y: 0}: ".",
					{X: 4, Y: 0}: "#",
					{X: 0, Y: 1}: "#",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: ".",
					{X: 3, Y: 1}: "#",
					{X: 4, Y: 1}: ".",
				},
			},
			tile: tile.Tile{
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: ".",
					{X: 3, Y: 0}: "#",
					{X: 4, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: "#",
					{X: 3, Y: 1}: ".",
					{X: 4, Y: 1}: "#",
				},
			},
			want: true,
		},
		{
			name: "returns false if the bottom row (y = t.Height) of t doesn't match the top row (y = 0) of tile",
			t: tile.Tile{
				Height: 1,
				Width:  4,
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: ".",
					{X: 3, Y: 0}: "#",
					{X: 4, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: "#",
					{X: 3, Y: 1}: "#",
					{X: 4, Y: 1}: "#",
				},
			},
			tile: tile.Tile{
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: ".",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 3, Y: 0}: ".",
					{X: 4, Y: 0}: "#",
					{X: 0, Y: 1}: "#",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: ".",
					{X: 3, Y: 1}: "#",
					{X: 4, Y: 1}: ".",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.t.IsAdjacentTop(tt.tile)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTile_IsAdjacentBottom(t *testing.T) {
	tests := []struct {
		name string
		t    tile.Tile
		tile tile.Tile
		want bool
	}{
		{
			name: "returns true if the top row (y = 0) of t matches the bottom row (y = tile.Height) of tile",
			t: tile.Tile{
				Width: 4,
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: ".",
					{X: 3, Y: 0}: "#",
					{X: 4, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: "#",
					{X: 3, Y: 1}: ".",
					{X: 4, Y: 1}: "#",
				},
			},
			tile: tile.Tile{
				Height: 1,
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: ".",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 3, Y: 0}: ".",
					{X: 4, Y: 0}: "#",
					{X: 0, Y: 1}: "#",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: ".",
					{X: 3, Y: 1}: "#",
					{X: 4, Y: 1}: ".",
				},
			},
			want: true,
		},
		{
			name: "returns false if the top row (y = 0) of t doesn't match the bottom row (y = tile.Height) of tile",
			t: tile.Tile{
				Height: 1,
				Width:  4,
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: ".",
					{X: 1, Y: 0}: "#",
					{X: 2, Y: 0}: "#",
					{X: 3, Y: 0}: ".",
					{X: 4, Y: 0}: "#",
					{X: 0, Y: 1}: "#",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: ".",
					{X: 3, Y: 1}: "#",
					{X: 4, Y: 1}: ".",
				},
			},
			tile: tile.Tile{
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: ".",
					{X: 3, Y: 0}: "#",
					{X: 4, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: "#",
					{X: 3, Y: 1}: ".",
					{X: 4, Y: 1}: "#",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.t.IsAdjacentBottom(tt.tile)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTile_IsAdjacentLeft(t *testing.T) {
	tests := []struct {
		name string
		t    tile.Tile
		tile tile.Tile
		want bool
	}{
		{
			name: "returns true if the right column (x = t.Width) of t matches the left column (x = 0) of tile",
			t: tile.Tile{
				Height: 4,
				Width:  1,
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 0, Y: 2}: ".",
					{X: 0, Y: 3}: "#",
					{X: 0, Y: 4}: ".",
					{X: 1, Y: 0}: "#",
					{X: 1, Y: 1}: ".",
					{X: 1, Y: 2}: "#",
					{X: 1, Y: 3}: "#",
					{X: 1, Y: 4}: ".",
				},
			},
			tile: tile.Tile{
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 0, Y: 2}: "#",
					{X: 0, Y: 3}: "#",
					{X: 0, Y: 4}: ".",
					{X: 1, Y: 0}: "#",
					{X: 1, Y: 1}: ".",
					{X: 1, Y: 2}: ".",
					{X: 1, Y: 3}: "#",
					{X: 1, Y: 4}: "#",
				},
			},
			want: true,
		},
		{
			name: "returns false if the right column (x = tile.Width) of t doesn't match the left column (x = 0) of tile",
			t: tile.Tile{
				Height: 4,
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 0, Y: 2}: "#",
					{X: 0, Y: 3}: "#",
					{X: 0, Y: 4}: "#",
					{X: 1, Y: 0}: "#",
					{X: 1, Y: 1}: "#",
					{X: 1, Y: 2}: ".",
					{X: 1, Y: 3}: "#",
					{X: 1, Y: 4}: ".",
				},
			},
			tile: tile.Tile{
				Width: 1,
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 0, Y: 2}: ".",
					{X: 0, Y: 3}: "#",
					{X: 0, Y: 4}: ".",
					{X: 1, Y: 0}: ".",
					{X: 1, Y: 1}: ".",
					{X: 1, Y: 2}: "#",
					{X: 1, Y: 3}: "#",
					{X: 1, Y: 4}: "#",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.t.IsAdjacentLeft(tt.tile)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTile_IsAdjacentRight(t *testing.T) {
	tests := []struct {
		name string
		t    tile.Tile
		tile tile.Tile
		want bool
	}{
		{
			name: "returns true if the left column (x = 0) of t matches the right column (x = tile.Width) of tile",
			t: tile.Tile{
				Height: 4,
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 0, Y: 2}: "#",
					{X: 0, Y: 3}: "#",
					{X: 0, Y: 4}: "#",
					{X: 1, Y: 0}: "#",
					{X: 1, Y: 1}: "#",
					{X: 1, Y: 2}: ".",
					{X: 1, Y: 3}: "#",
					{X: 1, Y: 4}: ".",
				},
			},
			tile: tile.Tile{
				Width: 1,
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 0, Y: 2}: ".",
					{X: 0, Y: 3}: "#",
					{X: 0, Y: 4}: ".",
					{X: 1, Y: 0}: ".",
					{X: 1, Y: 1}: ".",
					{X: 1, Y: 2}: "#",
					{X: 1, Y: 3}: "#",
					{X: 1, Y: 4}: "#",
				},
			},
			want: true,
		},
		{
			name: "returns false if the left column (x = 0) of t doesn't match the right column (x = tile.Width) of tile",
			t: tile.Tile{
				Height: 4,
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 0, Y: 2}: ".",
					{X: 0, Y: 3}: "#",
					{X: 0, Y: 4}: ".",
					{X: 1, Y: 0}: "#",
					{X: 1, Y: 1}: ".",
					{X: 1, Y: 2}: "#",
					{X: 1, Y: 3}: "#",
					{X: 1, Y: 4}: ".",
				},
			},
			tile: tile.Tile{
				Width: 1,
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 0, Y: 2}: "#",
					{X: 0, Y: 3}: "#",
					{X: 0, Y: 4}: ".",
					{X: 1, Y: 0}: "#",
					{X: 1, Y: 1}: ".",
					{X: 1, Y: 2}: ".",
					{X: 1, Y: 3}: "#",
					{X: 1, Y: 4}: "#",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.t.IsAdjacentRight(tt.tile)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTile_IsAdjacentTo(t *testing.T) {
	tests := []struct {
		name string
		t    tile.Tile
		tile tile.Tile
		want bool
	}{
		{
			name: "returns true if t.adjacentTiles.Top is equal to tile.ID",
			t: tile.Tile{
				AdjacentTiles: tile.AdjacentTiles{
					Top:    "tile-top",
					Bottom: "tile-bottom",
					Left:   "tile-left",
					Right:  "tile-right",
				},
			},
			tile: tile.Tile{
				ID: "tile-top",
			},
			want: true,
		},
		{
			name: "returns true if t.adjacentTiles.Bottom is equal to tile.ID",
			t: tile.Tile{
				AdjacentTiles: tile.AdjacentTiles{
					Top:    "tile-top",
					Bottom: "tile-bottom",
					Left:   "tile-left",
					Right:  "tile-right",
				},
			},
			tile: tile.Tile{
				ID: "tile-bottom",
			},
			want: true,
		},
		{
			name: "returns true if t.adjacentTiles.Left is equal to tile.ID",
			t: tile.Tile{
				AdjacentTiles: tile.AdjacentTiles{
					Top:    "tile-top",
					Bottom: "tile-bottom",
					Left:   "tile-left",
					Right:  "tile-right",
				},
			},
			tile: tile.Tile{
				ID: "tile-left",
			},
			want: true,
		},
		{
			name: "returns true if t.adjacentTiles.Right is equal to tile.ID",
			t: tile.Tile{
				AdjacentTiles: tile.AdjacentTiles{
					Top:    "tile-top",
					Bottom: "tile-bottom",
					Left:   "tile-left",
					Right:  "tile-right",
				},
			},
			tile: tile.Tile{
				ID: "tile-right",
			},
			want: true,
		},
		{
			name: "returns false if none of t.AdjacentTiles equals tile.ID",
			t: tile.Tile{
				AdjacentTiles: tile.AdjacentTiles{
					Top:    "tile-top",
					Bottom: "tile-bottom",
					Left:   "tile-left",
					Right:  "tile-right",
				},
			},
			tile: tile.Tile{
				ID: "tile-another",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.t.IsAdjacentTo(tt.tile)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTile_NumAdjacent(t *testing.T) {
	tests := []struct {
		name string
		t    tile.Tile
		want int
	}{
		{
			name: "returns 0 if tile t has no adjacent tiles",
			t:    tile.Tile{},
			want: 0,
		},
		{
			name: "returns 1 if tile t has one adjacent tile (top)",
			t: tile.Tile{
				AdjacentTiles: tile.AdjacentTiles{
					Top: "tile",
				},
			},
			want: 1,
		},
		{
			name: "returns 1 if tile t has one adjacent tile (bottom)",
			t: tile.Tile{
				AdjacentTiles: tile.AdjacentTiles{
					Bottom: "tile",
				},
			},
			want: 1,
		},
		{
			name: "returns 1 if tile t has one adjacent tile (left)",
			t: tile.Tile{
				AdjacentTiles: tile.AdjacentTiles{
					Left: "tile",
				},
			},
			want: 1,
		},
		{
			name: "returns 1 if tile t has one adjacent tile (right)",
			t: tile.Tile{
				AdjacentTiles: tile.AdjacentTiles{
					Right: "tile",
				},
			},
			want: 1,
		},
		{
			name: "returns 2 if tile t has two adjacent tiles",
			t: tile.Tile{
				AdjacentTiles: tile.AdjacentTiles{
					Top:  "til1",
					Left: "tile2",
				},
			},
			want: 2,
		},
		{
			name: "returns 4 if tile t has all adjacent tiles",
			t: tile.Tile{
				AdjacentTiles: tile.AdjacentTiles{
					Top:    "tile1",
					Bottom: "tile2",
					Left:   "tile3",
					Right:  "tile4",
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.t.NumAdjacent()
			assert.Equal(t, tt.want, got)
		})
	}
}
