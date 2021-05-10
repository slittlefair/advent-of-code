package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
)

// type coordinate struct {
// 	X int
// 	Y int
// }

// var part1Lights = make(map[coordinate]bool)
// var part2Lights = make(map[coordinate]int)

// func updateStatus(x int, y int, instruction string) {
// 	co := coordinate{x, y}
// 	switch instruction {
// 	case "on":
// 		part1Lights[co] = true
// 	case "off":
// 		part1Lights[co] = false
// 	case "through":
// 		part1Lights[co] = !part1Lights[co]
// 	}
// }

// func updateBrightness(x int, y int, instruction string) {
// 	co := coordinate{x, y}
// 	switch instruction {
// 	case "on":
// 		part2Lights[co] = part2Lights[co] + 1
// 	case "off":
// 		if part2Lights[co] > 0 {
// 			part2Lights[co] = part2Lights[co] - 1
// 		}
// 	case "through":
// 		part2Lights[co] = part2Lights[co] + 2
// 	}
// }

type Lights map[helpers.Co]bool

func populateLights() *Lights {
	lights := Lights{}
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			lights[helpers.Co{X: x, Y: y}] = false
		}
	}
	return &lights
}

func (l Lights) turnLightsOn(nums []int) {
	for x := nums[0]; x <= nums[2]; x++ {
		for y := nums[1]; y <= nums[3]; y++ {
			l[helpers.Co{X: x, Y: y}] = true
		}
	}
}

func (l Lights) turnLightsOff(nums []int) {
	for x := nums[0]; x <= nums[2]; x++ {
		for y := nums[1]; y <= nums[3]; y++ {
			l[helpers.Co{X: x, Y: y}] = false
		}
	}
}

func (l Lights) toggleLights(nums []int) {
	for x := nums[0]; x <= nums[2]; x++ {
		for y := nums[1]; y <= nums[3]; y++ {
			l[helpers.Co{X: x, Y: y}] = !l[helpers.Co{X: x, Y: y}]
		}
	}
}

func (l *Lights) followInstructions(input []string) error {
	re := regexp.MustCompile(`\d+`)
	reOn := regexp.MustCompile(`on`)
	reOff := regexp.MustCompile(`off`)
	reToggle := regexp.MustCompile(`toggle`)
	for _, inst := range input {
		nums := re.FindAllString(inst, -1)
		if len(nums) != 4 {
			return fmt.Errorf("something went wrong, got nums %v", nums)
		}
		intNums := helpers.StringSliceToIntSlice(nums)
		if reOn.MatchString(inst) {
			l.turnLightsOn(intNums)
			continue
		}
		if reOff.MatchString(inst) {
			l.turnLightsOff(intNums)
			continue
		}
		if reToggle.MatchString(inst) {
			l.toggleLights(intNums)
			continue
		}
		return fmt.Errorf("something went wrong, got instruction %s", inst)
	}
	return nil
}

func (l *Lights) countLights() int {
	count := 0
	for _, light := range *l {
		if light {
			count++
		}
	}
	return count
}

func main() {
	input := helpers.ReadFile()
	lights := populateLights()
	lights.followInstructions(input)
	fmt.Println("Part 1:", lights.countLights())
	// for x := 0; x < 1000; x++ {
	// 	for y := 0; y < 1000; y++ {
	// 		part1Lights[coordinate{x, y}] = false
	// 		part2Lights[coordinate{x, y}] = 0
	// 	}
	// }

	// instructions := helpers.ReadFile()
	// coordsRe := regexp.MustCompile("\\d+")
	// wordsRe := regexp.MustCompile("[a-z]+")

	// for _, inst := range instructions {
	// 	points := helpers.StringSliceToIntSlice(coordsRe.FindAllString(inst, -1))
	// 	words := wordsRe.FindAllString(inst, -1)
	// 	for x := points[0]; x <= points[2]; x++ {
	// 		for y := points[1]; y <= points[3]; y++ {
	// 			updateStatus(x, y, words[1])
	// 			updateBrightness(x, y, words[1])
	// 		}
	// 	}
	// }

	// totalOn := 0
	// totalBrightness := 0
	// for co := range part1Lights {
	// 	if part1Lights[co] {
	// 		totalOn++
	// 	}
	// 	totalBrightness += part2Lights[co]
	// }
	// fmt.Println("Part 1:", totalOn)
	// fmt.Println("Part 2:", totalBrightness)
}
