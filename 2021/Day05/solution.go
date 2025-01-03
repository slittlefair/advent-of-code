package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"Advent-of-Code/maths"
	"Advent-of-Code/regex"
	"fmt"
	"strconv"
)

type Grid map[graph.Co]int

func convertToCos(minX, maxX, minY, maxY int) []graph.Co {
	cos := []graph.Co{}
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			cos = append(cos, graph.Co{X: x, Y: y})
		}
	}
	return cos
}

func convertDiagonalToCos(startX, endX, startY, endY int) []graph.Co {
	cos := []graph.Co{}
	xCoords := []int{}
	yCoords := []int{}
	if startX < endX {
		for x := startX; x <= endX; x++ {
			xCoords = append(xCoords, x)
		}
	} else {
		for x := startX; x >= endX; x-- {
			xCoords = append(xCoords, x)
		}
	}
	if startY < endY {
		for y := startY; y <= endY; y++ {
			yCoords = append(yCoords, y)
		}
	} else {
		for y := startY; y >= endY; y-- {
			yCoords = append(yCoords, y)
		}
	}
	for i := 0; i < len(xCoords); i++ {
		cos = append(cos, graph.Co{X: xCoords[i], Y: yCoords[i]})
	}
	return cos
}

func validCos(input []string, part2 bool) ([]graph.Co, error) {
	cos := []graph.Co{}
	for _, line := range input {
		matches := regex.MatchNums.FindAllString(line, -1)
		if len(matches) != 4 {
			return nil, fmt.Errorf("expected 4 matches, got %d in %v", len(matches), matches)
		}
		matchInts := []int{}
		for _, m := range matches {
			// matches will be able to be converted due to regex match, so we know we won't get an error
			conv, _ := strconv.Atoi(m)
			matchInts = append(matchInts, conv)
		}
		m0 := matchInts[0]
		m1 := matchInts[1]
		m2 := matchInts[2]
		m3 := matchInts[3]
		if m0 == m2 || m1 == m3 {
			cos = append(
				cos,
				convertToCos(maths.Min(m0, m2), maths.Max(m0, m2), maths.Min(m1, m3), maths.Max(m1, m3))...)
		} else if part2 {
			cos = append(cos, convertDiagonalToCos(m0, m2, m1, m3)...)
		}
	}
	return cos, nil
}

func populateGrid(cos []graph.Co) Grid {
	g := Grid{}
	for _, co := range cos {
		if val, ok := g[co]; !ok {
			g[co] = 1
		} else {
			g[co] = val + 1
		}
	}
	return g
}

func (g Grid) countOverlaps() int {
	count := 0
	for _, val := range g {
		if val > 1 {
			count++
		}
	}
	return count
}

func findSolution(input []string, part2 bool) (int, error) {
	cos, err := validCos(input, part2)
	if err != nil {
		return -1, err
	}
	g := populateGrid(cos)
	return g.countOverlaps(), nil
}

func main() {
	input := file.Read()
	part1, err := findSolution(input, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)

	part2, err := findSolution(input, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 2:", part2)
}
