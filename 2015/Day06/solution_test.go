package main

import (
	"reflect"
	"testing"
)

func TestLights_turnLightsOn(t *testing.T) {
	tests := []struct {
		name string
		l    Lights
		nums []int
		want Lights
	}{
		{
			name: "turns on some lights",
			l: Lights{
				{X: 0, Y: 0}: false,
				{X: 1, Y: 0}: false,
				{X: 2, Y: 0}: false,
				{X: 0, Y: 1}: false,
				{X: 1, Y: 1}: false,
				{X: 2, Y: 1}: false,
			},
			nums: []int{0, 0, 1, 1},
			want: Lights{
				{X: 0, Y: 0}: true,
				{X: 1, Y: 0}: true,
				{X: 2, Y: 0}: false,
				{X: 0, Y: 1}: true,
				{X: 1, Y: 1}: true,
				{X: 2, Y: 1}: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.turnLightsOn(tt.nums)
			if !reflect.DeepEqual(tt.l, tt.want) {
				t.Errorf("Lights.turnLightsOn() = %v, want %v", tt.l, tt.want)
			}
		})
	}
}
