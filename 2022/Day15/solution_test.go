package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var aocExampleSensors = Sensors{
	{X: 2, Y: 18}:  {X: -2, Y: 15},
	{X: 9, Y: 16}:  {X: 10, Y: 16},
	{X: 13, Y: 2}:  {X: 15, Y: 3},
	{X: 12, Y: 14}: {X: 10, Y: 16},
	{X: 10, Y: 20}: {X: 10, Y: 16},
	{X: 14, Y: 17}: {X: 10, Y: 16},
	{X: 8, Y: 7}:   {X: 2, Y: 10},
	{X: 2, Y: 0}:   {X: 2, Y: 10},
	{X: 0, Y: 11}:  {X: 2, Y: 10},
	{X: 20, Y: 14}: {X: 25, Y: 17},
	{X: 17, Y: 20}: {X: 21, Y: 22},
	{X: 16, Y: 7}:  {X: 15, Y: 3},
	{X: 14, Y: 3}:  {X: 15, Y: 3},
	{X: 20, Y: 1}:  {X: 15, Y: 3},
}

func TestParseInput(t *testing.T) {
	t.Run("returns an error if one of the input lines cannot be parsed", func(t *testing.T) {
		input := []string{
			"Sensor at x=2, y=18: closest beacon is at x=-2, y=15",
			"Sensor at x=9, y=16: closest beacon is at x=10, y=16",
			"Sensor at x=13, y=2: closest beacon is at x=15, y=3",
			"Sensor at x=12, y=14: closest beacon is at x=10, y=16",
			"Sensor at x=10, y=20: closest beacon is at x=10, y=16",
			"Sensor at x=14, y=17: closest beacon is at x=10, y=16",
			"Sensor at x=8, y=7: closest beacon is at x=2, y=10",
			"Sensor at x=2, y=0: closest beacon is at x=2, y=10",
			"Sensor at x=0, y=11: closest beacn is at x=2, y=10",
			"Sensor at x=20, y=14: closest beacon is at x=25, y=17",
			"Sensor at x=17, y=20: closest beacon is at x=21, y=22",
			"Sensor at x=16, y=7: closest beacon is at x=15, y=3",
			"Sensor at x=14, y=3: closest beacon is at x=15, y=3",
			"Sensor at x=20, y=1: closest beacon is at x=15, y=3",
		}
		got, err := parseInput(input)
		assert.Nil(t, got)
		assert.Error(t, err)
	})

	t.Run("creates a Sensors map from input, advent of code example", func(t *testing.T) {
		input := []string{
			"Sensor at x=2, y=18: closest beacon is at x=-2, y=15",
			"Sensor at x=9, y=16: closest beacon is at x=10, y=16",
			"Sensor at x=13, y=2: closest beacon is at x=15, y=3",
			"Sensor at x=12, y=14: closest beacon is at x=10, y=16",
			"Sensor at x=10, y=20: closest beacon is at x=10, y=16",
			"Sensor at x=14, y=17: closest beacon is at x=10, y=16",
			"Sensor at x=8, y=7: closest beacon is at x=2, y=10",
			"Sensor at x=2, y=0: closest beacon is at x=2, y=10",
			"Sensor at x=0, y=11: closest beacon is at x=2, y=10",
			"Sensor at x=20, y=14: closest beacon is at x=25, y=17",
			"Sensor at x=17, y=20: closest beacon is at x=21, y=22",
			"Sensor at x=16, y=7: closest beacon is at x=15, y=3",
			"Sensor at x=14, y=3: closest beacon is at x=15, y=3",
			"Sensor at x=20, y=1: closest beacon is at x=15, y=3",
		}
		want := aocExampleSensors
		got, err := parseInput(input)
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})
}

func TestFindTakenSpaces(t *testing.T) {
	t.Run("returns number of taken spaces on a given y axis, advent of code example", func(t *testing.T) {
		got := aocExampleSensors.findTakenSpaces(10)
		assert.Equal(t, 26, got)
	})
}

func TestFindBeacon(t *testing.T) {
	t.Run("returns an error if the missing beacon can't be found", func(t *testing.T) {
		got, err := aocExampleSensors.findBeacon(10)
		assert.Error(t, err)
		assert.Equal(t, -1, got)
	})

	t.Run("returns the missing beacon's value, advent of code example", func(t *testing.T) {
		got, err := aocExampleSensors.findBeacon(20)
		assert.NoError(t, err)
		assert.Equal(t, 56000011, got)
	})
}

func TestFindSolutions(t *testing.T) {
	t.Run("returns correct solutions for advent of code example", func(t *testing.T) {
		input := []string{
			"Sensor at x=2, y=18: closest beacon is at x=-2, y=15",
			"Sensor at x=9, y=16: closest beacon is at x=10, y=16",
			"Sensor at x=13, y=2: closest beacon is at x=15, y=3",
			"Sensor at x=12, y=14: closest beacon is at x=10, y=16",
			"Sensor at x=10, y=20: closest beacon is at x=10, y=16",
			"Sensor at x=14, y=17: closest beacon is at x=10, y=16",
			"Sensor at x=8, y=7: closest beacon is at x=2, y=10",
			"Sensor at x=2, y=0: closest beacon is at x=2, y=10",
			"Sensor at x=0, y=11: closest beacon is at x=2, y=10",
			"Sensor at x=20, y=14: closest beacon is at x=25, y=17",
			"Sensor at x=17, y=20: closest beacon is at x=21, y=22",
			"Sensor at x=16, y=7: closest beacon is at x=15, y=3",
			"Sensor at x=14, y=3: closest beacon is at x=15, y=3",
			"Sensor at x=20, y=1: closest beacon is at x=15, y=3",
		}
		got, got1, err := findSolutions(input, 10)
		assert.NoError(t, err)
		assert.Equal(t, 26, got)
		assert.Equal(t, 56000011, got1)
	})
}
