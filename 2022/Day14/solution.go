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

func (w Waterfall) PrintGrid(animate, actualInput, slowMo, overwrite bool, co graph.Co, count int) {
	// If we don't want to animate, return early
	if !animate {
		return
	}
	xStart := w.minX
	xEnd := w.maxX
	yStart := w.minY
	yEnd := w.maxY
	if actualInput {
		xStart = co.X - 10
		xEnd = co.X + 10
		yStart = co.Y - 5
		yEnd = co.Y + 5
	}
	// We always want to overwrite the last grid unless it's the first one we're drawing
	if overwrite {
		if slowMo {
			time.Sleep(100 * time.Millisecond)
		} else {
			time.Sleep(30 * time.Millisecond)
		}
		// For every line we've written, write the "cursor up" ANSI escape code
		for i := yStart; i <= yEnd+1; i++ {
			fmt.Printf("\033[A")
		}
	}

	for y := yStart; y <= yEnd; y++ {
		for x := xStart - 1; x <= xEnd+1; x++ {
			printChar := " "
			if v, ok := w.grid[graph.Co{X: x, Y: y}]; x == 500 && y == 0 && v != "." {
				printChar = "+"
			} else if ok {
				printChar = v
			}
			fmt.Print(printChar)
		}
		fmt.Println()
	}
	// Also print the current coordinate of the falling sand
	fmt.Printf("X: %03d, Y: %03d, count %d\n", co.X, co.Y, count)
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

			// For ease of looping set start coordinates as the lower and end as the higher
			startX := maths.Min(x1, x2)
			endX := maths.Max(x1, x2)
			startY := maths.Min(y1, y2)
			endY := maths.Max(y1, y2)

			// Update waterfall's min and max values
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

	// For part 2 increase waterfall.maxY by 2 and add the floor. As the sand can fall diagonally
	// we need to extend the
	w.maxY += 2
	for x := w.minX - 2; x <= w.maxX+2; x++ {
		w.grid[graph.Co{X: x, Y: w.maxY}] = "#"
	}
	return w, nil
}

// If we have reached the floor and are at the edge of the grid then we need to extend it
func (w *Waterfall) extendGrid(sand graph.Co) {
	if sand.X < w.minX {
		w.minX = sand.X
		w.grid[graph.Co{X: sand.X - 1, Y: w.maxY}] = "#"
	}
	if sand.X > w.maxX {
		w.maxX = sand.X
		w.grid[graph.Co{X: sand.X + 1, Y: w.maxY}] = "#"
	}
}

func (w *Waterfall) releaseTheSand(animate, actualInput, slowMo bool) (int, int) {
	sand := graph.Co{X: 500, Y: 0}
	sandCount := 0
	w.PrintGrid(animate, actualInput, slowMo, false, sand, sandCount)
	var part1, part2 int

	for {
		// If we have reached the floor the first time then we have the solution to part 1.
		// Every time we reach the floor we need to extend it if necessary.
		if sand.Y == w.maxY-1 {
			if part1 == 0 {
				part1 = sandCount
			}
			w.extendGrid(sand)
		}

		w.PrintGrid(animate, actualInput, slowMo, true, sand, sandCount)

		// Try and move the sand down, if we can then update the grid and move onto the next loop.
		co := graph.Co{X: sand.X, Y: sand.Y + 1}
		if _, ok := w.grid[co]; !ok {
			delete(w.grid, sand)
			sand = co
			w.grid[sand] = "."
			continue
		}

		// Try and move the sand down and left, if we can then update the grid and move onto the next loop.
		co = graph.Co{X: sand.X - 1, Y: sand.Y + 1}
		if _, ok := w.grid[co]; !ok {
			delete(w.grid, sand)
			sand = co
			w.grid[sand] = "."
			continue
		}

		// Try and move the sand down and right, if we can then update the grid and move onto the next loop.
		co = graph.Co{X: sand.X + 1, Y: sand.Y + 1}
		if _, ok := w.grid[co]; !ok {
			delete(w.grid, sand)
			sand = co
			w.grid[sand] = "."
			continue
		}

		// If we can't move the sand then it comes to rest.
		// If the sand is stopped at the origin then we've found the solution to part2. Update the
		// grid and print it a final time if we're animating.
		// Otherwise create a new grain of sand at the top and start moving that one.
		sandCount++
		if sand.X == 500 && sand.Y == 0 {
			part2 = sandCount
			sand = graph.Co{X: 500, Y: 0}
			w.grid[sand] = "."
			if animate {
				w.PrintGrid(animate, actualInput, slowMo, true, sand, sandCount)
			}
			break
		}
		sand = graph.Co{X: 500, Y: 0}
	}
	return part1, part2
}

func main() {
	input := file.Read()
	waterfall, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	part1, part2 := waterfall.releaseTheSand(false, len(input) > 2, true)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
