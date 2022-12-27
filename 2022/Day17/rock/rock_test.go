package rock_test

import (
	"Advent-of-Code/2022/Day17/rock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDash(t *testing.T) {
	t.Run("returns a rock with correct coordinates if passed 0", func(t *testing.T) {
		want := rock.Rock{
			{X: 3, Y: 4},
			{X: 4, Y: 4},
			{X: 5, Y: 4},
			{X: 6, Y: 4},
		}
		got := rock.Dash(0)
		assert.Equal(t, want, got)
	})

	t.Run("returns a rock with correct coordinates if passed non zero value", func(t *testing.T) {
		want := rock.Rock{
			{X: 3, Y: 9},
			{X: 4, Y: 9},
			{X: 5, Y: 9},
			{X: 6, Y: 9},
		}
		got := rock.Dash(5)
		assert.Equal(t, want, got)
	})
}

func TestCross(t *testing.T) {
	t.Run("returns a rock with correct coordinates if passed 0", func(t *testing.T) {
		want := rock.Rock{
			{X: 3, Y: 5},
			{X: 4, Y: 4},
			{X: 4, Y: 5},
			{X: 4, Y: 6},
			{X: 5, Y: 5},
		}
		got := rock.Cross(0)
		assert.Equal(t, want, got)
	})

	t.Run("returns a rock with correct coordinates if passed non zero value", func(t *testing.T) {
		want := rock.Rock{
			{X: 3, Y: 14},
			{X: 4, Y: 13},
			{X: 4, Y: 14},
			{X: 4, Y: 15},
			{X: 5, Y: 14},
		}
		got := rock.Cross(9)
		assert.Equal(t, want, got)
	})
}

func TestL(t *testing.T) {
	t.Run("returns a rock with correct coordinates if passed 0", func(t *testing.T) {
		want := rock.Rock{
			{X: 3, Y: 4},
			{X: 4, Y: 4},
			{X: 5, Y: 4},
			{X: 5, Y: 5},
			{X: 5, Y: 6},
		}
		got := rock.L(0)
		assert.Equal(t, want, got)
	})

	t.Run("returns a rock with correct coordinates if passed non zero value", func(t *testing.T) {
		want := rock.Rock{
			{X: 3, Y: 6},
			{X: 4, Y: 6},
			{X: 5, Y: 6},
			{X: 5, Y: 7},
			{X: 5, Y: 8},
		}
		got := rock.L(2)
		assert.Equal(t, want, got)
	})
}

func TestI(t *testing.T) {
	t.Run("returns a rock with correct coordinates if passed 0", func(t *testing.T) {
		want := rock.Rock{
			{X: 3, Y: 4},
			{X: 3, Y: 5},
			{X: 3, Y: 6},
			{X: 3, Y: 7},
		}
		got := rock.I(0)
		assert.Equal(t, want, got)
	})

	t.Run("returns a rock with correct coordinates if passed non zero value", func(t *testing.T) {
		want := rock.Rock{
			{X: 3, Y: 41},
			{X: 3, Y: 42},
			{X: 3, Y: 43},
			{X: 3, Y: 44},
		}
		got := rock.I(37)
		assert.Equal(t, want, got)
	})
}

func TestSquare(t *testing.T) {
	t.Run("returns a rock with correct coordinates if passed 0", func(t *testing.T) {
		want := rock.Rock{
			{X: 3, Y: 4},
			{X: 4, Y: 4},
			{X: 3, Y: 5},
			{X: 4, Y: 5},
		}
		got := rock.Square(0)
		assert.Equal(t, want, got)
	})

	t.Run("returns a rock with correct coordinates if passed non zero value", func(t *testing.T) {
		want := rock.Rock{
			{X: 3, Y: 103},
			{X: 4, Y: 103},
			{X: 3, Y: 104},
			{X: 4, Y: 104},
		}
		got := rock.Square(99)
		assert.Equal(t, want, got)
	})
}
