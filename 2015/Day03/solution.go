package main

import (
	"Advent-of-Code"
	"fmt"
)

type coordinate struct {
	X int
	Y int
}

type houses map[coordinate]int

var delivered = houses{
	coordinate{0, 0}: 1,
}

func (delivered houses) addPresent(co coordinate) {
	if num, ok := delivered[co]; !ok {
		delivered[co] = 1
	} else {
		delivered[co] = num + 1
	}
}

func (co coordinate) move(d rune) coordinate {
	var offset [2]int
	switch string(d) {
	case ">":
		offset = [2]int{1, 0}
	case "<":
		offset = [2]int{-1, 0}
	case "^":
		offset = [2]int{0, 1}
	case "v":
		offset = [2]int{0, -1}
	}
	co = coordinate{co.X + offset[0], co.Y + offset[1]}
	delivered.addPresent(co)
	return co
}

func main() {
	directions := helpers.ReadFile()[0]
	co := coordinate{0, 0}
	for _, d := range directions {
		co = co.move(d)
	}
	fmt.Println("Part 1:", len(delivered))
	delivered = houses{
		coordinate{0, 0}: 1,
	}
	co1 := coordinate{0, 0}
	co2 := coordinate{0, 0}
	for i, d := range directions {
		if i%2 == 0 {
			co1 = co1.move(d)
		} else {
			co2 = co2.move(d)
		}
	}
	fmt.Println("Part 2:", len(delivered))
}
