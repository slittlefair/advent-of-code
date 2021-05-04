package picture

import (
	helpers "Advent-of-Code"
	tile "Advent-of-Code/2020/Day20/tile"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Picture struct {
	Height  int
	Width   int
	Pixels  map[helpers.Co]string
	TileMap map[helpers.Co]tile.Tile
	Tiles   []tile.Tile
}

// Populate tiles takes the string of lines, input, and populates the tiles accordingly
func (p *Picture) PopulateTiles(input []string) {
	re := regexp.MustCompile(`\d+`)
	t := tile.Tile{
		Pixels: make(map[helpers.Co]string),
	}
	var i int
	for _, line := range input {
		if line == "" {
			p.Tiles = append(p.Tiles, t)
			t = tile.Tile{
				Pixels: make(map[helpers.Co]string),
			}
			continue
		}
		if match := re.FindString(line); match != "" {
			t.ID = match
			i = 0
			continue
		}
		t.Height = i
		t.Width = len(line) - 1
		for j, char := range line {
			t.Pixels[helpers.Co{X: j, Y: i}] = string(char)
		}
		i++
	}
	p.Tiles = append(p.Tiles, t)
}

// FindMatchesForTile finds any tiles that are adjacent to tile t and sets them as either "top",
// "bottom", "left" or "right" on the respective tiles
func (p Picture) FindMatchesForTile(t tile.Tile, index int) {
	for k, tile := range p.Tiles {
		if tile.ID == t.ID {
			continue
		}
		if t.IsAdjacentTo(tile) {
			continue
		}
		if t.NumAdjacent() == 4 {
			continue
		}
		if tile.NumAdjacent() == 4 {
			continue
		}
		for j := 0; j < 2; j++ {
			for i := 0; i < 4; i++ {
				if t.IsAdjacentTop(tile) {
					tileAtK := &p.Tiles[k]
					tileAtK.AdjacentTiles.Top = t.ID
					tileAtK.Pixels = tile.Pixels
					tileAtIndex := &p.Tiles[index]
					tileAtIndex.AdjacentTiles.Bottom = tile.ID
					p.FindMatchesForTile(*tileAtK, k)
					goto out
				}
				if t.IsAdjacentBottom(tile) {
					tileAtK := &p.Tiles[k]
					tileAtK.AdjacentTiles.Bottom = t.ID
					tileAtK.Pixels = tile.Pixels
					tileAtIndex := &p.Tiles[index]
					tileAtIndex.AdjacentTiles.Top = tile.ID
					p.FindMatchesForTile(*tileAtK, k)
					goto out
				}
				if t.IsAdjacentLeft(tile) {
					tileAtK := &p.Tiles[k]
					tileAtK.AdjacentTiles.Left = t.ID
					tileAtK.Pixels = tile.Pixels
					tileAtIndex := &p.Tiles[index]
					tileAtIndex.AdjacentTiles.Right = tile.ID
					p.FindMatchesForTile(*tileAtK, k)
					goto out
				}
				if t.IsAdjacentRight(tile) {
					tileAtK := &p.Tiles[k]
					tileAtK.AdjacentTiles.Right = t.ID
					tileAtK.Pixels = tile.Pixels
					tileAtIndex := &p.Tiles[index]
					tileAtIndex.AdjacentTiles.Left = tile.ID
					p.FindMatchesForTile(*tileAtK, k)
					goto out
				}
				tile.RotateTile90()
			}
			tile.FlipTile()
		}
	out:
	}
}

// CalculateCornerIDs returns the product of the IDs for the four corner tile in picture p. Since
// the IDs are kept as strings they need to be converted to ints, so we return an error if this
// conversion can't be done. The value returned here is the solution to Part 1 of the challenge
func (p Picture) CalculateCornerIDs() (int, error) {
	cornerID := 1
	for _, tile := range p.Tiles {
		if tile.NumAdjacent() == 2 {
			numericID, err := strconv.Atoi(tile.ID)
			if err != nil {
				return 0, err
			}
			cornerID *= numericID
		}
	}
	return cornerID, nil
}

// getTileFromID returns the tile that has the given ID. If no such tile exists then return an
// error, as something has obviously gone wrong.
func (p Picture) getTileFromID(id string) (tile.Tile, error) {
	for _, tile := range p.Tiles {
		if tile.ID == id {
			return tile, nil
		}
	}
	return tile.Tile{}, errors.New(fmt.Sprintln("could not find tile for id:", id))
}

// getTopLeftTile returns the tile at the top left of the picture. We do this so we can start
// creating a picture from this tile, with all other tiles relative to it. At the point we call this
// we expect there to be a top left tile, so return an error if one doesn't exist
func (p Picture) getTopLeftTile() (tile.Tile, error) {
	for _, t := range p.Tiles {
		if t.AdjacentTiles.Bottom != "" && t.AdjacentTiles.Right != "" && t.AdjacentTiles.Left == "" && t.AdjacentTiles.Top == "" {
			return t, nil
		}
	}
	return tile.Tile{}, errors.New("no tile found at top left")
}

// populatePictureWithTile gets the individual tiles and puts their pixels into the picture p's
// pixels, into one large picture
func (p *Picture) populatePictureWithTile(t tile.Tile, x, y int) {
	p.TileMap[helpers.Co{X: x, Y: y}] = t
	for i := 1; i < t.Height; i++ {
		for j := 1; j < t.Width; j++ {
			xValue := (x * (t.Width + 1)) + j - 2*x - 1
			yValue := (y * (t.Height + 1)) + i - 2*y - 1
			p.Pixels[helpers.Co{X: xValue, Y: yValue}] = t.Pixels[helpers.Co{X: j, Y: i}]
			if yValue > p.Height {
				p.Height = yValue
			}
			if xValue > p.Width {
				p.Width = xValue
			}
		}
	}
}

// PopulateTileMap gets the individual tiles and determines their relativity to each other, in order
// to create a single set of pixels
func (p *Picture) PopulateTileMap() error {
	// Find the top left tile first, then build out from there
	var x, y int
	tile, err := p.getTopLeftTile()
	if err != nil {
		return err
	}
	p.populatePictureWithTile(tile, x, y)
	for {
		if tile.AdjacentTiles.Bottom != "" {
			t, err := p.getTileFromID(tile.AdjacentTiles.Bottom)
			if err != nil {
				return err
			}
			y++
			p.populatePictureWithTile(t, x, y)
			tile = t
		} else {
			tile = p.TileMap[helpers.Co{X: x, Y: 0}]
			if tile.AdjacentTiles.Right == "" {
				return nil
			}
			t, err := p.getTileFromID(tile.AdjacentTiles.Right)
			if err != nil {
				return err
			}
			x++
			y = 0
			p.populatePictureWithTile(t, x, y)
			tile = t
		}
	}
}

// PrintPictureMap is a helper method which prints a visual representation of the pixels of a tile
// in the console
func (p Picture) PrintPictureMap() {
	for h := 0; h <= p.Height; h++ {
		for w := 0; w <= p.Width; w++ {
			fmt.Print(p.Pixels[helpers.Co{X: w, Y: h}])
		}
		fmt.Println()
	}
}

// rotatePicture90 rotates the pixels in the picture p by 90 degrees
func (p *Picture) rotatePicture90() {
	newPixels := make(map[helpers.Co]string)
	for co, val := range p.Pixels {
		newPixels[helpers.Co{X: p.Width - co.Y, Y: co.X}] = val
	}
	p.Pixels = newPixels
}

// flipPicture flips (reflects) the pixels of a tile along the vertical centre
func (p *Picture) flipPicture() {
	newPixels := make(map[helpers.Co]string)
	for co, val := range p.Pixels {
		newPixels[helpers.Co{X: p.Width - co.X, Y: co.Y}] = val
	}
	p.Pixels = newPixels
}

// markSeaMonster changes the pixels in the picture that represent a sea monster from "#" to "O"
func (p *Picture) markSeaMonster(co helpers.Co, seaMonster []helpers.Co) {
	for _, smCo := range seaMonster {
		p.Pixels[helpers.Co{X: co.X + smCo.X, Y: co.Y + smCo.Y}] = "O"
	}
}

// checkSeaMonsterAtCo checks to see if a sea monster is in the picture relative to the given
// coordinate, co
func (p *Picture) checkSeaMonsterAtCo(co helpers.Co, seaMonster []helpers.Co) bool {
	for _, smCo := range seaMonster {
		c := helpers.Co{X: co.X + smCo.X, Y: co.Y + smCo.Y}
		if val, ok := p.Pixels[c]; val != "#" || !ok {
			return false
		}
	}
	p.markSeaMonster(co, seaMonster)
	return true
}

// FindSeaMonster iterates through the picture p and finds any sea monsters. They only exist on one
// orientation of the picture, so if we find at least one then we can find the rest in that
// orientation then return early. Otherwise we keep rotating the picture, then flip it before
// rotating it again.
func (p *Picture) FindSeaMonster(seaMonster []helpers.Co) {
	var found bool
	for j := 0; j < 2; j++ {
		for i := 0; i < 4; i++ {
			for co := range p.Pixels {
				if p.checkSeaMonsterAtCo(co, seaMonster) {
					found = true
				}
			}
			if found {
				return
			}
			p.rotatePicture90()
		}
		p.flipPicture()
	}
}

// CountWaterRoughness returns the number of rough patches of water "#" in the picture. The value
// returned here is the solution to Part 2 of the challenge, as long as it's done after we exclude
// all sea monster pixels via markSeaMonster
func (p Picture) CountWaterRoughness() int {
	count := 0
	for _, val := range p.Pixels {
		if val == "#" {
			count++
		}
	}
	return count
}
