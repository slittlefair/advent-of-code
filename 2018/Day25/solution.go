package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/slice"
	"fmt"
	"math"
	"regexp"
)

type coordinate struct {
	X int
	Y int
	Z int
	T int
}

var allCoords = make(map[coordinate][]coordinate)
var coordInConstellation = make(map[coordinate]bool)

type constellation []coordinate

var allConstellations []constellation

func manhattan(startCo coordinate, endCo coordinate) int {
	xDist := math.Abs(float64(startCo.X - endCo.X))
	yDist := math.Abs(float64(startCo.Y - endCo.Y))
	zDist := math.Abs(float64(startCo.Z - endCo.Z))
	tDist := math.Abs(float64(startCo.T - endCo.T))
	return int(xDist + yDist + zDist + tDist)
}

func (co coordinate) addNearbyToConstellation(cons constellation) constellation {
	for _, nearby := range allCoords[co] {
		if _, ok := coordInConstellation[nearby]; !ok {
			coordInConstellation[nearby] = true
			cons = append(cons, nearby)
			cons = nearby.addNearbyToConstellation(cons)
		}
	}
	return cons
}

func main() {
	re := regexp.MustCompile(`-?\d+`)
	lines := file.Read()
	for _, line := range lines {
		points, err := slice.StringSliceToIntSlice(re.FindAllString(line, -1))
		if err != nil {
			fmt.Println(err)
			return
		}
		co := coordinate{
			X: points[0],
			Y: points[1],
			Z: points[2],
			T: points[3],
		}
		allCoords[co] = []coordinate{}
	}

	for startCo := range allCoords {
		nearby := []coordinate{}
		for endCo := range allCoords {
			if dist := manhattan(startCo, endCo); dist <= 3 && startCo != endCo {
				nearby = append(nearby, endCo)
			}
		}
		allCoords[startCo] = nearby
	}

	for co := range allCoords {
		if _, ok := coordInConstellation[co]; !ok {
			cons := constellation{co}
			cons = co.addNearbyToConstellation(cons)
			allConstellations = append(allConstellations, cons)
		}
	}

	fmt.Println("Part 1:", len(allConstellations))
}
