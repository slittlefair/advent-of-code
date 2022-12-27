package chamber_test

import (
	"Advent-of-Code/2022/Day17/chamber"
	"Advent-of-Code/2022/Day17/rock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChamber(t *testing.T) {
	t.Run("creates an empty chamber", func(t *testing.T) {
		want := chamber.Chamber{
			{X: 0, Y: 0}: true,
			{X: 1, Y: 0}: true,
			{X: 2, Y: 0}: true,
			{X: 3, Y: 0}: true,
			{X: 4, Y: 0}: true,
			{X: 5, Y: 0}: true,
			{X: 6, Y: 0}: true,
			{X: 7, Y: 0}: true,
			{X: 8, Y: 0}: true,
			{X: 0, Y: 1}: true,
			{X: 0, Y: 2}: true,
			{X: 0, Y: 3}: true,
			{X: 0, Y: 4}: true,
			{X: 8, Y: 1}: true,
			{X: 8, Y: 2}: true,
			{X: 8, Y: 3}: true,
			{X: 8, Y: 4}: true,
		}
		got := chamber.CreateChamber()
		assert.Equal(t, want, got)
	})
}

func TestExtendWalls(t *testing.T) {
	t.Run("extends walls by four from the given number", func(t *testing.T) {
		c := chamber.Chamber{
			{X: 0, Y: 0}: true,
			{X: 1, Y: 0}: true,
			{X: 2, Y: 0}: true,
			{X: 3, Y: 0}: true,
			{X: 4, Y: 0}: true,
			{X: 5, Y: 0}: true,
			{X: 6, Y: 0}: true,
			{X: 7, Y: 0}: true,
			{X: 8, Y: 0}: true,
			{X: 0, Y: 1}: true,
			{X: 0, Y: 2}: true,
			{X: 0, Y: 3}: true,
			{X: 0, Y: 4}: true,
			{X: 8, Y: 1}: true,
			{X: 8, Y: 2}: true,
			{X: 8, Y: 3}: true,
			{X: 8, Y: 4}: true,
		}
		want := chamber.Chamber{
			{X: 0, Y: 0}: true,
			{X: 1, Y: 0}: true,
			{X: 2, Y: 0}: true,
			{X: 3, Y: 0}: true,
			{X: 4, Y: 0}: true,
			{X: 5, Y: 0}: true,
			{X: 6, Y: 0}: true,
			{X: 7, Y: 0}: true,
			{X: 8, Y: 0}: true,
			{X: 0, Y: 1}: true,
			{X: 0, Y: 2}: true,
			{X: 0, Y: 3}: true,
			{X: 0, Y: 4}: true,
			{X: 8, Y: 1}: true,
			{X: 8, Y: 2}: true,
			{X: 8, Y: 3}: true,
			{X: 8, Y: 4}: true,
			{X: 8, Y: 5}: true,
			{X: 8, Y: 6}: true,
			{X: 8, Y: 7}: true,
			{X: 8, Y: 8}: true,
			{X: 0, Y: 5}: true,
			{X: 0, Y: 6}: true,
			{X: 0, Y: 7}: true,
			{X: 0, Y: 8}: true,
		}
		c.ExtendWalls(5)
		assert.Equal(t, want, c)
	})
}

func TestHighestPoint(t *testing.T) {
	t.Run("returns 0 for an empty chamber", func(t *testing.T) {
		c := chamber.Chamber{}
		got := c.HighestPoint()
		assert.Equal(t, 0, got)
	})

	t.Run("returns the greatest y value in chamber coordinates where x is not 0 or 8", func(t *testing.T) {
		c := chamber.Chamber{
			{X: 7, Y: 0}:  true,
			{X: 8, Y: 20}: true,
			{X: 2, Y: 1}:  true,
			{X: 5, Y: 12}: true,
			{X: 0, Y: 3}:  true,
			{X: 0, Y: 54}: true,
			{X: 8, Y: 61}: true,
			{X: 4, Y: 52}: true,
			{X: 2, Y: 19}: true,
			{X: 1, Y: 12}: true,
			{X: 3, Y: 49}: true,
			{X: 5, Y: 56}: true,
			{X: 8, Y: 33}: true,
			{X: 8, Y: 28}: true,
		}
		got := c.HighestPoint()
		assert.Equal(t, 56, got)
	})
}

