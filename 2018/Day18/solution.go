package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
)

type coordinate struct {
	X int
	Y int
}

var maxX, maxY = 0, 0

type forestTemplate map[coordinate]string

func populateForest(acres []string) (forest forestTemplate) {
	forest = make(forestTemplate)
	for y, row := range acres {
		for x, acre := range row {
			co := coordinate{X: x, Y: y}
			forest[co] = string(acre)
			maxX = x
			maxY = y
		}
	}
	return
}

func (forest forestTemplate) printForest() {
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			fmt.Printf(forest[coordinate{x, y}])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (forest forestTemplate) checkOpen(co coordinate) string {
	adjacentTrees := 0
	for y := co.Y - 1; y <= co.Y+1; y++ {
		for x := co.X - 1; x <= co.X+1; x++ {
			if val, ok := forest[coordinate{x, y}]; ok && val == "|" {
				adjacentTrees++
			}
		}
	}
	if adjacentTrees >= 3 {
		return "|"
	}
	return "."
}

func (forest forestTemplate) checkTrees(co coordinate) string {
	adjacentLumber := 0
	for y := co.Y - 1; y <= co.Y+1; y++ {
		for x := co.X - 1; x <= co.X+1; x++ {
			if val, ok := forest[coordinate{x, y}]; ok && val == "#" {
				adjacentLumber++
			}
		}
	}
	if adjacentLumber >= 3 {
		return "#"
	}
	return "|"
}

func (forest forestTemplate) checkLumber(co coordinate) string {
	adjacentTrees := 0
	adjacentLumber := 0
	for y := co.Y - 1; y <= co.Y+1; y++ {
		for x := co.X - 1; x <= co.X+1; x++ {
			if val, ok := forest[coordinate{x, y}]; ok && val == "|" {
				adjacentTrees++
			}
			if val, ok := forest[coordinate{x, y}]; ok && val == "#" {
				adjacentLumber++
			}
		}
	}
	if forest[co] == "#" {
		adjacentLumber--
	}
	if adjacentTrees >= 1 && adjacentLumber >= 1 {
		return "#"
	}
	return "."
}

func (forest forestTemplate) countTotal() {
	totalTrees := 0
	totalLumber := 0
	for _, acre := range forest {
		if acre == "|" {
			totalTrees++
		}
		if acre == "#" {
			totalLumber++
		}
	}
	fmt.Println(totalTrees * totalLumber)
}

func main() {
	forest := populateForest(utils.ReadFile())
	// forest.printForest()
	newForest := make(forestTemplate)
	iterations := 1000000000
	for i := 1; i < iterations; i++ {
		for y := 0; y <= maxY; y++ {
			for x := 0; x <= maxX; x++ {
				co := coordinate{x, y}
				switch forest[co] {
				case ".":
					newForest[co] = forest.checkOpen(co)
				case "|":
					newForest[co] = forest.checkTrees(co)
				case "#":
					newForest[co] = forest.checkLumber(co)
				}
			}
		}
		forest = newForest
		newForest = make(forestTemplate)
		mod := iterations % 56
		if i > 1500 && i%56 == mod {
			forest.printForest()
			forest.countTotal()
			return
		}
	}
	// forest.printForest()
	// forest.countTotal()
}
