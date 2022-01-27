package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"strings"
)

type Cave struct {
	id         string
	small      bool
	neighbours map[string]*Cave
}

type System struct {
	caves map[string]*Cave
	paths []Path
}

type Path []*Cave

type canVisit func(cave *Cave, path Path) bool

func makeCave(id string) *Cave {
	return &Cave{
		id:         id,
		neighbours: map[string]*Cave{},
		small:      utils.IsLower(id),
	}
}

func parseInput(input []string) (*System, error) {
	neighbours := map[string][]string{}
	system := &System{
		caves: map[string]*Cave{},
	}
	for _, line := range input {
		caves := strings.Split(line, "-")
		if len(caves) != 2 {
			return nil, fmt.Errorf("expected 2 caves from line %v, got %v", line, caves)
		}
		c0, c1 := caves[0], caves[1]
		neighbours[c0] = append(neighbours[c0], c1)
		neighbours[c1] = append(neighbours[c1], c0)
		if _, ok := system.caves[c0]; !ok {
			system.caves[c0] = makeCave(c0)
		}
		if _, ok := system.caves[c1]; !ok {
			system.caves[c1] = makeCave(c1)
		}
	}
	for k, v := range system.caves {
		for _, nbs := range neighbours[k] {
			v.neighbours[nbs] = system.caves[nbs]
		}
	}
	return system, nil
}

func canVisitPart1(cave *Cave, path Path) bool {
	if !cave.small {
		return true
	}
	for _, c := range path {
		if c == cave {
			return false
		}
	}
	return true
}

func canVisitPart2(cave *Cave, path Path) bool {
	if cave.id == "start" {
		return false
	}
	if !cave.small {
		return true
	}
	smallCavesVisited := map[string]int{}
	for _, c := range path {
		if c.small {
			smallCavesVisited[c.id]++
		}
	}
	visitedASmallCaveMoreThanOnce := false
	for _, v := range smallCavesVisited {
		if v > 1 {
			visitedASmallCaveMoreThanOnce = true
		}
	}
	if smallCavesVisited[cave.id] == 2 {
		return false
	}
	if smallCavesVisited[cave.id] == 1 && visitedASmallCaveMoreThanOnce {
		return false
	}
	return true
}

// Debugging
// func printPath(path Path) {
// 	for i := 0; i < len(path)-1; i++ {
// 		fmt.Printf("%s,", path[i].id)
// 	}
// 	fmt.Println(path[len(path)-1].id)
// }

func (s *System) getNextCave(currentCave *Cave, currentPath Path, canVisit canVisit) {
	currentPath = append(currentPath, currentCave)
	if currentCave == s.caves["end"] {
		s.paths = append(s.paths, currentPath)
		return
	}
	for _, n := range currentCave.neighbours {
		if canVisit(n, currentPath) {
			s.getNextCave(n, currentPath, canVisit)
		}
	}
}

func (s *System) findNumberOfPaths(canVisit canVisit) int {
	s.getNextCave(s.caves["start"], Path{}, canVisit)
	return len(s.paths)
}

func findSolutions(input []string) (int, int, error) {
	s, err := parseInput(input)
	if err != nil {
		return -1, -1, err
	}
	part1 := s.findNumberOfPaths(canVisitPart1)
	s.paths = []Path{}
	part2 := s.findNumberOfPaths(canVisitPart2)
	return part1, part2, nil
}

func main() {
	input := utils.ReadFile()
	part1, part2, err := findSolutions(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
