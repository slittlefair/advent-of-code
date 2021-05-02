package tile

import (
	helpers "Advent-of-Code"
	"fmt"
)

type AdjacentTiles struct {
	Top    string
	Bottom string
	Left   string
	Right  string
}

type Tile struct {
	ID            string
	Pixels        map[helpers.Co]string
	Height        int
	Width         int
	AdjacentTiles AdjacentTiles
}

// PrintPixels is a helper method which prints a visual representation of the pixels of a tile in
// the console
func (t Tile) PrintPixels() {
	for h := 0; h <= t.Height; h++ {
		for w := 0; w <= t.Width; w++ {
			fmt.Print(t.Pixels[helpers.Co{X: w, Y: h}])
		}
		fmt.Println()
	}
}

// RotateTile90 rotates the pixels of a tile by 90 degrees
func (t *Tile) RotateTile90() {
	newPixels := make(map[helpers.Co]string)
	for co, val := range t.Pixels {
		newPixels[helpers.Co{X: t.Width - co.Y, Y: co.X}] = val
	}
	t.Pixels = newPixels
}

// FlipTile flips (reflects) the pixels of a tile along the vertical centre
func (t *Tile) FlipTile() {
	newPixels := make(map[helpers.Co]string)
	for co, val := range t.Pixels {
		newPixels[helpers.Co{X: t.Width - co.X, Y: co.Y}] = val
	}
	t.Pixels = newPixels
}

// IsAdjacentTop returns whether the tile "t" is directly above the tile "tile"
func (t Tile) IsAdjacentTop(tile Tile) bool {
	for x := 0; x <= t.Width; x++ {
		if t.Pixels[helpers.Co{X: x, Y: t.Height}] != tile.Pixels[helpers.Co{X: x, Y: 0}] {
			return false
		}
	}
	return true
}

// IsAdjacentTop returns whether the tile "t" is directly below the tile "tile"
func (t Tile) IsAdjacentBottom(tile Tile) bool {
	for x := 0; x <= t.Width; x++ {
		if t.Pixels[helpers.Co{X: x, Y: 0}] != tile.Pixels[helpers.Co{X: x, Y: tile.Height}] {
			return false
		}
	}
	return true
}

// IsAdjacentTop returns whether the tile "t" is directly to the left of tile "tile"
func (t Tile) IsAdjacentLeft(tile Tile) bool {
	for y := 0; y <= t.Height; y++ {
		if t.Pixels[helpers.Co{X: t.Width, Y: y}] != tile.Pixels[helpers.Co{X: 0, Y: y}] {
			return false
		}
	}
	return true
}

// IsAdjacentTop returns whether the tile "t" is directly to the right of tile "tile"
func (t Tile) IsAdjacentRight(tile Tile) bool {
	for y := 0; y <= t.Height; y++ {
		if t.Pixels[helpers.Co{X: 0, Y: y}] != tile.Pixels[helpers.Co{X: tile.Width, Y: y}] {
			return false
		}
	}
	return true
}

// isAdjacentTo returns whether the tile "tile" is adjacent to the tile "t"
func (t Tile) IsAdjacentTo(tile Tile) bool {
	if t.AdjacentTiles.Top == tile.ID {
		return true
	}
	if t.AdjacentTiles.Bottom == tile.ID {
		return true
	}
	if t.AdjacentTiles.Left == tile.ID {
		return true
	}
	if t.AdjacentTiles.Right == tile.ID {
		return true
	}
	return false
}

// NumAdjacent returns how many tiles are adjacent to the tile "t"
func (t Tile) NumAdjacent() int {
	numAdjacent := 0
	if t.AdjacentTiles.Top != "" {
		numAdjacent++
	}
	if t.AdjacentTiles.Bottom != "" {
		numAdjacent++
	}
	if t.AdjacentTiles.Left != "" {
		numAdjacent++
	}
	if t.AdjacentTiles.Right != "" {
		numAdjacent++
	}
	return numAdjacent
}
