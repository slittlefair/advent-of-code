package chamber

import (
	"Advent-of-Code/2022/Day17/rock"
	"Advent-of-Code/graph"
	"Advent-of-Code/maths"
	"Advent-of-Code/slice"
)

type Chamber map[graph.Co]bool

// CreateChamber creates the initial, empty chamber. Side walls are extended as
// and when.
func CreateChamber() Chamber {
	chamber := Chamber{}
	for x := 0; x <= 8; x++ {
		chamber[graph.Co{X: x, Y: 0}] = true
	}
	for y := 1; y <= 4; y++ {
		chamber[graph.Co{X: 0, Y: y}] = true
		chamber[graph.Co{X: 8, Y: y}] = true
	}
	return chamber
}

// ExtendWalls increases the side walls by 4 (the tallest rock) from the given y value
func (c Chamber) ExtendWalls(n int) {
	for y := n; y <= n+3; y++ {
		c[graph.Co{X: 0, Y: y}] = true
		c[graph.Co{X: 8, Y: y}] = true
	}
}

// HighestPoint returns the highest rock that isn't part of the side walls.
func (c Chamber) HighestPoint() int {
	maximum := 0
	for co := range c {
		if co.X > 0 && co.X < 8 {
			maximum = maths.Max(maximum, co.Y)
		}
	}
	return maximum
}

// Move alters a rock piece position the given x and y amounts. The chamber is updated with these
// positions. The rock piece's new coordinates are returned as well as whether the piece was able to
// be moved.
func (c Chamber) Move(r rock.Rock, x int, y int) (rock.Rock, bool) {
	movedRock := rock.Rock{}
	for _, co := range r {
		movedRock = append(movedRock, graph.Co{X: co.X + x, Y: co.Y + y})
	}
	for _, mCo := range movedRock {
		for co := range c {
			// If the original rock doesn't contain the moved coordinate but matches a coordinate
			// already in the chamber then it can't move, so return the original rock piece position
			// and false.
			if !slice.Contains(r, mCo) && co == mCo {
				return r, false
			}
		}
	}
	// Delete the old rock from the chamber and ad the new one
	for _, co := range r {
		delete(c, co)
	}
	for _, co := range movedRock {
		c[co] = true
	}
	return movedRock, true
}

// GetColumnHeightsAndMinimiseChamber returns an array of the column heights of the fallen rock
// pieces. This is after we minimise the chamber by removing all rocks below the uppermost rocks in
// each column, since these can no longer be interacted with by new pieces, which improves
// performance significantly.
func (c Chamber) GetColumnHeightsAndMinimiseChamber() [7]int {
	hp := [7]int{}
	for co := range c {
		if co.X > 0 && co.X < 8 {
			hp[co.X-1] = maths.Max(hp[co.X-1], co.Y)
		}
	}
	lowestHP := maths.Infinity
	for _, v := range hp {
		lowestHP = maths.Min(lowestHP, v)
	}
	for i, v := range hp {
		hp[i] = v - lowestHP
	}
	for co := range c {
		if co.Y < lowestHP {
			delete(c, co)
		}
	}
	return hp
}
