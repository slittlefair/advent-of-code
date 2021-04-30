package tile

import (
	helpers "Advent-of-Code"
)

type AdjacentTiles struct {
	Top    string
	Bottom string
	Left   string
	Right  string
}

type Tile struct {
	ID            string
	Pixels        map[helpers.Coordinate]string
	Height        int
	Width         int
	AdjacentTiles AdjacentTiles
}

// func (t Tile) printPixels() {
// 	for h := 0; h <= t.height; h++ {
// 		for w := 0; w <= t.width; w++ {
// 			fmt.Print(t.pixels[helpers.Coordinate{X: w, Y: h}])
// 		}
// 		fmt.Println()
// 	}
// }

// TODO add comments to exported functions

func (t *Tile) RotateTile90() {
	newPixels := make(map[helpers.Coordinate]string)
	for co, val := range t.Pixels {
		newPixels[helpers.Coordinate{X: t.Width - co.Y, Y: co.X}] = val
	}
	t.Pixels = newPixels
}

func (t *Tile) FlipTile() {
	newPixels := make(map[helpers.Coordinate]string)
	for co, val := range t.Pixels {
		newPixels[helpers.Coordinate{X: t.Width - co.X, Y: co.Y}] = val
	}
	t.Pixels = newPixels
}

func (t Tile) IsAdjacentTop(tile Tile) bool {
	for x := 0; x <= t.Width; x++ {
		if t.Pixels[helpers.Coordinate{X: x, Y: t.Height}] != tile.Pixels[helpers.Coordinate{X: x, Y: 0}] {
			return false
		}
	}
	return true
}

func (t Tile) IsAdjacentBottom(tile Tile) bool {
	for x := 0; x <= t.Width; x++ {
		if t.Pixels[helpers.Coordinate{X: x, Y: 0}] != tile.Pixels[helpers.Coordinate{X: x, Y: t.Height}] {
			return false
		}
	}
	return true
}

func (t Tile) IsAdjacentLeft(tile Tile) bool {
	for y := 0; y <= t.Height; y++ {
		if t.Pixels[helpers.Coordinate{X: 0, Y: y}] != tile.Pixels[helpers.Coordinate{X: tile.Width, Y: y}] {
			return false
		}
	}
	return true
}

func (t Tile) IsAdjacentRight(tile Tile) bool {
	for y := 0; y <= t.Height; y++ {
		if t.Pixels[helpers.Coordinate{X: t.Width, Y: y}] != tile.Pixels[helpers.Coordinate{X: 0, Y: y}] {
			return false
		}
	}
	return true
}

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
