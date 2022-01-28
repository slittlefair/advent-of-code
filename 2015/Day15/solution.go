package main

import (
	"Advent-of-Code/file"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Ingredient struct {
	Name       string
	Capacity   int
	Durability int
	Flavour    int
	Texture    int
	Calories   int
}

type Ingredients struct {
	Ingredients          []*Ingredient
	MaxScore             int
	MaxScoreCalorieLimit int
}

type SpoonfulsMap map[string]int

func parseInput(input []string) Ingredients {
	ingredients := Ingredients{}
	re := regexp.MustCompile(`-?\d+`)
	for _, i := range input {
		split := strings.Split(i, " ")
		name := strings.TrimRight(split[0], ":")
		nums := re.FindAllString(i, -1)
		// We can ignore the errors as we know they'll convert due to regex match
		capacity, _ := strconv.Atoi(nums[0])
		durability, _ := strconv.Atoi(nums[1])
		flavour, _ := strconv.Atoi(nums[2])
		texture, _ := strconv.Atoi(nums[3])
		calories, _ := strconv.Atoi(nums[4])
		ingredients.Ingredients = append(ingredients.Ingredients, &Ingredient{
			Name:       name,
			Capacity:   capacity,
			Durability: durability,
			Flavour:    flavour,
			Texture:    texture,
			Calories:   calories,
		})
	}
	return ingredients
}

func (in Ingredients) getScore(spoonfuls SpoonfulsMap) int {
	totalScores := &Ingredient{}
	for _, i := range in.Ingredients {
		s := spoonfuls[i.Name]
		totalScores.Capacity += i.Capacity * s
		totalScores.Durability += i.Durability * s
		totalScores.Flavour += i.Flavour * s
		totalScores.Texture += i.Texture * s
	}
	if totalScores.Capacity <= 0 {
		return 0
	}
	if totalScores.Durability <= 0 {
		return 0
	}
	if totalScores.Flavour <= 0 {
		return 0
	}
	if totalScores.Texture <= 0 {
		return 0
	}
	return totalScores.Capacity * totalScores.Durability * totalScores.Flavour * totalScores.Texture
}

func (s SpoonfulsMap) countSpoonfuls() int {
	count := 0
	for _, val := range s {
		count += val
	}
	return count
}

func (in Ingredients) is500Calories(spoonfuls SpoonfulsMap) bool {
	calorieCount := 0
	for _, i := range in.Ingredients {
		s := spoonfuls[i.Name]
		calorieCount += i.Calories * s
	}
	return calorieCount == 500
}

func (in *Ingredients) compareMaxScore(score int) {
	if score > in.MaxScore {
		in.MaxScore = score
	}
}

func (in *Ingredients) compareMaxScoreWithCalorieLimit(score int) {
	if score > in.MaxScoreCalorieLimit {
		in.MaxScoreCalorieLimit = score
	}
}

func (in *Ingredients) findOptimumSpoonfuls(spoonfuls SpoonfulsMap, level int) (SpoonfulsMap, int) {
	for i := 0; i <= 100; i++ {
		spoonfuls[in.Ingredients[level].Name] = i
		count := spoonfuls.countSpoonfuls()
		if count == 100 {
			score := in.getScore(spoonfuls)
			in.compareMaxScore(score)
			if in.is500Calories(spoonfuls) {
				in.compareMaxScoreWithCalorieLimit(score)
			}
		}
		if level > 0 {
			spoonfuls, level = in.findOptimumSpoonfuls(spoonfuls, level-1)
		}
	}
	return spoonfuls, level + 1
}

func main() {
	input := file.Read()
	ingredients := parseInput(input)
	ingredients.findOptimumSpoonfuls(SpoonfulsMap{}, len(ingredients.Ingredients)-1)
	fmt.Println("Part 1:", ingredients.MaxScore)
	fmt.Println("Part 2:", ingredients.MaxScoreCalorieLimit)
}
