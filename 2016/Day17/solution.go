package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"Advent-of-Code/maths"
	"Advent-of-Code/slice"
	"crypto/md5"
	"fmt"
)

type direction struct {
	letter string
	dir    graph.Co
}

type directions []direction

type floors map[graph.Co]struct{}

type solution struct {
	input              string
	currentRoom        graph.Co
	currentPath        string
	shortestPathLength int
	shortestPath       string
	directions         directions
	floors             floors
	longestPath        string
}

func setupSolution(input string) solution {
	floors := floors{}
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			floors[graph.Co{X: x, Y: y}] = struct{}{}
		}
	}
	dirs := directions{
		{letter: "U", dir: graph.Co{Y: -1}},
		{letter: "D", dir: graph.Co{Y: 1}},
		{letter: "L", dir: graph.Co{X: -1}},
		{letter: "R", dir: graph.Co{X: 1}},
	}
	return solution{
		currentRoom:        graph.Co{X: 0, Y: 0},
		input:              input,
		shortestPathLength: maths.Infinity,
		directions:         dirs,
		floors:             floors,
		shortestPath:       "",
		longestPath:        "",
	}
}

func (s *solution) reachedEnd() bool {
	if s.currentRoom.X != 3 || s.currentRoom.Y != 3 {
		return false
	}
	if len(s.currentPath) < s.shortestPathLength {
		s.shortestPathLength = len(s.currentPath)
		s.shortestPath = s.currentPath
	}
	if l := len(s.longestPath); len(s.currentPath) > l {
		s.longestPath = s.currentPath
	}
	return true
}

func (s solution) getLockStatus() string {
	hash := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%s", s.input, s.currentPath))))
	return hash[:4]
}

func (s solution) isValidRoom(char string, diff graph.Co) bool {
	if !slice.Contains([]string{"b", "c", "d", "e", "f"}, char) {
		return false
	}
	co := graph.Co{X: s.currentRoom.X + diff.X, Y: s.currentRoom.Y + diff.Y}
	_, ok := s.floors[co]
	return ok
}

func (s *solution) getNextDoor(newPath string, newX, newY int) {
	s.currentPath = newPath
	s.currentRoom = graph.Co{X: newX, Y: newY}
	if s.reachedEnd() {
		return
	}
	statuses := s.getLockStatus()
	for i := 0; i < 4; i++ {
		s.currentPath = newPath
		s.currentRoom = graph.Co{X: newX, Y: newY}
		d := s.directions[i]
		if s.isValidRoom(string(statuses[i]), d.dir) {
			s.getNextDoor(s.currentPath+d.letter, s.currentRoom.X+d.dir.X, s.currentRoom.Y+d.dir.Y)
		}
	}
}

func findSolutions(input string) (string, int) {
	sol := setupSolution(input)
	sol.getNextDoor("", 0, 0)
	return sol.shortestPath, len(sol.longestPath)
}

func main() {
	input := file.Read()[0]
	part1, part2 := findSolutions(input)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
