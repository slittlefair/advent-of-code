package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReindeer_move(t *testing.T) {
	tests := []struct {
		name string
		r    *Reindeer
		want *Reindeer
	}{
		{
			name: "moves a reindeer",
			r: &Reindeer{
				TotalDistance: 48,
				Speed:         16,
				CurrentMove:   17,
				Duration:      22,
				IsFlying:      true,
			},
			want: &Reindeer{
				TotalDistance: 64,
				Speed:         16,
				CurrentMove:   18,
				Duration:      22,
				IsFlying:      true,
			},
		},
		{
			name: "moves a reindeer and starts rest",
			r: &Reindeer{
				TotalDistance: 48,
				Speed:         16,
				CurrentMove:   21,
				Duration:      22,
				IsFlying:      true,
			},
			want: &Reindeer{
				TotalDistance: 64,
				Speed:         16,
				CurrentMove:   0,
				Duration:      22,
				IsFlying:      false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			r.move()
			assert.Equal(t, tt.want, r)
		})
	}
}

func TestReindeer_rest(t *testing.T) {
	tests := []struct {
		name string
		r    *Reindeer
		want *Reindeer
	}{
		{
			name: "rests a reindeer",
			r: &Reindeer{
				CurrentRest: 38,
				Rest:        41,
				IsFlying:    false,
			},
			want: &Reindeer{
				CurrentRest: 39,
				Rest:        41,
				IsFlying:    false,
			},
		},
		{
			name: "rests a reindeer and starts moving them",
			r: &Reindeer{
				CurrentRest: 40,
				Rest:        41,
				IsFlying:    false,
			},
			want: &Reindeer{
				CurrentRest: 0,
				Rest:        41,
				IsFlying:    true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			r.rest()
			assert.Equal(t, tt.want, r)
		})
	}
}

func TestRacers_givePoints(t *testing.T) {
	tests := []struct {
		name string
		rc   Racers
		want Racers
	}{
		{
			name: "increases points of the leading reindeer",
			rc: Racers{
				"Comet": {
					Name:          "Comet",
					TotalDistance: 4563,
					TotalPoints:   8799,
				},
				"Donner": {
					Name:          "Donner",
					TotalDistance: 7251,
					TotalPoints:   986,
				},
				"Vixen": {
					Name:          "Vixen",
					TotalDistance: 1998,
					TotalPoints:   2234,
				},
			},
			want: Racers{
				"Comet": {
					Name:          "Comet",
					TotalDistance: 4563,
					TotalPoints:   8799,
				},
				"Donner": {
					Name:          "Donner",
					TotalDistance: 7251,
					TotalPoints:   987,
				},
				"Vixen": {
					Name:          "Vixen",
					TotalDistance: 1998,
					TotalPoints:   2234,
				},
			},
		},
		{
			name: "increases points of the leading reindeers if joint leaders",
			rc: Racers{
				"Comet": {
					Name:          "Comet",
					TotalDistance: 8611,
					TotalPoints:   8799,
				},
				"Donner": {
					Name:          "Donner",
					TotalDistance: 7251,
					TotalPoints:   986,
				},
				"Vixen": {
					Name:          "Vixen",
					TotalDistance: 8611,
					TotalPoints:   2234,
				},
			},
			want: Racers{
				"Comet": {
					Name:          "Comet",
					TotalDistance: 8611,
					TotalPoints:   8800,
				},
				"Donner": {
					Name:          "Donner",
					TotalDistance: 7251,
					TotalPoints:   986,
				},
				"Vixen": {
					Name:          "Vixen",
					TotalDistance: 8611,
					TotalPoints:   2235,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := tt.rc
			rc.givePoints()
			assert.Equal(t, tt.want, rc)
		})
	}
}

func TestRacers_runRace(t *testing.T) {
	t.Run("returns winning distance and points, advent of code example", func(t *testing.T) {
		rc := Racers{
			"Comet": {
				Name:     "Comet",
				Speed:    14,
				Duration: 10,
				Rest:     127,
				IsFlying: true,
			},
			"Dancer": {
				Name:     "Dancer",
				Speed:    16,
				Duration: 11,
				Rest:     162,
				IsFlying: true,
			},
		}
		got, got1 := rc.runRace(1000)
		assert.Equal(t, 1120, got)
		assert.Equal(t, 689, got1)
	})
}

func Test_parseInput(t *testing.T) {
	t.Run("corretcly parses input to create reindeer and racers, advent of code example", func(t *testing.T) {
		arg := []string{
			"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.",
			"Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
		}
		want := Racers{
			"Comet": {
				Name:     "Comet",
				Speed:    14,
				Duration: 10,
				Rest:     127,
				IsFlying: true,
			},
			"Dancer": {
				Name:     "Dancer",
				Speed:    16,
				Duration: 11,
				Rest:     162,
				IsFlying: true,
			},
		}
		got := parseInput(arg)
		assert.Equal(t, want, got)
	})
}
