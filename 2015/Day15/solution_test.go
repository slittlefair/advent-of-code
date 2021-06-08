package main

import (
	"reflect"
	"testing"
)

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		want Ingredients
	}{
		{
			name: "parses a simple list of 2 ingredients, advent of code example",
			arg: []string{
				"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8",
				"Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3",
			},
			want: Ingredients{
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
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
			if got := tt.in.getScore(tt.arg); got != tt.want {
				t.Errorf("Ingredients.getScore() = %v, want %v", got, tt.want)
			}
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
			if got := tt.s.countSpoonfuls(); got != tt.want {
				t.Errorf("SpoonfulsMap.countSpoonfuls() = %v, want %v", got, tt.want)
			}
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
			if got := tt.in.is500Calories(tt.arg); got != tt.want {
				t.Errorf("Ingredients.is500Calories() = %v, want %v", got, tt.want)
			}
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
			if got := in.MaxScore; got != tt.want {
				t.Errorf("Ingredients.compareMaxScore() = %v, want %v", got, tt.want)
			}
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
			if got := in.MaxScoreCalorieLimit; got != tt.want {
				t.Errorf("Ingredients.compareMaxScoreWithCalorieLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIngredients_findOptimumSpoonfuls(t *testing.T) {
	type args struct {
		spoonfuls SpoonfulsMap
		maxScore  int
		level     int
	}
	tests := []struct {
		name  string
		in    Ingredients
		args  args
		want  int
		want1 int
	}{
		{
			name: "finds optimum score for two ingredient list, advent of code example 1",
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
			args: args{
				spoonfuls: SpoonfulsMap{},
				maxScore:  0,
				level:     1,
			},
			want:  62842880,
			want1: 57600000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := tt.in
			in.findOptimumSpoonfuls(tt.args.spoonfuls, tt.args.level)
			if got := in.MaxScore; got != tt.want {
				t.Errorf("Ingredients.findOptimumSpoonfuls() got = %v, want %v", in.MaxScore, tt.want)
			}
			if got1 := in.MaxScoreCalorieLimit; got1 != tt.want1 {
				t.Errorf("Ingredients.findOptimumSpoonfuls() got1 = %v, want %v", in.MaxScoreCalorieLimit, tt.want1)
			}
		})
	}
}
