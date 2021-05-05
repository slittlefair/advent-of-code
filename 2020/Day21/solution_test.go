package main

import (
	"reflect"
	"testing"
)

func TestFoods_parseInput(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  *Foods
	}{
		{
			name: "advent of code example",
			input: []string{
				"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
				"trh fvjkl sbzzf mxmxvkd (contains dairy)",
				"sqjhc fvjkl (contains soy)",
				"sqjhc mxmxvkd sbzzf (contains fish)",
			},
			want: &Foods{
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Foods{
				Contains:        make(map[string]map[string]bool),
				Input:           make(map[string][]map[string]bool),
				SafeIngredients: make(map[string]bool),
			}
			f.parseInput(tt.input)
			if !reflect.DeepEqual(f, tt.want) {
				t.Errorf("f.parseInput() = %v, want %v", f, tt.want)
			}
		})
	}
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
			if !reflect.DeepEqual(f, tt.want) {
				t.Errorf("f.removeSafeFoods() = %v, want %v", f, tt.want)
			}
		})
	}
}

func TestFoods_countSafeIngredients(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "advent of code example",
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
			if got := f.countSafeIngredients(); got != tt.want {
				t.Errorf("Foods.countSafeIngredients() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoods_assignIngredientToAllergen(t *testing.T) {
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
			if !reflect.DeepEqual(f, tt.want) {
				t.Errorf("f.assignIngredientToAllergen() = %v, want %v", f, tt.want)
			}
		})
	}
}

func TestFoods_createDangerousIngredientsList(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "advent of code example",
			want: "mxmxvkd,sqjhc,fvjkl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Foods{
				AllergenToIngredient: map[string]string{
					"dairy": "mxmxvkd",
					"soy":   "fvjkl",
					"fish":  "sqjhc",
				},
			}
			if got := f.createDangerousIngredientsList(); got != tt.want {
				t.Errorf("Foods.createDangerousIngredientsList() = %v, want %v", got, tt.want)
			}
		})
	}
}
