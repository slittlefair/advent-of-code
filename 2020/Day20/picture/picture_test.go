package picture_test

import (
	"Advent-of-Code/2020/Day20/picture"
	tile "Advent-of-Code/2020/Day20/tile"
	"Advent-of-Code/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPicture_PopulateTiles(t *testing.T) {
	t.Run("populates a simple picture from input", func(t *testing.T) {
		input := []string{
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
		}
		want := picture.Picture{
			Tiles: []tile.Tile{
				{
					ID:     "7",
					Height: 1,
					Width:  1,
					Pixels: map[graph.Co]string{
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
					Pixels: map[graph.Co]string{
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
					Pixels: map[graph.Co]string{
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
					Pixels: map[graph.Co]string{
						{X: 0, Y: 0}: ".",
						{X: 1, Y: 0}: ".",
						{X: 0, Y: 1}: ".",
						{X: 1, Y: 1}: ".",
					},
				},
			},
		}
		p := &picture.Picture{}
		p.PopulateTiles(input)
		assert.Equal(t, want, *p)
	})
}

func TestPicture_FindMatchesForTile(t *testing.T) {
	tests := []struct {
		name  string
		p     picture.Picture
		t     tile.Tile
		index int
		want  picture.Picture
	}{
		{
			name: "skips tile in picture if its ID matches the given tile ID",
			p: picture.Picture{
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
			want: picture.Picture{
				Tiles: []tile.Tile{
					{
						ID: "tile-1",
					},
				},
			},
		},
		{
			name: "skips tile in picture if it is already known to be adjacent to the given tile ID",
			p: picture.Picture{
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
			want: picture.Picture{
				Tiles: []tile.Tile{
					{
						ID: "tile-2",
					},
				},
			},
		},
		{
			name: "skips tile in picture if it is already adjacent to 4 tiles",
			p: picture.Picture{
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
			want: picture.Picture{
				Tiles: []tile.Tile{
					{
						ID: "tile-2",
					},
				},
			},
		},
		{
			name: "skips tile in picture if it is already adjacent to 4 tiles",
			p: picture.Picture{
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
			want: picture.Picture{
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
			p: picture.Picture{
				Tiles: []tile.Tile{
					{
						ID: "tile-1",
						Pixels: map[graph.Co]string{
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
						Pixels: map[graph.Co]string{
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
						Pixels: map[graph.Co]string{
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
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: ".",
					{X: 1, Y: 0}: "#",
					{X: 0, Y: 1}: "#",
					{X: 1, Y: 1}: ".",
				},
				Height: 1,
				Width:  1,
			},
			index: 0,
			want: picture.Picture{
				Tiles: []tile.Tile{
					{
						ID: "tile-1",
						Pixels: map[graph.Co]string{
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
						Pixels: map[graph.Co]string{
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
						Pixels: map[graph.Co]string{
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
			assert.Equal(t, tt.want, tt.p)
		})
	}

	tests2 := []struct {
		name          string
		tile1Pixels   map[graph.Co]string
		tile2Pixels   map[graph.Co]string
		wantPixels    map[graph.Co]string
		want1Adjacent tile.AdjacentTiles
		want2Adjacent tile.AdjacentTiles
	}{
		{
			name: "successfully edits adajcent tiles when tile1 is above tile2, rotate = 0, flip = 0",
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			tile1Pixels: map[graph.Co]string{
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
			tile2Pixels: map[graph.Co]string{
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
			wantPixels: map[graph.Co]string{
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
			p := picture.Picture{
				Tiles: []tile.Tile{tile1, tile2},
			}
			want := picture.Picture{
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
			assert.Equal(t, want, p)
		})
	}
}

func TestPicture_CalculateCornerIDs(t *testing.T) {
	tests := []struct {
		name               string
		Tiles              []tile.Tile
		want               int
		errorAssertionFunc assert.ErrorAssertionFunc
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
			want:               0,
			errorAssertionFunc: assert.Error,
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
			want:               3003,
			errorAssertionFunc: assert.NoError,
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
			want:               110517,
			errorAssertionFunc: assert.NoError,
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
			want:               17391,
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := picture.Picture{
				Tiles: tt.Tiles,
			}
			got, err := p.CalculateCornerIDs()
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPicture_getTileFromID(t *testing.T) {
	tests := []struct {
		name               string
		Tiles              []tile.Tile
		id                 string
		want               tile.Tile
		errorAssertionFunc assert.ErrorAssertionFunc
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
			id:                 "5",
			want:               tile.Tile{},
			errorAssertionFunc: assert.Error,
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
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := picture.Picture{
				Tiles: tt.Tiles,
			}
			got, err := p.GetTileFromID(tt.id)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPicture_getTopLeftTile(t *testing.T) {
	tests := []struct {
		name               string
		Tiles              []tile.Tile
		want               tile.Tile
		errorAssertionFunc assert.ErrorAssertionFunc
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
			want:               tile.Tile{},
			errorAssertionFunc: assert.Error,
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
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := picture.Picture{
				Tiles: tt.Tiles,
			}
			got, err := p.GetTopLeftTile()
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPicture_populatePictureWithTile(t *testing.T) {
	t.Run("correctly populates the given tile into the picture", func(t *testing.T) {
		p := &picture.Picture{
			Pixels:  make(map[graph.Co]string),
			TileMap: make(map[graph.Co]tile.Tile),
			Tiles: []tile.Tile{
				{
					ID: "123",
					Pixels: map[graph.Co]string{
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
		inputTile := tile.Tile{
			ID: "123",
			Pixels: map[graph.Co]string{
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
		}
		want := picture.Picture{
			Height: 5,
			Width:  5,
			Pixels: map[graph.Co]string{
				{X: 4, Y: 4}: "#",
				{X: 5, Y: 4}: ".",
				{X: 4, Y: 5}: ".",
				{X: 5, Y: 5}: "#",
			},
			TileMap: map[graph.Co]tile.Tile{
				{X: 2, Y: 2}: {
					ID: "123",
					Pixels: map[graph.Co]string{
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
					Pixels: map[graph.Co]string{
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
		p.PopulatePictureWithTile(inputTile, 2, 2)
		assert.Equal(t, &want, p)
	})
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
		Pixels: map[graph.Co]string{
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
		Pixels: map[graph.Co]string{
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
		Pixels: map[graph.Co]string{
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
		Pixels: map[graph.Co]string{
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
		Pixels: map[graph.Co]string{
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
		Pixels: map[graph.Co]string{
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
		Pixels: map[graph.Co]string{
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
		Pixels: map[graph.Co]string{
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
		Pixels: map[graph.Co]string{
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
		Height             int
		Width              int
		Pixels             map[graph.Co]string
		TileMap            map[graph.Co]tile.Tile
		errorAssertionFunc assert.ErrorAssertionFunc
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
				Pixels:             make(map[graph.Co]string),
				TileMap:            make(map[graph.Co]tile.Tile),
				errorAssertionFunc: assert.Error,
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
				Pixels: make(map[graph.Co]string),
				TileMap: map[graph.Co]tile.Tile{
					{X: 0, Y: 0}: {
						ID: "2",
						AdjacentTiles: tile.AdjacentTiles{
							Bottom: "100",
							Right:  "4",
						},
					},
				},
				errorAssertionFunc: assert.Error,
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
				Pixels: make(map[graph.Co]string),
				TileMap: map[graph.Co]tile.Tile{
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
				errorAssertionFunc: assert.Error,
			},
		},
		{
			name:  "populates a simple picture from the given tiles",
			Tiles: []tile.Tile{tile4, tile6, tile7, tile1, tile3, tile5, tile8, tile2, tile9},
			want: want{
				Height: 2,
				Width:  2,
				Pixels: map[graph.Co]string{
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
				TileMap: map[graph.Co]tile.Tile{
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
				errorAssertionFunc: assert.NoError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &picture.Picture{
				Pixels:  make(map[graph.Co]string),
				TileMap: make(map[graph.Co]tile.Tile),
				Tiles:   tt.Tiles,
			}
			err := p.PopulateTileMap()
			tt.want.errorAssertionFunc(t, err)
			want := &picture.Picture{
				Height:  tt.want.Height,
				Width:   tt.want.Width,
				Pixels:  tt.want.Pixels,
				TileMap: tt.want.TileMap,
				Tiles:   tt.Tiles,
			}
			assert.Equal(t, want, p)
		})
	}
}

func TestPicture_rotatePicture90(t *testing.T) {
	tests := []struct {
		name   string
		Height int
		Width  int
		Pixels map[graph.Co]string
		want   map[graph.Co]string
	}{
		{name: "test1",
			Height: 1,
			Width:  1,
			Pixels: map[graph.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
			},
			want: map[graph.Co]string{
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
			p := &picture.Picture{
				Height: tt.Height,
				Width:  tt.Width,
				Pixels: tt.Pixels,
			}
			p.RotatePicture90()
			assert.Equal(t, tt.want, p.Pixels)
		})
	}
}

func TestPicture_flipPicture(t *testing.T) {
	tests := []struct {
		name   string
		Width  int
		Pixels map[graph.Co]string
		want   map[graph.Co]string
	}{
		{
			name:  "test1",
			Width: 1,
			Pixels: map[graph.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
			},
			want: map[graph.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
			},
		},
		{
			name:  "test2",
			Width: 2,
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
			p := &picture.Picture{
				Width:  tt.Width,
				Pixels: tt.Pixels,
			}
			p.FlipPicture()
			assert.Equal(t, tt.want, p.Pixels)
		})
	}
}

func TestPicture_markSeaMonster(t *testing.T) {
	t.Run("it marks pixels as sea monster at the given coordinate", func(t *testing.T) {
		p := &picture.Picture{
			Pixels: map[graph.Co]string{
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
		}
		seaMonster := []graph.Co{
			{X: 0, Y: 0},
			{X: 1, Y: 0},
			{X: 1, Y: 1},
		}
		want := map[graph.Co]string{
			{X: 0, Y: 0}: "#",
			{X: 1, Y: 0}: "#",
			{X: 2, Y: 0}: ".",
			{X: 0, Y: 1}: "#",
			{X: 1, Y: 1}: "O",
			{X: 2, Y: 1}: "O",
			{X: 0, Y: 2}: ".",
			{X: 1, Y: 2}: "#",
			{X: 2, Y: 2}: "O",
		}
		p.MarkSeaMonster(graph.Co{X: 1, Y: 1}, seaMonster)
		assert.Equal(t, want, p.Pixels)
	})
}

func TestPicture_checkSeaMonsterAtCo(t *testing.T) {
	tests := []struct {
		name       string
		Pixels     map[graph.Co]string
		co         graph.Co
		seaMonster []graph.Co
		want       bool
		want1      map[graph.Co]string
	}{
		{
			name: "it doesn't mark a sea monster at the given co if the sea monster isn't there",
			Pixels: map[graph.Co]string{
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
			co: graph.Co{X: 1, Y: 1},
			seaMonster: []graph.Co{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 1, Y: 1},
			},
			want: false,
			want1: map[graph.Co]string{
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
			Pixels: map[graph.Co]string{
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
			co: graph.Co{X: 1, Y: 1},
			seaMonster: []graph.Co{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 100, Y: 1},
			},
			want: false,
			want1: map[graph.Co]string{
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
			co:   graph.Co{X: 1, Y: 1},
			seaMonster: []graph.Co{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 1, Y: 1},
			},
			Pixels: map[graph.Co]string{
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
			want1: map[graph.Co]string{
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
			p := &picture.Picture{
				Pixels: tt.Pixels,
			}
			got := p.CheckSeaMonsterAtCo(tt.co, tt.seaMonster)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, p.Pixels)
		})
	}
}

func TestPicture_FindSeaMonster(t *testing.T) {
	tests := []struct {
		name       string
		Pixels     map[graph.Co]string
		seaMonster []graph.Co
		want       map[graph.Co]string
	}{
		{
			name: "does not affect the picture's pixels if sea monster is not present",
			Pixels: map[graph.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: "#",
				{X: 3, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: ".",
				{X: 3, Y: 1}: "#",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: "#",
				{X: 3, Y: 2}: ".",
				{X: 0, Y: 3}: ".",
				{X: 1, Y: 3}: "#",
				{X: 2, Y: 3}: ".",
				{X: 3, Y: 3}: "#",
			},
			seaMonster: []graph.Co{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
			},
			want: map[graph.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: "#",
				{X: 3, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: ".",
				{X: 3, Y: 1}: "#",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: "#",
				{X: 3, Y: 2}: ".",
				{X: 0, Y: 3}: ".",
				{X: 1, Y: 3}: "#",
				{X: 2, Y: 3}: ".",
				{X: 3, Y: 3}: "#",
			},
		},
		{
			name: "marks a found sea monster in the picture, rotate = 0, flip = 0",
			Pixels: map[graph.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 3, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 3, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
				{X: 3, Y: 2}: ".",
				{X: 0, Y: 3}: "#",
				{X: 1, Y: 3}: ".",
				{X: 2, Y: 3}: ".",
				{X: 3, Y: 3}: ".",
			},
			seaMonster: []graph.Co{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 3, Y: 0},
				{X: 1, Y: 1},
			},
			want: map[graph.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 3, Y: 0}: ".",
				{X: 0, Y: 1}: "O",
				{X: 1, Y: 1}: "O",
				{X: 2, Y: 1}: "O",
				{X: 3, Y: 1}: "O",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "O",
				{X: 2, Y: 2}: "#",
				{X: 3, Y: 2}: ".",
				{X: 0, Y: 3}: "#",
				{X: 1, Y: 3}: ".",
				{X: 2, Y: 3}: ".",
				{X: 3, Y: 3}: ".",
			},
		},
		{
			name: "marks a found sea monster in the picture, rotate = 1, flip = 0",
			Pixels: map[graph.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 3, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 3, Y: 1}: ".",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
				{X: 3, Y: 2}: ".",
				{X: 0, Y: 3}: "#",
				{X: 1, Y: 3}: ".",
				{X: 2, Y: 3}: ".",
				{X: 3, Y: 3}: "#",
			},
			seaMonster: []graph.Co{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 3, Y: 0},
				{X: 1, Y: 1},
			},
			want: map[graph.Co]string{
				{X: 0, Y: 0}: "O",
				{X: 1, Y: 0}: "O",
				{X: 2, Y: 0}: "O",
				{X: 3, Y: 0}: "O",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: "O",
				{X: 2, Y: 1}: ".",
				{X: 3, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
				{X: 3, Y: 2}: ".",
				{X: 0, Y: 3}: "#",
				{X: 1, Y: 3}: ".",
				{X: 2, Y: 3}: ".",
				{X: 3, Y: 3}: ".",
			},
		},
		{
			name: "marks a found sea monster in the picture, rotate = 2, flip = 0",
			Pixels: map[graph.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: "#",
				{X: 3, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 3, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
				{X: 3, Y: 2}: "#",
				{X: 0, Y: 3}: "#",
				{X: 1, Y: 3}: "#",
				{X: 2, Y: 3}: "#",
				{X: 3, Y: 3}: "#",
			},
			seaMonster: []graph.Co{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 3, Y: 0},
				{X: 1, Y: 1},
			},
			want: map[graph.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 3, Y: 0}: "#",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 3, Y: 1}: ".",
				{X: 0, Y: 2}: "O",
				{X: 1, Y: 2}: "O",
				{X: 2, Y: 2}: "O",
				{X: 3, Y: 2}: "O",
				{X: 0, Y: 3}: ".",
				{X: 1, Y: 3}: "O",
				{X: 2, Y: 3}: ".",
				{X: 3, Y: 3}: ".",
			},
		},
		{
			name: "marks a found sea monster in the picture, rotate = 3, flip = 0",
			Pixels: map[graph.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: "#",
				{X: 3, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 3, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: "#",
				{X: 3, Y: 2}: ".",
				{X: 0, Y: 3}: ".",
				{X: 1, Y: 3}: ".",
				{X: 2, Y: 3}: "#",
				{X: 3, Y: 3}: "#",
			},
			seaMonster: []graph.Co{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 3, Y: 0},
				{X: 1, Y: 1},
			},
			want: map[graph.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 3, Y: 0}: "#",
				{X: 0, Y: 1}: "O",
				{X: 1, Y: 1}: "O",
				{X: 2, Y: 1}: "O",
				{X: 3, Y: 1}: "O",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "O",
				{X: 2, Y: 2}: ".",
				{X: 3, Y: 2}: ".",
				{X: 0, Y: 3}: ".",
				{X: 1, Y: 3}: "#",
				{X: 2, Y: 3}: ".",
				{X: 3, Y: 3}: ".",
			},
		},
		{
			name: "marks a found sea monster in the picture, rotate = 0, flip = 1",
			Pixels: map[graph.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 3, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 3, Y: 1}: ".",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
				{X: 3, Y: 2}: "#",
				{X: 0, Y: 3}: ".",
				{X: 1, Y: 3}: ".",
				{X: 2, Y: 3}: "#",
				{X: 3, Y: 3}: ".",
			},
			seaMonster: []graph.Co{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 3, Y: 0},
				{X: 1, Y: 1},
			},
			want: map[graph.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 3, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 3, Y: 1}: ".",
				{X: 0, Y: 2}: "O",
				{X: 1, Y: 2}: "O",
				{X: 2, Y: 2}: "O",
				{X: 3, Y: 2}: "O",
				{X: 0, Y: 3}: ".",
				{X: 1, Y: 3}: "O",
				{X: 2, Y: 3}: ".",
				{X: 3, Y: 3}: ".",
			},
		},
		{
			name: "marks a found sea monster in the picture, rotate = 1, flip = 1",
			Pixels: map[graph.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 3, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 3, Y: 1}: "#",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: ".",
				{X: 3, Y: 2}: ".",
				{X: 0, Y: 3}: "#",
				{X: 1, Y: 3}: ".",
				{X: 2, Y: 3}: ".",
				{X: 3, Y: 3}: ".",
			},
			seaMonster: []graph.Co{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 3, Y: 0},
				{X: 1, Y: 1},
			},
			want: map[graph.Co]string{
				{X: 0, Y: 0}: "O",
				{X: 1, Y: 0}: "O",
				{X: 2, Y: 0}: "O",
				{X: 3, Y: 0}: "O",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: "O",
				{X: 2, Y: 1}: ".",
				{X: 3, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
				{X: 3, Y: 2}: ".",
				{X: 0, Y: 3}: ".",
				{X: 1, Y: 3}: "#",
				{X: 2, Y: 3}: ".",
				{X: 3, Y: 3}: ".",
			},
		},
		{
			name: "marks a found sea monster in the picture, rotate = 2, flip = 1",
			Pixels: map[graph.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 3, Y: 0}: ".",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "#",
				{X: 3, Y: 1}: "#",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: "#",
				{X: 3, Y: 2}: "#",
				{X: 0, Y: 3}: "#",
				{X: 1, Y: 3}: "#",
				{X: 2, Y: 3}: "#",
				{X: 3, Y: 3}: "#",
			},
			seaMonster: []graph.Co{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 3, Y: 0},
				{X: 1, Y: 1},
			},
			want: map[graph.Co]string{
				{X: 0, Y: 0}: "O",
				{X: 1, Y: 0}: "O",
				{X: 2, Y: 0}: "O",
				{X: 3, Y: 0}: "O",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "O",
				{X: 2, Y: 1}: ".",
				{X: 3, Y: 1}: ".",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: ".",
				{X: 3, Y: 2}: ".",
				{X: 0, Y: 3}: ".",
				{X: 1, Y: 3}: ".",
				{X: 2, Y: 3}: ".",
				{X: 3, Y: 3}: ".",
			},
		},
		{
			name: "marks a found sea monster in the picture, rotate = 3, flip = 1",
			Pixels: map[graph.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: "#",
				{X: 3, Y: 0}: ".",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: "#",
				{X: 3, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
				{X: 3, Y: 2}: ".",
				{X: 0, Y: 3}: "#",
				{X: 1, Y: 3}: ".",
				{X: 2, Y: 3}: "#",
				{X: 3, Y: 3}: ".",
			},
			seaMonster: []graph.Co{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 3, Y: 0},
				{X: 1, Y: 1},
			},
			want: map[graph.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 3, Y: 0}: ".",
				{X: 0, Y: 1}: "O",
				{X: 1, Y: 1}: "O",
				{X: 2, Y: 1}: "O",
				{X: 3, Y: 1}: "O",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "O",
				{X: 2, Y: 2}: ".",
				{X: 3, Y: 2}: ".",
				{X: 0, Y: 3}: "#",
				{X: 1, Y: 3}: ".",
				{X: 2, Y: 3}: "#",
				{X: 3, Y: 3}: "#",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &picture.Picture{
				Pixels: tt.Pixels,
				Height: 3,
				Width:  3,
			}
			p.FindSeaMonster(tt.seaMonster)
			assert.Equal(t, tt.want, p.Pixels)
		})
	}
}

func TestPicture_CountWaterRoughness(t *testing.T) {
	tests := []struct {
		name   string
		Pixels map[graph.Co]string
		want   int
	}{
		{
			name: "returns 0 if there are no water symbols",
			Pixels: map[graph.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 3, Y: 0}: ".",
				{X: 0, Y: 1}: "O",
				{X: 1, Y: 1}: "O",
				{X: 2, Y: 1}: "O",
				{X: 3, Y: 1}: "O",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "O",
				{X: 2, Y: 2}: ".",
				{X: 3, Y: 2}: ".",
				{X: 0, Y: 3}: ".",
				{X: 1, Y: 3}: ".",
				{X: 2, Y: 3}: ".",
				{X: 3, Y: 3}: ".",
			},
			want: 0,
		},
		{
			name: "returns 1 if there is 1 water symbol",
			Pixels: map[graph.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: ".",
				{X: 3, Y: 0}: ".",
				{X: 0, Y: 1}: "O",
				{X: 1, Y: 1}: "O",
				{X: 2, Y: 1}: "O",
				{X: 3, Y: 1}: "O",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "O",
				{X: 2, Y: 2}: ".",
				{X: 3, Y: 2}: ".",
				{X: 0, Y: 3}: ".",
				{X: 1, Y: 3}: "#",
				{X: 2, Y: 3}: ".",
				{X: 3, Y: 3}: ".",
			},
			want: 1,
		},
		{
			name: "returns the correct number of water symbols in the picture",
			Pixels: map[graph.Co]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 3, Y: 0}: ".",
				{X: 0, Y: 1}: "O",
				{X: 1, Y: 1}: "O",
				{X: 2, Y: 1}: "O",
				{X: 3, Y: 1}: "O",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "O",
				{X: 2, Y: 2}: ".",
				{X: 3, Y: 2}: "#",
				{X: 0, Y: 3}: "#",
				{X: 1, Y: 3}: ".",
				{X: 2, Y: 3}: "#",
				{X: 3, Y: 3}: "#",
			},
			want: 7,
		},
		{
			name: "returns the maximum nunber (number of pixels) if all are water symbols",
			Pixels: map[graph.Co]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: "#",
				{X: 3, Y: 0}: "#",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: "#",
				{X: 2, Y: 1}: "#",
				{X: 3, Y: 1}: "#",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
				{X: 3, Y: 2}: "#",
				{X: 0, Y: 3}: "#",
				{X: 1, Y: 3}: "#",
				{X: 2, Y: 3}: "#",
				{X: 3, Y: 3}: "#",
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := picture.Picture{
				Pixels: tt.Pixels,
			}
			got := p.CountWaterRoughness()
			assert.Equal(t, tt.want, got)
		})
	}
}
