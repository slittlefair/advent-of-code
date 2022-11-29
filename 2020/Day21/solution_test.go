package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoods_parseInput(t *testing.T) {
	t.Run("advent of code example", func(t *testing.T) {
		input := []string{
			"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
			"trh fvjkl sbzzf mxmxvkd (contains dairy)",
			"sqjhc fvjkl (contains soy)",
			"sqjhc mxmxvkd sbzzf (contains fish)",
		}
		want := &Foods{
			Contains: map[string]map[string]bool{
				"dairy": {
					"mxmxvkd": true,
					"kfcds":   true,
					"sqjhc":   true,
					"nhms":    true,
					"trh":     true,
					"fvjkl":   true,
					"sbzzf":   true,
				},
				"fish": {
					"mxmxvkd": true,
					"kfcds":   true,
					"sqjhc":   true,
					"nhms":    true,
					"sbzzf":   true,
				},
				"soy": {
					"sqjhc": true,
					"fvjkl": true,
				},
			},
			Input: map[string][]map[string]bool{
				"dairy": {
					{
						"mxmxvkd": true,
						"kfcds":   true,
						"sqjhc":   true,
						"nhms":    true,
					},
					{
						"trh":     true,
						"fvjkl":   true,
						"sbzzf":   true,
						"mxmxvkd": true,
					},
				},
				"fish": {
					{
						"mxmxvkd": true,
						"kfcds":   true,
						"sqjhc":   true,
						"nhms":    true,
					},
					{
						"sqjhc":   true,
						"mxmxvkd": true,
						"sbzzf":   true,
					},
				},
				"soy": {
					{
						"sqjhc": true,
						"fvjkl": true,
					},
				},
			},
			OriginalInput: [][]string{
				{"mxmxvkd", "kfcds", "sqjhc", "nhms"},
				{"trh", "fvjkl", "sbzzf", "mxmxvkd"},
				{"sqjhc", "fvjkl"},
				{"sqjhc", "mxmxvkd", "sbzzf"},
			},
			SafeIngredients: map[string]bool{
				"mxmxvkd": true,
				"kfcds":   true,
				"sqjhc":   true,
				"nhms":    true,
				"trh":     true,
				"fvjkl":   true,
				"sbzzf":   true,
			},
		}
		f := &Foods{
			Contains:        make(map[string]map[string]bool),
			Input:           make(map[string][]map[string]bool),
			SafeIngredients: make(map[string]bool),
		}
		f.parseInput(input)
		assert.Equal(t, want, f)
	})
}

