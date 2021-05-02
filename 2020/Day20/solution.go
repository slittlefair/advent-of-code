package main

import (
	helpers "Advent-of-Code"
	"Advent-of-Code/2020/Day20/picture"
	"Advent-of-Code/2020/Day20/tile"
	"fmt"
)

// TODO:
// - Tests

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
	picture.FindSeaMonster()
	fmt.Println("Part 2:", picture.CountWaterRoughness())
}
