package main

import (
	"Advent-of-Code"
	"regexp"
	"fmt"
)

type ingredient struct {
	name string
	capacity int
	durability int
	flavour int
	texture int
	calories int
}

var nameToIngredient = make(map[string]ingredient)
var ingredientsToSpoonfuls = make(map[string]int)
var spoonfulCombinations []map[string]int

var totalSpoonfuls = 100
var maxScore = 0
var maxScoreWithCalories = 0

func cookieScore(ingredients map[string]int) int {
	capacityScore := 0
	durabilityScore := 0
	flavourScore := 0
	textureScore := 0
	for ingredient, spoonfuls := range ingredients {
		capacityScore += nameToIngredient[ingredient].capacity * spoonfuls
		durabilityScore += nameToIngredient[ingredient].durability * spoonfuls
		flavourScore += nameToIngredient[ingredient].flavour * spoonfuls
		textureScore += nameToIngredient[ingredient].texture * spoonfuls
	}
	if capacityScore < 0 || durabilityScore < 0 || flavourScore < 0 || textureScore < 0 {
		return 0
	}
	return capacityScore * durabilityScore * flavourScore * textureScore
}

func calories(ingredients map[string]int) (calories int) {
	for ingredient, spoonfuls := range ingredients {
		calories += nameToIngredient[ingredient].calories * spoonfuls
	}
	return
}

func assignSpoonfuls(ingredients []string, ingredientsToSpoonfuls map[string]int, level int) {
	for i := 0; i <= 100; i++ {
		ingredientsToSpoonfuls[ingredients[level]] = i
		totalSpoonfuls := 0
		spoonfuls := []int{}
		for _, val := range ingredientsToSpoonfuls {
			totalSpoonfuls += val
			spoonfuls = append(spoonfuls, val)
		}
		if totalSpoonfuls == 100 {
			if score := cookieScore(ingredientsToSpoonfuls); score > maxScore {
				maxScore = score
			}
			if cals := calories(ingredientsToSpoonfuls); cals == 500 {
				if score := cookieScore(ingredientsToSpoonfuls); score > maxScoreWithCalories {
					maxScoreWithCalories = score
				}
			}
		}
		if level > 0 {
			assignSpoonfuls(ingredients, ingredientsToSpoonfuls, level-1)
		}
	}
}

func main() {
	lines := helpers.ReadFile()
	nameRe := regexp.MustCompile("[A-Z][a-z]+")
	numRe := regexp.MustCompile("-?\\d+")
	for _, l := range lines {
		name := nameRe.FindAllString(l, -1)[0]
		nums := helpers.StringSliceToIntSlice(numRe.FindAllString(l, -1))
		nameToIngredient[name] = ingredient{
			name: name,
			capacity: nums[0],
			durability: nums[1],
			flavour: nums[2],
			texture: nums[3],
			calories: nums[4],
		}
	}
	ingredients := []string{}
	for name := range nameToIngredient {
		ingredientsToSpoonfuls[name] = 0
		ingredients = append(ingredients, name)
	}
	assignSpoonfuls(ingredients, ingredientsToSpoonfuls, len(ingredientsToSpoonfuls)-1)
	fmt.Println("Part 1:", maxScore)
	fmt.Println("Part 2:", maxScoreWithCalories)
}