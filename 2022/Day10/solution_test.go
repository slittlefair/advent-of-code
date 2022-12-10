package main

import (
	"Advent-of-Code/graph"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckCycle(t *testing.T) {
	t.Run("does not update the signal if the cycle is not the right iteration", func(t *testing.T) {
		cpu := &cpu{
			cycle:  125,
			x:      13,
			pixels: make(map[graph.Co]string),
		}
		cpu.checkCycle()
		assert.Equal(t, 0, cpu.signal)
	})

	testsPart1 := []struct {
		cpu  *cpu
		want int
	}{
		{
			cpu: &cpu{
				cycle: 19,
				x:     21,
			},
			want: 420,
		},
		{
			cpu: &cpu{
				cycle: 59,
				x:     19,
			},
			want: 1140,
		},
		{
			cpu: &cpu{
				cycle: 99,
				x:     18,
			},
			want: 1800,
		},
		{
			cpu: &cpu{
				cycle: 139,
				x:     21,
			},
			want: 2940,
		},
		{
			cpu: &cpu{
				cycle: 179,
				x:     16,
			},
			want: 2880,
		},
		{
			cpu: &cpu{
				cycle: 219,
				x:     18,
			},
			want: 3960,
		},
	}
	for i, tt := range testsPart1 {
		t.Run(fmt.Sprintf("increases the cpu signal by the correct amount at the right time, advent of code example %d", i+1), func(t *testing.T) {
			tt.cpu.pixels = make(map[graph.Co]string)
			tt.cpu.checkCycle()
			assert.Equal(t, tt.want, tt.cpu.signal)
		})
	}

	testsPart2 := []struct {
		name string
		co   graph.Co
		x    int
		want string
	}{
		{
			name: "sets the current pixel to empty space if the x value is greater than the sprite",
			co:   graph.Co{X: 13, Y: 0},
			x:    6,
			want: " ",
		},
		{
			name: "sets the current pixel to empty space if the x value is less than the sprite",
			co:   graph.Co{X: 13, Y: 0},
			x:    36,
			want: " ",
		},
		{
			name: "sets the current pixel to a block if the x value is on the left of the sprite",
			co:   graph.Co{X: 10, Y: 0},
			x:    11,
			want: "\u2588",
		},
		{
			name: "sets the current pixel to a block if the x value is on the right of the sprite",
			co:   graph.Co{X: 12, Y: 0},
			x:    11,
			want: "\u2588",
		},
		{
			name: "sets the current pixel to a block if the x value is in the centre of the sprite",
			co:   graph.Co{X: 10, Y: 0},
			x:    10,
			want: "\u2588",
		},
	}

	for _, tt := range testsPart2 {
		t.Run(tt.name, func(t *testing.T) {
			cpu := &cpu{
				pixels: make(map[graph.Co]string),
				x:      tt.x,
				co:     tt.co,
			}
			cpu.checkCycle()
			assert.Equal(t, tt.want, cpu.pixels[tt.co])
		})
	}

	t.Run("it moves the sprite along one space if it has not reached the end of the current row", func(t *testing.T) {
		cpu := &cpu{
			pixels: make(map[graph.Co]string),
			x:      26,
			co:     graph.Co{X: 17, Y: 2},
		}
		cpu.checkCycle()
		assert.Equal(t, graph.Co{X: 18, Y: 2}, cpu.co)
	})

	t.Run("it moves the sprite onto the next row if it has reached the end of the current row", func(t *testing.T) {
		cpu := &cpu{
			pixels: make(map[graph.Co]string),
			x:      26,
			co:     graph.Co{X: 39, Y: 2},
		}
		cpu.checkCycle()
		assert.Equal(t, graph.Co{X: 0, Y: 3}, cpu.co)
	})
}

func TestHandleInstruction(t *testing.T) {
	t.Run("returns an error if the instruction cannot be scanned", func(t *testing.T) {
		cpu := &cpu{}
		err := cpu.handleInstruction("add3")
		assert.Error(t, err)
	})

	t.Run("runs a noop instruction", func(t *testing.T) {
		c := &cpu{
			x:      1,
			pixels: make(map[graph.Co]string),
		}
		want := &cpu{
			x:     1,
			cycle: 1,
			co:    graph.Co{X: 1, Y: 0},
			pixels: map[graph.Co]string{
				{X: 0, Y: 0}: "\u2588",
			},
		}
		err := c.handleInstruction("noop")
		assert.NoError(t, err)
		assert.Equal(t, want, c)
	})

	t.Run("runs a move instruction", func(t *testing.T) {
		c := &cpu{
			x:     1,
			cycle: 1,
			co:    graph.Co{X: 1, Y: 0},
			pixels: map[graph.Co]string{
				{X: 0, Y: 0}: "\u2588",
			},
		}
		want := &cpu{
			cycle: 3,
			x:     4,
			co:    graph.Co{X: 3, Y: 0},
			pixels: map[graph.Co]string{
				{X: 0, Y: 0}: "\u2588",
				{X: 1, Y: 0}: "\u2588",
				{X: 2, Y: 0}: "\u2588",
			},
		}
		err := c.handleInstruction("addx 3")
		assert.NoError(t, err)
		assert.Equal(t, want, c)
	})
}

func TestCompleteCycles(t *testing.T) {
	t.Run("returns an error if an instruction cannot be parsed", func(t *testing.T) {
		c := &cpu{
			x:      1,
			pixels: make(map[graph.Co]string),
		}
		input := []string{
			"noop",
			"addx 3",
			"addx-5",
		}
		err := c.completeCycles(input)
		assert.Error(t, err)
	})

	t.Run("completes a simple set of instructions", func(t *testing.T) {
		c := &cpu{
			x:      1,
			pixels: make(map[graph.Co]string),
		}
		input := []string{
			"noop",
			"addx 3",
			"addx -5",
		}
		want := &cpu{
			x:     -1,
			cycle: 5,
			co:    graph.Co{X: 5, Y: 0},
			pixels: map[graph.Co]string{
				{X: 0, Y: 0}: "\u2588",
				{X: 1, Y: 0}: "\u2588",
				{X: 2, Y: 0}: "\u2588",
				{X: 3, Y: 0}: "\u2588",
				{X: 4, Y: 0}: "\u2588",
			},
		}
		err := c.completeCycles(input)
		assert.NoError(t, err)
		assert.Equal(t, want, c)
	})
}