func TestMove(t *testing.T) {
	t.Run("does not move a piece if moving it right would move it into an existing coordinate", func(t *testing.T) {
		c := chamber.Chamber{
			{X: 0, Y: 0}: true,
			{X: 1, Y: 0}: true,
			{X: 2, Y: 0}: true,
			{X: 3, Y: 0}: true,
			{X: 4, Y: 0}: true,
			{X: 5, Y: 0}: true,
			{X: 6, Y: 0}: true,
			{X: 7, Y: 0}: true,
			{X: 8, Y: 0}: true,
			{X: 0, Y: 1}: true,
			{X: 0, Y: 2}: true,
			{X: 0, Y: 3}: true,
			{X: 0, Y: 4}: true,
			{X: 8, Y: 1}: true,
			{X: 8, Y: 2}: true,
			{X: 8, Y: 3}: true,
			{X: 8, Y: 4}: true,
			{X: 4, Y: 4}: true,
			{X: 5, Y: 4}: true,
			{X: 6, Y: 4}: true,
			{X: 7, Y: 4}: true,
		}
		r := rock.Rock{
			{X: 4, Y: 4},
			{X: 5, Y: 4},
			{X: 6, Y: 4},
			{X: 7, Y: 4},
		}
		want := chamber.Chamber{
			{X: 0, Y: 0}: true,
			{X: 1, Y: 0}: true,
			{X: 2, Y: 0}: true,
			{X: 3, Y: 0}: true,
			{X: 4, Y: 0}: true,
			{X: 5, Y: 0}: true,
			{X: 6, Y: 0}: true,
			{X: 7, Y: 0}: true,
			{X: 8, Y: 0}: true,
			{X: 0, Y: 1}: true,
			{X: 0, Y: 2}: true,
			{X: 0, Y: 3}: true,
			{X: 0, Y: 4}: true,
			{X: 8, Y: 1}: true,
			{X: 8, Y: 2}: true,
			{X: 8, Y: 3}: true,
			{X: 8, Y: 4}: true,
			{X: 4, Y: 4}: true,
			{X: 5, Y: 4}: true,
			{X: 6, Y: 4}: true,
			{X: 7, Y: 4}: true,
		}
		want1 := rock.Rock{
			{X: 4, Y: 4},
			{X: 5, Y: 4},
			{X: 6, Y: 4},
			{X: 7, Y: 4},
		}
		want2 := false
		got, got1 := c.Move(r, 1, 0)
		assert.Equal(t, want, c)
		assert.Equal(t, want1, got)
		assert.Equal(t, want2, got1)
	})

	t.Run("does not move a piece if moving it left would move it into an existing coordinate", func(t *testing.T) {
		c := chamber.Chamber{
			{X: 0, Y: 0}: true,
			{X: 1, Y: 0}: true,
			{X: 2, Y: 0}: true,
			{X: 3, Y: 0}: true,
			{X: 4, Y: 0}: true,
			{X: 5, Y: 0}: true,
			{X: 6, Y: 0}: true,
			{X: 7, Y: 0}: true,
			{X: 8, Y: 0}: true,
			{X: 0, Y: 1}: true,
			{X: 0, Y: 2}: true,
			{X: 0, Y: 3}: true,
			{X: 0, Y: 4}: true,
			{X: 8, Y: 1}: true,
			{X: 8, Y: 2}: true,
			{X: 8, Y: 3}: true,
			{X: 8, Y: 4}: true,
			{X: 1, Y: 4}: true,
			{X: 2, Y: 4}: true,
			{X: 3, Y: 4}: true,
			{X: 3, Y: 5}: true,
			{X: 3, Y: 6}: true,
		}
		r := rock.Rock{
			{X: 1, Y: 4},
			{X: 2, Y: 4},
			{X: 3, Y: 4},
			{X: 3, Y: 5},
			{X: 3, Y: 6},
		}
		want := chamber.Chamber{
			{X: 0, Y: 0}: true,
			{X: 1, Y: 0}: true,
			{X: 2, Y: 0}: true,
			{X: 3, Y: 0}: true,
			{X: 4, Y: 0}: true,
			{X: 5, Y: 0}: true,
			{X: 6, Y: 0}: true,
			{X: 7, Y: 0}: true,
			{X: 8, Y: 0}: true,
			{X: 0, Y: 1}: true,
			{X: 0, Y: 2}: true,
			{X: 0, Y: 3}: true,
			{X: 0, Y: 4}: true,
			{X: 8, Y: 1}: true,
			{X: 8, Y: 2}: true,
			{X: 8, Y: 3}: true,
			{X: 8, Y: 4}: true,
			{X: 1, Y: 4}: true,
			{X: 2, Y: 4}: true,
			{X: 3, Y: 4}: true,
			{X: 3, Y: 5}: true,
			{X: 3, Y: 6}: true,
		}
		want1 := rock.Rock{
			{X: 1, Y: 4},
			{X: 2, Y: 4},
			{X: 3, Y: 4},
			{X: 3, Y: 5},
			{X: 3, Y: 6},
		}
		want2 := false
		got, got1 := c.Move(r, -1, 0)
		assert.Equal(t, want, c)
		assert.Equal(t, want1, got)
		assert.Equal(t, want2, got1)
	})

	t.Run("does not move a piece if moving it down would move it into an existing coordinate", func(t *testing.T) {
		c := chamber.Chamber{
			{X: 0, Y: 0}: true,
			{X: 1, Y: 0}: true,
			{X: 2, Y: 0}: true,
			{X: 3, Y: 0}: true,
			{X: 4, Y: 0}: true,
			{X: 5, Y: 0}: true,
			{X: 6, Y: 0}: true,
			{X: 7, Y: 0}: true,
			{X: 8, Y: 0}: true,
			{X: 0, Y: 1}: true,
			{X: 0, Y: 2}: true,
			{X: 0, Y: 3}: true,
			{X: 0, Y: 4}: true,
			{X: 8, Y: 1}: true,
			{X: 8, Y: 2}: true,
			{X: 8, Y: 3}: true,
			{X: 8, Y: 4}: true,
			{X: 3, Y: 1}: true,
			{X: 4, Y: 1}: true,
			{X: 3, Y: 2}: true,
			{X: 4, Y: 2}: true,
		}
		r := rock.Rock{
			{X: 3, Y: 1},
			{X: 4, Y: 1},
			{X: 3, Y: 2},
			{X: 4, Y: 2},
		}
		want := chamber.Chamber{
			{X: 0, Y: 0}: true,
			{X: 1, Y: 0}: true,
			{X: 2, Y: 0}: true,
			{X: 3, Y: 0}: true,
			{X: 4, Y: 0}: true,
			{X: 5, Y: 0}: true,
			{X: 6, Y: 0}: true,
			{X: 7, Y: 0}: true,
			{X: 8, Y: 0}: true,
			{X: 0, Y: 1}: true,
			{X: 0, Y: 2}: true,
			{X: 0, Y: 3}: true,
			{X: 0, Y: 4}: true,
			{X: 8, Y: 1}: true,
			{X: 8, Y: 2}: true,
			{X: 8, Y: 3}: true,
			{X: 8, Y: 4}: true,
			{X: 3, Y: 1}: true,
			{X: 4, Y: 1}: true,
			{X: 3, Y: 2}: true,
			{X: 4, Y: 2}: true,
		}
		want1 := rock.Rock{
			{X: 3, Y: 1},
			{X: 4, Y: 1},
			{X: 3, Y: 2},
			{X: 4, Y: 2},
		}
		want2 := false
		got, got1 := c.Move(r, 0, -1)
		assert.Equal(t, want, c)
		assert.Equal(t, want1, got)
		assert.Equal(t, want2, got1)
	})

	t.Run("moves a piece right", func(t *testing.T) {
		c := chamber.Chamber{
			{X: 0, Y: 0}: true,
			{X: 1, Y: 0}: true,
			{X: 2, Y: 0}: true,
			{X: 3, Y: 0}: true,
			{X: 4, Y: 0}: true,
			{X: 5, Y: 0}: true,
			{X: 6, Y: 0}: true,
			{X: 7, Y: 0}: true,
			{X: 8, Y: 0}: true,
			{X: 0, Y: 1}: true,
			{X: 0, Y: 2}: true,
			{X: 0, Y: 3}: true,
			{X: 0, Y: 4}: true,
			{X: 8, Y: 1}: true,
			{X: 8, Y: 2}: true,
			{X: 8, Y: 3}: true,
			{X: 8, Y: 4}: true,
			{X: 1, Y: 4}: true,
			{X: 2, Y: 4}: true,
			{X: 3, Y: 4}: true,
			{X: 4, Y: 4}: true,
		}
		r := rock.Rock{
			{X: 1, Y: 4},
			{X: 2, Y: 4},
			{X: 3, Y: 4},
			{X: 4, Y: 4},
		}
		want := chamber.Chamber{
			{X: 0, Y: 0}: true,
			{X: 1, Y: 0}: true,
			{X: 2, Y: 0}: true,
			{X: 3, Y: 0}: true,
			{X: 4, Y: 0}: true,
			{X: 5, Y: 0}: true,
			{X: 6, Y: 0}: true,
			{X: 7, Y: 0}: true,
			{X: 8, Y: 0}: true,
			{X: 0, Y: 1}: true,
			{X: 0, Y: 2}: true,
			{X: 0, Y: 3}: true,
			{X: 0, Y: 4}: true,
			{X: 8, Y: 1}: true,
			{X: 8, Y: 2}: true,
			{X: 8, Y: 3}: true,
			{X: 8, Y: 4}: true,
			{X: 2, Y: 4}: true,
			{X: 3, Y: 4}: true,
			{X: 4, Y: 4}: true,
			{X: 5, Y: 4}: true,
		}
		want1 := rock.Rock{
			{X: 2, Y: 4},
			{X: 3, Y: 4},
			{X: 4, Y: 4},
			{X: 5, Y: 4},
		}
		want2 := true
		got, got1 := c.Move(r, 1, 0)
		assert.Equal(t, want, c)
		assert.Equal(t, want1, got)
		assert.Equal(t, want2, got1)
	})

	t.Run("moves a piece left", func(t *testing.T) {
		c := chamber.Chamber{
			{X: 0, Y: 0}: true,
			{X: 1, Y: 0}: true,
			{X: 2, Y: 0}: true,
			{X: 3, Y: 0}: true,
			{X: 4, Y: 0}: true,
			{X: 5, Y: 0}: true,
			{X: 6, Y: 0}: true,
			{X: 7, Y: 0}: true,
			{X: 8, Y: 0}: true,
			{X: 0, Y: 1}: true,
			{X: 0, Y: 2}: true,
			{X: 0, Y: 3}: true,
			{X: 0, Y: 4}: true,
			{X: 8, Y: 1}: true,
			{X: 8, Y: 2}: true,
			{X: 8, Y: 3}: true,
			{X: 8, Y: 4}: true,
			{X: 5, Y: 4}: true,
			{X: 6, Y: 4}: true,
			{X: 7, Y: 4}: true,
			{X: 7, Y: 5}: true,
			{X: 7, Y: 6}: true,
		}
		r := rock.Rock{
			{X: 5, Y: 4},
			{X: 6, Y: 4},
			{X: 7, Y: 4},
			{X: 7, Y: 5},
			{X: 7, Y: 6},
		}
		want := chamber.Chamber{
			{X: 0, Y: 0}: true,
			{X: 1, Y: 0}: true,
			{X: 2, Y: 0}: true,
			{X: 3, Y: 0}: true,
			{X: 4, Y: 0}: true,
			{X: 5, Y: 0}: true,
			{X: 6, Y: 0}: true,
			{X: 7, Y: 0}: true,
			{X: 8, Y: 0}: true,
			{X: 0, Y: 1}: true,
			{X: 0, Y: 2}: true,
			{X: 0, Y: 3}: true,
			{X: 0, Y: 4}: true,
			{X: 8, Y: 1}: true,
			{X: 8, Y: 2}: true,
			{X: 8, Y: 3}: true,
			{X: 8, Y: 4}: true,
			{X: 4, Y: 4}: true,
			{X: 5, Y: 4}: true,
			{X: 6, Y: 4}: true,
			{X: 6, Y: 5}: true,
			{X: 6, Y: 6}: true,
		}
		want1 := rock.Rock{
			{X: 4, Y: 4},
			{X: 5, Y: 4},
			{X: 6, Y: 4},
			{X: 6, Y: 5},
			{X: 6, Y: 6},
		}
		want2 := true
		got, got1 := c.Move(r, -1, 0)
		assert.Equal(t, want, c)
		assert.Equal(t, want1, got)
		assert.Equal(t, want2, got1)
	})

	t.Run("moves a piece down", func(t *testing.T) {
		c := chamber.Chamber{
			{X: 0, Y: 0}: true,
			{X: 1, Y: 0}: true,
			{X: 2, Y: 0}: true,
			{X: 3, Y: 0}: true,
			{X: 4, Y: 0}: true,
			{X: 5, Y: 0}: true,
			{X: 6, Y: 0}: true,
			{X: 7, Y: 0}: true,
			{X: 8, Y: 0}: true,
			{X: 0, Y: 1}: true,
			{X: 0, Y: 2}: true,
			{X: 0, Y: 3}: true,
			{X: 0, Y: 4}: true,
			{X: 8, Y: 1}: true,
			{X: 8, Y: 2}: true,
			{X: 8, Y: 3}: true,
			{X: 8, Y: 4}: true,
			{X: 3, Y: 2}: true,
			{X: 4, Y: 2}: true,
			{X: 3, Y: 3}: true,
			{X: 4, Y: 3}: true,
		}
		r := rock.Rock{
			{X: 3, Y: 2},
			{X: 4, Y: 2},
			{X: 3, Y: 3},
			{X: 4, Y: 3},
		}
		want := chamber.Chamber{
			{X: 0, Y: 0}: true,
			{X: 1, Y: 0}: true,
			{X: 2, Y: 0}: true,
			{X: 3, Y: 0}: true,
			{X: 4, Y: 0}: true,
			{X: 5, Y: 0}: true,
			{X: 6, Y: 0}: true,
			{X: 7, Y: 0}: true,
			{X: 8, Y: 0}: true,
			{X: 0, Y: 1}: true,
			{X: 0, Y: 2}: true,
			{X: 0, Y: 3}: true,
			{X: 0, Y: 4}: true,
			{X: 8, Y: 1}: true,
			{X: 8, Y: 2}: true,
			{X: 8, Y: 3}: true,
			{X: 8, Y: 4}: true,
			{X: 3, Y: 1}: true,
			{X: 4, Y: 1}: true,
			{X: 3, Y: 2}: true,
			{X: 4, Y: 2}: true,
		}
		want1 := rock.Rock{
			{X: 3, Y: 1},
			{X: 4, Y: 1},
			{X: 3, Y: 2},
			{X: 4, Y: 2},
		}
		want2 := true
		got, got1 := c.Move(r, 0, -1)
		assert.Equal(t, want, c)
		assert.Equal(t, want1, got)
		assert.Equal(t, want2, got1)
	})
}

