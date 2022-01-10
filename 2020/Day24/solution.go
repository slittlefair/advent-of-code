package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"regexp"
)

type Tiles struct {
	Map         map[utils.Co]bool
	CurrentTile utils.Co
	MaxX        int
	MaxY        int
	MinX        int
	MinY        int
}

func parseInput(input []string) [][]string {
	re := regexp.MustCompile(`e|se|ne|w|sw|nw`)
	tileList := [][]string{}
	for _, line := range input {
		tileList = append(tileList, re.FindAllString(line, -1))
	}
	return tileList
}

func (t *Tiles) getETile(co utils.Co) (utils.Co, bool) {
	newCo := utils.Co{X: co.X + 1, Y: co.Y}
	return newCo, t.Map[newCo]
}

func (t *Tiles) getSETile(co utils.Co) (utils.Co, bool) {
	changeX := 0
	if co.Y%2 != 0 {
		changeX = 1
	}
	newCo := utils.Co{X: co.X + changeX, Y: co.Y + 1}
	return newCo, t.Map[newCo]
}

func (t *Tiles) getNETile(co utils.Co) (utils.Co, bool) {
	changeX := 0
	if co.Y%2 != 0 {
		changeX = 1
	}
	newCo := utils.Co{X: co.X + changeX, Y: co.Y - 1}
	return newCo, t.Map[newCo]
}

func (t *Tiles) getWTile(co utils.Co) (utils.Co, bool) {
	newCo := utils.Co{X: co.X - 1, Y: co.Y}
	return newCo, t.Map[newCo]
}

func (t *Tiles) getSWTile(co utils.Co) (utils.Co, bool) {
	changeX := 0
	if co.Y%2 == 0 {
		changeX = -1
	}
	newCo := utils.Co{X: co.X + changeX, Y: co.Y + 1}
	return newCo, t.Map[newCo]
}

func (t *Tiles) getNWTile(co utils.Co) (utils.Co, bool) {
	changeX := 0
	if co.Y%2 == 0 {
		changeX = -1
	}
	newCo := utils.Co{X: co.X + changeX, Y: co.Y - 1}
	return newCo, t.Map[newCo]
}

func (t *Tiles) moveTile(dir string) {
	switch dir {
	case "e":
		t.CurrentTile, _ = t.getETile(t.CurrentTile)
	case "se":
		t.CurrentTile, _ = t.getSETile(t.CurrentTile)
	case "ne":
		t.CurrentTile, _ = t.getNETile(t.CurrentTile)
	case "w":
		t.CurrentTile, _ = t.getWTile(t.CurrentTile)
	case "sw":
		t.CurrentTile, _ = t.getSWTile(t.CurrentTile)
	case "nw":
		t.CurrentTile, _ = t.getNWTile(t.CurrentTile)
	}
	if t.CurrentTile.X < t.MinX {
		t.MinX = t.CurrentTile.X
	}
	if t.CurrentTile.X > t.MaxX {
		t.MaxX = t.CurrentTile.X
	}
	if t.CurrentTile.Y < t.MinY {
		t.MinY = t.CurrentTile.Y
	}
	if t.CurrentTile.Y > t.MaxY {
		t.MaxY = t.CurrentTile.Y
	}
}

func (t *Tiles) flipTiles(co utils.Co) {
	if val, ok := t.Map[co]; !ok {
		t.Map[co] = true
	} else {
		t.Map[co] = !val
	}
}

func (t *Tiles) moveThroughList(tiles []string) {
	t.CurrentTile = utils.Co{X: 0, Y: 0}
	for _, tile := range tiles {
		t.moveTile(tile)
	}
	t.flipTiles(t.CurrentTile)
}

func (t Tiles) countTiles() int {
	var count int
	for _, tile := range t.Map {
		if tile {
			count++
		}
	}
	return count
}

func (t *Tiles) populateMissingTiles() {
	t.MinX--
	t.MaxX++
	t.MinY--
	t.MaxY++
	for x := t.MinX; x <= t.MaxX; x++ {
		for y := t.MinY; y <= t.MaxY; y++ {
			if _, ok := t.Map[utils.Co{X: x, Y: y}]; !ok {
				t.Map[utils.Co{X: x, Y: y}] = false
			}
		}
	}
}

func (t *Tiles) shouldFlip(co utils.Co) bool {
	count := 0
	if _, blackTile := t.getETile(co); blackTile {
		count++
	}
	if _, blackTile := t.getSETile(co); blackTile {
		count++
	}
	if _, blackTile := t.getNETile(co); blackTile {
		count++
	}
	if _, blackTile := t.getWTile(co); blackTile {
		count++
	}
	if _, blackTile := t.getSWTile(co); blackTile {
		count++
	}
	if _, blackTile := t.getNWTile(co); blackTile {
		count++
	}
	if t.Map[co] && (count == 0 || count > 2) {
		return true
	}
	if !t.Map[co] && count == 2 {
		return true
	}
	return false
}

func (t *Tiles) decideWhichTilesToFlip() []utils.Co {
	tiles := []utils.Co{}
	for co := range t.Map {
		if t.shouldFlip(co) {
			tiles = append(tiles, co)
		}
	}
	return tiles
}

func (t *Tiles) doFlips() {
	tiles := t.decideWhichTilesToFlip()
	for _, co := range tiles {
		t.Map[co] = !t.Map[co]
	}
}

func (t *Tiles) countTilesAfterDays(days int) int {
	for i := 1; i <= days; i++ {
		t.populateMissingTiles()
		t.doFlips()
	}
	return t.countTiles()
}

func main() {
	input := utils.ReadFile()
	tileList := parseInput(input)
	tiles := &Tiles{
		Map: make(map[utils.Co]bool),
	}
	for _, list := range tileList {
		tiles.moveThroughList(list)
	}
	fmt.Println("Part 1:", tiles.countTiles())
	fmt.Println("Part 2:", tiles.countTilesAfterDays(100))
}
