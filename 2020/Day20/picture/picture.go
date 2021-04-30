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
	Pixels  map[helpers.Coordinate]string
	TileMap map[helpers.Coordinate]tile.Tile
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
		Pixels: make(map[helpers.Coordinate]string),
	}
	var i int
	for _, line := range input {
		if line == "" {
			p.Tiles = append(p.Tiles, t)
			t = tile.Tile{
				Pixels: make(map[helpers.Coordinate]string),
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
			t.Pixels[helpers.Coordinate{X: j, Y: i}] = string(char)
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

func (p *Picture) PopulateTileMap() error {
	// First the top left tile, then build out from there
	var tile tile.Tile
	var x, y int
	for _, t := range p.Tiles {
		if t.AdjacentTiles.Bottom != "" && t.AdjacentTiles.Right != "" && t.AdjacentTiles.Left == "" && t.AdjacentTiles.Top == "" {
			tile = t
		}
		p.TileMap[helpers.Coordinate{X: x, Y: y}] = tile
	}
	for {
		if tile.AdjacentTiles.Bottom != "" {
			t, err := p.getTileFromID(tile.AdjacentTiles.Bottom)
			if err != nil {
				return err
			}
			y++
			p.TileMap[helpers.Coordinate{X: x, Y: y}] = t
			tile = t
		} else {
			tile = p.TileMap[helpers.Coordinate{X: x, Y: 0}]
			if tile.AdjacentTiles.Right == "" {
				return nil
			}
			t, err := p.getTileFromID(tile.AdjacentTiles.Right)
			if err != nil {
				return err
			}
			x++
			y = 0
			p.TileMap[helpers.Coordinate{X: x, Y: y}] = t
			tile = t
		}
	}
}