func TestGetColumnHeightsAndMinimiseChamber(t *testing.T) {
	t.Run("returns the heights of each column and removes excess coordinates", func(t *testing.T) {
		c := chamber.Chamber{
			{X: 0, Y: 0}:    true,
			{X: 1, Y: 0}:    true,
			{X: 2, Y: 0}:    true,
			{X: 3, Y: 0}:    true,
			{X: 4, Y: 0}:    true,
			{X: 5, Y: 0}:    true,
			{X: 6, Y: 0}:    true,
			{X: 7, Y: 0}:    true,
			{X: 8, Y: 0}:    true,
			{X: 0, Y: 1}:    true,
			{X: 0, Y: 2}:    true,
			{X: 0, Y: 3}:    true,
			{X: 0, Y: 4}:    true,
			{X: 8, Y: 1}:    true,
			{X: 8, Y: 2}:    true,
			{X: 8, Y: 3}:    true,
			{X: 8, Y: 4}:    true,
			{X: 1, Y: 1}:    true,
			{X: 1, Y: 2}:    true,
			{X: 1, Y: 3}:    true,
			{X: 2, Y: 5}:    true,
			{X: 2, Y: 4}:    true,
			{X: 3, Y: 5}:    true,
			{X: 3, Y: 15}:   true,
			{X: 3, Y: 25}:   true,
			{X: 3, Y: 35}:   true,
			{X: 4, Y: 85}:   true,
			{X: 5, Y: 2}:    true,
			{X: 5, Y: 22}:   true,
			{X: 6, Y: 8}:    true,
			{X: 6, Y: 11}:   true,
			{X: 6, Y: 10}:   true,
			{X: 6, Y: 9}:    true,
			{X: 7, Y: 4655}: true,
			{X: 7, Y: 8742}: true,
			{X: 7, Y: 1111}: true,
			{X: 7, Y: 1873}: true,
			{X: 7, Y: 3799}: true,
		}
		want := chamber.Chamber{
			{X: 0, Y: 3}:    true,
			{X: 0, Y: 4}:    true,
			{X: 8, Y: 3}:    true,
			{X: 8, Y: 4}:    true,
			{X: 1, Y: 3}:    true,
			{X: 2, Y: 5}:    true,
			{X: 2, Y: 4}:    true,
			{X: 3, Y: 5}:    true,
			{X: 3, Y: 15}:   true,
			{X: 3, Y: 25}:   true,
			{X: 3, Y: 35}:   true,
			{X: 4, Y: 85}:   true,
			{X: 5, Y: 22}:   true,
			{X: 6, Y: 8}:    true,
			{X: 6, Y: 11}:   true,
			{X: 6, Y: 10}:   true,
			{X: 6, Y: 9}:    true,
			{X: 7, Y: 4655}: true,
			{X: 7, Y: 8742}: true,
			{X: 7, Y: 1111}: true,
			{X: 7, Y: 1873}: true,
			{X: 7, Y: 3799}: true,
		}
		want1 := [7]int{0, 2, 32, 82, 19, 8, 8739}
		got := c.GetColumnHeightsAndMinimiseChamber()
		assert.Equal(t, want, c)
		assert.Equal(t, want1, got)
	})
}
