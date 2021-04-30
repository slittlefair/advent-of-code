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

// TODO add comments to exported functions

func (p Picture) FindMatchesForTile(t tile.Tile, index int) {
	for k, tile := range p.Tiles {
		if tile.ID == t.ID || t.IsAdjacentTo(tile) || t.NumAdjacent() == 4 || tile.NumAdjacent() == 4 {
			continue
		}
		for j := 0; j < 2; j++ {
			for i := 0; i < 4; i++ {
				if t.IsAdjacentTop(tile) {
					p.Tiles[k].AdjacentTiles.Bottom = t.ID
					p.Tiles[index].AdjacentTiles.Top = tile.ID
					p.Tiles[k].Pixels = tile.Pixels
					p.FindMatchesForTile(tile, k)
					break
				}
				if t.IsAdjacentBottom(tile) {
					p.Tiles[k].AdjacentTiles.Top = t.ID
					p.Tiles[index].AdjacentTiles.Bottom = tile.ID
					p.Tiles[k].Pixels = tile.Pixels
					p.FindMatchesForTile(tile, k)
					break
				}
				if t.IsAdjacentLeft(tile) {
					p.Tiles[k].AdjacentTiles.Right = t.ID
					p.Tiles[index].AdjacentTiles.Left = tile.ID
					p.Tiles[k].Pixels = tile.Pixels
					p.FindMatchesForTile(tile, k)
					break
				}
				if t.IsAdjacentRight(tile) {
					p.Tiles[k].AdjacentTiles.Left = t.ID
					p.Tiles[index].AdjacentTiles.Right = tile.ID
					p.Tiles[k].Pixels = tile.Pixels
					p.FindMatchesForTile(tile, k)
					break
				}
				tile.RotateTile90()
			}
			tile.FlipTile()
		}
	}
}

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

func (p Picture) getTileFromID(id string) (tile.Tile, error) {
	for _, tile := range p.Tiles {
		if tile.ID == id {
			return tile, nil
		}
	}
	return tile.Tile{}, errors.New(fmt.Sprintln("could not find tile for id:", id))
}

func (p Picture) getTopLeftTile() (tile.Tile, error) {
	for _, t := range p.Tiles {
		if t.AdjacentTiles.Bottom != "" && t.AdjacentTiles.Right != "" && t.AdjacentTiles.Left == "" && t.AdjacentTiles.Top == "" {
			return t, nil
		}
	}
	return tile.Tile{}, errors.New("no tile found at top left")
}

func (p *Picture) populatePictureWithTile(t tile.Tile, x, y int) {
	p.TileMap[helpers.Co{X: x, Y: y}] = t
	for i := 1; i < t.Height; i++ {
		for j := 1; j < t.Width; j++ {
			xValue := (x * (t.Width + 1)) + j - 2*x - 1
			yValue := (y * (t.Height + 1)) + i - 2*y - 1
			// for some reason the pixels in the tiles are flipped vertically, so populate with t.Height - i rather than just i
			p.Pixels[helpers.Co{X: xValue, Y: yValue}] = t.Pixels[helpers.Co{X: j, Y: t.Height - i}]
			if yValue > p.Height {
				p.Height = yValue
			}
			if xValue > p.Width {
				p.Width = xValue
			}
		}
	}
}

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

func (p Picture) PrintPictureMap() {
	for h := 0; h <= p.Height; h++ {
		for w := 0; w <= p.Width; w++ {
			fmt.Print(p.Pixels[helpers.Co{X: w, Y: h}])
		}
		fmt.Println()
	}
}

var seaMonster = []helpers.Co{
	{X: 0, Y: 1},
	{X: 1, Y: 2},
	{X: 4, Y: 2},
	{X: 5, Y: 1},
	{X: 6, Y: 1},
	{X: 7, Y: 2},
	{X: 10, Y: 2},
	{X: 11, Y: 1},
	{X: 12, Y: 1},
	{X: 13, Y: 2},
	{X: 16, Y: 2},
	{X: 17, Y: 1},
	{X: 18, Y: 0},
	{X: 18, Y: 1},
	{X: 19, Y: 1},
}

func (p *Picture) rotatePicture90() {
	newPixels := make(map[helpers.Co]string)
	for co, val := range p.Pixels {
		newPixels[helpers.Co{X: p.Width - co.Y, Y: co.X}] = val
	}
	p.Pixels = newPixels
}

func (p *Picture) flipPicture() {
	newPixels := make(map[helpers.Co]string)
	for co, val := range p.Pixels {
		newPixels[helpers.Co{X: p.Width - co.X, Y: co.Y}] = val
	}
	p.Pixels = newPixels
}

func (p *Picture) FindSeaMonster() {
	var found bool
	for j := 0; j < 2; j++ {
		for i := 0; i < 4; i++ {
			for co := range p.Pixels {
				if p.checkSeaMonsterAtCo(co) {
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

func (p *Picture) checkSeaMonsterAtCo(co helpers.Co) bool {
	for _, smCo := range seaMonster {
		c := helpers.Co{X: co.X + smCo.X, Y: co.Y + smCo.Y}
		if val, ok := p.Pixels[c]; val != "#" || !ok {
			return false
		}
	}
	p.markSeaMonster(co)
	return true
}

func (p *Picture) markSeaMonster(co helpers.Co) {
	for _, smCo := range seaMonster {
		p.Pixels[helpers.Co{X: co.X + smCo.X, Y: co.Y + smCo.Y}] = "O"
	}
}

func (p Picture) CountWaterRoughness() int {
	count := 0
	for _, val := range p.Pixels {
		if val == "#" {
			count++
		}
	}
	return count
}
