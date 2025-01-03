package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/regex"
	"fmt"
	"strconv"
)

func triangleIsValid(triangle []int) bool {
	if triangle[0]+triangle[1] <= triangle[2] {
		return false
	}
	if triangle[0]+triangle[2] <= triangle[1] {
		return false
	}
	if triangle[1]+triangle[2] <= triangle[0] {
		return false
	}
	return true
}

func checkTriangles(triangles [][]int) int {
	count := 0
	for _, t := range triangles {
		if triangleIsValid(t) {
			count++
		}
	}
	return count
}

func validateHorizontalTriangles(input []string) int {
	count := 0
	for _, line := range input {
		matches := regex.MatchNums.FindAllString(line, -1)
		triangle := []int{}
		for _, m := range matches {
			// we know this won't error since we only deal with matches to regex
			i, _ := strconv.Atoi(m)
			triangle = append(triangle, i)
		}
		count += checkTriangles([][]int{triangle})
	}
	return count
}

func validateVerticalTriangles(input []string) int {
	count := 0
	triangles := [][]int{{}, {}, {}}
	for i, line := range input {
		matches := regex.MatchNums.FindAllString(line, -1)
		for j, m := range matches {
			// we know this won't error since we only deal with matches to regex
			t, _ := strconv.Atoi(m)
			triangles[j] = append(triangles[j], t)
		}
		if (i+1)%3 == 0 {
			count += checkTriangles(triangles)
			triangles = [][]int{{}, {}, {}}
		}
	}
	return count
}

func main() {
	input := file.Read()
	fmt.Println("Part 1:", validateHorizontalTriangles(input))
	fmt.Println("Part 2:", validateVerticalTriangles(input))
}
