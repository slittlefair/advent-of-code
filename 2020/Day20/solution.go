package main

import (
	helpers "Advent-of-Code"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Picture struct {
	pixels  map[helpers.Coordinate]string
	tileMap map[helpers.Coordinate]Tile
	tiles   []Tile
}

type AdjacentTiles struct {
	top    string
	bottom string
	left   string
	right  string
}

type Tile struct {
	id            string
	pixels        map[helpers.Coordinate]string
	height        int
	width         int
	adjacentTiles AdjacentTiles
}

// func (t Tile) printPixels() {
// 	for h := 0; h <= t.height; h++ {
// 		for w := 0; w <= t.width; w++ {
// 			fmt.Print(t.pixels[helpers.Coordinate{X: w, Y: h}])
// 		}
// 		fmt.Println()
// 	}
// }

func (t *Tile) rotateTile90() {
	newPixels := make(map[helpers.Coordinate]string)
	for co, val := range t.pixels {
		newPixels[helpers.Coordinate{X: t.width - co.Y, Y: co.X}] = val
	}
	t.pixels = newPixels
}

func (t *Tile) flipTile() {
	newPixels := make(map[helpers.Coordinate]string)
	for co, val := range t.pixels {
		newPixels[helpers.Coordinate{X: t.width - co.X, Y: co.Y}] = val
	}
	t.pixels = newPixels
}

func (t Tile) isAdjacentTop(tile Tile) bool {
	for x := 0; x <= t.width; x++ {
		if t.pixels[helpers.Coordinate{X: x, Y: t.height}] != tile.pixels[helpers.Coordinate{X: x, Y: 0}] {
			return false
		}
	}
	return true
}

func (t Tile) isAdjacentBottom(tile Tile) bool {
	for x := 0; x <= t.width; x++ {
		if t.pixels[helpers.Coordinate{X: x, Y: 0}] != tile.pixels[helpers.Coordinate{X: x, Y: t.height}] {
			return false
		}
	}
	return true
}

func (t Tile) isAdjacentLeft(tile Tile) bool {
	for y := 0; y <= t.height; y++ {
		if t.pixels[helpers.Coordinate{X: 0, Y: y}] != tile.pixels[helpers.Coordinate{X: tile.width, Y: y}] {
			return false
		}
	}
	return true
}

func (t Tile) isAdjacentRight(tile Tile) bool {
	for y := 0; y <= t.height; y++ {
		if t.pixels[helpers.Coordinate{X: t.width, Y: y}] != tile.pixels[helpers.Coordinate{X: 0, Y: y}] {
			return false
		}
	}
	return true
}

func (t Tile) isAdjacentTo(tile Tile) bool {
	if t.adjacentTiles.top == tile.id {
		return true
	}
	if t.adjacentTiles.bottom == tile.id {
		return true
	}
	if t.adjacentTiles.left == tile.id {
		return true
	}
	if t.adjacentTiles.right == tile.id {
		return true
	}
	return false
}

func (p Picture) findMatchesForTile(t Tile, index int) {
	for k, tile := range p.tiles {
		if tile.id == t.id || t.isAdjacentTo(tile) || t.numAdjacent() == 4 || tile.numAdjacent() == 4 {
			continue
		}
		for j := 0; j < 2; j++ {
			for i := 0; i < 4; i++ {
				if t.isAdjacentTop(tile) {
					p.tiles[k].adjacentTiles.bottom = t.id
					p.tiles[index].adjacentTiles.top = tile.id
					p.tiles[k].pixels = tile.pixels
					p.findMatchesForTile(tile, k)
					break
				}
				if t.isAdjacentBottom(tile) {
					p.tiles[k].adjacentTiles.top = t.id
					p.tiles[index].adjacentTiles.bottom = tile.id
					p.tiles[k].pixels = tile.pixels
					p.findMatchesForTile(tile, k)
					break
				}
				if t.isAdjacentLeft(tile) {
					p.tiles[k].adjacentTiles.right = t.id
					p.tiles[index].adjacentTiles.left = tile.id
					p.tiles[k].pixels = tile.pixels
					p.findMatchesForTile(tile, k)
					break
				}
				if t.isAdjacentRight(tile) {
					p.tiles[k].adjacentTiles.left = t.id
					p.tiles[index].adjacentTiles.right = tile.id
					p.tiles[k].pixels = tile.pixels
					p.findMatchesForTile(tile, k)
					break
				}
				tile.rotateTile90()
			}
			tile.flipTile()
		}
	}
}

func (p *Picture) populateTiles(input []string) {
	re := regexp.MustCompile(`\d+`)
	tile := Tile{
		pixels: make(map[helpers.Coordinate]string),
	}
	var i int
	for _, line := range input {
		if line == "" {
			p.tiles = append(p.tiles, tile)
			tile = Tile{
				pixels: make(map[helpers.Coordinate]string),
			}
			continue
		}
		if match := re.FindString(line); match != "" {
			tile.id = match
			i = 0
			continue
		}
		tile.height = i
		tile.width = len(line) - 1
		for j, char := range line {
			tile.pixels[helpers.Coordinate{X: j, Y: i}] = string(char)
		}
		i++
	}
	p.tiles = append(p.tiles, tile)
}

func (t Tile) numAdjacent() int {
	numAdjacent := 0
	if t.adjacentTiles.top != "" {
		numAdjacent++
	}
	if t.adjacentTiles.bottom != "" {
		numAdjacent++
	}
	if t.adjacentTiles.left != "" {
		numAdjacent++
	}
	if t.adjacentTiles.right != "" {
		numAdjacent++
	}
	return numAdjacent
}

func (p Picture) calculateCornerIDs() (int, error) {
	cornerID := 1
	for _, tile := range p.tiles {
		if tile.numAdjacent() == 2 {
			numericID, err := strconv.Atoi(tile.id)
			if err != nil {
				return 0, err
			}
			cornerID *= numericID
		}
	}
	return cornerID, nil
}

func (p Picture) getTileFromID(id string) (Tile, error) {
	for _, tile := range p.tiles {
		if tile.id == id {
			return tile, nil
		}
	}
	return Tile{}, errors.New(fmt.Sprintln("could not find tile for id:", id))
}

func (p *Picture) populateTileMap() error {
	// First the top left tile, then build out from there
	var tile Tile
	var x, y int
	for _, t := range p.tiles {
		if t.adjacentTiles.bottom != "" && t.adjacentTiles.right != "" && t.adjacentTiles.left == "" && t.adjacentTiles.top == "" {
			tile = t
		}
		p.tileMap[helpers.Coordinate{X: x, Y: y}] = tile
	}
	for {
		if tile.adjacentTiles.bottom != "" {
			t, err := p.getTileFromID(tile.adjacentTiles.bottom)
			if err != nil {
				return err
			}
			y++
			p.tileMap[helpers.Coordinate{X: x, Y: y}] = t
			tile = t
		} else {
			tile = p.tileMap[helpers.Coordinate{X: x, Y: 0}]
			if tile.adjacentTiles.right == "" {
				return nil
			}
			t, err := p.getTileFromID(tile.adjacentTiles.right)
			if err != nil {
				return err
			}
			x++
			y = 0
			p.tileMap[helpers.Coordinate{X: x, Y: y}] = t
			tile = t
		}
	}
}

func main() {
	input := helpers.ReadFile()
	picture := &Picture{
		tileMap: make(map[helpers.Coordinate]Tile),
	}
	picture.populateTiles(input)
	for i, tile := range picture.tiles {
		picture.findMatchesForTile(tile, i)
	}
	sol, err := picture.calculateCornerIDs()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", sol)
	err = picture.populateTileMap()
	if err != nil {
		fmt.Println(err)
		return
	}
	for co, tile := range picture.tileMap {
		fmt.Println(co, tile.id)
	}
}
