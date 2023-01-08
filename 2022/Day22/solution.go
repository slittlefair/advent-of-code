package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"Advent-of-Code/maths"
	"fmt"
	"regexp"
	"strconv"
)

type Facing int

const (
	Right Facing = iota
	Down
	Left
	Up
)

type ForceField struct {
	grid                   map[graph.Co]bool
	co                     graph.Co
	facing                 Facing
	facingDirections       []graph.Co
	minX, maxX, minY, maxY int
	path                   map[graph.Co]string
	facingPrint            []string
}

func (ff ForceField) PrintGrid() {
	for y := ff.minY; y <= ff.maxY; y++ {
		for x := ff.minX; x <= ff.maxX; x++ {
			co := graph.Co{X: x, Y: y}
			if v, ok := ff.grid[co]; !ok {
				fmt.Print(" ")
			} else if p, ok := ff.path[co]; ok {
				fmt.Print(p)
			} else if v {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func parseInput(input []string) (ForceField, []string, []int, error) {
	ff := ForceField{
		minX:             1,
		minY:             1,
		maxY:             len(input) - 2,
		grid:             make(map[graph.Co]bool),
		facing:           Right,
		facingDirections: []graph.Co{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 0, Y: -1}},
		facingPrint:      []string{">", "v", "<", "^"},
	}
	for y := 0; y < len(input)-2; y++ {
		line := input[y]
		for x := 0; x < len(line); x++ {
			ff.maxX = maths.Max(ff.maxX, x)
			if string(line[x]) == "." {
				ff.grid[graph.Co{X: x + 1, Y: y + 1}] = true
				if ff.path == nil {
					ff.co = graph.Co{X: x + 1, Y: y + 1}
					ff.path = map[graph.Co]string{
						{X: x + 1, Y: y + 1}: ">",
					}
				}
			} else if string(line[x]) == "#" {
				ff.grid[graph.Co{X: x + 1, Y: y + 1}] = false
			}
		}
	}

	instLine := input[len(input)-1]
	stringRe := regexp.MustCompile(`[LR]`)
	strings := stringRe.FindAllString(instLine, -1)

	intRe := regexp.MustCompile(`\d+`)
	intStrings := intRe.FindAllString(instLine, -1)

	ints := []int{}
	for _, i := range intStrings {
		// We know all elements of ints can be converted safely due to regex matching, so this error
		// can be safely ignored
		v, _ := strconv.Atoi(i)
		ints = append(ints, v)
	}

	if len(ints) != len(strings)+1 {
		return ff, nil, nil, fmt.Errorf("something went wrong parsing instructions %s: got strings %v, ints %v", instLine, strings, ints)
	}

	return ff, strings, ints, nil
}

func (ff *ForceField) handleLoop(co graph.Co) error {
	switch ff.facing {
	case Right:
		for x := ff.minX; x <= ff.maxX; x++ {
			newCo := graph.Co{X: x, Y: co.Y}
			if v, ok := ff.grid[newCo]; ok {
				if v {
					ff.co = newCo
				}
				return nil
			}
		}
	case Down:
		for y := ff.minY; y <= ff.maxY; y++ {
			newCo := graph.Co{X: co.X, Y: y}
			if v, ok := ff.grid[newCo]; ok {
				if v {
					ff.co = newCo
				}
				return nil
			}
		}
	case Left:
		for x := ff.maxX; x >= ff.minX; x-- {
			newCo := graph.Co{X: x, Y: co.Y}
			if v, ok := ff.grid[newCo]; ok {
				if v {
					ff.co = newCo
				}
				return nil
			}
		}
	case Up:
		for y := ff.maxY; y >= ff.minY; y-- {
			newCo := graph.Co{X: co.X, Y: y}
			if v, ok := ff.grid[newCo]; ok {
				if v {
					ff.co = newCo
				}
				return nil
			}
		}
	default:
		return fmt.Errorf("invalid facing direction %v", ff.facing)
	}
	return fmt.Errorf("could not loop around %v", co)
}

func (ff *ForceField) move(steps int) error {
	for i := 0; i < steps; i++ {
		moveCo := ff.facingDirections[ff.facing]
		newCo := graph.Co{X: ff.co.X + moveCo.X, Y: ff.co.Y + moveCo.Y}
		if v, ok := ff.grid[newCo]; !ok {
			err := ff.handleLoop(newCo)
			if err != nil {
				return err
			}
		} else if v {
			ff.co = newCo
		} else {
			return nil
		}
		ff.path[ff.co] = ff.facingPrint[int(ff.facing)]
	}
	return nil
}

func (ff *ForceField) turn(direction string) error {
	var newFacing int
	switch direction {
	case "L":
		newFacing = int(ff.facing) - 1
	case "R":
		newFacing = int(ff.facing) + 1
	default:
		return fmt.Errorf("invalid direction: %s", direction)
	}
	ff.facing = Facing(maths.Modulo(newFacing, len(ff.facingDirections)))
	ff.path[ff.co] = ff.facingPrint[int(ff.facing)]
	return nil
}

func (ff *ForceField) followInstructions(steps []int, turns []string) error {
	for i := 0; i < len(turns); i++ {
		err := ff.move(steps[i])
		if err != nil {
			return err
		}
		err = ff.turn(turns[i])
		if err != nil {
			return err
		}
	}
	return ff.move(steps[len(steps)-1])
}

func (ff ForceField) password() int {
	return (ff.co.Y * 1000) + (ff.co.X * 4) + int(ff.facing)
}

func main() {
	input := file.Read()
	ff, s, i, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ff.followInstructions(i, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	ff.PrintGrid()

	fmt.Println("Part 1:", ff.password())
}
