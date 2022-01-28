package main

import (
	"Advent-of-Code/file"
	"Advent-of-Code/graph"
	"fmt"
)

// type grid map[graph.Co]string

type scanner struct {
	id      int
	beacons map[graph.Co]struct{}
}

func parseInput(input []string) {}

func main() {
	input := file.Read()
	fmt.Println(input)
}
