package main

import (
	"fmt"
)

var puzzleInput = 681901
var elf1 = 0
var elf2 = 1

var recipes = []int{3, 7}

func getNextRecipes() {
	sum := recipes[elf1] + recipes[elf2]
	if sum > 9 {
		recipes = append(recipes, 1)
	}
	recipes = append(recipes, sum%10)
}

func moveElf(elf int) int {
	elfScore := elf + recipes[elf] + 1
	return (elfScore % len(recipes))
}

func printRecipes() {
	for i, v := range recipes {
		if i == elf1 {
			fmt.Printf("(%v)", v)
		} else if i == elf2 {
			fmt.Printf("[%v]", v)
		} else {
			fmt.Printf(" %v ", v)
		}
	}
	fmt.Println()
}

func main() {
	for {
		// printRecipes()
		if len(recipes) > puzzleInput+10 {
			fmt.Println(recipes[puzzleInput : puzzleInput+10])
			return
		}
		getNextRecipes()
		elf1 = moveElf(elf1)
		elf2 = moveElf(elf2)
	}
}
