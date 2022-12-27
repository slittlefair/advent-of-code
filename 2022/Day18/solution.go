package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"Advent-of-Code/maths"
	"fmt"
)

type Blob struct {
	grid                               map[graph.Co]bool
	minX, maxX, minY, maxY, minZ, maxZ int
}

func parseInput(input []string) (*Blob, error) {
	blob := &Blob{
		grid: map[graph.Co]bool{},
		minX: maths.Infinity,
		maxX: 0,
		minY: maths.Infinity,
		maxY: 0,
		minZ: maths.Infinity,
		maxZ: 0,
	}
	for _, line := range input {
		var x, y, z int
		_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if err != nil {
			return nil, err
		}
		blob.minX = maths.Min(blob.minX, x)
		blob.maxX = maths.Max(blob.maxX, x)
		blob.minY = maths.Min(blob.minY, y)
		blob.maxY = maths.Max(blob.maxY, y)
		blob.minZ = maths.Min(blob.minZ, z)
		blob.maxZ = maths.Max(blob.maxZ, z)
		blob.grid[graph.Co{X: x, Y: y, Z: z}] = true
	}
	for z := blob.minZ - 1; z <= blob.maxZ+1; z++ {
		for y := blob.minY - 1; y <= blob.maxY+1; y++ {
			for x := blob.minX - 1; x <= blob.maxX+1; x++ {
				co := graph.Co{X: x, Y: y, Z: z}
				if _, ok := blob.grid[co]; !ok {
					blob.grid[co] = false
				}
			}
		}
	}
	return blob, nil
}

func (b Blob) calculateSurfaceArea() (int, int) {
	var part1, part2 int
	adjacent := []graph.Co{
		{X: -1},
		{X: 1},
		{Y: -1},
		{Y: 1},
		{Z: -1},
		{Z: 1},
	}
	for co, coVal := range b.grid {
		isGap := 0
		for _, adj := range adjacent {
			adjCo := graph.Co{X: co.X + adj.X, Y: co.Y + adj.Y, Z: co.Z + adj.Z}
			if adjCoVal := b.grid[adjCo]; coVal && !adjCoVal {
				part1++
				part2++
			} else if !coVal && adjCoVal {
				isGap++
			}
		}
		if isGap == 6 {
			part2 -= 6
		}
	}
	fmt.Println((part2-part1)%6 == 0)
	return part1, part2
}

func main() {
	input := file.Read()
	blob, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	part1, part2 := blob.calculateSurfaceArea()
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
