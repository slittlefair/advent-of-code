package main

import (
	helpers "Advent-of-Code"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Foods struct {
	AllergenToIngredient map[string]string
	Contains             map[string]map[string]bool
	Input                map[string][]map[string]bool
	OriginalInput        [][]string
	SafeIngredients      map[string]bool
}

func (f *Foods) parseInput(input []string) {
	re := regexp.MustCompile(`\w+`)
	for _, line := range input {
		split := strings.Split(line, "contains")
		ingredients := re.FindAllString(split[0], -1)
		f.OriginalInput = append(f.OriginalInput, ingredients)
		allergens := re.FindAllString(split[1], -1)
		for _, a := range allergens {
			if _, ok := f.Contains[a]; !ok {
				f.Contains[a] = map[string]bool{}
			}
			val := f.Contains[a]
			newMap := map[string]bool{}
			for _, i := range ingredients {
				val[i] = true
				f.SafeIngredients[i] = true
				newMap[i] = true
			}
			f.Input[a] = append(f.Input[a], newMap)
		}
	}
}

func (f Foods) removeSafeFoods() {
	for a, maps := range f.Input {
		for pf := range f.Contains[a] {
			for _, mp := range maps {
				if !mp[pf] {
					delete(f.Contains[a], pf)
					goto out
				}
			}
		out:
		}
	}
	for _, c := range f.Contains {
		for cc := range c {
			delete(f.SafeIngredients, cc)
		}
	}
}

func (f Foods) countSafeIngredients() int {
	count := 0
	for _, ingredients := range f.OriginalInput {
		for _, i := range ingredients {
			if f.SafeIngredients[i] {
				count++
			}
		}
	}
	return count
}

func (f *Foods) assignIngredientToAllergen() {
	for {
		if len(f.Contains) == len(f.AllergenToIngredient) {
			return
		}
		for allergen, contains := range f.Contains {
			if len(contains) == 1 {
				var ingredient string
				for c := range contains {
					f.AllergenToIngredient[allergen] = c
					ingredient = c
				}
				for allergen2, contains2 := range f.Contains {
					if allergen2 != allergen {
						delete(contains2, ingredient)
					}
				}
			}
		}
	}
}

func (f Foods) createDangerousIngredientsList() string {
	allergens := []string{}
	for allergen := range f.AllergenToIngredient {
		allergens = append(allergens, allergen)
	}
	sort.Strings(allergens)
	ingredients := []string{}
	for _, allergen := range allergens {
		ingredients = append(ingredients, f.AllergenToIngredient[allergen])
	}
	return strings.Join(ingredients, ",")
}

func main() {
	input := helpers.ReadFile()
	f := Foods{
		AllergenToIngredient: make(map[string]string),
		Contains:             make(map[string]map[string]bool),
		Input:                make(map[string][]map[string]bool),
		SafeIngredients:      make(map[string]bool),
	}
	f.parseInput(input)
	f.removeSafeFoods()
	fmt.Println("Part 1:", f.countSafeIngredients())

	f.assignIngredientToAllergen()
	fmt.Println("Part 2:", f.createDangerousIngredientsList())
}
