package picture

import (
	helpers "Advent-of-Code"
	tile "Advent-of-Code/2020/Day20/tile"
	"reflect"
	"testing"
)

func TestPicture_PopulateTiles(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  Picture
	}{
		{
			name: "populates a simple picture from input",
			input: []string{
				"Tile 7:",
				"..",
				"##",
				"",
				"Tile 5:",
				".#",
				"#.",
				"",
				"Tile 11:",
				".#",
				".#",
				"",
				"Tile 3:",
				"..",
				"..",
			},
			want: Picture{
				Tiles: []tile.Tile{
					{
						ID:     "7",
						Height: 1,
						Width:  1,
						Pixels: map[helpers.Co]string{
							{X: 0, Y: 0}: ".",
							{X: 1, Y: 0}: ".",
							{X: 0, Y: 1}: "#",
							{X: 1, Y: 1}: "#",
						},
					},
					{
						ID:     "5",
						Height: 1,
						Width:  1,
						Pixels: map[helpers.Co]string{
							{X: 0, Y: 0}: ".",
							{X: 1, Y: 0}: "#",
							{X: 0, Y: 1}: "#",
							{X: 1, Y: 1}: ".",
						},
					},
					{
						ID:     "11",
						Height: 1,
						Width:  1,
						Pixels: map[helpers.Co]string{
							{X: 0, Y: 0}: ".",
							{X: 1, Y: 0}: "#",
							{X: 0, Y: 1}: ".",
							{X: 1, Y: 1}: "#",
						},
					},
					{
						ID:     "3",
						Height: 1,
						Width:  1,
						Pixels: map[helpers.Co]string{
							{X: 0, Y: 0}: ".",
							{X: 1, Y: 0}: ".",
							{X: 0, Y: 1}: ".",
							{X: 1, Y: 1}: ".",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Picture{}
			p.PopulateTiles(tt.input)
			if !reflect.DeepEqual(*p, tt.want) {
				t.Errorf("got %v, want %v", p, tt.want)
			}
		})
	}
}

func TestPicture_FindMatchesForTile(t *testing.T) {
	tests := []struct {
		name  string
		p     Picture
		t     tile.Tile
		index int
		want  Picture
	}{
		{
			name: "skips tile in picture if its ID matches the given tile ID",
			p: Picture{
				Tiles: []tile.Tile{
					{
						ID: "tile-1",
					},
				},
			},
			t: tile.Tile{
				ID: "tile-1",
			},
			index: 0,
			want: Picture{
				Tiles: []tile.Tile{
					{
						ID: "tile-1",
					},
				},
			},
		},
		{
			name: "skips tile in picture if it is already known to be adjacent to the given tile ID",
			p: Picture{
				Tiles: []tile.Tile{
					{
						ID: "tile-2",
					},
				},
			},
			t: tile.Tile{
				ID: "tile-1",
				AdjacentTiles: tile.AdjacentTiles{
					Top: "tile-2",
				},
			},
			index: 0,
			want: Picture{
				Tiles: []tile.Tile{
					{
						ID: "tile-2",
					},
				},
			},
		},
		{
			name: "skips tile in picture if it is already adjacent to 4 tiles",
			p: Picture{
				Tiles: []tile.Tile{
					{
						ID: "tile-2",
					},
				},
			},
			t: tile.Tile{
				ID: "tile-1",
				AdjacentTiles: tile.AdjacentTiles{
					Top:    "tile-3",
					Bottom: "tile-4",
					Left:   "tile-5",
					Right:  "tile-6",
				},
			},
			index: 0,
			want: Picture{
				Tiles: []tile.Tile{
					{
						ID: "tile-2",
					},
				},
			},
		},
		{
			name: "skips tile in picture if it is already adjacent to 4 tiles",
			p: Picture{
				Tiles: []tile.Tile{
					{
						ID: "tile-2",
						AdjacentTiles: tile.AdjacentTiles{
							Top:    "tile-3",
							Bottom: "tile-4",
							Left:   "tile-5",
							Right:  "tile-6",
						},
					},
				},
			},
			t: tile.Tile{
				ID: "tile-1",
			},
			index: 0,
			want: Picture{
				Tiles: []tile.Tile{
					{
						ID: "tile-2",
						AdjacentTiles: tile.AdjacentTiles{
							Top:    "tile-3",
							Bottom: "tile-4",
							Left:   "tile-5",
							Right:  "tile-6",
						},
					},
				},
			},
		},
		{
			name: "does not affect the picture if the given tile is not adjacent to any tile in the picture",
			p: Picture{
				Tiles: []tile.Tile{
					{
						ID: "tile-1",
						Pixels: map[helpers.Co]string{
							{X: 0, Y: 0}: ".",
							{X: 1, Y: 0}: "#",
							{X: 0, Y: 1}: "#",
							{X: 1, Y: 1}: ".",
						},
						Height: 1,
						Width:  1,
					},
					{
						ID: "tile-2",
						Pixels: map[helpers.Co]string{
							{X: 0, Y: 0}: ".",
							{X: 1, Y: 0}: ".",
							{X: 0, Y: 1}: ".",
							{X: 1, Y: 1}: ".",
						},
						Height: 1,
						Width:  1,
					},
					{
						ID: "tile-3",
						Pixels: map[helpers.Co]string{
							{X: 0, Y: 0}: "#",
							{X: 1, Y: 0}: "#",
							{X: 0, Y: 1}: "#",
							{X: 1, Y: 1}: "#",
						},
						Height: 1,
						Width:  1,
					},
				},
			},
			t: tile.Tile{
				ID: "tile-1",
				Pixels: map[helpers.Co]string{
					{X: 0, Y: 0}: ".",
					{X: 1, Y: 0}: "#",
					{X: 0, Y: 1}: "#",
					{X: 1, Y: 1}: ".",
				},
				Height: 1,
				Width:  1,
			},
			index: 0,
			want: Picture{
				Tiles: []tile.Tile{
					{
						ID: "tile-1",
						Pixels: map[helpers.Co]string{
							{X: 0, Y: 0}: ".",
							{X: 1, Y: 0}: "#",
							{X: 0, Y: 1}: "#",
							{X: 1, Y: 1}: ".",
						},
						Height: 1,
						Width:  1,
					},
					{
						ID: "tile-2",
						Pixels: map[helpers.Co]string{
							{X: 0, Y: 0}: ".",
							{X: 1, Y: 0}: ".",
							{X: 0, Y: 1}: ".",
							{X: 1, Y: 1}: ".",
						},
						Height: 1,
						Width:  1,
					},
					{
						ID: "tile-3",
						Pixels: map[helpers.Co]string{
							{X: 0, Y: 0}: "#",
							{X: 1, Y: 0}: "#",
							{X: 0, Y: 1}: "#",
							{X: 1, Y: 1}: "#",
						},
						Height: 1,
						Width:  1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.FindMatchesForTile(tt.t, tt.index)
			if !reflect.DeepEqual(tt.p, tt.want) {
				t.Errorf("got %+v, want %+v", tt.p, tt.want)
			}
		})
	}

	tests2 := []struct {
		name          string
		tile1Pixels   map[helpers.Co]string
		tile2Pixels   map[helpers.Co]string
		wantPixels    map[helpers.Co]string
		want1Adjacent tile.AdjacentTiles
		want2Adjacent tile.AdjacentTiles
	}{
		{
			name: "successfully edits adajcent tiles when tile1 is above tile2, rotate = 0, flip = 0",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "-",
				{X: 1, Y: 0}: "-",
				{X: 2, Y: 0}: "-",
				{X: 0, Y: 1}: "-",
				{X: 1, Y: 1}: "-",
				{X: 2, Y: 1}: "-",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Bottom: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Top: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is below tile2, rotate = 0, flip = 0",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: "-",
				{X: 2, Y: 1}: "-",
				{X: 0, Y: 2}: "-",
				{X: 1, Y: 2}: "-",
				{X: 2, Y: 2}: "-",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
			},
			want1Adjacent: tile.AdjacentTiles{
				Top: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Bottom: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is to the left of tile2, rotate = 0, flip = 0",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "-",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: "-",
				{X: 1, Y: 1}: "-",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: "-",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: "#",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Right: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Left: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is to the right of tile2, rotate = 0, flip = 0",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "-",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "-",
				{X: 2, Y: 2}: "-",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: "#",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: "#",
			},
			want1Adjacent: tile.AdjacentTiles{
				Left: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Right: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is above tile2, rotate = 1, flip = 0",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Bottom: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Top: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is below tile2, rotate = 1, flip = 0",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Top: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Bottom: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is to the left of tile2, rotate = 1, flip = 0",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Right: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Left: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is to the right of tile2, rotate = 1, flip = 0",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Left: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Right: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is above tile2, rotate = 2, flip = 0",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Bottom: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Top: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is below tile2, rotate = 2, flip = 0",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Top: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Bottom: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is to the left of tile2, rotate = 2, flip = 0",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Right: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Left: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is to the right of tile2, rotate = 2, flip = 0",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Left: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Right: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is above tile2, rotate = 1, flip = 0",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Bottom: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Top: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is below tile2, rotate = 1, flip = 0",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Top: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Bottom: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is to the left of tile2, rotate = 1, flip = 0",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Right: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Left: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is to the right of tile2, rotate = 1, flip = 0",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Left: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Right: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is above tile2, rotate = 0, flip = 1",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "-",
				{X: 1, Y: 0}: "-",
				{X: 2, Y: 0}: "-",
				{X: 0, Y: 1}: "-",
				{X: 1, Y: 1}: "-",
				{X: 2, Y: 1}: "-",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Bottom: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Top: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is below tile2, rotate = 0, flip = 1",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "-",
				{X: 0, Y: 2}: "-",
				{X: 1, Y: 2}: "-",
				{X: 2, Y: 2}: ".",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
			},
			want1Adjacent: tile.AdjacentTiles{
				Top: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Bottom: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is to the left of tile2, rotate = 0, flip = 0",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "-",
				{X: 1, Y: 0}: "-",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: "-",
				{X: 1, Y: 1}: "-",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: "-",
				{X: 1, Y: 2}: "-",
				{X: 2, Y: 2}: ".",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Right: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Left: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is to the right of tile2, rotate = 0, flip = 0",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "-",
				{X: 2, Y: 0}: "-",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "-",
				{X: 2, Y: 1}: "-",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "-",
				{X: 2, Y: 2}: "-",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Left: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Right: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is above tile2, rotate = 1, flip = 1",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "-",
				{X: 1, Y: 0}: "-",
				{X: 2, Y: 0}: "-",
				{X: 0, Y: 1}: "-",
				{X: 1, Y: 1}: "-",
				{X: 2, Y: 1}: "-",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Bottom: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Top: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is below tile2, rotate = 1, flip = 1",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: "-",
				{X: 1, Y: 1}: "-",
				{X: 2, Y: 1}: "-",
				{X: 0, Y: 2}: "-",
				{X: 1, Y: 2}: "-",
				{X: 2, Y: 2}: "-",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: "#",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
			},
			want1Adjacent: tile.AdjacentTiles{
				Top: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Bottom: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is to the left of tile2, rotate = 1, flip = 1",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "-",
				{X: 1, Y: 0}: "-",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: "-",
				{X: 1, Y: 1}: "-",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: "-",
				{X: 1, Y: 2}: "-",
				{X: 2, Y: 2}: ".",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Right: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Left: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is to the right of tile2, rotate = 1, flip = 1",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "-",
				{X: 2, Y: 0}: "-",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "-",
				{X: 2, Y: 1}: "-",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "-",
				{X: 2, Y: 2}: "-",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Left: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Right: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is above tile2, rotate = 2, flip = 1",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "-",
				{X: 1, Y: 0}: "-",
				{X: 2, Y: 0}: "-",
				{X: 0, Y: 1}: "-",
				{X: 1, Y: 1}: "-",
				{X: 2, Y: 1}: "-",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Bottom: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Top: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is below tile2, rotate = 2, flip = 1",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "-",
				{X: 1, Y: 1}: "-",
				{X: 2, Y: 1}: "-",
				{X: 0, Y: 2}: "-",
				{X: 1, Y: 2}: "-",
				{X: 2, Y: 2}: "-",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Top: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Bottom: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is to the left of tile2, rotate = 2, flip = 1",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "-",
				{X: 1, Y: 0}: "-",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: "-",
				{X: 1, Y: 1}: "-",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: "-",
				{X: 1, Y: 2}: "-",
				{X: 2, Y: 2}: ".",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Right: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Left: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is to the right of tile2, rotate = 2, flip = 1",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "-",
				{X: 2, Y: 0}: "-",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "-",
				{X: 2, Y: 1}: "-",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "-",
				{X: 2, Y: 2}: "-",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Left: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Right: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is above tile2, rotate = 3, flip = 1",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "-",
				{X: 1, Y: 0}: "-",
				{X: 2, Y: 0}: "-",
				{X: 0, Y: 1}: "-",
				{X: 1, Y: 1}: "-",
				{X: 2, Y: 1}: "-",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: "#",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Bottom: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Top: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is below tile2, rotate = 3, flip = 1",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "-",
				{X: 1, Y: 1}: "-",
				{X: 2, Y: 1}: "-",
				{X: 0, Y: 2}: "-",
				{X: 1, Y: 2}: "-",
				{X: 2, Y: 2}: "-",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Top: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Bottom: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is to the left of tile2, rotate = 3, flip = 1",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "-",
				{X: 1, Y: 0}: "-",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: "-",
				{X: 1, Y: 1}: "-",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: "-",
				{X: 1, Y: 2}: "-",
				{X: 2, Y: 2}: ".",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Right: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Left: "tile-1",
			},
		},
		{
			name: "successfully edits adajcent tiles when tile1 is to the right of tile2, rotate = 3, flip = 1",
			tile1Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "-",
				{X: 2, Y: 0}: "-",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "-",
				{X: 2, Y: 1}: "-",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "-",
				{X: 2, Y: 2}: "-",
			},
			tile2Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			wantPixels: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
			},
			want1Adjacent: tile.AdjacentTiles{
				Left: "tile-2",
			},
			want2Adjacent: tile.AdjacentTiles{
				Right: "tile-1",
			},
		},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			tile1 := tile.Tile{
				ID:     "tile-1",
				Pixels: tt.tile1Pixels,
				Height: 2,
				Width:  2,
			}
			tile2 := tile.Tile{
				ID:     "tile-2",
				Pixels: tt.tile2Pixels,
				Height: 2,
				Width:  2,
			}
			p := Picture{
				Tiles: []tile.Tile{tile1, tile2},
			}
			want := Picture{
				Tiles: []tile.Tile{
					{
						ID:            tile1.ID,
						Pixels:        tile1.Pixels,
						Height:        tile1.Height,
						Width:         tile1.Width,
						AdjacentTiles: tt.want1Adjacent,
					},
					{
						ID:            tile2.ID,
						Pixels:        tt.wantPixels,
						Height:        tile2.Height,
						Width:         tile2.Width,
						AdjacentTiles: tt.want2Adjacent,
					},
				},
			}
			p.FindMatchesForTile(tile1, 0)
			if !reflect.DeepEqual(p, want) {
				t.Errorf("Picture.FindMatchesForTile() = %v, want %v", p, want)
			}
		})
	}
}

