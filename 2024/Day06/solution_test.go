package main

import (
	"Advent-of-Code/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseInput(t *testing.T) {
	t.Run("returns a floor for a given input", func(t *testing.T) {
		input := []string{
			"....#.....",
			".........#",
			"..........",
			"..#.......",
			".......#..",
			"..........",
			".#..^.....",
			"........#.",
			"#.........",
			"......#...",
		}
		expectedFloor := &Floor{
			guard: &Guard{
				co:  graph.Co{X: 4, Y: 6},
				dir: 0,
			},
			grid: &graph.Grid{
				MaxX: 9,
				MaxY: 9,
				Graph: graph.Graph{
					{X: 4, Y: 0}: "#",
					{X: 9, Y: 1}: "#",
					{X: 2, Y: 3}: "#",
					{X: 7, Y: 4}: "#",
					{X: 1, Y: 6}: "#",
					{X: 8, Y: 7}: "#",
					{X: 0, Y: 8}: "#",
					{X: 6, Y: 9}: "#",
				},
			},
			visitedSteps: map[graph.Co]map[int]bool{
				{X: 4, Y: 6}: {
					0: true,
				},
			},
		}
		assert.Equal(t, expectedFloor, parseInput(input))
	})
}

func Test_runPatrol(t *testing.T) {
	t.Run("runs a patrol where the guard exits, changing the floor", func(t *testing.T) {
		floor := &Floor{
			guard: &Guard{
				co:  graph.Co{X: 4, Y: 6},
				dir: 0,
			},
			grid: &graph.Grid{
				MaxX: 9,
				MaxY: 9,
				Graph: graph.Graph{
					{X: 4, Y: 0}: "#",
					{X: 9, Y: 1}: "#",
					{X: 2, Y: 3}: "#",
					{X: 7, Y: 4}: "#",
					{X: 1, Y: 6}: "#",
					{X: 8, Y: 7}: "#",
					{X: 0, Y: 8}: "#",
					{X: 6, Y: 9}: "#",
				},
			},
			visitedSteps: map[graph.Co]map[int]bool{
				{X: 4, Y: 6}: {
					0: true,
				},
			},
		}

		stuckInloop := floor.runPatrol()
		assert.False(t, stuckInloop)
		assert.Equal(t, 41, len(floor.visitedSteps))
	})

	t.Run("runs a patrol where the guard is stuck in a loop", func(t *testing.T) {
		floor := &Floor{
			guard: &Guard{
				co:  graph.Co{X: 4, Y: 6},
				dir: 0,
			},
			grid: &graph.Grid{
				MaxX: 9,
				MaxY: 9,
				Graph: graph.Graph{
					{X: 4, Y: 0}: "#",
					{X: 9, Y: 1}: "#",
					{X: 2, Y: 3}: "#",
					{X: 7, Y: 4}: "#",
					{X: 1, Y: 6}: "#",
					{X: 8, Y: 7}: "#",
					{X: 0, Y: 8}: "#",
					{X: 6, Y: 9}: "#",
					{X: 3, Y: 6}: "#",
				},
			},
			visitedSteps: map[graph.Co]map[int]bool{
				{X: 4, Y: 6}: {
					0: true,
				},
			},
		}

		assert.True(t, floor.runPatrol())
	})
}

func assertVisitedStepsCount(t *testing.T, floor *Floor, expected int) {
	l := 0
	for _, v := range floor.visitedSteps {
		l += len(v)
	}
	assert.Equal(t, expected, l)
}

