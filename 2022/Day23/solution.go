package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"Advent-of-Code/maths"
	"fmt"
)

type Grove struct {
	minX, maxX, minY, maxY int
	elves                  map[graph.Co]bool
}

type Direction int

const (
	North Direction = iota
	South
	West
	East
)

type move struct {
	m      graph.Co
	checks []graph.Co
}

var moves = map[Direction]move{
	North: {m: graph.Co{Y: -1}, checks: []graph.Co{{X: -1}, {}, {X: 1}}},
	South: {m: graph.Co{Y: 1}, checks: []graph.Co{{X: -1}, {}, {X: 1}}},
	West:  {m: graph.Co{X: -1}, checks: []graph.Co{{Y: -1}, {}, {Y: 1}}},
	East:  {m: graph.Co{X: 1}, checks: []graph.Co{{Y: -1}, {}, {Y: 1}}},
}

func (g Grove) PrintGrove() {
	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			if _, ok := g.elves[graph.Co{X: x, Y: y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// newExtremities calculates new max and min values given the elf values
func (g *Grove) newExtremities(elves []graph.Co) {
	g.maxX, g.maxY, g.minX, g.minY = 0, 0, maths.Infinity, maths.Infinity
	for _, elf := range elves {
		g.minX = maths.Min(g.minX, elf.X)
		g.minY = maths.Min(g.minY, elf.Y)
		g.maxX = maths.Max(g.maxX, elf.X)
		g.maxY = maths.Max(g.maxY, elf.Y)
	}
}

func parseInput(input []string) *Grove {
	g := &Grove{
		minX: maths.Infinity,
		maxX: 0,
		minY: maths.Infinity,
		maxY: 0,
	}
	elves := map[graph.Co]bool{}
	elfCos := []graph.Co{}
	for y, line := range input {
		for x, r := range line {
			if string(r) == "#" {
				elf := graph.Co{X: x, Y: y}
				elves[elf] = true
				elfCos = append(elfCos, elf)
			}
		}
	}
	g.newExtremities(elfCos)
	g.elves = elves
	return g
}

// proposeMove takes an elf and initial direction of travel and determines where the elf wants to
// move, returning that new value as well as whether it wants to move at all.
func (g Grove) proposeMove(elf graph.Co, currentDirection Direction) (graph.Co, bool) {
	// If the elf has no adjacent elves it won't move
	adjTiles := graph.AdjacentCos(elf, true)
	willMove := false
	for _, at := range adjTiles {
		if _, ok := g.elves[at]; ok {
			willMove = true
			break
		}
	}
	if !willMove {
		return elf, false
	}
	// Loop over directions and see if the elf can move that way. Once it finds a direction it can
	// move it returns that value as well as that it moved.
out:
	for i := int(currentDirection); i < (int(currentDirection) + len(moves)); i++ {
		move := moves[Direction(i%len(moves))]
		for _, c := range move.checks {
			x := elf.X + move.m.X + c.X
			y := elf.Y + move.m.Y + c.Y
			if _, ok := g.elves[graph.Co{X: x, Y: y}]; ok {
				continue out
			}
		}
		return graph.Co{X: elf.X + move.m.X, Y: elf.Y + move.m.Y}, true
	}
	// If the elf could not move in any direction it returns its original location.
	return elf, false
}

// proposeMoves takes the current direction and tries to move all of its elves and updates its max
// and min coordinates. It returns whether or not there were no elves that moved, in which case all
// elves are in their final positions.
func (g *Grove) proposeMoves(currentDirection Direction) bool {
	// Map of old positions to proposed new positions
	newPositions := map[graph.Co]graph.Co{}
	// Frequencies of proposed new positions. If any new positon has a frequency over 1 then
	// multiple elves want to move there and so they cannot.
	positionFrequencies := map[graph.Co]int{}
	atLeastOneMoved := false
	// For each elf find its proposed new position.
	for elf := range g.elves {
		newPosition, didMove := g.proposeMove(elf, currentDirection)
		newPositions[elf] = newPosition
		positionFrequencies[newPosition]++
		if didMove {
			atLeastOneMoved = true
		}
	}
	// If no elves moved then we are finished, return true.
	if !atLeastOneMoved {
		return true
	}
	newElves := map[graph.Co]bool{}
	elfCos := []graph.Co{}
	// For old and proposed new positions, if more than 1 elf wants to move there then they
	// remain in their original position, so their new locations is the same as their old.
	for old, new := range newPositions {
		if positionFrequencies[new] > 1 {
			new = old
		}
		newElves[new] = true
		elfCos = append(elfCos, new)
	}
	// Populate the grove with the new elf positions after moves and update its extremities
	g.elves = newElves
	g.newExtremities(elfCos)
	return false
}

// findNumEmptyTiles finds all spaces in the grove that don't have an elf in them. Do so by getting
// the square that all elves are in and subtracting the number of elves in that square.
func (g Grove) findNumEmptyTiles() int {
	return ((g.maxX - g.minX + 1) * (g.maxY - g.minY + 1)) - len(g.elves)
}

func findSolutions(input []string) (int, int) {
	grove := parseInput(input)
	round := 0
	part1 := 0
	currentDir := North
	for {
		round++
		if grove.proposeMoves(currentDir) {
			return part1, round
		}
		currentDir = Direction(((int(currentDir) + 1) % len(moves)))
		if round == 10 {
			part1 = grove.findNumEmptyTiles()
		}
	}
}

func main() {
	input := file.Read()
	part1, part2 := findSolutions(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
