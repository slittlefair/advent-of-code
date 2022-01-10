package main

import (
	utils "Advent-of-Code/utils"
	"fmt"
)

type CodeConstructor struct {
	currentCo utils.Co
	code      []string
}

func (cc *CodeConstructor) move(dir string, keypad map[utils.Co]string) {
	newCo := utils.Co{X: cc.currentCo.X, Y: cc.currentCo.Y}
	if dir == "U" {
		newCo.Y--
	}
	if dir == "D" {
		newCo.Y++
	}
	if dir == "L" {
		newCo.X--
	}
	if dir == "R" {
		newCo.X++
	}
	if _, ok := keypad[newCo]; ok {
		cc.currentCo = newCo
	}
}

func (cc *CodeConstructor) followDirections(line string, keypad map[utils.Co]string) {
	for _, d := range line {
		cc.move(string(d), keypad)
	}
	cc.code = append(cc.code, keypad[cc.currentCo])
}

func (cc CodeConstructor) getCode() string {
	code := ""
	for _, c := range cc.code {
		code += c
	}
	return code
}

func getSolution(input []string, keypad map[utils.Co]string, startingCo utils.Co) string {
	cc := &CodeConstructor{
		currentCo: startingCo,
	}
	for _, line := range input {
		cc.followDirections(line, keypad)
	}
	return cc.getCode()
}

func main() {
	input := utils.ReadFile()
	keypad := map[utils.Co]string{
		{X: 0, Y: 0}: "1",
		{X: 1, Y: 0}: "2",
		{X: 2, Y: 0}: "3",
		{X: 0, Y: 1}: "4",
		{X: 1, Y: 1}: "5",
		{X: 2, Y: 1}: "6",
		{X: 0, Y: 2}: "7",
		{X: 1, Y: 2}: "8",
		{X: 2, Y: 2}: "9",
	}
	fmt.Println("Part 1:", getSolution(input, keypad, utils.Co{X: 1, Y: 1}))
	keypad = map[utils.Co]string{
		{X: 2, Y: 0}: "1",
		{X: 1, Y: 1}: "2",
		{X: 2, Y: 1}: "3",
		{X: 3, Y: 1}: "4",
		{X: 0, Y: 2}: "5",
		{X: 1, Y: 2}: "6",
		{X: 2, Y: 2}: "7",
		{X: 3, Y: 2}: "8",
		{X: 4, Y: 2}: "9",
		{X: 1, Y: 3}: "A",
		{X: 2, Y: 3}: "B",
		{X: 3, Y: 3}: "C",
		{X: 2, Y: 4}: "D",
	}
	fmt.Println("Part 2:", getSolution(input, keypad, utils.Co{X: 0, Y: 2}))
}
