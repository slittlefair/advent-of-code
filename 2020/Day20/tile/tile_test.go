package tile

import (
	helpers "Advent-of-Code"
	"reflect"
	"testing"
)

func TestTile_RotateTile90(t *testing.T) {
	type fields struct {
		Pixels map[helpers.Coordinate]string
		Height int
		Width  int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[helpers.Coordinate]string
	}{
		{
			name: "test",
			fields: fields{
				Pixels: map[helpers.Coordinate]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
				},
				Width:  1,
				Height: 1,
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
				Pixels: map[helpers.Coordinate]string{
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
				Width:  2,
				Height: 2,
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
				Pixels: tt.fields.Pixels,
				Height: tt.fields.Height,
				Width:  tt.fields.Width,
			}
			tr.RotateTile90()
			if !reflect.DeepEqual(tr.Pixels, tt.want) {
				t.Errorf("got %v, want %v, %d", tr.Pixels, tt.want, tr.Width)
			}
		})
	}
}

func TestTile_FlipTile(t *testing.T) {
	type fields struct {
		Pixels map[helpers.Coordinate]string
		Height int
		Width  int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[helpers.Coordinate]string
	}{
		{
			name: "test",
			fields: fields{
				Pixels: map[helpers.Coordinate]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: ".",
					{X: 0, Y: 1}: ".",
					{X: 1, Y: 1}: ".",
				},
				Width: 1,
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
				Pixels: map[helpers.Coordinate]string{
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
				Width: 2,
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
				Pixels: tt.fields.Pixels,
				Height: tt.fields.Height,
				Width:  tt.fields.Width,
			}
			tr.FlipTile()
			if !reflect.DeepEqual(tr.Pixels, tt.want) {
				t.Errorf("got %v, want %v, %d", tr.Pixels, tt.want, tr.Width)
			}
		})
	}
}
