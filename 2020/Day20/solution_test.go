package main

import (
	helpers "Advent-of-Code"
	"reflect"
	"testing"
)

func TestTile_rotateTile90right(t *testing.T) {
	type fields struct {
		id            string
		pixels        map[helpers.Coordinate]string
		height        int
		width         int
		adjacentTiles map[string]bool
		edges         []string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[helpers.Coordinate]string
	}{
		{
			name: "test",
			fields: fields{
				pixels: map[helpers.Coordinate]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
				},
				width:  1,
				height: 1,
			},
			want: map[helpers.Coordinate]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
			},
		},
		{
			name: "test",
			fields: fields{
				pixels: map[helpers.Coordinate]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: "#",
					{X: 0, Y: 2}: ".",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: ".",
				},
				width:  2,
				height: 2,
			},
			want: map[helpers.Coordinate]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tile{
				id:            tt.fields.id,
				pixels:        tt.fields.pixels,
				height:        tt.fields.height,
				width:         tt.fields.width,
				adjacentTiles: tt.fields.adjacentTiles,
				edges:         tt.fields.edges,
			}
			tr.rotateTile90()
			if !reflect.DeepEqual(tr.pixels, tt.want) {
				t.Errorf("got %v, want %v, %d", tr.pixels, tt.want, tr.width)
			}
		})
	}
}

func TestTile_flipTile(t *testing.T) {
	type fields struct {
		id            string
		pixels        map[helpers.Coordinate]string
		height        int
		width         int
		adjacentTiles map[string]bool
		edges         []string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[helpers.Coordinate]string
	}{
		{
			name: "test",
			fields: fields{
				pixels: map[helpers.Coordinate]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
				},
				width: 1,
			},
			want: map[helpers.Coordinate]string{
				{X: 0, Y: 0}: ".",
				{X: 1, Y: 0}: "#",
				{X: 0, Y: 1}: ".",
				{X: 1, Y: 1}: ".",
			},
		},
		{
			name: "test",
			fields: fields{
				pixels: map[helpers.Coordinate]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
					{X: 2, Y: 1}: "#",
					{X: 0, Y: 2}: "#",
					{X: 1, Y: 2}: ".",
					{X: 2, Y: 2}: ".",
				},
				width: 2,
			},
			want: map[helpers.Coordinate]string{
				{X: 0, Y: 0}: "#",
				{X: 1, Y: 0}: ".",
				{X: 2, Y: 0}: "#",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: ".",
				{X: 2, Y: 1}: ".",
				{X: 0, Y: 2}: ".",
				{X: 1, Y: 2}: ".",
				{X: 2, Y: 2}: "#",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tile{
				id:            tt.fields.id,
				pixels:        tt.fields.pixels,
				height:        tt.fields.height,
				width:         tt.fields.width,
				adjacentTiles: tt.fields.adjacentTiles,
				edges:         tt.fields.edges,
			}
			tr.flipTile()
			if !reflect.DeepEqual(tr.pixels, tt.want) {
				t.Errorf("got %v, want %v, %d", tr.pixels, tt.want, tr.width)
			}
		})
	}
}
