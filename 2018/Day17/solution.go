package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"regexp"
)

var allCoords = map[utils.Co]string{{X: 500, Y: 0}: "+"}

var (
	minX = 10000000
	maxX = 0
	minY = 0
	maxY = 0
)

var boundaryCheck = make(map[utils.Co]bool)

func updateBoundaries(co utils.Co) {
	if co.X < minX {
		minX = co.X
	}
	if co.X > maxX {
		maxX = co.X + 10
	}
	if co.Y > maxY {
		maxY = co.Y
	}
}

func populateRocks(lines []string) {
	re, err := regexp.Compile(`[xy]|\d+`)
	utils.Check(err)
	for _, line := range lines {
		match := re.FindAllString(line, -1)
		if len(match) == 6 {
			fmt.Println("PANIC!!!!", line)
		} else if len(match) == 5 {
			xIndex, yIndex := 0, 0
			for i, v := range match {
				if v == "x" {
					xIndex = i
				} else if v == "y" {
					yIndex = i
				}
			}
			if yIndex == 2 {
				for y := utils.StringToInt(match[3]); y <= utils.StringToInt(match[4]); y++ {
					co := utils.Co{X: utils.StringToInt(match[1]), Y: y}
					allCoords[co] = "#"
					updateBoundaries(co)
				}
			} else if yIndex == 3 {
				for x := utils.StringToInt(match[1]); x <= utils.StringToInt(match[2]); x++ {
					co := utils.Co{X: x, Y: utils.StringToInt(match[4])}
					allCoords[co] = "#"
					updateBoundaries(co)
				}
			} else if xIndex == 2 {
				for x := utils.StringToInt(match[3]); x <= utils.StringToInt(match[4]); x++ {
					co := utils.Co{X: x, Y: utils.StringToInt(match[1])}
					allCoords[co] = "#"
					updateBoundaries(co)
				}
			} else if xIndex == 3 {
				for y := utils.StringToInt(match[1]); y <= utils.StringToInt(match[2]); y++ {
					co := utils.Co{X: utils.StringToInt(match[4]), Y: y}
					allCoords[co] = "#"
					updateBoundaries(co)
				}
			} else {
				fmt.Println("PANIC!!!!!!", line)
			}
		} else {
			if match[0] == "x" {
				co := utils.Co{X: utils.StringToInt(match[1]), Y: utils.StringToInt(match[3])}
				allCoords[co] = "#"
				updateBoundaries(co)
			} else {
				co := utils.Co{X: utils.StringToInt(match[3]), Y: utils.StringToInt(match[1])}
				allCoords[co] = "#"
				updateBoundaries(co)
			}
		}
	}
}

func populateSand() {
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			co := utils.Co{X: x, Y: y}
			if _, ok := allCoords[co]; !ok {
				allCoords[co] = "."
			}
		}
	}
}

func printLandscape(part1 bool) {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			co := utils.Co{X: x, Y: y}
			if allCoords[co] == "." {
				fmt.Printf(" ")
			} else if part1 {
				fmt.Printf(allCoords[co])
			} else {
				if allCoords[co] == "|" {
					fmt.Printf(" ")
				} else {
					fmt.Printf(allCoords[co])
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func nextWater(c []utils.Co) (newCoords []utils.Co) {
	for _, co := range c {
		if _, ok := boundaryCheck[co]; ok {
			newCo := utils.Co{X: co.X, Y: co.Y + 1}
			if allCoords[co] == "~" {
				newCoords = append(newCoords, utils.Co{X: co.X, Y: co.Y - 1})
			} else if allCoords[newCo] == "|" {
			} else if allCoords[newCo] == "." {
				allCoords[newCo] = "|"
				newCoords = append(newCoords, newCo)
			} else {
				layer := []utils.Co{{X: co.X, Y: co.Y}}
				fullLayer := true
				offsets := [2]int{-1, 1}
				for _, direction := range offsets {
					traversing := true
					directionCo := utils.Co{X: co.X + direction, Y: co.Y}
					for traversing {
						if (allCoords[directionCo] == "." || allCoords[directionCo] == "|") && allCoords[utils.Co{X: directionCo.X, Y: directionCo.Y + 1}] != "." {
							layer = append(layer, directionCo)
							directionCo = utils.Co{X: directionCo.X + direction, Y: directionCo.Y}
						} else if (allCoords[directionCo] == "." || allCoords[directionCo] == "|") && allCoords[utils.Co{X: directionCo.X, Y: directionCo.Y + 1}] == "." {
							if allCoords[utils.Co{X: directionCo.X - direction, Y: directionCo.Y + 1}] == "|" {
								fullLayer = false
								traversing = false
							} else {
								layer = append(layer, directionCo)
								fullLayer = false
								traversing = false
								newCoords = append(newCoords, directionCo)
							}
						} else {
							traversing = false
						}
					}
				}
				if fullLayer {
					for _, cell := range layer {
						allCoords[cell] = "~"
					}
					newCoords = append(newCoords, utils.Co{X: co.X, Y: co.Y - 1})
				} else {
					for _, cell := range layer {
						allCoords[cell] = "|"
					}
				}
			}
		}
	}
	boundaryCheck = make(map[utils.Co]bool)
	for _, c := range newCoords {
		if c.Y != maxY && c.Y >= minY {
			boundaryCheck[c] = false
		}
	}
	// fmt.Println(newCoords)
	return
}

func calculateTotal() {
	total := 0
	for co, val := range allCoords {
		if val == "~" || val == "|" {
			if co.Y != maxY {
				total++
			}
		}
	}
	printLandscape(false)
	fmt.Println("Part1:", total)
	for co, val := range allCoords {
		if val == "|" {
			if co.Y != maxY {
				total--
			}
		}
	}
	fmt.Println("Part2:", total)
}

func main() {
	lines := utils.ReadFile()
	populateRocks(lines)
	populateSand()
	co := []utils.Co{{X: 500, Y: 0}}
	boundaryCheck[utils.Co{X: 500, Y: 0}] = false
	for {
		if len(boundaryCheck) == 0 {
			calculateTotal()
			return
		}
		co = nextWater(co)
	}
}
