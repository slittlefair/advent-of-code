package tile

import (
	helpers "Advent-of-Code"
	"reflect"
	"testing"
)

func TestTile_RotateTile90(t *testing.T) {
	type fields struct {
		Pixels map[helpers.Co]string
		Height int
		Width  int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[helpers.Co]string
	}{
		{
			name: "test",
			fields: fields{
				Pixels: map[helpers.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
				},
				Width:  1,
				Height: 1,
			},
			want: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
			},
		},
		{
			name: "test",
			fields: fields{
				Pixels: map[helpers.Co]string{
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
			want: map[helpers.Co]string{
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
			tr := &Tile{
				Pixels: tt.fields.Pixels,
				Height: tt.fields.Height,
				Width:  tt.fields.Width,
			}
			tr.RotateTile90()
			if !reflect.DeepEqual(tr.Pixels, tt.want) {
				t.Errorf("got %v, want %v, %d", tr.Pixels, tt.want, tr.Width)
			}
		})
	}
}

func TestTile_FlipTile(t *testing.T) {
	type fields struct {
		Pixels map[helpers.Co]string
		Height int
		Width  int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[helpers.Co]string
	}{
		{
			name: "test",
			fields: fields{
				Pixels: map[helpers.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
				},
				Width: 1,
			},
			want: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
			},
		},
		{
			name: "test",
			fields: fields{
				Pixels: map[helpers.Co]string{
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
			want: map[helpers.Co]string{
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
			tr := &Tile{
				Pixels: tt.fields.Pixels,
				Height: tt.fields.Height,
				Width:  tt.fields.Width,
			}
			tr.FlipTile()
			if !reflect.DeepEqual(tr.Pixels, tt.want) {
				t.Errorf("got %v, want %v, %d", tr.Pixels, tt.want, tr.Width)
			}
		})
	}
}

func TestTile_IsAdjacentTop(t *testing.T) {
	tests := []struct {
		name string
		t    Tile
		tile Tile
		want bool
	}{
		{
			name: "returns true if the bottom row (y = t.Height) of t matches the top row (y = 0) of tile",
			t: Tile{
				Height: 1,
				Width:  4,
				Pixels: map[helpers.Co]string{
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
			tile: Tile{
				Pixels: map[helpers.Co]string{
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
			t: Tile{
				Height: 1,
				Width:  4,
				Pixels: map[helpers.Co]string{
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
			tile: Tile{
				Pixels: map[helpers.Co]string{
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
			if got := tt.t.IsAdjacentTop(tt.tile); got != tt.want {
				t.Errorf("Tile.IsAdjacentTop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTile_IsAdjacentBottom(t *testing.T) {
	tests := []struct {
		name string
		t    Tile
		tile Tile
		want bool
	}{
		{
			name: "returns true if the top row (y = 0) of t matches the bottom row (y = tile.Height) of tile",
			t: Tile{
				Width: 4,
				Pixels: map[helpers.Co]string{
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
			tile: Tile{
				Height: 1,
				Pixels: map[helpers.Co]string{
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
			t: Tile{
				Height: 1,
				Width:  4,
				Pixels: map[helpers.Co]string{
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
			tile: Tile{
				Pixels: map[helpers.Co]string{
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
			if got := tt.t.IsAdjacentBottom(tt.tile); got != tt.want {
				t.Errorf("Tile.IsAdjacentBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTile_IsAdjacentLeft(t *testing.T) {
	tests := []struct {
		name string
		t    Tile
		tile Tile
		want bool
	}{
		{
			name: "returns true if the right column (x = t.Width) of t matches the left column (x = 0) of tile",
			t: Tile{
				Height: 4,
				Width:  1,
				Pixels: map[helpers.Co]string{
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
			tile: Tile{
				Pixels: map[helpers.Co]string{
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
			t: Tile{
				Height: 4,
				Pixels: map[helpers.Co]string{
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
			tile: Tile{
				Width: 1,
				Pixels: map[helpers.Co]string{
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
			if got := tt.t.IsAdjacentLeft(tt.tile); got != tt.want {
				t.Errorf("Tile.IsAdjacentLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTile_IsAdjacentRight(t *testing.T) {
	tests := []struct {
		name string
		t    Tile
		tile Tile
		want bool
	}{
		{
			name: "returns true if the left column (x = 0) of t matches the right column (x = tile.Width) of tile",
			t: Tile{
				Height: 4,
				Pixels: map[helpers.Co]string{
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
			tile: Tile{
				Width: 1,
				Pixels: map[helpers.Co]string{
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
			t: Tile{
				Height: 4,
				Pixels: map[helpers.Co]string{
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
			tile: Tile{
				Width: 1,
				Pixels: map[helpers.Co]string{
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
			if got := tt.t.IsAdjacentRight(tt.tile); got != tt.want {
				t.Errorf("Tile.IsAdjacentRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTile_IsAdjacentTo(t *testing.T) {
	tests := []struct {
		name string
		t    Tile
		tile Tile
		want bool
	}{
		{
			name: "returns true if t.adjacentTiles.Top is equal to tile.ID",
			t: Tile{
				AdjacentTiles: AdjacentTiles{
					Top:    "tile-top",
					Bottom: "tile-bottom",
					Left:   "tile-left",
					Right:  "tile-right",
				},
			},
			tile: Tile{
				ID: "tile-top",
			},
			want: true,
		},
		{
			name: "returns true if t.adjacentTiles.Bottom is equal to tile.ID",
			t: Tile{
				AdjacentTiles: AdjacentTiles{
					Top:    "tile-top",
					Bottom: "tile-bottom",
					Left:   "tile-left",
					Right:  "tile-right",
				},
			},
			tile: Tile{
				ID: "tile-bottom",
			},
			want: true,
		},
		{
			name: "returns true if t.adjacentTiles.Left is equal to tile.ID",
			t: Tile{
				AdjacentTiles: AdjacentTiles{
					Top:    "tile-top",
					Bottom: "tile-bottom",
					Left:   "tile-left",
					Right:  "tile-right",
				},
			},
			tile: Tile{
				ID: "tile-left",
			},
			want: true,
		},
		{
			name: "returns true if t.adjacentTiles.Right is equal to tile.ID",
			t: Tile{
				AdjacentTiles: AdjacentTiles{
					Top:    "tile-top",
					Bottom: "tile-bottom",
					Left:   "tile-left",
					Right:  "tile-right",
				},
			},
			tile: Tile{
				ID: "tile-right",
			},
			want: true,
		},
		{
			name: "returns false if none of t.AdjacentTiles equals tile.ID",
			t: Tile{
				AdjacentTiles: AdjacentTiles{
					Top:    "tile-top",
					Bottom: "tile-bottom",
					Left:   "tile-left",
					Right:  "tile-right",
				},
			},
			tile: Tile{
				ID: "tile-another",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.IsAdjacentTo(tt.tile); got != tt.want {
				t.Errorf("Tile.IsAdjacentTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTile_NumAdjacent(t *testing.T) {
	tests := []struct {
		name string
		t    Tile
		want int
	}{
		{
			name: "returns 0 if tile t has no adjacent tiles",
			t:    Tile{},
			want: 0,
		},
		{
			name: "returns 1 if tile t has one adjacent tile (top)",
			t: Tile{
				AdjacentTiles: AdjacentTiles{
					Top: "tile",
				},
			},
			want: 1,
		},
		{
			name: "returns 1 if tile t has one adjacent tile (bottom)",
			t: Tile{
				AdjacentTiles: AdjacentTiles{
					Bottom: "tile",
				},
			},
			want: 1,
		},
		{
			name: "returns 1 if tile t has one adjacent tile (left)",
			t: Tile{
				AdjacentTiles: AdjacentTiles{
					Left: "tile",
				},
			},
			want: 1,
		},
		{
			name: "returns 1 if tile t has one adjacent tile (right)",
			t: Tile{
				AdjacentTiles: AdjacentTiles{
					Right: "tile",
				},
			},
			want: 1,
		},
		{
			name: "returns 2 if tile t has two adjacent tiles",
			t: Tile{
				AdjacentTiles: AdjacentTiles{
					Top:  "til1",
					Left: "tile2",
				},
			},
			want: 2,
		},
		{
			name: "returns 4 if tile t has all adjacent tiles",
			t: Tile{
				AdjacentTiles: AdjacentTiles{
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
			if got := tt.t.NumAdjacent(); got != tt.want {
				t.Errorf("Tile.NumAdjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}
