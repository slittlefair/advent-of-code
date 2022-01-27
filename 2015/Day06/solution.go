package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"regexp"
)

type Lights struct {
	Analogue map[utils.Co]bool
	Digital  map[utils.Co]int
}

func populateLights() *Lights {
	lights := Lights{
		Analogue: make(map[utils.Co]bool),
		Digital:  make(map[utils.Co]int),
	}
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			lights.Analogue[utils.Co{X: x, Y: y}] = false
			lights.Digital[utils.Co{X: x, Y: y}] = 0
		}
	}
	return &lights
}

func (l Lights) turnLightsOn(nums []int) {
	for x := nums[0]; x <= nums[2]; x++ {
		for y := nums[1]; y <= nums[3]; y++ {
			co := utils.Co{X: x, Y: y}
			l.Analogue[co] = true
			l.Digital[co]++
		}
	}
}

func (l Lights) turnLightsOff(nums []int) {
	for x := nums[0]; x <= nums[2]; x++ {
		for y := nums[1]; y <= nums[3]; y++ {
			co := utils.Co{X: x, Y: y}
			l.Analogue[co] = false
			if l.Digital[co] > 0 {
				l.Digital[co]--
			}
		}
	}
}

func (l Lights) toggleLights(nums []int) {
	for x := nums[0]; x <= nums[2]; x++ {
		for y := nums[1]; y <= nums[3]; y++ {
			co := utils.Co{X: x, Y: y}
			l.Analogue[co] = !l.Analogue[co]
			l.Digital[co] += 2
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
		intNums := utils.StringSliceToIntSlice(nums)
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

func (l *Lights) countAnalogueBrightness() int {
	count := 0
	for _, light := range l.Analogue {
		if light {
			count++
		}
	}
	return count
}

func (l *Lights) countDigitalBrightness() int {
	count := 0
	for _, light := range l.Digital {
		count += light
	}
	return count
}

func main() {
	input := utils.ReadFile()
	lights := populateLights()
	err := lights.followInstructions(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", lights.countAnalogueBrightness())
	fmt.Println("Part 2:", lights.countDigitalBrightness())
}