func Test_step(t *testing.T) {
	t.Run("turns the guard if they are facing an obstacle", func(t *testing.T) {
		floor := &Floor{
			guard: &Guard{
				co:  graph.Co{X: 4, Y: 1},
				dir: 0,
			},
			grid: &graph.Grid{
				MaxX: 9,
				MaxY: 9,
				Graph: graph.Graph{
					{X: 4, Y: 0}: "#",
					{X: 9, Y: 1}: "#",
					{X: 2, Y: 3}: "#",
					{X: 7, Y: 4}: "#",
					{X: 1, Y: 6}: "#",
					{X: 8, Y: 7}: "#",
					{X: 0, Y: 8}: "#",
					{X: 6, Y: 9}: "#",
				},
			},
			visitedSteps: map[graph.Co]map[int]bool{
				{X: 4, Y: 6}: {0: true},
				{X: 4, Y: 5}: {0: true},
				{X: 4, Y: 4}: {0: true},
				{X: 4, Y: 3}: {0: true},
				{X: 4, Y: 2}: {0: true},
				{X: 4, Y: 1}: {0: true},
			},
		}
		inBounds, visitedBefore := floor.step()
		assert.True(t, inBounds)
		assert.False(t, visitedBefore)
		assertVisitedStepsCount(t, floor, 6)
		assert.Equal(t, &Guard{co: graph.Co{X: 4, Y: 1}, dir: 1}, floor.guard)
		assert.Equal(t, 6, len(floor.visitedSteps))
	})

	t.Run("moves the guard one step forwards to a space not yet visited, staying in bounds", func(t *testing.T) {
		floor := &Floor{
			guard: &Guard{
				co:  graph.Co{X: 4, Y: 2},
				dir: 0,
			},
			grid: &graph.Grid{
				MaxX: 9,
				MaxY: 9,
				Graph: graph.Graph{
					{X: 4, Y: 0}: "#",
					{X: 9, Y: 1}: "#",
					{X: 2, Y: 3}: "#",
					{X: 7, Y: 4}: "#",
					{X: 1, Y: 6}: "#",
					{X: 8, Y: 7}: "#",
					{X: 0, Y: 8}: "#",
					{X: 6, Y: 9}: "#",
				},
			},
			visitedSteps: map[graph.Co]map[int]bool{
				{X: 4, Y: 6}: {0: true},
				{X: 4, Y: 5}: {0: true},
				{X: 4, Y: 4}: {0: true},
				{X: 4, Y: 3}: {0: true},
				{X: 4, Y: 2}: {0: true},
			},
		}
		inBounds, visitedBefore := floor.step()
		assert.True(t, inBounds)
		assert.False(t, visitedBefore)
		assert.Equal(t, &Guard{co: graph.Co{X: 4, Y: 1}, dir: 0}, floor.guard)
		assertVisitedStepsCount(t, floor, 6)
		v := floor.visitedSteps[graph.Co{X: 4, Y: 1}]
		assert.Equal(t, map[int]bool{0: true}, v)
	})

	t.Run("moves the guard one step forwards to a space visited but different direction", func(t *testing.T) {
		floor := &Floor{
			guard: &Guard{
				co:  graph.Co{X: 4, Y: 2},
				dir: 1,
			},
			grid: &graph.Grid{
				MaxX: 9,
				MaxY: 9,
				Graph: graph.Graph{
					{X: 4, Y: 0}: "#",
					{X: 9, Y: 1}: "#",
					{X: 2, Y: 3}: "#",
					{X: 7, Y: 4}: "#",
					{X: 1, Y: 6}: "#",
					{X: 8, Y: 7}: "#",
					{X: 0, Y: 8}: "#",
					{X: 6, Y: 9}: "#",
				},
			},
			visitedSteps: map[graph.Co]map[int]bool{
				{X: 4, Y: 6}: {0: true},
				{X: 4, Y: 5}: {0: true},
				{X: 4, Y: 4}: {0: true},
				{X: 4, Y: 3}: {0: true},
				{X: 4, Y: 2}: {0: true},
				{X: 5, Y: 2}: {0: true},
			},
		}
		inBounds, visitedBefore := floor.step()
		assert.True(t, inBounds)
		assert.False(t, visitedBefore)
		assert.Equal(t, &Guard{co: graph.Co{X: 5, Y: 2}, dir: 1}, floor.guard)
		assertVisitedStepsCount(t, floor, 7)
		v := floor.visitedSteps[graph.Co{X: 5, Y: 2}]
		assert.Equal(t, map[int]bool{0: true, 1: true}, v)
	})

	t.Run("moves the guard outside of the floor", func(t *testing.T) {
		floor := &Floor{
			guard: &Guard{
				co:  graph.Co{X: 9, Y: 9},
				dir: 2,
			},
			grid: &graph.Grid{
				MaxX: 9,
				MaxY: 9,
				Graph: graph.Graph{
					{X: 4, Y: 0}: "#",
					{X: 9, Y: 1}: "#",
					{X: 2, Y: 3}: "#",
					{X: 7, Y: 4}: "#",
					{X: 1, Y: 6}: "#",
					{X: 8, Y: 7}: "#",
					{X: 0, Y: 8}: "#",
					{X: 6, Y: 9}: "#",
				},
			},
			visitedSteps: map[graph.Co]map[int]bool{
				{X: 4, Y: 6}: {0: true},
				{X: 4, Y: 5}: {0: true},
				{X: 4, Y: 4}: {0: true},
				{X: 4, Y: 3}: {0: true},
				{X: 4, Y: 2}: {0: true},
				{X: 5, Y: 2}: {0: true},
			},
		}
		inBounds, visitedBefore := floor.step()
		assert.False(t, inBounds)
		assert.False(t, visitedBefore)
		assertVisitedStepsCount(t, floor, 6)
	})

	t.Run("returns if the guard is in a space and position they ave been in before", func(t *testing.T) {
		floor := &Floor{
			guard: &Guard{
				co:  graph.Co{X: 4, Y: 2},
				dir: 1,
			},
			grid: &graph.Grid{
				MaxX: 9,
				MaxY: 9,
				Graph: graph.Graph{
					{X: 4, Y: 0}: "#",
					{X: 9, Y: 1}: "#",
					{X: 2, Y: 3}: "#",
					{X: 7, Y: 4}: "#",
					{X: 1, Y: 6}: "#",
					{X: 8, Y: 7}: "#",
					{X: 0, Y: 8}: "#",
					{X: 6, Y: 9}: "#",
				},
			},
			visitedSteps: map[graph.Co]map[int]bool{
				{X: 4, Y: 6}: {0: true},
				{X: 4, Y: 5}: {0: true},
				{X: 4, Y: 4}: {0: true},
				{X: 4, Y: 3}: {0: true},
				{X: 4, Y: 2}: {0: true},
				{X: 5, Y: 2}: {0: true, 1: true},
			},
		}
		inBounds, visitedBefore := floor.step()
		assert.True(t, inBounds)
		assert.True(t, visitedBefore)
		assertVisitedStepsCount(t, floor, 7)
	})
}

