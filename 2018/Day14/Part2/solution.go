package main

import (
	"fmt"
)

var puzzleInput = []int{6, 8, 1, 9, 0, 1}
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

// debugging
// func printRecipes() {
// 	for i, v := range recipes {
// 		if i == elf1 {
// 			fmt.Printf("(%v)", v)
// 		} else if i == elf2 {
// 			fmt.Printf("[%v]", v)
// 		} else {
// 			fmt.Printf(" %v ", v)
// 		}
// 	}
// 	fmt.Println()
// }

func main() {
	for {
		// printRecipes()
		length := len(recipes)
		if len(recipes) >= 7 {
			if recipes[length-7] == puzzleInput[0] {
				if recipes[length-6] == puzzleInput[1] {
					if recipes[length-5] == puzzleInput[2] {
						if recipes[length-4] == puzzleInput[3] {
							if recipes[length-3] == puzzleInput[4] {
								if recipes[length-2] == puzzleInput[5] {
									fmt.Println(length - 7)
									return
								}
							}
						}
					}
				}
			}
		}
		getNextRecipes()
		elf1 = moveElf(elf1)
		elf2 = moveElf(elf2)
	}
}
