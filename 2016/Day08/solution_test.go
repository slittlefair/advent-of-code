package main

import (
	"Advent-of-Code/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_constructLights(t *testing.T) {
	t.Run("constructs a set of lights with the given height and width", func(t *testing.T) {
		want := &Lights{
			Pixels: map[graph.Co]string{
				{X: 0, Y: 0}: " ",
				{X: 1, Y: 0}: " ",
				{X: 2, Y: 0}: " ",
				{X: 3, Y: 0}: " ",
				{X: 0, Y: 1}: " ",
				{X: 1, Y: 1}: " ",
				{X: 2, Y: 1}: " ",
				{X: 3, Y: 1}: " ",
			},
			Height: 2,
			Width:  4,
		}
		got := constructLights(2, 4)
		assert.Equal(t, want, got)

	})
}

func TestLights_followInstruction(t *testing.T) {
	type fields struct {
		Pixels map[graph.Co]string
		Height int
		Width  int
	}
	tests := []struct {
		name               string
		fields             fields
		inst               string
		errorAssertionFunc assert.ErrorAssertionFunc
		want               *Lights
	}{
		{
			name:               "returns an error if an instruction contains fewer than 2 numbers",
			fields:             fields{},
			inst:               "rotate column x=1 by a few",
			errorAssertionFunc: assert.Error,
			want:               &Lights{},
		},
		{
			name:               "returns an error if an instruction contains more than 2 numbers",
			fields:             fields{},
			inst:               "rect 3x2x1",
			errorAssertionFunc: assert.Error,
			want:               &Lights{},
		},
		{
			name: "it follows a rect instruction",
			fields: fields{
				Pixels: map[graph.Co]string{},
			},
			inst:               "rect 3x2",
			errorAssertionFunc: assert.NoError,
			want: &Lights{
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: "#",
					{X: 2, Y: 0}: "#",
					{X: 0, Y: 1}: "#",
					{X: 1, Y: 1}: "#",
					{X: 2, Y: 1}: "#",
				},
			},
		},
		{
			name: "it follows a rotate column instruction",
			fields: fields{
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: "#",
					{X: 2, Y: 0}: " ",
					{X: 3, Y: 0}: " ",
					{X: 0, Y: 1}: " ",
					{X: 1, Y: 1}: " ",
					{X: 2, Y: 1}: " ",
					{X: 3, Y: 1}: " ",
				},
				Height: 2,
				Width:  4,
			},
			inst:               "rotate column x=1 by 1",
			errorAssertionFunc: assert.NoError,
			want: &Lights{
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: " ",
					{X: 2, Y: 0}: " ",
					{X: 3, Y: 0}: " ",
					{X: 0, Y: 1}: " ",
					{X: 1, Y: 1}: "#",
					{X: 2, Y: 1}: " ",
					{X: 3, Y: 1}: " ",
				},
				Height: 2,
				Width:  4,
			},
		},
		{
			name: "it follows a rotate row instruction",
			fields: fields{
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: "#",
					{X: 1, Y: 0}: " ",
					{X: 2, Y: 0}: " ",
					{X: 3, Y: 0}: " ",
					{X: 0, Y: 1}: " ",
					{X: 1, Y: 1}: "#",
					{X: 2, Y: 1}: " ",
					{X: 3, Y: 1}: " ",
				},
				Height: 2,
				Width:  4,
			},
			inst:               "rotate row y=0 by 6",
			errorAssertionFunc: assert.NoError,
			want: &Lights{
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: " ",
					{X: 1, Y: 0}: " ",
					{X: 2, Y: 0}: "#",
					{X: 3, Y: 0}: " ",
					{X: 0, Y: 1}: " ",
					{X: 1, Y: 1}: "#",
					{X: 2, Y: 1}: " ",
					{X: 3, Y: 1}: " ",
				},
				Height: 2,
				Width:  4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lights{
				Pixels: tt.fields.Pixels,
				Height: tt.fields.Height,
				Width:  tt.fields.Width,
			}
			err := l.followInstruction(tt.inst)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, l)
		})
	}
}

func TestLights_followInstructions(t *testing.T) {
	type fields struct {
		Pixels map[graph.Co]string
		Height int
		Width  int
	}
	tests := []struct {
		name               string
		fields             fields
		input              []string
		errorAssertionFunc assert.ErrorAssertionFunc
		want               *Lights
	}{
		{
			name:   "returns an error if an instruction contains fewer than 2 numbers",
			fields: fields{},
			input: []string{
				"rotate column x=1 by a few",
			},
			errorAssertionFunc: assert.Error,
			want:               &Lights{},
		},
		{
			name:   "returns an error if an instruction contains more than 2 numbers",
			fields: fields{},
			input: []string{
				"rect 3x2x1",
			},
			errorAssertionFunc: assert.Error,
			want:               &Lights{},
		},
		{
			name: "follows a set of instructions",
			fields: fields{
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: " ",
					{X: 1, Y: 0}: " ",
					{X: 2, Y: 0}: " ",
					{X: 3, Y: 0}: " ",
					{X: 0, Y: 1}: " ",
					{X: 1, Y: 1}: " ",
					{X: 2, Y: 1}: " ",
					{X: 3, Y: 1}: " ",
				},
				Height: 2,
				Width:  4,
			},
			input: []string{
				"rect 2x1",
				"rotate column x=1 by 1",
				"rotate row y=0 by 2",
			},
			errorAssertionFunc: assert.NoError,
			want: &Lights{
				Pixels: map[graph.Co]string{
					{X: 0, Y: 0}: " ",
					{X: 1, Y: 0}: " ",
					{X: 2, Y: 0}: "#",
					{X: 3, Y: 0}: " ",
					{X: 0, Y: 1}: " ",
					{X: 1, Y: 1}: "#",
					{X: 2, Y: 1}: " ",
					{X: 3, Y: 1}: " ",
				},
				Height: 2,
				Width:  4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lights{
				Pixels: tt.fields.Pixels,
				Height: tt.fields.Height,
				Width:  tt.fields.Width,
			}
			err := l.followInstructions(tt.input)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, tt.want, l)
		})
	}
}

func TestLights_countLightsOn(t *testing.T) {
	t.Run(`returns number of pixels set to "#"`, func(t *testing.T) {
		l := Lights{
			Pixels: map[graph.Co]string{
				{X: 0, Y: 0}: " ",
				{X: 1, Y: 0}: "#",
				{X: 2, Y: 0}: " ",
				{X: 3, Y: 0}: "#",
				{X: 0, Y: 1}: "#",
				{X: 1, Y: 1}: " ",
				{X: 2, Y: 1}: "#",
				{X: 3, Y: 1}: " ",
				{X: 0, Y: 2}: "#",
				{X: 1, Y: 2}: "#",
				{X: 2, Y: 2}: "#",
				{X: 3, Y: 2}: " ",
			},
		}
		got := l.countLightsOn()
		assert.Equal(t, 7, got)
	})
}
