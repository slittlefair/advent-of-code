package main

import (
	helpers "Advent-of-Code"
	"fmt"
)

type Lights map[helpers.Co]string

type Grid struct {
	Lights Lights
	Height int
	Width  int
}

func parseInput(input []string) Grid {
	grid := Grid{
		Lights: make(Lights),
		Height: len(input) - 1,
	}
	for y, row := range input {
		grid.Width = len(row) - 1
		for x, col := range row {
			grid.Lights[helpers.Co{X: x, Y: y}] = string(col)
		}
	}
	return grid
}

// PrintLights for debugging
// func (g Grid) PrintLights() {
// 	for y := 0; y <= g.Height; y++ {
// 		for x := 0; x <= g.Width; x++ {
// 			fmt.Printf(g.Lights[helpers.Co{X: x, Y: y}])
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println()
// }

func (g *Grid) LightStaysOn(co helpers.Co) bool {
	count := 0
	for _, adjCo := range helpers.AdjacentCos(co, true) {
		if g.Lights[adjCo] == "#" {
			count++
		}
	}
	if g.Lights[co] == "#" && (count == 2 || count == 3) {
		return true
	}
	if g.Lights[co] == "." && count == 3 {
		return true
	}
	return false
}

func (g *Grid) ChangeLights() {
	newLights := make(Lights)
	for co := range g.Lights {
		if g.LightStaysOn(co) {
			newLights[co] = "#"
		} else {
			newLights[co] = "."
		}
	}
	g.Lights = newLights
}

func (g Grid) CountLightsOn() int {
	count := 0
	for _, val := range g.Lights {
		if val == "#" {
			count++
		}
	}
	return count
}

func (g *Grid) TurnCornersOn() {
	g.Lights[helpers.Co{X: 0, Y: 0}] = "#"
	g.Lights[helpers.Co{X: 0, Y: g.Height}] = "#"
	g.Lights[helpers.Co{X: g.Width, Y: 0}] = "#"
	g.Lights[helpers.Co{X: g.Width, Y: g.Height}] = "#"
}

func (g *Grid) RunStepsPart1(steps int) {
	for i := 0; i < steps; i++ {
		g.ChangeLights()
	}
}

func (g *Grid) RunStepsPart2(steps int) {
	g.TurnCornersOn()
	for i := 0; i < steps; i++ {
		g.ChangeLights()
		g.TurnCornersOn()
	}
}

func runAndCountLightsPart1(input []string, steps int) int {
	grid := parseInput(input)
	grid.RunStepsPart1(steps)
	return grid.CountLightsOn()
}

func runAndCountLightsPart2(input []string, steps int) int {
	grid := parseInput(input)
	grid.RunStepsPart2(steps)
	return grid.CountLightsOn()
}

func main() {
	input := helpers.ReadFile()
	fmt.Println("Part 1:", runAndCountLightsPart1(input, 100))
	fmt.Println("Part 2:", runAndCountLightsPart2(input, 100))
}
