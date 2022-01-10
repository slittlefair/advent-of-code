package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"math"
)

type coordinate struct {
	X int
	Y int
}

type creature struct {
	letter string
	hp     int
	co     coordinate
}

type ac struct {
	elves   creaturesMap
	goblins creaturesMap
}

var allCreatures = ac{
	elves:   make(creaturesMap),
	goblins: make(creaturesMap),
}

var enemyReference = map[string]creaturesMap{
	"E": allCreatures.goblins,
	"G": allCreatures.elves,
}

type creaturesMap map[coordinate]creature

var maxRow, maxCol int

var cavePieces = make(map[coordinate]string)

func (c creature) nearestEnemy() (e creature) {
	enemies := enemyReference[c.letter]
	shortestDist := float64(10000000)
	var nearestCreature creature
	for i, v := range enemies {
		dist := math.Abs(float64(c.co.X-i.X)) + math.Abs(float64(c.co.Y-i.Y))
		if dist < shortestDist {
			shortestDist = dist
			nearestCreature = v
		}
	}
	// TODO check we don't pass through walls or other creatures
	// TODO use the first creature in reading order if there are two nearestEnemies
	return nearestCreature
}

func populateMap() {
	for j, row := range utils.ReadFile() {
		for i, col := range row {
			if i > maxRow {
				maxRow = i
			}
			if j > maxCol {
				maxCol = j
			}
			if string(col) == "E" {
				coord := coordinate{i, j}
				cavePieces[coord] = "."
				elf := creature{letter: "E", hp: 200, co: coord}
				allCreatures.elves[coord] = elf
			} else if string(col) == "G" {
				coord := coordinate{i, j}
				cavePieces[coord] = "."
				goblin := creature{letter: "G", hp: 200, co: coord}
				allCreatures.goblins[coord] = goblin
			}
			cavePieces[coordinate{i, j}] = string(col)
		}
	}
}

func printMap() {
	for j := 0; j <= maxRow; j++ {
		for i := 0; i <= maxCol; i++ {
			coord := coordinate{i, j}
			if e, ok := allCreatures.elves[coord]; ok {
				fmt.Print(e.letter)
			} else if g, ok := allCreatures.goblins[coord]; ok {
				fmt.Print(g.letter)
			} else {
				fmt.Print(cavePieces[coord])
			}
		}
		fmt.Println()
	}
}

func main() {
	populateMap()
	printMap()
	fmt.Println(cavePieces)
	fmt.Println(allCreatures.elves)
	fmt.Println(allCreatures.goblins)
	for k, e := range allCreatures.elves {
		fmt.Println(k, e)
	}
	for k, g := range allCreatures.goblins {
		fmt.Println(k, g)
	}
	fmt.Println(allCreatures.elves[coordinate{4, 2}].nearestEnemy())
}
