package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInput(t *testing.T) {
	t.Run("parses a simple list of 2 ingredients, advent of code example", func(t *testing.T) {
		arg := []string{
			"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8",
			"Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3",
		}
		want := Ingredients{
			Ingredients: []*Ingredient{
				{
					Name:       "Butterscotch",
					Capacity:   -1,
					Durability: -2,
					Flavour:    6,
					Texture:    3,
					Calories:   8,
				},
				{
					Name:       "Cinnamon",
					Capacity:   2,
					Durability: 3,
					Flavour:    -2,
					Texture:    -1,
					Calories:   3,
				},
			},
		}
		got := parseInput(arg)
		assert.Equal(t, want, got)
	})
}

func TestIngredients_getScore(t *testing.T) {
	tests := []struct {
		name string
		in   Ingredients
		arg  SpoonfulsMap
		want int
	}{
		{
			name: "returns the correct score for the given amount of ingredients, advent of code example",
			in: Ingredients{
				Ingredients: []*Ingredient{
					{
						Name:       "Butterscotch",
						Capacity:   -1,
						Durability: -2,
						Flavour:    6,
						Texture:    3,
						Calories:   8,
					},
					{
						Name:       "Cinnamon",
						Capacity:   2,
						Durability: 3,
						Flavour:    -2,
						Texture:    -1,
						Calories:   3,
					},
				},
			},
			arg: SpoonfulsMap{
				"Butterscotch": 44,
				"Cinnamon":     56,
			},
			want: 62842880,
		},
		{
			name: "returns a score of 0 if Capacity score is less than 0",
			in: Ingredients{
				Ingredients: []*Ingredient{
					{
						Name:       "Butterscotch",
						Capacity:   -1,
						Durability: -2,
						Flavour:    6,
						Texture:    3,
						Calories:   8,
					},
					{
						Name:       "Cinnamon",
						Capacity:   2,
						Durability: 3,
						Flavour:    -2,
						Texture:    -1,
						Calories:   3,
					},
				},
			},
			arg: SpoonfulsMap{
				"Butterscotch": 99,
				"Cinnamon":     1,
			},
			want: 0,
		},
		{
			name: "returns a score of 0 if Durability score is less than 0",
			in: Ingredients{
				Ingredients: []*Ingredient{
					{
						Name:       "Butterscotch",
						Capacity:   -1,
						Durability: -2,
						Flavour:    6,
						Texture:    3,
						Calories:   8,
					},
					{
						Name:       "Cinnamon",
						Capacity:   2,
						Durability: 3,
						Flavour:    -2,
						Texture:    -1,
						Calories:   3,
					},
				},
			},
			arg: SpoonfulsMap{
				"Butterscotch": 62,
				"Cinnamon":     38,
			},
			want: 0,
		},
		{
			name: "returns a score of 0 if Flavour score is less than 0",
			in: Ingredients{
				Ingredients: []*Ingredient{
					{
						Name:       "Butterscotch",
						Capacity:   -1,
						Durability: -2,
						Flavour:    6,
						Texture:    3,
						Calories:   8,
					},
					{
						Name:       "Cinnamon",
						Capacity:   2,
						Durability: 3,
						Flavour:    -2,
						Texture:    -1,
						Calories:   3,
					},
				},
			},
			arg: SpoonfulsMap{
				"Butterscotch": 1,
				"Cinnamon":     10,
			},
			want: 0,
		},
		{
			name: "returns a score of 0 if Texture score is less than 0",
			in: Ingredients{
				Ingredients: []*Ingredient{
					{
						Name:       "Butterscotch",
						Capacity:   -1,
						Durability: -2,
						Flavour:    6,
						Texture:    3,
						Calories:   8,
					},
					{
						Name:       "Cinnamon",
						Capacity:   2,
						Durability: 3,
						Flavour:    2,
						Texture:    -1,
						Calories:   3,
					},
				},
			},
			arg: SpoonfulsMap{
				"Butterscotch": 1,
				"Cinnamon":     10,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in.getScore(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSpoonfulsMap_countSpoonfuls(t *testing.T) {
	tests := []struct {
		name string
		s    SpoonfulsMap
		want int
	}{
		{
			name: "counts the total number of spoonfuls in an empty map",
			want: 0,
		},
		{
			name: "counts the total number of spoonfuls in a 2 ingredient map",
			s: SpoonfulsMap{
				"Butterscotch": 9,
				"Cinnamon":     44,
			},
			want: 53,
		},
		{
			name: "counts the total number of spoonfuls in a 4 ingredient map",
			s: SpoonfulsMap{
				"Butterscotch": 9,
				"Cinnamon":     44,
				"Nutmeg":       1,
				"Sugar":        199,
			},
			want: 253,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.countSpoonfuls()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIngredients_is500Calories(t *testing.T) {
	tests := []struct {
		name string
		in   Ingredients
		arg  SpoonfulsMap
		want bool
	}{
		{
			name: "returns false if calories are less than 500",
			in: Ingredients{
				Ingredients: []*Ingredient{
					{
						Name:       "Butterscotch",
						Capacity:   -1,
						Durability: -2,
						Flavour:    6,
						Texture:    3,
						Calories:   8,
					},
					{
						Name:       "Cinnamon",
						Capacity:   2,
						Durability: 3,
						Flavour:    -2,
						Texture:    -1,
						Calories:   3,
					},
				},
			},
			arg: SpoonfulsMap{
				"Butterscotch": 50,
				"Cinnamon":     33,
			},
			want: false,
		},
		{
			name: "returns false if calories are more than 500",
			in: Ingredients{
				Ingredients: []*Ingredient{
					{
						Name:       "Butterscotch",
						Capacity:   -1,
						Durability: -2,
						Flavour:    6,
						Texture:    3,
						Calories:   8,
					},
					{
						Name:       "Cinnamon",
						Capacity:   2,
						Durability: 3,
						Flavour:    -2,
						Texture:    -1,
						Calories:   3,
					},
				},
			},
			arg: SpoonfulsMap{
				"Butterscotch": 60,
				"Cinnamon":     7,
			},
			want: false,
		},
		{
			name: "returns true if calories are exactly 500",
			in: Ingredients{
				Ingredients: []*Ingredient{
					{
						Name:       "Butterscotch",
						Capacity:   -1,
						Durability: -2,
						Flavour:    6,
						Texture:    3,
						Calories:   8,
					},
					{
						Name:       "Cinnamon",
						Capacity:   2,
						Durability: 3,
						Flavour:    -2,
						Texture:    -1,
						Calories:   3,
					},
				},
			},
			arg: SpoonfulsMap{
				"Butterscotch": 40,
				"Cinnamon":     60,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in.is500Calories(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIngredients_compareMaxScore(t *testing.T) {
	tests := []struct {
		name string
		in   *Ingredients
		arg  int
		want int
	}{
		{
			name: "it doesn't change ingredients.MaxScore if the provided score is less than the current MaxScore",
			in: &Ingredients{
				MaxScore: 183999,
			},
			arg:  100547,
			want: 183999,
		},
		{
			name: "it doesn't change ingredients.MaxScore if the provided score is equal to the current MaxScore",
			in: &Ingredients{
				MaxScore: 451092,
			},
			arg:  451092,
			want: 451092,
		},
		{
			name: "it does change ingredients.MaxScore if the provided score is greater than the current MaxScore",
			in: &Ingredients{
				MaxScore: 183999,
			},
			arg:  342865,
			want: 342865,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := tt.in
			in.compareMaxScore(tt.arg)
			assert.Equal(t, tt.want, in.MaxScore)
		})
	}
}

func TestIngredients_compareMaxScoreWithCalorieLimit(t *testing.T) {
	tests := []struct {
		name string
		in   *Ingredients
		arg  int
		want int
	}{
		{
			name: "it doesn't change ingredients.MaxScoreCalorieLimit if the provided score is less than the current MaxScoreCalorieLimit",
			in: &Ingredients{
				MaxScoreCalorieLimit: 183999,
			},
			arg:  100547,
			want: 183999,
		},
		{
			name: "it doesn't change ingredients.MaxScoreCalorieLimit if the provided score is equal to the current MaxScoreCalorieLimit",
			in: &Ingredients{
				MaxScoreCalorieLimit: 451092,
			},
			arg:  451092,
			want: 451092,
		},
		{
			name: "it does change ingredients.MaxScoreCalorieLimit if the provided score is greater than the current MaxScoreCalorieLimit",
			in: &Ingredients{
				MaxScoreCalorieLimit: 183999,
			},
			arg:  342865,
			want: 342865,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := tt.in
			in.compareMaxScoreWithCalorieLimit(tt.arg)
			assert.Equal(t, tt.want, in.MaxScoreCalorieLimit)
		})
	}
}

func TestIngredients_findOptimumSpoonfuls(t *testing.T) {
	t.Run("finds optimum score for two ingredient list, advent of code example 1", func(t *testing.T) {
		in := Ingredients{
			Ingredients: []*Ingredient{
				{
					Name:       "Butterscotch",
					Capacity:   -1,
					Durability: -2,
					Flavour:    6,
					Texture:    3,
					Calories:   8,
				},
				{
					Name:       "Cinnamon",
					Capacity:   2,
					Durability: 3,
					Flavour:    -2,
					Texture:    -1,
					Calories:   3,
				},
			},
		}

		in.findOptimumSpoonfuls(SpoonfulsMap{}, 1)
		assert.Equal(t, 62842880, in.MaxScore)
		assert.Equal(t, 57600000, in.MaxScoreCalorieLimit)
	})
}
