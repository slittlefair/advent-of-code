package maths

import "testing"

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
