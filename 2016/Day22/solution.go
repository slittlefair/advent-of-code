package main

import (
	"Advent-of-Code/file"
	"fmt"
	"regexp"
	"strconv"
)

type Node struct {
	X          int
	Y          int
	Size       int
	Used       int
	Avail      int
	UsePercent int
}

var re = regexp.MustCompile(`\d+`)

func main() {
	input := file.Read()
	fileSystem := []Node{}
	for i := 2; i < len(input); i++ {
		// Parse input
		matches := re.FindAllString(input[i], -1)
		if len(matches) != 6 {
			fmt.Printf("Error parsing input: %s\n", input[i])
			return
		}
		x, _ := strconv.Atoi(matches[0])
		y, _ := strconv.Atoi(matches[1])
		size, _ := strconv.Atoi(matches[2])
		used, _ := strconv.Atoi(matches[3])
		avail, _ := strconv.Atoi(matches[4])
		usePercent, _ := strconv.Atoi(matches[5])
		fileSystem = append(fileSystem, Node{X: x, Y: y, Size: size, Used: used, Avail: avail, UsePercent: usePercent})
	}

	viablePairs := 0
	for i, n := range fileSystem {
		for j, nn := range fileSystem {
			if i == j {
				continue
			}
			if n.Used == 0 {
				continue
			}
			if n.Used <= nn.Avail {
				viablePairs++
			}
		}
	}
	fmt.Printf("Part 1: %v\n", viablePairs)
}
