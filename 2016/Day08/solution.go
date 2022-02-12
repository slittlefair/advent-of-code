package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Lights struct {
	Pixels map[graph.Co]string
	Height int
	Width  int
}

func constructLights(height, width int) *Lights {
	pixels := make(map[graph.Co]string)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pixels[graph.Co{X: x, Y: y}] = " "
		}
	}
	return &Lights{
		Pixels: pixels,
		Height: height,
		Width:  width,
	}
}

func (l *Lights) followInstruction(inst string) error {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(inst, -1)
	if l := len(matches); l != 2 {
		return fmt.Errorf("expected 2 numbers, got %d from %s", l, inst)
	}
	// We can ignore these errors since we know they can be converted due to regex match.
	n1, _ := strconv.Atoi(matches[0])
	n2, _ := strconv.Atoi(matches[1])
	if strings.Contains(inst, "rect") {
		for x := 0; x < n1; x++ {
			for y := 0; y < n2; y++ {
				l.Pixels[graph.Co{X: x, Y: y}] = "#"
			}
		}
		return nil
	}
	pixels := map[graph.Co]string{}
	for k, v := range l.Pixels {
		pixels[k] = v
	}
	if strings.Contains(inst, "column") {
		for y := 0; y < l.Height; y++ {
			pixels[graph.Co{X: n1, Y: (y + n2) % l.Height}] = l.Pixels[graph.Co{X: n1, Y: y}]
		}
		l.Pixels = pixels
		return nil
	}
	for x := 0; x < l.Width; x++ {
		pixels[graph.Co{X: (x + n2) % l.Width, Y: n1}] = l.Pixels[graph.Co{X: x, Y: n1}]
	}
	l.Pixels = pixels
	return nil
}

func (l *Lights) followInstructions(input []string) error {
	for _, inst := range input {
		err := l.followInstruction(inst)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l Lights) countLightsOn() int {
	count := 0
	for _, v := range l.Pixels {
		if v == "#" {
			count++
		}
	}
	return count
}

// debugging
// func (l Lights) printLights() {
// 	for y := 0; y < l.Height; y++ {
// 		for x := 0; x < l.Width; x++ {
// 			fmt.Print(l.Pixels[graph.Co{X: x, Y: y}])
// 		}
// 		fmt.Println()
// 	}
// }

func main() {
	input := file.Read()
	lights := constructLights(6, 50)
	err := lights.followInstructions(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", lights.countLightsOn())
	fmt.Println("Part 2:")
}