func test_runPatrol(t *testing.T) {
	t.Run("runs a patrol for a given grid 1, leaving the grid", func(t *testing.T) {
		floor := &Floor{
			guard: &Guard{
				co:  graph.Co{X: 4, Y: 6},
				dir: 0,
			},
			grid: &graph.Grid{
				MaxX: 9,
				MaxY: 9,
				Graph: graph.Graph{
					{X: 4, Y: 0}: "#",
					{X: 9, Y: 1}: "#",
					{X: 2, Y: 3}: "#",
					{X: 7, Y: 4}: "#",
					{X: 1, Y: 6}: "#",
					{X: 8, Y: 7}: "#",
					{X: 0, Y: 8}: "#",
					{X: 6, Y: 9}: "#",
				},
			},
			visitedSteps: map[graph.Co]map[int]bool{
				{X: 4, Y: 6}: {
					0: true,
				},
			},
		}
		stuckInLoop := floor.runPatrol()
		assert.False(t, stuckInLoop)
		assert.Len(t, 41, len(floor.visitedSteps))
	})

	t.Run("runs a patrol for a given grid 2, leaving the grid", func(t *testing.T) {
		floor := &Floor{
			guard: &Guard{
				co:  graph.Co{X: 4, Y: 1},
				dir: 0,
			},
			grid: &graph.Grid{
				MaxX: 9,
				MaxY: 9,
				Graph: graph.Graph{
					{X: 4, Y: 0}: "#",
					{X: 9, Y: 1}: "#",
					{X: 2, Y: 3}: "#",
					{X: 7, Y: 4}: "#",
					{X: 1, Y: 6}: "#",
					{X: 8, Y: 7}: "#",
					{X: 0, Y: 8}: "#",
					{X: 6, Y: 9}: "#",
					{X: 4, Y: 4}: "#",
				},
			},
			visitedSteps: map[graph.Co]map[int]bool{
				{X: 4, Y: 6}: {
					0: true,
				},
			},
		}
		stuckInLoop := floor.runPatrol()
		assert.False(t, stuckInLoop)
		assert.Len(t, 10, len(floor.visitedSteps))
	})

	t.Run("runs a patrol for a given grid 3, getting stuck in a loop", func(t *testing.T) {
		floor := &Floor{
			guard: &Guard{
				co:  graph.Co{X: 4, Y: 1},
				dir: 0,
			},
			grid: &graph.Grid{
				MaxX: 9,
				MaxY: 9,
				Graph: graph.Graph{
					{X: 4, Y: 0}: "#",
					{X: 9, Y: 1}: "#",
					{X: 2, Y: 3}: "#",
					{X: 7, Y: 4}: "#",
					{X: 1, Y: 6}: "#",
					{X: 8, Y: 7}: "#",
					{X: 0, Y: 8}: "#",
					{X: 6, Y: 9}: "#",
					{X: 3, Y: 6}: "#",
				},
			},
			visitedSteps: map[graph.Co]map[int]bool{
				{X: 4, Y: 6}: {
					0: true,
				},
			},
		}
		stuckInLoop := floor.runPatrol()
		assert.True(t, stuckInLoop)
	})
}

