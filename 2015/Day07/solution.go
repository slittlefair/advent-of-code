package main

import (
	"Advent-of-Code"
	"fmt"
	"regexp"
)

type v map[string]int

var reNum = regexp.MustCompile("\\d+")
var reCaps = regexp.MustCompile("[A-Z]+")
var reIdentifier = regexp.MustCompile("[a-z]+")

func (values v) assign(value int, id string) {
	values[id] = value
}

func (values v) iterateValues(instructions []string) {
	addingValues := true
	for addingValues {
		for _, inst := range instructions {
			nums := helpers.StringSliceToIntSlice(reNum.FindAllString(inst, -1))
			caps := reCaps.FindAllString(inst, -1)
			id := reIdentifier.FindAllString(inst, -1)
			if len(caps) == 0 {
				if len(nums) != 0 {
					if _, ok := values[id[0]]; !ok {
						values.assign(nums[0], id[0])
					}
				} else if _, ok := values[id[0]]; ok {
					values.assign(values[id[0]], id[1])
				}
			} else if _, ok := values[id[0]]; ok {
				switch caps[0] {
				case "AND":
					if _, ok := values[id[1]]; len(nums) == 0 && ok {
						values.assign(values[id[0]]&values[id[1]], id[2])
					} else if len(nums) > 0 {
						values.assign(values[id[0]]&nums[0], id[1])
					}
				case "OR":
					if _, ok := values[id[1]]; len(nums) == 0 && ok {
						values.assign(values[id[0]]|values[id[1]], id[2])
					} else if len(nums) > 0 {
						values.assign(values[id[0]]|nums[0], id[1])
					}
				case "LSHIFT":
					values.assign(values[id[0]]<<uint(nums[0]), id[1])
				case "RSHIFT":
					values.assign(values[id[0]]>>uint(nums[0]), id[1])
				case "NOT":
					values.assign(65535-values[id[0]], id[1])
				}
			}
		}
		if len(values) == len(instructions) {
			addingValues = false
		}
	}
}

func main() {
	instructions := helpers.ReadFile()
	part1Values := make(v)
	part1Values.iterateValues(instructions)
	fmt.Println("Part 1:", part1Values["a"])

	part2Values := v{"b": part1Values["a"]}
	part2Values.iterateValues(instructions)
	fmt.Println("Part 2:", part2Values["a"])
}
