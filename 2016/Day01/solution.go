package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Position struct {
	direction int
	location  utils.Co
	seen      map[utils.Co]bool
	hq        utils.Co
}

func (p *Position) turnLeft() {
	if p.direction == 0 {
		p.direction = 3
	} else {
		p.direction--
	}
}

func (p *Position) turnRight() {
	if p.direction == 3 {
		p.direction = 0
	} else {
		p.direction++
	}
}

func (p *Position) move() {
	switch p.direction {
	case 0:
		p.location.Y += 1
	case 1:
		p.location.X += 1
	case 2:
		p.location.Y -= 1
	case 3:
		p.location.X -= 1
	}
	// assume that we don't start at the hq
	if p.seen[p.location] && p.hq.X == 0 && p.hq.Y == 0 {
		p.hq = p.location
	}
	p.seen[p.location] = true
}

func (p *Position) followInstruction(inst string) error {
	strRe := regexp.MustCompile(`[A-Z]`)
	letters := strRe.FindAllString(inst, -1)
	if len(letters) != 1 {
		return fmt.Errorf("malformed instruction %s", inst)
	}
	if dir := letters[0]; dir == "L" {
		p.turnLeft()
	} else if dir == "R" {
		p.turnRight()
	} else {
		return fmt.Errorf("malformed instruction %s", inst)
	}
	intRe := regexp.MustCompile(`\d+`)
	ints := intRe.FindAllString(inst, -1)
	if len(ints) != 1 {
		return fmt.Errorf("malformed instruction %s", inst)
	}
	// Ignore this error as we know it converts due to regex match
	intConv, _ := strconv.Atoi(ints[0])
	for i := 0; i < intConv; i++ {
		p.move()
	}
	// p.move(intConv)
	return nil
}

func (p *Position) followSteps(input []string) (int, error) {
	for _, inst := range strings.Split(input[0], " ") {
		err := p.followInstruction(inst)
		if err != nil {
			return -1, err
		}
	}
	return utils.CalculateManhattanDistance(p.location, utils.Co{}), nil
}

func main() {
	input := utils.ReadFile()
	p := &Position{
		seen: map[utils.Co]bool{
			{X: 0, Y: 0}: true,
		},
	}
	part1, err := p.followSteps(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", utils.CalculateManhattanDistance(p.hq, utils.Co{}))
}
