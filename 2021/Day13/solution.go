package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
	"regexp"
	"strconv"
)

type Instruction struct {
	Dir string
	Val int
}

type Dots map[utils.Co]struct{}

type Paper struct {
	Dots         Dots
	Instructions []Instruction
	MaxX         int
	MaxY         int
}

func parseInput(input []string) (*Paper, error) {
	blankIndex := -1
	for i, line := range input {
		if line == "" {
			blankIndex = i
		}
	}
	if blankIndex == -1 {
		return nil, fmt.Errorf("expected blank line, couldn't find one")
	}
	p := &Paper{
		Dots: map[utils.Co]struct{}{},
	}
	reNum := regexp.MustCompile(`\d+`)
	for i := 0; i < blankIndex; i++ {
		matches := reNum.FindAllString(input[i], -1)
		if len(matches) != 2 {
			return nil, fmt.Errorf("expected a valid coordinate, got %s", input[i])
		}
		// We know that matches can be converted to ints due to regex matching, so errors here will be nil
		x, _ := strconv.Atoi(matches[0])
		y, _ := strconv.Atoi(matches[1])
		p.Dots[utils.Co{X: x, Y: y}] = struct{}{}
		p.MaxX = utils.Max(p.MaxX, x)
		p.MaxY = utils.Max(p.MaxY, y)
	}
	reFold := regexp.MustCompile(`(\w)=(\d+)`)
	for i := blankIndex + 1; i < len(input); i++ {
		matches := reFold.FindStringSubmatch(input[i])
		if len(matches) != 3 {
			return nil, fmt.Errorf("expected a valid instruction, got %s", input[i])
		}
		dir := matches[1]
		if dir != "x" && dir != "y" {
			return nil, fmt.Errorf("expected a valid instruction, got %s", input[i])
		}
		// We know that match can be converted to int due to regex matching, so errors here will be nil
		val, _ := strconv.Atoi(matches[2])
		p.Instructions = append(p.Instructions, Instruction{
			Dir: dir,
			Val: val,
		})
	}
	return p, nil
}

func (p Paper) printPaper() {
	for y := 0; y <= p.MaxY; y++ {
		for x := 0; x <= p.MaxX; x++ {
			if _, ok := p.Dots[utils.Co{X: x, Y: y}]; ok {
				fmt.Print("\u2588")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (p *Paper) doFold(idx int) {
	if inst := p.Instructions[idx]; inst.Dir == "x" {
		p.doFoldLeft(inst.Val)
	} else {
		p.doFoldUp(inst.Val)
	}
}

func (p *Paper) doFoldUp(v int) {
	newDots := Dots{}
	for co := range p.Dots {
		if co.Y < v {
			newDots[co] = struct{}{}
		} else {
			newDots[utils.Co{X: co.X, Y: v - (co.Y - v)}] = struct{}{}
		}
	}
	p.Dots = newDots
	p.MaxY = (p.MaxY - 1) / 2
}

func (p *Paper) doFoldLeft(v int) {
	newDots := Dots{}
	for co := range p.Dots {
		if co.X < v {
			newDots[co] = struct{}{}
		} else {
			newDots[utils.Co{X: v - (co.X - v), Y: co.Y}] = struct{}{}
		}
	}
	p.Dots = newDots
	p.MaxX = (p.MaxX - 1) / 2
}

func findSolutions(input []string) (int, *Paper, error) {
	p, err := parseInput(input)
	if err != nil {
		return -1, nil, err
	}
	p.doFold(0)
	part1 := len(p.Dots)
	for i := 1; i < len(p.Instructions); i++ {
		p.doFold(i)
	}
	return part1, p, nil
}

func main() {
	input := utils.ReadFile()
	part1, part2, err := findSolutions(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:")
	part2.printPaper()
}
