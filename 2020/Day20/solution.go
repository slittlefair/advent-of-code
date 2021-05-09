package main

import (
	helpers "Advent-of-Code"
	"Advent-of-Code/2020/Day20/picture"
	"Advent-of-Code/2020/Day20/tile"
	"fmt"
)

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

func main() {
	input := helpers.ReadFile()
	picture := &picture.Picture{
		Pixels:  make(map[helpers.Co]string),
		TileMap: make(map[helpers.Co]tile.Tile),
	}
	picture.PopulateTiles(input)
	for i, tile := range picture.Tiles {
		picture.FindMatchesForTile(tile, i)
	}
	sol, err := picture.CalculateCornerIDs()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", sol)
	err = picture.PopulateTileMap()
	if err != nil {
		fmt.Println(err)
		return
	}
	picture.FindSeaMonster(seaMonster)
	fmt.Println("Part 2:", picture.CountWaterRoughness())
}
