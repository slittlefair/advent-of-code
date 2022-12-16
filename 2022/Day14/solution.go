package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"Advent-of-Code/maths"
	"fmt"
	"strings"
	"time"
)

type Grid map[graph.Co]string

type Waterfall struct {
	grid                   Grid
	minX, minY, maxX, maxY int
}

func (w Waterfall) PrintGrid(overwrite bool, slowMo bool) {
	if overwrite {
		if slowMo {
			time.Sleep(100 * time.Millisecond)
		} else {
			time.Sleep(30 * time.Millisecond)
		}
		for i := w.minY; i <= w.maxY+1; i++ {
			fmt.Printf("\033[A")
		}
	}

	for y := w.minY; y <= w.maxY+1; y++ {
		for x := w.minX - 1; x <= w.maxX+1; x++ {
			if x == 500 && y == 0 {
				fmt.Print("+")
			} else if v, ok := w.grid[graph.Co{X: x, Y: y}]; ok {
				fmt.Print(v)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func parseInput(input []string) (*Waterfall, error) {
	// We know sand comes from {X: 500, Y: 0}, so minY will always be 0
	w := &Waterfall{
		grid: Grid{},
		minX: maths.Infinity,
		maxX: 0,
		maxY: 0,
	}
	for _, line := range input {
		corners := strings.Split(line, " -> ")
		for i := 0; i < len(corners)-1; i++ {
			start, end := corners[i], corners[i+1]
			var x1, x2, y1, y2 int
			_, err := fmt.Sscanf(start, "%d,%d", &x1, &y1)
			if err != nil {
				return nil, err
			}
			_, err = fmt.Sscanf(end, "%d,%d", &x2, &y2)
			if err != nil {
				return nil, err
			}

			startX := maths.Min(x1, x2)
			endX := maths.Max(x1, x2)
			startY := maths.Min(y1, y2)
			endY := maths.Max(y1, y2)

			w.minX = maths.Min(w.minX, startX)
			w.maxX = maths.Max(w.maxX, endX)
			w.maxY = maths.Max(w.maxY, endY)

			for x := startX; x <= endX; x++ {
				for y := startY; y <= endY; y++ {
					w.grid[graph.Co{X: x, Y: y}] = "#"
				}
			}
		}
	}
	return w, nil
}

func (w *Waterfall) releaseTheSand(animate, slowMo bool) int {
	sand := graph.Co{X: 500, Y: 0}
	sandCount := 0
	if animate {
		w.PrintGrid(false, slowMo)
	}
	for sand.Y < w.maxY+2 {
		if animate {
			w.PrintGrid(true, slowMo)
		}
		co := graph.Co{X: sand.X, Y: sand.Y + 1}
		if _, ok := w.grid[co]; !ok {
			delete(w.grid, sand)
			sand = co
			w.grid[sand] = "o"
			continue
		}
		co = graph.Co{X: sand.X - 1, Y: sand.Y + 1}
		if _, ok := w.grid[co]; !ok {
			delete(w.grid, sand)
			sand = co
			w.grid[sand] = "o"
			continue
		}
		co = graph.Co{X: sand.X + 1, Y: sand.Y + 1}
		if _, ok := w.grid[co]; !ok {
			delete(w.grid, sand)
			sand = co
			w.grid[sand] = "o"
			continue
		}

		sand = graph.Co{X: 500, Y: 0}
		sandCount++
	}
	return sandCount
}

func main() {
	input := file.Read()
	waterfall, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Only usable for part 1!
	part1 := waterfall.releaseTheSand(false, false)
	fmt.Println("Part 1:", part1)
}
