package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	t.Run("returns an error if line of input can't be parsed correctly", func(t *testing.T) {
		input := []string{
			"Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 obsideon. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.",
			"Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.",
		}
		got, err := parseInput(input)
		assert.Error(t, err)
		assert.Nil(t, got)
	})

	t.Run("parses an input into a set of blueprints, advent of code example", func(t *testing.T) {
		input := []string{
			"Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.",
			"Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.",
		}
		want := Blueprints{
			{
				id:            1,
				oreRobot:      Robot{robotType: Ore, oreCost: 4},
				clayRobot:     Robot{robotType: Clay, oreCost: 2},
				obsidianRobot: Robot{robotType: Obsidian, oreCost: 3, clayCost: 14},
				geodeRobot:    Robot{robotType: Geode, oreCost: 2, obsidianCost: 7},
				state:         State{robots: [4]int{1, 0, 0, 0}},
				cache:         Cache{},
			},
			{
				id:            2,
				oreRobot:      Robot{robotType: Ore, oreCost: 2},
				clayRobot:     Robot{robotType: Clay, oreCost: 3},
				obsidianRobot: Robot{robotType: Obsidian, oreCost: 3, clayCost: 8},
				geodeRobot:    Robot{robotType: Geode, oreCost: 3, obsidianCost: 12},
				state:         State{robots: [4]int{1, 0, 0, 0}},
				cache:         Cache{},
			},
		}
		got, err := parseInput(input)
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})
}

func TestCanAffordRobot(t *testing.T) {
	t.Run("returns false if state doesn't have enough ore for robot", func(t *testing.T) {
		s := &State{minute: 24, robots: [4]int{1, 0, 0, 0}, ore: 2, clay: 9, obsidian: 10, geode: 1}
		r := Robot{robotType: Obsidian, oreCost: 3, clayCost: 8}
		got := s.canAffordRobot(r)
		assert.False(t, got)
	})

	t.Run("returns false if state doesn't have enough clay for robot", func(t *testing.T) {
		s := &State{minute: 24, robots: [4]int{1, 0, 0, 0}, ore: 5, clay: 9, obsidian: 10, geode: 1}
		r := Robot{robotType: Obsidian, oreCost: 3, clayCost: 10}
		got := s.canAffordRobot(r)
		assert.False(t, got)
	})

	t.Run("returns false if state doesn't have enough obsidian for robot", func(t *testing.T) {
		s := &State{minute: 24, robots: [4]int{1, 0, 0, 0}, ore: 5, clay: 9, obsidian: 0, geode: 1}
		r := Robot{robotType: Obsidian, oreCost: 3, obsidianCost: 9}
		got := s.canAffordRobot(r)
		assert.False(t, got)
	})

	t.Run("returns true if state has resources for robot", func(t *testing.T) {
		s := &State{minute: 24, robots: [4]int{1, 0, 0, 0}, ore: 5, clay: 9, obsidian: 10, geode: 1}
		r := Robot{robotType: Obsidian, oreCost: 3, clayCost: 8, obsidianCost: 7}
		got := s.canAffordRobot(r)
		assert.True(t, got)
	})
}

func TestPayForRobot(t *testing.T) {
	t.Run("removes reources from state when building a robot", func(t *testing.T) {
		s := &State{minute: 24, robots: [4]int{1, 0, 0, 0}, ore: 5, clay: 9, obsidian: 10, geode: 1}
		r := Robot{robotType: Obsidian, oreCost: 3, clayCost: 8, obsidianCost: 7}
		s.buildRobot(r)
		want := &State{minute: 24, robots: [4]int{1, 0, 1, 0}, ore: 2, clay: 1, obsidian: 3, geode: 1}
		assert.Equal(t, want, s)
	})
}

func TestCollectMinerals(t *testing.T) {
	t.Run("collects minerals from each robot", func(t *testing.T) {
		s := &State{minute: 24, robots: [4]int{9, 3, 6, 4}, ore: 5, clay: 9, obsidian: 10, geode: 1}
		s.collectMinerals()
		want := &State{minute: 24, robots: [4]int{9, 3, 6, 4}, ore: 14, clay: 12, obsidian: 16, geode: 5}
		assert.Equal(t, want, s)
	})
}