func TestPicture_CalculateCornerIDs(t *testing.T) {
	tests := []struct {
		name    string
		Tiles   []tile.Tile
		want    int
		wantErr bool
	}{
		{
			name: "returns an error if a corner tile ID can't be converted into an int",
			Tiles: []tile.Tile{
				{
					ID: "7",
					AdjacentTiles: tile.AdjacentTiles{
						Top:   "6",
						Right: "1",
					},
				},
				{
					ID: "10",
					AdjacentTiles: tile.AdjacentTiles{
						Top:  "9",
						Left: "2",
					},
				},
				{
					ID: "77",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "0",
						Left:   "3",
					},
				},
				{
					ID: "string",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "13",
						Right:  "17",
					},
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "returns the product of corner tiles (only corners)",
			Tiles: []tile.Tile{
				{
					ID: "7",
					AdjacentTiles: tile.AdjacentTiles{
						Top:   "6",
						Right: "1",
					},
				},
				{
					ID: "3",
					AdjacentTiles: tile.AdjacentTiles{
						Top:  "9",
						Left: "2",
					},
				},
				{
					ID: "11",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "0",
						Left:   "3",
					},
				},
				{
					ID: "13",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "13",
						Right:  "17",
					},
				},
			},
			want:    3003,
			wantErr: false,
		},
		{
			name: "returns the product of corner tiles",
			Tiles: []tile.Tile{
				{
					ID: "7",
					AdjacentTiles: tile.AdjacentTiles{
						Top: "6",
					},
				},
				{
					ID: "17",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "13",
						Right:  "17",
					},
				},
				{
					ID: "3",
					AdjacentTiles: tile.AdjacentTiles{
						Top:  "9",
						Left: "2",
					},
				},
				{
					ID: "11",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "0",
						Left:   "3",
					},
				},
				{
					ID: "13",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "13",
					},
				},
				{
					ID: "197",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "13",
						Right:  "17",
					},
				},
				{
					ID: "19",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "13",
						Right:  "17",
						Left:   "1001",
						Top:    "97",
					},
				},
			},
			want:    110517,
			wantErr: false,
		},
		{
			name: "returns the product of corner tiles (non corner, non numeric ID)",
			Tiles: []tile.Tile{
				{
					ID: "7",
					AdjacentTiles: tile.AdjacentTiles{
						Top: "6",
					},
				},
				{
					ID: "17",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "13",
						Right:  "17",
					},
				},
				{
					ID: "3",
					AdjacentTiles: tile.AdjacentTiles{
						Top:  "9",
						Left: "2",
					},
				},
				{
					ID: "11",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "0",
						Left:   "3",
					},
				},
				{
					ID: "string",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "13",
					},
				},
				{
					ID: "31",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "13",
						Right:  "17",
					},
				},
				{
					ID: "19",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "13",
						Right:  "17",
						Left:   "1001",
						Top:    "97",
					},
				},
			},
			want:    17391,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Picture{
				Tiles: tt.Tiles,
			}
			got, err := p.CalculateCornerIDs()
			if (err != nil) != tt.wantErr {
				t.Errorf("Picture.CalculateCornerIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Picture.CalculateCornerIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPicture_getTileFromID(t *testing.T) {
	tests := []struct {
		name    string
		Tiles   []tile.Tile
		id      string
		want    tile.Tile
		wantErr bool
	}{
		{
			name: "returns an error if picture contains no tile with the given ID",
			Tiles: []tile.Tile{
				{
					ID: "1",
				},
				{
					ID: "2",
				},
				{
					ID: "3",
				},
				{
					ID: "4",
				},
			},
			id:      "5",
			want:    tile.Tile{},
			wantErr: true,
		},
		{
			name: "returns the correct tile if picture contains a tile with the given ID",
			Tiles: []tile.Tile{
				{
					ID: "1",
				},
				{
					ID: "2",
				},
				{
					ID: "3",
				},
				{
					ID: "4",
				},
			},
			id: "3",
			want: tile.Tile{
				ID: "3",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Picture{
				Tiles: tt.Tiles,
			}
			got, err := p.getTileFromID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Picture.getTileFromID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Picture.getTileFromID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPicture_getTopLeftTile(t *testing.T) {
	tests := []struct {
		name    string
		Tiles   []tile.Tile
		want    tile.Tile
		wantErr bool
	}{
		{
			name: "returns an error if there is no top left tile",
			Tiles: []tile.Tile{
				{
					ID: "top",
					AdjacentTiles: tile.AdjacentTiles{
						Top: "a",
					},
				},
				{
					ID: "top-right",
					AdjacentTiles: tile.AdjacentTiles{
						Top:   "a",
						Right: "b",
					},
				},
				{
					ID: "right",
					AdjacentTiles: tile.AdjacentTiles{
						Right: "a",
					},
				},
				{
					ID: "top-left",
					AdjacentTiles: tile.AdjacentTiles{
						Top:  "a",
						Left: "b",
					},
				},
				{
					ID: "bottom",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "a",
					},
				},
				{
					ID: "bottom-left",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "a",
						Left:   "b",
					},
				},
				{
					ID: "left",
					AdjacentTiles: tile.AdjacentTiles{
						Left: "a",
					},
				},
				{
					ID: "top-left-right",
					AdjacentTiles: tile.AdjacentTiles{
						Top:   "a",
						Left:  "b",
						Right: "c",
					},
				},
				{
					ID: "top-left-bottom",
					AdjacentTiles: tile.AdjacentTiles{
						Top:    "a",
						Left:   "b",
						Bottom: "c",
					},
				},
				{
					ID: "top-right-bottom",
					AdjacentTiles: tile.AdjacentTiles{
						Top:    "a",
						Right:  "b",
						Bottom: "c",
					},
				},
				{
					ID: "bottom-left-right",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "a",
						Left:   "b",
						Right:  "c",
					},
				},
			},
			want:    tile.Tile{},
			wantErr: true,
		},
		{
			name: "returns an error if there is no top left tile",
			Tiles: []tile.Tile{
				{
					ID: "top",
					AdjacentTiles: tile.AdjacentTiles{
						Top: "a",
					},
				},
				{
					ID: "top-right",
					AdjacentTiles: tile.AdjacentTiles{
						Top:   "a",
						Right: "b",
					},
				},
				{
					ID: "right",
					AdjacentTiles: tile.AdjacentTiles{
						Right: "a",
					},
				},
				{
					ID: "top-left",
					AdjacentTiles: tile.AdjacentTiles{
						Top:  "a",
						Left: "b",
					},
				},
				{
					ID: "bottom",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "a",
					},
				},
				{
					ID: "bottom-left",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "a",
						Left:   "b",
					},
				},
				{
					ID: "left",
					AdjacentTiles: tile.AdjacentTiles{
						Left: "a",
					},
				},
				{
					ID: "top-left-right",
					AdjacentTiles: tile.AdjacentTiles{
						Top:   "a",
						Left:  "b",
						Right: "c",
					},
				},
				{
					ID: "top-left-bottom",
					AdjacentTiles: tile.AdjacentTiles{
						Top:    "a",
						Left:   "b",
						Bottom: "c",
					},
				},
				{
					ID: "bottom-right",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "a",
						Right:  "b",
					},
				},
				{
					ID: "top-right-bottom",
					AdjacentTiles: tile.AdjacentTiles{
						Top:    "a",
						Right:  "b",
						Bottom: "c",
					},
				},
				{
					ID: "bottom-left-right",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "a",
						Left:   "b",
						Right:  "c",
					},
				},
			},
			want: tile.Tile{
				ID: "bottom-right",
				AdjacentTiles: tile.AdjacentTiles{
					Bottom: "a",
					Right:  "b",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Picture{
				Tiles: tt.Tiles,
			}
			got, err := p.getTopLeftTile()
			if (err != nil) != tt.wantErr {
				t.Errorf("Picture.getTopLeftTile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Picture.getTopLeftTile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPicture_populatePictureWithTile(t *testing.T) {
	type args struct {
		t tile.Tile
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want Picture
	}{
		{
			name: "correctly populates the given tile into the picture",
			args: args{
				t: tile.Tile{
					ID: "123",
					Pixels: map[helpers.Co]string{
						{X: 0, Y: 0}: ".",
						{X: 1, Y: 0}: "#",
						{X: 2, Y: 0}: "#",
						{X: 3, Y: 0}: ".",
						{X: 0, Y: 1}: ".",
						{X: 1, Y: 1}: "#",
						{X: 2, Y: 1}: ".",
						{X: 3, Y: 1}: ".",
						{X: 0, Y: 2}: ".",
						{X: 1, Y: 2}: ".",
						{X: 2, Y: 2}: "#",
						{X: 3, Y: 2}: "#",
						{X: 0, Y: 3}: "#",
						{X: 1, Y: 3}: "#",
						{X: 2, Y: 3}: "#",
						{X: 3, Y: 3}: ".",
					},
					Height: 3,
					Width:  3,
				},
				x: 2,
				y: 2,
			},
			want: Picture{
				Height: 5,
				Width:  5,
				Pixels: map[helpers.Co]string{
					{X: 4, Y: 4}: "#",
					{X: 5, Y: 4}: ".",
					{X: 4, Y: 5}: ".",
					{X: 5, Y: 5}: "#",
				},
				TileMap: map[helpers.Co]tile.Tile{
					{X: 2, Y: 2}: {
						ID: "123",
						Pixels: map[helpers.Co]string{
							{X: 0, Y: 0}: ".",
							{X: 1, Y: 0}: "#",
							{X: 2, Y: 0}: "#",
							{X: 3, Y: 0}: ".",
							{X: 0, Y: 1}: ".",
							{X: 1, Y: 1}: "#",
							{X: 2, Y: 1}: ".",
							{X: 3, Y: 1}: ".",
							{X: 0, Y: 2}: ".",
							{X: 1, Y: 2}: ".",
							{X: 2, Y: 2}: "#",
							{X: 3, Y: 2}: "#",
							{X: 0, Y: 3}: "#",
							{X: 1, Y: 3}: "#",
							{X: 2, Y: 3}: "#",
							{X: 3, Y: 3}: ".",
						},
						Height: 3,
						Width:  3,
					},
				},
				Tiles: []tile.Tile{
					{
						ID: "123",
						Pixels: map[helpers.Co]string{
							{X: 0, Y: 0}: ".",
							{X: 1, Y: 0}: "#",
							{X: 2, Y: 0}: "#",
							{X: 3, Y: 0}: ".",
							{X: 0, Y: 1}: ".",
							{X: 1, Y: 1}: "#",
							{X: 2, Y: 1}: ".",
							{X: 3, Y: 1}: ".",
							{X: 0, Y: 2}: ".",
							{X: 1, Y: 2}: ".",
							{X: 2, Y: 2}: "#",
							{X: 3, Y: 2}: "#",
							{X: 0, Y: 3}: "#",
							{X: 1, Y: 3}: "#",
							{X: 2, Y: 3}: "#",
							{X: 3, Y: 3}: ".",
						},
						Height: 3,
						Width:  3,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Picture{
				Pixels:  make(map[helpers.Co]string),
				TileMap: make(map[helpers.Co]tile.Tile),
				Tiles: []tile.Tile{
					{
						ID: "123",
						Pixels: map[helpers.Co]string{
							{X: 0, Y: 0}: ".",
							{X: 1, Y: 0}: "#",
							{X: 2, Y: 0}: "#",
							{X: 3, Y: 0}: ".",
							{X: 0, Y: 1}: ".",
							{X: 1, Y: 1}: "#",
							{X: 2, Y: 1}: ".",
							{X: 3, Y: 1}: ".",
							{X: 0, Y: 2}: ".",
							{X: 1, Y: 2}: ".",
							{X: 2, Y: 2}: "#",
							{X: 3, Y: 2}: "#",
							{X: 0, Y: 3}: "#",
							{X: 1, Y: 3}: "#",
							{X: 2, Y: 3}: "#",
							{X: 3, Y: 3}: ".",
						},
						Height: 3,
						Width:  3,
					},
				},
			}
			p.populatePictureWithTile(tt.args.t, tt.args.x, tt.args.y)
			if !reflect.DeepEqual(p, &tt.want) {
				t.Errorf("Picture.getTopLeftTile() = %v, want &%v", p, tt.want)
			}
		})
	}
}

func TestPicture_PopulateTileMap(t *testing.T) {
	tile1 := tile.Tile{
		ID:     "1",
		Height: 2,
		Width:  2,
		AdjacentTiles: tile.AdjacentTiles{
			Bottom: "4",
			Right:  "2",
		},
		Pixels: map[helpers.Co]string{
			{X: 0, Y: 0}: ".",
			{X: 1, Y: 0}: "#",
			{X: 2, Y: 0}: "#",
			{X: 0, Y: 1}: ".",
			{X: 1, Y: 1}: "#",
			{X: 2, Y: 1}: ".",
			{X: 0, Y: 2}: "#",
			{X: 1, Y: 2}: ".",
			{X: 2, Y: 2}: "#",
		},
	}
	tile2 := tile.Tile{
		ID:     "2",
		Height: 2,
		Width:  2,
		AdjacentTiles: tile.AdjacentTiles{
			Bottom: "5",
			Right:  "3",
			Left:   "1",
		},
		Pixels: map[helpers.Co]string{
			{X: 0, Y: 0}: "#",
			{X: 1, Y: 0}: "#",
			{X: 2, Y: 0}: ".",
			{X: 0, Y: 1}: ".",
			{X: 1, Y: 1}: ".",
			{X: 2, Y: 1}: ".",
			{X: 0, Y: 2}: "#",
			{X: 1, Y: 2}: "#",
			{X: 2, Y: 2}: "#",
		},
	}
	tile3 := tile.Tile{
		ID:     "3",
		Height: 2,
		Width:  2,
		AdjacentTiles: tile.AdjacentTiles{
			Bottom: "6",
			Left:   "2",
		},
		Pixels: map[helpers.Co]string{
			{X: 0, Y: 0}: ".",
			{X: 1, Y: 0}: "#",
			{X: 2, Y: 0}: "#",
			{X: 0, Y: 1}: ".",
			{X: 1, Y: 1}: "#",
			{X: 2, Y: 1}: "#",
			{X: 0, Y: 2}: ".",
			{X: 1, Y: 2}: ".",
			{X: 2, Y: 2}: ".",
		},
	}
	tile4 := tile.Tile{
		ID:     "4",
		Height: 2,
		Width:  2,
		AdjacentTiles: tile.AdjacentTiles{
			Top:    "1",
			Right:  "5",
			Bottom: "7",
		},
		Pixels: map[helpers.Co]string{
			{X: 0, Y: 0}: ".",
			{X: 1, Y: 0}: ".",
			{X: 2, Y: 0}: "#",
			{X: 0, Y: 1}: "#",
			{X: 1, Y: 1}: "#",
			{X: 2, Y: 1}: "#",
			{X: 0, Y: 2}: ".",
			{X: 1, Y: 2}: ".",
			{X: 2, Y: 2}: "#",
		},
	}
	tile5 := tile.Tile{
		ID:     "5",
		Height: 2,
		Width:  2,
		AdjacentTiles: tile.AdjacentTiles{
			Top:    "2",
			Left:   "4",
			Right:  "6",
			Bottom: "8",
		},
		Pixels: map[helpers.Co]string{
			{X: 0, Y: 0}: "#",
			{X: 1, Y: 0}: ".",
			{X: 2, Y: 0}: ".",
			{X: 0, Y: 1}: "#",
			{X: 1, Y: 1}: ".",
			{X: 2, Y: 1}: ".",
			{X: 0, Y: 2}: ".",
			{X: 1, Y: 2}: "#",
			{X: 2, Y: 2}: ".",
		},
	}
	tile6 := tile.Tile{
		ID:     "6",
		Height: 2,
		Width:  2,
		AdjacentTiles: tile.AdjacentTiles{
			Top:    "3",
			Left:   "5",
			Bottom: "9",
		},
		Pixels: map[helpers.Co]string{
			{X: 0, Y: 0}: ".",
			{X: 1, Y: 0}: ".",
			{X: 2, Y: 0}: ".",
			{X: 0, Y: 1}: ".",
			{X: 1, Y: 1}: ".",
			{X: 2, Y: 1}: "#",
			{X: 0, Y: 2}: ".",
			{X: 1, Y: 2}: "#",
			{X: 2, Y: 2}: "#",
		},
	}
	tile7 := tile.Tile{
		ID:     "7",
		Height: 2,
		Width:  2,
		AdjacentTiles: tile.AdjacentTiles{
			Top:   "4",
			Right: "8",
		},
		Pixels: map[helpers.Co]string{
			{X: 0, Y: 0}: ".",
			{X: 1, Y: 0}: "#",
			{X: 2, Y: 0}: ".",
			{X: 0, Y: 1}: ".",
			{X: 1, Y: 1}: "#",
			{X: 2, Y: 1}: ".",
			{X: 0, Y: 2}: ".",
			{X: 1, Y: 2}: "#",
			{X: 2, Y: 2}: "#",
		},
	}
	tile8 := tile.Tile{
		ID:     "8",
		Height: 2,
		Width:  2,
		AdjacentTiles: tile.AdjacentTiles{
			Top:   "5",
			Left:  "7",
			Right: "9",
		},
		Pixels: map[helpers.Co]string{
			{X: 0, Y: 0}: ".",
			{X: 1, Y: 0}: "#",
			{X: 2, Y: 0}: "#",
			{X: 0, Y: 1}: ".",
			{X: 1, Y: 1}: "#",
			{X: 2, Y: 1}: ".",
			{X: 0, Y: 2}: ".",
			{X: 1, Y: 2}: ".",
			{X: 2, Y: 2}: "#",
		},
	}
	tile9 := tile.Tile{
		ID:     "9",
		Height: 2,
		Width:  2,
		AdjacentTiles: tile.AdjacentTiles{
			Top:  "6",
			Left: "8",
		},
		Pixels: map[helpers.Co]string{
			{X: 0, Y: 0}: "#",
			{X: 1, Y: 0}: "#",
			{X: 2, Y: 0}: ".",
			{X: 0, Y: 1}: ".",
			{X: 1, Y: 1}: ".",
			{X: 2, Y: 1}: "#",
			{X: 0, Y: 2}: ".",
			{X: 1, Y: 2}: "#",
			{X: 2, Y: 2}: "#",
		},
	}
	type want struct {
		Height  int
		Width   int
		Pixels  map[helpers.Co]string
		TileMap map[helpers.Co]tile.Tile
		wantErr bool
	}
	tests := []struct {
		name  string
		Tiles []tile.Tile
		want  want
	}{
		{
			name: "returns an error if there is no top left tile",
			Tiles: []tile.Tile{
				{
					ID: "2",
					AdjacentTiles: tile.AdjacentTiles{
						Top:   "1",
						Right: "2",
					},
				},
				{
					ID: "3",
					AdjacentTiles: tile.AdjacentTiles{
						Left: "1",
					},
				},
			},
			want: want{
				Pixels:  make(map[helpers.Co]string),
				TileMap: make(map[helpers.Co]tile.Tile),
				wantErr: true,
			},
		},
		{
			name: "returns an error if a bottom tile's ID cannot be found in picture's tiles",
			Tiles: []tile.Tile{
				{
					ID: "2",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "100",
						Right:  "4",
					},
				},
				{
					ID: "1",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "2",
					},
				},
				{
					ID: "3",
					AdjacentTiles: tile.AdjacentTiles{
						Left: "1",
					},
				},
			},
			want: want{
				Pixels: make(map[helpers.Co]string),
				TileMap: map[helpers.Co]tile.Tile{
					{X: 0, Y: 0}: {
						ID: "2",
						AdjacentTiles: tile.AdjacentTiles{
							Bottom: "100",
							Right:  "4",
						},
					},
				},
				wantErr: true,
			},
		},
		{
			name: "returns an error if a right tile's ID cannot be found in picture's tiles",
			Tiles: []tile.Tile{
				{
					ID: "2",
					AdjacentTiles: tile.AdjacentTiles{
						Bottom: "1",
						Right:  "66",
					},
				},
				{
					ID: "1",
					AdjacentTiles: tile.AdjacentTiles{
						Right: "3",
						Left:  "0",
					},
				},
				{
					ID: "3",
					AdjacentTiles: tile.AdjacentTiles{
						Left: "1",
					},
				},
			},
			want: want{
				Pixels: make(map[helpers.Co]string),
				TileMap: map[helpers.Co]tile.Tile{
					{X: 0, Y: 0}: {
						ID: "2",
						AdjacentTiles: tile.AdjacentTiles{
							Bottom: "1",
							Right:  "66",
						},
					},
					{X: 0, Y: 1}: {
						ID: "1",
						AdjacentTiles: tile.AdjacentTiles{
							Right: "3",
							Left:  "0",
						},
					},
				},
				wantErr: true,
			},
		},
		{
			name:  "populates a simple picture from the given tiles",
			Tiles: []tile.Tile{tile4, tile6, tile7, tile1, tile3, tile5, tile8, tile2, tile9},
			want: want{
				Height: 2,
				Width:  2,
				Pixels: map[helpers.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: "#",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: ".",
					{X: 0, Y: 2}: "#",
					{X: 1, Y: 2}: "#",
					{X: 2, Y: 2}: ".",
				},
				TileMap: map[helpers.Co]tile.Tile{
					{X: 0, Y: 0}: tile1,
					{X: 1, Y: 0}: tile2,
					{X: 2, Y: 0}: tile3,
					{X: 0, Y: 1}: tile4,
					{X: 1, Y: 1}: tile5,
					{X: 2, Y: 1}: tile6,
					{X: 0, Y: 2}: tile7,
					{X: 1, Y: 2}: tile8,
					{X: 2, Y: 2}: tile9,
				},
				wantErr: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Picture{
				Pixels:  make(map[helpers.Co]string),
				TileMap: make(map[helpers.Co]tile.Tile),
				Tiles:   tt.Tiles,
			}
			if err := p.PopulateTileMap(); (err != nil) != tt.want.wantErr {
				t.Errorf("Picture.PopulateTileMap() error = %v, wantErr %v", err, tt.want.wantErr)
			}
			want := &Picture{
				Height:  tt.want.Height,
				Width:   tt.want.Width,
				Pixels:  tt.want.Pixels,
				TileMap: tt.want.TileMap,
				Tiles:   tt.Tiles,
			}
			if !reflect.DeepEqual(p, want) {
				t.Errorf("Picture.PopulateTileMap() = %v, want %v", p, want)
			}
		})
	}
}

func TestPicture_rotatePicture90(t *testing.T) {
	tests := []struct {
		name   string
		Height int
		Width  int
		Pixels map[helpers.Co]string
		want   map[helpers.Co]string
	}{
		{name: "test1",
			Height: 1,
			Width:  1,
			Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
			},
			want: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
			},
		},
		{
			name:   "test2",
			Height: 2,
			Width:  2,
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
			p := &Picture{
				Height: tt.Height,
				Width:  tt.Width,
				Pixels: tt.Pixels,
			}
			p.rotatePicture90()
			if !reflect.DeepEqual(p.Pixels, tt.want) {
				t.Errorf("Picture.rotatePicture90() = %v, want %v", p, tt.want)
			}
		})
	}
}

func TestPicture_flipPicture(t *testing.T) {
	tests := []struct {
		name   string
		Width  int
		Pixels map[helpers.Co]string
		want   map[helpers.Co]string
	}{
		{
			name:  "test1",
			Width: 1,
			Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
			},
			want: map[helpers.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
			},
		},
		{
			name:  "test2",
			Width: 2,
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
			p := &Picture{
				Width:  tt.Width,
				Pixels: tt.Pixels,
			}
			p.flipPicture()
			if !reflect.DeepEqual(p.Pixels, tt.want) {
				t.Errorf("Picture.flipPicture() = %v, want %v", p, tt.want)
			}
		})
	}
}

