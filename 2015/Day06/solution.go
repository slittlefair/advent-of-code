package main

import (
	"Advent-of-Code"
	"fmt"
	"regexp"
)

type coordinate struct {
	X int
	Y int
}

var part1Lights = make(map[coordinate]bool)
var part2Lights = make(map[coordinate]int)

func updateStatus(x int, y int, instruction string) {
	co := coordinate{x, y}
	switch instruction {
	case "on":
		part1Lights[co] = true
	case "off":
		part1Lights[co] = false
	case "through":
		part1Lights[co] = !part1Lights[co]
	}
}

func updateBrightness(x int, y int, instruction string) {
	co := coordinate{x, y}
	switch instruction {
	case "on":
		part2Lights[co] = part2Lights[co] + 1
	case "off":
		if part2Lights[co] > 0 {
			part2Lights[co] = part2Lights[co] - 1
		}
	case "through":
		part2Lights[co] = part2Lights[co] + 2
	}
}

func main() {
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			part1Lights[coordinate{x, y}] = false
			part2Lights[coordinate{x, y}] = 0
		}
	}

	instructions := helpers.ReadFile()
	coordsRe := regexp.MustCompile("\\d+")
	wordsRe := regexp.MustCompile("[a-z]+")

	for _, inst := range instructions {
		points := helpers.StringSliceToIntSlice(coordsRe.FindAllString(inst, -1))
		words := wordsRe.FindAllString(inst, -1)
		for x := points[0]; x <= points[2]; x++ {
			for y := points[1]; y <= points[3]; y++ {
				updateStatus(x, y, words[1])
				updateBrightness(x, y, words[1])
			}
		}
	}

	totalOn := 0
	totalBrightness := 0
	for co := range part1Lights {
		if part1Lights[co] {
			totalOn++
		}
		totalBrightness += part2Lights[co]
	}
	fmt.Println("Part 1:", totalOn)
	fmt.Println("Part 2:", totalBrightness)
}