func TestFoods_removeSafeFoods(t *testing.T) {
	tests := []struct {
		name string
		want *Foods
	}{
		{
			name: "advent of code example",
			want: &Foods{
				Contains: map[string]map[string]bool{
					"dairy": {
						"mxmxvkd": true,
					},
					"fish": {
						"mxmxvkd": true,
						"sqjhc":   true,
					},
					"soy": {
						"sqjhc": true,
						"fvjkl": true,
					},
				},
				SafeIngredients: map[string]bool{
					"kfcds": true,
					"nhms":  true,
					"trh":   true,
					"sbzzf": true,
				},
				Input: map[string][]map[string]bool{
					"dairy": {
						{
							"mxmxvkd": true,
							"kfcds":   true,
							"sqjhc":   true,
							"nhms":    true,
						},
						{
							"trh":     true,
							"fvjkl":   true,
							"sbzzf":   true,
							"mxmxvkd": true,
						},
					},
					"fish": {
						{
							"mxmxvkd": true,
							"kfcds":   true,
							"sqjhc":   true,
							"nhms":    true,
						},
						{
							"sqjhc":   true,
							"mxmxvkd": true,
							"sbzzf":   true,
						},
					},
					"soy": {
						{
							"sqjhc": true,
							"fvjkl": true,
						},
					},
				},
				OriginalInput: [][]string{
					{"mxmxvkd", "kfcds", "sqjhc", "nhms"},
					{"trh", "fvjkl", "sbzzf", "mxmxvkd"},
					{"sqjhc", "fvjkl"},
					{"sqjhc", "mxmxvkd", "sbzzf"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := &Foods{
				Contains: map[string]map[string]bool{
					"dairy": {
						"mxmxvkd": true,
					},
					"fish": {
						"mxmxvkd": true,
						"sqjhc":   true,
					},
					"soy": {
						"sqjhc": true,
						"fvjkl": true,
					},
				},
				SafeIngredients: map[string]bool{
					"kfcds": true,
					"nhms":  true,
					"trh":   true,
					"sbzzf": true,
				},
				Input: map[string][]map[string]bool{
					"dairy": {
						{
							"mxmxvkd": true,
							"kfcds":   true,
							"sqjhc":   true,
							"nhms":    true,
						},
						{
							"trh":     true,
							"fvjkl":   true,
							"sbzzf":   true,
							"mxmxvkd": true,
						},
					},
					"fish": {
						{
							"mxmxvkd": true,
							"kfcds":   true,
							"sqjhc":   true,
							"nhms":    true,
						},
						{
							"sqjhc":   true,
							"mxmxvkd": true,
							"sbzzf":   true,
						},
					},
					"soy": {
						{
							"sqjhc": true,
							"fvjkl": true,
						},
					},
				},
				OriginalInput: [][]string{
					{"mxmxvkd", "kfcds", "sqjhc", "nhms"},
					{"trh", "fvjkl", "sbzzf", "mxmxvkd"},
					{"sqjhc", "fvjkl"},
					{"sqjhc", "mxmxvkd", "sbzzf"},
				},
			}
			f := &Foods{
				Contains: map[string]map[string]bool{
					"dairy": {
						"mxmxvkd": true,
						"kfcds":   true,
						"sqjhc":   true,
						"nhms":    true,
						"trh":     true,
						"fvjkl":   true,
						"sbzzf":   true,
					},
					"fish": {
						"mxmxvkd": true,
						"kfcds":   true,
						"sqjhc":   true,
						"nhms":    true,
						"sbzzf":   true,
					},
					"soy": {
						"sqjhc": true,
						"fvjkl": true,
					},
				},
				Input: map[string][]map[string]bool{
					"dairy": {
						{
							"mxmxvkd": true,
							"kfcds":   true,
							"sqjhc":   true,
							"nhms":    true,
						},
						{
							"trh":     true,
							"fvjkl":   true,
							"sbzzf":   true,
							"mxmxvkd": true,
						},
					},
					"fish": {
						{
							"mxmxvkd": true,
							"kfcds":   true,
							"sqjhc":   true,
							"nhms":    true,
						},
						{
							"sqjhc":   true,
							"mxmxvkd": true,
							"sbzzf":   true,
						},
					},
					"soy": {
						{
							"sqjhc": true,
							"fvjkl": true,
						},
					},
				},
				OriginalInput: [][]string{
					{"mxmxvkd", "kfcds", "sqjhc", "nhms"},
					{"trh", "fvjkl", "sbzzf", "mxmxvkd"},
					{"sqjhc", "fvjkl"},
					{"sqjhc", "mxmxvkd", "sbzzf"},
				},
				SafeIngredients: map[string]bool{
					"mxmxvkd": true,
					"kfcds":   true,
					"sqjhc":   true,
					"nhms":    true,
					"trh":     true,
					"fvjkl":   true,
					"sbzzf":   true,
				},
			}
			f.removeSafeFoods()
			assert.Equal(t, want, f)
		})
	}
}

func TestFoods_countSafeIngredients(t *testing.T) {
	t.Run("advent of code example", func(t *testing.T) {
		f := &Foods{
			Contains: map[string]map[string]bool{
				"dairy": {
					"mxmxvkd": true,
				},
				"fish": {
					"mxmxvkd": true,
					"sqjhc":   true,
				},
				"soy": {
					"sqjhc": true,
					"fvjkl": true,
				},
			},
			SafeIngredients: map[string]bool{
				"kfcds": true,
				"nhms":  true,
				"trh":   true,
				"sbzzf": true,
			},
			Input: map[string][]map[string]bool{
				"dairy": {
					{
						"mxmxvkd": true,
						"kfcds":   true,
						"sqjhc":   true,
						"nhms":    true,
					},
					{
						"trh":     true,
						"fvjkl":   true,
						"sbzzf":   true,
						"mxmxvkd": true,
					},
				},
				"fish": {
					{
						"mxmxvkd": true,
						"kfcds":   true,
						"sqjhc":   true,
						"nhms":    true,
					},
					{
						"sqjhc":   true,
						"mxmxvkd": true,
						"sbzzf":   true,
					},
				},
				"soy": {
					{
						"sqjhc": true,
						"fvjkl": true,
					},
				},
			},
			OriginalInput: [][]string{
				{"mxmxvkd", "kfcds", "sqjhc", "nhms"},
				{"trh", "fvjkl", "sbzzf", "mxmxvkd"},
				{"sqjhc", "fvjkl"},
				{"sqjhc", "mxmxvkd", "sbzzf"},
			},
		}
		got := f.countSafeIngredients()
		assert.Equal(t, 5, got)
	})
}

func TestFoods_assignIngredientToAllergen(t *testing.T) {
	t.Run("advent of code example", func(t *testing.T) {
		want := &Foods{
			Contains: map[string]map[string]bool{
				"dairy": {
					"mxmxvkd": true,
				},
				"fish": {
					"sqjhc": true,
				},
				"soy": {
					"fvjkl": true,
				},
			},
			SafeIngredients: map[string]bool{
				"kfcds": true,
				"nhms":  true,
				"trh":   true,
				"sbzzf": true,
			},
			Input: map[string][]map[string]bool{
				"dairy": {
					{
						"mxmxvkd": true,
						"kfcds":   true,
						"sqjhc":   true,
						"nhms":    true,
					},
					{
						"trh":     true,
						"fvjkl":   true,
						"sbzzf":   true,
						"mxmxvkd": true,
					},
				},
				"fish": {
					{
						"mxmxvkd": true,
						"kfcds":   true,
						"sqjhc":   true,
						"nhms":    true,
					},
					{
						"sqjhc":   true,
						"mxmxvkd": true,
						"sbzzf":   true,
					},
				},
				"soy": {
					{
						"sqjhc": true,
						"fvjkl": true,
					},
				},
			},
			OriginalInput: [][]string{
				{"mxmxvkd", "kfcds", "sqjhc", "nhms"},
				{"trh", "fvjkl", "sbzzf", "mxmxvkd"},
				{"sqjhc", "fvjkl"},
				{"sqjhc", "mxmxvkd", "sbzzf"},
			},
			AllergenToIngredient: map[string]string{
				"dairy": "mxmxvkd",
				"soy":   "fvjkl",
				"fish":  "sqjhc",
			},
		}
		f := &Foods{
			Contains: map[string]map[string]bool{
				"dairy": {
					"mxmxvkd": true,
				},
				"fish": {
					"mxmxvkd": true,
					"sqjhc":   true,
				},
				"soy": {
					"sqjhc": true,
					"fvjkl": true,
				},
			},
			SafeIngredients: map[string]bool{
				"kfcds": true,
				"nhms":  true,
				"trh":   true,
				"sbzzf": true,
			},
			Input: map[string][]map[string]bool{
				"dairy": {
					{
						"mxmxvkd": true,
						"kfcds":   true,
						"sqjhc":   true,
						"nhms":    true,
					},
					{
						"trh":     true,
						"fvjkl":   true,
						"sbzzf":   true,
						"mxmxvkd": true,
					},
				},
				"fish": {
					{
						"mxmxvkd": true,
						"kfcds":   true,
						"sqjhc":   true,
						"nhms":    true,
					},
					{
						"sqjhc":   true,
						"mxmxvkd": true,
						"sbzzf":   true,
					},
				},
				"soy": {
					{
						"sqjhc": true,
						"fvjkl": true,
					},
				},
			},
			OriginalInput: [][]string{
				{"mxmxvkd", "kfcds", "sqjhc", "nhms"},
				{"trh", "fvjkl", "sbzzf", "mxmxvkd"},
				{"sqjhc", "fvjkl"},
				{"sqjhc", "mxmxvkd", "sbzzf"},
			},
			AllergenToIngredient: make(map[string]string),
		}
		f.assignIngredientToAllergen()
		assert.Equal(t, want, f)
	})
}

func TestFoods_createDangerousIngredientsList(t *testing.T) {
	t.Run("advent of code example", func(t *testing.T) {
		f := &Foods{
			AllergenToIngredient: map[string]string{
				"dairy": "mxmxvkd",
				"soy":   "fvjkl",
				"fish":  "sqjhc",
			},
		}
		got := f.createDangerousIngredientsList()
		assert.Equal(t, "mxmxvkd,sqjhc,fvjkl", got)
	})
}
