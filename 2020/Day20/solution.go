package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
	"strconv"
)

type Picture []Tile

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
	for k, tile := range p {
		if tile.id == t.id || t.isAdjacentTo(tile) || t.numAdjacent() == 4 || tile.numAdjacent() == 4 {
			continue
		}
		for j := 0; j < 2; j++ {
			for i := 0; i < 4; i++ {
				if t.isAdjacentTop(tile) {
					p[k].adjacentTiles.bottom = t.id
					p[index].adjacentTiles.top = tile.id
					p[k].pixels = tile.pixels
					p.findMatchesForTile(tile, k)
					break
				}
				if t.isAdjacentBottom(tile) {
					p[k].adjacentTiles.top = t.id
					p[index].adjacentTiles.bottom = tile.id
					p[k].pixels = tile.pixels
					p.findMatchesForTile(tile, k)
					break
				}
				if t.isAdjacentLeft(tile) {
					p[k].adjacentTiles.right = t.id
					p[index].adjacentTiles.left = tile.id
					p[k].pixels = tile.pixels
					p.findMatchesForTile(tile, k)
					break
				}
				if t.isAdjacentRight(tile) {
					p[k].adjacentTiles.left = t.id
					p[index].adjacentTiles.right = tile.id
					p[k].pixels = tile.pixels
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
			*p = append(*p, tile)
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
	*p = append(*p, tile)
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
	for _, tile := range p {
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

func main() {
	input := helpers.ReadFile()
	picture := &Picture{}
	picture.populateTiles(input)
	for i, tile := range *picture {
		picture.findMatchesForTile(tile, i)
	}
	sol, err := picture.calculateCornerIDs()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", sol)
}
