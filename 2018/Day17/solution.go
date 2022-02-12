package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"fmt"
	"regexp"
	"strconv"
)

var allCoords = map[graph.Co]string{{X: 500, Y: 0}: "+"}

var (
	minX = 10000000
	maxX = 0
	minY = 0
	maxY = 0
)

var boundaryCheck = make(map[graph.Co]bool)

func updateBoundaries(co graph.Co) {
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

func populateRocks(lines []string) error {
	re, err := regexp.Compile(`[xy]|\d+`)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, line := range lines {
		match := re.FindAllString(line, -1)
		if len(match) == 6 {
			return fmt.Errorf("got six matches for line %s", line)
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
				min, err := strconv.Atoi(match[3])
				if err != nil {
					return err
				}
				max, err := strconv.Atoi(match[4])
				if err != nil {
					return err
				}
				for y := min; y <= max; y++ {
					x, err := strconv.Atoi(match[1])
					if err != nil {
						return err
					}
					co := graph.Co{X: x, Y: y}
					allCoords[co] = "#"
					updateBoundaries(co)
				}
			} else if yIndex == 3 {
				min, err := strconv.Atoi(match[1])
				if err != nil {
					return err
				}
				max, err := strconv.Atoi(match[2])
				if err != nil {
					return err
				}
				for x := min; x <= max; x++ {
					y, err := strconv.Atoi(match[4])
					if err != nil {
						return err
					}
					co := graph.Co{X: x, Y: y}
					allCoords[co] = "#"
					updateBoundaries(co)
				}
			} else if xIndex == 2 {
				min, err := strconv.Atoi(match[3])
				if err != nil {
					return err
				}
				max, err := strconv.Atoi(match[4])
				if err != nil {
					return err
				}
				for x := min; x <= max; x++ {
					y, err := strconv.Atoi(match[1])
					if err != nil {
						return err
					}
					co := graph.Co{X: x, Y: y}
					allCoords[co] = "#"
					updateBoundaries(co)
				}
			} else if xIndex == 3 {
				min, err := strconv.Atoi(match[1])
				if err != nil {
					return err
				}
				max, err := strconv.Atoi(match[2])
				if err != nil {
					return err
				}
				for y := min; y <= max; y++ {
					x, err := strconv.Atoi(match[4])
					if err != nil {
						return err
					}
					co := graph.Co{X: x, Y: y}
					allCoords[co] = "#"
					updateBoundaries(co)
				}
			} else {
				return fmt.Errorf("something went wrong, line : %s", line)
			}
		} else {
			if match[0] == "x" {
				x, err := strconv.Atoi(match[1])
				if err != nil {
					return err
				}
				y, err := strconv.Atoi(match[3])
				if err != nil {
					return err
				}
				co := graph.Co{X: x, Y: y}
				allCoords[co] = "#"
				updateBoundaries(co)
			} else {
				x, err := strconv.Atoi(match[3])
				if err != nil {
					return err
				}
				y, err := strconv.Atoi(match[1])
				if err != nil {
					return err
				}
				co := graph.Co{X: x, Y: y}
				allCoords[co] = "#"
				updateBoundaries(co)
			}
		}
	}
	return fmt.Errorf("something went wrong for line: %v", lines)
}

func populateSand() {
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			co := graph.Co{X: x, Y: y}
			if _, ok := allCoords[co]; !ok {
				allCoords[co] = "."
			}
		}
	}
}

func printLandscape(part1 bool) {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			co := graph.Co{X: x, Y: y}
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

func nextWater(c []graph.Co) (newCoords []graph.Co) {
	for _, co := range c {
		if _, ok := boundaryCheck[co]; ok {
			newCo := graph.Co{X: co.X, Y: co.Y + 1}
			if allCoords[co] == "~" {
				newCoords = append(newCoords, graph.Co{X: co.X, Y: co.Y - 1})
			} else if allCoords[newCo] == "|" {
			} else if allCoords[newCo] == "." {
				allCoords[newCo] = "|"
				newCoords = append(newCoords, newCo)
			} else {
				layer := []graph.Co{{X: co.X, Y: co.Y}}
				fullLayer := true
				offsets := [2]int{-1, 1}
				for _, direction := range offsets {
					traversing := true
					directionCo := graph.Co{X: co.X + direction, Y: co.Y}
					for traversing {
						if (allCoords[directionCo] == "." || allCoords[directionCo] == "|") && allCoords[graph.Co{X: directionCo.X, Y: directionCo.Y + 1}] != "." {
							layer = append(layer, directionCo)
							directionCo = graph.Co{X: directionCo.X + direction, Y: directionCo.Y}
						} else if (allCoords[directionCo] == "." || allCoords[directionCo] == "|") && allCoords[graph.Co{X: directionCo.X, Y: directionCo.Y + 1}] == "." {
							if allCoords[graph.Co{X: directionCo.X - direction, Y: directionCo.Y + 1}] == "|" {
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
					newCoords = append(newCoords, graph.Co{X: co.X, Y: co.Y - 1})
				} else {
					for _, cell := range layer {
						allCoords[cell] = "|"
					}
				}
			}
		}
	}
	boundaryCheck = make(map[graph.Co]bool)
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
	lines := file.Read()
	if err := populateRocks(lines); err != nil {
		fmt.Println(err)
		return
	}
	populateSand()
	co := []graph.Co{{X: 500, Y: 0}}
	boundaryCheck[graph.Co{X: 500, Y: 0}] = false
	for {
		if len(boundaryCheck) == 0 {
			calculateTotal()
			return
		}
		co = nextWater(co)
	}
}
