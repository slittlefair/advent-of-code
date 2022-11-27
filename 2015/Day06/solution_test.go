package main

import (
	"Advent-of-Code/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_populateLights(t *testing.T) {
	t.Run("generates Lights correctly", func(t *testing.T) {
		lights := populateLights()
		assert.Equal(t, 1000000, len(lights.Analogue))
		assert.Equal(t, 1000000, len(lights.Digital))
		for x := 0; x < 1000; x++ {
			for y := 0; y < 1000; y++ {
				assert.Zero(t, lights.Analogue[graph.Co{X: x, Y: y}])
				assert.Zero(t, lights.Digital[graph.Co{X: x, Y: y}])
			}
		}
	})
}

func TestLights_turnLightsOn(t *testing.T) {
	t.Run("turns on analogue and digital lights", func(t *testing.T) {
		l := Lights{
			Analogue: map[graph.Co]bool{
				{X: 0, Y: 0}: false,
				{X: 1, Y: 0}: true,
				{X: 2, Y: 0}: false,
				{X: 0, Y: 1}: false,
				{X: 1, Y: 1}: false,
				{X: 2, Y: 1}: false,
			},
			Digital: map[graph.Co]int{
				{X: 0, Y: 0}: 0,
				{X: 1, Y: 0}: 0,
				{X: 2, Y: 0}: 0,
				{X: 0, Y: 1}: 1,
				{X: 1, Y: 1}: 0,
				{X: 2, Y: 1}: 0,
			},
		}
		want := Lights{
			Analogue: map[graph.Co]bool{
				{X: 0, Y: 0}: true,
				{X: 1, Y: 0}: true,
				{X: 2, Y: 0}: false,
				{X: 0, Y: 1}: true,
				{X: 1, Y: 1}: true,
				{X: 2, Y: 1}: false,
			},
			Digital: map[graph.Co]int{
				{X: 0, Y: 0}: 1,
				{X: 1, Y: 0}: 1,
				{X: 2, Y: 0}: 0,
				{X: 0, Y: 1}: 2,
				{X: 1, Y: 1}: 1,
				{X: 2, Y: 1}: 0,
			},
		}
		l.turnLightsOn([]int{0, 0, 1, 1})
		assert.Equal(t, l, want)
	})
}

func TestLights_turnLightsOff(t *testing.T) {
	t.Run("turns off analogue and digital lights", func(t *testing.T) {
		l := Lights{
			Analogue: map[graph.Co]bool{
				{X: 0, Y: 0}: false,
				{X: 1, Y: 0}: true,
				{X: 2, Y: 0}: false,
				{X: 0, Y: 1}: true,
				{X: 1, Y: 1}: false,
				{X: 2, Y: 1}: true,
			},
			Digital: map[graph.Co]int{
				{X: 0, Y: 0}: 1,
				{X: 1, Y: 0}: 0,
				{X: 2, Y: 0}: 3,
				{X: 0, Y: 1}: 1,
				{X: 1, Y: 1}: 2,
				{X: 2, Y: 1}: 0,
			},
		}
		want := Lights{
			Analogue: map[graph.Co]bool{
				{X: 0, Y: 0}: false,
				{X: 1, Y: 0}: false,
				{X: 2, Y: 0}: false,
				{X: 0, Y: 1}: false,
				{X: 1, Y: 1}: false,
				{X: 2, Y: 1}: true,
			},
			Digital: map[graph.Co]int{
				{X: 0, Y: 0}: 0,
				{X: 1, Y: 0}: 0,
				{X: 2, Y: 0}: 3,
				{X: 0, Y: 1}: 0,
				{X: 1, Y: 1}: 1,
				{X: 2, Y: 1}: 0,
			},
		}
		l.turnLightsOff([]int{0, 0, 1, 1})
		assert.Equal(t, l, want)
	})
}

