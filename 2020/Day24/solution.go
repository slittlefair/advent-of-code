package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
)

type Tiles struct {
	Map         map[helpers.Co]bool
	CurrentTile helpers.Co
	MaxX        int
	MaxY        int
	MinX        int
	MinY        int
}

func (t *Tiles) populateMissingTiles() {
	for x := t.MinX - 1; x <= t.MaxX+1; x++ {
		for y := t.MinY - 1; y <= t.MaxY+1; y++ {
			if _, ok := t.Map[helpers.Co{X: x, Y: y}]; !ok {
				t.Map[helpers.Co{X: x, Y: y}] = false
			}
		}
	}
}

func (t *Tiles) doFlipEachDay(days int) {
	for i := 0; i < days; i++ {
		t.populateMissingTiles()
		t.doFlips()
		fmt.Printf("Day %d: %d\n", i+1, t.countTiles())
	}
}

func (t *Tiles) doFlips() {
	tiles := t.decideWhichTilesToFlip()
	for _, co := range tiles {
		t.Map[co] = !t.Map[co]
	}
}

func (t *Tiles) shouldFlip(co helpers.Co) bool {
	count := 0
	if t.Map[helpers.Co{X: co.X + 1, Y: co.Y}] {
		count++
	}
	if t.Map[helpers.Co{X: co.X + 1, Y: co.Y + 1}] {
		count++
	}
	if t.Map[helpers.Co{X: co.X + 1, Y: co.Y - 1}] {
		count++
	}
	if t.Map[helpers.Co{X: co.X - 1, Y: co.Y}] {
		count++
	}
	if t.Map[helpers.Co{X: co.X - 1, Y: co.Y + 1}] {
		count++
	}
	if t.Map[helpers.Co{X: co.X - 1, Y: co.Y - 1}] {
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

func (t *Tiles) decideWhichTilesToFlip() []helpers.Co {
	tiles := []helpers.Co{}
	for co := range t.Map {
		if t.shouldFlip(co) {
			tiles = append(tiles, co)
		}
	}
	return tiles
}

func (t *Tiles) moveThroughList(tiles []string) {
	t.CurrentTile = helpers.Co{X: 0, Y: 0}
	for _, tile := range tiles {
		t.moveTile(tile)
		if t.CurrentTile.X < t.MinX {
			t.MinX = t.CurrentTile.X
		}
		if t.CurrentTile.Y < t.MinY {
			t.MinY = t.CurrentTile.Y
		}
		if t.CurrentTile.X > t.MaxX {
			t.MaxX = t.CurrentTile.X
		}
		if t.CurrentTile.Y > t.MaxY {
			t.MaxY = t.CurrentTile.Y
		}
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
	fmt.Println("Part 1:", tiles.countTiles())

	tiles.doFlipEachDay(10)
	// fmt.Println("Part 2:", tiles.countTiles())
}
