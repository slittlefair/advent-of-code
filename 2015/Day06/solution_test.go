package main

import (
	utils "Advent-of-Code/utils"
	"reflect"
	"testing"
)

func Test_populateLights(t *testing.T) {
	t.Run("generates Lights correctly", func(t *testing.T) {
		lights := populateLights()
		if len(lights.Analogue) != 1000000 {
			t.Errorf("did not generate the correct number of analogue lights, got %d, want 1000000", len(lights.Analogue))
		}
		if len(lights.Digital) != 1000000 {
			t.Errorf("did not generate the correct number of digital lights, got %d, want 1000000", len(lights.Digital))
		}
		for x := 0; x < 1000; x++ {
			for y := 0; y < 1000; y++ {
				if val, ok := lights.Analogue[utils.Co{X: x, Y: y}]; !ok {
					t.Errorf("analogue light %v not in lights map", utils.Co{X: x, Y: y})
				} else if val != false {
					t.Errorf("analogue light %v has the incorrect value, got %t, want false", utils.Co{X: x, Y: y}, val)
				}
				if val, ok := lights.Digital[utils.Co{X: x, Y: y}]; !ok {
					t.Errorf("digital light %v not in lights map", utils.Co{X: x, Y: y})
				} else if val != 0 {
					t.Errorf("digital light %v has the incorrect value, got %d, want 0", utils.Co{X: x, Y: y}, val)
				}
			}
		}
	})
}

func TestLights_turnLightsOn(t *testing.T) {
	type fields struct {
		Analogue map[utils.Co]bool
		Digital  map[utils.Co]int
	}
	tests := []struct {
		name   string
		fields fields
		nums   []int
		want   fields
	}{
		{
			name: "turns on analogue and digital lights",
			fields: fields{
				Analogue: map[utils.Co]bool{
					{X: 0, Y: 0}: false,
					{X: 1, Y: 0}: true,
					{X: 2, Y: 0}: false,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 1}: false,
					{X: 2, Y: 1}: false,
				},
				Digital: map[utils.Co]int{
					{X: 0, Y: 0}: 0,
					{X: 1, Y: 0}: 0,
					{X: 2, Y: 0}: 0,
					{X: 0, Y: 1}: 1,
					{X: 1, Y: 1}: 0,
					{X: 2, Y: 1}: 0,
				},
			},
			nums: []int{0, 0, 1, 1},
			want: fields{
				Analogue: map[utils.Co]bool{
					{X: 0, Y: 0}: true,
					{X: 1, Y: 0}: true,
					{X: 2, Y: 0}: false,
					{X: 0, Y: 1}: true,
					{X: 1, Y: 1}: true,
					{X: 2, Y: 1}: false,
				},
				Digital: map[utils.Co]int{
					{X: 0, Y: 0}: 1,
					{X: 1, Y: 0}: 1,
					{X: 2, Y: 0}: 0,
					{X: 0, Y: 1}: 2,
					{X: 1, Y: 1}: 1,
					{X: 2, Y: 1}: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Lights{
				Analogue: tt.fields.Analogue,
				Digital:  tt.fields.Digital,
			}
			want := Lights{
				Analogue: tt.want.Analogue,
				Digital:  tt.want.Digital,
			}
			l.turnLightsOn(tt.nums)
			if !reflect.DeepEqual(l, want) {
				t.Errorf("Lights.turnLightsOn() = %v, want %v", l, want)
			}
		})
	}
}