func TestLights_toggleLights(t *testing.T) {
	t.Run("toggles analogue and digital lights", func(t *testing.T) {
		l := Lights{
			Analogue: map[graph.Co]bool{
				{X: 0, Y: 0}: false,
				{X: 1, Y: 0}: true,
				{X: 2, Y: 0}: false,
				{X: 0, Y: 1}: true,
				{X: 1, Y: 1}: false,
				{X: 2, Y: 1}: true,
			},
			Digital: map[graph.Co]int{
				{X: 0, Y: 0}: 1,
				{X: 1, Y: 0}: 0,
				{X: 2, Y: 0}: 3,
				{X: 0, Y: 1}: 1,
				{X: 1, Y: 1}: 2,
				{X: 2, Y: 1}: 0,
			},
		}
		want := Lights{
			Analogue: map[graph.Co]bool{
				{X: 0, Y: 0}: true,
				{X: 1, Y: 0}: false,
				{X: 2, Y: 0}: false,
				{X: 0, Y: 1}: false,
				{X: 1, Y: 1}: true,
				{X: 2, Y: 1}: true,
			},
			Digital: map[graph.Co]int{
				{X: 0, Y: 0}: 3,
				{X: 1, Y: 0}: 2,
				{X: 2, Y: 0}: 3,
				{X: 0, Y: 1}: 3,
				{X: 1, Y: 1}: 4,
				{X: 2, Y: 1}: 0,
			},
		}
		l.toggleLights([]int{0, 0, 1, 1})
		assert.Equal(t, l, want)
	})
}

