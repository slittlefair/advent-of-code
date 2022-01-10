package main

import (
	utils "Advent-of-Code/utils"
	"reflect"
	"testing"
)

func Test_constructLights(t *testing.T) {
	type args struct {
		height int
		width  int
	}
	tests := []struct {
		name string
		args args
		want *Lights
	}{
		{
			name: "constructs a set of lights with the given height and width",
			args: args{
				height: 2,
				width:  4,
			},
			want: &Lights{
				Pixels: map[utils.Co]string{
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructLights(tt.args.height, tt.args.width); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructLights() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLights_followInstruction(t *testing.T) {
	type fields struct {
		Pixels map[utils.Co]string
		Height int
		Width  int
	}
	tests := []struct {
		name    string
		fields  fields
		inst    string
		wantErr bool
		want    *Lights
	}{
		{
			name:    "returns an error if an instruction contains fewer than 2 numbers",
			fields:  fields{},
			inst:    "rotate column x=1 by a few",
			wantErr: true,
			want:    &Lights{},
		},
		{
			name:    "returns an error if an instruction contains more than 2 numbers",
			fields:  fields{},
			inst:    "rect 3x2x1",
			wantErr: true,
			want:    &Lights{},
		},
		{
			name: "it follows a rect instruction",
			fields: fields{
				Pixels: map[utils.Co]string{},
			},
			inst:    "rect 3x2",
			wantErr: false,
			want: &Lights{
				Pixels: map[utils.Co]string{
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
				Pixels: map[utils.Co]string{
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
			inst:    "rotate column x=1 by 1",
			wantErr: false,
			want: &Lights{
				Pixels: map[utils.Co]string{
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
				Pixels: map[utils.Co]string{
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
			inst:    "rotate row y=0 by 6",
			wantErr: false,
			want: &Lights{
				Pixels: map[utils.Co]string{
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
			if err := l.followInstruction(tt.inst); (err != nil) != tt.wantErr {
				t.Errorf("Lights.followInstruction() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(l, tt.want) {
				t.Errorf("Lights.followInstructions() = %v, want %v", l, tt.want)
			}
		})
	}
}

func TestLights_followInstructions(t *testing.T) {
	type fields struct {
		Pixels map[utils.Co]string
		Height int
		Width  int
	}
	tests := []struct {
		name    string
		fields  fields
		input   []string
		wantErr bool
		want    *Lights
	}{
		{
			name:   "returns an error if an instruction contains fewer than 2 numbers",
			fields: fields{},
			input: []string{
				"rotate column x=1 by a few",
			},
			wantErr: true,
			want:    &Lights{},
		},
		{
			name:   "returns an error if an instruction contains more than 2 numbers",
			fields: fields{},
			input: []string{
				"rect 3x2x1",
			},
			wantErr: true,
			want:    &Lights{},
		},
		{
			name: "follows a set of instructions",
			fields: fields{
				Pixels: map[utils.Co]string{
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
			wantErr: false,
			want: &Lights{
				Pixels: map[utils.Co]string{
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
			if err := l.followInstructions(tt.input); (err != nil) != tt.wantErr {
				t.Errorf("Lights.followInstructions() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(l, tt.want) {
				t.Errorf("Lights.followInstructions() = %v, want %v", l, tt.want)
			}
		})
	}
}

func TestLights_countLightsOn(t *testing.T) {
	type fields struct {
		Pixels map[utils.Co]string
		Height int
		Width  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: `returns number of pixels set to "#"`,
			fields: fields{
				Pixels: map[utils.Co]string{
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
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Lights{
				Pixels: tt.fields.Pixels,
				Height: tt.fields.Height,
				Width:  tt.fields.Width,
			}
			if got := l.countLightsOn(); got != tt.want {
				t.Errorf("Lights.countLightsOn() = %v, want %v", got, tt.want)
			}
		})
	}
}
