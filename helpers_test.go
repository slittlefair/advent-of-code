package helpers

import (
	"reflect"
	"testing"
)

func TestMin(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{
			name: "returns x if x is less than y",
			x:    23,
			y:    78,
			want: 23,
		},
		{
			name: "returns y if y is less than x",
			x:    9,
			y:    2,
			want: 2,
		},
		{
			name: "returns y if x and y are equal",
			x:    10,
			y:    10,
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.x, tt.y); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{
			name: "returns x if x is greater than y",
			x:    23,
			y:    7,
			want: 23,
		},
		{
			name: "returns y if y is greater than x",
			x:    9,
			y:    25,
			want: 25,
		},
		{
			name: "returns y if x and y are equal",
			x:    10,
			y:    10,
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.x, tt.y); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermutations(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		want [][]string
	}{
		{
			name: "returns all permutations of two element slice",
			arg: []string{
				"Alligator",
				"Broccoli",
			},
			want: [][]string{
				{"Alligator", "Broccoli"},
				{"Broccoli", "Alligator"},
			},
		},
		{
			name: "returns all permutations of three element slice",
			arg: []string{
				"Alligator",
				"Broccoli",
				"Calcium",
			},
			want: [][]string{
				{"Alligator", "Broccoli", "Calcium"},
				{"Alligator", "Calcium", "Broccoli"},
				{"Broccoli", "Calcium", "Alligator"},
				{"Broccoli", "Alligator", "Calcium"},
				{"Calcium", "Broccoli", "Alligator"},
				{"Calcium", "Alligator", "Broccoli"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Permutations(tt.arg)
			if len(got) != len(tt.want) {
				t.Errorf("Permutations() = %v, want %v", got, tt.want)
			}
			for _, g := range got {
				for _, w := range tt.want {
					if reflect.DeepEqual(g, w) {
						goto out
					}
				}
				t.Errorf("Permutations() = %v, want %v", got, tt.want)
			out:
			}
		})
	}
}

func TestIntSlicesAreEqual(t *testing.T) {
	type args struct {
		slice1 []int
		slice2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns false if the given decks aren't of equal length",
			args: args{
				slice1: []int{1, 2, 3, 4},
				slice2: []int{1, 2},
			},
			want: false,
		},
		{
			name: "returns false if the given decks aren't equal",
			args: args{
				slice1: []int{1, 2, 3, 4, 6, 5, 7},
				slice2: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: false,
		},
		{
			name: "returns true if the given decks are equal",
			args: args{
				slice1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				slice2: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntSlicesAreEqual(tt.args.slice1, tt.args.slice2); got != tt.want {
				t.Errorf("isEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		s []int
		i int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		{
			name: "it removes first element from given slice",
			args: args{
				s: []int{1, 2, 3, 4},
				i: 0,
			},
			want:  []int{2, 3, 4},
			want1: []int{1, 2, 3, 4},
		},
		{
			name: "it removes last element from given slice",
			args: args{
				s: []int{1, 2, 3, 4},
				i: 3,
			},
			want:  []int{1, 2, 3},
			want1: []int{1, 2, 3, 4},
		},
		{
			name: "it removes a middle element from given slice",
			args: args{
				s: []int{1, 2, 3, 4},
				i: 1,
			},
			want:  []int{1, 3, 4},
			want1: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove(tt.args.s, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.args.s, tt.want1) {
				t.Errorf("original slice = %v, want %v", tt.args.s, tt.want1)
			}
		})
	}
}

func TestCalculateManhattanDistance(t *testing.T) {
	type args struct {
		co1 Co
		co2 Co
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "returns distance between a coordinate and origin",
			args: args{
				co1: Co{X: 7, Y: 8},
				co2: Co{},
			},
			want: 15,
		},
		{
			name: "returns distance between a positive and negative coordinate",
			args: args{
				co1: Co{X: 9, Y: 1},
				co2: Co{X: -9, Y: -7},
			},
			want: 26,
		},
		{
			name: "returns distance when where difference between the two will be negative",
			args: args{
				co1: Co{X: 1, Y: 1},
				co2: Co{X: 8, Y: 11},
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateManhattanDistance(tt.args.co1, tt.args.co2); got != tt.want {
				t.Errorf("CalculateManhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaesarCipher(t *testing.T) {
	type args struct {
		text     string
		shiftNum int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "applies Caesar Cipher to the given text shifted number of supplied times",
			args: args{
				text:     "qZmt-zixMtkozy-Ivhz-343",
				shiftNum: 343,
			},
			want: "vEry-encRypted-Name-343",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CaesarCipher(tt.args.text, tt.args.shiftNum); got != tt.want {
				t.Errorf("CaesarCipher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMedian(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want float64
	}{
		{
			name: "returns median of odd length input",
			nums: []int{1, 9, 6, 5, 2},
			want: 5,
		},
		{
			name: "returns median of even length input, same numbers",
			nums: []int{1, 3, 6, 5, 2, 3},
			want: 3,
		},
		{
			name: "returns median of even length input, different numbers but int",
			nums: []int{1, 5, 6, 5, 2, 3},
			want: 4,
		},
		{
			name: "returns median of even length input, different numbers but decimal",
			nums: []int{1, 5, 6, 5, 2, 2},
			want: 3.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Median(tt.nums); got != tt.want {
				t.Errorf("Median() = %v, want %v", got, tt.want)
			}
		})
	}
}