func TestLights_turnLightsOff(t *testing.T) {
	type fields struct {
		Analogue map[utils.Co]bool
		Digital  map[utils.Co]int
	}
	tests := []struct {
		name   string
		fields fields
		nums   []int
		want   fields
	}{
		{
			name: "turns off analogue and digital lights",
			fields: fields{
				Analogue: map[utils.Co]bool{
					{X: 0, Y: 0}: false,
					{X: 1, Y: 0}: true,
					{X: 2, Y: 0}: false,
					{X: 0, Y: 1}: true,
					{X: 1, Y: 1}: false,
					{X: 2, Y: 1}: true,
				},
				Digital: map[utils.Co]int{
					{X: 0, Y: 0}: 1,
					{X: 1, Y: 0}: 0,
					{X: 2, Y: 0}: 3,
					{X: 0, Y: 1}: 1,
					{X: 1, Y: 1}: 2,
					{X: 2, Y: 1}: 0,
				},
			},
			nums: []int{0, 0, 1, 1},
			want: fields{
				Analogue: map[utils.Co]bool{
					{X: 0, Y: 0}: false,
					{X: 1, Y: 0}: false,
					{X: 2, Y: 0}: false,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 1}: false,
					{X: 2, Y: 1}: true,
				},
				Digital: map[utils.Co]int{
					{X: 0, Y: 0}: 0,
					{X: 1, Y: 0}: 0,
					{X: 2, Y: 0}: 3,
					{X: 0, Y: 1}: 0,
					{X: 1, Y: 1}: 1,
					{X: 2, Y: 1}: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Lights{
				Analogue: tt.fields.Analogue,
				Digital:  tt.fields.Digital,
			}
			want := Lights{
				Analogue: tt.want.Analogue,
				Digital:  tt.want.Digital,
			}
			l.turnLightsOff(tt.nums)
			if !reflect.DeepEqual(l, want) {
				t.Errorf("Lights.turnLightsOff() = %v, want %v", l, want)
			}
		})
	}
}

func TestLights_toggleLights(t *testing.T) {
	type fields struct {
		Analogue map[utils.Co]bool
		Digital  map[utils.Co]int
	}
	tests := []struct {
		name   string
		fields fields
		nums   []int
		want   fields
	}{
		{
			name: "toggles analogue and digital lights",
			fields: fields{
				Analogue: map[utils.Co]bool{
					{X: 0, Y: 0}: false,
					{X: 1, Y: 0}: true,
					{X: 2, Y: 0}: false,
					{X: 0, Y: 1}: true,
					{X: 1, Y: 1}: false,
					{X: 2, Y: 1}: true,
				},
				Digital: map[utils.Co]int{
					{X: 0, Y: 0}: 1,
					{X: 1, Y: 0}: 0,
					{X: 2, Y: 0}: 3,
					{X: 0, Y: 1}: 1,
					{X: 1, Y: 1}: 2,
					{X: 2, Y: 1}: 0,
				},
			},
			nums: []int{0, 0, 1, 1},
			want: fields{
				Analogue: map[utils.Co]bool{
					{X: 0, Y: 0}: true,
					{X: 1, Y: 0}: false,
					{X: 2, Y: 0}: false,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 1}: true,
					{X: 2, Y: 1}: true,
				},
				Digital: map[utils.Co]int{
					{X: 0, Y: 0}: 3,
					{X: 1, Y: 0}: 2,
					{X: 2, Y: 0}: 3,
					{X: 0, Y: 1}: 3,
					{X: 1, Y: 1}: 4,
					{X: 2, Y: 1}: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Lights{
				Analogue: tt.fields.Analogue,
				Digital:  tt.fields.Digital,
			}
			want := Lights{
				Analogue: tt.want.Analogue,
				Digital:  tt.want.Digital,
			}
			l.toggleLights(tt.nums)
			if !reflect.DeepEqual(l, want) {
				t.Errorf("Lights.toggleLights() = %v, want %v", l, want)
			}
		})
	}
}