func TestLights_followInstructions(t *testing.T) {
	type fields struct {
		Analogue map[graph.Co]bool
		Digital  map[graph.Co]int
	}
	tests := []struct {
		name               string
		fields             fields
		input              []string
		want               fields
		errorAssertionFunc assert.ErrorAssertionFunc
	}{
		{
			name: "returns an error if a row of instructions has less than 4 numbers",
			fields: fields{
				Analogue: map[graph.Co]bool{
					{X: 0, Y: 0}: false,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 0}: false,
					{X: 1, Y: 1}: false,
					{X: 3, Y: 3}: false,
				},
				Digital: map[graph.Co]int{
					{X: 0, Y: 0}: 0,
					{X: 0, Y: 1}: 0,
					{X: 1, Y: 0}: 0,
					{X: 1, Y: 1}: 0,
					{X: 3, Y: 3}: 0,
				},
			},
			input: []string{
				"turn on 0,0 through 1,1",
				"toggle 0,1 through 1,1",
				"turn off only 3,3",
				"turn off 0,1 through 0,1",
			},
			want: fields{
				Analogue: map[graph.Co]bool{
					{X: 0, Y: 0}: true,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 0}: true,
					{X: 1, Y: 1}: false,
					{X: 3, Y: 3}: false,
				},
				Digital: map[graph.Co]int{
					{X: 0, Y: 0}: 1,
					{X: 0, Y: 1}: 3,
					{X: 1, Y: 0}: 1,
					{X: 1, Y: 1}: 3,
					{X: 3, Y: 3}: 0,
				},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if a row of instructions has more than 4 numbers",
			fields: fields{
				Analogue: map[graph.Co]bool{
					{X: 0, Y: 0}: false,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 0}: false,
					{X: 1, Y: 1}: false,
					{X: 3, Y: 3}: false,
				},
				Digital: map[graph.Co]int{
					{X: 0, Y: 0}: 0,
					{X: 0, Y: 1}: 0,
					{X: 1, Y: 0}: 0,
					{X: 1, Y: 1}: 0,
					{X: 3, Y: 3}: 0,
				},
			},
			input: []string{
				"turn on 0,0 through 1,1",
				"toggle 0,1 through 1,1 and also 9,9",
				"turn off 3,3 through 3,3",
				"turn off 0,1 through 0,1",
			},
			want: fields{
				Analogue: map[graph.Co]bool{
					{X: 0, Y: 0}: true,
					{X: 0, Y: 1}: true,
					{X: 1, Y: 0}: true,
					{X: 1, Y: 1}: true,
					{X: 3, Y: 3}: false,
				},
				Digital: map[graph.Co]int{
					{X: 0, Y: 0}: 1,
					{X: 0, Y: 1}: 1,
					{X: 1, Y: 0}: 1,
					{X: 1, Y: 1}: 1,
					{X: 3, Y: 3}: 0,
				},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "returns an error if an instruction doesn't contain 'on', 'off' or 'toggle' keywords",
			fields: fields{
				Analogue: map[graph.Co]bool{
					{X: 0, Y: 0}: false,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 0}: false,
					{X: 1, Y: 1}: false,
					{X: 3, Y: 3}: false,
				},
				Digital: map[graph.Co]int{
					{X: 0, Y: 0}: 0,
					{X: 0, Y: 1}: 0,
					{X: 1, Y: 0}: 0,
					{X: 1, Y: 1}: 0,
					{X: 3, Y: 3}: 0,
				},
			},
			input: []string{
				"turn on 0,0 through 1,1",
				"toggle 0,1 through 1,1",
				"turn off 3,3 through 3,3",
				"do something to 0,1 through 0,1",
			},
			want: fields{
				Analogue: map[graph.Co]bool{
					{X: 0, Y: 0}: true,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 0}: true,
					{X: 1, Y: 1}: false,
					{X: 3, Y: 3}: false,
				},
				Digital: map[graph.Co]int{
					{X: 0, Y: 0}: 1,
					{X: 0, Y: 1}: 3,
					{X: 1, Y: 0}: 1,
					{X: 1, Y: 1}: 3,
					{X: 3, Y: 3}: 0,
				},
			},
			errorAssertionFunc: assert.Error,
		},
		{
			name: "follows a valid list of instructions",
			fields: fields{
				Analogue: map[graph.Co]bool{
					{X: 0, Y: 0}: false,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 0}: false,
					{X: 1, Y: 1}: false,
					{X: 3, Y: 3}: false,
				},
				Digital: map[graph.Co]int{
					{X: 0, Y: 0}: 0,
					{X: 0, Y: 1}: 0,
					{X: 1, Y: 0}: 0,
					{X: 1, Y: 1}: 0,
					{X: 3, Y: 3}: 0,
				},
			},
			input: []string{
				"turn on 0,0 through 1,1",
				"toggle 0,1 through 1,1",
				"turn off 3,3 through 3,3",
				"turn off 0,1 through 0,1",
			},
			want: fields{
				Analogue: map[graph.Co]bool{
					{X: 0, Y: 0}: true,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 0}: true,
					{X: 1, Y: 1}: false,
					{X: 3, Y: 3}: false,
				},
				Digital: map[graph.Co]int{
					{X: 0, Y: 0}: 1,
					{X: 0, Y: 1}: 2,
					{X: 1, Y: 0}: 1,
					{X: 1, Y: 1}: 3,
					{X: 3, Y: 3}: 0,
				},
			},
			errorAssertionFunc: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lights{
				Analogue: tt.fields.Analogue,
				Digital:  tt.fields.Digital,
			}
			want := &Lights{
				Analogue: tt.want.Analogue,
				Digital:  tt.want.Digital,
			}
			err := l.followInstructions(tt.input)
			tt.errorAssertionFunc(t, err)
			assert.Equal(t, want, l)
		})
	}
}

func TestLights_countAnalogueBrightness(t *testing.T) {
	t.Run("returns number of analogue lights that are on (count of true)", func(t *testing.T) {
		l := &Lights{
			Analogue: map[graph.Co]bool{
				{X: 0, Y: 0}: true,
				{X: 0, Y: 1}: false,
				{X: 1, Y: 0}: true,
				{X: 1, Y: 1}: false,
				{X: 3, Y: 3}: false,
			},
		}
		got := l.countAnalogueBrightness()
		assert.Equal(t, 2, got)
	})
}

func TestLights_countDigitalBrightness(t *testing.T) {
	t.Run("returns brightness of digital lights that are on (count of values)", func(t *testing.T) {
		l := &Lights{
			Digital: map[graph.Co]int{
				{X: 0, Y: 0}: 1,
				{X: 0, Y: 1}: 2,
				{X: 1, Y: 0}: 1,
				{X: 1, Y: 1}: 3,
				{X: 3, Y: 3}: 0,
			},
		}
		got := l.countDigitalBrightness()
		assert.Equal(t, 7, got)
	})
}
