package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"fmt"
	"time"
)

type wideBox struct {
	opposite graph.Co
	value    string
}

type Warehouse struct {
	graph.Grid[string]
	robot      graph.Co
	walls      map[graph.Co]bool
	boxes      map[graph.Co]bool
	wideBoxes  map[graph.Co]wideBox
	directions []graph.Co
}

func parseInput(input []string) *Warehouse {
	w := &Warehouse{
		Grid: graph.Grid[string]{
			MaxX: len(input[0]) - 1,
		},
		walls: map[graph.Co]bool{},
		boxes: map[graph.Co]bool{},
	}

	var breakline int
	for y, line := range input {
		if line == "" {
			breakline = y
			break
		}
		for x, c := range line {
			s := string(c)
			co := graph.Co{X: x, Y: y}
			switch s {
			case "#":
				w.walls[co] = true
			case "@":
				w.robot = co
			case "O":
				w.boxes[co] = true
			}
		}
	}
	w.MaxY = breakline - 1

	dirs := []graph.Co{}
	for i := breakline + 1; i < len(input); i++ {
		line := input[i]
		for _, char := range line {
			switch string(char) {
			case ">":
				dirs = append(dirs, graph.Co{X: 1})
			case "<":
				dirs = append(dirs, graph.Co{X: -1})
			case "v":
				dirs = append(dirs, graph.Co{Y: 1})
			case "^":
				dirs = append(dirs, graph.Co{Y: -1})
			}
		}
	}
	w.directions = dirs

	return w
}

func (w Warehouse) widenWarehouse() *Warehouse {
	ww := &Warehouse{
		Grid: graph.Grid[string]{
			MaxX: w.MaxX*2 + 1,
			MaxY: w.MaxY,
		},
		robot:     graph.Co{X: w.robot.X * 2, Y: w.robot.Y},
		walls:     map[graph.Co]bool{},
		boxes:     map[graph.Co]bool{},
		wideBoxes: map[graph.Co]wideBox{},
	}
	for wall := range w.walls {
		ww.walls[graph.Co{X: wall.X * 2, Y: wall.Y}] = true
		ww.walls[graph.Co{X: wall.X*2 + 1, Y: wall.Y}] = true
	}
	for box := range w.boxes {
		leftCo := graph.Co{X: box.X * 2, Y: box.Y}
		rightCo := graph.Co{X: box.X*2 + 1, Y: box.Y}
		ww.wideBoxes[leftCo] = wideBox{opposite: rightCo, value: "["}
		ww.wideBoxes[rightCo] = wideBox{opposite: leftCo, value: "]"}
	}
	return ww
}

type option func(*generateConfig)
type generateConfig struct {
	withDirs, wide bool
}

func withDirs() option {
	return func(c *generateConfig) {
		c.withDirs = true
	}
}

func withWide() option {
	return func(c *generateConfig) {
		c.wide = true
	}
}

func getGenerateConfig(opts []option) *generateConfig {
	cfg := &generateConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func (w Warehouse) print(opts ...option) {
	gen := getGenerateConfig(opts)
	for y := w.MinY; y <= w.MaxY; y++ {
		for x := w.MinX; x <= w.MaxX; x++ {
			co := graph.Co{X: x, Y: y}
			if w.robot == co {
				fmt.Print("@")
			} else if _, ok := w.walls[co]; ok {
				fmt.Print("#")
			} else if _, ok := w.boxes[co]; ok && !gen.wide {
				fmt.Print("O")
			} else if v, ok := w.wideBoxes[co]; ok && gen.wide {
				fmt.Print(v.value)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()

	if !gen.withDirs {
		return
	}

	for _, co := range w.directions {
		fmt.Printf("%+v\n", co)
	}
}

func dirString(dir graph.Co) string {
	if dir.X == -1 {
		return "<"
	}
	if dir.X == 1 {
		return ">"
	}
	if dir.Y == -1 {
		return "^"
	}
	return "v"
}

func (w *Warehouse) animate() {
	// print thing
	for i, dir := range w.directions {
		w.move(dir)
		w.print()
		fmt.Printf("%v", dirString(dir))
		time.Sleep(450 * time.Millisecond)
		if i != len(w.directions)-1 {
			// For every line we've written, write the "cursor up" ANSI escape code
			for i := w.MinY; i <= w.MaxY+1; i++ {
				fmt.Printf("\033[A")
			}
		}
	}
}

func (w *Warehouse) canMove(co, dir graph.Co, opts ...option) *graph.Co {
	cfg := getGenerateConfig(opts)
	if _, ok := w.walls[co]; ok {
		return nil
	}
	if cfg.wide {
		if _, ok := w.wideBoxes[co]; ok {
			newCo := co.Add(dir)
			return w.canMove(newCo, dir, withWide())
		}
	} else {
		if _, ok := w.boxes[co]; ok {
			newCo := co.Add(dir)
			return w.canMove(newCo, dir)
		}
	}
	return &co
}

func (w *Warehouse) moveHorizontally(dir graph.Co) {
	co := w.robot.Add(dir)
	endCo := w.canMove(co, dir, withWide())
	if endCo == nil {
		return
	}
	if _, ok := w.boxes[co]; !ok {
		w.robot = co
		return
	}

	// delete(w.boxes, co)
	// for {
	// 	newCo := co.Add(dir)
	// 	if _, ok := w.boxes[newCo]; !ok {
	// 		w.boxes[newCo] = true
	// 		return
	// 	}
	// 	co = newCo
	// }
}

func (w *Warehouse) move(dir graph.Co) {
	co := w.robot.Add(dir)
	endCo := w.canMove(co, dir)
	if endCo == nil {
		return
	}
	w.robot = co
	if _, ok := w.boxes[co]; !ok {
		return
	}
	delete(w.boxes, co)
	w.boxes[*endCo] = true
}

func (w *Warehouse) sumGPS() int {
	sum := 0
	for co := range w.boxes {
		sum += (100*co.Y + co.X)
	}
	return sum
}

func findSolutions(input []string) (int, int) {
	part1 := 0
	part2 := 0
	warehouse := parseInput(input)
	// wideWarehouse := warehouse.widenWarehouse()
	// wideWarehouse.print(withWide())
	for _, dir := range warehouse.directions {
		warehouse.move(dir)
	}
	// warehouse.print()
	// warehouse.animate()
	part1 = warehouse.sumGPS()
	return part1, part2
}

func main() {
	input := file.Read()
	part1, part2 := findSolutions(input)
	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}