func TestLights_followInstructions(t *testing.T) {
	type fields struct {
		Analogue map[utils.Co]bool
		Digital  map[utils.Co]int
	}
	tests := []struct {
		name    string
		fields  fields
		input   []string
		want    fields
		wantErr bool
	}{
		{
			name: "returns an error if a row of instructions has less than 4 numbers",
			fields: fields{
				Analogue: map[utils.Co]bool{
					{X: 0, Y: 0}: false,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 0}: false,
					{X: 1, Y: 1}: false,
					{X: 3, Y: 3}: false,
				},
				Digital: map[utils.Co]int{
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
				Analogue: map[utils.Co]bool{
					{X: 0, Y: 0}: true,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 0}: true,
					{X: 1, Y: 1}: false,
					{X: 3, Y: 3}: false,
				},
				Digital: map[utils.Co]int{
					{X: 0, Y: 0}: 1,
					{X: 0, Y: 1}: 3,
					{X: 1, Y: 0}: 1,
					{X: 1, Y: 1}: 3,
					{X: 3, Y: 3}: 0,
				},
			},
			wantErr: true,
		},
		{
			name: "returns an error if a row of instructions has more than 4 numbers",
			fields: fields{
				Analogue: map[utils.Co]bool{
					{X: 0, Y: 0}: false,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 0}: false,
					{X: 1, Y: 1}: false,
					{X: 3, Y: 3}: false,
				},
				Digital: map[utils.Co]int{
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
				Analogue: map[utils.Co]bool{
					{X: 0, Y: 0}: true,
					{X: 0, Y: 1}: true,
					{X: 1, Y: 0}: true,
					{X: 1, Y: 1}: true,
					{X: 3, Y: 3}: false,
				},
				Digital: map[utils.Co]int{
					{X: 0, Y: 0}: 1,
					{X: 0, Y: 1}: 1,
					{X: 1, Y: 0}: 1,
					{X: 1, Y: 1}: 1,
					{X: 3, Y: 3}: 0,
				},
			},
			wantErr: true,
		},
		{
			name: "returns an error if an instruction doesn't contain 'on', 'off' or 'toggle' keywords",
			fields: fields{
				Analogue: map[utils.Co]bool{
					{X: 0, Y: 0}: false,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 0}: false,
					{X: 1, Y: 1}: false,
					{X: 3, Y: 3}: false,
				},
				Digital: map[utils.Co]int{
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
				Analogue: map[utils.Co]bool{
					{X: 0, Y: 0}: true,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 0}: true,
					{X: 1, Y: 1}: false,
					{X: 3, Y: 3}: false,
				},
				Digital: map[utils.Co]int{
					{X: 0, Y: 0}: 1,
					{X: 0, Y: 1}: 3,
					{X: 1, Y: 0}: 1,
					{X: 1, Y: 1}: 3,
					{X: 3, Y: 3}: 0,
				},
			},
			wantErr: true,
		},
		{
			name: "follows a valid list of instructions",
			fields: fields{
				Analogue: map[utils.Co]bool{
					{X: 0, Y: 0}: false,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 0}: false,
					{X: 1, Y: 1}: false,
					{X: 3, Y: 3}: false,
				},
				Digital: map[utils.Co]int{
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
				Analogue: map[utils.Co]bool{
					{X: 0, Y: 0}: true,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 0}: true,
					{X: 1, Y: 1}: false,
					{X: 3, Y: 3}: false,
				},
				Digital: map[utils.Co]int{
					{X: 0, Y: 0}: 1,
					{X: 0, Y: 1}: 2,
					{X: 1, Y: 0}: 1,
					{X: 1, Y: 1}: 3,
					{X: 3, Y: 3}: 0,
				},
			},
			wantErr: false,
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
			if (err != nil) != tt.wantErr {
				t.Errorf("Lights.followInstructions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(l, want) {
				t.Errorf("part1() = %v, want %v", l, want)
			}
		})
	}
}

func TestLights_countAnalogueBrightness(t *testing.T) {
	type fields struct {
		Analogue map[utils.Co]bool
		Digital  map[utils.Co]int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "returns number of analogue lights that are on (count of true)",
			fields: fields{
				Analogue: map[utils.Co]bool{
					{X: 0, Y: 0}: true,
					{X: 0, Y: 1}: false,
					{X: 1, Y: 0}: true,
					{X: 1, Y: 1}: false,
					{X: 3, Y: 3}: false,
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lights{
				Analogue: tt.fields.Analogue,
				Digital:  tt.fields.Digital,
			}
			if got := l.countAnalogueBrightness(); got != tt.want {
				t.Errorf("Lights.countAnalogueBrightness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLights_countDigitalBrightness(t *testing.T) {
	type fields struct {
		Analogue map[utils.Co]bool
		Digital  map[utils.Co]int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "returns brightness of digital lights that are on (count of values)",
			fields: fields{
				Digital: map[utils.Co]int{
					{X: 0, Y: 0}: 1,
					{X: 0, Y: 1}: 2,
					{X: 1, Y: 0}: 1,
					{X: 1, Y: 1}: 3,
					{X: 3, Y: 3}: 0,
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lights{
				Analogue: tt.fields.Analogue,
				Digital:  tt.fields.Digital,
			}
			if got := l.countDigitalBrightness(); got != tt.want {
				t.Errorf("Lights.countDigitalBrightness() = %v, want %v", got, tt.want)
			}
		})
	}
}