func TestPicture_markSeaMonster(t *testing.T) {
	tests := []struct {
		name       string
		co         helpers.Co
		seaMonster []helpers.Co
		pixels     map[helpers.Co]string
		want       map[helpers.Co]string
	}{
		{
			name: "it marks pixels as sea monster at the given coordinate",
			co:   helpers.Co{X: 1, Y: 1},
			seaMonster: []helpers.Co{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 1, Y: 1},
			},
			pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
			},
			want: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "O",
				{X: 2, Y: 1}: "O",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "O",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Picture{
				Pixels: tt.pixels,
			}
			p.markSeaMonster(tt.co, tt.seaMonster)
			if !reflect.DeepEqual(p.Pixels, tt.want) {
				t.Errorf("Picture.markSeaMonster() = %v, want %v", p.Pixels, tt.want)
			}
		})
	}
}

func TestPicture_checkSeaMonsterAtCo(t *testing.T) {
	tests := []struct {
		name       string
		Pixels     map[helpers.Co]string
		co         helpers.Co
		seaMonster []helpers.Co
		want       bool
		want1      map[helpers.Co]string
	}{
		{
			name: "it doesn't mark a sea monster at the given co if the sea monster isn't there",
			Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			co: helpers.Co{X: 1, Y: 1},
			seaMonster: []helpers.Co{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 1, Y: 1},
			},
			want: false,
			want1: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
		},
		{
			name: "it doesn't mark a sea monster at the given co if the sea monster runs out the the pixel's boundaries",
			Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
			co: helpers.Co{X: 1, Y: 1},
			seaMonster: []helpers.Co{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 100, Y: 1},
			},
			want: false,
			want1: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
			},
		},
		{
			name: "returns true and marks a sea monster if one is found in the picture",
			co:   helpers.Co{X: 1, Y: 1},
			seaMonster: []helpers.Co{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 1, Y: 1},
			},
			Pixels: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
			},
			want: true,
			want1: map[helpers.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "O",
				{X: 2, Y: 1}: "O",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "O",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Picture{
				Pixels: tt.Pixels,
			}
			if got := p.checkSeaMonsterAtCo(tt.co, tt.seaMonster); got != tt.want {
				t.Errorf("Picture.checkSeaMonsterAtCo() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(p.Pixels, tt.want1) {
				t.Errorf("Picture.checkSeaMonster() = %v, want %v", p.Pixels, tt.want)
			}
		})
	}
}

func TestPicture_FindSeaMonster(t *testing.T) {
	tests := []struct {
		name       string
		Pixels     map[helpers.Co]string
		seaMonster []helpers.Co
		want       map[helpers.Co]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Picture{
				Pixels: tt.Pixels,
			}
			p.FindSeaMonster(tt.seaMonster)
			if !reflect.DeepEqual(p.Pixels, tt.want) {
				t.Errorf("Picture.FindSeaMonster() = %v, want %v", p.Pixels, tt.want)
			}
		})
	}
}