func Test_resetFloor(t *testing.T) {
	t.Run("resets a floor to the values provided", func(t *testing.T) {
		floor := &Floor{
			guard: &Guard{
				co:  graph.Co{X: 4, Y: 6},
				dir: 0,
			},
			grid: &graph.Grid{
				MaxX: 9,
				MaxY: 9,
				Graph: graph.Graph{
					{X: 4, Y: 0}: "#",
					{X: 9, Y: 1}: "#",
					{X: 2, Y: 3}: "#",
					{X: 7, Y: 4}: "#",
					{X: 1, Y: 6}: "#",
					{X: 8, Y: 7}: "#",
					{X: 0, Y: 8}: "#",
					{X: 6, Y: 9}: "#",
					{X: 3, Y: 6}: "#",
				},
			},
			visitedSteps: map[graph.Co]map[int]bool{
				{X: 4, Y: 6}: {
					0: true,
					1: true,
				},
				{X: 9, Y: 0}: {
					3: true,
				},
				{X: 1, Y: 2}: {
					0: true,
					1: true,
					3: true,
					2: true,
				},
				{X: 4, Y: 7}: {
					3: true,
				},
			},
		}

		obstacle := graph.Co{X: 3, Y: 6}
		originalGuard := &Guard{
			co:  graph.Co{X: 4, Y: 6},
			dir: 0,
		}

		floor.resetFloor(obstacle, *originalGuard)
		assert.Equal(t, originalGuard, floor.guard)
		assert.Equal(t, graph.Graph{
			{X: 4, Y: 0}: "#",
			{X: 9, Y: 1}: "#",
			{X: 2, Y: 3}: "#",
			{X: 7, Y: 4}: "#",
			{X: 1, Y: 6}: "#",
			{X: 8, Y: 7}: "#",
			{X: 0, Y: 8}: "#",
			{X: 6, Y: 9}: "#",
		}, floor.grid.Graph)
		assert.Equal(t, map[graph.Co]map[int]bool{
			{X: 4, Y: 6}: {
				0: true,
			},
		}, floor.visitedSteps)
	})
}

func Test_findSolutions(t *testing.T) {
	t.Run("returns solutions for a given input", func(t *testing.T) {
		input := []string{
			"....#.....",
			".........#",
			"..........",
			"..#.......",
			".......#..",
			"..........",
			".#..^.....",
			"........#.",
			"#.........",
			"......#...",
		}
		part1, part2 := findSolutions(input)
		assert.Equal(t, 41, part1)
		assert.Equal(t, 6, part2)
	})
}
