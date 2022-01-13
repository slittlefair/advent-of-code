package main

import (
	"Advent-of-Code/utils"
	"fmt"
	"regexp"
	"strconv"
)

type TargetArea struct {
	MinX, MaxX, MinY, MaxY int
	GreatestSuccessfulY    int
	SuccessfulTrajectories int
}

func parseInput(input string) (*TargetArea, error) {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(input, -1)
	if len(matches) != 4 {
		return nil, fmt.Errorf("expected 4 numbers, received %d: %v from line %s", len(matches), matches, input)
	}
	conv := [4]int{}
	for i, m := range matches {
		c, err := strconv.Atoi(m)
		if err != nil {
			return nil, err
		}
		conv[i] = c
	}
	ta := &TargetArea{
		MinX:                conv[0],
		MaxX:                conv[1],
		MinY:                conv[2],
		MaxY:                conv[3],
		GreatestSuccessfulY: conv[2],
	}
	return ta, nil
}

func (ta TargetArea) isInTargetArea(co utils.Co) bool {
	return co.X >= ta.MinX && co.X <= ta.MaxX && co.Y >= ta.MinY && co.Y <= ta.MaxY
}

func (ta *TargetArea) evaluateTrajectory(x, y int) {
	velocity := utils.Co{X: x, Y: y}
	currentPosition := utils.Co{X: 0, Y: 0}
	highestYForShot := y
	for {
		currentPosition.X += velocity.X
		currentPosition.Y += velocity.Y
		highestYForShot = utils.Max(highestYForShot, currentPosition.Y)
		if ta.isInTargetArea(currentPosition) {
			ta.SuccessfulTrajectories++
			ta.GreatestSuccessfulY = utils.Max(ta.GreatestSuccessfulY, highestYForShot)
			break
		}
		// This trajectory won't treach x as we haven't reached it and x is now 0
		if velocity.X == 0 && currentPosition.X < ta.MinX {
			break
		}
		// This trajectory has passed the target area along the x axis
		if currentPosition.X > ta.MaxX {
			break
		}
		// This trajectory has passed the target area along the y axis
		if currentPosition.Y < ta.MinY {
			break
		}
		if velocity.X > 0 {
			velocity.X--
		}
		velocity.Y--
	}
}

func (ta *TargetArea) findTrajectories() (int, int) {
	for y := ta.MinY; y <= -ta.MinY; y++ {
		for x := 0; x <= ta.MaxX; x++ {
			ta.evaluateTrajectory(x, y)
		}
	}
	return ta.GreatestSuccessfulY, ta.SuccessfulTrajectories
}

func main() {
	input := utils.ReadFile()
	ta, err := parseInput(input[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	part1, part2 := ta.findTrajectories()
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
