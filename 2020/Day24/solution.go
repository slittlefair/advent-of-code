package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
)

type Tiles struct {
	Map         map[helpers.Co]bool
	CurrentTile helpers.Co
}

func (t *Tiles) moveThroughList(tiles []string) {
	t.CurrentTile = helpers.Co{X: 0, Y: 0}
	for _, tile := range tiles {
		t.moveTile(tile)
	}
	t.flipTiles(t.CurrentTile)
}

func (t *Tiles) moveTile(dir string) {
	var changeX, changeY int
	switch dir {
	case "e":
		changeX = 1
	case "se":
		changeY = 1
		if t.CurrentTile.Y%2 != 0 {
			changeX = 1
		}
	case "ne":
		changeY = -1
		if t.CurrentTile.Y%2 != 0 {
			changeX = 1
		}
	case "w":
		changeX = -1
	case "sw":
		changeY = 1
		if t.CurrentTile.Y%2 == 0 {
			changeX = -1
		}
	case "nw":
		changeY = -1
		if t.CurrentTile.Y%2 == 0 {
			changeX = -1
		}
	}
	t.CurrentTile.X += changeX
	t.CurrentTile.Y += changeY
}

func (t *Tiles) flipTiles(co helpers.Co) {
	if val, ok := t.Map[co]; !ok {
		t.Map[co] = true
	} else {
		t.Map[co] = !val
	}
}

func parseInput(input []string) [][]string {
	re := regexp.MustCompile(`e|se|ne|w|sw|nw`)
	tileList := [][]string{}
	for _, line := range input {
		tileList = append(tileList, re.FindAllString(line, -1))
	}
	return tileList
}

func (t Tiles) countTiles() int {
	var black int
	for _, tile := range t.Map {
		if tile {
			black++
		}
	}
	return black
}

func main() {
	input := helpers.ReadFile()
	tileList := parseInput(input)
	tiles := &Tiles{
		Map: make(map[helpers.Co]bool),
	}
	for _, list := range tileList {
		tiles.moveThroughList(list)
	}
	blackTiles := tiles.countTiles()
	fmt.Println("Part 1:", blackTiles)
}
