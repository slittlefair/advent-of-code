package main

import (
	"reflect"
	"testing"
)

var exampleBagMap = &BagMap{
	"light red": map[string]int{
		"bright white": 1,
		"muted yellow": 2,
	},
	"dark orange": map[string]int{
		"bright white": 3,
		"muted yellow": 4,
	},
	"bright white": map[string]int{
		"shiny gold": 1,
	},
	"muted yellow": map[string]int{
		"shiny gold": 2,
		"faded blue": 9,
	},
	"shiny gold": map[string]int{
		"dark olive":   1,
		"vibrant plum": 2,
	},
	"dark olive": map[string]int{
		"faded blue":   3,
		"dotted black": 4,
	},
	"vibrant plum": map[string]int{
		"faded blue":   5,
		"dotted black": 6,
	},
	"faded blue":   map[string]int{},
	"dotted black": map[string]int{},
}

func TestBagMap_parseBag(t *testing.T) {
	tests := []struct {
		name    string
		bm      *BagMap
		entries []string
		want    *BagMap
	}{
		{
			name: "advent of code example 1",
			bm:   &BagMap{},
			entries: []string{
				"light red bags contain 1 bright white bag, 2 muted yellow bags.",
				"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
				"bright white bags contain 1 shiny gold bag.",
				"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
				"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
				"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
				"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
				"faded blue bags contain no other bags.",
				"dotted black bags contain no other bags.",
			},
			want: exampleBagMap,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bm.parseBag(tt.entries)
			if !reflect.DeepEqual(tt.bm, tt.want) {
				t.Errorf("parseBag() bm = %v, want %v", tt.bm, tt.want)
			}
		})
	}
}

func TestBagMap_traverseBags(t *testing.T) {
	type args struct {
		bag       string
		bagsFound map[string]bool
	}
	tests := []struct {
		name string
		bm   BagMap
		args args
		want map[string]bool
	}{
		{
			name: "generates no bags found if my bag isn't in the map",
			bm:   *exampleBagMap,
			args: args{
				bag:       "awful amber",
				bagsFound: map[string]bool{},
			},
			want: map[string]bool{},
		},
		{
			name: "generates no bags found if my bag isn't contained by any other bags",
			bm:   *exampleBagMap,
			args: args{
				bag:       "light red",
				bagsFound: map[string]bool{},
			},
			want: map[string]bool{},
		},
		{
			name: "advent of code example",
			bm:   *exampleBagMap,
			args: args{
				bag:       "shiny gold",
				bagsFound: map[string]bool{},
			},
			want: map[string]bool{
				"bright white": true,
				"muted yellow": true,
				"dark orange":  true,
				"light red":    true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bm.traverseBags(tt.args.bag, tt.args.bagsFound)
			if !reflect.DeepEqual(tt.args.bagsFound, tt.want) {
				t.Errorf("parseBag() args.bagsFound = %v, want %v", tt.args.bagsFound, tt.want)
			}
		})
	}
}

func TestBagMap_getBagsFound(t *testing.T) {
	tests := []struct {
		name  string
		bm    BagMap
		myBag string
		want  int
	}{
		{
			name:  "returns 0 if bag map is empty",
			bm:    BagMap{},
			myBag: "shiny gold",
			want:  0,
		},
		{
			name:  "returns 0 if my bag is not in bag map",
			bm:    *exampleBagMap,
			myBag: "boring brown",
			want:  0,
		},
		{
			name:  "advent of code example",
			bm:    *exampleBagMap,
			myBag: "shiny gold",
			want:  4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bm.getBagsFound(tt.myBag); got != tt.want {
				t.Errorf("BagMap.getBagsFound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBagMap_countBags(t *testing.T) {
	tests := []struct {
		name string
		bm   BagMap
		bag  string
		want int
	}{
		{
			name: "returns 0 if the given bag isn't in bag map",
			bm:   *exampleBagMap,
			bag:  "rubbish red",
			want: 0,
		},
		{
			name: "returns 0 if the given bag contains no other bags",
			bm:   *exampleBagMap,
			bag:  "dotted black",
			want: 0,
		},
		{
			name: "advent of code example 1",
			bm:   *exampleBagMap,
			bag:  "shiny gold",
			want: 32,
		},
		{
			name: "advent of code example 2",
			bm: BagMap{
				"shiny gold": map[string]int{
					"dark red": 2,
				},
				"dark red": map[string]int{
					"dark orange": 2,
				},
				"dark orange": map[string]int{
					"dark yellow": 2,
				},
				"dark yellow": map[string]int{
					"dark green": 2,
				},
				"dark green": map[string]int{
					"dark blue": 2,
				},
				"dark blue": map[string]int{
					"dark violet": 2,
				},
				"dark violet": map[string]int{},
			},
			bag:  "shiny gold",
			want: 126,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bm.countBags(tt.bag); got != tt.want {
				t.Errorf("BagMap.countBags() = %v, want %v", got, tt.want)
			}
		})
	}
}
