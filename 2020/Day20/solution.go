package main

import (
	helpers "Advent-of-Code"
	"Advent-of-Code/2020/Day20/picture"
	"Advent-of-Code/2020/Day20/tile"
	"fmt"
)

func main() {
	input := helpers.ReadFile()
	picture := &picture.Picture{
		Pixels:  make(map[helpers.Coordinate]string),
		TileMap: make(map[helpers.Coordinate]tile.Tile),
	}
	picture.PopulateTiles(input)
	for _, tile := range picture.Tiles {
		fmt.Println(tile.ID)
		tile.PrintPixels()
	}
	for i, tile := range picture.Tiles {
		picture.FindMatchesForTile(tile, i)
	}
	for _, tile := range picture.Tiles {
		fmt.Println(tile.ID)
		tile.PrintPixels()
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
	picture.PrintPictureMap()
}