func TestSeenState(t *testing.T) {
	cache := Cache{
		{minute: 24, robots: [4]int{1, 0, 0, 0}, ore: 5, clay: 9, obsidian: 10}:           true,
		{minute: 23, robots: [4]int{1, 0, 0, 0}, ore: 5, clay: 9, obsidian: 10, geode: 1}: true,
		{minute: 24, robots: [4]int{1, 2, 0, 0}, ore: 5, clay: 9, obsidian: 10, geode: 1}: true,
		{minute: 24, robots: [4]int{1, 0, 0, 0}, ore: 8, clay: 9, obsidian: 10, geode: 1}: true,
		{minute: 24, robots: [4]int{1, 0, 0, 0}, ore: 5, clay: 9, obsidian: 10, geode: 1}: true,
		{minute: 24, robots: [4]int{1, 0, 0, 0}, ore: 5, clay: 9, obsidian: 1, geode: 1}:  true,
	}

	t.Run("returns true if the given state is in the cache", func(t *testing.T) {
		b := Blueprint{
			id:            1,
			oreRobot:      Robot{robotType: Ore, oreCost: 4},
			clayRobot:     Robot{robotType: Clay, oreCost: 2},
			obsidianRobot: Robot{robotType: Obsidian, oreCost: 3, clayCost: 14},
			geodeRobot:    Robot{robotType: Geode, oreCost: 2, obsidianCost: 7},
			state:         State{minute: 24, robots: [4]int{1, 0, 0, 0}},
			cache:         cache,
		}
		state := State{minute: 24, robots: [4]int{1, 0, 0, 0}, ore: 5, clay: 9, obsidian: 10, geode: 1}
		got := b.seenState(state)
		assert.True(t, got)
	})

	t.Run("returns false if the given state is not in the cache", func(t *testing.T) {
		b := Blueprint{
			id:            1,
			oreRobot:      Robot{robotType: Ore, oreCost: 4},
			clayRobot:     Robot{robotType: Clay, oreCost: 2},
			obsidianRobot: Robot{robotType: Obsidian, oreCost: 3, clayCost: 14},
			geodeRobot:    Robot{robotType: Geode, oreCost: 2, obsidianCost: 7},
			state:         State{minute: 24, robots: [4]int{1, 0, 0, 0}},
			cache:         cache,
		}
		state := State{minute: 2, robots: [4]int{1, 0, 0, 0}, ore: 5, clay: 9, obsidian: 10, geode: 1}
		got := b.seenState(state)
		assert.False(t, got)
	})
}

func TestMaxGeodesCanBuild(t *testing.T) {
	t.Run("TODO", func(t *testing.T) {
		state := State{minute: 20, robots: [4]int{1, 0, 0, 3}, ore: 5, clay: 9, obsidian: 10, geode: 1}
		got := state.maxGeodesCanBuild()
		assert.Equal(t, 19, got)
	})
}

// func TestOneMinute(t *testing.T) {
// 	t.Run("returns highest quality level, advent of code example 1", func(t *testing.T) {
// 		bp := &Blueprint{
// 			id:            1,
// 			oreRobot:      Robot{robotType: Ore, oreCost: 4},
// 			clayRobot:     Robot{robotType: Clay, oreCost: 2},
// 			obsidianRobot: Robot{robotType: Obsidian, oreCost: 3, clayCost: 14},
// 			geodeRobot:    Robot{robotType: Geode, oreCost: 2, obsidianCost: 7},
// 			state:         State{minute: 24, robots: [4]int{1, 0, 0, 0}},
// 		}
// 		bp.oneMinute(bp.state)
// 		assert.Equal(t, 9, bp.QualityLevel)
// 	})

// 	t.Run("returns highest quality level, advent of code example 2", func(t *testing.T) {
// 		bp := &Blueprint{
// 			id:            2,
// 			oreRobot:      Robot{robotType: Ore, oreCost: 2},
// 			clayRobot:     Robot{robotType: Clay, oreCost: 3},
// 			obsidianRobot: Robot{robotType: Obsidian, oreCost: 3, clayCost: 8},
// 			geodeRobot:    Robot{robotType: Geode, oreCost: 3, obsidianCost: 12},
// 			state:         State{minute: 24, robots: [4]int{1, 0, 0, 0}},
// 		}
// 		bp.oneMinute(bp.state)
// 		assert.Equal(t, 24, bp.QualityLevel)
// 	})
// }

// func TestFindSolution(t *testing.T) {
// 	t.Run("returns sum of highest quality levels, advent of code example", func(t *testing.T) {
// 		blueprints := Blueprints{
// 			{
// 				id:            1,
// 				oreRobot:      Robot{robotType: Ore, oreCost: 4},
// 				clayRobot:     Robot{robotType: Clay, oreCost: 2},
// 				obsidianRobot: Robot{robotType: Obsidian, oreCost: 3, clayCost: 14},
// 				geodeRobot:    Robot{robotType: Geode, oreCost: 2, obsidianCost: 7},
// 				state:         State{minute: 24, robots: [4]int{1, 0, 0, 0}},
// 			},
// 			{
// 				id:            2,
// 				oreRobot:      Robot{robotType: Ore, oreCost: 2},
// 				clayRobot:     Robot{robotType: Clay, oreCost: 3},
// 				obsidianRobot: Robot{robotType: Obsidian, oreCost: 3, clayCost: 8},
// 				geodeRobot:    Robot{robotType: Geode, oreCost: 3, obsidianCost: 12},
// 				state:         State{minute: 24, robots: [4]int{1, 0, 0, 0}},
// 			},
// 		}
// 		got := blueprints.findSolution()
// 		assert.Equal(t, 33, got)
// 	})
// }
