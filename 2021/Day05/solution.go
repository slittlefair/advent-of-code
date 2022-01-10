package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"regexp"
	"strconv"
)

type Grid map[utils.Co]int

func convertToCos(minX, maxX, minY, maxY int) []utils.Co {
	cos := []utils.Co{}
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			cos = append(cos, utils.Co{X: x, Y: y})
		}
	}
	return cos
}

func convertDiagonalToCos(startX, endX, startY, endY int) []utils.Co {
	cos := []utils.Co{}
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
		cos = append(cos, utils.Co{X: xCoords[i], Y: yCoords[i]})
	}
	return cos
}

func validCos(input []string, part2 bool) ([]utils.Co, error) {
	cos := []utils.Co{}
	reNum := regexp.MustCompile(`\d+`)
	for _, line := range input {
		matches := reNum.FindAllString(line, -1)
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
			cos = append(cos, convertToCos(utils.Min(m0, m2), utils.Max(m0, m2), utils.Min(m1, m3), utils.Max(m1, m3))...)
		} else if part2 {
			cos = append(cos, convertDiagonalToCos(m0, m2, m1, m3)...)
		}
	}
	return cos, nil
}

func populateGrid(cos []utils.Co) Grid {
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
	input := utils.ReadFile()
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
