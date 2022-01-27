package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type co struct {
	X float64
	Y float64
}

var m = make(map[co]int)
var coords []co

var inf = make(map[co]bool)

var re = regexp.MustCompile(`\d+`)

var maxX, maxY float64

var safeRegions int
var maxSafe = float64(10000)

func main() {
	coordinates := utils.ReadFile()
	for _, val := range coordinates {
		match := re.FindAllString(val, -1)
		var ints []float64
		for _, num := range match {
			i, err := strconv.ParseFloat(num, 64)
			utils.Check(err)
			ints = append(ints, i)
		}
		coords = append(coords, co{ints[0], ints[1]})
		if ints[0] > maxX {
			maxX = ints[0]
		}
		if ints[1] > maxY {
			maxY = ints[1]
		}
	}

	for y := float64(0); y < maxY; y++ {
		for x := float64(0); x < maxX; x++ {
			var mc = co{}
			var minDist = float64(1000000)
			var tot float64
			for _, c := range coords {
				dist := math.Abs(x-c.X) + math.Abs(y-c.Y)
				if dist < minDist {
					minDist = dist
					mc = c
				} else if dist == minDist {
					mc = co{-1, -1}
				}
				tot += dist
			}

			// Part 1
			if x == 0 || y == 0 || x == maxX || y == maxY {
				inf[mc] = true
			}
			m[mc]++

			// Part 2
			if tot < maxSafe {
				safeRegions++
			}
		}
	}

	maxArea := 0
	for key, val := range m {
		if _, ok := inf[key]; val > maxArea && !ok {
			maxArea = val
		}
	}
	fmt.Println("Part 1:", maxArea)
	fmt.Println("Part 2:", safeRegions)
}
