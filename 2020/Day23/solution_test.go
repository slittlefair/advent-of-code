package main

import (
	"reflect"
	"testing"
)

func TestCups_seenNumsBefore(t *testing.T) {
	tests := []struct {
		name string
		Nums []int
		Seen [][]int
		want bool
	}{
		{
			name: "returns false if nums hasn't been seen before",
			Nums: []int{1, 2, 3},
			Seen: [][]int{
				{1, 2},
				{1, 2, 3, 4},
				{3, 2, 1},
			},
			want: false,
		},
		{
			name: "returns true if nums has been seen before",
			Nums: []int{1, 2, 3},
			Seen: [][]int{
				{3, 2, 1},
				{3, 1, 2},
				{2, 3, 1},
				{1, 3, 2},
				{1, 2, 3},
				{2, 1, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cups{
				Nums: tt.Nums,
			}
			if got := c.seenNumsBefore(tt.Seen); got != tt.want {
				t.Errorf("Cups.seenNumsBefore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCups_indexOfNum(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		Nums []int
		want int
	}{
		{
			name: "returns -1 if arg is not in Nums",
			arg:  0,
			Nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: -1,
		},
		{
			name: "returns 0 if arg is the first int in Nums",
			arg:  1,
			Nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 0,
		},
		{
			name: "returns length of Nums if arg is the last int in Nums",
			arg:  9,
			Nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 8,
		},
		{
			name: "returns the index of Nums if 1 is in the middle of Nums",
			arg:  7,
			Nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cups{
				Nums: tt.Nums,
			}
			got := c.indexOfNum(tt.arg)
			if got != tt.want {
				t.Errorf("Cups.indexOfNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCups_getOrderString(t *testing.T) {
	tests := []struct {
		name    string
		Nums    []int
		want    string
		wantErr bool
	}{
		{
			name:    "returns an error if 1 is not in Nums",
			Nums:    []int{2, 3, 4, 5, 6},
			want:    "",
			wantErr: true,
		},
		{
			name:    "returns a string of ints after 1, simple example",
			Nums:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want:    "23456789",
			wantErr: false,
		},
		{
			name:    "advent of code example 1",
			Nums:    []int{5, 8, 3, 7, 4, 1, 9, 2, 6},
			want:    "92658374",
			wantErr: false,
		},
		{
			name:    "advent of code example 1",
			Nums:    []int{7, 3, 8, 4, 5, 2, 9, 1, 6},
			want:    "67384529",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cups{
				Nums: tt.Nums,
			}
			got, err := c.getOrderString()
			if (err != nil) != tt.wantErr {
				t.Errorf("Cups.getOrderString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Cups.getOrderString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCups_getDestinationIndex(t *testing.T) {
	type fields struct {
		CurrentCup int
		Nums       []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "advent of code example 1",
			fields: fields{
				CurrentCup: 3,
				Nums:       []int{3, 2, 5, 4, 6, 7},
			},
			want: 1,
		},
		{
			name: "advent of code example 2",
			fields: fields{
				CurrentCup: 2,
				Nums:       []int{3, 2, 5, 4, 6, 7},
			},
			want: 5,
		},
		{
			name: "advent of code example 3",
			fields: fields{
				CurrentCup: 5,
				Nums:       []int{3, 2, 5, 8, 9, 1},
			},
			want: 0,
		},
		{
			name: "advent of code example 4",
			fields: fields{
				CurrentCup: 1,
				Nums:       []int{9, 2, 5, 8, 4, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cups{
				CurrentCup: tt.fields.CurrentCup,
				Nums:       tt.fields.Nums,
			}
			got := c.getDestinationIndex()
			if got != tt.want {
				t.Errorf("Cups.getDestinationIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCups_shiftCups(t *testing.T) {
	tests := []struct {
		name string
		Nums []int
		want *Cups
	}{
		{
			name: "shifts the first 3 cups to the end of the slice",
			Nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: &Cups{
				Nums: []int{4, 5, 6, 7, 8, 9, 1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cups{
				Nums: tt.Nums,
			}
			c.shiftCups()
			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("Cups.shiftCups() = %v, want %v", c, tt.want)
			}
		})
	}
}

func TestCups_doPickUp(t *testing.T) {
	type fields struct {
		CurrentCup int
		Nums       []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
		want1  *Cups
	}{
		{
			name: "simple pick up, advent of code example 1",
			fields: fields{
				CurrentCup: 3,
				Nums:       []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			},
			want: []int{8, 9, 1},
			want1: &Cups{
				CurrentCup: 3,
				Nums:       []int{3, 2, 5, 4, 6, 7},
			},
		},
		{
			name: "simple pick up, advent of code example 2",
			fields: fields{
				CurrentCup: 1,
				Nums:       []int{9, 2, 5, 8, 4, 1, 3, 6, 7},
			},
			want: []int{3, 6, 7},
			want1: &Cups{
				CurrentCup: 1,
				Nums:       []int{9, 2, 5, 8, 4, 1},
			},
		},
		{
			name: "pick up with shift of nums, advent of code example 3",
			fields: fields{
				CurrentCup: 9,
				Nums:       []int{7, 2, 5, 8, 4, 1, 9, 3, 6},
			},
			want: []int{3, 6, 7},
			want1: &Cups{
				CurrentCup: 9,
				Nums:       []int{8, 4, 1, 9, 2, 5},
			},
		},
		{
			name: "pick up with shift of nums, advent of code example 4",
			fields: fields{
				CurrentCup: 2,
				Nums:       []int{8, 3, 6, 7, 4, 1, 9, 2, 5},
			},
			want: []int{5, 8, 3},
			want1: &Cups{
				CurrentCup: 2,
				Nums:       []int{7, 4, 1, 9, 2, 6},
			},
		},
		{
			name: "pick up with shift of nums, advent of code example 5",
			fields: fields{
				CurrentCup: 6,
				Nums:       []int{7, 4, 1, 5, 8, 3, 9, 2, 6},
			},
			want: []int{7, 4, 1},
			want1: &Cups{
				CurrentCup: 6,
				Nums:       []int{5, 8, 3, 9, 2, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cups{
				CurrentCup: tt.fields.CurrentCup,
				Nums:       tt.fields.Nums,
			}
			if got := c.doPickUp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cups.doPickUp() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(c, tt.want1) {
				t.Errorf("Cups.doPickUp() = %v, want %v", c, tt.want1)
			}
		})
	}
}

func TestCups_doPutDown(t *testing.T) {
	type fields struct {
		CurrentCup int
		Nums       []int
	}
	tests := []struct {
		name   string
		fields fields
		pickup []int
		want   *Cups
	}{
		{
			name: "advent of code example 1",
			fields: fields{
				CurrentCup: 3,
				Nums:       []int{3, 2, 5, 4, 6, 7},
			},
			pickup: []int{8, 9, 1},
			want: &Cups{
				CurrentCup: 2,
				Nums:       []int{3, 2, 8, 9, 1, 5, 4, 6, 7},
			},
		},
		{
			name: "advent of code example 2",
			fields: fields{
				CurrentCup: 2,
				Nums:       []int{3, 2, 5, 4, 6, 7},
			},
			pickup: []int{8, 9, 1},
			want: &Cups{
				CurrentCup: 5,
				Nums:       []int{3, 2, 5, 4, 6, 7, 8, 9, 1},
			},
		},
		{
			name: "advent of code example 3 after shift",
			fields: fields{
				CurrentCup: 2,
				Nums:       []int{7, 4, 1, 9, 2, 6},
			},
			pickup: []int{5, 8, 3},
			want: &Cups{
				CurrentCup: 6,
				Nums:       []int{7, 4, 1, 5, 8, 3, 9, 2, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cups{
				CurrentCup: tt.fields.CurrentCup,
				Nums:       tt.fields.Nums,
			}
			c.doPutDown(tt.pickup)
			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("Cups.doPutDown() = %v, want %v", c, tt.want)
			}
		})
	}
}

func TestCups_doMove(t *testing.T) {
	type fields struct {
		CurrentCup int
		Nums       []int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Cups
	}{
		{
			name: "advent of code example 1, no shift",
			fields: fields{
				CurrentCup: 3,
				Nums:       []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			},
			want: &Cups{
				CurrentCup: 2,
				Nums:       []int{3, 2, 8, 9, 1, 5, 4, 6, 7},
			},
		},
		{
			name: "advent of code example 2, with shift",
			fields: fields{
				CurrentCup: 9,
				Nums:       []int{7, 2, 5, 8, 4, 1, 9, 3, 6},
			},
			want: &Cups{
				CurrentCup: 2,
				Nums:       []int{8, 3, 6, 7, 4, 1, 9, 2, 5},
			},
		},
		{
			name: "advent of code example 3, with shift",
			fields: fields{
				CurrentCup: 2,
				Nums:       []int{8, 3, 6, 7, 4, 1, 9, 2, 5},
			},
			want: &Cups{
				CurrentCup: 6,
				Nums:       []int{7, 4, 1, 5, 8, 3, 9, 2, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cups{
				CurrentCup: tt.fields.CurrentCup,
				Nums:       tt.fields.Nums,
			}
			c.doMove()
			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("Cups.doMove() = %v, want %v", c, tt.want)
			}
		})
	}
}

func TestCups_playGame(t *testing.T) {
	type fields struct {
		CurrentCup int
		Nums       []int
	}
	tests := []struct {
		name   string
		fields fields
		rounds int
		want   []int
	}{
		{
			name: "advent of code example, 10 rounds",
			fields: fields{
				CurrentCup: 3,
				Nums:       []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			},
			rounds: 10,
			want:   []int{9, 2, 6, 5, 8, 3, 7, 4, 1},
		},
		{
			name: "advent of code example, 100 rounds",
			fields: fields{
				CurrentCup: 3,
				Nums:       []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			},
			rounds: 100,
			want:   []int{9, 1, 6, 7, 3, 8, 4, 5, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cups{
				CurrentCup: tt.fields.CurrentCup,
				Nums:       tt.fields.Nums,
			}
			c.playGame(tt.rounds)
			if !reflect.DeepEqual(c.Nums, tt.want) {
				t.Errorf("Cups.playGame() Nums = %v, want %v", c.Nums, tt.want)
			}
		})
	}
}

func TestCups_populateCupsForPart2(t *testing.T) {
	t.Run("advent of code example 1", func(t *testing.T) {
		c := &Cups{
			CurrentCup: 5,
			Nums:       []int{5, 4, 3, 2, 1},
		}
		c.populateCupsForPart2()
		if !reflect.DeepEqual(c.Nums[:5], []int{5, 4, 3, 2, 1}) {
			t.Errorf("Cups.populateCupsForPart2() = %v, want %v", c.Nums[:5], []int{5, 4, 3, 2, 1})
		}
		for i := 6; i <= 1000000; i++ {
			if c.Nums[i-1] != i {
				t.Errorf("Cups.populateCupsForPart2() = %d at %d, want %d", c.Nums[i-1], i, i-1)
			}
		}
	})

	t.Run("advent of code example 2", func(t *testing.T) {
		c := &Cups{
			CurrentCup: 3,
			Nums:       []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
		}
		c.populateCupsForPart2()
		if !reflect.DeepEqual(c.Nums[:9], []int{3, 8, 9, 1, 2, 5, 4, 6, 7}) {
			t.Errorf("Cups.populateCupsForPart2() = %v, want %v", c.Nums[:9], []int{3, 8, 9, 1, 2, 5, 4, 6, 7})
		}
		for i := 10; i <= 1000000; i++ {
			if c.Nums[i-1] != i {
				t.Errorf("Cups.populateCupsForPart2() = %d at %d, want %d", c.Nums[i-1], i, i-1)
			}
		}
	})
}

func TestCups_productOfCupsToRightOf1(t *testing.T) {
	type fields struct {
		CurrentCup int
		Nums       []int
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name: "returns an error if 1 not in Nums",
			fields: fields{
				CurrentCup: 2,
				Nums:       []int{2, 6, 4, 3, 5, 9, 8, 7},
			},
			want:    -1,
			wantErr: true,
		},
		{
			name: "returns product of the 2 cups to the right of 1 in Nums",
			fields: fields{
				CurrentCup: 2,
				Nums:       []int{2, 6, 4, 1, 3, 5, 9, 8, 7},
			},
			want:    15,
			wantErr: false,
		},
		{
			name: "returns product of the 2 cups to the right of 1 in Nums, but with their index greater than length of Nums",
			fields: fields{
				CurrentCup: 2,
				Nums:       []int{5, 9, 8, 7, 2, 6, 4, 1, 3},
			},
			want:    15,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cups{
				CurrentCup: tt.fields.CurrentCup,
				Nums:       tt.fields.Nums,
			}
			got, err := c.productOfCupsToRightOf1()
			if (err != nil) != tt.wantErr {
				t.Errorf("Cups.productOfCupsToRightOf1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Cups.productOfCupsToRightOf1() = %v, want %v", got, tt.want)
			}
		})
	}
}
