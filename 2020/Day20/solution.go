package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
	"strconv"
)

type Picture []Tile

type Tile struct {
	id            string
	pixels        map[helpers.Coordinate]string
	height        int
	width         int
	adjacentTiles map[string]bool
	edges         []string
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
		newPixels[helpers.Coordinate{X: int((float64(t.width) / 2) - (float64(co.Y) - (float64(t.height) / 2))), Y: int(float64(co.X) - float64(t.width)/2 + float64(t.height)/2)}] = val
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

func (t Tile) isMatch(tile Tile) bool {
	for _, edge1 := range t.edges {
		for _, edge2 := range tile.edges {
			if edge1 == edge2 {
				return true
			}
		}
	}
	return false
}

func (p Picture) findMatchesForTile(t Tile) {
	for _, tile := range p {
		if tile.id == t.id || t.adjacentTiles[tile.id] {
			continue
		}
		match := t.isMatch(tile)
		t.adjacentTiles[tile.id] = match
		tile.adjacentTiles[t.id] = match
	}
}

func (t *Tile) populateEdges() {
	edge := ""
	for x := 0; x <= t.width; x++ {
		edge += t.pixels[helpers.Coordinate{X: x, Y: 0}]
	}
	t.edges = append(t.edges, edge)

	edge = ""
	for x := 0; x <= t.width; x++ {
		edge += t.pixels[helpers.Coordinate{X: x, Y: t.height}]
	}
	t.edges = append(t.edges, edge)

	edge = ""
	for y := 0; y <= t.height; y++ {
		edge += t.pixels[helpers.Coordinate{X: 0, Y: y}]
	}
	t.edges = append(t.edges, edge)

	edge = ""
	for y := 0; y <= t.height; y++ {
		edge += t.pixels[helpers.Coordinate{X: t.height, Y: y}]
	}
	t.edges = append(t.edges, edge)

	edge = ""
	for x := 0; x <= t.width; x++ {
		edge += t.pixels[helpers.Coordinate{X: t.width - x, Y: 0}]
	}
	t.edges = append(t.edges, edge)

	edge = ""
	for x := 0; x <= t.width; x++ {
		edge += t.pixels[helpers.Coordinate{X: t.width - x, Y: t.height}]
	}
	t.edges = append(t.edges, edge)

	edge = ""
	for y := 0; y <= t.height; y++ {
		edge += t.pixels[helpers.Coordinate{X: 0, Y: t.height - y}]
	}
	t.edges = append(t.edges, edge)

	edge = ""
	for y := 0; y <= t.height; y++ {
		edge += t.pixels[helpers.Coordinate{X: t.height, Y: t.height - y}]
	}
	t.edges = append(t.edges, edge)
}

func (p *Picture) populateTiles(input []string) {
	re := regexp.MustCompile(`\d+`)
	tile := Tile{
		adjacentTiles: make(map[string]bool),
		pixels:        make(map[helpers.Coordinate]string),
	}
	var i int
	for _, line := range input {
		if line == "" {
			tile.populateEdges()
			*p = append(*p, tile)
			tile = Tile{
				adjacentTiles: make(map[string]bool),
				pixels:        make(map[helpers.Coordinate]string),
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
	tile.populateEdges()
	*p = append(*p, tile)
}

func (t Tile) numAdjacent() int {
	numAdjacent := 0
	for _, isAdj := range t.adjacentTiles {
		if isAdj {
			numAdjacent++
		}
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
	for _, tile := range *picture {
		picture.findMatchesForTile(tile)
	}
	sol, err := picture.calculateCornerIDs()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", sol)
}
