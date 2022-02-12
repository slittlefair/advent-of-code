package main

import (
	"Advent-of-Code/file"
	"fmt"
)

type coordinate struct {
	X int
	Y int
}

type info struct {
	gate    string
	offsets [8]int
}

func directionInts(s string) info {
	switch s {
	case "N":
		return info{
			gate:    "-",
			offsets: [8]int{0, -2, 0, -1, -1, -1, 1, -1},
		}
	case "S":
		return info{
			gate:    "-",
			offsets: [8]int{0, 2, 0, 1, -1, 1, -1, 1},
		}
	case "E":
		return info{
			gate:    "|",
			offsets: [8]int{2, 0, 1, 0, 1, 1, -1, 1},
		}
	case "W":
		return info{
			gate:    "|",
			offsets: [8]int{-2, 0, -1, 0, -1, 1, -1, -1},
		}
	}
	panic("PANIC!!!")
}

var savedCoords []coordinate

func (f fullMap) fillInMap(l string, co coordinate) coordinate {
	if l == "N" || l == "S" || l == "E" || l == "W" {
		info := directionInts(l)
		o := info.offsets
		distanceTotal++
		if val, ok := dm[coordinate{co.X + o[0], co.Y + o[1]}]; !ok {
			dm[coordinate{co.X + o[0], co.Y + o[1]}] = distanceTotal
		} else {
			distanceTotal = val
		}
		f[coordinate{co.X + o[0], co.Y + o[1]}] = "."
		f[coordinate{co.X + o[2], co.Y + o[3]}] = info.gate
		f[coordinate{co.X + o[4], co.Y + o[5]}] = "#"
		f[coordinate{co.X + o[6], co.Y + o[7]}] = "#"
		checkLimits(coordinate{co.X + o[0], co.Y + o[1]})
		return coordinate{co.X + o[0], co.Y + o[1]}
	} else if l == "(" {
		savedCoords = append(savedCoords, co)
		return co
	} else if l == "|" {
		distanceTotal = dm[savedCoords[len(savedCoords)-1]]
		return savedCoords[len(savedCoords)-1]
	} else if l == ")" {
		savedCoords = savedCoords[:len(savedCoords)-1]
		return co
	}
	panic(fmt.Sprintln("s=", l))
}

var (
	minX int
	maxX int
	minY int
	maxY int
)

type fullMap map[coordinate]string

var f = fullMap{coordinate{0, 0}: "X"}

type distanceMap map[coordinate]int

var dm = distanceMap{coordinate{0, 0}: 0}

var distanceTotal int

func (f fullMap) printMap() {
	for y := minY - 1; y <= maxY+1; y++ {
		for x := minX - 1; x <= maxX+1; x++ {
			if val, ok := f[coordinate{x, y}]; !ok {
				fmt.Printf("#")
			} else {
				fmt.Print(val)
			}
		}
		fmt.Println()
	}
}

func checkLimits(co coordinate) {
	if co.X < minX {
		minX = co.X
	}
	if co.X > maxX {
		maxX = co.X
	}
	if co.Y < minY {
		minY = co.Y
	}
	if co.Y > maxY {
		maxY = co.Y
	}
}

// for debugging
// func printLimits() {
// 	fmt.Println("minX:", minX)
// 	fmt.Println("maxX:", maxX)
// 	fmt.Println("minY:", minY)
// 	fmt.Println("maxY:", maxY)
// }

func (dm distanceMap) greatestDistance() {
	greatestDistance := 0
	for _, val := range dm {
		if val > greatestDistance {
			greatestDistance = val
		}
	}
	fmt.Println("Part A:", greatestDistance)
}

func (dm distanceMap) distancesOver1000() {
	rooms := 0
	for _, val := range dm {
		if val >= 1000 {
			rooms++
		}
	}
	fmt.Println("Part B:", rooms)
}

func main() {
	str := file.Read()
	co := coordinate{0, 0}
	for i := 1; i < len(str[0])-1; i++ {
		co = f.fillInMap(string(str[0][i]), co)
	}
	fmt.Println(co)
	fmt.Println()
	f.printMap()
	dm.greatestDistance()
	dm.distancesOver1000()
}
